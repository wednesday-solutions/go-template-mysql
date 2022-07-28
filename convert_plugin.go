//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	gbgen "github.com/web-ridge/gqlgen-sqlboiler/v3"
)

func main() {
	output := gbgen.Config{
		Directory:   "helpers", // supports root or sub directories
		PackageName: "helpers",
	}
	backend := gbgen.Config{
		Directory:   "models",
		PackageName: "models",
	}
	frontend := gbgen.Config{
		Directory:   "gqlmodels",
		PackageName: "gqlmodels",
	}

	if err := gbgen.SchemaWrite(gbgen.SchemaConfig{
		BoilerModelDirectory: backend,
		// Directives:           []string{"IsAuthenticated"},
		// GenerateBatchCreate:  false, // Not implemented yet
		GenerateMutations:   true,
		GenerateBatchDelete: true,
		GenerateBatchUpdate: true,
	}, "schema.graphql", gbgen.SchemaGenerateConfig{
		MergeSchema: false, // uses three way merge to keep your customization
	}); err != nil {
		fmt.Println("error while trying to generate schema.graphql")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	if err = api.Generate(cfg,
		api.AddPlugin(gbgen.NewConvertPlugin(
			output,   // directory where convert.go, convert_input.go and preload.go should live
			backend,  // directory where sqlboiler files are put
			frontend, // directory where gqlgen models live
			gbgen.ConvertPluginConfig{
				DatabaseDriver: gbgen.MySQL, // or gbgen.PostgreSQL,
			},
		)),
		api.AddPlugin(gbgen.NewResolverPlugin(
			output,
			backend,
			frontend,
			gbgen.ResolverPluginConfig{
				// See example for AuthorizationScopes here: https://github.com/web-ridge/gqlgen-sqlboiler-examples/blob/main/social-network/convert_plugin.go#L66
			},
		)),
	); err != nil {
		fmt.Println("error while trying generate resolver and converts")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
