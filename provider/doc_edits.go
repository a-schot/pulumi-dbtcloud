package dbtcloud

import (
	"bytes"
	"regexp"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
)

const preProjectExample string = `data "dbtcloud_project" "test_project" {
  name = "My project name"
}`

const postProjectExample string = `data "dbtcloud_project" "another_test_project" {
  name = "My other project name"
}`

func editRules(defaults []tfbridge.DocsEdit) []tfbridge.DocsEdit {
	return append(
		defaults,
		simpleReplace("appplication_id", "applicationId"),
		simpleReplace("appplication_id", "application_id"),
		simpleReplace(preProjectExample, postProjectExample),
		simpleReplace(
			"```console\ndbt_cloud_api_result",
			"```javascript\ndbt_cloud_api_result",
		),

		// remove the header which warns for legacy dbt_cloud_ prefixed resources
		reReplace("// use dbt.*\n// legacy names will be removed from 0.3 onwards\n\n", ""),
		// reReplace("(?s)run the following commands in the Google Chrome console:.*Alternatively, you can ", ""),
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

func reReplace(from string, to string) tfbridge.DocsEdit {
	fromR, toB := regexp.MustCompile(from), []byte(to)
	return tfbridge.DocsEdit{
		Path: "*",
		Edit: func(_ string, content []byte) ([]byte, error) {
			return fromR.ReplaceAll(content, toB), nil
		},
	}
}
