package server

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/manfromdownunder/eKonyv/irc" // Import the IRC package

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

//go:embed app/dist
var reactClient embed.FS

// Store IRC connections mapped to usernames
var ircConnections = make(map[string]*irc.Conn) // Map for tracking IRC connections per username

func (server *server) registerRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Handle("/*", server.staticFilesHandler("app/dist"))
	router.Get("/ws", server.serveWs())
	router.Get("/stats", server.statsHandler())
	router.Get("/servers", server.serverListHandler())

	router.Group(func(r chi.Router) {
		r.Use(server.requireUser)
		r.Get("/library", server.getAllBooksHandler())
		r.Delete("/library/{fileName}", server.deleteBooksHandler())
		r.Get("/library/*", server.getBookHandler())
	})

	return router
}

// serveWs handles websocket requests from the peer.
func (server *server) serveWs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("OpenBooks")
		if errors.Is(err, http.ErrNoCookie) {
			cookie = &http.Cookie{
				Name:     "OpenBooks",
				Value:    uuid.New().String(),
				Secure:   false,
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour * 24 * 7),
				SameSite: http.SameSiteStrictMode,
			}
			w.Header().Add("Set-Cookie", cookie.String())
		}

		userId, err := uuid.Parse(cookie.Value)
		_, alreadyConnected := server.clients[userId]

		// If invalid UUID or the same browser tries to connect again or multiple browser connections
		// Don't connect to IRC or create new client
		if err != nil || alreadyConnected || len(server.clients) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		upgrader.CheckOrigin = func(req *http.Request) bool {
			return true
		}

		conn, err := upgrader.Upgrade(w, r, w.Header())
		if err != nil {
			server.log.Println(err)
			return
		}

		// Get or create IRC connection for the username
		username := fmt.Sprintf("user-%s", userId.String()) // Maintain the same username across connections

		// Check if there's already an existing IRC connection for the username
		var ircConn *irc.Conn
		if existingConn, exists := ircConnections[username]; exists {
			ircConn = existingConn
		} else {
			// If no existing connection, create a new IRC connection
			ircConn = irc.New(username, server.config.UserAgent)
			err = ircConn.Connect(server.config.Server, true) // true for TLS
			if err != nil {
				server.log.Println("Error connecting to IRC:", err)
				return
			}

			// Store the IRC connection in the map
			ircConnections[username] = ircConn
		}

		client := &Client{
			conn: conn,
			send: make(chan interface{}, 128),
			uuid: userId,
			irc:  ircConn, // Use the existing or newly created IRC connection
			log:  log.New(os.Stdout, fmt.Sprintf("CLIENT (%s): ", server.config.UserName), log.LstdFlags|log.Lmsgprefix),
			ctx:  context.Background(),
		}

		server.log.Printf("Client connected from %s\n", conn.RemoteAddr().String())
		client.log.Println("New client created.")

		// Register the client in the server
		server.register <- client

		go server.writePump(client)
		go server.readPump(client)
	}
}

func (server *server) staticFilesHandler(assetPath string) http.Handler {
	// Update the embedded file system's tree so that index.html is at the root
	app, err := fs.Sub(reactClient, assetPath)
	if err != nil {
		server.log.Println(err)
	}

	// Strip the predefined base path and serve the static file
	return http.StripPrefix(server.config.Basepath, http.FileServer(http.FS(app)))
}

func (server *server) statsHandler() http.HandlerFunc {
	type statsResponse struct {
		UUID string `json:"uuid"`
		IP   string `json:"ip"`
		Name string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		result := make([]statsResponse, 0, len(server.clients))

		// Gather the stats for all connected clients
		for _, client := range server.clients {
			details := statsResponse{
				UUID: client.uuid.String(),
				Name: client.irc.Username,
				IP:   client.conn.RemoteAddr().String(),
			}

			result = append(result, details)
		}

		// Send the client stats as a JSON response
		json.NewEncoder(w).Encode(result)
	}
}

func (server *server) serverListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Send the list of IRC servers
		json.NewEncoder(w).Encode(server.repository.servers)
	}
}

func (server *server) getAllBooksHandler() http.HandlerFunc {
	type download struct {
		Name         string    `json:"name"`
		DownloadLink string    `json:"downloadLink"`
		Time         time.Time `json:"time"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if !server.config.Persist {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Fetch and list all the books in the library
		libraryDir := filepath.Join(server.config.DownloadDir, "books")
		books, err := os.ReadDir(libraryDir)
		if err != nil {
			server.log.Printf("Unable to list books. %s\n", err)
		}

		output := make([]download, 0)
		for _, book := range books {
			if book.IsDir() || strings.HasPrefix(book.Name(), ".") || filepath.Ext(book.Name()) == ".temp" {
				continue
			}

			info, err := book.Info()
			if err != nil {
				server.log.Println(err)
			}

			dl := download{
				Name:         book.Name(),
				DownloadLink: path.Join("library", book.Name()),
				Time:         info.ModTime(),
			}

			output = append(output, dl)
		}

		// Send the list of books as a JSON response
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(output)
	}
}

func (server *server) getBookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, fileName := path.Split(r.URL.Path)
		bookPath := filepath.Join(server.config.DownloadDir, "books", fileName)

		// Serve the requested book file
		http.ServeFile(w, r, bookPath)

		// Optionally remove the file if Persist is not enabled
		if !server.config.Persist {
			err := os.Remove(bookPath)
			if err != nil {
				server.log.Printf("Error when deleting book file. %s", err)
			}
		}
	}
}

func (server *server) deleteBooksHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileName, err := url.PathUnescape(chi.URLParam(r, "fileName"))
		if err != nil {
			server.log.Printf("Error unescaping path: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		// Delete the specified book from the library
		err = os.Remove(filepath.Join(server.config.DownloadDir, "books", fileName))
		if err != nil {
			server.log.Printf("Error deleting book file: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
