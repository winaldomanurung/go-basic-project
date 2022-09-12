package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/winaldomanurung/go-basic-web-app/pkg/config"
	"github.com/winaldomanurung/go-basic-web-app/pkg/handlers"
	"github.com/winaldomanurung/go-basic-web-app/pkg/render"
)

// Jika mau deploy harus hapus 127.0.0.1
const portNumber = "127.0.0.1:8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}