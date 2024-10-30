#!/bin/bash

# MacOS Intel 64位
GOOS=darwin GOARCH=amd64 go build -o kss-cli-mac-amd64 main.go

# MacOS Apple Silicon (M1/M2)
GOOS=darwin GOARCH=arm64 go build -o kss-cli-mac-arm64 main.go

# Windows 64位
GOOS=windows GOARCH=amd64 go build -o kss-cli-windows-amd64.exe main.go

# Linux 64位
GOOS=linux GOARCH=amd64 go build -o kss-cli-linux-amd64 main.go

echo "Build finished for all platforms."