#!/bin/bash 
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o ./target/go-webserver ./src/webserver.go || exit 1 
