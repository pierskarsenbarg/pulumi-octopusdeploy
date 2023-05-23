// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages user roles in Octopus Deploy.
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
//			_, err := octopusdeploy.NewUserRole(ctx, "example", &octopusdeploy.UserRoleArgs{
//				CanBeDeleted: pulumi.Bool(true),
//				Description:  pulumi.String("Responsible for all development-related operations."),
//				GrantedSpacePermissions: pulumi.StringArray{
//					pulumi.String("DeploymentCreate"),
//					pulumi.String("DeploymentDelete"),
//					pulumi.String("DeploymentView"),
//				},
//				GrantedSystemPermissions: pulumi.StringArray{
//					pulumi.String("SpaceCreate"),
//				},
//				SpacePermissionDescriptions: pulumi.StringArray{
//					pulumi.String("Delete deployments (restrictable to Environments, Projects, Tenants)"),
//					pulumi.String("Deploy releases to target environments (restrictable to Environments, Projects, Tenants)"),
//					pulumi.String("View deployments (restrictable to Environments, Projects, Tenants)"),
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
//	$ pulumi import octopusdeploy:index/userRole:UserRole [options] octopusdeploy_user_role.<name> <user-role-id>
//
// ```
type UserRole struct {
	pulumi.CustomResourceState

	CanBeDeleted pulumi.BoolOutput `pulumi:"canBeDeleted"`
	// The description of this user role.
	Description              pulumi.StringPtrOutput   `pulumi:"description"`
	GrantedSpacePermissions  pulumi.StringArrayOutput `pulumi:"grantedSpacePermissions"`
	GrantedSystemPermissions pulumi.StringArrayOutput `pulumi:"grantedSystemPermissions"`
	// The name of this resource.
	Name                         pulumi.StringOutput      `pulumi:"name"`
	SpacePermissionDescriptions  pulumi.StringArrayOutput `pulumi:"spacePermissionDescriptions"`
	SupportedRestrictions        pulumi.StringArrayOutput `pulumi:"supportedRestrictions"`
	SystemPermissionDescriptions pulumi.StringArrayOutput `pulumi:"systemPermissionDescriptions"`
}

// NewUserRole registers a new resource with the given unique name, arguments, and options.
func NewUserRole(ctx *pulumi.Context,
	name string, args *UserRoleArgs, opts ...pulumi.ResourceOption) (*UserRole, error) {
	if args == nil {
		args = &UserRoleArgs{}
	}

	var resource UserRole
	err := ctx.RegisterResource("octopusdeploy:index/userRole:UserRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserRole gets an existing UserRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserRoleState, opts ...pulumi.ResourceOption) (*UserRole, error) {
	var resource UserRole
	err := ctx.ReadResource("octopusdeploy:index/userRole:UserRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserRole resources.
type userRoleState struct {
	CanBeDeleted *bool `pulumi:"canBeDeleted"`
	// The description of this user role.
	Description              *string  `pulumi:"description"`
	GrantedSpacePermissions  []string `pulumi:"grantedSpacePermissions"`
	GrantedSystemPermissions []string `pulumi:"grantedSystemPermissions"`
	// The name of this resource.
	Name                         *string  `pulumi:"name"`
	SpacePermissionDescriptions  []string `pulumi:"spacePermissionDescriptions"`
	SupportedRestrictions        []string `pulumi:"supportedRestrictions"`
	SystemPermissionDescriptions []string `pulumi:"systemPermissionDescriptions"`
}

type UserRoleState struct {
	CanBeDeleted pulumi.BoolPtrInput
	// The description of this user role.
	Description              pulumi.StringPtrInput
	GrantedSpacePermissions  pulumi.StringArrayInput
	GrantedSystemPermissions pulumi.StringArrayInput
	// The name of this resource.
	Name                         pulumi.StringPtrInput
	SpacePermissionDescriptions  pulumi.StringArrayInput
	SupportedRestrictions        pulumi.StringArrayInput
	SystemPermissionDescriptions pulumi.StringArrayInput
}

func (UserRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*userRoleState)(nil)).Elem()
}

type userRoleArgs struct {
	CanBeDeleted *bool `pulumi:"canBeDeleted"`
	// The description of this user role.
	Description              *string  `pulumi:"description"`
	GrantedSpacePermissions  []string `pulumi:"grantedSpacePermissions"`
	GrantedSystemPermissions []string `pulumi:"grantedSystemPermissions"`
	// The name of this resource.
	Name                         *string  `pulumi:"name"`
	SpacePermissionDescriptions  []string `pulumi:"spacePermissionDescriptions"`
	SupportedRestrictions        []string `pulumi:"supportedRestrictions"`
	SystemPermissionDescriptions []string `pulumi:"systemPermissionDescriptions"`
}

// The set of arguments for constructing a UserRole resource.
type UserRoleArgs struct {
	CanBeDeleted pulumi.BoolPtrInput
	// The description of this user role.
	Description              pulumi.StringPtrInput
	GrantedSpacePermissions  pulumi.StringArrayInput
	GrantedSystemPermissions pulumi.StringArrayInput
	// The name of this resource.
	Name                         pulumi.StringPtrInput
	SpacePermissionDescriptions  pulumi.StringArrayInput
	SupportedRestrictions        pulumi.StringArrayInput
	SystemPermissionDescriptions pulumi.StringArrayInput
}

func (UserRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userRoleArgs)(nil)).Elem()
}

