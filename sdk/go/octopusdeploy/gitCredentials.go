// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages Git credentials in Octopus Deploy.
type GitCredentials struct {
	pulumi.CustomResourceState

	// The description of this Git credential.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the Git credential. This name must be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password for the Git credential.
	Password pulumi.StringOutput `pulumi:"password"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// The Git credential authentication type.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The username for the Git credential.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewGitCredentials registers a new resource with the given unique name, arguments, and options.
func NewGitCredentials(ctx *pulumi.Context,
	name string, args *GitCredentialsArgs, opts ...pulumi.ResourceOption) (*GitCredentials, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource GitCredentials
	err := ctx.RegisterResource("octopusdeploy:index/gitCredentials:GitCredentials", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGitCredentials gets an existing GitCredentials resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGitCredentials(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GitCredentialsState, opts ...pulumi.ResourceOption) (*GitCredentials, error) {
	var resource GitCredentials
	err := ctx.ReadResource("octopusdeploy:index/gitCredentials:GitCredentials", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GitCredentials resources.
type gitCredentialsState struct {
	// The description of this Git credential.
	Description *string `pulumi:"description"`
	// The name of the Git credential. This name must be unique.
	Name *string `pulumi:"name"`
	// The password for the Git credential.
	Password *string `pulumi:"password"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// The Git credential authentication type.
	Type *string `pulumi:"type"`
	// The username for the Git credential.
	Username *string `pulumi:"username"`
}

type GitCredentialsState struct {
	// The description of this Git credential.
	Description pulumi.StringPtrInput
	// The name of the Git credential. This name must be unique.
	Name pulumi.StringPtrInput
	// The password for the Git credential.
	Password pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// The Git credential authentication type.
	Type pulumi.StringPtrInput
	// The username for the Git credential.
	Username pulumi.StringPtrInput
}

func (GitCredentialsState) ElementType() reflect.Type {
	return reflect.TypeOf((*gitCredentialsState)(nil)).Elem()
}

