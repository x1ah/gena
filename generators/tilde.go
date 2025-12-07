package generators

import (
	"io"
	"log"
	"strings"
	"text/template"

	"github.com/x1ah/gena"
)

// TildeGenerator generate tilde-enhanced template
type TildeGenerator struct{}

// Run implement Generator.Run
func (tg *TildeGenerator) Run(cfg *gena.Config, writer io.Writer) {
	tmpl := template.Must(template.New("tilde.tmpl").Funcs(map[string]interface{}{
		"icon": favicon,
		"escapeJS": func(s string) string {
			s = strings.ReplaceAll(s, "\\", "\\\\")
			s = strings.ReplaceAll(s, "'", "\\'")
			s = strings.ReplaceAll(s, "\n", "\\n")
			s = strings.ReplaceAll(s, "\r", "\\r")
			return s
		},
		"getFirstChar": func(s string) string {
			s = strings.TrimSpace(s)
			if len(s) == 0 {
				return ""
			}
			first := strings.ToLower(string(s[0]))
			// Only return alphanumeric characters
			if (first >= "a" && first <= "z") || (first >= "0" && first <= "9") {
				return first
			}
			return ""
		},
		"getSearchURL": func() string {
			if cfg.Tilde != nil && cfg.Tilde.Search != nil && cfg.Tilde.Search.URL != "" {
				return cfg.Tilde.Search.URL
			}
			return "https://duckduckgo.com/?q="
		},
		"getSearchPlaceholder": func() string {
			if cfg.Tilde != nil && cfg.Tilde.Search != nil && cfg.Tilde.Search.Placeholder != "" {
				return cfg.Tilde.Search.Placeholder
			}
			return "Search or enter URL"
		},
		"getTheme": func() string {
			if cfg.Tilde != nil && cfg.Tilde.Theme != "" {
				return cfg.Tilde.Theme
			}
			return "dark"
		},
		"showKeys": func() bool {
			return cfg.Tilde != nil && cfg.Tilde.ShowKeys
		},
	}).ParseFS(gena.Templates, "templates/tilde.tmpl"))

	if err := tmpl.Execute(writer, cfg); err != nil {
		log.Fatal("[tilde] Render template error: ", err.Error())
	}
}

