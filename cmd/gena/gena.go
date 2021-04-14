package main

import (
	"flag"
	"log"

	"github.com/x1ah/gena"
)

func main() {
	tmpl := flag.String("t", "webstack", "Template name, available: webstack")
	cfgFile := flag.String("c", "config.yml", "Config file")

	flag.Parse()

	cfg, err := gena.ParseConfig(*cfgFile)
	if err != nil {
		log.Fatal("Parse Config file error: ", err.Error())
	}

	var generator gena.Generator
	switch *tmpl {
	case "webstack":
		generator = &gena.WebStackGenerator{}
	default:
		log.Fatal("Invalid template name, expected: webstack")
	}
	generator.Run(cfg)
}
