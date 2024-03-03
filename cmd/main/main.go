package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scriptnsam/url-shortner/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.ShortnerRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
