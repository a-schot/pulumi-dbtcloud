//go:build dotnet || all
// +build dotnet all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccProjectCSharp(t *testing.T) {
	test := getCSharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "project", "csharp"),
		})

	integration.ProgramTest(t, &test)
}

func getCSharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseCSharp := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"ASchot.Pulumi.Dbtcloud",
		},
	})

	return baseCSharp
}
