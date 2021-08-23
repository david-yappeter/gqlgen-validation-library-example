package plugin

import (
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin"
	"github.com/vektah/gqlparser/v2/ast"
)

// Define our generator struct
type generator struct{}

// Initialize 'generator' to plugin.Plugin
func New() plugin.Plugin {
	return &generator{}
}

// Name of our plugin
func (g *generator) Name() string {
	return "gql_generator"
}

// Add our custom directive @binding to configuration
func (g *generator) MutateConfig(cfg *config.Config) error {
	cfg.Directives["binding"] = config.DirectiveConfig{SkipRuntime: false}
	return nil
}

// Inject a file before generating the resolver (doesn't generate a new file)
func (g *generator) InjectSourceEarly() *ast.Source {
	return &ast.Source{
		Name:    "validationDirective.graphqls",
		BuiltIn: false,
		Input: `
            directive @binding(constraint: String!) on INPUT_FIELD_DEFINITION | ARGUMENT_DEFINITION
        `,
	}
}
