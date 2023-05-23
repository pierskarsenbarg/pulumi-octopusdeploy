// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages listening tentacle deployment targets in Octopus Deploy.
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
//	$ pulumi import octopusdeploy:index/listeningTentacleDeploymentTarget:ListeningTentacleDeploymentTarget [options] octopusdeploy_listening_tentacle_deployment_target.<name> <machine-id>
//
// ```
type ListeningTentacleDeploymentTarget struct {
	pulumi.CustomResourceState

	CertificateSignatureAlgorithm pulumi.StringOutput `pulumi:"certificateSignatureAlgorithm"`
	// A list of environment IDs associated with this listening tentacle.
	Environments      pulumi.StringArrayOutput `pulumi:"environments"`
	HasLatestCalamari pulumi.BoolOutput        `pulumi:"hasLatestCalamari"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus pulumi.StringOutput `pulumi:"healthStatus"`
	// Represents the disabled status of this deployment target.
	IsDisabled pulumi.BoolOutput `pulumi:"isDisabled"`
	// Represents the in-process status of this deployment target.
	IsInProcess pulumi.BoolOutput `pulumi:"isInProcess"`
	// The machine policy ID that is associated with this deployment target.
	MachinePolicyId pulumi.StringOutput `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The operating system that is associated with this deployment target.
	OperatingSystem pulumi.StringOutput `pulumi:"operatingSystem"`
	// The proxy ID that is associated with this deployment target.
	ProxyId pulumi.StringOutput `pulumi:"proxyId"`
	// A list of role IDs that are associated with this deployment target.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The shell name associated with this deployment target.
	ShellName pulumi.StringOutput `pulumi:"shellName"`
	// The shell version associated with this deployment target.
	ShellVersion pulumi.StringOutput `pulumi:"shellVersion"`
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
	Tenants pulumi.StringArrayOutput `pulumi:"tenants"`
	// The tenant URL of this deployment target.
	TentacleUrl            pulumi.StringOutput                                               `pulumi:"tentacleUrl"`
	TentacleVersionDetails ListeningTentacleDeploymentTargetTentacleVersionDetailArrayOutput `pulumi:"tentacleVersionDetails"`
	// The thumbprint of this deployment target.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// The URI of this deployment target.
	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewListeningTentacleDeploymentTarget registers a new resource with the given unique name, arguments, and options.
func NewListeningTentacleDeploymentTarget(ctx *pulumi.Context,
	name string, args *ListeningTentacleDeploymentTargetArgs, opts ...pulumi.ResourceOption) (*ListeningTentacleDeploymentTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environments == nil {
		return nil, errors.New("invalid value for required argument 'Environments'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.TentacleUrl == nil {
		return nil, errors.New("invalid value for required argument 'TentacleUrl'")
	}
	if args.Thumbprint == nil {
		return nil, errors.New("invalid value for required argument 'Thumbprint'")
	}
	var resource ListeningTentacleDeploymentTarget
	err := ctx.RegisterResource("octopusdeploy:index/listeningTentacleDeploymentTarget:ListeningTentacleDeploymentTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListeningTentacleDeploymentTarget gets an existing ListeningTentacleDeploymentTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListeningTentacleDeploymentTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListeningTentacleDeploymentTargetState, opts ...pulumi.ResourceOption) (*ListeningTentacleDeploymentTarget, error) {
	var resource ListeningTentacleDeploymentTarget
	err := ctx.ReadResource("octopusdeploy:index/listeningTentacleDeploymentTarget:ListeningTentacleDeploymentTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ListeningTentacleDeploymentTarget resources.
type listeningTentacleDeploymentTargetState struct {
	CertificateSignatureAlgorithm *string `pulumi:"certificateSignatureAlgorithm"`
	// A list of environment IDs associated with this listening tentacle.
	Environments      []string `pulumi:"environments"`
	HasLatestCalamari *bool    `pulumi:"hasLatestCalamari"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus *string `pulumi:"healthStatus"`
	// Represents the disabled status of this deployment target.
	IsDisabled *bool `pulumi:"isDisabled"`
	// Represents the in-process status of this deployment target.
	IsInProcess *bool `pulumi:"isInProcess"`
	// The machine policy ID that is associated with this deployment target.
	MachinePolicyId *string `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The operating system that is associated with this deployment target.
	OperatingSystem *string `pulumi:"operatingSystem"`
	// The proxy ID that is associated with this deployment target.
	ProxyId *string `pulumi:"proxyId"`
	// A list of role IDs that are associated with this deployment target.
	Roles []string `pulumi:"roles"`
	// The shell name associated with this deployment target.
	ShellName *string `pulumi:"shellName"`
	// The shell version associated with this deployment target.
	ShellVersion *string `pulumi:"shellVersion"`
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
	Tenants []string `pulumi:"tenants"`
	// The tenant URL of this deployment target.
	TentacleUrl            *string                                                  `pulumi:"tentacleUrl"`
	TentacleVersionDetails []ListeningTentacleDeploymentTargetTentacleVersionDetail `pulumi:"tentacleVersionDetails"`
	// The thumbprint of this deployment target.
	Thumbprint *string `pulumi:"thumbprint"`
	// The URI of this deployment target.
	Uri *string `pulumi:"uri"`
}

type ListeningTentacleDeploymentTargetState struct {
	CertificateSignatureAlgorithm pulumi.StringPtrInput
	// A list of environment IDs associated with this listening tentacle.
	Environments      pulumi.StringArrayInput
	HasLatestCalamari pulumi.BoolPtrInput
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus pulumi.StringPtrInput
	// Represents the disabled status of this deployment target.
	IsDisabled pulumi.BoolPtrInput
	// Represents the in-process status of this deployment target.
	IsInProcess pulumi.BoolPtrInput
	// The machine policy ID that is associated with this deployment target.
	MachinePolicyId pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The operating system that is associated with this deployment target.
	OperatingSystem pulumi.StringPtrInput
	// The proxy ID that is associated with this deployment target.
	ProxyId pulumi.StringPtrInput
	// A list of role IDs that are associated with this deployment target.
	Roles pulumi.StringArrayInput
	// The shell name associated with this deployment target.
	ShellName pulumi.StringPtrInput
	// The shell version associated with this deployment target.
	ShellVersion pulumi.StringPtrInput
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
	Tenants pulumi.StringArrayInput
	// The tenant URL of this deployment target.
	TentacleUrl            pulumi.StringPtrInput
	TentacleVersionDetails ListeningTentacleDeploymentTargetTentacleVersionDetailArrayInput
	// The thumbprint of this deployment target.
	Thumbprint pulumi.StringPtrInput
	// The URI of this deployment target.
	Uri pulumi.StringPtrInput
}

func (ListeningTentacleDeploymentTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*listeningTentacleDeploymentTargetState)(nil)).Elem()
}

type listeningTentacleDeploymentTargetArgs struct {
	CertificateSignatureAlgorithm *string `pulumi:"certificateSignatureAlgorithm"`
	// A list of environment IDs associated with this listening tentacle.
	Environments []string `pulumi:"environments"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus *string `pulumi:"healthStatus"`
	// Represents the disabled status of this deployment target.
	IsDisabled *bool `pulumi:"isDisabled"`
	// Represents the in-process status of this deployment target.
	IsInProcess *bool `pulumi:"isInProcess"`
	// The machine policy ID that is associated with this deployment target.
	MachinePolicyId *string `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The operating system that is associated with this deployment target.
	OperatingSystem *string `pulumi:"operatingSystem"`
	// The proxy ID that is associated with this deployment target.
	ProxyId *string `pulumi:"proxyId"`
	// A list of role IDs that are associated with this deployment target.
	Roles []string `pulumi:"roles"`
	// The shell name associated with this deployment target.
	ShellName *string `pulumi:"shellName"`
	// The shell version associated with this deployment target.
	ShellVersion *string `pulumi:"shellVersion"`
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
	Tenants []string `pulumi:"tenants"`
	// The tenant URL of this deployment target.
	TentacleUrl            string                                                   `pulumi:"tentacleUrl"`
	TentacleVersionDetails []ListeningTentacleDeploymentTargetTentacleVersionDetail `pulumi:"tentacleVersionDetails"`
	// The thumbprint of this deployment target.
	Thumbprint string `pulumi:"thumbprint"`
	// The URI of this deployment target.
	Uri *string `pulumi:"uri"`
}

// The set of arguments for constructing a ListeningTentacleDeploymentTarget resource.
type ListeningTentacleDeploymentTargetArgs struct {
	CertificateSignatureAlgorithm pulumi.StringPtrInput
	// A list of environment IDs associated with this listening tentacle.
	Environments pulumi.StringArrayInput
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus pulumi.StringPtrInput
	// Represents the disabled status of this deployment target.
	IsDisabled pulumi.BoolPtrInput
	// Represents the in-process status of this deployment target.
	IsInProcess pulumi.BoolPtrInput
	// The machine policy ID that is associated with this deployment target.
	MachinePolicyId pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The operating system that is associated with this deployment target.
	OperatingSystem pulumi.StringPtrInput
	// The proxy ID that is associated with this deployment target.
	ProxyId pulumi.StringPtrInput
	// A list of role IDs that are associated with this deployment target.
	Roles pulumi.StringArrayInput
	// The shell name associated with this deployment target.
	ShellName pulumi.StringPtrInput
	// The shell version associated with this deployment target.
	ShellVersion pulumi.StringPtrInput
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
	Tenants pulumi.StringArrayInput
	// The tenant URL of this deployment target.
	TentacleUrl            pulumi.StringInput
	TentacleVersionDetails ListeningTentacleDeploymentTargetTentacleVersionDetailArrayInput
	// The thumbprint of this deployment target.
	Thumbprint pulumi.StringInput
	// The URI of this deployment target.
	Uri pulumi.StringPtrInput
}

func (ListeningTentacleDeploymentTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listeningTentacleDeploymentTargetArgs)(nil)).Elem()
}

type ListeningTentacleDeploymentTargetInput interface {
	pulumi.Input

	ToListeningTentacleDeploymentTargetOutput() ListeningTentacleDeploymentTargetOutput
	ToListeningTentacleDeploymentTargetOutputWithContext(ctx context.Context) ListeningTentacleDeploymentTargetOutput
}

func (*ListeningTentacleDeploymentTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**ListeningTentacleDeploymentTarget)(nil)).Elem()
}

