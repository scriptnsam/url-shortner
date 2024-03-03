package routes

import (
	"github.com/gorilla/mux"
	"github.com/scriptnsam/url-shortner/pkg/controllers"
)

var ShortnerRoutes = func(routes *mux.Router) {
	routes.HandleFunc("/", controllers.GetUrls).Methods("GET")
	routes.HandleFunc("/{id}", controllers.RedirectHandler).Methods("GET")
	routes.HandleFunc("/create", controllers.CreateLinkHandler).Methods("POST")
}
