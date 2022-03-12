#!/bin/bash 
if [ "$(docker ps -q -f name=go-web-server)" ]; then 
 docker stop $(docker ps -a -q --filter name=go-web-server --format="{{.ID}}") 
fi 
