// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages Azure web app deployment targets in Octopus Deploy.
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
//			_, err := octopusdeploy.NewAzureWebAppDeploymentTarget(ctx, "example", &octopusdeploy.AzureWebAppDeploymentTargetArgs{
//				AccountId: pulumi.String("Accounts-123"),
//				Environments: pulumi.StringArray{
//					pulumi.String("Environments-123"),
//				},
//				ResourceGroupName: pulumi.String("resource-group-name"),
//				Roles: pulumi.StringArray{
//					pulumi.String("Development Team"),
//					pulumi.String("System Administrators"),
//				},
//				TenantedDeploymentParticipation: pulumi.String("Untenanted"),
//				WebAppName:                      pulumi.String("web-app-name"),
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
//	$ pulumi import octopusdeploy:index/azureWebAppDeploymentTarget:AzureWebAppDeploymentTarget [options] octopusdeploy_azure_web_app_deployment_target.<name> <machine-id>
//
// ```
type AzureWebAppDeploymentTarget struct {
	pulumi.CustomResourceState

	AccountId pulumi.StringOutput                            `pulumi:"accountId"`
	Endpoints AzureWebAppDeploymentTargetEndpointArrayOutput `pulumi:"endpoints"`
	// A list of environment IDs associated with this resource.
	Environments      pulumi.StringArrayOutput `pulumi:"environments"`
	HasLatestCalamari pulumi.BoolOutput        `pulumi:"hasLatestCalamari"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    pulumi.StringOutput `pulumi:"healthStatus"`
	IsDisabled      pulumi.BoolOutput   `pulumi:"isDisabled"`
	IsInProcess     pulumi.BoolOutput   `pulumi:"isInProcess"`
	MachinePolicyId pulumi.StringOutput `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name              pulumi.StringOutput      `pulumi:"name"`
	OperatingSystem   pulumi.StringOutput      `pulumi:"operatingSystem"`
	ResourceGroupName pulumi.StringOutput      `pulumi:"resourceGroupName"`
	Roles             pulumi.StringArrayOutput `pulumi:"roles"`
	ShellName         pulumi.StringOutput      `pulumi:"shellName"`
	ShellVersion      pulumi.StringOutput      `pulumi:"shellVersion"`
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
	Tenants        pulumi.StringArrayOutput `pulumi:"tenants"`
	Thumbprint     pulumi.StringOutput      `pulumi:"thumbprint"`
	Uri            pulumi.StringOutput      `pulumi:"uri"`
	WebAppName     pulumi.StringOutput      `pulumi:"webAppName"`
	WebAppSlotName pulumi.StringPtrOutput   `pulumi:"webAppSlotName"`
}

