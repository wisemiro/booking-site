package models

import "github.com/wycemiro/booking-site/internal/forms"

//TemplateData to send data from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CRSFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
