// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages teams in Octopus Deploy.
//
// > **NOTE on Team User Roles and Scoped User Roles:** We currently
// provides both a standalone Scoped User Role resource
// and a Team resource with `userRoles` blocks defined in-line. At this time you
// cannot use a Team with in-line userRoles in conjunction with any Scoped User Role
// resources. Doing so will cause a conflict of user role settings and will overwrite
// user roles.
type Team struct {
	pulumi.CustomResourceState

	CanBeDeleted     pulumi.BoolOutput `pulumi:"canBeDeleted"`
	CanBeRenamed     pulumi.BoolOutput `pulumi:"canBeRenamed"`
	CanChangeMembers pulumi.BoolOutput `pulumi:"canChangeMembers"`
	CanChangeRoles   pulumi.BoolOutput `pulumi:"canChangeRoles"`
	// The user-friendly description of this team.
	Description            pulumi.StringPtrOutput               `pulumi:"description"`
	ExternalSecurityGroups TeamExternalSecurityGroupArrayOutput `pulumi:"externalSecurityGroups"`
	// The name of this team.
	Name pulumi.StringOutput `pulumi:"name"`
	// The space associated with this team.
	SpaceId   pulumi.StringOutput     `pulumi:"spaceId"`
	UserRoles TeamUserRoleArrayOutput `pulumi:"userRoles"`
	// A list of user IDs designated to be members of this team.
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewTeam registers a new resource with the given unique name, arguments, and options.
func NewTeam(ctx *pulumi.Context,
	name string, args *TeamArgs, opts ...pulumi.ResourceOption) (*Team, error) {
	if args == nil {
		args = &TeamArgs{}
	}

	var resource Team
	err := ctx.RegisterResource("octopusdeploy:index/team:Team", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeam gets an existing Team resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamState, opts ...pulumi.ResourceOption) (*Team, error) {
	var resource Team
	err := ctx.ReadResource("octopusdeploy:index/team:Team", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Team resources.
type teamState struct {
	CanBeDeleted     *bool `pulumi:"canBeDeleted"`
	CanBeRenamed     *bool `pulumi:"canBeRenamed"`
	CanChangeMembers *bool `pulumi:"canChangeMembers"`
	CanChangeRoles   *bool `pulumi:"canChangeRoles"`
	// The user-friendly description of this team.
	Description            *string                     `pulumi:"description"`
	ExternalSecurityGroups []TeamExternalSecurityGroup `pulumi:"externalSecurityGroups"`
	// The name of this team.
	Name *string `pulumi:"name"`
	// The space associated with this team.
	SpaceId   *string        `pulumi:"spaceId"`
	UserRoles []TeamUserRole `pulumi:"userRoles"`
	// A list of user IDs designated to be members of this team.
	Users []string `pulumi:"users"`
}

type TeamState struct {
	CanBeDeleted     pulumi.BoolPtrInput
	CanBeRenamed     pulumi.BoolPtrInput
	CanChangeMembers pulumi.BoolPtrInput
	CanChangeRoles   pulumi.BoolPtrInput
	// The user-friendly description of this team.
	Description            pulumi.StringPtrInput
	ExternalSecurityGroups TeamExternalSecurityGroupArrayInput
	// The name of this team.
	Name pulumi.StringPtrInput
	// The space associated with this team.
	SpaceId   pulumi.StringPtrInput
	UserRoles TeamUserRoleArrayInput
	// A list of user IDs designated to be members of this team.
	Users pulumi.StringArrayInput
}

func (TeamState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamState)(nil)).Elem()
}

type teamArgs struct {
	CanBeDeleted     *bool `pulumi:"canBeDeleted"`
	CanBeRenamed     *bool `pulumi:"canBeRenamed"`
	CanChangeMembers *bool `pulumi:"canChangeMembers"`
	CanChangeRoles   *bool `pulumi:"canChangeRoles"`
	// The user-friendly description of this team.
	Description            *string                     `pulumi:"description"`
	ExternalSecurityGroups []TeamExternalSecurityGroup `pulumi:"externalSecurityGroups"`
	// The name of this team.
	Name *string `pulumi:"name"`
	// The space associated with this team.
	SpaceId   *string        `pulumi:"spaceId"`
	UserRoles []TeamUserRole `pulumi:"userRoles"`
	// A list of user IDs designated to be members of this team.
	Users []string `pulumi:"users"`
}

// The set of arguments for constructing a Team resource.
type TeamArgs struct {
	CanBeDeleted     pulumi.BoolPtrInput
	CanBeRenamed     pulumi.BoolPtrInput
	CanChangeMembers pulumi.BoolPtrInput
	CanChangeRoles   pulumi.BoolPtrInput
	// The user-friendly description of this team.
	Description            pulumi.StringPtrInput
	ExternalSecurityGroups TeamExternalSecurityGroupArrayInput
	// The name of this team.
	Name pulumi.StringPtrInput
	// The space associated with this team.
	SpaceId   pulumi.StringPtrInput
	UserRoles TeamUserRoleArrayInput
	// A list of user IDs designated to be members of this team.
	Users pulumi.StringArrayInput
}

func (TeamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamArgs)(nil)).Elem()
}

