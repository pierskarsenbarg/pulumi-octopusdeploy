// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages spaces in Octopus Deploy.
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
//			_, err := octopusdeploy.NewSpace(ctx, "example", &octopusdeploy.SpaceArgs{
//				Description:        pulumi.String("A space for the development team."),
//				IsDefault:          pulumi.Bool(false),
//				IsTaskQueueStopped: pulumi.Bool(false),
//				SpaceManagersTeamMembers: pulumi.StringArray{
//					pulumi.String("Users-123"),
//					pulumi.String("Users-321"),
//				},
//				SpaceManagersTeams: pulumi.StringArray{
//					pulumi.String("teams-everyone"),
//				},
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
//	$ pulumi import octopusdeploy:index/space:Space [options] octopusdeploy_space.<name> <space-id>
//
// ```
type Space struct {
	pulumi.CustomResourceState

	// The description of this space.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies if this space is the default space in Octopus.
	IsDefault pulumi.BoolPtrOutput `pulumi:"isDefault"`
	// Specifies the status of the task queue for this space.
	IsTaskQueueStopped pulumi.BoolPtrOutput `pulumi:"isTaskQueueStopped"`
	// The name of this resource, no more than 20 characters long
	Name pulumi.StringOutput `pulumi:"name"`
	// The unique slug of this space.
	Slug pulumi.StringOutput `pulumi:"slug"`
	// A list of user IDs designated to be managers of this space.
	SpaceManagersTeamMembers pulumi.StringArrayOutput `pulumi:"spaceManagersTeamMembers"`
	// A list of team IDs designated to be managers of this space.
	SpaceManagersTeams pulumi.StringArrayOutput `pulumi:"spaceManagersTeams"`
}

// NewSpace registers a new resource with the given unique name, arguments, and options.
func NewSpace(ctx *pulumi.Context,
	name string, args *SpaceArgs, opts ...pulumi.ResourceOption) (*Space, error) {
	if args == nil {
		args = &SpaceArgs{}
	}

	var resource Space
	err := ctx.RegisterResource("octopusdeploy:index/space:Space", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpace gets an existing Space resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpaceState, opts ...pulumi.ResourceOption) (*Space, error) {
	var resource Space
	err := ctx.ReadResource("octopusdeploy:index/space:Space", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Space resources.
type spaceState struct {
	// The description of this space.
	Description *string `pulumi:"description"`
	// Specifies if this space is the default space in Octopus.
	IsDefault *bool `pulumi:"isDefault"`
	// Specifies the status of the task queue for this space.
	IsTaskQueueStopped *bool `pulumi:"isTaskQueueStopped"`
	// The name of this resource, no more than 20 characters long
	Name *string `pulumi:"name"`
	// The unique slug of this space.
	Slug *string `pulumi:"slug"`
	// A list of user IDs designated to be managers of this space.
	SpaceManagersTeamMembers []string `pulumi:"spaceManagersTeamMembers"`
	// A list of team IDs designated to be managers of this space.
	SpaceManagersTeams []string `pulumi:"spaceManagersTeams"`
}

type SpaceState struct {
	// The description of this space.
	Description pulumi.StringPtrInput
	// Specifies if this space is the default space in Octopus.
	IsDefault pulumi.BoolPtrInput
	// Specifies the status of the task queue for this space.
	IsTaskQueueStopped pulumi.BoolPtrInput
	// The name of this resource, no more than 20 characters long
	Name pulumi.StringPtrInput
	// The unique slug of this space.
	Slug pulumi.StringPtrInput
	// A list of user IDs designated to be managers of this space.
	SpaceManagersTeamMembers pulumi.StringArrayInput
	// A list of team IDs designated to be managers of this space.
	SpaceManagersTeams pulumi.StringArrayInput
}

func (SpaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*spaceState)(nil)).Elem()
}

type spaceArgs struct {
	// The description of this space.
	Description *string `pulumi:"description"`
	// Specifies if this space is the default space in Octopus.
	IsDefault *bool `pulumi:"isDefault"`
	// Specifies the status of the task queue for this space.
	IsTaskQueueStopped *bool `pulumi:"isTaskQueueStopped"`
	// The name of this resource, no more than 20 characters long
	Name *string `pulumi:"name"`
	// The unique slug of this space.
	Slug *string `pulumi:"slug"`
	// A list of user IDs designated to be managers of this space.
	SpaceManagersTeamMembers []string `pulumi:"spaceManagersTeamMembers"`
	// A list of team IDs designated to be managers of this space.
	SpaceManagersTeams []string `pulumi:"spaceManagersTeams"`
}

// The set of arguments for constructing a Space resource.
type SpaceArgs struct {
	// The description of this space.
	Description pulumi.StringPtrInput
	// Specifies if this space is the default space in Octopus.
	IsDefault pulumi.BoolPtrInput
	// Specifies the status of the task queue for this space.
	IsTaskQueueStopped pulumi.BoolPtrInput
	// The name of this resource, no more than 20 characters long
	Name pulumi.StringPtrInput
	// The unique slug of this space.
	Slug pulumi.StringPtrInput
	// A list of user IDs designated to be managers of this space.
	SpaceManagersTeamMembers pulumi.StringArrayInput
	// A list of team IDs designated to be managers of this space.
	SpaceManagersTeams pulumi.StringArrayInput
}

func (SpaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spaceArgs)(nil)).Elem()
}

