// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages channels in Octopus Deploy.
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
//			_, err := octopusdeploy.NewChannel(ctx, "example", &octopusdeploy.ChannelArgs{
//				ProjectId: pulumi.String("Projects-123"),
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
//	$ pulumi import octopusdeploy:index/channel:Channel [options] octopusdeploy_channel.<name> <channel-id>
//
// ```
type Channel struct {
	pulumi.CustomResourceState

	// The description of this channel.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates if this is the default channel for the associated project.
	IsDefault pulumi.BoolPtrOutput `pulumi:"isDefault"`
	// The lifecycle ID associated with this channel.
	LifecycleId pulumi.StringPtrOutput `pulumi:"lifecycleId"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project ID associated with this channel.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// A list of rules associated with this channel.
	Rules ChannelRuleArrayOutput `pulumi:"rules"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayOutput `pulumi:"tenantTags"`
}

// NewChannel registers a new resource with the given unique name, arguments, and options.
func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource Channel
	err := ctx.RegisterResource("octopusdeploy:index/channel:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannel gets an existing Channel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("octopusdeploy:index/channel:Channel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Channel resources.
type channelState struct {
	// The description of this channel.
	Description *string `pulumi:"description"`
	// Indicates if this is the default channel for the associated project.
	IsDefault *bool `pulumi:"isDefault"`
	// The lifecycle ID associated with this channel.
	LifecycleId *string `pulumi:"lifecycleId"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The project ID associated with this channel.
	ProjectId *string `pulumi:"projectId"`
	// A list of rules associated with this channel.
	Rules []ChannelRule `pulumi:"rules"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
}

type ChannelState struct {
	// The description of this channel.
	Description pulumi.StringPtrInput
	// Indicates if this is the default channel for the associated project.
	IsDefault pulumi.BoolPtrInput
	// The lifecycle ID associated with this channel.
	LifecycleId pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The project ID associated with this channel.
	ProjectId pulumi.StringPtrInput
	// A list of rules associated with this channel.
	Rules ChannelRuleArrayInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
}

func (ChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelState)(nil)).Elem()
}

type channelArgs struct {
	// The description of this channel.
	Description *string `pulumi:"description"`
	// Indicates if this is the default channel for the associated project.
	IsDefault *bool `pulumi:"isDefault"`
	// The lifecycle ID associated with this channel.
	LifecycleId *string `pulumi:"lifecycleId"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The project ID associated with this channel.
	ProjectId string `pulumi:"projectId"`
	// A list of rules associated with this channel.
	Rules []ChannelRule `pulumi:"rules"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
}

// The set of arguments for constructing a Channel resource.
type ChannelArgs struct {
	// The description of this channel.
	Description pulumi.StringPtrInput
	// Indicates if this is the default channel for the associated project.
	IsDefault pulumi.BoolPtrInput
	// The lifecycle ID associated with this channel.
	LifecycleId pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The project ID associated with this channel.
	ProjectId pulumi.StringInput
	// A list of rules associated with this channel.
	Rules ChannelRuleArrayInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
}

func (ChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelArgs)(nil)).Elem()
}

type ChannelInput interface {
	pulumi.Input

	ToChannelOutput() ChannelOutput
	ToChannelOutputWithContext(ctx context.Context) ChannelOutput
}

func (*Channel) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (i *Channel) ToChannelOutput() ChannelOutput {
	return i.ToChannelOutputWithContext(context.Background())
}

func (i *Channel) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelOutput)
}

// ChannelArrayInput is an input type that accepts ChannelArray and ChannelArrayOutput values.
// You can construct a concrete instance of `ChannelArrayInput` via:
//
//	ChannelArray{ ChannelArgs{...} }
type ChannelArrayInput interface {
	pulumi.Input

	ToChannelArrayOutput() ChannelArrayOutput
	ToChannelArrayOutputWithContext(context.Context) ChannelArrayOutput
}

type ChannelArray []ChannelInput

func (ChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Channel)(nil)).Elem()
}

func (i ChannelArray) ToChannelArrayOutput() ChannelArrayOutput {
	return i.ToChannelArrayOutputWithContext(context.Background())
}

func (i ChannelArray) ToChannelArrayOutputWithContext(ctx context.Context) ChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelArrayOutput)
}

// ChannelMapInput is an input type that accepts ChannelMap and ChannelMapOutput values.
// You can construct a concrete instance of `ChannelMapInput` via:
//
//	ChannelMap{ "key": ChannelArgs{...} }
type ChannelMapInput interface {
	pulumi.Input

	ToChannelMapOutput() ChannelMapOutput
	ToChannelMapOutputWithContext(context.Context) ChannelMapOutput
}

type ChannelMap map[string]ChannelInput

func (ChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Channel)(nil)).Elem()
}

func (i ChannelMap) ToChannelMapOutput() ChannelMapOutput {
	return i.ToChannelMapOutputWithContext(context.Background())
}

func (i ChannelMap) ToChannelMapOutputWithContext(ctx context.Context) ChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelMapOutput)
}

type ChannelOutput struct{ *pulumi.OutputState }

func (ChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (o ChannelOutput) ToChannelOutput() ChannelOutput {
	return o
}

func (o ChannelOutput) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return o
}

// The description of this channel.
func (o ChannelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates if this is the default channel for the associated project.
func (o ChannelOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.BoolPtrOutput { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// The lifecycle ID associated with this channel.
func (o ChannelOutput) LifecycleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.LifecycleId }).(pulumi.StringPtrOutput)
}

// The name of this resource.
func (o ChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project ID associated with this channel.
func (o ChannelOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// A list of rules associated with this channel.
func (o ChannelOutput) Rules() ChannelRuleArrayOutput {
	return o.ApplyT(func(v *Channel) ChannelRuleArrayOutput { return v.Rules }).(ChannelRuleArrayOutput)
}

// The space ID associated with this resource.
func (o ChannelOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// A list of tenant tags associated with this resource.
func (o ChannelOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringArrayOutput { return v.TenantTags }).(pulumi.StringArrayOutput)
}

type ChannelArrayOutput struct{ *pulumi.OutputState }

func (ChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Channel)(nil)).Elem()
}

func (o ChannelArrayOutput) ToChannelArrayOutput() ChannelArrayOutput {
	return o
}

func (o ChannelArrayOutput) ToChannelArrayOutputWithContext(ctx context.Context) ChannelArrayOutput {
	return o
}

func (o ChannelArrayOutput) Index(i pulumi.IntInput) ChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Channel {
		return vs[0].([]*Channel)[vs[1].(int)]
	}).(ChannelOutput)
}

type ChannelMapOutput struct{ *pulumi.OutputState }

func (ChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Channel)(nil)).Elem()
}

func (o ChannelMapOutput) ToChannelMapOutput() ChannelMapOutput {
	return o
}

func (o ChannelMapOutput) ToChannelMapOutputWithContext(ctx context.Context) ChannelMapOutput {
	return o
}

func (o ChannelMapOutput) MapIndex(k pulumi.StringInput) ChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Channel {
		return vs[0].(map[string]*Channel)[vs[1].(string)]
	}).(ChannelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelInput)(nil)).Elem(), &Channel{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelArrayInput)(nil)).Elem(), ChannelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelMapInput)(nil)).Elem(), ChannelMap{})
	pulumi.RegisterOutputType(ChannelOutput{})
	pulumi.RegisterOutputType(ChannelArrayOutput{})
	pulumi.RegisterOutputType(ChannelMapOutput{})
}
