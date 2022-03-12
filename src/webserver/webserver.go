package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	aero "github.com/aerospike/aerospike-client-go"
)

func main() {
	client, err := aero.NewClient("192.168.88.190", 3000)
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

func userAccessor(c Client) func(func(error), string, string) {
	return func(returnIfError func(error), userId string, apiKey string) (string, error) {
		intClientId, err := strconv.Atoi(userId)
		key, err := aero.NewKey("test", "users", intClientId)
		returnIfError(err)

	}
}

func userAccess(w http.ResponseWriter, r *http.Request) {
	userid := strings.TrimPrefix(r.URL.Path, "/user/")
	apikey := r.URL.Query().Get("api_key")
	response, _ := getUser(userid, apikey, returnError(w))
	fmt.Fprintf(w, response)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func returnError(w http.ResponseWriter) func(error) {
	return func(err error) {

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			resp := make(map[string]string)
			resp["message"] = fmt.Sprintf("Resource Not Found, %s", err)
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)
		}

		return
	}
}
func getUser(userId string, apiKey string, returnIfError func(error)) (string, error) {
	client, err := aero.NewClient("192.168.88.190", 3000)
	returnIfError(err)
	intClientId, err := strconv.Atoi(userId)
	key, err := aero.NewKey("test", "users", intClientId)
	returnIfError(err)

	// read it back!
	rec, err := client.Get(nil, key)
	returnIfError(err)

	return string(fmt.Sprintf("Hello! running jwAreoSpike search for user %s with api_key %s found %s", userId, apiKey, rec)), nil
}
