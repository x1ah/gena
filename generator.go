package gena

import (
	"html/template"
	"log"
	"os"
)

// Generator is interface of template generator
type Generator interface {
	Run(cfg *Config)
}

// WebStackGenerator generate webstack template
type WebStackGenerator struct{}

// Run implement Generator.Run
func (ws *WebStackGenerator) Run(cfg *Config) {
	tmpl := template.Must(template.ParseFiles("templates/webstack.tmpl"))
	if err := tmpl.Execute(os.Stdout, cfg); err != nil {
		log.Fatal("[webstack] Render template error: ", err.Error())
	}
}
