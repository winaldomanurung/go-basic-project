package handlers

import (
	"net/http"

	"github.com/winaldomanurung/go-basic-web-app/pkg/config"
	"github.com/winaldomanurung/go-basic-web-app/pkg/models"
	"github.com/winaldomanurung/go-basic-web-app/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct{
	App *config.AppConfig
}

// NewRepo creates a new repository
// menerima argument a yang merefer ke *config.AppConfig
// mereturn sesuatu yang merefer ke *Repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

// NewHandlers set the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}

// kita bind Repository struct dengan function ini
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

