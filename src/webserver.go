package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/user/", indexHandlerHelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	userid := strings.TrimPrefix(r.URL.Path, "/user/")
	apikey := r.URL.Query().Get("api_key")
	fmt.Fprintf(w, "Hello! running search for user %s with api_key %s", userid, apikey)
}
