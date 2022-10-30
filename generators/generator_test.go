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
	default:
		log.Fatal("Invalid template name, expected: webstack")
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
