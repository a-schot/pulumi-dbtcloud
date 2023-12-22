// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dbtcloud.NewProjectRepository(ctx, "dbtProjectRepository", &dbtcloud.ProjectRepositoryArgs{
//				ProjectId:    pulumi.Any(dbtcloud_project.Dbt_project.Id),
//				RepositoryId: pulumi.Any(dbtcloud_repository.Dbt_repository.Repository_id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Import using a project ID and Connection ID found in the URL or via the API.
//
// ```sh
//
//	$ pulumi import dbtcloud:index/projectRepository:ProjectRepository my_project "project_id:repository_id"
//
// ```
//
// ```sh
//
//	$ pulumi import dbtcloud:index/projectRepository:ProjectRepository my_project 12345:5678
//
// ```
type ProjectRepository struct {
	pulumi.CustomResourceState

	// Project ID
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// Repository ID
	RepositoryId pulumi.IntOutput `pulumi:"repositoryId"`
}

// NewProjectRepository registers a new resource with the given unique name, arguments, and options.
func NewProjectRepository(ctx *pulumi.Context,
	name string, args *ProjectRepositoryArgs, opts ...pulumi.ResourceOption) (*ProjectRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectRepository
	err := ctx.RegisterResource("dbtcloud:index/projectRepository:ProjectRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectRepository gets an existing ProjectRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectRepositoryState, opts ...pulumi.ResourceOption) (*ProjectRepository, error) {
	var resource ProjectRepository
	err := ctx.ReadResource("dbtcloud:index/projectRepository:ProjectRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectRepository resources.
type projectRepositoryState struct {
	// Project ID
	ProjectId *int `pulumi:"projectId"`
	// Repository ID
	RepositoryId *int `pulumi:"repositoryId"`
}

type ProjectRepositoryState struct {
	// Project ID
	ProjectId pulumi.IntPtrInput
	// Repository ID
	RepositoryId pulumi.IntPtrInput
}

func (ProjectRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectRepositoryState)(nil)).Elem()
}

type projectRepositoryArgs struct {
	// Project ID
	ProjectId int `pulumi:"projectId"`
	// Repository ID
	RepositoryId int `pulumi:"repositoryId"`
}

// The set of arguments for constructing a ProjectRepository resource.
type ProjectRepositoryArgs struct {
	// Project ID
	ProjectId pulumi.IntInput
	// Repository ID
	RepositoryId pulumi.IntInput
}

func (ProjectRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectRepositoryArgs)(nil)).Elem()
}

type ProjectRepositoryInput interface {
	pulumi.Input

	ToProjectRepositoryOutput() ProjectRepositoryOutput
	ToProjectRepositoryOutputWithContext(ctx context.Context) ProjectRepositoryOutput
}

func (*ProjectRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectRepository)(nil)).Elem()
}

func (i *ProjectRepository) ToProjectRepositoryOutput() ProjectRepositoryOutput {
	return i.ToProjectRepositoryOutputWithContext(context.Background())
}

func (i *ProjectRepository) ToProjectRepositoryOutputWithContext(ctx context.Context) ProjectRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectRepositoryOutput)
}

// ProjectRepositoryArrayInput is an input type that accepts ProjectRepositoryArray and ProjectRepositoryArrayOutput values.
// You can construct a concrete instance of `ProjectRepositoryArrayInput` via:
//
//	ProjectRepositoryArray{ ProjectRepositoryArgs{...} }
type ProjectRepositoryArrayInput interface {
	pulumi.Input

	ToProjectRepositoryArrayOutput() ProjectRepositoryArrayOutput
	ToProjectRepositoryArrayOutputWithContext(context.Context) ProjectRepositoryArrayOutput
}

type ProjectRepositoryArray []ProjectRepositoryInput

func (ProjectRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectRepository)(nil)).Elem()
}

