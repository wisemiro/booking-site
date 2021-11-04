package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InProduction  bool
	Sessions      *scs.SessionManager
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
}
