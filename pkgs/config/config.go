package config

import "text/template"

type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
}
