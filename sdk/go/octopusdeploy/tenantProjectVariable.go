// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages tenant project variables in Octopus Deploy.
type TenantProjectVariable struct {
	pulumi.CustomResourceState

	EnvironmentId pulumi.StringOutput    `pulumi:"environmentId"`
	ProjectId     pulumi.StringOutput    `pulumi:"projectId"`
	TemplateId    pulumi.StringOutput    `pulumi:"templateId"`
	TenantId      pulumi.StringOutput    `pulumi:"tenantId"`
	Value         pulumi.StringPtrOutput `pulumi:"value"`
}

// NewTenantProjectVariable registers a new resource with the given unique name, arguments, and options.
func NewTenantProjectVariable(ctx *pulumi.Context,
	name string, args *TenantProjectVariableArgs, opts ...pulumi.ResourceOption) (*TenantProjectVariable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	var resource TenantProjectVariable
	err := ctx.RegisterResource("octopusdeploy:index/tenantProjectVariable:TenantProjectVariable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenantProjectVariable gets an existing TenantProjectVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenantProjectVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantProjectVariableState, opts ...pulumi.ResourceOption) (*TenantProjectVariable, error) {
	var resource TenantProjectVariable
	err := ctx.ReadResource("octopusdeploy:index/tenantProjectVariable:TenantProjectVariable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TenantProjectVariable resources.
type tenantProjectVariableState struct {
	EnvironmentId *string `pulumi:"environmentId"`
	ProjectId     *string `pulumi:"projectId"`
	TemplateId    *string `pulumi:"templateId"`
	TenantId      *string `pulumi:"tenantId"`
	Value         *string `pulumi:"value"`
}

type TenantProjectVariableState struct {
	EnvironmentId pulumi.StringPtrInput
	ProjectId     pulumi.StringPtrInput
	TemplateId    pulumi.StringPtrInput
	TenantId      pulumi.StringPtrInput
	Value         pulumi.StringPtrInput
}

func (TenantProjectVariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantProjectVariableState)(nil)).Elem()
}

type tenantProjectVariableArgs struct {
	EnvironmentId string  `pulumi:"environmentId"`
	ProjectId     string  `pulumi:"projectId"`
	TemplateId    string  `pulumi:"templateId"`
	TenantId      string  `pulumi:"tenantId"`
	Value         *string `pulumi:"value"`
}

// The set of arguments for constructing a TenantProjectVariable resource.
type TenantProjectVariableArgs struct {
	EnvironmentId pulumi.StringInput
	ProjectId     pulumi.StringInput
	TemplateId    pulumi.StringInput
	TenantId      pulumi.StringInput
	Value         pulumi.StringPtrInput
}

func (TenantProjectVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantProjectVariableArgs)(nil)).Elem()
}

type TenantProjectVariableInput interface {
	pulumi.Input

	ToTenantProjectVariableOutput() TenantProjectVariableOutput
	ToTenantProjectVariableOutputWithContext(ctx context.Context) TenantProjectVariableOutput
}

func (*TenantProjectVariable) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantProjectVariable)(nil)).Elem()
}

func (i *TenantProjectVariable) ToTenantProjectVariableOutput() TenantProjectVariableOutput {
	return i.ToTenantProjectVariableOutputWithContext(context.Background())
}

func (i *TenantProjectVariable) ToTenantProjectVariableOutputWithContext(ctx context.Context) TenantProjectVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantProjectVariableOutput)
}

// TenantProjectVariableArrayInput is an input type that accepts TenantProjectVariableArray and TenantProjectVariableArrayOutput values.
// You can construct a concrete instance of `TenantProjectVariableArrayInput` via:
//
//	TenantProjectVariableArray{ TenantProjectVariableArgs{...} }
type TenantProjectVariableArrayInput interface {
	pulumi.Input

	ToTenantProjectVariableArrayOutput() TenantProjectVariableArrayOutput
	ToTenantProjectVariableArrayOutputWithContext(context.Context) TenantProjectVariableArrayOutput
}

type TenantProjectVariableArray []TenantProjectVariableInput

func (TenantProjectVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TenantProjectVariable)(nil)).Elem()
}

