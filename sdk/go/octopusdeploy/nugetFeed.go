// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages a NuGet feed in Octopus Deploy.
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
//			_, err := octopusdeploy.NewNugetFeed(ctx, "example", &octopusdeploy.NugetFeedArgs{
//				DownloadAttempts:            pulumi.Int(1),
//				DownloadRetryBackoffSeconds: pulumi.Int(30),
//				FeedUri:                     pulumi.String("https://api.nuget.org/v3/index.json"),
//				IsEnhancedMode:              pulumi.Bool(true),
//				Password:                    pulumi.String("test-password"),
//				Username:                    pulumi.String("test-username"),
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
//	$ pulumi import octopusdeploy:index/nugetFeed:NugetFeed [options] octopusdeploy_nuget_feed.<name> <feed-id>
//
// ```
type NugetFeed struct {
	pulumi.CustomResourceState

	// The number of times a deployment should attempt to download a package from this feed before failing.
	DownloadAttempts pulumi.IntPtrOutput `pulumi:"downloadAttempts"`
	// The number of seconds to apply as a linear back off between download attempts.
	DownloadRetryBackoffSeconds pulumi.IntPtrOutput `pulumi:"downloadRetryBackoffSeconds"`
	// The feed URI can be a URL or a folder path.
	FeedUri pulumi.StringOutput `pulumi:"feedUri"`
	// This will improve performance of the NuGet feed but may not be supported by some older feeds. Disable if the operation, Create Release does not return the latest version for a package.
	IsEnhancedMode pulumi.BoolPtrOutput `pulumi:"isEnhancedMode"`
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

// NewNugetFeed registers a new resource with the given unique name, arguments, and options.
func NewNugetFeed(ctx *pulumi.Context,
	name string, args *NugetFeedArgs, opts ...pulumi.ResourceOption) (*NugetFeed, error) {
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
	var resource NugetFeed
	err := ctx.RegisterResource("octopusdeploy:index/nugetFeed:NugetFeed", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNugetFeed gets an existing NugetFeed resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNugetFeed(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NugetFeedState, opts ...pulumi.ResourceOption) (*NugetFeed, error) {
	var resource NugetFeed
	err := ctx.ReadResource("octopusdeploy:index/nugetFeed:NugetFeed", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NugetFeed resources.
type nugetFeedState struct {
	// The number of times a deployment should attempt to download a package from this feed before failing.
	DownloadAttempts *int `pulumi:"downloadAttempts"`
	// The number of seconds to apply as a linear back off between download attempts.
	DownloadRetryBackoffSeconds *int `pulumi:"downloadRetryBackoffSeconds"`
	// The feed URI can be a URL or a folder path.
	FeedUri *string `pulumi:"feedUri"`
	// This will improve performance of the NuGet feed but may not be supported by some older feeds. Disable if the operation, Create Release does not return the latest version for a package.
	IsEnhancedMode *bool `pulumi:"isEnhancedMode"`
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

type NugetFeedState struct {
	// The number of times a deployment should attempt to download a package from this feed before failing.
	DownloadAttempts pulumi.IntPtrInput
	// The number of seconds to apply as a linear back off between download attempts.
	DownloadRetryBackoffSeconds pulumi.IntPtrInput
	// The feed URI can be a URL or a folder path.
	FeedUri pulumi.StringPtrInput
	// This will improve performance of the NuGet feed but may not be supported by some older feeds. Disable if the operation, Create Release does not return the latest version for a package.
	IsEnhancedMode pulumi.BoolPtrInput
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

func (NugetFeedState) ElementType() reflect.Type {
	return reflect.TypeOf((*nugetFeedState)(nil)).Elem()
}

type nugetFeedArgs struct {
	// The number of times a deployment should attempt to download a package from this feed before failing.
	DownloadAttempts *int `pulumi:"downloadAttempts"`
	// The number of seconds to apply as a linear back off between download attempts.
	DownloadRetryBackoffSeconds *int `pulumi:"downloadRetryBackoffSeconds"`
	// The feed URI can be a URL or a folder path.
	FeedUri string `pulumi:"feedUri"`
	// This will improve performance of the NuGet feed but may not be supported by some older feeds. Disable if the operation, Create Release does not return the latest version for a package.
	IsEnhancedMode *bool `pulumi:"isEnhancedMode"`
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

// The set of arguments for constructing a NugetFeed resource.
type NugetFeedArgs struct {
	// The number of times a deployment should attempt to download a package from this feed before failing.
	DownloadAttempts pulumi.IntPtrInput
	// The number of seconds to apply as a linear back off between download attempts.
	DownloadRetryBackoffSeconds pulumi.IntPtrInput
	// The feed URI can be a URL or a folder path.
	FeedUri pulumi.StringInput
	// This will improve performance of the NuGet feed but may not be supported by some older feeds. Disable if the operation, Create Release does not return the latest version for a package.
	IsEnhancedMode pulumi.BoolPtrInput
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

func (NugetFeedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nugetFeedArgs)(nil)).Elem()
}

type NugetFeedInput interface {
	pulumi.Input

	ToNugetFeedOutput() NugetFeedOutput
	ToNugetFeedOutputWithContext(ctx context.Context) NugetFeedOutput
}

func (*NugetFeed) ElementType() reflect.Type {
	return reflect.TypeOf((**NugetFeed)(nil)).Elem()
}

func (i *NugetFeed) ToNugetFeedOutput() NugetFeedOutput {
	return i.ToNugetFeedOutputWithContext(context.Background())
}

func (i *NugetFeed) ToNugetFeedOutputWithContext(ctx context.Context) NugetFeedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NugetFeedOutput)
}

// NugetFeedArrayInput is an input type that accepts NugetFeedArray and NugetFeedArrayOutput values.
// You can construct a concrete instance of `NugetFeedArrayInput` via:
//
//	NugetFeedArray{ NugetFeedArgs{...} }
type NugetFeedArrayInput interface {
	pulumi.Input

	ToNugetFeedArrayOutput() NugetFeedArrayOutput
	ToNugetFeedArrayOutputWithContext(context.Context) NugetFeedArrayOutput
}

type NugetFeedArray []NugetFeedInput

func (NugetFeedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NugetFeed)(nil)).Elem()
}

func (i NugetFeedArray) ToNugetFeedArrayOutput() NugetFeedArrayOutput {
	return i.ToNugetFeedArrayOutputWithContext(context.Background())
}

func (i NugetFeedArray) ToNugetFeedArrayOutputWithContext(ctx context.Context) NugetFeedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NugetFeedArrayOutput)
}

// NugetFeedMapInput is an input type that accepts NugetFeedMap and NugetFeedMapOutput values.
// You can construct a concrete instance of `NugetFeedMapInput` via:
//
//	NugetFeedMap{ "key": NugetFeedArgs{...} }
type NugetFeedMapInput interface {
	pulumi.Input

	ToNugetFeedMapOutput() NugetFeedMapOutput
	ToNugetFeedMapOutputWithContext(context.Context) NugetFeedMapOutput
}

type NugetFeedMap map[string]NugetFeedInput

func (NugetFeedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NugetFeed)(nil)).Elem()
}

