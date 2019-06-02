package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"projects/starter/renderer"
)

// New Router function
func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	ServeAPIRoutes(router)
	ServeWEBRoutes(router)
	ServeStaticFile(router)
	// 404 handler
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	return router
}

func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	if err := renderer.HTML(w, http.StatusNotFound, "404", nil); err != nil {
		println(err.Error())
	}
}

func ServeStaticFile(router *mux.Router) {
	// serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
}
