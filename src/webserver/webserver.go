package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	aero "github.com/aerospike/aerospike-client-go"
)

func main() {
	client, err := aero.NewClient("aerospike", 3000)
	panicOnError(err)

	key, err := aero.NewKey("test", "users", 11)
	panicOnError(err)

	// define some bins with data
	bins := aero.BinMap{
		"api_key":    "42",
		"first_name": "jonathan",
		"last_name":  "wilson",
		"company":    "mindbodyengineer",
	}

	// write the bins
	err = client.Put(nil, key, bins)
	panicOnError(err)

	http.HandleFunc("/user/", userAccessor(client))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func userAccessor(c *aero.Client) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := strings.TrimPrefix(r.URL.Path, "/user/")
		apikey := r.URL.Query().Get("api_key")
		response, err := getUser(userid, apikey, c)
		if err == nil {
			fmt.Fprintf(w, response)
		} else {
			returnError(w)(err)
			fmt.Fprintf(w, "")
		}
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func returnError(w http.ResponseWriter) func(error) {
	return func(err error) {
		if err != nil {
			if err.Error() == "401" {
				w.WriteHeader(http.StatusUnauthorized)
				w.Header().Set("Content-Type", "application/json")
				resp := make(map[string]string)
				resp["message"] = "Unauthorized"
				jsonResp, err := json.Marshal(resp)
				if err != nil {
					log.Fatalf("Error happened in JSON marshal. Err: %s", err)
				}
				w.Write(jsonResp)
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "application/json")
				resp := make(map[string]string)
				resp["message"] = "Resource Not Found"
				jsonResp, err := json.Marshal(resp)
				if err != nil {
					log.Fatalf("Error happened in JSON marshal. Err: %s", err)
				}
				w.Write(jsonResp)
			}
		}

	}
}

func getUser(userId string, apiKey string, client *aero.Client) (string, error) {
	intClientId, err := strconv.Atoi(userId)
	key, err := aero.NewKey("test", "users", intClientId)
	if err == nil {
		rec, err := client.Get(nil, key)
		if err == nil {
			if rec.Bins["api_key"] == apiKey {
				jsonResp, err := json.Marshal(rec.Bins)
				return string(jsonResp), err
			} else {
				return "", errors.New("401")
			}
		} else {
			return "", errors.New("404")
		}
	} else {
		return "", err
	}
}
