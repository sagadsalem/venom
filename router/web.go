package router

import (
	"github.com/gorilla/mux"
	"projects/starter/controllers"
)

// WEBRoutes function for the web request only
func WEBRoutes(r *mux.Router) {
	web := r.PathPrefix("/").Subrouter().StrictSlash(true)
	web.HandleFunc("/", controllers.HomeController().HomeHandler).Methods("GET")
	// middleware
	// web.Use(app.WebAuthentication)
}
