package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTweets(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(beers)
}
