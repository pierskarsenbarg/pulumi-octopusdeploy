// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages a Helm feed in Octopus Deploy.
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
//			_, err := octopusdeploy.NewHelmFeed(ctx, "example", &octopusdeploy.HelmFeedArgs{
//				FeedUri:  pulumi.String("https://charts.helm.sh/stable"),
//				Password: pulumi.String("test-password"),
//				Username: pulumi.String("test-username"),
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
//	$ pulumi import octopusdeploy:index/helmFeed:HelmFeed [options] octopusdeploy_helm_feed.<name> <feed-id>
//
// ```
type HelmFeed struct {
	pulumi.CustomResourceState

	FeedUri pulumi.StringOutput `pulumi:"feedUri"`
	// A short, memorable, unique name for this feed. Example: ACME Builds.
	Name                              pulumi.StringOutput      `pulumi:"name"`
	PackageAcquisitionLocationOptions pulumi.StringArrayOutput `pulumi:"packageAcquisitionLocationOptions"`
	// The password associated with this resource.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// The username associated with this resource.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewHelmFeed registers a new resource with the given unique name, arguments, and options.
func NewHelmFeed(ctx *pulumi.Context,
	name string, args *HelmFeedArgs, opts ...pulumi.ResourceOption) (*HelmFeed, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FeedUri == nil {
		return nil, errors.New("invalid value for required argument 'FeedUri'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	if args.Username != nil {
		args.Username = pulumi.ToSecret(args.Username).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
		"username",
	})
	opts = append(opts, secrets)
	var resource HelmFeed
	err := ctx.RegisterResource("octopusdeploy:index/helmFeed:HelmFeed", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHelmFeed gets an existing HelmFeed resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHelmFeed(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HelmFeedState, opts ...pulumi.ResourceOption) (*HelmFeed, error) {
	var resource HelmFeed
	err := ctx.ReadResource("octopusdeploy:index/helmFeed:HelmFeed", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HelmFeed resources.
type helmFeedState struct {
	FeedUri *string `pulumi:"feedUri"`
	// A short, memorable, unique name for this feed. Example: ACME Builds.
	Name                              *string  `pulumi:"name"`
	PackageAcquisitionLocationOptions []string `pulumi:"packageAcquisitionLocationOptions"`
	// The password associated with this resource.
	Password *string `pulumi:"password"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// The username associated with this resource.
	Username *string `pulumi:"username"`
}

type HelmFeedState struct {
	FeedUri pulumi.StringPtrInput
	// A short, memorable, unique name for this feed. Example: ACME Builds.
	Name                              pulumi.StringPtrInput
	PackageAcquisitionLocationOptions pulumi.StringArrayInput
	// The password associated with this resource.
	Password pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// The username associated with this resource.
	Username pulumi.StringPtrInput
}

func (HelmFeedState) ElementType() reflect.Type {
	return reflect.TypeOf((*helmFeedState)(nil)).Elem()
}

type helmFeedArgs struct {
	FeedUri string `pulumi:"feedUri"`
	// A short, memorable, unique name for this feed. Example: ACME Builds.
	Name                              *string  `pulumi:"name"`
	PackageAcquisitionLocationOptions []string `pulumi:"packageAcquisitionLocationOptions"`
	// The password associated with this resource.
	Password *string `pulumi:"password"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// The username associated with this resource.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a HelmFeed resource.
type HelmFeedArgs struct {
	FeedUri pulumi.StringInput
	// A short, memorable, unique name for this feed. Example: ACME Builds.
	Name                              pulumi.StringPtrInput
	PackageAcquisitionLocationOptions pulumi.StringArrayInput
	// The password associated with this resource.
	Password pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// The username associated with this resource.
	Username pulumi.StringPtrInput
}

func (HelmFeedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*helmFeedArgs)(nil)).Elem()
}

type HelmFeedInput interface {
	pulumi.Input

	ToHelmFeedOutput() HelmFeedOutput
	ToHelmFeedOutputWithContext(ctx context.Context) HelmFeedOutput
}

func (*HelmFeed) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmFeed)(nil)).Elem()
}

func (i *HelmFeed) ToHelmFeedOutput() HelmFeedOutput {
	return i.ToHelmFeedOutputWithContext(context.Background())
}

func (i *HelmFeed) ToHelmFeedOutputWithContext(ctx context.Context) HelmFeedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmFeedOutput)
}

