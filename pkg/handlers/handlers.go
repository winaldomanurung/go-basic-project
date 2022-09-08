package handlers

import (
	"net/http"

	"github.com/winaldomanurung/go-basic-web-app/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.page.tmpl")
}

