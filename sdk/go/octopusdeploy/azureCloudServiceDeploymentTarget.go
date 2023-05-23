// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages Azure cloud service deployment targets in Octopus Deploy.
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
//			_, err := octopusdeploy.NewAzureCloudServiceDeploymentTarget(ctx, "example", &octopusdeploy.AzureCloudServiceDeploymentTargetArgs{
//				AccountId:        pulumi.String("Accounts-123"),
//				CloudServiceName: pulumi.String("[cloud_service_name]"),
//				Environments: pulumi.StringArray{
//					pulumi.String("Environments-123"),
//					pulumi.String("Environment-321"),
//				},
//				Roles: pulumi.StringArray{
//					pulumi.String("Development Team"),
//					pulumi.String("System Administrators"),
//				},
//				StorageAccountName:              pulumi.String("[storage_account_name]"),
//				TenantedDeploymentParticipation: pulumi.String("Untenanted"),
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
//	$ pulumi import octopusdeploy:index/azureCloudServiceDeploymentTarget:AzureCloudServiceDeploymentTarget [options] octopusdeploy_azure_cloud_service_deployment_target.<name> <machine-id>
//
// ```
type AzureCloudServiceDeploymentTarget struct {
	pulumi.CustomResourceState

	AccountId           pulumi.StringOutput                                  `pulumi:"accountId"`
	CloudServiceName    pulumi.StringOutput                                  `pulumi:"cloudServiceName"`
	DefaultWorkerPoolId pulumi.StringPtrOutput                               `pulumi:"defaultWorkerPoolId"`
	Endpoints           AzureCloudServiceDeploymentTargetEndpointArrayOutput `pulumi:"endpoints"`
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
	Slot            pulumi.StringPtrOutput   `pulumi:"slot"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status pulumi.StringOutput `pulumi:"status"`
	// A summary elaborating on the status of this resource.
	StatusSummary      pulumi.StringOutput  `pulumi:"statusSummary"`
	StorageAccountName pulumi.StringOutput  `pulumi:"storageAccountName"`
	SwapIfPossible     pulumi.BoolPtrOutput `pulumi:"swapIfPossible"`
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayOutput `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringOutput `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants                 pulumi.StringArrayOutput `pulumi:"tenants"`
	Thumbprint              pulumi.StringOutput      `pulumi:"thumbprint"`
	Uri                     pulumi.StringOutput      `pulumi:"uri"`
	UseCurrentInstanceCount pulumi.BoolPtrOutput     `pulumi:"useCurrentInstanceCount"`
}

// NewAzureCloudServiceDeploymentTarget registers a new resource with the given unique name, arguments, and options.
func NewAzureCloudServiceDeploymentTarget(ctx *pulumi.Context,
	name string, args *AzureCloudServiceDeploymentTargetArgs, opts ...pulumi.ResourceOption) (*AzureCloudServiceDeploymentTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.CloudServiceName == nil {
		return nil, errors.New("invalid value for required argument 'CloudServiceName'")
	}
	if args.Environments == nil {
		return nil, errors.New("invalid value for required argument 'Environments'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	var resource AzureCloudServiceDeploymentTarget
	err := ctx.RegisterResource("octopusdeploy:index/azureCloudServiceDeploymentTarget:AzureCloudServiceDeploymentTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureCloudServiceDeploymentTarget gets an existing AzureCloudServiceDeploymentTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureCloudServiceDeploymentTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureCloudServiceDeploymentTargetState, opts ...pulumi.ResourceOption) (*AzureCloudServiceDeploymentTarget, error) {
	var resource AzureCloudServiceDeploymentTarget
	err := ctx.ReadResource("octopusdeploy:index/azureCloudServiceDeploymentTarget:AzureCloudServiceDeploymentTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureCloudServiceDeploymentTarget resources.
type azureCloudServiceDeploymentTargetState struct {
	AccountId           *string                                     `pulumi:"accountId"`
	CloudServiceName    *string                                     `pulumi:"cloudServiceName"`
	DefaultWorkerPoolId *string                                     `pulumi:"defaultWorkerPoolId"`
	Endpoints           []AzureCloudServiceDeploymentTargetEndpoint `pulumi:"endpoints"`
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
	Slot            *string  `pulumi:"slot"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status *string `pulumi:"status"`
	// A summary elaborating on the status of this resource.
	StatusSummary      *string `pulumi:"statusSummary"`
	StorageAccountName *string `pulumi:"storageAccountName"`
	SwapIfPossible     *bool   `pulumi:"swapIfPossible"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants                 []string `pulumi:"tenants"`
	Thumbprint              *string  `pulumi:"thumbprint"`
	Uri                     *string  `pulumi:"uri"`
	UseCurrentInstanceCount *bool    `pulumi:"useCurrentInstanceCount"`
}

type AzureCloudServiceDeploymentTargetState struct {
	AccountId           pulumi.StringPtrInput
	CloudServiceName    pulumi.StringPtrInput
	DefaultWorkerPoolId pulumi.StringPtrInput
	Endpoints           AzureCloudServiceDeploymentTargetEndpointArrayInput
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
	Slot            pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status pulumi.StringPtrInput
	// A summary elaborating on the status of this resource.
	StatusSummary      pulumi.StringPtrInput
	StorageAccountName pulumi.StringPtrInput
	SwapIfPossible     pulumi.BoolPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants                 pulumi.StringArrayInput
	Thumbprint              pulumi.StringPtrInput
	Uri                     pulumi.StringPtrInput
	UseCurrentInstanceCount pulumi.BoolPtrInput
}

func (AzureCloudServiceDeploymentTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureCloudServiceDeploymentTargetState)(nil)).Elem()
}

type azureCloudServiceDeploymentTargetArgs struct {
	AccountId           string                                      `pulumi:"accountId"`
	CloudServiceName    string                                      `pulumi:"cloudServiceName"`
	DefaultWorkerPoolId *string                                     `pulumi:"defaultWorkerPoolId"`
	Endpoints           []AzureCloudServiceDeploymentTargetEndpoint `pulumi:"endpoints"`
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
	Slot            *string  `pulumi:"slot"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status *string `pulumi:"status"`
	// A summary elaborating on the status of this resource.
	StatusSummary      *string `pulumi:"statusSummary"`
	StorageAccountName string  `pulumi:"storageAccountName"`
	SwapIfPossible     *bool   `pulumi:"swapIfPossible"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants                 []string `pulumi:"tenants"`
	Thumbprint              *string  `pulumi:"thumbprint"`
	Uri                     *string  `pulumi:"uri"`
	UseCurrentInstanceCount *bool    `pulumi:"useCurrentInstanceCount"`
}

// The set of arguments for constructing a AzureCloudServiceDeploymentTarget resource.
type AzureCloudServiceDeploymentTargetArgs struct {
	AccountId           pulumi.StringInput
	CloudServiceName    pulumi.StringInput
	DefaultWorkerPoolId pulumi.StringPtrInput
	Endpoints           AzureCloudServiceDeploymentTargetEndpointArrayInput
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
	Slot            pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status pulumi.StringPtrInput
	// A summary elaborating on the status of this resource.
	StatusSummary      pulumi.StringPtrInput
	StorageAccountName pulumi.StringInput
	SwapIfPossible     pulumi.BoolPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants                 pulumi.StringArrayInput
	Thumbprint              pulumi.StringPtrInput
	Uri                     pulumi.StringPtrInput
	UseCurrentInstanceCount pulumi.BoolPtrInput
}

func (AzureCloudServiceDeploymentTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureCloudServiceDeploymentTargetArgs)(nil)).Elem()
}

type AzureCloudServiceDeploymentTargetInput interface {
	pulumi.Input

	ToAzureCloudServiceDeploymentTargetOutput() AzureCloudServiceDeploymentTargetOutput
	ToAzureCloudServiceDeploymentTargetOutputWithContext(ctx context.Context) AzureCloudServiceDeploymentTargetOutput
}

func (*AzureCloudServiceDeploymentTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCloudServiceDeploymentTarget)(nil)).Elem()
}

func (i *AzureCloudServiceDeploymentTarget) ToAzureCloudServiceDeploymentTargetOutput() AzureCloudServiceDeploymentTargetOutput {
	return i.ToAzureCloudServiceDeploymentTargetOutputWithContext(context.Background())
}

func (i *AzureCloudServiceDeploymentTarget) ToAzureCloudServiceDeploymentTargetOutputWithContext(ctx context.Context) AzureCloudServiceDeploymentTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureCloudServiceDeploymentTargetOutput)
}

// AzureCloudServiceDeploymentTargetArrayInput is an input type that accepts AzureCloudServiceDeploymentTargetArray and AzureCloudServiceDeploymentTargetArrayOutput values.
// You can construct a concrete instance of `AzureCloudServiceDeploymentTargetArrayInput` via:
//
//	AzureCloudServiceDeploymentTargetArray{ AzureCloudServiceDeploymentTargetArgs{...} }
type AzureCloudServiceDeploymentTargetArrayInput interface {
	pulumi.Input

	ToAzureCloudServiceDeploymentTargetArrayOutput() AzureCloudServiceDeploymentTargetArrayOutput
	ToAzureCloudServiceDeploymentTargetArrayOutputWithContext(context.Context) AzureCloudServiceDeploymentTargetArrayOutput
}

type AzureCloudServiceDeploymentTargetArray []AzureCloudServiceDeploymentTargetInput

func (AzureCloudServiceDeploymentTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureCloudServiceDeploymentTarget)(nil)).Elem()
}

func (i AzureCloudServiceDeploymentTargetArray) ToAzureCloudServiceDeploymentTargetArrayOutput() AzureCloudServiceDeploymentTargetArrayOutput {
	return i.ToAzureCloudServiceDeploymentTargetArrayOutputWithContext(context.Background())
}

func (i AzureCloudServiceDeploymentTargetArray) ToAzureCloudServiceDeploymentTargetArrayOutputWithContext(ctx context.Context) AzureCloudServiceDeploymentTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureCloudServiceDeploymentTargetArrayOutput)
}

// AzureCloudServiceDeploymentTargetMapInput is an input type that accepts AzureCloudServiceDeploymentTargetMap and AzureCloudServiceDeploymentTargetMapOutput values.
// You can construct a concrete instance of `AzureCloudServiceDeploymentTargetMapInput` via:
//
//	AzureCloudServiceDeploymentTargetMap{ "key": AzureCloudServiceDeploymentTargetArgs{...} }
type AzureCloudServiceDeploymentTargetMapInput interface {
	pulumi.Input

	ToAzureCloudServiceDeploymentTargetMapOutput() AzureCloudServiceDeploymentTargetMapOutput
	ToAzureCloudServiceDeploymentTargetMapOutputWithContext(context.Context) AzureCloudServiceDeploymentTargetMapOutput
}

type AzureCloudServiceDeploymentTargetMap map[string]AzureCloudServiceDeploymentTargetInput

func (AzureCloudServiceDeploymentTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureCloudServiceDeploymentTarget)(nil)).Elem()
}

func (i AzureCloudServiceDeploymentTargetMap) ToAzureCloudServiceDeploymentTargetMapOutput() AzureCloudServiceDeploymentTargetMapOutput {
	return i.ToAzureCloudServiceDeploymentTargetMapOutputWithContext(context.Background())
}

func (i AzureCloudServiceDeploymentTargetMap) ToAzureCloudServiceDeploymentTargetMapOutputWithContext(ctx context.Context) AzureCloudServiceDeploymentTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureCloudServiceDeploymentTargetMapOutput)
}

type AzureCloudServiceDeploymentTargetOutput struct{ *pulumi.OutputState }

func (AzureCloudServiceDeploymentTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCloudServiceDeploymentTarget)(nil)).Elem()
}

func (o AzureCloudServiceDeploymentTargetOutput) ToAzureCloudServiceDeploymentTargetOutput() AzureCloudServiceDeploymentTargetOutput {
	return o
}

func (o AzureCloudServiceDeploymentTargetOutput) ToAzureCloudServiceDeploymentTargetOutputWithContext(ctx context.Context) AzureCloudServiceDeploymentTargetOutput {
	return o
}

func (o AzureCloudServiceDeploymentTargetOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) CloudServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.CloudServiceName }).(pulumi.StringOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) DefaultWorkerPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringPtrOutput { return v.DefaultWorkerPoolId }).(pulumi.StringPtrOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) Endpoints() AzureCloudServiceDeploymentTargetEndpointArrayOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) AzureCloudServiceDeploymentTargetEndpointArrayOutput {
		return v.Endpoints
	}).(AzureCloudServiceDeploymentTargetEndpointArrayOutput)
}

// A list of environment IDs associated with this resource.
func (o AzureCloudServiceDeploymentTargetOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringArrayOutput { return v.Environments }).(pulumi.StringArrayOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) HasLatestCalamari() pulumi.BoolOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.BoolOutput { return v.HasLatestCalamari }).(pulumi.BoolOutput)
}

// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
func (o AzureCloudServiceDeploymentTargetOutput) HealthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.HealthStatus }).(pulumi.StringOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.BoolOutput { return v.IsDisabled }).(pulumi.BoolOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) IsInProcess() pulumi.BoolOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.BoolOutput { return v.IsInProcess }).(pulumi.BoolOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) MachinePolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.MachinePolicyId }).(pulumi.StringOutput)
}

// The name of this resource.
func (o AzureCloudServiceDeploymentTargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) OperatingSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.OperatingSystem }).(pulumi.StringOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) ShellName() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.ShellName }).(pulumi.StringOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) ShellVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.ShellVersion }).(pulumi.StringOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) Slot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringPtrOutput { return v.Slot }).(pulumi.StringPtrOutput)
}

// The space ID associated with this resource.
func (o AzureCloudServiceDeploymentTargetOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
func (o AzureCloudServiceDeploymentTargetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A summary elaborating on the status of this resource.
func (o AzureCloudServiceDeploymentTargetOutput) StatusSummary() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.StatusSummary }).(pulumi.StringOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) SwapIfPossible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.BoolPtrOutput { return v.SwapIfPossible }).(pulumi.BoolPtrOutput)
}

// A list of tenant tags associated with this resource.
func (o AzureCloudServiceDeploymentTargetOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringArrayOutput { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
func (o AzureCloudServiceDeploymentTargetOutput) TenantedDeploymentParticipation() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput {
		return v.TenantedDeploymentParticipation
	}).(pulumi.StringOutput)
}

// A list of tenant IDs associated with this resource.
func (o AzureCloudServiceDeploymentTargetOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringArrayOutput { return v.Tenants }).(pulumi.StringArrayOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

func (o AzureCloudServiceDeploymentTargetOutput) UseCurrentInstanceCount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureCloudServiceDeploymentTarget) pulumi.BoolPtrOutput { return v.UseCurrentInstanceCount }).(pulumi.BoolPtrOutput)
}

type AzureCloudServiceDeploymentTargetArrayOutput struct{ *pulumi.OutputState }

func (AzureCloudServiceDeploymentTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureCloudServiceDeploymentTarget)(nil)).Elem()
}

func (o AzureCloudServiceDeploymentTargetArrayOutput) ToAzureCloudServiceDeploymentTargetArrayOutput() AzureCloudServiceDeploymentTargetArrayOutput {
	return o
}

func (o AzureCloudServiceDeploymentTargetArrayOutput) ToAzureCloudServiceDeploymentTargetArrayOutputWithContext(ctx context.Context) AzureCloudServiceDeploymentTargetArrayOutput {
	return o
}

func (o AzureCloudServiceDeploymentTargetArrayOutput) Index(i pulumi.IntInput) AzureCloudServiceDeploymentTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AzureCloudServiceDeploymentTarget {
		return vs[0].([]*AzureCloudServiceDeploymentTarget)[vs[1].(int)]
	}).(AzureCloudServiceDeploymentTargetOutput)
}

type AzureCloudServiceDeploymentTargetMapOutput struct{ *pulumi.OutputState }

func (AzureCloudServiceDeploymentTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureCloudServiceDeploymentTarget)(nil)).Elem()
}

func (o AzureCloudServiceDeploymentTargetMapOutput) ToAzureCloudServiceDeploymentTargetMapOutput() AzureCloudServiceDeploymentTargetMapOutput {
	return o
}

func (o AzureCloudServiceDeploymentTargetMapOutput) ToAzureCloudServiceDeploymentTargetMapOutputWithContext(ctx context.Context) AzureCloudServiceDeploymentTargetMapOutput {
	return o
}

func (o AzureCloudServiceDeploymentTargetMapOutput) MapIndex(k pulumi.StringInput) AzureCloudServiceDeploymentTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AzureCloudServiceDeploymentTarget {
		return vs[0].(map[string]*AzureCloudServiceDeploymentTarget)[vs[1].(string)]
	}).(AzureCloudServiceDeploymentTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureCloudServiceDeploymentTargetInput)(nil)).Elem(), &AzureCloudServiceDeploymentTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureCloudServiceDeploymentTargetArrayInput)(nil)).Elem(), AzureCloudServiceDeploymentTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureCloudServiceDeploymentTargetMapInput)(nil)).Elem(), AzureCloudServiceDeploymentTargetMap{})
	pulumi.RegisterOutputType(AzureCloudServiceDeploymentTargetOutput{})
	pulumi.RegisterOutputType(AzureCloudServiceDeploymentTargetArrayOutput{})
	pulumi.RegisterOutputType(AzureCloudServiceDeploymentTargetMapOutput{})
}