func (i *ListeningTentacleDeploymentTarget) ToListeningTentacleDeploymentTargetOutput() ListeningTentacleDeploymentTargetOutput {
	return i.ToListeningTentacleDeploymentTargetOutputWithContext(context.Background())
}

func (i *ListeningTentacleDeploymentTarget) ToListeningTentacleDeploymentTargetOutputWithContext(ctx context.Context) ListeningTentacleDeploymentTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListeningTentacleDeploymentTargetOutput)
}

// ListeningTentacleDeploymentTargetArrayInput is an input type that accepts ListeningTentacleDeploymentTargetArray and ListeningTentacleDeploymentTargetArrayOutput values.
// You can construct a concrete instance of `ListeningTentacleDeploymentTargetArrayInput` via:
//
//	ListeningTentacleDeploymentTargetArray{ ListeningTentacleDeploymentTargetArgs{...} }
type ListeningTentacleDeploymentTargetArrayInput interface {
	pulumi.Input

	ToListeningTentacleDeploymentTargetArrayOutput() ListeningTentacleDeploymentTargetArrayOutput
	ToListeningTentacleDeploymentTargetArrayOutputWithContext(context.Context) ListeningTentacleDeploymentTargetArrayOutput
}

type ListeningTentacleDeploymentTargetArray []ListeningTentacleDeploymentTargetInput

func (ListeningTentacleDeploymentTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ListeningTentacleDeploymentTarget)(nil)).Elem()
}

func (i ListeningTentacleDeploymentTargetArray) ToListeningTentacleDeploymentTargetArrayOutput() ListeningTentacleDeploymentTargetArrayOutput {
	return i.ToListeningTentacleDeploymentTargetArrayOutputWithContext(context.Background())
}

func (i ListeningTentacleDeploymentTargetArray) ToListeningTentacleDeploymentTargetArrayOutputWithContext(ctx context.Context) ListeningTentacleDeploymentTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListeningTentacleDeploymentTargetArrayOutput)
}

// ListeningTentacleDeploymentTargetMapInput is an input type that accepts ListeningTentacleDeploymentTargetMap and ListeningTentacleDeploymentTargetMapOutput values.
// You can construct a concrete instance of `ListeningTentacleDeploymentTargetMapInput` via:
//
//	ListeningTentacleDeploymentTargetMap{ "key": ListeningTentacleDeploymentTargetArgs{...} }
type ListeningTentacleDeploymentTargetMapInput interface {
	pulumi.Input

	ToListeningTentacleDeploymentTargetMapOutput() ListeningTentacleDeploymentTargetMapOutput
	ToListeningTentacleDeploymentTargetMapOutputWithContext(context.Context) ListeningTentacleDeploymentTargetMapOutput
}