// NewAzureWebAppDeploymentTarget registers a new resource with the given unique name, arguments, and options.
func NewAzureWebAppDeploymentTarget(ctx *pulumi.Context,
	name string, args *AzureWebAppDeploymentTargetArgs, opts ...pulumi.ResourceOption) (*AzureWebAppDeploymentTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.Environments == nil {
		return nil, errors.New("invalid value for required argument 'Environments'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.WebAppName == nil {
		return nil, errors.New("invalid value for required argument 'WebAppName'")
	}
	var resource AzureWebAppDeploymentTarget
	err := ctx.RegisterResource("octopusdeploy:index/azureWebAppDeploymentTarget:AzureWebAppDeploymentTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureWebAppDeploymentTarget gets an existing AzureWebAppDeploymentTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureWebAppDeploymentTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureWebAppDeploymentTargetState, opts ...pulumi.ResourceOption) (*AzureWebAppDeploymentTarget, error) {
	var resource AzureWebAppDeploymentTarget
	err := ctx.ReadResource("octopusdeploy:index/azureWebAppDeploymentTarget:AzureWebAppDeploymentTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureWebAppDeploymentTarget resources.
type azureWebAppDeploymentTargetState struct {
	AccountId *string                               `pulumi:"accountId"`
	Endpoints []AzureWebAppDeploymentTargetEndpoint `pulumi:"endpoints"`
	// A list of environment IDs associated with this resource.
	Environments      []string `pulumi:"environments"`
	HasLatestCalamari *bool    `pulumi:"hasLatestCalamari"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    *string `pulumi:"healthStatus"`
	IsDisabled      *bool   `pulumi:"isDisabled"`
	IsInProcess     *bool   `pulumi:"isInProcess"`
	MachinePolicyId *string `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name              *string  `pulumi:"name"`
	OperatingSystem   *string  `pulumi:"operatingSystem"`
	ResourceGroupName *string  `pulumi:"resourceGroupName"`
	Roles             []string `pulumi:"roles"`
	ShellName         *string  `pulumi:"shellName"`
	ShellVersion      *string  `pulumi:"shellVersion"`
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
	Tenants        []string `pulumi:"tenants"`
	Thumbprint     *string  `pulumi:"thumbprint"`
	Uri            *string  `pulumi:"uri"`
	WebAppName     *string  `pulumi:"webAppName"`
	WebAppSlotName *string  `pulumi:"webAppSlotName"`
}

type AzureWebAppDeploymentTargetState struct {
	AccountId pulumi.StringPtrInput
	Endpoints AzureWebAppDeploymentTargetEndpointArrayInput
	// A list of environment IDs associated with this resource.
	Environments      pulumi.StringArrayInput
	HasLatestCalamari pulumi.BoolPtrInput
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    pulumi.StringPtrInput
	IsDisabled      pulumi.BoolPtrInput
	IsInProcess     pulumi.BoolPtrInput
	MachinePolicyId pulumi.StringPtrInput
	// The name of this resource.
	Name              pulumi.StringPtrInput
	OperatingSystem   pulumi.StringPtrInput
	ResourceGroupName pulumi.StringPtrInput
	Roles             pulumi.StringArrayInput
	ShellName         pulumi.StringPtrInput
	ShellVersion      pulumi.StringPtrInput
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
	Tenants        pulumi.StringArrayInput
	Thumbprint     pulumi.StringPtrInput
	Uri            pulumi.StringPtrInput
	WebAppName     pulumi.StringPtrInput
	WebAppSlotName pulumi.StringPtrInput
}

func (AzureWebAppDeploymentTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureWebAppDeploymentTargetState)(nil)).Elem()
}

type azureWebAppDeploymentTargetArgs struct {
	AccountId string                                `pulumi:"accountId"`
	Endpoints []AzureWebAppDeploymentTargetEndpoint `pulumi:"endpoints"`
	// A list of environment IDs associated with this resource.
	Environments []string `pulumi:"environments"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    *string `pulumi:"healthStatus"`
	IsDisabled      *bool   `pulumi:"isDisabled"`
	MachinePolicyId *string `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name              *string  `pulumi:"name"`
	OperatingSystem   *string  `pulumi:"operatingSystem"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Roles             []string `pulumi:"roles"`
	ShellName         *string  `pulumi:"shellName"`
	ShellVersion      *string  `pulumi:"shellVersion"`
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
	Tenants        []string `pulumi:"tenants"`
	Thumbprint     *string  `pulumi:"thumbprint"`
	Uri            *string  `pulumi:"uri"`
	WebAppName     string   `pulumi:"webAppName"`
	WebAppSlotName *string  `pulumi:"webAppSlotName"`
}

// The set of arguments for constructing a AzureWebAppDeploymentTarget resource.
type AzureWebAppDeploymentTargetArgs struct {
	AccountId pulumi.StringInput
	Endpoints AzureWebAppDeploymentTargetEndpointArrayInput
	// A list of environment IDs associated with this resource.
	Environments pulumi.StringArrayInput
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    pulumi.StringPtrInput
	IsDisabled      pulumi.BoolPtrInput
	MachinePolicyId pulumi.StringPtrInput
	// The name of this resource.
	Name              pulumi.StringPtrInput
	OperatingSystem   pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Roles             pulumi.StringArrayInput
	ShellName         pulumi.StringPtrInput
	ShellVersion      pulumi.StringPtrInput
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
	Tenants        pulumi.StringArrayInput
	Thumbprint     pulumi.StringPtrInput
	Uri            pulumi.StringPtrInput
	WebAppName     pulumi.StringInput
	WebAppSlotName pulumi.StringPtrInput
}

func (AzureWebAppDeploymentTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureWebAppDeploymentTargetArgs)(nil)).Elem()
}

type AzureWebAppDeploymentTargetInput interface {
	pulumi.Input

	ToAzureWebAppDeploymentTargetOutput() AzureWebAppDeploymentTargetOutput
	ToAzureWebAppDeploymentTargetOutputWithContext(ctx context.Context) AzureWebAppDeploymentTargetOutput
}

func (*AzureWebAppDeploymentTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureWebAppDeploymentTarget)(nil)).Elem()
}

func (i *AzureWebAppDeploymentTarget) ToAzureWebAppDeploymentTargetOutput() AzureWebAppDeploymentTargetOutput {
	return i.ToAzureWebAppDeploymentTargetOutputWithContext(context.Background())
}

func (i *AzureWebAppDeploymentTarget) ToAzureWebAppDeploymentTargetOutputWithContext(ctx context.Context) AzureWebAppDeploymentTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWebAppDeploymentTargetOutput)
}

// AzureWebAppDeploymentTargetArrayInput is an input type that accepts AzureWebAppDeploymentTargetArray and AzureWebAppDeploymentTargetArrayOutput values.
// You can construct a concrete instance of `AzureWebAppDeploymentTargetArrayInput` via:
//
//	AzureWebAppDeploymentTargetArray{ AzureWebAppDeploymentTargetArgs{...} }
type AzureWebAppDeploymentTargetArrayInput interface {
	pulumi.Input

	ToAzureWebAppDeploymentTargetArrayOutput() AzureWebAppDeploymentTargetArrayOutput
	ToAzureWebAppDeploymentTargetArrayOutputWithContext(context.Context) AzureWebAppDeploymentTargetArrayOutput
}

type AzureWebAppDeploymentTargetArray []AzureWebAppDeploymentTargetInput

func (AzureWebAppDeploymentTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureWebAppDeploymentTarget)(nil)).Elem()
}

func (i AzureWebAppDeploymentTargetArray) ToAzureWebAppDeploymentTargetArrayOutput() AzureWebAppDeploymentTargetArrayOutput {
	return i.ToAzureWebAppDeploymentTargetArrayOutputWithContext(context.Background())
}

func (i AzureWebAppDeploymentTargetArray) ToAzureWebAppDeploymentTargetArrayOutputWithContext(ctx context.Context) AzureWebAppDeploymentTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWebAppDeploymentTargetArrayOutput)
}

// AzureWebAppDeploymentTargetMapInput is an input type that accepts AzureWebAppDeploymentTargetMap and AzureWebAppDeploymentTargetMapOutput values.
// You can construct a concrete instance of `AzureWebAppDeploymentTargetMapInput` via:
//
//	AzureWebAppDeploymentTargetMap{ "key": AzureWebAppDeploymentTargetArgs{...} }
type AzureWebAppDeploymentTargetMapInput interface {
	pulumi.Input

	ToAzureWebAppDeploymentTargetMapOutput() AzureWebAppDeploymentTargetMapOutput
	ToAzureWebAppDeploymentTargetMapOutputWithContext(context.Context) AzureWebAppDeploymentTargetMapOutput
}

type AzureWebAppDeploymentTargetMap map[string]AzureWebAppDeploymentTargetInput

func (AzureWebAppDeploymentTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureWebAppDeploymentTarget)(nil)).Elem()
}

func (i AzureWebAppDeploymentTargetMap) ToAzureWebAppDeploymentTargetMapOutput() AzureWebAppDeploymentTargetMapOutput {
	return i.ToAzureWebAppDeploymentTargetMapOutputWithContext(context.Background())
}

func (i AzureWebAppDeploymentTargetMap) ToAzureWebAppDeploymentTargetMapOutputWithContext(ctx context.Context) AzureWebAppDeploymentTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWebAppDeploymentTargetMapOutput)
}

type AzureWebAppDeploymentTargetOutput struct{ *pulumi.OutputState }

func (AzureWebAppDeploymentTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureWebAppDeploymentTarget)(nil)).Elem()
}

func (o AzureWebAppDeploymentTargetOutput) ToAzureWebAppDeploymentTargetOutput() AzureWebAppDeploymentTargetOutput {
	return o
}

func (o AzureWebAppDeploymentTargetOutput) ToAzureWebAppDeploymentTargetOutputWithContext(ctx context.Context) AzureWebAppDeploymentTargetOutput {
	return o
}

func (o AzureWebAppDeploymentTargetOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o AzureWebAppDeploymentTargetOutput) Endpoints() AzureWebAppDeploymentTargetEndpointArrayOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) AzureWebAppDeploymentTargetEndpointArrayOutput {
		return v.Endpoints
	}).(AzureWebAppDeploymentTargetEndpointArrayOutput)
}

// A list of environment IDs associated with this resource.
func (o AzureWebAppDeploymentTargetOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringArrayOutput { return v.Environments }).(pulumi.StringArrayOutput)
}

func (o AzureWebAppDeploymentTargetOutput) HasLatestCalamari() pulumi.BoolOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.BoolOutput { return v.HasLatestCalamari }).(pulumi.BoolOutput)
}

// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
func (o AzureWebAppDeploymentTargetOutput) HealthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.HealthStatus }).(pulumi.StringOutput)
}

func (o AzureWebAppDeploymentTargetOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.BoolOutput { return v.IsDisabled }).(pulumi.BoolOutput)
}

func (o AzureWebAppDeploymentTargetOutput) IsInProcess() pulumi.BoolOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.BoolOutput { return v.IsInProcess }).(pulumi.BoolOutput)
}

func (o AzureWebAppDeploymentTargetOutput) MachinePolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.MachinePolicyId }).(pulumi.StringOutput)
}

// The name of this resource.
func (o AzureWebAppDeploymentTargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AzureWebAppDeploymentTargetOutput) OperatingSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.OperatingSystem }).(pulumi.StringOutput)
}

func (o AzureWebAppDeploymentTargetOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o AzureWebAppDeploymentTargetOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o AzureWebAppDeploymentTargetOutput) ShellName() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.ShellName }).(pulumi.StringOutput)
}

func (o AzureWebAppDeploymentTargetOutput) ShellVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.ShellVersion }).(pulumi.StringOutput)
}

// The space ID associated with this resource.
func (o AzureWebAppDeploymentTargetOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
func (o AzureWebAppDeploymentTargetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A summary elaborating on the status of this resource.
func (o AzureWebAppDeploymentTargetOutput) StatusSummary() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.StatusSummary }).(pulumi.StringOutput)
}

// A list of tenant tags associated with this resource.
func (o AzureWebAppDeploymentTargetOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringArrayOutput { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
func (o AzureWebAppDeploymentTargetOutput) TenantedDeploymentParticipation() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.TenantedDeploymentParticipation }).(pulumi.StringOutput)
}

// A list of tenant IDs associated with this resource.
func (o AzureWebAppDeploymentTargetOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringArrayOutput { return v.Tenants }).(pulumi.StringArrayOutput)
}

func (o AzureWebAppDeploymentTargetOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o AzureWebAppDeploymentTargetOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

func (o AzureWebAppDeploymentTargetOutput) WebAppName() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringOutput { return v.WebAppName }).(pulumi.StringOutput)
}

func (o AzureWebAppDeploymentTargetOutput) WebAppSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureWebAppDeploymentTarget) pulumi.StringPtrOutput { return v.WebAppSlotName }).(pulumi.StringPtrOutput)
}

type AzureWebAppDeploymentTargetArrayOutput struct{ *pulumi.OutputState }

func (AzureWebAppDeploymentTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureWebAppDeploymentTarget)(nil)).Elem()
}

func (o AzureWebAppDeploymentTargetArrayOutput) ToAzureWebAppDeploymentTargetArrayOutput() AzureWebAppDeploymentTargetArrayOutput {
	return o
}

func (o AzureWebAppDeploymentTargetArrayOutput) ToAzureWebAppDeploymentTargetArrayOutputWithContext(ctx context.Context) AzureWebAppDeploymentTargetArrayOutput {
	return o
}

func (o AzureWebAppDeploymentTargetArrayOutput) Index(i pulumi.IntInput) AzureWebAppDeploymentTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AzureWebAppDeploymentTarget {
		return vs[0].([]*AzureWebAppDeploymentTarget)[vs[1].(int)]
	}).(AzureWebAppDeploymentTargetOutput)
}

type AzureWebAppDeploymentTargetMapOutput struct{ *pulumi.OutputState }

func (AzureWebAppDeploymentTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureWebAppDeploymentTarget)(nil)).Elem()
}

func (o AzureWebAppDeploymentTargetMapOutput) ToAzureWebAppDeploymentTargetMapOutput() AzureWebAppDeploymentTargetMapOutput {
	return o
}

func (o AzureWebAppDeploymentTargetMapOutput) ToAzureWebAppDeploymentTargetMapOutputWithContext(ctx context.Context) AzureWebAppDeploymentTargetMapOutput {
	return o
}

func (o AzureWebAppDeploymentTargetMapOutput) MapIndex(k pulumi.StringInput) AzureWebAppDeploymentTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AzureWebAppDeploymentTarget {
		return vs[0].(map[string]*AzureWebAppDeploymentTarget)[vs[1].(string)]
	}).(AzureWebAppDeploymentTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureWebAppDeploymentTargetInput)(nil)).Elem(), &AzureWebAppDeploymentTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureWebAppDeploymentTargetArrayInput)(nil)).Elem(), AzureWebAppDeploymentTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureWebAppDeploymentTargetMapInput)(nil)).Elem(), AzureWebAppDeploymentTargetMap{})
	pulumi.RegisterOutputType(AzureWebAppDeploymentTargetOutput{})
	pulumi.RegisterOutputType(AzureWebAppDeploymentTargetArrayOutput{})
	pulumi.RegisterOutputType(AzureWebAppDeploymentTargetMapOutput{})
}
