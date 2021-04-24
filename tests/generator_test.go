package tests

import (
	"bytes"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"testing"

	"github.com/x1ah/gena"
	"github.com/x1ah/gena/generators"
)

func compareOutput(t *testing.T, configPath string) {
	if !strings.HasSuffix(configPath, ".yml") {
		return
	}

	outputPath := configPath + ".output"
	cfg, err := gena.ParseConfig(configPath)
	if err != nil {
		t.Error("parse config error: ", configPath, err.Error())
		return
	}
	var generator generators.Generator
	switch cfg.Template {
	case "webstack":
		generator = &generators.WebStackGenerator{}
	default:
		log.Fatal("Invalid template name, expected: webstack")
	}
	writer := new(bytes.Buffer)
	generator.Run(cfg, writer)

	outout, err := ioutil.ReadFile(outputPath)
	if err != nil {
		t.Error("read output file error: ", err.Error())
		return
	}
	if string(outout) != writer.String() {
		t.Fatal("output compare error: ", configPath, outputPath)
	}
}

func TestGenerator(t *testing.T) {
	var walk = func(path string, info fs.FileInfo, err error) error {
		compareOutput(t, path)
		return nil
	}
	if err := filepath.Walk("spec", walk); err != nil {
		t.Error(err)
	}
}
