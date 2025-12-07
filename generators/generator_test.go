package generators

import (
	"bytes"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/x1ah/gena"
)

func compareOutput(t *testing.T, configPath string) {
	if !strings.HasSuffix(configPath, ".yml") {
		return
	}

	t.Logf("Compare %s output", configPath)
	outputPath := configPath + ".output"
	cfg, err := gena.ParseConfig(configPath)
	require.Nil(t, err)
	var generator Generator
	switch cfg.Template {
	case "webstack":
		generator = &WebStackGenerator{}
	case "tilde":
		generator = &TildeGenerator{}
	default:
		log.Fatal("Invalid template name, expected: webstack or tilde")
	}
	writer := new(bytes.Buffer)
	generator.Run(cfg, writer)

	outout, err := ioutil.ReadFile(outputPath)
	require.Nil(t, err)
	require.Equal(t, string(outout), writer.String())
	t.Logf("Compare %s output success", configPath)
}

func TestGenerator(t *testing.T) {
	var walk = func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			t.Fatal("walk spec error: ", err)
			return err
		}
		compareOutput(t, path)
		return nil
	}
	err := filepath.Walk("../tests/spec", walk)
	require.Nil(t, err)
}

func TestTildeGenerator(t *testing.T) {
	cfg := &gena.Config{
		Title:       "Test",
		Description: "Test description",
		Template:    "tilde",
		Favicon:     "https://example.com/favicon.ico",
		URL:         "https://example.com",
		Github:      "https://github.com/test",
		Footer:      "© 2024",
		Tilde: &gena.TildeConf{
			Search: &gena.TildeSearchConf{
				URL:         "https://duckduckgo.com/?q=",
				Placeholder: "Search",
			},
			Theme:    "dark",
			ShowKeys: false,
		},
		Content: &gena.Content{
			Categories: []*gena.Category{
				{
					Name: "Test Category",
					Sites: []*gena.Site{
						{
							Name:        "Test Site",
							Description: "Test description",
							URL:         "https://example.com",
							Icon:        "https://example.com/icon.png",
						},
					},
				},
			},
		},
	}

	generator := &TildeGenerator{}
	writer := new(bytes.Buffer)
	generator.Run(cfg, writer)

	output := writer.String()
	require.Contains(t, output, "<title>Test</title>")
	require.Contains(t, output, "Test Category")
	require.Contains(t, output, "Test Site")
	require.Contains(t, output, "https://example.com")
}

func TestTildeGeneratorWithShowKeys(t *testing.T) {
	cfg := &gena.Config{
		Title:    "Test",
		Template: "tilde",
		Tilde: &gena.TildeConf{
			ShowKeys: true,
		},
		Content: &gena.Content{
			Categories: []*gena.Category{
				{
					Name: "Test",
					Sites: []*gena.Site{
						{
							Name: "GitHub",
							URL:  "https://github.com",
						},
					},
				},
			},
		},
	}

	generator := &TildeGenerator{}
	writer := new(bytes.Buffer)
	generator.Run(cfg, writer)

	output := writer.String()
	require.Contains(t, output, "site-key")
}

func TestTildeGeneratorDefaultConfig(t *testing.T) {
	cfg := &gena.Config{
		Title:    "Test",
		Template: "tilde",
		Content: &gena.Content{
			Categories: []*gena.Category{
				{
					Name: "Test",
					Sites: []*gena.Site{
						{
							Name: "Test",
							URL:  "https://example.com",
						},
					},
				},
			},
		},
	}

	generator := &TildeGenerator{}
	writer := new(bytes.Buffer)
	generator.Run(cfg, writer)

	output := writer.String()
	// Should use default search URL
	require.Contains(t, output, "duckduckgo.com")
	// Should use default theme (dark)
	require.Contains(t, output, "#1a1a1a")
}