type TeamInput interface {
	pulumi.Input

	ToTeamOutput() TeamOutput
	ToTeamOutputWithContext(ctx context.Context) TeamOutput
}

func (*Team) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (i *Team) ToTeamOutput() TeamOutput {
	return i.ToTeamOutputWithContext(context.Background())
}

func (i *Team) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamOutput)
}

// TeamArrayInput is an input type that accepts TeamArray and TeamArrayOutput values.
// You can construct a concrete instance of `TeamArrayInput` via:
//
//	TeamArray{ TeamArgs{...} }
type TeamArrayInput interface {
	pulumi.Input

	ToTeamArrayOutput() TeamArrayOutput
	ToTeamArrayOutputWithContext(context.Context) TeamArrayOutput
}

type TeamArray []TeamInput

func (TeamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (i TeamArray) ToTeamArrayOutput() TeamArrayOutput {
	return i.ToTeamArrayOutputWithContext(context.Background())
}

func (i TeamArray) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamArrayOutput)
}

// TeamMapInput is an input type that accepts TeamMap and TeamMapOutput values.
// You can construct a concrete instance of `TeamMapInput` via:
//
//	TeamMap{ "key": TeamArgs{...} }
type TeamMapInput interface {
	pulumi.Input

	ToTeamMapOutput() TeamMapOutput
	ToTeamMapOutputWithContext(context.Context) TeamMapOutput
}

type TeamMap map[string]TeamInput

func (TeamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (i TeamMap) ToTeamMapOutput() TeamMapOutput {
	return i.ToTeamMapOutputWithContext(context.Background())
}

func (i TeamMap) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMapOutput)
}

type TeamOutput struct{ *pulumi.OutputState }

func (TeamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (o TeamOutput) ToTeamOutput() TeamOutput {
	return o
}

func (o TeamOutput) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return o
}

func (o TeamOutput) CanBeDeleted() pulumi.BoolOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolOutput { return v.CanBeDeleted }).(pulumi.BoolOutput)
}

func (o TeamOutput) CanBeRenamed() pulumi.BoolOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolOutput { return v.CanBeRenamed }).(pulumi.BoolOutput)
}

func (o TeamOutput) CanChangeMembers() pulumi.BoolOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolOutput { return v.CanChangeMembers }).(pulumi.BoolOutput)
}

func (o TeamOutput) CanChangeRoles() pulumi.BoolOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolOutput { return v.CanChangeRoles }).(pulumi.BoolOutput)
}

// The user-friendly description of this team.
func (o TeamOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TeamOutput) ExternalSecurityGroups() TeamExternalSecurityGroupArrayOutput {
	return o.ApplyT(func(v *Team) TeamExternalSecurityGroupArrayOutput { return v.ExternalSecurityGroups }).(TeamExternalSecurityGroupArrayOutput)
}

// The name of this team.
func (o TeamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The space associated with this team.
func (o TeamOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

func (o TeamOutput) UserRoles() TeamUserRoleArrayOutput {
	return o.ApplyT(func(v *Team) TeamUserRoleArrayOutput { return v.UserRoles }).(TeamUserRoleArrayOutput)
}

// A list of user IDs designated to be members of this team.
func (o TeamOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Team) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
}

type TeamArrayOutput struct{ *pulumi.OutputState }

func (TeamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (o TeamArrayOutput) ToTeamArrayOutput() TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) Index(i pulumi.IntInput) TeamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Team {
		return vs[0].([]*Team)[vs[1].(int)]
	}).(TeamOutput)
}

type TeamMapOutput struct{ *pulumi.OutputState }

func (TeamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (o TeamMapOutput) ToTeamMapOutput() TeamMapOutput {
	return o
}

func (o TeamMapOutput) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return o
}

func (o TeamMapOutput) MapIndex(k pulumi.StringInput) TeamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Team {
		return vs[0].(map[string]*Team)[vs[1].(string)]
	}).(TeamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamInput)(nil)).Elem(), &Team{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamArrayInput)(nil)).Elem(), TeamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMapInput)(nil)).Elem(), TeamMap{})
	pulumi.RegisterOutputType(TeamOutput{})
	pulumi.RegisterOutputType(TeamArrayOutput{})
	pulumi.RegisterOutputType(TeamMapOutput{})
}