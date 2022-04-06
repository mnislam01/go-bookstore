package main

import (
	"BookStoreAPIs/core/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RoutesRegistry(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
