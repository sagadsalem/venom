package controllers

import (
	"net/http"
	"projects/starter/utils"
)

// Home Struct
type Home struct{}

// HomeController function return home struct with all his function
func HomeController() *Home { return &Home{} }

// HomeHandler is starter handler
func (home Home) HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.Tpl.ExecuteTemplate(w, "web/home.html", nil)
	return
}
