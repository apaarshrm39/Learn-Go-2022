package handlers

import (
	"net/http"

	"example.com/html/pkg/config"
	"example.com/html/pkg/models"
	"example.com/html/pkg/render"
)

// Repo the repository used by the handlers.
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	StringMap := make(map[string]string)
	StringMap["test"] = "Hello Again"
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: StringMap,
	})
}
