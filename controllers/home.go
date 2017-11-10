package controllers

import (
	"net/http"
)

// Home page
func (m Main) Home(w http.ResponseWriter, r *http.Request) error {
	return m.Render.HTML(w, http.StatusOK, "main/home", nil)
}
