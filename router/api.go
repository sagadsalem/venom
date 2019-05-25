package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"projects/starter/utils"
)

// APIRoutes function for the api request only
func APIRoutes(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter().StrictSlash(true)
	api.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		utils.Respond(w, utils.Message(true, "api routes and functions goes here"))
		return
	}).Methods("GET")

	// middleware
	// api.Use(app.JwtAuthentication)
}
