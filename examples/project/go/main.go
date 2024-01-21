package main

import (
	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		project, err := dbtcloud.NewProject(ctx, "go-project", nil)
		if err != nil {
			return err
		}
		ctx.Export("project_name", project.Name)
		return nil
	})
}
