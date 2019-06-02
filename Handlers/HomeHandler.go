package handlers

import (
	"net/http"
	"projects/starter/renderer"
)

// Home struct
type Home struct{}

// HomeHandlers function return home struct with all his function
func HomeHandlers() *Home { return &Home{} }

// Index is starter handler
func (home Home) Index(w http.ResponseWriter, r *http.Request) {
	if err := renderer.HTML(w, http.StatusOK, "home", nil); err != nil {
		panic(err.Error())
	}
}