func (i ProjectRepositoryArray) ToProjectRepositoryArrayOutput() ProjectRepositoryArrayOutput {
	return i.ToProjectRepositoryArrayOutputWithContext(context.Background())
}

func (i ProjectRepositoryArray) ToProjectRepositoryArrayOutputWithContext(ctx context.Context) ProjectRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectRepositoryArrayOutput)
}

// ProjectRepositoryMapInput is an input type that accepts ProjectRepositoryMap and ProjectRepositoryMapOutput values.
// You can construct a concrete instance of `ProjectRepositoryMapInput` via:
//
//	ProjectRepositoryMap{ "key": ProjectRepositoryArgs{...} }
type ProjectRepositoryMapInput interface {
	pulumi.Input

	ToProjectRepositoryMapOutput() ProjectRepositoryMapOutput
	ToProjectRepositoryMapOutputWithContext(context.Context) ProjectRepositoryMapOutput
}

type ProjectRepositoryMap map[string]ProjectRepositoryInput

func (ProjectRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectRepository)(nil)).Elem()
}

func (i ProjectRepositoryMap) ToProjectRepositoryMapOutput() ProjectRepositoryMapOutput {
	return i.ToProjectRepositoryMapOutputWithContext(context.Background())
}

func (i ProjectRepositoryMap) ToProjectRepositoryMapOutputWithContext(ctx context.Context) ProjectRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectRepositoryMapOutput)
}

type ProjectRepositoryOutput struct{ *pulumi.OutputState }

func (ProjectRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectRepository)(nil)).Elem()
}

func (o ProjectRepositoryOutput) ToProjectRepositoryOutput() ProjectRepositoryOutput {
	return o
}

func (o ProjectRepositoryOutput) ToProjectRepositoryOutputWithContext(ctx context.Context) ProjectRepositoryOutput {
	return o
}

// Project ID
func (o ProjectRepositoryOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectRepository) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// Repository ID
func (o ProjectRepositoryOutput) RepositoryId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectRepository) pulumi.IntOutput { return v.RepositoryId }).(pulumi.IntOutput)
}

type ProjectRepositoryArrayOutput struct{ *pulumi.OutputState }

func (ProjectRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectRepository)(nil)).Elem()
}

func (o ProjectRepositoryArrayOutput) ToProjectRepositoryArrayOutput() ProjectRepositoryArrayOutput {
	return o
}

func (o ProjectRepositoryArrayOutput) ToProjectRepositoryArrayOutputWithContext(ctx context.Context) ProjectRepositoryArrayOutput {
	return o
}

func (o ProjectRepositoryArrayOutput) Index(i pulumi.IntInput) ProjectRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectRepository {
		return vs[0].([]*ProjectRepository)[vs[1].(int)]
	}).(ProjectRepositoryOutput)
}

type ProjectRepositoryMapOutput struct{ *pulumi.OutputState }

func (ProjectRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectRepository)(nil)).Elem()
}

func (o ProjectRepositoryMapOutput) ToProjectRepositoryMapOutput() ProjectRepositoryMapOutput {
	return o
}

func (o ProjectRepositoryMapOutput) ToProjectRepositoryMapOutputWithContext(ctx context.Context) ProjectRepositoryMapOutput {
	return o
}

func (o ProjectRepositoryMapOutput) MapIndex(k pulumi.StringInput) ProjectRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectRepository {
		return vs[0].(map[string]*ProjectRepository)[vs[1].(string)]
	}).(ProjectRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectRepositoryInput)(nil)).Elem(), &ProjectRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectRepositoryArrayInput)(nil)).Elem(), ProjectRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectRepositoryMapInput)(nil)).Elem(), ProjectRepositoryMap{})
	pulumi.RegisterOutputType(ProjectRepositoryOutput{})
	pulumi.RegisterOutputType(ProjectRepositoryArrayOutput{})
	pulumi.RegisterOutputType(ProjectRepositoryMapOutput{})
}