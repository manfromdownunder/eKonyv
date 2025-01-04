//go:build windows && webview

package desktop

import "github.com/inkeliz/gowebview"

func StartWebView(url string, debug bool) {
	config := &gowebview.Config{
		Debug: debug,
		URL:   url,
		WindowConfig: &gowebview.WindowConfig{
			Title:      "eKonyv",
			Visibility: gowebview.VisibilityDefault,
			Size: &gowebview.Point{
				X: 1200,
				Y: 800,
			},
		},
	}

	w, err := gowebview.New(config)
	if err != nil {
		panic(err)
	}

	defer w.Destroy()
	w.Run()
}
