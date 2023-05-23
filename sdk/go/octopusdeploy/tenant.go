// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages tenants in Octopus Deploy.
type Tenant struct {
	pulumi.CustomResourceState

	// The ID of the tenant from which this tenant was cloned.
	ClonedFromTenantId pulumi.StringPtrOutput `pulumi:"clonedFromTenantId"`
	// The description of this tenant.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of this resource.
	Name                pulumi.StringOutput                 `pulumi:"name"`
	ProjectEnvironments TenantProjectEnvironmentArrayOutput `pulumi:"projectEnvironments"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayOutput `pulumi:"tenantTags"`
}

// NewTenant registers a new resource with the given unique name, arguments, and options.
func NewTenant(ctx *pulumi.Context,
	name string, args *TenantArgs, opts ...pulumi.ResourceOption) (*Tenant, error) {
	if args == nil {
		args = &TenantArgs{}
	}

	var resource Tenant
	err := ctx.RegisterResource("octopusdeploy:index/tenant:Tenant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenant gets an existing Tenant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantState, opts ...pulumi.ResourceOption) (*Tenant, error) {
	var resource Tenant
	err := ctx.ReadResource("octopusdeploy:index/tenant:Tenant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tenant resources.
type tenantState struct {
	// The ID of the tenant from which this tenant was cloned.
	ClonedFromTenantId *string `pulumi:"clonedFromTenantId"`
	// The description of this tenant.
	Description *string `pulumi:"description"`
	// The name of this resource.
	Name                *string                    `pulumi:"name"`
	ProjectEnvironments []TenantProjectEnvironment `pulumi:"projectEnvironments"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
}

type TenantState struct {
	// The ID of the tenant from which this tenant was cloned.
	ClonedFromTenantId pulumi.StringPtrInput
	// The description of this tenant.
	Description pulumi.StringPtrInput
	// The name of this resource.
	Name                pulumi.StringPtrInput
	ProjectEnvironments TenantProjectEnvironmentArrayInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
}

func (TenantState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantState)(nil)).Elem()
}

type tenantArgs struct {
	// The ID of the tenant from which this tenant was cloned.
	ClonedFromTenantId *string `pulumi:"clonedFromTenantId"`
	// The description of this tenant.
	Description *string `pulumi:"description"`
	// The name of this resource.
	Name                *string                    `pulumi:"name"`
	ProjectEnvironments []TenantProjectEnvironment `pulumi:"projectEnvironments"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
}

// The set of arguments for constructing a Tenant resource.
type TenantArgs struct {
	// The ID of the tenant from which this tenant was cloned.
	ClonedFromTenantId pulumi.StringPtrInput
	// The description of this tenant.
	Description pulumi.StringPtrInput
	// The name of this resource.
	Name                pulumi.StringPtrInput
	ProjectEnvironments TenantProjectEnvironmentArrayInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
}

func (TenantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantArgs)(nil)).Elem()
}

type TenantInput interface {
	pulumi.Input

	ToTenantOutput() TenantOutput
	ToTenantOutputWithContext(ctx context.Context) TenantOutput
}

func (*Tenant) ElementType() reflect.Type {
	return reflect.TypeOf((**Tenant)(nil)).Elem()
}

func (i *Tenant) ToTenantOutput() TenantOutput {
	return i.ToTenantOutputWithContext(context.Background())
}

func (i *Tenant) ToTenantOutputWithContext(ctx context.Context) TenantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantOutput)
}

// TenantArrayInput is an input type that accepts TenantArray and TenantArrayOutput values.
// You can construct a concrete instance of `TenantArrayInput` via:
//
//	TenantArray{ TenantArgs{...} }
type TenantArrayInput interface {
	pulumi.Input

	ToTenantArrayOutput() TenantArrayOutput
	ToTenantArrayOutputWithContext(context.Context) TenantArrayOutput
}

type TenantArray []TenantInput

func (TenantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tenant)(nil)).Elem()
}

func (i TenantArray) ToTenantArrayOutput() TenantArrayOutput {
	return i.ToTenantArrayOutputWithContext(context.Background())
}

func (i TenantArray) ToTenantArrayOutputWithContext(ctx context.Context) TenantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantArrayOutput)
}

// TenantMapInput is an input type that accepts TenantMap and TenantMapOutput values.
// You can construct a concrete instance of `TenantMapInput` via:
//
//	TenantMap{ "key": TenantArgs{...} }
type TenantMapInput interface {
	pulumi.Input

	ToTenantMapOutput() TenantMapOutput
	ToTenantMapOutputWithContext(context.Context) TenantMapOutput
}

type TenantMap map[string]TenantInput

func (TenantMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tenant)(nil)).Elem()
}

func (i TenantMap) ToTenantMapOutput() TenantMapOutput {
	return i.ToTenantMapOutputWithContext(context.Background())
}

func (i TenantMap) ToTenantMapOutputWithContext(ctx context.Context) TenantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantMapOutput)
}

type TenantOutput struct{ *pulumi.OutputState }

func (TenantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tenant)(nil)).Elem()
}

func (o TenantOutput) ToTenantOutput() TenantOutput {
	return o
}

func (o TenantOutput) ToTenantOutputWithContext(ctx context.Context) TenantOutput {
	return o
}

// The ID of the tenant from which this tenant was cloned.
func (o TenantOutput) ClonedFromTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tenant) pulumi.StringPtrOutput { return v.ClonedFromTenantId }).(pulumi.StringPtrOutput)
}

// The description of this tenant.
func (o TenantOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tenant) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of this resource.
func (o TenantOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Tenant) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TenantOutput) ProjectEnvironments() TenantProjectEnvironmentArrayOutput {
	return o.ApplyT(func(v *Tenant) TenantProjectEnvironmentArrayOutput { return v.ProjectEnvironments }).(TenantProjectEnvironmentArrayOutput)
}

// The space ID associated with this resource.
func (o TenantOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Tenant) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// A list of tenant tags associated with this resource.
func (o TenantOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Tenant) pulumi.StringArrayOutput { return v.TenantTags }).(pulumi.StringArrayOutput)
}

type TenantArrayOutput struct{ *pulumi.OutputState }

func (TenantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tenant)(nil)).Elem()
}

func (o TenantArrayOutput) ToTenantArrayOutput() TenantArrayOutput {
	return o
}

func (o TenantArrayOutput) ToTenantArrayOutputWithContext(ctx context.Context) TenantArrayOutput {
	return o
}

func (o TenantArrayOutput) Index(i pulumi.IntInput) TenantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Tenant {
		return vs[0].([]*Tenant)[vs[1].(int)]
	}).(TenantOutput)
}

type TenantMapOutput struct{ *pulumi.OutputState }

func (TenantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tenant)(nil)).Elem()
}

func (o TenantMapOutput) ToTenantMapOutput() TenantMapOutput {
	return o
}

func (o TenantMapOutput) ToTenantMapOutputWithContext(ctx context.Context) TenantMapOutput {
	return o
}

func (o TenantMapOutput) MapIndex(k pulumi.StringInput) TenantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Tenant {
		return vs[0].(map[string]*Tenant)[vs[1].(string)]
	}).(TenantOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TenantInput)(nil)).Elem(), &Tenant{})
	pulumi.RegisterInputType(reflect.TypeOf((*TenantArrayInput)(nil)).Elem(), TenantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TenantMapInput)(nil)).Elem(), TenantMap{})
	pulumi.RegisterOutputType(TenantOutput{})
	pulumi.RegisterOutputType(TenantArrayOutput{})
	pulumi.RegisterOutputType(TenantMapOutput{})
}
