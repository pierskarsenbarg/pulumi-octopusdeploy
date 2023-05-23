// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages AWS accounts in Octopus Deploy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-octopusdeploy/sdk/go/octopusdeploy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := octopusdeploy.NewAwsAccount(ctx, "example", &octopusdeploy.AwsAccountArgs{
//				AccessKey: pulumi.String("access-key"),
//				SecretKey: pulumi.String("###########"),
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
// ```sh
//
//	$ pulumi import octopusdeploy:index/awsAccount:AwsAccount [options] octopusdeploy_aws_account.<name> <account-id>
//
// ```
type AwsAccount struct {
	pulumi.CustomResourceState

	// The access key associated with this AWS account.
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// A user-friendly description of this AWS account.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of environment IDs associated with this resource.
	Environments pulumi.StringArrayOutput `pulumi:"environments"`
	// The name of this AWS account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The secret key associated with this resource.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayOutput `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringOutput `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants pulumi.StringArrayOutput `pulumi:"tenants"`
}

// NewAwsAccount registers a new resource with the given unique name, arguments, and options.
func NewAwsAccount(ctx *pulumi.Context,
	name string, args *AwsAccountArgs, opts ...pulumi.ResourceOption) (*AwsAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessKey == nil {
		return nil, errors.New("invalid value for required argument 'AccessKey'")
	}
	if args.SecretKey == nil {
		return nil, errors.New("invalid value for required argument 'SecretKey'")
	}
	var resource AwsAccount
	err := ctx.RegisterResource("octopusdeploy:index/awsAccount:AwsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAwsAccount gets an existing AwsAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAwsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsAccountState, opts ...pulumi.ResourceOption) (*AwsAccount, error) {
	var resource AwsAccount
	err := ctx.ReadResource("octopusdeploy:index/awsAccount:AwsAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsAccount resources.
type awsAccountState struct {
	// The access key associated with this AWS account.
	AccessKey *string `pulumi:"accessKey"`
	// A user-friendly description of this AWS account.
	Description *string `pulumi:"description"`
	// A list of environment IDs associated with this resource.
	Environments []string `pulumi:"environments"`
	// The name of this AWS account.
	Name *string `pulumi:"name"`
	// The secret key associated with this resource.
	SecretKey *string `pulumi:"secretKey"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants []string `pulumi:"tenants"`
}

type AwsAccountState struct {
	// The access key associated with this AWS account.
	AccessKey pulumi.StringPtrInput
	// A user-friendly description of this AWS account.
	Description pulumi.StringPtrInput
	// A list of environment IDs associated with this resource.
	Environments pulumi.StringArrayInput
	// The name of this AWS account.
	Name pulumi.StringPtrInput
	// The secret key associated with this resource.
	SecretKey pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants pulumi.StringArrayInput
}

func (AwsAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsAccountState)(nil)).Elem()
}

type awsAccountArgs struct {
	// The access key associated with this AWS account.
	AccessKey string `pulumi:"accessKey"`
	// A user-friendly description of this AWS account.
	Description *string `pulumi:"description"`
	// A list of environment IDs associated with this resource.
	Environments []string `pulumi:"environments"`
	// The name of this AWS account.
	Name *string `pulumi:"name"`
	// The secret key associated with this resource.
	SecretKey string `pulumi:"secretKey"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants []string `pulumi:"tenants"`
}

// The set of arguments for constructing a AwsAccount resource.
type AwsAccountArgs struct {
	// The access key associated with this AWS account.
	AccessKey pulumi.StringInput
	// A user-friendly description of this AWS account.
	Description pulumi.StringPtrInput
	// A list of environment IDs associated with this resource.
	Environments pulumi.StringArrayInput
	// The name of this AWS account.
	Name pulumi.StringPtrInput
	// The secret key associated with this resource.
	SecretKey pulumi.StringInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants pulumi.StringArrayInput
}

func (AwsAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsAccountArgs)(nil)).Elem()
}

type AwsAccountInput interface {
	pulumi.Input

	ToAwsAccountOutput() AwsAccountOutput
	ToAwsAccountOutputWithContext(ctx context.Context) AwsAccountOutput
}

func (*AwsAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsAccount)(nil)).Elem()
}

func (i *AwsAccount) ToAwsAccountOutput() AwsAccountOutput {
	return i.ToAwsAccountOutputWithContext(context.Background())
}

func (i *AwsAccount) ToAwsAccountOutputWithContext(ctx context.Context) AwsAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsAccountOutput)
}

// AwsAccountArrayInput is an input type that accepts AwsAccountArray and AwsAccountArrayOutput values.
// You can construct a concrete instance of `AwsAccountArrayInput` via:
//
//	AwsAccountArray{ AwsAccountArgs{...} }
type AwsAccountArrayInput interface {
	pulumi.Input

	ToAwsAccountArrayOutput() AwsAccountArrayOutput
	ToAwsAccountArrayOutputWithContext(context.Context) AwsAccountArrayOutput
}

