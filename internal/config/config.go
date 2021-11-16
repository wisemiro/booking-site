package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
	"github.com/wycemiro/booking-site/internal/models"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InProduction  bool
	Sessions      *scs.SessionManager
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	MailChan      chan models.MailData
}
