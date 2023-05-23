// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages tenant common variables in Octopus Deploy.
type TenantCommonVariable struct {
	pulumi.CustomResourceState

	LibraryVariableSetId pulumi.StringOutput    `pulumi:"libraryVariableSetId"`
	TemplateId           pulumi.StringOutput    `pulumi:"templateId"`
	TenantId             pulumi.StringOutput    `pulumi:"tenantId"`
	Value                pulumi.StringPtrOutput `pulumi:"value"`
}

// NewTenantCommonVariable registers a new resource with the given unique name, arguments, and options.
func NewTenantCommonVariable(ctx *pulumi.Context,
	name string, args *TenantCommonVariableArgs, opts ...pulumi.ResourceOption) (*TenantCommonVariable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LibraryVariableSetId == nil {
		return nil, errors.New("invalid value for required argument 'LibraryVariableSetId'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	var resource TenantCommonVariable
	err := ctx.RegisterResource("octopusdeploy:index/tenantCommonVariable:TenantCommonVariable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenantCommonVariable gets an existing TenantCommonVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenantCommonVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantCommonVariableState, opts ...pulumi.ResourceOption) (*TenantCommonVariable, error) {
	var resource TenantCommonVariable
	err := ctx.ReadResource("octopusdeploy:index/tenantCommonVariable:TenantCommonVariable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TenantCommonVariable resources.
type tenantCommonVariableState struct {
	LibraryVariableSetId *string `pulumi:"libraryVariableSetId"`
	TemplateId           *string `pulumi:"templateId"`
	TenantId             *string `pulumi:"tenantId"`
	Value                *string `pulumi:"value"`
}

type TenantCommonVariableState struct {
	LibraryVariableSetId pulumi.StringPtrInput
	TemplateId           pulumi.StringPtrInput
	TenantId             pulumi.StringPtrInput
	Value                pulumi.StringPtrInput
}

func (TenantCommonVariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantCommonVariableState)(nil)).Elem()
}

type tenantCommonVariableArgs struct {
	LibraryVariableSetId string  `pulumi:"libraryVariableSetId"`
	TemplateId           string  `pulumi:"templateId"`
	TenantId             string  `pulumi:"tenantId"`
	Value                *string `pulumi:"value"`
}

// The set of arguments for constructing a TenantCommonVariable resource.
type TenantCommonVariableArgs struct {
	LibraryVariableSetId pulumi.StringInput
	TemplateId           pulumi.StringInput
	TenantId             pulumi.StringInput
	Value                pulumi.StringPtrInput
}

func (TenantCommonVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantCommonVariableArgs)(nil)).Elem()
}

type TenantCommonVariableInput interface {
	pulumi.Input

	ToTenantCommonVariableOutput() TenantCommonVariableOutput
	ToTenantCommonVariableOutputWithContext(ctx context.Context) TenantCommonVariableOutput
}

func (*TenantCommonVariable) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantCommonVariable)(nil)).Elem()
}

func (i *TenantCommonVariable) ToTenantCommonVariableOutput() TenantCommonVariableOutput {
	return i.ToTenantCommonVariableOutputWithContext(context.Background())
}

func (i *TenantCommonVariable) ToTenantCommonVariableOutputWithContext(ctx context.Context) TenantCommonVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantCommonVariableOutput)
}

// TenantCommonVariableArrayInput is an input type that accepts TenantCommonVariableArray and TenantCommonVariableArrayOutput values.
// You can construct a concrete instance of `TenantCommonVariableArrayInput` via:
//
//	TenantCommonVariableArray{ TenantCommonVariableArgs{...} }
type TenantCommonVariableArrayInput interface {
	pulumi.Input

	ToTenantCommonVariableArrayOutput() TenantCommonVariableArrayOutput
	ToTenantCommonVariableArrayOutputWithContext(context.Context) TenantCommonVariableArrayOutput
}

type TenantCommonVariableArray []TenantCommonVariableInput

func (TenantCommonVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TenantCommonVariable)(nil)).Elem()
}

func (i TenantCommonVariableArray) ToTenantCommonVariableArrayOutput() TenantCommonVariableArrayOutput {
	return i.ToTenantCommonVariableArrayOutputWithContext(context.Background())
}