type AwsAccountArray []AwsAccountInput

func (AwsAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsAccount)(nil)).Elem()
}

func (i AwsAccountArray) ToAwsAccountArrayOutput() AwsAccountArrayOutput {
	return i.ToAwsAccountArrayOutputWithContext(context.Background())
}

func (i AwsAccountArray) ToAwsAccountArrayOutputWithContext(ctx context.Context) AwsAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsAccountArrayOutput)
}

// AwsAccountMapInput is an input type that accepts AwsAccountMap and AwsAccountMapOutput values.
// You can construct a concrete instance of `AwsAccountMapInput` via:
//
//	AwsAccountMap{ "key": AwsAccountArgs{...} }
type AwsAccountMapInput interface {
	pulumi.Input

	ToAwsAccountMapOutput() AwsAccountMapOutput
	ToAwsAccountMapOutputWithContext(context.Context) AwsAccountMapOutput
}

type AwsAccountMap map[string]AwsAccountInput

func (AwsAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsAccount)(nil)).Elem()
}

func (i AwsAccountMap) ToAwsAccountMapOutput() AwsAccountMapOutput {
	return i.ToAwsAccountMapOutputWithContext(context.Background())
}

func (i AwsAccountMap) ToAwsAccountMapOutputWithContext(ctx context.Context) AwsAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsAccountMapOutput)
}

type AwsAccountOutput struct{ *pulumi.OutputState }

func (AwsAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsAccount)(nil)).Elem()
}

func (o AwsAccountOutput) ToAwsAccountOutput() AwsAccountOutput {
	return o
}

func (o AwsAccountOutput) ToAwsAccountOutputWithContext(ctx context.Context) AwsAccountOutput {
	return o
}

// The access key associated with this AWS account.
func (o AwsAccountOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsAccount) pulumi.StringOutput { return v.AccessKey }).(pulumi.StringOutput)
}

// A user-friendly description of this AWS account.
func (o AwsAccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsAccount) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A list of environment IDs associated with this resource.
func (o AwsAccountOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AwsAccount) pulumi.StringArrayOutput { return v.Environments }).(pulumi.StringArrayOutput)
}

// The name of this AWS account.
func (o AwsAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The secret key associated with this resource.
func (o AwsAccountOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsAccount) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

// The space ID associated with this resource.
func (o AwsAccountOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsAccount) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// A list of tenant tags associated with this resource.
func (o AwsAccountOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AwsAccount) pulumi.StringArrayOutput { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
func (o AwsAccountOutput) TenantedDeploymentParticipation() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsAccount) pulumi.StringOutput { return v.TenantedDeploymentParticipation }).(pulumi.StringOutput)
}

// A list of tenant IDs associated with this resource.
func (o AwsAccountOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AwsAccount) pulumi.StringArrayOutput { return v.Tenants }).(pulumi.StringArrayOutput)
}

type AwsAccountArrayOutput struct{ *pulumi.OutputState }

func (AwsAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsAccount)(nil)).Elem()
}

func (o AwsAccountArrayOutput) ToAwsAccountArrayOutput() AwsAccountArrayOutput {
	return o
}

func (o AwsAccountArrayOutput) ToAwsAccountArrayOutputWithContext(ctx context.Context) AwsAccountArrayOutput {
	return o
}

func (o AwsAccountArrayOutput) Index(i pulumi.IntInput) AwsAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AwsAccount {
		return vs[0].([]*AwsAccount)[vs[1].(int)]
	}).(AwsAccountOutput)
}

type AwsAccountMapOutput struct{ *pulumi.OutputState }

func (AwsAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsAccount)(nil)).Elem()
}

func (o AwsAccountMapOutput) ToAwsAccountMapOutput() AwsAccountMapOutput {
	return o
}

func (o AwsAccountMapOutput) ToAwsAccountMapOutputWithContext(ctx context.Context) AwsAccountMapOutput {
	return o
}

func (o AwsAccountMapOutput) MapIndex(k pulumi.StringInput) AwsAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AwsAccount {
		return vs[0].(map[string]*AwsAccount)[vs[1].(string)]
	}).(AwsAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AwsAccountInput)(nil)).Elem(), &AwsAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsAccountArrayInput)(nil)).Elem(), AwsAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsAccountMapInput)(nil)).Elem(), AwsAccountMap{})
	pulumi.RegisterOutputType(AwsAccountOutput{})
	pulumi.RegisterOutputType(AwsAccountArrayOutput{})
	pulumi.RegisterOutputType(AwsAccountMapOutput{})
}