type SpaceInput interface {
	pulumi.Input

	ToSpaceOutput() SpaceOutput
	ToSpaceOutputWithContext(ctx context.Context) SpaceOutput
}

func (*Space) ElementType() reflect.Type {
	return reflect.TypeOf((**Space)(nil)).Elem()
}

func (i *Space) ToSpaceOutput() SpaceOutput {
	return i.ToSpaceOutputWithContext(context.Background())
}

func (i *Space) ToSpaceOutputWithContext(ctx context.Context) SpaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpaceOutput)
}

// SpaceArrayInput is an input type that accepts SpaceArray and SpaceArrayOutput values.
// You can construct a concrete instance of `SpaceArrayInput` via:
//
//	SpaceArray{ SpaceArgs{...} }
type SpaceArrayInput interface {
	pulumi.Input

	ToSpaceArrayOutput() SpaceArrayOutput
	ToSpaceArrayOutputWithContext(context.Context) SpaceArrayOutput
}

type SpaceArray []SpaceInput

func (SpaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Space)(nil)).Elem()
}

func (i SpaceArray) ToSpaceArrayOutput() SpaceArrayOutput {
	return i.ToSpaceArrayOutputWithContext(context.Background())
}

func (i SpaceArray) ToSpaceArrayOutputWithContext(ctx context.Context) SpaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpaceArrayOutput)
}

// SpaceMapInput is an input type that accepts SpaceMap and SpaceMapOutput values.
// You can construct a concrete instance of `SpaceMapInput` via:
//
//	SpaceMap{ "key": SpaceArgs{...} }
type SpaceMapInput interface {
	pulumi.Input

	ToSpaceMapOutput() SpaceMapOutput
	ToSpaceMapOutputWithContext(context.Context) SpaceMapOutput
}

type SpaceMap map[string]SpaceInput

func (SpaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Space)(nil)).Elem()
}

func (i SpaceMap) ToSpaceMapOutput() SpaceMapOutput {
	return i.ToSpaceMapOutputWithContext(context.Background())
}

func (i SpaceMap) ToSpaceMapOutputWithContext(ctx context.Context) SpaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpaceMapOutput)
}

type SpaceOutput struct{ *pulumi.OutputState }

func (SpaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Space)(nil)).Elem()
}

func (o SpaceOutput) ToSpaceOutput() SpaceOutput {
	return o
}

func (o SpaceOutput) ToSpaceOutputWithContext(ctx context.Context) SpaceOutput {
	return o
}

// The description of this space.
func (o SpaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Space) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies if this space is the default space in Octopus.
func (o SpaceOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Space) pulumi.BoolPtrOutput { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// Specifies the status of the task queue for this space.
func (o SpaceOutput) IsTaskQueueStopped() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Space) pulumi.BoolPtrOutput { return v.IsTaskQueueStopped }).(pulumi.BoolPtrOutput)
}

// The name of this resource, no more than 20 characters long
func (o SpaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Space) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The unique slug of this space.
func (o SpaceOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Space) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// A list of user IDs designated to be managers of this space.
func (o SpaceOutput) SpaceManagersTeamMembers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Space) pulumi.StringArrayOutput { return v.SpaceManagersTeamMembers }).(pulumi.StringArrayOutput)
}

// A list of team IDs designated to be managers of this space.
func (o SpaceOutput) SpaceManagersTeams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Space) pulumi.StringArrayOutput { return v.SpaceManagersTeams }).(pulumi.StringArrayOutput)
}

type SpaceArrayOutput struct{ *pulumi.OutputState }

func (SpaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Space)(nil)).Elem()
}

func (o SpaceArrayOutput) ToSpaceArrayOutput() SpaceArrayOutput {
	return o
}

func (o SpaceArrayOutput) ToSpaceArrayOutputWithContext(ctx context.Context) SpaceArrayOutput {
	return o
}

func (o SpaceArrayOutput) Index(i pulumi.IntInput) SpaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Space {
		return vs[0].([]*Space)[vs[1].(int)]
	}).(SpaceOutput)
}

type SpaceMapOutput struct{ *pulumi.OutputState }

func (SpaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Space)(nil)).Elem()
}

func (o SpaceMapOutput) ToSpaceMapOutput() SpaceMapOutput {
	return o
}

func (o SpaceMapOutput) ToSpaceMapOutputWithContext(ctx context.Context) SpaceMapOutput {
	return o
}

func (o SpaceMapOutput) MapIndex(k pulumi.StringInput) SpaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Space {
		return vs[0].(map[string]*Space)[vs[1].(string)]
	}).(SpaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpaceInput)(nil)).Elem(), &Space{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpaceArrayInput)(nil)).Elem(), SpaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpaceMapInput)(nil)).Elem(), SpaceMap{})
	pulumi.RegisterOutputType(SpaceOutput{})
	pulumi.RegisterOutputType(SpaceArrayOutput{})
	pulumi.RegisterOutputType(SpaceMapOutput{})
}
