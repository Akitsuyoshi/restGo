package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter exports each routes, according to its path
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	r := router.PathPrefix("/v1/").Subrouter()

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return r
}
