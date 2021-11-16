package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/wycemiro/booking-site/internal/config"
	"github.com/wycemiro/booking-site/internal/driver"
	"github.com/wycemiro/booking-site/internal/forms"
	"github.com/wycemiro/booking-site/internal/helpers"
	"github.com/wycemiro/booking-site/internal/models"
	"github.com/wycemiro/booking-site/internal/renders"
	"github.com/wycemiro/booking-site/internal/repository"
	"github.com/wycemiro/booking-site/internal/repository/dbrepo"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is handler for the home page
func (b *Repository) Home(w http.ResponseWriter, r *http.Request) {
	renders.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

//About is handler for the about page.
func (b *Repository) About(w http.ResponseWriter, r *http.Request) {
	renders.Template(w, r, "about.page.tmpl", &models.TemplateData{})
}

//Contact is handler for contact page.
func (b *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renders.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
}

//Major
func (b *Repository) Major(w http.ResponseWriter, r *http.Request) {
	renders.Template(w, r, "majors.page.tmpl", &models.TemplateData{})
}

//search-availability
func (b *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	renders.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

//post search-availability
func (b *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	layout := "2006-01-02"

	startDate, err := time.Parse(layout, start)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rooms, err := b.DB.SearchAvailabilityForAllRooms(startDate, endDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	for _, i := range rooms {
		b.App.InfoLog.Println("ROOM:", i.ID, i.RoomName)
	}
	if len(rooms) == 0 {
		b.App.Sessions.Put(r.Context(), "error", "No availability")
		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	res := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}
	b.App.Sessions.Put(r.Context(), "reservation", res)

	renders.Template(w, r, "choose-room.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

//General is handler for the generals page.
func (b *Repository) General(w http.ResponseWriter, r *http.Request) {
	renders.Template(w, r, "generals.page.tmpl", &models.TemplateData{})
}

type jsonResp struct {
	OK        bool   `json:"ok"`
	Message   string `json:"message"`
	RoomID    string `json:"room_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

//AvailabilityJson handles req for availability and retuns json
func (b *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {

	sd := r.Form.Get("start")
	ed := r.Form.Get("end")
	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, sd)
	endDate, _ := time.Parse(layout, ed)

	roomID, _ := strconv.Atoi(r.Form.Get("room_id"))

	available, _ := b.DB.SearchAvailabilityByDatesByRoomID(startDate, endDate, roomID)
	resp := jsonResp{
		OK:        available,
		Message:   "",
		StartDate: sd,
		EndDate:   ed,
		RoomID:    strconv.Itoa(roomID),
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
	res, ok := b.App.Sessions.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("cant get reservations from sessions"))
		return
	}
	room, err := b.DB.GetRoomByID(res.RoomID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	res.Room.RoomName = room.RoomName

	b.App.Sessions.Put(r.Context(), "reservation", res)
	sd := res.StartDate.Format("2006-01-02")
	ed := res.EndDate.Format("2006-01-02")
	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed
	data := make(map[string]interface{})
	data["reservation"] = res
	renders.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

//PostReservation posts a reservation
func (b *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {

	reservation, ok := b.App.Sessions.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("cant get from session"))
		return
	}

	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation.FirstName = r.Form.Get("first_name")
	reservation.LastName = r.Form.Get("last_name")
	reservation.Email = r.Form.Get("email")
	reservation.Phone = r.Form.Get("phone")

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 4, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		renders.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	newReservationID, err := b.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
	}

	restriction := models.RoomRestriction{
		RoomID:        reservation.RoomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
		StartDate:     reservation.StartDate,
		EndDate:       reservation.EndDate,
	}

	err = b.DB.InstertRoomRestriction(restriction)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	b.App.Sessions.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (b *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := b.App.Sessions.Get(r.Context(), "reservation").(models.Reservation) //assertion
	if !ok {
		log.Println("cant get item from session.")
		b.App.Sessions.Put(r.Context(), "error", "Cant get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	b.App.Sessions.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["resercation"] = reservation
	sd := reservation.StartDate.Format("2006-01-02")
	ed := reservation.EndDate.Format("2006-01-02")
	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	renders.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})
}

func (b *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res, ok := b.App.Sessions.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, err)
		return
	}
	res.RoomID = roomID
	b.App.Sessions.Put(r.Context(), "reservation", res)
	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}

func (b *Repository) BookRoom(w http.ResponseWriter, r *http.Request) {
	roomID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	sd := r.URL.Query().Get("s")
	ed := r.URL.Query().Get("e")

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, sd)
	endDate, _ := time.Parse(layout, ed)

	var res models.Reservation
	res.RoomID = roomID
	res.StartDate = startDate
	res.EndDate = endDate

	room, err := b.DB.GetRoomByID(res.RoomID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	res.Room.RoomName = room.RoomName
	b.App.Sessions.Put(r.Context(), "reservation", res)
	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}
