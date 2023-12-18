package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {

	r := chi.NewRouter()
	// log all requests
	r.Use(middleware.Logger)
	// recover from panics without crashing server
	r.Use(middleware.Recoverer)
	// set content type to json
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	http.ListenAndServe(":8080", r)
}
