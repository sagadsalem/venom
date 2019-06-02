package router

import (
	"github.com/gorilla/mux"
	"projects/starter/handlers"
	"projects/starter/middleware"
)

// WEBRoutes function for the web request only
func ServeWEBRoutes(r *mux.Router) {
	web := r.PathPrefix("/").Subrouter().StrictSlash(true)
	web.HandleFunc("/", handlers.HomeHandlers().Index).Methods("GET")

	// middleware
	web.Use(middleware.RequestPathLogger)
}
