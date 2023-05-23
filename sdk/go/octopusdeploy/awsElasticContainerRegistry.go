// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages an AWS Elastic Container Registry in Octopus Deploy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pierskarsenbarg/pulumi-octopusdeploy/sdk/go/octopusdeploy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := octopusdeploy.NewAwsElasticContainerRegistry(ctx, "example", &octopusdeploy.AwsElasticContainerRegistryArgs{
//				AccessKey: pulumi.String("access-key"),
//				Region:    pulumi.String("us-east-1"),
//				SecretKey: pulumi.String("secret-key"),
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
//	$ pulumi import octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry [options] octopusdeploy_aws_elastic_container_registry.<name> <feed-id>
//
// ```
type AwsElasticContainerRegistry struct {
	pulumi.CustomResourceState

	// The AWS access key to use when authenticating against Amazon Web Services.
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// A short, memorable, unique name for this feed. Example: ACME Builds.
	Name                              pulumi.StringOutput      `pulumi:"name"`
	PackageAcquisitionLocationOptions pulumi.StringArrayOutput `pulumi:"packageAcquisitionLocationOptions"`
	// The AWS region where the registry resides.
	Region pulumi.StringOutput `pulumi:"region"`
	// The AWS secret key to use when authenticating against Amazon Web Services.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// The space ID associated with this feed.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
}

// NewAwsElasticContainerRegistry registers a new resource with the given unique name, arguments, and options.
func NewAwsElasticContainerRegistry(ctx *pulumi.Context,
	name string, args *AwsElasticContainerRegistryArgs, opts ...pulumi.ResourceOption) (*AwsElasticContainerRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessKey == nil {
		return nil, errors.New("invalid value for required argument 'AccessKey'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.SecretKey == nil {
		return nil, errors.New("invalid value for required argument 'SecretKey'")
	}
	if args.SecretKey != nil {
		args.SecretKey = pulumi.ToSecret(args.SecretKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretKey",
	})
	opts = append(opts, secrets)
	var resource AwsElasticContainerRegistry
	err := ctx.RegisterResource("octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAwsElasticContainerRegistry gets an existing AwsElasticContainerRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAwsElasticContainerRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsElasticContainerRegistryState, opts ...pulumi.ResourceOption) (*AwsElasticContainerRegistry, error) {
	var resource AwsElasticContainerRegistry
	err := ctx.ReadResource("octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsElasticContainerRegistry resources.
type awsElasticContainerRegistryState struct {
	// The AWS access key to use when authenticating against Amazon Web Services.
	AccessKey *string `pulumi:"accessKey"`
	// A short, memorable, unique name for this feed. Example: ACME Builds.
	Name                              *string  `pulumi:"name"`
	PackageAcquisitionLocationOptions []string `pulumi:"packageAcquisitionLocationOptions"`
	// The AWS region where the registry resides.
	Region *string `pulumi:"region"`
	// The AWS secret key to use when authenticating against Amazon Web Services.
	SecretKey *string `pulumi:"secretKey"`
	// The space ID associated with this feed.
	SpaceId *string `pulumi:"spaceId"`
}

type AwsElasticContainerRegistryState struct {
	// The AWS access key to use when authenticating against Amazon Web Services.
	AccessKey pulumi.StringPtrInput
	// A short, memorable, unique name for this feed. Example: ACME Builds.
	Name                              pulumi.StringPtrInput
	PackageAcquisitionLocationOptions pulumi.StringArrayInput
	// The AWS region where the registry resides.
	Region pulumi.StringPtrInput
	// The AWS secret key to use when authenticating against Amazon Web Services.
	SecretKey pulumi.StringPtrInput
	// The space ID associated with this feed.
	SpaceId pulumi.StringPtrInput
}

func (AwsElasticContainerRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsElasticContainerRegistryState)(nil)).Elem()
}

type awsElasticContainerRegistryArgs struct {
	// The AWS access key to use when authenticating against Amazon Web Services.
	AccessKey string `pulumi:"accessKey"`
	// A short, memorable, unique name for this feed. Example: ACME Builds.
	Name                              *string  `pulumi:"name"`
	PackageAcquisitionLocationOptions []string `pulumi:"packageAcquisitionLocationOptions"`
	// The AWS region where the registry resides.
	Region string `pulumi:"region"`
	// The AWS secret key to use when authenticating against Amazon Web Services.
	SecretKey string `pulumi:"secretKey"`
	// The space ID associated with this feed.
	SpaceId *string `pulumi:"spaceId"`
}

// The set of arguments for constructing a AwsElasticContainerRegistry resource.
type AwsElasticContainerRegistryArgs struct {
	// The AWS access key to use when authenticating against Amazon Web Services.
	AccessKey pulumi.StringInput
	// A short, memorable, unique name for this feed. Example: ACME Builds.
	Name                              pulumi.StringPtrInput
	PackageAcquisitionLocationOptions pulumi.StringArrayInput
	// The AWS region where the registry resides.
	Region pulumi.StringInput
	// The AWS secret key to use when authenticating against Amazon Web Services.
	SecretKey pulumi.StringInput
	// The space ID associated with this feed.
	SpaceId pulumi.StringPtrInput
}

func (AwsElasticContainerRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsElasticContainerRegistryArgs)(nil)).Elem()
}

