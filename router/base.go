package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"projects/starter/utils"
)

// New Router function
func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	APIRoutes(router)
	WEBRoutes(router)

	// serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// 404 view
	router.NotFoundHandler = http.HandlerFunc(utils.NotFoundHandler)
	return router
}
