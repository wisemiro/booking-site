package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wycemiro/booking-site/internal/config"
	"github.com/wycemiro/booking-site/internal/models"
	"github.com/wycemiro/booking-site/internal/renders"
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
	renders.RenderTemplates(w, r, "home.page.tmpl", &models.TemplateData{})
}

//About is handler for the about page.
func (b *Repository) About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplates(w, r, "about.page.tmpl", &models.TemplateData{})
}

//Contact is handler for contact page.
func (b *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplates(w, r, "contact.page.tmpl", &models.TemplateData{})
}

//Major
func (b *Repository) Major(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplates(w, r, "majors.page.tmpl", &models.TemplateData{})
}

//search-availability
func (b *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplates(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

//post search-availability
func (b *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {

}

//General is handler for the generals page.
func (b *Repository) General(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplates(w, r, "generals.page.tmpl", &models.TemplateData{})
}

type jsonResp struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

//AvailabilityJson handles req for availability and retuns json
func (b *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResp{
		OK:      true,
		Message: "Available!",
	}
	out, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
