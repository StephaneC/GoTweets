package main

import (
	"net/http"
	"encoding/json"
)

func GetTweets(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getTweets(0))
}
