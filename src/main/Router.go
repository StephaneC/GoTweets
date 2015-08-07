package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{"GetTweets", "GET", "/tweets", GetTweets},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		fmt.Printf("routing : " + route.Pattern)
		router.
			Methods(route.Method).Path(route.Pattern).
			Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}
