package gena

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
	"text/template"
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
	tmpl := template.Must(template.New("webstack.tmpl").Funcs(map[string]interface{}{
		"icon": icon,
	}).ParseFS(f, "templates/webstack.tmpl"))

	if err := tmpl.Execute(os.Stdout, cfg); err != nil {
		log.Fatal("[webstack] Render template error: ", err.Error())
	}
}

// icon return favicon of url
func icon(rawurl string) string {
	base := "https://f.start.me/%s"
	rawurl = strings.TrimSpace(rawurl)
	u, err := url.Parse(rawurl)
	if err != nil {
		return fmt.Sprintf(base, "o.oo")
	}
	host := u.Host
	if strings.Contains(host, ":") {
		host, _, _ = net.SplitHostPort(host)
	}
	return fmt.Sprintf(base, host)
}
