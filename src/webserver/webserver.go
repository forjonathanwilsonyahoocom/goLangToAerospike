package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/user/", userAccess)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func userAccess(w http.ResponseWriter, r *http.Request) {
	userid := strings.TrimPrefix(r.URL.Path, "/user/")
	apikey := r.URL.Query().Get("api_key")
	response, _ := getUser(userid, apikey)
	fmt.Fprintf(w, response)
}

func getUser(userId string, apiKey string) (string, error) {

	return string(fmt.Sprintf("Hello! running jwAreoSpike search for user %s with api_key %s", userId, apiKey)), nil
}