func (i TenantProjectVariableArray) ToTenantProjectVariableArrayOutput() TenantProjectVariableArrayOutput {
	return i.ToTenantProjectVariableArrayOutputWithContext(context.Background())
}

func (i TenantProjectVariableArray) ToTenantProjectVariableArrayOutputWithContext(ctx context.Context) TenantProjectVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantProjectVariableArrayOutput)
}

// TenantProjectVariableMapInput is an input type that accepts TenantProjectVariableMap and TenantProjectVariableMapOutput values.
// You can construct a concrete instance of `TenantProjectVariableMapInput` via:
//
//	TenantProjectVariableMap{ "key": TenantProjectVariableArgs{...} }
type TenantProjectVariableMapInput interface {
	pulumi.Input

	ToTenantProjectVariableMapOutput() TenantProjectVariableMapOutput
	ToTenantProjectVariableMapOutputWithContext(context.Context) TenantProjectVariableMapOutput
}

type TenantProjectVariableMap map[string]TenantProjectVariableInput

func (TenantProjectVariableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TenantProjectVariable)(nil)).Elem()
}

func (i TenantProjectVariableMap) ToTenantProjectVariableMapOutput() TenantProjectVariableMapOutput {
	return i.ToTenantProjectVariableMapOutputWithContext(context.Background())
}

func (i TenantProjectVariableMap) ToTenantProjectVariableMapOutputWithContext(ctx context.Context) TenantProjectVariableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantProjectVariableMapOutput)
}

type TenantProjectVariableOutput struct{ *pulumi.OutputState }

func (TenantProjectVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantProjectVariable)(nil)).Elem()
}

func (o TenantProjectVariableOutput) ToTenantProjectVariableOutput() TenantProjectVariableOutput {
	return o
}

func (o TenantProjectVariableOutput) ToTenantProjectVariableOutputWithContext(ctx context.Context) TenantProjectVariableOutput {
	return o
}

func (o TenantProjectVariableOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantProjectVariable) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

func (o TenantProjectVariableOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantProjectVariable) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

func (o TenantProjectVariableOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantProjectVariable) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

func (o TenantProjectVariableOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantProjectVariable) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o TenantProjectVariableOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TenantProjectVariable) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

type TenantProjectVariableArrayOutput struct{ *pulumi.OutputState }

func (TenantProjectVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TenantProjectVariable)(nil)).Elem()
}

func (o TenantProjectVariableArrayOutput) ToTenantProjectVariableArrayOutput() TenantProjectVariableArrayOutput {
	return o
}

func (o TenantProjectVariableArrayOutput) ToTenantProjectVariableArrayOutputWithContext(ctx context.Context) TenantProjectVariableArrayOutput {
	return o
}

func (o TenantProjectVariableArrayOutput) Index(i pulumi.IntInput) TenantProjectVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TenantProjectVariable {
		return vs[0].([]*TenantProjectVariable)[vs[1].(int)]
	}).(TenantProjectVariableOutput)
}

type TenantProjectVariableMapOutput struct{ *pulumi.OutputState }

func (TenantProjectVariableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TenantProjectVariable)(nil)).Elem()
}

func (o TenantProjectVariableMapOutput) ToTenantProjectVariableMapOutput() TenantProjectVariableMapOutput {
	return o
}

func (o TenantProjectVariableMapOutput) ToTenantProjectVariableMapOutputWithContext(ctx context.Context) TenantProjectVariableMapOutput {
	return o
}

func (o TenantProjectVariableMapOutput) MapIndex(k pulumi.StringInput) TenantProjectVariableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TenantProjectVariable {
		return vs[0].(map[string]*TenantProjectVariable)[vs[1].(string)]
	}).(TenantProjectVariableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TenantProjectVariableInput)(nil)).Elem(), &TenantProjectVariable{})
	pulumi.RegisterInputType(reflect.TypeOf((*TenantProjectVariableArrayInput)(nil)).Elem(), TenantProjectVariableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TenantProjectVariableMapInput)(nil)).Elem(), TenantProjectVariableMap{})
	pulumi.RegisterOutputType(TenantProjectVariableOutput{})
	pulumi.RegisterOutputType(TenantProjectVariableArrayOutput{})
	pulumi.RegisterOutputType(TenantProjectVariableMapOutput{})
}
