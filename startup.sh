#!/bin/bash 
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o ./target/go-webserver ./src/webserver.go || exit 1 

if [ "$(docker ps -q -f name=go-web-server)" ]; then 
 docker stop $(docker ps -a -q --filter name=go-web-server --format="{{.ID}}") 
fi 

docker build -f ./deploy/Dockerfile -t go-web-server . || exit 1 

docker run -p 8000:8080 go-web-server || exit 1
