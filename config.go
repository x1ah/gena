package gena

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config is struct of gena's config
type Config struct {
	Title           string
	Description     string
	Template        string
	URL             string
	Logo            string
	Favicon         string
	Github          string
	Footer          string
	Content         *Content
	GoogleAnalytics string `yaml:"google_analytics"`
	WebStack        *WebStackConf
	Tilde           *TildeConf
}

// Content is struct of categories
type Content struct {
	Categories []*Category
}

// Category struct
type Category struct {
	Name  string
	Sites []*Site
}

// Site struct
type Site struct {
	Name        string
	Description string
	URL         string
	Icon        string // 网站图标
}

// WebStackConf is config of webstack
type WebStackConf struct {
	Search *WebStackSearchConf
}

// WebStackSearchConf search engine config
type WebStackSearchConf struct {
	Enabled bool
	Default string
	Engines []string
}

// TildeConf is config of tilde-enhanced
type TildeConf struct {
	Search *TildeSearchConf
	Theme  string
	ShowKeys bool `yaml:"show_keys"`
}

// TildeSearchConf search config for tilde
type TildeSearchConf struct {
	URL         string
	Placeholder string
}

// ParseConfig parse config from file
func ParseConfig(file string) (*Config, error) {
	buf, err := ioutil.ReadFile(filepath.Clean(file))
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(buf, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
