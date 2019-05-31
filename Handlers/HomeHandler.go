package handlers

import (
	"net/http"
	"projects/starter/utils"
)

// Home Struct
type Home struct{}

// HomeHandlers function return home struct with all his function
func HomeHandlers() *Home { return &Home{} }

// Index is starter handler
func (home Home) Index(w http.ResponseWriter, r *http.Request) {
	utils.Tpl.ExecuteTemplate(w, "home.html", nil)
	return
}
