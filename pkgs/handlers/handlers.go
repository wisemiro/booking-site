package handlers

import (
	"learn/pkgs/config"
	"learn/pkgs/models"
	"learn/pkgs/renders"
	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is handler for the home page
func (b *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "yawh"
	renders.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
