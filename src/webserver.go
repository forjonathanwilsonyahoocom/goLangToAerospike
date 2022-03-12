package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("/", indexHandlerHelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, I'm running search on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}
