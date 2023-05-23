// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about existing channels.
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
//			_, err := octopusdeploy.GetChannels(ctx, &GetChannelsArgs{
//				Ids: []string{
//					"Channels-123",
//					"Channels-321",
//				},
//				PartialName: pulumi.StringRef("Defau"),
//				Skip:        pulumi.IntRef(5),
//				Take:        pulumi.IntRef(100),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetChannels(ctx *pulumi.Context, args *GetChannelsArgs, opts ...pulumi.InvokeOption) (*GetChannelsResult, error) {
	var rv GetChannelsResult
	err := ctx.Invoke("octopusdeploy:index/getChannels:getChannels", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getChannels.
type GetChannelsArgs struct {
	// A channel that matches the specified filter(s).
	Channels []GetChannelsChannel `pulumi:"channels"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A filter to search by the partial match of a name.
	PartialName *string `pulumi:"partialName"`
	// A filter to specify the number of items to skip in the response.
	Skip *int `pulumi:"skip"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
}

// A collection of values returned by getChannels.
type GetChannelsResult struct {
	// A channel that matches the specified filter(s).
	Channels []GetChannelsChannel `pulumi:"channels"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A filter to search by the partial match of a name.
	PartialName *string `pulumi:"partialName"`
	// A filter to specify the number of items to skip in the response.
	Skip *int `pulumi:"skip"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
}

func GetChannelsOutput(ctx *pulumi.Context, args GetChannelsOutputArgs, opts ...pulumi.InvokeOption) GetChannelsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetChannelsResult, error) {
			args := v.(GetChannelsArgs)
			r, err := GetChannels(ctx, &args, opts...)
			var s GetChannelsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetChannelsResultOutput)
}

// A collection of arguments for invoking getChannels.
type GetChannelsOutputArgs struct {
	// A channel that matches the specified filter(s).
	Channels GetChannelsChannelArrayInput `pulumi:"channels"`
	// A filter to search by a list of IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A filter to search by the partial match of a name.
	PartialName pulumi.StringPtrInput `pulumi:"partialName"`
	// A filter to specify the number of items to skip in the response.
	Skip pulumi.IntPtrInput `pulumi:"skip"`
	// A filter to specify the number of items to take (or return) in the response.
	Take pulumi.IntPtrInput `pulumi:"take"`
}

func (GetChannelsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetChannelsArgs)(nil)).Elem()
}

// A collection of values returned by getChannels.
type GetChannelsResultOutput struct{ *pulumi.OutputState }

func (GetChannelsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetChannelsResult)(nil)).Elem()
}

func (o GetChannelsResultOutput) ToGetChannelsResultOutput() GetChannelsResultOutput {
	return o
}

func (o GetChannelsResultOutput) ToGetChannelsResultOutputWithContext(ctx context.Context) GetChannelsResultOutput {
	return o
}

// A channel that matches the specified filter(s).
func (o GetChannelsResultOutput) Channels() GetChannelsChannelArrayOutput {
	return o.ApplyT(func(v GetChannelsResult) []GetChannelsChannel { return v.Channels }).(GetChannelsChannelArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetChannelsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetChannelsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A filter to search by a list of IDs.
func (o GetChannelsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetChannelsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A filter to search by the partial match of a name.
func (o GetChannelsResultOutput) PartialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetChannelsResult) *string { return v.PartialName }).(pulumi.StringPtrOutput)
}

// A filter to specify the number of items to skip in the response.
func (o GetChannelsResultOutput) Skip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetChannelsResult) *int { return v.Skip }).(pulumi.IntPtrOutput)
}

// A filter to specify the number of items to take (or return) in the response.
func (o GetChannelsResultOutput) Take() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetChannelsResult) *int { return v.Take }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetChannelsResultOutput{})
}