#!/bin/bash 
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o ./bin/go-webserver ./src/webserver/webserver.go || exit 1 