type ListeningTentacleDeploymentTargetMap map[string]ListeningTentacleDeploymentTargetInput

func (ListeningTentacleDeploymentTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ListeningTentacleDeploymentTarget)(nil)).Elem()
}

func (i ListeningTentacleDeploymentTargetMap) ToListeningTentacleDeploymentTargetMapOutput() ListeningTentacleDeploymentTargetMapOutput {
	return i.ToListeningTentacleDeploymentTargetMapOutputWithContext(context.Background())
}

func (i ListeningTentacleDeploymentTargetMap) ToListeningTentacleDeploymentTargetMapOutputWithContext(ctx context.Context) ListeningTentacleDeploymentTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListeningTentacleDeploymentTargetMapOutput)
}

type ListeningTentacleDeploymentTargetOutput struct{ *pulumi.OutputState }

func (ListeningTentacleDeploymentTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ListeningTentacleDeploymentTarget)(nil)).Elem()
}

func (o ListeningTentacleDeploymentTargetOutput) ToListeningTentacleDeploymentTargetOutput() ListeningTentacleDeploymentTargetOutput {
	return o
}

func (o ListeningTentacleDeploymentTargetOutput) ToListeningTentacleDeploymentTargetOutputWithContext(ctx context.Context) ListeningTentacleDeploymentTargetOutput {
	return o
}

func (o ListeningTentacleDeploymentTargetOutput) CertificateSignatureAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.CertificateSignatureAlgorithm }).(pulumi.StringOutput)
}

// A list of environment IDs associated with this listening tentacle.
func (o ListeningTentacleDeploymentTargetOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringArrayOutput { return v.Environments }).(pulumi.StringArrayOutput)
}

func (o ListeningTentacleDeploymentTargetOutput) HasLatestCalamari() pulumi.BoolOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.BoolOutput { return v.HasLatestCalamari }).(pulumi.BoolOutput)
}

// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
func (o ListeningTentacleDeploymentTargetOutput) HealthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.HealthStatus }).(pulumi.StringOutput)
}

// Represents the disabled status of this deployment target.
func (o ListeningTentacleDeploymentTargetOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.BoolOutput { return v.IsDisabled }).(pulumi.BoolOutput)
}

// Represents the in-process status of this deployment target.
func (o ListeningTentacleDeploymentTargetOutput) IsInProcess() pulumi.BoolOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.BoolOutput { return v.IsInProcess }).(pulumi.BoolOutput)
}

// The machine policy ID that is associated with this deployment target.
func (o ListeningTentacleDeploymentTargetOutput) MachinePolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.MachinePolicyId }).(pulumi.StringOutput)
}

// The name of this resource.
func (o ListeningTentacleDeploymentTargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The operating system that is associated with this deployment target.
func (o ListeningTentacleDeploymentTargetOutput) OperatingSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.OperatingSystem }).(pulumi.StringOutput)
}

// The proxy ID that is associated with this deployment target.
func (o ListeningTentacleDeploymentTargetOutput) ProxyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.ProxyId }).(pulumi.StringOutput)
}

// A list of role IDs that are associated with this deployment target.
func (o ListeningTentacleDeploymentTargetOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// The shell name associated with this deployment target.
func (o ListeningTentacleDeploymentTargetOutput) ShellName() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.ShellName }).(pulumi.StringOutput)
}

// The shell version associated with this deployment target.
func (o ListeningTentacleDeploymentTargetOutput) ShellVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.ShellVersion }).(pulumi.StringOutput)
}