type AwsElasticContainerRegistryInput interface {
	pulumi.Input

	ToAwsElasticContainerRegistryOutput() AwsElasticContainerRegistryOutput
	ToAwsElasticContainerRegistryOutputWithContext(ctx context.Context) AwsElasticContainerRegistryOutput
}

func (*AwsElasticContainerRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsElasticContainerRegistry)(nil)).Elem()
}

func (i *AwsElasticContainerRegistry) ToAwsElasticContainerRegistryOutput() AwsElasticContainerRegistryOutput {
	return i.ToAwsElasticContainerRegistryOutputWithContext(context.Background())
}

func (i *AwsElasticContainerRegistry) ToAwsElasticContainerRegistryOutputWithContext(ctx context.Context) AwsElasticContainerRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsElasticContainerRegistryOutput)
}

// AwsElasticContainerRegistryArrayInput is an input type that accepts AwsElasticContainerRegistryArray and AwsElasticContainerRegistryArrayOutput values.
// You can construct a concrete instance of `AwsElasticContainerRegistryArrayInput` via:
//
//	AwsElasticContainerRegistryArray{ AwsElasticContainerRegistryArgs{...} }
type AwsElasticContainerRegistryArrayInput interface {
	pulumi.Input

	ToAwsElasticContainerRegistryArrayOutput() AwsElasticContainerRegistryArrayOutput
	ToAwsElasticContainerRegistryArrayOutputWithContext(context.Context) AwsElasticContainerRegistryArrayOutput
}

type AwsElasticContainerRegistryArray []AwsElasticContainerRegistryInput

func (AwsElasticContainerRegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsElasticContainerRegistry)(nil)).Elem()
}

func (i AwsElasticContainerRegistryArray) ToAwsElasticContainerRegistryArrayOutput() AwsElasticContainerRegistryArrayOutput {
	return i.ToAwsElasticContainerRegistryArrayOutputWithContext(context.Background())
}

func (i AwsElasticContainerRegistryArray) ToAwsElasticContainerRegistryArrayOutputWithContext(ctx context.Context) AwsElasticContainerRegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsElasticContainerRegistryArrayOutput)
}

// AwsElasticContainerRegistryMapInput is an input type that accepts AwsElasticContainerRegistryMap and AwsElasticContainerRegistryMapOutput values.
// You can construct a concrete instance of `AwsElasticContainerRegistryMapInput` via:
//
//	AwsElasticContainerRegistryMap{ "key": AwsElasticContainerRegistryArgs{...} }
type AwsElasticContainerRegistryMapInput interface {
	pulumi.Input

	ToAwsElasticContainerRegistryMapOutput() AwsElasticContainerRegistryMapOutput
	ToAwsElasticContainerRegistryMapOutputWithContext(context.Context) AwsElasticContainerRegistryMapOutput
}

type AwsElasticContainerRegistryMap map[string]AwsElasticContainerRegistryInput

func (AwsElasticContainerRegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsElasticContainerRegistry)(nil)).Elem()
}

