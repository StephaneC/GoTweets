package main

import (
	"net/http"
	"encoding/json"
	                "fmt"
	//"github.com/gorilla/mux"
	//"strconv"
)

func GetTweetsHandler(w http.ResponseWriter, r *http.Request) {
	//get vars from request
	//vars := mux.Vars(r)
	fmt.Printf("Get Tweets handler")

	//ts, _ := strconv.ParseInt(r.Queries("ts", "0"), 10, 64)
	json.NewEncoder(w).Encode(getTweets(0))
}
