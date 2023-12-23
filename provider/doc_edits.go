package dbtcloud

import (
	"bytes"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
)

func editRules(defaults []tfbridge.DocsEdit) []tfbridge.DocsEdit {
	return append(
		defaults,
		simpleReplace(`appplicationId`, `applicationId`),
	)
}

func simpleReplace(from, to string) tfbridge.DocsEdit {
	fromB, toB := []byte(from), []byte(to)
	return tfbridge.DocsEdit{
		Path: "*",
		Edit: func(_ string, content []byte) ([]byte, error) {
			return bytes.ReplaceAll(content, fromB, toB), nil
		},
	}
}
