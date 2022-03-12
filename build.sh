#!/bin/bash 
cd ./src/webserver
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o ../../bin/go-webserver webserver.go || exit 1 
cd ../../
