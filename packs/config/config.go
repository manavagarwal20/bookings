package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type App struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Infolog       *log.Logger
	InProd        bool
	Session       *scs.SessionManager
}