func (i TenantCommonVariableArray) ToTenantCommonVariableArrayOutputWithContext(ctx context.Context) TenantCommonVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantCommonVariableArrayOutput)
}

// TenantCommonVariableMapInput is an input type that accepts TenantCommonVariableMap and TenantCommonVariableMapOutput values.
// You can construct a concrete instance of `TenantCommonVariableMapInput` via:
//
//	TenantCommonVariableMap{ "key": TenantCommonVariableArgs{...} }
type TenantCommonVariableMapInput interface {
	pulumi.Input

	ToTenantCommonVariableMapOutput() TenantCommonVariableMapOutput
	ToTenantCommonVariableMapOutputWithContext(context.Context) TenantCommonVariableMapOutput
}

type TenantCommonVariableMap map[string]TenantCommonVariableInput

func (TenantCommonVariableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TenantCommonVariable)(nil)).Elem()
}

func (i TenantCommonVariableMap) ToTenantCommonVariableMapOutput() TenantCommonVariableMapOutput {
	return i.ToTenantCommonVariableMapOutputWithContext(context.Background())
}

func (i TenantCommonVariableMap) ToTenantCommonVariableMapOutputWithContext(ctx context.Context) TenantCommonVariableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantCommonVariableMapOutput)
}

type TenantCommonVariableOutput struct{ *pulumi.OutputState }

func (TenantCommonVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantCommonVariable)(nil)).Elem()
}

func (o TenantCommonVariableOutput) ToTenantCommonVariableOutput() TenantCommonVariableOutput {
	return o
}

func (o TenantCommonVariableOutput) ToTenantCommonVariableOutputWithContext(ctx context.Context) TenantCommonVariableOutput {
	return o
}

func (o TenantCommonVariableOutput) LibraryVariableSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantCommonVariable) pulumi.StringOutput { return v.LibraryVariableSetId }).(pulumi.StringOutput)
}

func (o TenantCommonVariableOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantCommonVariable) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

func (o TenantCommonVariableOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantCommonVariable) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o TenantCommonVariableOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TenantCommonVariable) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

type TenantCommonVariableArrayOutput struct{ *pulumi.OutputState }

func (TenantCommonVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TenantCommonVariable)(nil)).Elem()
}

func (o TenantCommonVariableArrayOutput) ToTenantCommonVariableArrayOutput() TenantCommonVariableArrayOutput {
	return o
}

func (o TenantCommonVariableArrayOutput) ToTenantCommonVariableArrayOutputWithContext(ctx context.Context) TenantCommonVariableArrayOutput {
	return o
}

func (o TenantCommonVariableArrayOutput) Index(i pulumi.IntInput) TenantCommonVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TenantCommonVariable {
		return vs[0].([]*TenantCommonVariable)[vs[1].(int)]
	}).(TenantCommonVariableOutput)
}

type TenantCommonVariableMapOutput struct{ *pulumi.OutputState }

func (TenantCommonVariableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TenantCommonVariable)(nil)).Elem()
}

func (o TenantCommonVariableMapOutput) ToTenantCommonVariableMapOutput() TenantCommonVariableMapOutput {
	return o
}

func (o TenantCommonVariableMapOutput) ToTenantCommonVariableMapOutputWithContext(ctx context.Context) TenantCommonVariableMapOutput {
	return o
}

func (o TenantCommonVariableMapOutput) MapIndex(k pulumi.StringInput) TenantCommonVariableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TenantCommonVariable {
		return vs[0].(map[string]*TenantCommonVariable)[vs[1].(string)]
	}).(TenantCommonVariableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TenantCommonVariableInput)(nil)).Elem(), &TenantCommonVariable{})
	pulumi.RegisterInputType(reflect.TypeOf((*TenantCommonVariableArrayInput)(nil)).Elem(), TenantCommonVariableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TenantCommonVariableMapInput)(nil)).Elem(), TenantCommonVariableMap{})
	pulumi.RegisterOutputType(TenantCommonVariableOutput{})
	pulumi.RegisterOutputType(TenantCommonVariableArrayOutput{})
	pulumi.RegisterOutputType(TenantCommonVariableMapOutput{})
}
