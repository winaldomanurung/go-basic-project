package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Jika mau deploy harus hapus 127.0.0.1
const portNumber = "127.0.0.1:8080"


func Home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate (w http.ResponseWriter, tmpl string){
	// kita parse salah satu template
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)
	if err !=nil {
		fmt.Println("error parsing template:", err)
		return
	}
}


func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}