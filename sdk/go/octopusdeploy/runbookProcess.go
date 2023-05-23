// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages runbook processes in Octopus Deploy.
type RunbookProcess struct {
	pulumi.CustomResourceState

	// Read only value containing the last snapshot ID.
	LastSnapshotId pulumi.StringPtrOutput `pulumi:"lastSnapshotId"`
	// The project ID associated with this runbook process.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The runbook ID associated with this runbook process.
	RunbookId pulumi.StringOutput `pulumi:"runbookId"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput           `pulumi:"spaceId"`
	Steps   RunbookProcessStepArrayOutput `pulumi:"steps"`
	// The version number of this runbook process.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewRunbookProcess registers a new resource with the given unique name, arguments, and options.
func NewRunbookProcess(ctx *pulumi.Context,
	name string, args *RunbookProcessArgs, opts ...pulumi.ResourceOption) (*RunbookProcess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RunbookId == nil {
		return nil, errors.New("invalid value for required argument 'RunbookId'")
	}
	var resource RunbookProcess
	err := ctx.RegisterResource("octopusdeploy:index/runbookProcess:RunbookProcess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRunbookProcess gets an existing RunbookProcess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRunbookProcess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RunbookProcessState, opts ...pulumi.ResourceOption) (*RunbookProcess, error) {
	var resource RunbookProcess
	err := ctx.ReadResource("octopusdeploy:index/runbookProcess:RunbookProcess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RunbookProcess resources.
type runbookProcessState struct {
	// Read only value containing the last snapshot ID.
	LastSnapshotId *string `pulumi:"lastSnapshotId"`
	// The project ID associated with this runbook process.
	ProjectId *string `pulumi:"projectId"`
	// The runbook ID associated with this runbook process.
	RunbookId *string `pulumi:"runbookId"`
	// The space ID associated with this resource.
	SpaceId *string              `pulumi:"spaceId"`
	Steps   []RunbookProcessStep `pulumi:"steps"`
	// The version number of this runbook process.
	Version *int `pulumi:"version"`
}

type RunbookProcessState struct {
	// Read only value containing the last snapshot ID.
	LastSnapshotId pulumi.StringPtrInput
	// The project ID associated with this runbook process.
	ProjectId pulumi.StringPtrInput
	// The runbook ID associated with this runbook process.
	RunbookId pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	Steps   RunbookProcessStepArrayInput
	// The version number of this runbook process.
	Version pulumi.IntPtrInput
}

func (RunbookProcessState) ElementType() reflect.Type {
	return reflect.TypeOf((*runbookProcessState)(nil)).Elem()
}

type runbookProcessArgs struct {
	// Read only value containing the last snapshot ID.
	LastSnapshotId *string `pulumi:"lastSnapshotId"`
	// The project ID associated with this runbook process.
	ProjectId *string `pulumi:"projectId"`
	// The runbook ID associated with this runbook process.
	RunbookId string `pulumi:"runbookId"`
	// The space ID associated with this resource.
	SpaceId *string              `pulumi:"spaceId"`
	Steps   []RunbookProcessStep `pulumi:"steps"`
	// The version number of this runbook process.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a RunbookProcess resource.
type RunbookProcessArgs struct {
	// Read only value containing the last snapshot ID.
	LastSnapshotId pulumi.StringPtrInput
	// The project ID associated with this runbook process.
	ProjectId pulumi.StringPtrInput
	// The runbook ID associated with this runbook process.
	RunbookId pulumi.StringInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	Steps   RunbookProcessStepArrayInput
	// The version number of this runbook process.
	Version pulumi.IntPtrInput
}

func (RunbookProcessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runbookProcessArgs)(nil)).Elem()
}

type RunbookProcessInput interface {
	pulumi.Input

	ToRunbookProcessOutput() RunbookProcessOutput
	ToRunbookProcessOutputWithContext(ctx context.Context) RunbookProcessOutput
}

func (*RunbookProcess) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookProcess)(nil)).Elem()
}

func (i *RunbookProcess) ToRunbookProcessOutput() RunbookProcessOutput {
	return i.ToRunbookProcessOutputWithContext(context.Background())
}

func (i *RunbookProcess) ToRunbookProcessOutputWithContext(ctx context.Context) RunbookProcessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookProcessOutput)
}

// RunbookProcessArrayInput is an input type that accepts RunbookProcessArray and RunbookProcessArrayOutput values.
// You can construct a concrete instance of `RunbookProcessArrayInput` via:
//
//	RunbookProcessArray{ RunbookProcessArgs{...} }
type RunbookProcessArrayInput interface {
	pulumi.Input

	ToRunbookProcessArrayOutput() RunbookProcessArrayOutput
	ToRunbookProcessArrayOutputWithContext(context.Context) RunbookProcessArrayOutput
}

type RunbookProcessArray []RunbookProcessInput

func (RunbookProcessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RunbookProcess)(nil)).Elem()
}

func (i RunbookProcessArray) ToRunbookProcessArrayOutput() RunbookProcessArrayOutput {
	return i.ToRunbookProcessArrayOutputWithContext(context.Background())
}

func (i RunbookProcessArray) ToRunbookProcessArrayOutputWithContext(ctx context.Context) RunbookProcessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookProcessArrayOutput)
}

// RunbookProcessMapInput is an input type that accepts RunbookProcessMap and RunbookProcessMapOutput values.
// You can construct a concrete instance of `RunbookProcessMapInput` via:
//
//	RunbookProcessMap{ "key": RunbookProcessArgs{...} }
type RunbookProcessMapInput interface {
	pulumi.Input

	ToRunbookProcessMapOutput() RunbookProcessMapOutput
	ToRunbookProcessMapOutputWithContext(context.Context) RunbookProcessMapOutput
}

type RunbookProcessMap map[string]RunbookProcessInput

func (RunbookProcessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RunbookProcess)(nil)).Elem()
}

func (i RunbookProcessMap) ToRunbookProcessMapOutput() RunbookProcessMapOutput {
	return i.ToRunbookProcessMapOutputWithContext(context.Background())
}

func (i RunbookProcessMap) ToRunbookProcessMapOutputWithContext(ctx context.Context) RunbookProcessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookProcessMapOutput)
}

type RunbookProcessOutput struct{ *pulumi.OutputState }

func (RunbookProcessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookProcess)(nil)).Elem()
}

func (o RunbookProcessOutput) ToRunbookProcessOutput() RunbookProcessOutput {
	return o
}

func (o RunbookProcessOutput) ToRunbookProcessOutputWithContext(ctx context.Context) RunbookProcessOutput {
	return o
}

// Read only value containing the last snapshot ID.
func (o RunbookProcessOutput) LastSnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookProcess) pulumi.StringPtrOutput { return v.LastSnapshotId }).(pulumi.StringPtrOutput)
}

// The project ID associated with this runbook process.
func (o RunbookProcessOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *RunbookProcess) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The runbook ID associated with this runbook process.
func (o RunbookProcessOutput) RunbookId() pulumi.StringOutput {
	return o.ApplyT(func(v *RunbookProcess) pulumi.StringOutput { return v.RunbookId }).(pulumi.StringOutput)
}

// The space ID associated with this resource.
func (o RunbookProcessOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RunbookProcess) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

func (o RunbookProcessOutput) Steps() RunbookProcessStepArrayOutput {
	return o.ApplyT(func(v *RunbookProcess) RunbookProcessStepArrayOutput { return v.Steps }).(RunbookProcessStepArrayOutput)
}

// The version number of this runbook process.
func (o RunbookProcessOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *RunbookProcess) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type RunbookProcessArrayOutput struct{ *pulumi.OutputState }

func (RunbookProcessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RunbookProcess)(nil)).Elem()
}

func (o RunbookProcessArrayOutput) ToRunbookProcessArrayOutput() RunbookProcessArrayOutput {
	return o
}

func (o RunbookProcessArrayOutput) ToRunbookProcessArrayOutputWithContext(ctx context.Context) RunbookProcessArrayOutput {
	return o
}

func (o RunbookProcessArrayOutput) Index(i pulumi.IntInput) RunbookProcessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RunbookProcess {
		return vs[0].([]*RunbookProcess)[vs[1].(int)]
	}).(RunbookProcessOutput)
}

type RunbookProcessMapOutput struct{ *pulumi.OutputState }

func (RunbookProcessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RunbookProcess)(nil)).Elem()
}

func (o RunbookProcessMapOutput) ToRunbookProcessMapOutput() RunbookProcessMapOutput {
	return o
}

func (o RunbookProcessMapOutput) ToRunbookProcessMapOutputWithContext(ctx context.Context) RunbookProcessMapOutput {
	return o
}

func (o RunbookProcessMapOutput) MapIndex(k pulumi.StringInput) RunbookProcessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RunbookProcess {
		return vs[0].(map[string]*RunbookProcess)[vs[1].(string)]
	}).(RunbookProcessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookProcessInput)(nil)).Elem(), &RunbookProcess{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookProcessArrayInput)(nil)).Elem(), RunbookProcessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookProcessMapInput)(nil)).Elem(), RunbookProcessMap{})
	pulumi.RegisterOutputType(RunbookProcessOutput{})
	pulumi.RegisterOutputType(RunbookProcessArrayOutput{})
	pulumi.RegisterOutputType(RunbookProcessMapOutput{})
}
