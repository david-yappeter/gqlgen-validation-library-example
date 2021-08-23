package main

import (
	"fmt"
	"os"

	"github.com/david-yappeter/gqlgen-validation-library-example/plugin"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
    // This will get config file from gqlgen.yml
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

    // Generate with our custom plugin
	err = api.Generate(cfg,
		api.AddPlugin(plugin.New()),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
