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
	return router
}

// NotFoundHandler function
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	utils.Respond(w, utils.Message(false, "This resources was not found on our server"))
	return
}