type UserRoleInput interface {
	pulumi.Input

	ToUserRoleOutput() UserRoleOutput
	ToUserRoleOutputWithContext(ctx context.Context) UserRoleOutput
}

func (*UserRole) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRole)(nil)).Elem()
}

func (i *UserRole) ToUserRoleOutput() UserRoleOutput {
	return i.ToUserRoleOutputWithContext(context.Background())
}

func (i *UserRole) ToUserRoleOutputWithContext(ctx context.Context) UserRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRoleOutput)
}

// UserRoleArrayInput is an input type that accepts UserRoleArray and UserRoleArrayOutput values.
// You can construct a concrete instance of `UserRoleArrayInput` via:
//
//	UserRoleArray{ UserRoleArgs{...} }
type UserRoleArrayInput interface {
	pulumi.Input

	ToUserRoleArrayOutput() UserRoleArrayOutput
	ToUserRoleArrayOutputWithContext(context.Context) UserRoleArrayOutput
}

type UserRoleArray []UserRoleInput

func (UserRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserRole)(nil)).Elem()
}

func (i UserRoleArray) ToUserRoleArrayOutput() UserRoleArrayOutput {
	return i.ToUserRoleArrayOutputWithContext(context.Background())
}

func (i UserRoleArray) ToUserRoleArrayOutputWithContext(ctx context.Context) UserRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRoleArrayOutput)
}

// UserRoleMapInput is an input type that accepts UserRoleMap and UserRoleMapOutput values.
// You can construct a concrete instance of `UserRoleMapInput` via:
//
//	UserRoleMap{ "key": UserRoleArgs{...} }
type UserRoleMapInput interface {
	pulumi.Input

	ToUserRoleMapOutput() UserRoleMapOutput
	ToUserRoleMapOutputWithContext(context.Context) UserRoleMapOutput
}

type UserRoleMap map[string]UserRoleInput

func (UserRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserRole)(nil)).Elem()
}

func (i UserRoleMap) ToUserRoleMapOutput() UserRoleMapOutput {
	return i.ToUserRoleMapOutputWithContext(context.Background())
}

func (i UserRoleMap) ToUserRoleMapOutputWithContext(ctx context.Context) UserRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRoleMapOutput)
}

type UserRoleOutput struct{ *pulumi.OutputState }

func (UserRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRole)(nil)).Elem()
}

func (o UserRoleOutput) ToUserRoleOutput() UserRoleOutput {
	return o
}

func (o UserRoleOutput) ToUserRoleOutputWithContext(ctx context.Context) UserRoleOutput {
	return o
}

func (o UserRoleOutput) CanBeDeleted() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserRole) pulumi.BoolOutput { return v.CanBeDeleted }).(pulumi.BoolOutput)
}

// The description of this user role.
func (o UserRoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserRole) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o UserRoleOutput) GrantedSpacePermissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserRole) pulumi.StringArrayOutput { return v.GrantedSpacePermissions }).(pulumi.StringArrayOutput)
}

func (o UserRoleOutput) GrantedSystemPermissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserRole) pulumi.StringArrayOutput { return v.GrantedSystemPermissions }).(pulumi.StringArrayOutput)
}

// The name of this resource.
func (o UserRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UserRoleOutput) SpacePermissionDescriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserRole) pulumi.StringArrayOutput { return v.SpacePermissionDescriptions }).(pulumi.StringArrayOutput)
}

func (o UserRoleOutput) SupportedRestrictions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserRole) pulumi.StringArrayOutput { return v.SupportedRestrictions }).(pulumi.StringArrayOutput)
}

func (o UserRoleOutput) SystemPermissionDescriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserRole) pulumi.StringArrayOutput { return v.SystemPermissionDescriptions }).(pulumi.StringArrayOutput)
}

type UserRoleArrayOutput struct{ *pulumi.OutputState }

func (UserRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserRole)(nil)).Elem()
}

func (o UserRoleArrayOutput) ToUserRoleArrayOutput() UserRoleArrayOutput {
	return o
}

func (o UserRoleArrayOutput) ToUserRoleArrayOutputWithContext(ctx context.Context) UserRoleArrayOutput {
	return o
}

func (o UserRoleArrayOutput) Index(i pulumi.IntInput) UserRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserRole {
		return vs[0].([]*UserRole)[vs[1].(int)]
	}).(UserRoleOutput)
}

type UserRoleMapOutput struct{ *pulumi.OutputState }

func (UserRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserRole)(nil)).Elem()
}

func (o UserRoleMapOutput) ToUserRoleMapOutput() UserRoleMapOutput {
	return o
}

func (o UserRoleMapOutput) ToUserRoleMapOutputWithContext(ctx context.Context) UserRoleMapOutput {
	return o
}

func (o UserRoleMapOutput) MapIndex(k pulumi.StringInput) UserRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserRole {
		return vs[0].(map[string]*UserRole)[vs[1].(string)]
	}).(UserRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserRoleInput)(nil)).Elem(), &UserRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserRoleArrayInput)(nil)).Elem(), UserRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserRoleMapInput)(nil)).Elem(), UserRoleMap{})
	pulumi.RegisterOutputType(UserRoleOutput{})
	pulumi.RegisterOutputType(UserRoleArrayOutput{})
	pulumi.RegisterOutputType(UserRoleMapOutput{})
}
