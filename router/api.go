package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"projects/starter/renderer"
)

// APIRoutes function for the api request only
func ServeAPIRoutes(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter().StrictSlash(true)
	api.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		if err := renderer.JSON(w, http.StatusOK,
			map[string]interface{}{"message": "Venom API routes is places in $ROOT/router/api.go"}); err != nil {
			panic(err.Error())
		}
	}).Methods("GET")

	// middleware
	// api.Use(app.JwtAuthentication)
}
