# eKonyv - fork of OpenBooks

OpenBooks does not appear to be maintained any longer.  Forked the project to keep adding functionality.

> NOTE: Going forward only the latest release will be supported. If you encounter any issues, be sure you are using the latest version.

[![Docker Pulls](https://img.shields.io/docker/pulls/manfromdownunder/eKonyv.svg)](https://hub.docker.com/r/manfromdownunder/eKonyv/)

eKonyv allows you to download ebooks from irc.irchighway.net quickly and easily.

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="./.github/home_v3_dark.png">
  <img alt="eKonyv screenshot" src="./.github/home_v3.png">
</picture>


## Getting Started

### Binary

1. Download the latest release for your platform from the [releases page](https://github.com/manfromdownunder/eKonyv/releases).
2. Run the binary
   - Linux users may have to run `chmod +x [binary name]` to make it executable
3. `./eKonyv --help`
   - This will display all possible configuration values and introduce the two modes; CLI or Server.

### Docker

- Basic config
  - `docker run -p 8080:80 manfromdownunder/eKonyv`
- Config to persist all eBook files to disk
  - `docker run -p 8080:80 -v /home/user/Downloads/eKonyv:/books manfromdownunder/eKonyv --persist`

### Setting the Base Path

eKonyv server doesn't have to be hosted at the root of your webserver. The basepath value allows you to host it behind a reverse proxy. The base path value must have opening and closing forward slashes (default "/").

- Docker
  - `docker run -p 8080:80 -e BASE_PATH=/eKonyv/ manfromdownunder/eKonyv`
- Binary
  - `./eKonyv server --basepath /eKonyv/`

## Usage

For a complete list of features use the `--help` flags on all subcommands.
For example `eKonyv cli --help or eKonyv cli download --help`. There are
two modes; Server or CLI. In CLI mode you interact and download books through
a terminal interface. In server mode the application runs as a web application
that you can visit in your browser.

Double clicking the executable will open the UI in your browser. In the future it may use [webviews](https://developer.microsoft.com/en-us/microsoft-edge/webview2/) to provide a "native-like" desktop application. 

## Development

### Install the dependencies

- `go get`
- `cd server/app && npm install`
- `cd ../..`
- `go run main.go`

### Build the React SPA and compile binaries for multiple platforms.

- Run `./build.sh`
- This will install npm packages, build the React app, and compile the executable.

### Build the go binary (if you haven't changed the frontend)

- `go build`

### Mock Development Server

- The mock server allows you to debug responses and requests to simplified IRC / DCC
  servers that mimic the responses received from IRC Highway.
- ```bash
  cd cmd/mock_server
  go run .
  # Another Terminal
  cd cmd/eKonyv
  go run . server --server localhost --log
  ```

### Desktop App
Compile eKonyv with experimental webview support:

``` shell
cd cmd/eKonyv
go build -tags webview
```


## Why / How

- I wrote this as an easier way to search and download books from irchighway.net. It handles all the extraction and data processing for you. You just have to click the book you want. Hopefully you find it much easier than the IRC interface.
- It was also interesting to learn how the [IRC](https://en.wikipedia.org/wiki/Internet_Relay_Chat) and [DCC](https://en.wikipedia.org/wiki/Direct_Client-to-Client) protocols work and write custom implementations.

## Technology

- Backend
  - Golang
  - Chi
  - gorilla/websocket
  - Archiver (extract files from various archive formats)
- Frontend
  - React.js
  - TypeScript
  - Redux / Redux Toolkit
  - Mantine UI / @emotion/react
  - Framer Motion