type gitCredentialsArgs struct {
	// The description of this Git credential.
	Description *string `pulumi:"description"`
	// The name of the Git credential. This name must be unique.
	Name *string `pulumi:"name"`
	// The password for the Git credential.
	Password string `pulumi:"password"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// The Git credential authentication type.
	Type *string `pulumi:"type"`
	// The username for the Git credential.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a GitCredentials resource.
type GitCredentialsArgs struct {
	// The description of this Git credential.
	Description pulumi.StringPtrInput
	// The name of the Git credential. This name must be unique.
	Name pulumi.StringPtrInput
	// The password for the Git credential.
	Password pulumi.StringInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// The Git credential authentication type.
	Type pulumi.StringPtrInput
	// The username for the Git credential.
	Username pulumi.StringInput
}

func (GitCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gitCredentialsArgs)(nil)).Elem()
}

type GitCredentialsInput interface {
	pulumi.Input

	ToGitCredentialsOutput() GitCredentialsOutput
	ToGitCredentialsOutputWithContext(ctx context.Context) GitCredentialsOutput
}

func (*GitCredentials) ElementType() reflect.Type {
	return reflect.TypeOf((**GitCredentials)(nil)).Elem()
}

func (i *GitCredentials) ToGitCredentialsOutput() GitCredentialsOutput {
	return i.ToGitCredentialsOutputWithContext(context.Background())
}

func (i *GitCredentials) ToGitCredentialsOutputWithContext(ctx context.Context) GitCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitCredentialsOutput)
}

// GitCredentialsArrayInput is an input type that accepts GitCredentialsArray and GitCredentialsArrayOutput values.
// You can construct a concrete instance of `GitCredentialsArrayInput` via:
//
//	GitCredentialsArray{ GitCredentialsArgs{...} }
type GitCredentialsArrayInput interface {
	pulumi.Input

	ToGitCredentialsArrayOutput() GitCredentialsArrayOutput
	ToGitCredentialsArrayOutputWithContext(context.Context) GitCredentialsArrayOutput
}

type GitCredentialsArray []GitCredentialsInput

func (GitCredentialsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GitCredentials)(nil)).Elem()
}

func (i GitCredentialsArray) ToGitCredentialsArrayOutput() GitCredentialsArrayOutput {
	return i.ToGitCredentialsArrayOutputWithContext(context.Background())
}

func (i GitCredentialsArray) ToGitCredentialsArrayOutputWithContext(ctx context.Context) GitCredentialsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitCredentialsArrayOutput)
}

// GitCredentialsMapInput is an input type that accepts GitCredentialsMap and GitCredentialsMapOutput values.
// You can construct a concrete instance of `GitCredentialsMapInput` via:
//
//	GitCredentialsMap{ "key": GitCredentialsArgs{...} }
type GitCredentialsMapInput interface {
	pulumi.Input

	ToGitCredentialsMapOutput() GitCredentialsMapOutput
	ToGitCredentialsMapOutputWithContext(context.Context) GitCredentialsMapOutput
}

type GitCredentialsMap map[string]GitCredentialsInput

func (GitCredentialsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GitCredentials)(nil)).Elem()
}

func (i GitCredentialsMap) ToGitCredentialsMapOutput() GitCredentialsMapOutput {
	return i.ToGitCredentialsMapOutputWithContext(context.Background())
}

func (i GitCredentialsMap) ToGitCredentialsMapOutputWithContext(ctx context.Context) GitCredentialsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitCredentialsMapOutput)
}

type GitCredentialsOutput struct{ *pulumi.OutputState }

func (GitCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitCredentials)(nil)).Elem()
}

func (o GitCredentialsOutput) ToGitCredentialsOutput() GitCredentialsOutput {
	return o
}

func (o GitCredentialsOutput) ToGitCredentialsOutputWithContext(ctx context.Context) GitCredentialsOutput {
	return o
}

// The description of this Git credential.
func (o GitCredentialsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitCredentials) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the Git credential. This name must be unique.
func (o GitCredentialsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GitCredentials) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The password for the Git credential.
func (o GitCredentialsOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *GitCredentials) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The space ID associated with this resource.
func (o GitCredentialsOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *GitCredentials) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// The Git credential authentication type.
func (o GitCredentialsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitCredentials) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// The username for the Git credential.
func (o GitCredentialsOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *GitCredentials) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type GitCredentialsArrayOutput struct{ *pulumi.OutputState }

func (GitCredentialsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GitCredentials)(nil)).Elem()
}

func (o GitCredentialsArrayOutput) ToGitCredentialsArrayOutput() GitCredentialsArrayOutput {
	return o
}

func (o GitCredentialsArrayOutput) ToGitCredentialsArrayOutputWithContext(ctx context.Context) GitCredentialsArrayOutput {
	return o
}

func (o GitCredentialsArrayOutput) Index(i pulumi.IntInput) GitCredentialsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GitCredentials {
		return vs[0].([]*GitCredentials)[vs[1].(int)]
	}).(GitCredentialsOutput)
}

type GitCredentialsMapOutput struct{ *pulumi.OutputState }

func (GitCredentialsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GitCredentials)(nil)).Elem()
}

func (o GitCredentialsMapOutput) ToGitCredentialsMapOutput() GitCredentialsMapOutput {
	return o
}

func (o GitCredentialsMapOutput) ToGitCredentialsMapOutputWithContext(ctx context.Context) GitCredentialsMapOutput {
	return o
}

func (o GitCredentialsMapOutput) MapIndex(k pulumi.StringInput) GitCredentialsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GitCredentials {
		return vs[0].(map[string]*GitCredentials)[vs[1].(string)]
	}).(GitCredentialsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GitCredentialsInput)(nil)).Elem(), &GitCredentials{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitCredentialsArrayInput)(nil)).Elem(), GitCredentialsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitCredentialsMapInput)(nil)).Elem(), GitCredentialsMap{})
	pulumi.RegisterOutputType(GitCredentialsOutput{})
	pulumi.RegisterOutputType(GitCredentialsArrayOutput{})
	pulumi.RegisterOutputType(GitCredentialsMapOutput{})
}
