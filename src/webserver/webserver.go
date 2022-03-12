package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	aero "github.com/aerospike/aerospike-client-go"
)

func main() {
	client, err := aero.NewClient("localhost", 3000)
	panicOnError(err)

	key, err := aero.NewKey("test", "users", 11)
	panicOnError(err)

	// define some bins with data
	bins := aero.BinMap{
		"api_key":    42,
		"first_name": "jonathan",
		"last_name":  "wilson",
		"company":    "mindbodyengineer",
	}

	// write the bins
	err = client.Put(nil, key, bins)
	panicOnError(err)

	http.HandleFunc("/user/", userAccess)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func userAccess(w http.ResponseWriter, r *http.Request) {
	userid := strings.TrimPrefix(r.URL.Path, "/user/")
	apikey := r.URL.Query().Get("api_key")
	response, _ := getUser(userid, apikey)
	fmt.Fprintf(w, response)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func getUser(userId string, apiKey string) (string, error) {
	client, err := aero.NewClient("127.0.0.1", 3000)
	panicOnError(err)
	intClientId, err := strconv.Atoi(userId)
	key, err := aero.NewKey("test", "users", intClientId)
	panicOnError(err)

	// read it back!
	rec, err := client.Get(nil, key)
	panicOnError(err)

	return string(fmt.Sprintf("Hello! running jwAreoSpike search for user %s with api_key %s found %s", userId, apiKey, rec)), nil
}