// The space ID associated with this resource.
func (o ListeningTentacleDeploymentTargetOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
func (o ListeningTentacleDeploymentTargetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A summary elaborating on the status of this resource.
func (o ListeningTentacleDeploymentTargetOutput) StatusSummary() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.StatusSummary }).(pulumi.StringOutput)
}

// A list of tenant tags associated with this resource.
func (o ListeningTentacleDeploymentTargetOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringArrayOutput { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
func (o ListeningTentacleDeploymentTargetOutput) TenantedDeploymentParticipation() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput {
		return v.TenantedDeploymentParticipation
	}).(pulumi.StringOutput)
}

// A list of tenant IDs associated with this resource.
func (o ListeningTentacleDeploymentTargetOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringArrayOutput { return v.Tenants }).(pulumi.StringArrayOutput)
}

// The tenant URL of this deployment target.
func (o ListeningTentacleDeploymentTargetOutput) TentacleUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.TentacleUrl }).(pulumi.StringOutput)
}

func (o ListeningTentacleDeploymentTargetOutput) TentacleVersionDetails() ListeningTentacleDeploymentTargetTentacleVersionDetailArrayOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) ListeningTentacleDeploymentTargetTentacleVersionDetailArrayOutput {
		return v.TentacleVersionDetails
	}).(ListeningTentacleDeploymentTargetTentacleVersionDetailArrayOutput)
}

// The thumbprint of this deployment target.
func (o ListeningTentacleDeploymentTargetOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

// The URI of this deployment target.
func (o ListeningTentacleDeploymentTargetOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *ListeningTentacleDeploymentTarget) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

type ListeningTentacleDeploymentTargetArrayOutput struct{ *pulumi.OutputState }

func (ListeningTentacleDeploymentTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ListeningTentacleDeploymentTarget)(nil)).Elem()
}

func (o ListeningTentacleDeploymentTargetArrayOutput) ToListeningTentacleDeploymentTargetArrayOutput() ListeningTentacleDeploymentTargetArrayOutput {
	return o
}

func (o ListeningTentacleDeploymentTargetArrayOutput) ToListeningTentacleDeploymentTargetArrayOutputWithContext(ctx context.Context) ListeningTentacleDeploymentTargetArrayOutput {
	return o
}

func (o ListeningTentacleDeploymentTargetArrayOutput) Index(i pulumi.IntInput) ListeningTentacleDeploymentTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ListeningTentacleDeploymentTarget {
		return vs[0].([]*ListeningTentacleDeploymentTarget)[vs[1].(int)]
	}).(ListeningTentacleDeploymentTargetOutput)
}

type ListeningTentacleDeploymentTargetMapOutput struct{ *pulumi.OutputState }

func (ListeningTentacleDeploymentTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ListeningTentacleDeploymentTarget)(nil)).Elem()
}

func (o ListeningTentacleDeploymentTargetMapOutput) ToListeningTentacleDeploymentTargetMapOutput() ListeningTentacleDeploymentTargetMapOutput {
	return o
}

func (o ListeningTentacleDeploymentTargetMapOutput) ToListeningTentacleDeploymentTargetMapOutputWithContext(ctx context.Context) ListeningTentacleDeploymentTargetMapOutput {
	return o
}

func (o ListeningTentacleDeploymentTargetMapOutput) MapIndex(k pulumi.StringInput) ListeningTentacleDeploymentTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ListeningTentacleDeploymentTarget {
		return vs[0].(map[string]*ListeningTentacleDeploymentTarget)[vs[1].(string)]
	}).(ListeningTentacleDeploymentTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListeningTentacleDeploymentTargetInput)(nil)).Elem(), &ListeningTentacleDeploymentTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListeningTentacleDeploymentTargetArrayInput)(nil)).Elem(), ListeningTentacleDeploymentTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListeningTentacleDeploymentTargetMapInput)(nil)).Elem(), ListeningTentacleDeploymentTargetMap{})
	pulumi.RegisterOutputType(ListeningTentacleDeploymentTargetOutput{})
	pulumi.RegisterOutputType(ListeningTentacleDeploymentTargetArrayOutput{})
	pulumi.RegisterOutputType(ListeningTentacleDeploymentTargetMapOutput{})
}