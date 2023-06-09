// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about existing feeds.
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
//			_, err := octopusdeploy.GetFeeds(ctx, &octopusdeploy.GetFeedsArgs{
//				FeedType: pulumi.StringRef("NuGet"),
//				Ids: []string{
//					"Feeds-123",
//					"Feeds-321",
//				},
//				PartialName: pulumi.StringRef("Develop"),
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
func GetFeeds(ctx *pulumi.Context, args *GetFeedsArgs, opts ...pulumi.InvokeOption) (*GetFeedsResult, error) {
	var rv GetFeedsResult
	err := ctx.Invoke("octopusdeploy:index/getFeeds:getFeeds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFeeds.
type GetFeedsArgs struct {
	// A filter to search by feed type. Valid feed types are `AwsElasticContainerRegistry`, `BuiltIn`, `Docker`, `GitHub`, `Helm`, `Maven`, `NuGet`, or `OctopusProject`.
	FeedType *string `pulumi:"feedType"`
	// A list of feeds that match the filter(s).
	Feeds []GetFeedsFeed `pulumi:"feeds"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A filter to search by name.
	Name *string `pulumi:"name"`
	// A filter to search by the partial match of a name.
	PartialName *string `pulumi:"partialName"`
	// A filter to specify the number of items to skip in the response.
	Skip *int `pulumi:"skip"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
}

// A collection of values returned by getFeeds.
type GetFeedsResult struct {
	// A filter to search by feed type. Valid feed types are `AwsElasticContainerRegistry`, `BuiltIn`, `Docker`, `GitHub`, `Helm`, `Maven`, `NuGet`, or `OctopusProject`.
	FeedType *string `pulumi:"feedType"`
	// A list of feeds that match the filter(s).
	Feeds []GetFeedsFeed `pulumi:"feeds"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A filter to search by name.
	Name *string `pulumi:"name"`
	// A filter to search by the partial match of a name.
	PartialName *string `pulumi:"partialName"`
	// A filter to specify the number of items to skip in the response.
	Skip *int `pulumi:"skip"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
}

func GetFeedsOutput(ctx *pulumi.Context, args GetFeedsOutputArgs, opts ...pulumi.InvokeOption) GetFeedsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFeedsResult, error) {
			args := v.(GetFeedsArgs)
			r, err := GetFeeds(ctx, &args, opts...)
			var s GetFeedsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFeedsResultOutput)
}

// A collection of arguments for invoking getFeeds.
type GetFeedsOutputArgs struct {
	// A filter to search by feed type. Valid feed types are `AwsElasticContainerRegistry`, `BuiltIn`, `Docker`, `GitHub`, `Helm`, `Maven`, `NuGet`, or `OctopusProject`.
	FeedType pulumi.StringPtrInput `pulumi:"feedType"`
	// A list of feeds that match the filter(s).
	Feeds GetFeedsFeedArrayInput `pulumi:"feeds"`
	// A filter to search by a list of IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A filter to search by name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A filter to search by the partial match of a name.
	PartialName pulumi.StringPtrInput `pulumi:"partialName"`
	// A filter to specify the number of items to skip in the response.
	Skip pulumi.IntPtrInput `pulumi:"skip"`
	// A filter to specify the number of items to take (or return) in the response.
	Take pulumi.IntPtrInput `pulumi:"take"`
}

func (GetFeedsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFeedsArgs)(nil)).Elem()
}

// A collection of values returned by getFeeds.
type GetFeedsResultOutput struct{ *pulumi.OutputState }

func (GetFeedsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFeedsResult)(nil)).Elem()
}

func (o GetFeedsResultOutput) ToGetFeedsResultOutput() GetFeedsResultOutput {
	return o
}

func (o GetFeedsResultOutput) ToGetFeedsResultOutputWithContext(ctx context.Context) GetFeedsResultOutput {
	return o
}

// A filter to search by feed type. Valid feed types are `AwsElasticContainerRegistry`, `BuiltIn`, `Docker`, `GitHub`, `Helm`, `Maven`, `NuGet`, or `OctopusProject`.
func (o GetFeedsResultOutput) FeedType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFeedsResult) *string { return v.FeedType }).(pulumi.StringPtrOutput)
}

// A list of feeds that match the filter(s).
func (o GetFeedsResultOutput) Feeds() GetFeedsFeedArrayOutput {
	return o.ApplyT(func(v GetFeedsResult) []GetFeedsFeed { return v.Feeds }).(GetFeedsFeedArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFeedsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFeedsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A filter to search by a list of IDs.
func (o GetFeedsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFeedsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A filter to search by name.
func (o GetFeedsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFeedsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A filter to search by the partial match of a name.
func (o GetFeedsResultOutput) PartialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFeedsResult) *string { return v.PartialName }).(pulumi.StringPtrOutput)
}

// A filter to specify the number of items to skip in the response.
func (o GetFeedsResultOutput) Skip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFeedsResult) *int { return v.Skip }).(pulumi.IntPtrOutput)
}

// A filter to specify the number of items to take (or return) in the response.
func (o GetFeedsResultOutput) Take() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFeedsResult) *int { return v.Take }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFeedsResultOutput{})
}
