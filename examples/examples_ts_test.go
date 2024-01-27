//go:build nodejs || all
// +build nodejs all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccProjectTS(t *testing.T) {
	test := getTSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "project", "ts"),
		})

	integration.ProgramTest(t, &test)
}

func getTSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseTS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@aschot/pulumi-dbtcloud",
		},
	})

	return baseTS
}
