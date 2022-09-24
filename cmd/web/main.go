package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/winaldomanurung/go-basic-project/pkg/config"
	"github.com/winaldomanurung/go-basic-project/pkg/handlers"
	"github.com/winaldomanurung/go-basic-project/pkg/render"
)

var app config.AppConfig
// Jika mau deploy harus hapus 127.0.0.1
const portNumber = "127.0.0.1:8080"
var session *scs.SessionManager

func main() {
	// app merupakan variable berisi UseCache dan TemplateCache

	app.InProduction = false


	session = scs.New()
	session.Lifetime = 24 * time.Hour
	// untuk ketika browser di close
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //di set true di production

	app.Session = session
	// CreateTemplateCache mereturn map berisi referan Template struct dan error (kalau ada)
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// app kita ubah nilainya
	// app.TemplateCache adalah map berisi referan Template struct
	app.TemplateCache = tc
	app.UseCache = true

	// NewRepo kita kasih argument app (refer ke AppConfig). Dia merupakan alamat dari app
	repo := handlers.NewRepo(&app)
	// NewHandlers menerima argument repo. Asalnya dia adalah merefer ke Repository struct, yang mana isinya adalah AppConfig.
	handlers.NewHandlers(repo)

	// ini untuk set config untuk template package
	render.NewTemplates(&app)

	// Home dan About available karena ada Receiver di Repository struct
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}