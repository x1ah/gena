package main

import (
	"flag"
	"log"
	"os"

	"github.com/x1ah/gena"
	"github.com/x1ah/gena/generators"
)

func main() {
	cfgFile := flag.String("c", "config.yml", "Config file")

	flag.Parse()

	cfg, err := gena.ParseConfig(*cfgFile)
	if err != nil {
		log.Fatal("Parse Config file error: ", err.Error())
	}

	var generator generators.Generator
	switch cfg.Template {
	case "webstack":
		generator = &generators.WebStackGenerator{}
	case "tilde":
		generator = &generators.TildeGenerator{}
	default:
		log.Fatal("Invalid template name, expected: webstack or tilde")
	}
	generator.Run(cfg, os.Stdout)
}
