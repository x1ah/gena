package generators

import (
	"io"
	"log"
	"text/template"

	"github.com/x1ah/gena"
)

// WebStackGenerator generate webstack template
type WebStackGenerator struct{}

// Run implement Generator.Run
func (ws *WebStackGenerator) Run(cfg *gena.Config, writer io.Writer) {
	tmpl := template.Must(template.New("webstack.tmpl").Funcs(map[string]interface{}{
		"icon": favicon,
	}).ParseFS(gena.Templates, "templates/webstack.tmpl"))

	if err := tmpl.Execute(writer, cfg); err != nil {
		log.Fatal("[webstack] Render template error: ", err.Error())
	}
}
