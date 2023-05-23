// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages offline package drop deployment targets in Octopus Deploy.
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
//			_, err := octopusdeploy.NewListeningTentacleDeploymentTarget(ctx, "example", &octopusdeploy.ListeningTentacleDeploymentTargetArgs{
//				Environments: pulumi.StringArray{
//					pulumi.String("Environments-123"),
//					pulumi.String("Environment-321"),
//				},
//				IsDisabled:      pulumi.Bool(true),
//				MachinePolicyId: pulumi.String("MachinePolicy-123"),
//				Roles: pulumi.StringArray{
//					pulumi.String("Development Team"),
//					pulumi.String("System Administrators"),
//				},
//				TenantedDeploymentParticipation: pulumi.String("Untenanted"),
//				TentacleUrl:                     pulumi.String("https://example.com:1234/"),
//				Thumbprint:                      pulumi.String("<thumbprint>"),
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
//	$ pulumi import octopusdeploy:index/offlinePackageDropDeploymentTarget:OfflinePackageDropDeploymentTarget [options] octopusdeploy_listening_tentacle_deployment_target.<name> <machine-id>
//
// ```
type OfflinePackageDropDeploymentTarget struct {
	pulumi.CustomResourceState

	ApplicationsDirectory pulumi.StringOutput                                   `pulumi:"applicationsDirectory"`
	Destination           OfflinePackageDropDeploymentTargetDestinationOutput   `pulumi:"destination"`
	Endpoints             OfflinePackageDropDeploymentTargetEndpointArrayOutput `pulumi:"endpoints"`
	// A list of environment IDs associated with this resource.
	Environments      pulumi.StringArrayOutput `pulumi:"environments"`
	HasLatestCalamari pulumi.BoolOutput        `pulumi:"hasLatestCalamari"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    pulumi.StringOutput `pulumi:"healthStatus"`
	IsDisabled      pulumi.BoolOutput   `pulumi:"isDisabled"`
	IsInProcess     pulumi.BoolOutput   `pulumi:"isInProcess"`
	MachinePolicyId pulumi.StringOutput `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name            pulumi.StringOutput      `pulumi:"name"`
	OperatingSystem pulumi.StringOutput      `pulumi:"operatingSystem"`
	Roles           pulumi.StringArrayOutput `pulumi:"roles"`
	ShellName       pulumi.StringOutput      `pulumi:"shellName"`
	ShellVersion    pulumi.StringOutput      `pulumi:"shellVersion"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status pulumi.StringOutput `pulumi:"status"`
	// A summary elaborating on the status of this resource.
	StatusSummary pulumi.StringOutput `pulumi:"statusSummary"`
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayOutput `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringOutput `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants          pulumi.StringArrayOutput `pulumi:"tenants"`
	Thumbprint       pulumi.StringOutput      `pulumi:"thumbprint"`
	Uri              pulumi.StringOutput      `pulumi:"uri"`
	WorkingDirectory pulumi.StringOutput      `pulumi:"workingDirectory"`
}

// NewOfflinePackageDropDeploymentTarget registers a new resource with the given unique name, arguments, and options.
func NewOfflinePackageDropDeploymentTarget(ctx *pulumi.Context,
	name string, args *OfflinePackageDropDeploymentTargetArgs, opts ...pulumi.ResourceOption) (*OfflinePackageDropDeploymentTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationsDirectory == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationsDirectory'")
	}
	if args.Environments == nil {
		return nil, errors.New("invalid value for required argument 'Environments'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.WorkingDirectory == nil {
		return nil, errors.New("invalid value for required argument 'WorkingDirectory'")
	}
	var resource OfflinePackageDropDeploymentTarget
	err := ctx.RegisterResource("octopusdeploy:index/offlinePackageDropDeploymentTarget:OfflinePackageDropDeploymentTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOfflinePackageDropDeploymentTarget gets an existing OfflinePackageDropDeploymentTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOfflinePackageDropDeploymentTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OfflinePackageDropDeploymentTargetState, opts ...pulumi.ResourceOption) (*OfflinePackageDropDeploymentTarget, error) {
	var resource OfflinePackageDropDeploymentTarget
	err := ctx.ReadResource("octopusdeploy:index/offlinePackageDropDeploymentTarget:OfflinePackageDropDeploymentTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OfflinePackageDropDeploymentTarget resources.
type offlinePackageDropDeploymentTargetState struct {
	ApplicationsDirectory *string                                        `pulumi:"applicationsDirectory"`
	Destination           *OfflinePackageDropDeploymentTargetDestination `pulumi:"destination"`
	Endpoints             []OfflinePackageDropDeploymentTargetEndpoint   `pulumi:"endpoints"`
	// A list of environment IDs associated with this resource.
	Environments      []string `pulumi:"environments"`
	HasLatestCalamari *bool    `pulumi:"hasLatestCalamari"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    *string `pulumi:"healthStatus"`
	IsDisabled      *bool   `pulumi:"isDisabled"`
	IsInProcess     *bool   `pulumi:"isInProcess"`
	MachinePolicyId *string `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name            *string  `pulumi:"name"`
	OperatingSystem *string  `pulumi:"operatingSystem"`
	Roles           []string `pulumi:"roles"`
	ShellName       *string  `pulumi:"shellName"`
	ShellVersion    *string  `pulumi:"shellVersion"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status *string `pulumi:"status"`
	// A summary elaborating on the status of this resource.
	StatusSummary *string `pulumi:"statusSummary"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants          []string `pulumi:"tenants"`
	Thumbprint       *string  `pulumi:"thumbprint"`
	Uri              *string  `pulumi:"uri"`
	WorkingDirectory *string  `pulumi:"workingDirectory"`
}

type OfflinePackageDropDeploymentTargetState struct {
	ApplicationsDirectory pulumi.StringPtrInput
	Destination           OfflinePackageDropDeploymentTargetDestinationPtrInput
	Endpoints             OfflinePackageDropDeploymentTargetEndpointArrayInput
	// A list of environment IDs associated with this resource.
	Environments      pulumi.StringArrayInput
	HasLatestCalamari pulumi.BoolPtrInput
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    pulumi.StringPtrInput
	IsDisabled      pulumi.BoolPtrInput
	IsInProcess     pulumi.BoolPtrInput
	MachinePolicyId pulumi.StringPtrInput
	// The name of this resource.
	Name            pulumi.StringPtrInput
	OperatingSystem pulumi.StringPtrInput
	Roles           pulumi.StringArrayInput
	ShellName       pulumi.StringPtrInput
	ShellVersion    pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status pulumi.StringPtrInput
	// A summary elaborating on the status of this resource.
	StatusSummary pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants          pulumi.StringArrayInput
	Thumbprint       pulumi.StringPtrInput
	Uri              pulumi.StringPtrInput
	WorkingDirectory pulumi.StringPtrInput
}

func (OfflinePackageDropDeploymentTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*offlinePackageDropDeploymentTargetState)(nil)).Elem()
}

type offlinePackageDropDeploymentTargetArgs struct {
	ApplicationsDirectory string                                         `pulumi:"applicationsDirectory"`
	Destination           *OfflinePackageDropDeploymentTargetDestination `pulumi:"destination"`
	Endpoints             []OfflinePackageDropDeploymentTargetEndpoint   `pulumi:"endpoints"`
	// A list of environment IDs associated with this resource.
	Environments []string `pulumi:"environments"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    *string `pulumi:"healthStatus"`
	IsDisabled      *bool   `pulumi:"isDisabled"`
	MachinePolicyId *string `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name            *string  `pulumi:"name"`
	OperatingSystem *string  `pulumi:"operatingSystem"`
	Roles           []string `pulumi:"roles"`
	ShellName       *string  `pulumi:"shellName"`
	ShellVersion    *string  `pulumi:"shellVersion"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status *string `pulumi:"status"`
	// A summary elaborating on the status of this resource.
	StatusSummary *string `pulumi:"statusSummary"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants          []string `pulumi:"tenants"`
	Thumbprint       *string  `pulumi:"thumbprint"`
	Uri              *string  `pulumi:"uri"`
	WorkingDirectory string   `pulumi:"workingDirectory"`
}

// The set of arguments for constructing a OfflinePackageDropDeploymentTarget resource.
type OfflinePackageDropDeploymentTargetArgs struct {
	ApplicationsDirectory pulumi.StringInput
	Destination           OfflinePackageDropDeploymentTargetDestinationPtrInput
	Endpoints             OfflinePackageDropDeploymentTargetEndpointArrayInput
	// A list of environment IDs associated with this resource.
	Environments pulumi.StringArrayInput
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    pulumi.StringPtrInput
	IsDisabled      pulumi.BoolPtrInput
	MachinePolicyId pulumi.StringPtrInput
	// The name of this resource.
	Name            pulumi.StringPtrInput
	OperatingSystem pulumi.StringPtrInput
	Roles           pulumi.StringArrayInput
	ShellName       pulumi.StringPtrInput
	ShellVersion    pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status pulumi.StringPtrInput
	// A summary elaborating on the status of this resource.
	StatusSummary pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants          pulumi.StringArrayInput
	Thumbprint       pulumi.StringPtrInput
	Uri              pulumi.StringPtrInput
	WorkingDirectory pulumi.StringInput
}

func (OfflinePackageDropDeploymentTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*offlinePackageDropDeploymentTargetArgs)(nil)).Elem()
}

type OfflinePackageDropDeploymentTargetInput interface {
	pulumi.Input

	ToOfflinePackageDropDeploymentTargetOutput() OfflinePackageDropDeploymentTargetOutput
	ToOfflinePackageDropDeploymentTargetOutputWithContext(ctx context.Context) OfflinePackageDropDeploymentTargetOutput
}

func (*OfflinePackageDropDeploymentTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**OfflinePackageDropDeploymentTarget)(nil)).Elem()
}

func (i *OfflinePackageDropDeploymentTarget) ToOfflinePackageDropDeploymentTargetOutput() OfflinePackageDropDeploymentTargetOutput {
	return i.ToOfflinePackageDropDeploymentTargetOutputWithContext(context.Background())
}

func (i *OfflinePackageDropDeploymentTarget) ToOfflinePackageDropDeploymentTargetOutputWithContext(ctx context.Context) OfflinePackageDropDeploymentTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfflinePackageDropDeploymentTargetOutput)
}

// OfflinePackageDropDeploymentTargetArrayInput is an input type that accepts OfflinePackageDropDeploymentTargetArray and OfflinePackageDropDeploymentTargetArrayOutput values.
// You can construct a concrete instance of `OfflinePackageDropDeploymentTargetArrayInput` via:
//
//	OfflinePackageDropDeploymentTargetArray{ OfflinePackageDropDeploymentTargetArgs{...} }
type OfflinePackageDropDeploymentTargetArrayInput interface {
	pulumi.Input

	ToOfflinePackageDropDeploymentTargetArrayOutput() OfflinePackageDropDeploymentTargetArrayOutput
	ToOfflinePackageDropDeploymentTargetArrayOutputWithContext(context.Context) OfflinePackageDropDeploymentTargetArrayOutput
}

type OfflinePackageDropDeploymentTargetArray []OfflinePackageDropDeploymentTargetInput

func (OfflinePackageDropDeploymentTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OfflinePackageDropDeploymentTarget)(nil)).Elem()
}

func (i OfflinePackageDropDeploymentTargetArray) ToOfflinePackageDropDeploymentTargetArrayOutput() OfflinePackageDropDeploymentTargetArrayOutput {
	return i.ToOfflinePackageDropDeploymentTargetArrayOutputWithContext(context.Background())
}

func (i OfflinePackageDropDeploymentTargetArray) ToOfflinePackageDropDeploymentTargetArrayOutputWithContext(ctx context.Context) OfflinePackageDropDeploymentTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfflinePackageDropDeploymentTargetArrayOutput)
}

// OfflinePackageDropDeploymentTargetMapInput is an input type that accepts OfflinePackageDropDeploymentTargetMap and OfflinePackageDropDeploymentTargetMapOutput values.
// You can construct a concrete instance of `OfflinePackageDropDeploymentTargetMapInput` via:
//
//	OfflinePackageDropDeploymentTargetMap{ "key": OfflinePackageDropDeploymentTargetArgs{...} }
type OfflinePackageDropDeploymentTargetMapInput interface {
	pulumi.Input

	ToOfflinePackageDropDeploymentTargetMapOutput() OfflinePackageDropDeploymentTargetMapOutput
	ToOfflinePackageDropDeploymentTargetMapOutputWithContext(context.Context) OfflinePackageDropDeploymentTargetMapOutput
}

type OfflinePackageDropDeploymentTargetMap map[string]OfflinePackageDropDeploymentTargetInput

func (OfflinePackageDropDeploymentTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OfflinePackageDropDeploymentTarget)(nil)).Elem()
}

func (i OfflinePackageDropDeploymentTargetMap) ToOfflinePackageDropDeploymentTargetMapOutput() OfflinePackageDropDeploymentTargetMapOutput {
	return i.ToOfflinePackageDropDeploymentTargetMapOutputWithContext(context.Background())
}

func (i OfflinePackageDropDeploymentTargetMap) ToOfflinePackageDropDeploymentTargetMapOutputWithContext(ctx context.Context) OfflinePackageDropDeploymentTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfflinePackageDropDeploymentTargetMapOutput)
}

type OfflinePackageDropDeploymentTargetOutput struct{ *pulumi.OutputState }

func (OfflinePackageDropDeploymentTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfflinePackageDropDeploymentTarget)(nil)).Elem()
}

func (o OfflinePackageDropDeploymentTargetOutput) ToOfflinePackageDropDeploymentTargetOutput() OfflinePackageDropDeploymentTargetOutput {
	return o
}

func (o OfflinePackageDropDeploymentTargetOutput) ToOfflinePackageDropDeploymentTargetOutputWithContext(ctx context.Context) OfflinePackageDropDeploymentTargetOutput {
	return o
}

func (o OfflinePackageDropDeploymentTargetOutput) ApplicationsDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.ApplicationsDirectory }).(pulumi.StringOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) Destination() OfflinePackageDropDeploymentTargetDestinationOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) OfflinePackageDropDeploymentTargetDestinationOutput {
		return v.Destination
	}).(OfflinePackageDropDeploymentTargetDestinationOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) Endpoints() OfflinePackageDropDeploymentTargetEndpointArrayOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) OfflinePackageDropDeploymentTargetEndpointArrayOutput {
		return v.Endpoints
	}).(OfflinePackageDropDeploymentTargetEndpointArrayOutput)
}

// A list of environment IDs associated with this resource.
func (o OfflinePackageDropDeploymentTargetOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringArrayOutput { return v.Environments }).(pulumi.StringArrayOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) HasLatestCalamari() pulumi.BoolOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.BoolOutput { return v.HasLatestCalamari }).(pulumi.BoolOutput)
}

// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
func (o OfflinePackageDropDeploymentTargetOutput) HealthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.HealthStatus }).(pulumi.StringOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.BoolOutput { return v.IsDisabled }).(pulumi.BoolOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) IsInProcess() pulumi.BoolOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.BoolOutput { return v.IsInProcess }).(pulumi.BoolOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) MachinePolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.MachinePolicyId }).(pulumi.StringOutput)
}

// The name of this resource.
func (o OfflinePackageDropDeploymentTargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) OperatingSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.OperatingSystem }).(pulumi.StringOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) ShellName() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.ShellName }).(pulumi.StringOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) ShellVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.ShellVersion }).(pulumi.StringOutput)
}

// The space ID associated with this resource.
func (o OfflinePackageDropDeploymentTargetOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
func (o OfflinePackageDropDeploymentTargetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A summary elaborating on the status of this resource.
func (o OfflinePackageDropDeploymentTargetOutput) StatusSummary() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.StatusSummary }).(pulumi.StringOutput)
}

// A list of tenant tags associated with this resource.
func (o OfflinePackageDropDeploymentTargetOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringArrayOutput { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
func (o OfflinePackageDropDeploymentTargetOutput) TenantedDeploymentParticipation() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput {
		return v.TenantedDeploymentParticipation
	}).(pulumi.StringOutput)
}

// A list of tenant IDs associated with this resource.
func (o OfflinePackageDropDeploymentTargetOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringArrayOutput { return v.Tenants }).(pulumi.StringArrayOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

func (o OfflinePackageDropDeploymentTargetOutput) WorkingDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v *OfflinePackageDropDeploymentTarget) pulumi.StringOutput { return v.WorkingDirectory }).(pulumi.StringOutput)
}

type OfflinePackageDropDeploymentTargetArrayOutput struct{ *pulumi.OutputState }

func (OfflinePackageDropDeploymentTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OfflinePackageDropDeploymentTarget)(nil)).Elem()
}

func (o OfflinePackageDropDeploymentTargetArrayOutput) ToOfflinePackageDropDeploymentTargetArrayOutput() OfflinePackageDropDeploymentTargetArrayOutput {
	return o
}

func (o OfflinePackageDropDeploymentTargetArrayOutput) ToOfflinePackageDropDeploymentTargetArrayOutputWithContext(ctx context.Context) OfflinePackageDropDeploymentTargetArrayOutput {
	return o
}

func (o OfflinePackageDropDeploymentTargetArrayOutput) Index(i pulumi.IntInput) OfflinePackageDropDeploymentTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OfflinePackageDropDeploymentTarget {
		return vs[0].([]*OfflinePackageDropDeploymentTarget)[vs[1].(int)]
	}).(OfflinePackageDropDeploymentTargetOutput)
}

type OfflinePackageDropDeploymentTargetMapOutput struct{ *pulumi.OutputState }

func (OfflinePackageDropDeploymentTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OfflinePackageDropDeploymentTarget)(nil)).Elem()
}

func (o OfflinePackageDropDeploymentTargetMapOutput) ToOfflinePackageDropDeploymentTargetMapOutput() OfflinePackageDropDeploymentTargetMapOutput {
	return o
}

func (o OfflinePackageDropDeploymentTargetMapOutput) ToOfflinePackageDropDeploymentTargetMapOutputWithContext(ctx context.Context) OfflinePackageDropDeploymentTargetMapOutput {
	return o
}

func (o OfflinePackageDropDeploymentTargetMapOutput) MapIndex(k pulumi.StringInput) OfflinePackageDropDeploymentTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OfflinePackageDropDeploymentTarget {
		return vs[0].(map[string]*OfflinePackageDropDeploymentTarget)[vs[1].(string)]
	}).(OfflinePackageDropDeploymentTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OfflinePackageDropDeploymentTargetInput)(nil)).Elem(), &OfflinePackageDropDeploymentTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfflinePackageDropDeploymentTargetArrayInput)(nil)).Elem(), OfflinePackageDropDeploymentTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfflinePackageDropDeploymentTargetMapInput)(nil)).Elem(), OfflinePackageDropDeploymentTargetMap{})
	pulumi.RegisterOutputType(OfflinePackageDropDeploymentTargetOutput{})
	pulumi.RegisterOutputType(OfflinePackageDropDeploymentTargetArrayOutput{})
	pulumi.RegisterOutputType(OfflinePackageDropDeploymentTargetMapOutput{})
}
