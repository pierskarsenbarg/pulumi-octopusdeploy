// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages tag sets in Octopus Deploy.
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
//			example, err := octopusdeploy.NewTagSet(ctx, "example", &octopusdeploy.TagSetArgs{
//				Description: pulumi.String("Provides tenants with access to certain early access programs."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = octopusdeploy.NewTag(ctx, "alpha", &octopusdeploy.TagArgs{
//				Color:    pulumi.String("#00FF00"),
//				TagSetId: example.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = octopusdeploy.NewTag(ctx, "beta", &octopusdeploy.TagArgs{
//				Color:    pulumi.String("#FF0000"),
//				TagSetId: example.ID(),
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
//	$ pulumi import octopusdeploy:index/tagSet:TagSet [options] octopusdeploy_tag_set.<name> <tag-set-id>
//
// ```
type TagSet struct {
	pulumi.CustomResourceState

	// The description of this tag set.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The sort order associated with this resource.
	SortOrder pulumi.IntOutput `pulumi:"sortOrder"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
}

// NewTagSet registers a new resource with the given unique name, arguments, and options.
func NewTagSet(ctx *pulumi.Context,
	name string, args *TagSetArgs, opts ...pulumi.ResourceOption) (*TagSet, error) {
	if args == nil {
		args = &TagSetArgs{}
	}

	var resource TagSet
	err := ctx.RegisterResource("octopusdeploy:index/tagSet:TagSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagSet gets an existing TagSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagSetState, opts ...pulumi.ResourceOption) (*TagSet, error) {
	var resource TagSet
	err := ctx.ReadResource("octopusdeploy:index/tagSet:TagSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagSet resources.
type tagSetState struct {
	// The description of this tag set.
	Description *string `pulumi:"description"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The sort order associated with this resource.
	SortOrder *int `pulumi:"sortOrder"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
}

type TagSetState struct {
	// The description of this tag set.
	Description pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The sort order associated with this resource.
	SortOrder pulumi.IntPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
}

func (TagSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagSetState)(nil)).Elem()
}

type tagSetArgs struct {
	// The description of this tag set.
	Description *string `pulumi:"description"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The sort order associated with this resource.
	SortOrder *int `pulumi:"sortOrder"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
}

// The set of arguments for constructing a TagSet resource.
type TagSetArgs struct {
	// The description of this tag set.
	Description pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The sort order associated with this resource.
	SortOrder pulumi.IntPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
}

func (TagSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagSetArgs)(nil)).Elem()
}

type TagSetInput interface {
	pulumi.Input

	ToTagSetOutput() TagSetOutput
	ToTagSetOutputWithContext(ctx context.Context) TagSetOutput
}

func (*TagSet) ElementType() reflect.Type {
	return reflect.TypeOf((**TagSet)(nil)).Elem()
}

func (i *TagSet) ToTagSetOutput() TagSetOutput {
	return i.ToTagSetOutputWithContext(context.Background())
}

func (i *TagSet) ToTagSetOutputWithContext(ctx context.Context) TagSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagSetOutput)
}

// TagSetArrayInput is an input type that accepts TagSetArray and TagSetArrayOutput values.
// You can construct a concrete instance of `TagSetArrayInput` via:
//
//	TagSetArray{ TagSetArgs{...} }
type TagSetArrayInput interface {
	pulumi.Input

	ToTagSetArrayOutput() TagSetArrayOutput
	ToTagSetArrayOutputWithContext(context.Context) TagSetArrayOutput
}

type TagSetArray []TagSetInput

func (TagSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagSet)(nil)).Elem()
}

func (i TagSetArray) ToTagSetArrayOutput() TagSetArrayOutput {
	return i.ToTagSetArrayOutputWithContext(context.Background())
}

func (i TagSetArray) ToTagSetArrayOutputWithContext(ctx context.Context) TagSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagSetArrayOutput)
}

// TagSetMapInput is an input type that accepts TagSetMap and TagSetMapOutput values.
// You can construct a concrete instance of `TagSetMapInput` via:
//
//	TagSetMap{ "key": TagSetArgs{...} }
type TagSetMapInput interface {
	pulumi.Input

	ToTagSetMapOutput() TagSetMapOutput
	ToTagSetMapOutputWithContext(context.Context) TagSetMapOutput
}

type TagSetMap map[string]TagSetInput

func (TagSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagSet)(nil)).Elem()
}

func (i TagSetMap) ToTagSetMapOutput() TagSetMapOutput {
	return i.ToTagSetMapOutputWithContext(context.Background())
}

func (i TagSetMap) ToTagSetMapOutputWithContext(ctx context.Context) TagSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagSetMapOutput)
}

type TagSetOutput struct{ *pulumi.OutputState }

func (TagSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagSet)(nil)).Elem()
}

func (o TagSetOutput) ToTagSetOutput() TagSetOutput {
	return o
}

func (o TagSetOutput) ToTagSetOutputWithContext(ctx context.Context) TagSetOutput {
	return o
}

// The description of this tag set.
func (o TagSetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagSet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of this resource.
func (o TagSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The sort order associated with this resource.
func (o TagSetOutput) SortOrder() pulumi.IntOutput {
	return o.ApplyT(func(v *TagSet) pulumi.IntOutput { return v.SortOrder }).(pulumi.IntOutput)
}

// The space ID associated with this resource.
func (o TagSetOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TagSet) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

type TagSetArrayOutput struct{ *pulumi.OutputState }

func (TagSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagSet)(nil)).Elem()
}

func (o TagSetArrayOutput) ToTagSetArrayOutput() TagSetArrayOutput {
	return o
}

func (o TagSetArrayOutput) ToTagSetArrayOutputWithContext(ctx context.Context) TagSetArrayOutput {
	return o
}

func (o TagSetArrayOutput) Index(i pulumi.IntInput) TagSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TagSet {
		return vs[0].([]*TagSet)[vs[1].(int)]
	}).(TagSetOutput)
}

type TagSetMapOutput struct{ *pulumi.OutputState }

func (TagSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagSet)(nil)).Elem()
}

func (o TagSetMapOutput) ToTagSetMapOutput() TagSetMapOutput {
	return o
}

func (o TagSetMapOutput) ToTagSetMapOutputWithContext(ctx context.Context) TagSetMapOutput {
	return o
}

func (o TagSetMapOutput) MapIndex(k pulumi.StringInput) TagSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TagSet {
		return vs[0].(map[string]*TagSet)[vs[1].(string)]
	}).(TagSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagSetInput)(nil)).Elem(), &TagSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagSetArrayInput)(nil)).Elem(), TagSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagSetMapInput)(nil)).Elem(), TagSetMap{})
	pulumi.RegisterOutputType(TagSetOutput{})
	pulumi.RegisterOutputType(TagSetArrayOutput{})
	pulumi.RegisterOutputType(TagSetMapOutput{})
}
