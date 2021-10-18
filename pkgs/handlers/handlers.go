package handlers

import (
	"github.com/wycemiro/booking-site/pkgs/config"
	"github.com/wycemiro/booking-site/pkgs/models"
	"github.com/wycemiro/booking-site/pkgs/renders"
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
	renders.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{})
}

//About is handler for the about page.
func (b *Repository) About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplates(w, "about.page.tmpl", &models.TemplateData{})
}

//Contact is handler for contact page.
func (b *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplates(w, "contact.page.tmpl", &models.TemplateData{})
}

//Major
func (b *Repository) Major(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplates(w, "majors.page.tmpl", &models.TemplateData{})
}

//search-availability
func (b *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{})
}

//General is handler for the generals page.
func (b *Repository) General(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplates(w, "generals.page.tmpl", &models.TemplateData{})
}