// HelmFeedArrayInput is an input type that accepts HelmFeedArray and HelmFeedArrayOutput values.
// You can construct a concrete instance of `HelmFeedArrayInput` via:
//
//	HelmFeedArray{ HelmFeedArgs{...} }
type HelmFeedArrayInput interface {
	pulumi.Input

	ToHelmFeedArrayOutput() HelmFeedArrayOutput
	ToHelmFeedArrayOutputWithContext(context.Context) HelmFeedArrayOutput
}

type HelmFeedArray []HelmFeedInput

func (HelmFeedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HelmFeed)(nil)).Elem()
}

func (i HelmFeedArray) ToHelmFeedArrayOutput() HelmFeedArrayOutput {
	return i.ToHelmFeedArrayOutputWithContext(context.Background())
}

func (i HelmFeedArray) ToHelmFeedArrayOutputWithContext(ctx context.Context) HelmFeedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmFeedArrayOutput)
}

// HelmFeedMapInput is an input type that accepts HelmFeedMap and HelmFeedMapOutput values.
// You can construct a concrete instance of `HelmFeedMapInput` via:
//
//	HelmFeedMap{ "key": HelmFeedArgs{...} }
type HelmFeedMapInput interface {
	pulumi.Input

	ToHelmFeedMapOutput() HelmFeedMapOutput
	ToHelmFeedMapOutputWithContext(context.Context) HelmFeedMapOutput
}

type HelmFeedMap map[string]HelmFeedInput

func (HelmFeedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HelmFeed)(nil)).Elem()
}

func (i HelmFeedMap) ToHelmFeedMapOutput() HelmFeedMapOutput {
	return i.ToHelmFeedMapOutputWithContext(context.Background())
}

func (i HelmFeedMap) ToHelmFeedMapOutputWithContext(ctx context.Context) HelmFeedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmFeedMapOutput)
}

type HelmFeedOutput struct{ *pulumi.OutputState }

func (HelmFeedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmFeed)(nil)).Elem()
}

func (o HelmFeedOutput) ToHelmFeedOutput() HelmFeedOutput {
	return o
}

func (o HelmFeedOutput) ToHelmFeedOutputWithContext(ctx context.Context) HelmFeedOutput {
	return o
}

func (o HelmFeedOutput) FeedUri() pulumi.StringOutput {
	return o.ApplyT(func(v *HelmFeed) pulumi.StringOutput { return v.FeedUri }).(pulumi.StringOutput)
}

// A short, memorable, unique name for this feed. Example: ACME Builds.
func (o HelmFeedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HelmFeed) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o HelmFeedOutput) PackageAcquisitionLocationOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HelmFeed) pulumi.StringArrayOutput { return v.PackageAcquisitionLocationOptions }).(pulumi.StringArrayOutput)
}

// The password associated with this resource.
func (o HelmFeedOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmFeed) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The space ID associated with this resource.
func (o HelmFeedOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *HelmFeed) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// The username associated with this resource.
func (o HelmFeedOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmFeed) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type HelmFeedArrayOutput struct{ *pulumi.OutputState }

func (HelmFeedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HelmFeed)(nil)).Elem()
}

func (o HelmFeedArrayOutput) ToHelmFeedArrayOutput() HelmFeedArrayOutput {
	return o
}

func (o HelmFeedArrayOutput) ToHelmFeedArrayOutputWithContext(ctx context.Context) HelmFeedArrayOutput {
	return o
}

func (o HelmFeedArrayOutput) Index(i pulumi.IntInput) HelmFeedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HelmFeed {
		return vs[0].([]*HelmFeed)[vs[1].(int)]
	}).(HelmFeedOutput)
}

type HelmFeedMapOutput struct{ *pulumi.OutputState }

func (HelmFeedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HelmFeed)(nil)).Elem()
}

func (o HelmFeedMapOutput) ToHelmFeedMapOutput() HelmFeedMapOutput {
	return o
}

func (o HelmFeedMapOutput) ToHelmFeedMapOutputWithContext(ctx context.Context) HelmFeedMapOutput {
	return o
}

func (o HelmFeedMapOutput) MapIndex(k pulumi.StringInput) HelmFeedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HelmFeed {
		return vs[0].(map[string]*HelmFeed)[vs[1].(string)]
	}).(HelmFeedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HelmFeedInput)(nil)).Elem(), &HelmFeed{})
	pulumi.RegisterInputType(reflect.TypeOf((*HelmFeedArrayInput)(nil)).Elem(), HelmFeedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HelmFeedMapInput)(nil)).Elem(), HelmFeedMap{})
	pulumi.RegisterOutputType(HelmFeedOutput{})
	pulumi.RegisterOutputType(HelmFeedArrayOutput{})
	pulumi.RegisterOutputType(HelmFeedMapOutput{})
}