func (i AwsElasticContainerRegistryMap) ToAwsElasticContainerRegistryMapOutput() AwsElasticContainerRegistryMapOutput {
	return i.ToAwsElasticContainerRegistryMapOutputWithContext(context.Background())
}

func (i AwsElasticContainerRegistryMap) ToAwsElasticContainerRegistryMapOutputWithContext(ctx context.Context) AwsElasticContainerRegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsElasticContainerRegistryMapOutput)
}

type AwsElasticContainerRegistryOutput struct{ *pulumi.OutputState }

func (AwsElasticContainerRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsElasticContainerRegistry)(nil)).Elem()
}

func (o AwsElasticContainerRegistryOutput) ToAwsElasticContainerRegistryOutput() AwsElasticContainerRegistryOutput {
	return o
}

func (o AwsElasticContainerRegistryOutput) ToAwsElasticContainerRegistryOutputWithContext(ctx context.Context) AwsElasticContainerRegistryOutput {
	return o
}

// The AWS access key to use when authenticating against Amazon Web Services.
func (o AwsElasticContainerRegistryOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsElasticContainerRegistry) pulumi.StringOutput { return v.AccessKey }).(pulumi.StringOutput)
}

// A short, memorable, unique name for this feed. Example: ACME Builds.
func (o AwsElasticContainerRegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsElasticContainerRegistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AwsElasticContainerRegistryOutput) PackageAcquisitionLocationOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AwsElasticContainerRegistry) pulumi.StringArrayOutput {
		return v.PackageAcquisitionLocationOptions
	}).(pulumi.StringArrayOutput)
}

// The AWS region where the registry resides.
func (o AwsElasticContainerRegistryOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsElasticContainerRegistry) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The AWS secret key to use when authenticating against Amazon Web Services.
func (o AwsElasticContainerRegistryOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsElasticContainerRegistry) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

// The space ID associated with this feed.
func (o AwsElasticContainerRegistryOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsElasticContainerRegistry) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

type AwsElasticContainerRegistryArrayOutput struct{ *pulumi.OutputState }

func (AwsElasticContainerRegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsElasticContainerRegistry)(nil)).Elem()
}

func (o AwsElasticContainerRegistryArrayOutput) ToAwsElasticContainerRegistryArrayOutput() AwsElasticContainerRegistryArrayOutput {
	return o
}

func (o AwsElasticContainerRegistryArrayOutput) ToAwsElasticContainerRegistryArrayOutputWithContext(ctx context.Context) AwsElasticContainerRegistryArrayOutput {
	return o
}

func (o AwsElasticContainerRegistryArrayOutput) Index(i pulumi.IntInput) AwsElasticContainerRegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AwsElasticContainerRegistry {
		return vs[0].([]*AwsElasticContainerRegistry)[vs[1].(int)]
	}).(AwsElasticContainerRegistryOutput)
}

type AwsElasticContainerRegistryMapOutput struct{ *pulumi.OutputState }

func (AwsElasticContainerRegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsElasticContainerRegistry)(nil)).Elem()
}

func (o AwsElasticContainerRegistryMapOutput) ToAwsElasticContainerRegistryMapOutput() AwsElasticContainerRegistryMapOutput {
	return o
}

func (o AwsElasticContainerRegistryMapOutput) ToAwsElasticContainerRegistryMapOutputWithContext(ctx context.Context) AwsElasticContainerRegistryMapOutput {
	return o
}

func (o AwsElasticContainerRegistryMapOutput) MapIndex(k pulumi.StringInput) AwsElasticContainerRegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AwsElasticContainerRegistry {
		return vs[0].(map[string]*AwsElasticContainerRegistry)[vs[1].(string)]
	}).(AwsElasticContainerRegistryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AwsElasticContainerRegistryInput)(nil)).Elem(), &AwsElasticContainerRegistry{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsElasticContainerRegistryArrayInput)(nil)).Elem(), AwsElasticContainerRegistryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsElasticContainerRegistryMapInput)(nil)).Elem(), AwsElasticContainerRegistryMap{})
	pulumi.RegisterOutputType(AwsElasticContainerRegistryOutput{})
	pulumi.RegisterOutputType(AwsElasticContainerRegistryArrayOutput{})
	pulumi.RegisterOutputType(AwsElasticContainerRegistryMapOutput{})
}
