package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Get Wrap the router for GET method
func Get(r *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.HandleFunc(path, f).Methods("GET")
}

// Post Wrap the router for POST method
func Post(r *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.HandleFunc(path, f).Methods("POST")
}

// Put Wrap the router for PUT method
func Put(r *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.HandleFunc(path, f).Methods("PUT")
}

// Delete Wrap the router for DELETE method
func Delete(r *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.HandleFunc(path, f).Methods("DELETE")
}
