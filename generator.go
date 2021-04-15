package gena

import (
	"embed"
	"html/template"
	"log"
	"os"
)

// Generator is interface of template generator
type Generator interface {
	Run(cfg *Config)
}

//go:embed templates/*
var f embed.FS

// WebStackGenerator generate webstack template
type WebStackGenerator struct{}

// Run implement Generator.Run
func (ws *WebStackGenerator) Run(cfg *Config) {
	tmpl := template.Must(template.ParseFS(f, "templates/webstack.tmpl"))
	if err := tmpl.Execute(os.Stdout, cfg); err != nil {
		log.Fatal("[webstack] Render template error: ", err.Error())
	}
}
