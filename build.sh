#!/bin/bash

echo "Building React App."
cd server/app
npm install
npm run build
cd ../../cmd/eKonyv

echo "Building binaries for various platforms."
env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../../build/eKonyv.exe
env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ../../build/eKonyv_mac
env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ../../build/eKonyv_mac_arm
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../build/eKonyv_linux
env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ../../build/eKonyv_linux_arm