func (i NugetFeedMap) ToNugetFeedMapOutput() NugetFeedMapOutput {
	return i.ToNugetFeedMapOutputWithContext(context.Background())
}

func (i NugetFeedMap) ToNugetFeedMapOutputWithContext(ctx context.Context) NugetFeedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NugetFeedMapOutput)
}

type NugetFeedOutput struct{ *pulumi.OutputState }

func (NugetFeedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NugetFeed)(nil)).Elem()
}

func (o NugetFeedOutput) ToNugetFeedOutput() NugetFeedOutput {
	return o
}

func (o NugetFeedOutput) ToNugetFeedOutputWithContext(ctx context.Context) NugetFeedOutput {
	return o
}

// The number of times a deployment should attempt to download a package from this feed before failing.
func (o NugetFeedOutput) DownloadAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NugetFeed) pulumi.IntPtrOutput { return v.DownloadAttempts }).(pulumi.IntPtrOutput)
}

// The number of seconds to apply as a linear back off between download attempts.
func (o NugetFeedOutput) DownloadRetryBackoffSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NugetFeed) pulumi.IntPtrOutput { return v.DownloadRetryBackoffSeconds }).(pulumi.IntPtrOutput)
}

// The feed URI can be a URL or a folder path.
func (o NugetFeedOutput) FeedUri() pulumi.StringOutput {
	return o.ApplyT(func(v *NugetFeed) pulumi.StringOutput { return v.FeedUri }).(pulumi.StringOutput)
}

// This will improve performance of the NuGet feed but may not be supported by some older feeds. Disable if the operation, Create Release does not return the latest version for a package.
func (o NugetFeedOutput) IsEnhancedMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NugetFeed) pulumi.BoolPtrOutput { return v.IsEnhancedMode }).(pulumi.BoolPtrOutput)
}

// A short, memorable, unique name for this feed. Example: ACME Builds.
func (o NugetFeedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NugetFeed) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NugetFeedOutput) PackageAcquisitionLocationOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NugetFeed) pulumi.StringArrayOutput { return v.PackageAcquisitionLocationOptions }).(pulumi.StringArrayOutput)
}

// The password associated with this resource.
func (o NugetFeedOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NugetFeed) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The space ID associated with this resource.
func (o NugetFeedOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *NugetFeed) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// The username associated with this resource.
func (o NugetFeedOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NugetFeed) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type NugetFeedArrayOutput struct{ *pulumi.OutputState }

func (NugetFeedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NugetFeed)(nil)).Elem()
}

func (o NugetFeedArrayOutput) ToNugetFeedArrayOutput() NugetFeedArrayOutput {
	return o
}

func (o NugetFeedArrayOutput) ToNugetFeedArrayOutputWithContext(ctx context.Context) NugetFeedArrayOutput {
	return o
}

func (o NugetFeedArrayOutput) Index(i pulumi.IntInput) NugetFeedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NugetFeed {
		return vs[0].([]*NugetFeed)[vs[1].(int)]
	}).(NugetFeedOutput)
}

type NugetFeedMapOutput struct{ *pulumi.OutputState }

func (NugetFeedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NugetFeed)(nil)).Elem()
}

func (o NugetFeedMapOutput) ToNugetFeedMapOutput() NugetFeedMapOutput {
	return o
}

func (o NugetFeedMapOutput) ToNugetFeedMapOutputWithContext(ctx context.Context) NugetFeedMapOutput {
	return o
}

func (o NugetFeedMapOutput) MapIndex(k pulumi.StringInput) NugetFeedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NugetFeed {
		return vs[0].(map[string]*NugetFeed)[vs[1].(string)]
	}).(NugetFeedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NugetFeedInput)(nil)).Elem(), &NugetFeed{})
	pulumi.RegisterInputType(reflect.TypeOf((*NugetFeedArrayInput)(nil)).Elem(), NugetFeedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NugetFeedMapInput)(nil)).Elem(), NugetFeedMap{})
	pulumi.RegisterOutputType(NugetFeedOutput{})
	pulumi.RegisterOutputType(NugetFeedArrayOutput{})
	pulumi.RegisterOutputType(NugetFeedMapOutput{})
}
