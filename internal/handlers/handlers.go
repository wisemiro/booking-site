package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wycemiro/booking-site/internal/config"
	"github.com/wycemiro/booking-site/internal/forms"
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

//Reservation renders reservation page
func (b *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	renders.RenderTemplates(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

//PostReservation posts a reservation
func (b *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	//pass user data from the form to the reservation model
	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
	}
	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 4, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		renders.RenderTemplates(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

}
