// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages Kubernets cluster deployment targets in Octopus Deploy.
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
//			_, err := octopusdeploy.NewKubernetesClusterDeploymentTarget(ctx, "k8s-target", &octopusdeploy.KubernetesClusterDeploymentTargetArgs{
//				AwsAccountAuthentication: &octopusdeploy.KubernetesClusterDeploymentTargetAwsAccountAuthenticationArgs{
//					AccountId:   pulumi.String("Accounts-123"),
//					ClusterName: pulumi.String("cluster-name"),
//				},
//				ClusterUrl: pulumi.String("https://example.com:1234/"),
//				Environments: pulumi.StringArray{
//					pulumi.String("Environments-123"),
//					pulumi.String("Environment-321"),
//				},
//				Roles: pulumi.StringArray{
//					pulumi.String("Development Team"),
//					pulumi.String("System Administrators"),
//				},
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
//	$ pulumi import octopusdeploy:index/kubernetesClusterDeploymentTarget:KubernetesClusterDeploymentTarget [options] octopusdeploy_kubernetes_cluster_deployment_target.<name> <machine-id>
//
// ```
type KubernetesClusterDeploymentTarget struct {
	pulumi.CustomResourceState

	Authentication                      KubernetesClusterDeploymentTargetAuthenticationOutput                      `pulumi:"authentication"`
	AwsAccountAuthentication            KubernetesClusterDeploymentTargetAwsAccountAuthenticationOutput            `pulumi:"awsAccountAuthentication"`
	AzureServicePrincipalAuthentication KubernetesClusterDeploymentTargetAzureServicePrincipalAuthenticationOutput `pulumi:"azureServicePrincipalAuthentication"`
	CertificateAuthentication           KubernetesClusterDeploymentTargetCertificateAuthenticationOutput           `pulumi:"certificateAuthentication"`
	ClusterCertificate                  pulumi.StringPtrOutput                                                     `pulumi:"clusterCertificate"`
	ClusterUrl                          pulumi.StringOutput                                                        `pulumi:"clusterUrl"`
	Containers                          KubernetesClusterDeploymentTargetContainerArrayOutput                      `pulumi:"containers"`
	DefaultWorkerPoolId                 pulumi.StringPtrOutput                                                     `pulumi:"defaultWorkerPoolId"`
	Endpoints                           KubernetesClusterDeploymentTargetEndpointArrayOutput                       `pulumi:"endpoints"`
	// A list of environment IDs associated with this resource.
	Environments             pulumi.StringArrayOutput                                        `pulumi:"environments"`
	GcpAccountAuthentication KubernetesClusterDeploymentTargetGcpAccountAuthenticationOutput `pulumi:"gcpAccountAuthentication"`
	HasLatestCalamari        pulumi.BoolOutput                                               `pulumi:"hasLatestCalamari"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    pulumi.StringOutput `pulumi:"healthStatus"`
	IsDisabled      pulumi.BoolOutput   `pulumi:"isDisabled"`
	IsInProcess     pulumi.BoolOutput   `pulumi:"isInProcess"`
	MachinePolicyId pulumi.StringOutput `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name                pulumi.StringOutput      `pulumi:"name"`
	Namespace           pulumi.StringPtrOutput   `pulumi:"namespace"`
	OperatingSystem     pulumi.StringOutput      `pulumi:"operatingSystem"`
	ProxyId             pulumi.StringPtrOutput   `pulumi:"proxyId"`
	Roles               pulumi.StringArrayOutput `pulumi:"roles"`
	RunningInContainer  pulumi.BoolPtrOutput     `pulumi:"runningInContainer"`
	ShellName           pulumi.StringOutput      `pulumi:"shellName"`
	ShellVersion        pulumi.StringOutput      `pulumi:"shellVersion"`
	SkipTlsVerification pulumi.BoolPtrOutput     `pulumi:"skipTlsVerification"`
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
	Tenants    pulumi.StringArrayOutput `pulumi:"tenants"`
	Thumbprint pulumi.StringOutput      `pulumi:"thumbprint"`
	Uri        pulumi.StringOutput      `pulumi:"uri"`
}

// NewKubernetesClusterDeploymentTarget registers a new resource with the given unique name, arguments, and options.
func NewKubernetesClusterDeploymentTarget(ctx *pulumi.Context,
	name string, args *KubernetesClusterDeploymentTargetArgs, opts ...pulumi.ResourceOption) (*KubernetesClusterDeploymentTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterUrl == nil {
		return nil, errors.New("invalid value for required argument 'ClusterUrl'")
	}
	if args.Environments == nil {
		return nil, errors.New("invalid value for required argument 'Environments'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	var resource KubernetesClusterDeploymentTarget
	err := ctx.RegisterResource("octopusdeploy:index/kubernetesClusterDeploymentTarget:KubernetesClusterDeploymentTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesClusterDeploymentTarget gets an existing KubernetesClusterDeploymentTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesClusterDeploymentTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesClusterDeploymentTargetState, opts ...pulumi.ResourceOption) (*KubernetesClusterDeploymentTarget, error) {
	var resource KubernetesClusterDeploymentTarget
	err := ctx.ReadResource("octopusdeploy:index/kubernetesClusterDeploymentTarget:KubernetesClusterDeploymentTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesClusterDeploymentTarget resources.
type kubernetesClusterDeploymentTargetState struct {
	Authentication                      *KubernetesClusterDeploymentTargetAuthentication                      `pulumi:"authentication"`
	AwsAccountAuthentication            *KubernetesClusterDeploymentTargetAwsAccountAuthentication            `pulumi:"awsAccountAuthentication"`
	AzureServicePrincipalAuthentication *KubernetesClusterDeploymentTargetAzureServicePrincipalAuthentication `pulumi:"azureServicePrincipalAuthentication"`
	CertificateAuthentication           *KubernetesClusterDeploymentTargetCertificateAuthentication           `pulumi:"certificateAuthentication"`
	ClusterCertificate                  *string                                                               `pulumi:"clusterCertificate"`
	ClusterUrl                          *string                                                               `pulumi:"clusterUrl"`
	Containers                          []KubernetesClusterDeploymentTargetContainer                          `pulumi:"containers"`
	DefaultWorkerPoolId                 *string                                                               `pulumi:"defaultWorkerPoolId"`
	Endpoints                           []KubernetesClusterDeploymentTargetEndpoint                           `pulumi:"endpoints"`
	// A list of environment IDs associated with this resource.
	Environments             []string                                                   `pulumi:"environments"`
	GcpAccountAuthentication *KubernetesClusterDeploymentTargetGcpAccountAuthentication `pulumi:"gcpAccountAuthentication"`
	HasLatestCalamari        *bool                                                      `pulumi:"hasLatestCalamari"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    *string `pulumi:"healthStatus"`
	IsDisabled      *bool   `pulumi:"isDisabled"`
	IsInProcess     *bool   `pulumi:"isInProcess"`
	MachinePolicyId *string `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name                *string  `pulumi:"name"`
	Namespace           *string  `pulumi:"namespace"`
	OperatingSystem     *string  `pulumi:"operatingSystem"`
	ProxyId             *string  `pulumi:"proxyId"`
	Roles               []string `pulumi:"roles"`
	RunningInContainer  *bool    `pulumi:"runningInContainer"`
	ShellName           *string  `pulumi:"shellName"`
	ShellVersion        *string  `pulumi:"shellVersion"`
	SkipTlsVerification *bool    `pulumi:"skipTlsVerification"`
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
	Tenants    []string `pulumi:"tenants"`
	Thumbprint *string  `pulumi:"thumbprint"`
	Uri        *string  `pulumi:"uri"`
}

type KubernetesClusterDeploymentTargetState struct {
	Authentication                      KubernetesClusterDeploymentTargetAuthenticationPtrInput
	AwsAccountAuthentication            KubernetesClusterDeploymentTargetAwsAccountAuthenticationPtrInput
	AzureServicePrincipalAuthentication KubernetesClusterDeploymentTargetAzureServicePrincipalAuthenticationPtrInput
	CertificateAuthentication           KubernetesClusterDeploymentTargetCertificateAuthenticationPtrInput
	ClusterCertificate                  pulumi.StringPtrInput
	ClusterUrl                          pulumi.StringPtrInput
	Containers                          KubernetesClusterDeploymentTargetContainerArrayInput
	DefaultWorkerPoolId                 pulumi.StringPtrInput
	Endpoints                           KubernetesClusterDeploymentTargetEndpointArrayInput
	// A list of environment IDs associated with this resource.
	Environments             pulumi.StringArrayInput
	GcpAccountAuthentication KubernetesClusterDeploymentTargetGcpAccountAuthenticationPtrInput
	HasLatestCalamari        pulumi.BoolPtrInput
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    pulumi.StringPtrInput
	IsDisabled      pulumi.BoolPtrInput
	IsInProcess     pulumi.BoolPtrInput
	MachinePolicyId pulumi.StringPtrInput
	// The name of this resource.
	Name                pulumi.StringPtrInput
	Namespace           pulumi.StringPtrInput
	OperatingSystem     pulumi.StringPtrInput
	ProxyId             pulumi.StringPtrInput
	Roles               pulumi.StringArrayInput
	RunningInContainer  pulumi.BoolPtrInput
	ShellName           pulumi.StringPtrInput
	ShellVersion        pulumi.StringPtrInput
	SkipTlsVerification pulumi.BoolPtrInput
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
	Tenants    pulumi.StringArrayInput
	Thumbprint pulumi.StringPtrInput
	Uri        pulumi.StringPtrInput
}

func (KubernetesClusterDeploymentTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterDeploymentTargetState)(nil)).Elem()
}

type kubernetesClusterDeploymentTargetArgs struct {
	Authentication                      *KubernetesClusterDeploymentTargetAuthentication                      `pulumi:"authentication"`
	AwsAccountAuthentication            *KubernetesClusterDeploymentTargetAwsAccountAuthentication            `pulumi:"awsAccountAuthentication"`
	AzureServicePrincipalAuthentication *KubernetesClusterDeploymentTargetAzureServicePrincipalAuthentication `pulumi:"azureServicePrincipalAuthentication"`
	CertificateAuthentication           *KubernetesClusterDeploymentTargetCertificateAuthentication           `pulumi:"certificateAuthentication"`
	ClusterCertificate                  *string                                                               `pulumi:"clusterCertificate"`
	ClusterUrl                          string                                                                `pulumi:"clusterUrl"`
	Containers                          []KubernetesClusterDeploymentTargetContainer                          `pulumi:"containers"`
	DefaultWorkerPoolId                 *string                                                               `pulumi:"defaultWorkerPoolId"`
	Endpoints                           []KubernetesClusterDeploymentTargetEndpoint                           `pulumi:"endpoints"`
	// A list of environment IDs associated with this resource.
	Environments             []string                                                   `pulumi:"environments"`
	GcpAccountAuthentication *KubernetesClusterDeploymentTargetGcpAccountAuthentication `pulumi:"gcpAccountAuthentication"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    *string `pulumi:"healthStatus"`
	IsDisabled      *bool   `pulumi:"isDisabled"`
	MachinePolicyId *string `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name                *string  `pulumi:"name"`
	Namespace           *string  `pulumi:"namespace"`
	OperatingSystem     *string  `pulumi:"operatingSystem"`
	ProxyId             *string  `pulumi:"proxyId"`
	Roles               []string `pulumi:"roles"`
	RunningInContainer  *bool    `pulumi:"runningInContainer"`
	ShellName           *string  `pulumi:"shellName"`
	ShellVersion        *string  `pulumi:"shellVersion"`
	SkipTlsVerification *bool    `pulumi:"skipTlsVerification"`
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
	Tenants    []string `pulumi:"tenants"`
	Thumbprint *string  `pulumi:"thumbprint"`
	Uri        *string  `pulumi:"uri"`
}

// The set of arguments for constructing a KubernetesClusterDeploymentTarget resource.
type KubernetesClusterDeploymentTargetArgs struct {
	Authentication                      KubernetesClusterDeploymentTargetAuthenticationPtrInput
	AwsAccountAuthentication            KubernetesClusterDeploymentTargetAwsAccountAuthenticationPtrInput
	AzureServicePrincipalAuthentication KubernetesClusterDeploymentTargetAzureServicePrincipalAuthenticationPtrInput
	CertificateAuthentication           KubernetesClusterDeploymentTargetCertificateAuthenticationPtrInput
	ClusterCertificate                  pulumi.StringPtrInput
	ClusterUrl                          pulumi.StringInput
	Containers                          KubernetesClusterDeploymentTargetContainerArrayInput
	DefaultWorkerPoolId                 pulumi.StringPtrInput
	Endpoints                           KubernetesClusterDeploymentTargetEndpointArrayInput
	// A list of environment IDs associated with this resource.
	Environments             pulumi.StringArrayInput
	GcpAccountAuthentication KubernetesClusterDeploymentTargetGcpAccountAuthenticationPtrInput
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    pulumi.StringPtrInput
	IsDisabled      pulumi.BoolPtrInput
	MachinePolicyId pulumi.StringPtrInput
	// The name of this resource.
	Name                pulumi.StringPtrInput
	Namespace           pulumi.StringPtrInput
	OperatingSystem     pulumi.StringPtrInput
	ProxyId             pulumi.StringPtrInput
	Roles               pulumi.StringArrayInput
	RunningInContainer  pulumi.BoolPtrInput
	ShellName           pulumi.StringPtrInput
	ShellVersion        pulumi.StringPtrInput
	SkipTlsVerification pulumi.BoolPtrInput
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
	Tenants    pulumi.StringArrayInput
	Thumbprint pulumi.StringPtrInput
	Uri        pulumi.StringPtrInput
}

func (KubernetesClusterDeploymentTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterDeploymentTargetArgs)(nil)).Elem()
}

type KubernetesClusterDeploymentTargetInput interface {
	pulumi.Input

	ToKubernetesClusterDeploymentTargetOutput() KubernetesClusterDeploymentTargetOutput
	ToKubernetesClusterDeploymentTargetOutputWithContext(ctx context.Context) KubernetesClusterDeploymentTargetOutput
}

func (*KubernetesClusterDeploymentTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesClusterDeploymentTarget)(nil)).Elem()
}

func (i *KubernetesClusterDeploymentTarget) ToKubernetesClusterDeploymentTargetOutput() KubernetesClusterDeploymentTargetOutput {
	return i.ToKubernetesClusterDeploymentTargetOutputWithContext(context.Background())
}

func (i *KubernetesClusterDeploymentTarget) ToKubernetesClusterDeploymentTargetOutputWithContext(ctx context.Context) KubernetesClusterDeploymentTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterDeploymentTargetOutput)
}

// KubernetesClusterDeploymentTargetArrayInput is an input type that accepts KubernetesClusterDeploymentTargetArray and KubernetesClusterDeploymentTargetArrayOutput values.
// You can construct a concrete instance of `KubernetesClusterDeploymentTargetArrayInput` via:
//
//	KubernetesClusterDeploymentTargetArray{ KubernetesClusterDeploymentTargetArgs{...} }
type KubernetesClusterDeploymentTargetArrayInput interface {
	pulumi.Input

	ToKubernetesClusterDeploymentTargetArrayOutput() KubernetesClusterDeploymentTargetArrayOutput
	ToKubernetesClusterDeploymentTargetArrayOutputWithContext(context.Context) KubernetesClusterDeploymentTargetArrayOutput
}

type KubernetesClusterDeploymentTargetArray []KubernetesClusterDeploymentTargetInput

func (KubernetesClusterDeploymentTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubernetesClusterDeploymentTarget)(nil)).Elem()
}

func (i KubernetesClusterDeploymentTargetArray) ToKubernetesClusterDeploymentTargetArrayOutput() KubernetesClusterDeploymentTargetArrayOutput {
	return i.ToKubernetesClusterDeploymentTargetArrayOutputWithContext(context.Background())
}

func (i KubernetesClusterDeploymentTargetArray) ToKubernetesClusterDeploymentTargetArrayOutputWithContext(ctx context.Context) KubernetesClusterDeploymentTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterDeploymentTargetArrayOutput)
}

// KubernetesClusterDeploymentTargetMapInput is an input type that accepts KubernetesClusterDeploymentTargetMap and KubernetesClusterDeploymentTargetMapOutput values.
// You can construct a concrete instance of `KubernetesClusterDeploymentTargetMapInput` via:
//
//	KubernetesClusterDeploymentTargetMap{ "key": KubernetesClusterDeploymentTargetArgs{...} }
type KubernetesClusterDeploymentTargetMapInput interface {
	pulumi.Input

	ToKubernetesClusterDeploymentTargetMapOutput() KubernetesClusterDeploymentTargetMapOutput
	ToKubernetesClusterDeploymentTargetMapOutputWithContext(context.Context) KubernetesClusterDeploymentTargetMapOutput
}

type KubernetesClusterDeploymentTargetMap map[string]KubernetesClusterDeploymentTargetInput

func (KubernetesClusterDeploymentTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubernetesClusterDeploymentTarget)(nil)).Elem()
}

func (i KubernetesClusterDeploymentTargetMap) ToKubernetesClusterDeploymentTargetMapOutput() KubernetesClusterDeploymentTargetMapOutput {
	return i.ToKubernetesClusterDeploymentTargetMapOutputWithContext(context.Background())
}

func (i KubernetesClusterDeploymentTargetMap) ToKubernetesClusterDeploymentTargetMapOutputWithContext(ctx context.Context) KubernetesClusterDeploymentTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterDeploymentTargetMapOutput)
}

type KubernetesClusterDeploymentTargetOutput struct{ *pulumi.OutputState }

func (KubernetesClusterDeploymentTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesClusterDeploymentTarget)(nil)).Elem()
}

func (o KubernetesClusterDeploymentTargetOutput) ToKubernetesClusterDeploymentTargetOutput() KubernetesClusterDeploymentTargetOutput {
	return o
}

func (o KubernetesClusterDeploymentTargetOutput) ToKubernetesClusterDeploymentTargetOutputWithContext(ctx context.Context) KubernetesClusterDeploymentTargetOutput {
	return o
}

func (o KubernetesClusterDeploymentTargetOutput) Authentication() KubernetesClusterDeploymentTargetAuthenticationOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) KubernetesClusterDeploymentTargetAuthenticationOutput {
		return v.Authentication
	}).(KubernetesClusterDeploymentTargetAuthenticationOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) AwsAccountAuthentication() KubernetesClusterDeploymentTargetAwsAccountAuthenticationOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) KubernetesClusterDeploymentTargetAwsAccountAuthenticationOutput {
		return v.AwsAccountAuthentication
	}).(KubernetesClusterDeploymentTargetAwsAccountAuthenticationOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) AzureServicePrincipalAuthentication() KubernetesClusterDeploymentTargetAzureServicePrincipalAuthenticationOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) KubernetesClusterDeploymentTargetAzureServicePrincipalAuthenticationOutput {
		return v.AzureServicePrincipalAuthentication
	}).(KubernetesClusterDeploymentTargetAzureServicePrincipalAuthenticationOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) CertificateAuthentication() KubernetesClusterDeploymentTargetCertificateAuthenticationOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) KubernetesClusterDeploymentTargetCertificateAuthenticationOutput {
		return v.CertificateAuthentication
	}).(KubernetesClusterDeploymentTargetCertificateAuthenticationOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) ClusterCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringPtrOutput { return v.ClusterCertificate }).(pulumi.StringPtrOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) ClusterUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput { return v.ClusterUrl }).(pulumi.StringOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) Containers() KubernetesClusterDeploymentTargetContainerArrayOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) KubernetesClusterDeploymentTargetContainerArrayOutput {
		return v.Containers
	}).(KubernetesClusterDeploymentTargetContainerArrayOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) DefaultWorkerPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringPtrOutput { return v.DefaultWorkerPoolId }).(pulumi.StringPtrOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) Endpoints() KubernetesClusterDeploymentTargetEndpointArrayOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) KubernetesClusterDeploymentTargetEndpointArrayOutput {
		return v.Endpoints
	}).(KubernetesClusterDeploymentTargetEndpointArrayOutput)
}

// A list of environment IDs associated with this resource.
func (o KubernetesClusterDeploymentTargetOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringArrayOutput { return v.Environments }).(pulumi.StringArrayOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) GcpAccountAuthentication() KubernetesClusterDeploymentTargetGcpAccountAuthenticationOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) KubernetesClusterDeploymentTargetGcpAccountAuthenticationOutput {
		return v.GcpAccountAuthentication
	}).(KubernetesClusterDeploymentTargetGcpAccountAuthenticationOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) HasLatestCalamari() pulumi.BoolOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.BoolOutput { return v.HasLatestCalamari }).(pulumi.BoolOutput)
}

// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
func (o KubernetesClusterDeploymentTargetOutput) HealthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput { return v.HealthStatus }).(pulumi.StringOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.BoolOutput { return v.IsDisabled }).(pulumi.BoolOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) IsInProcess() pulumi.BoolOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.BoolOutput { return v.IsInProcess }).(pulumi.BoolOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) MachinePolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput { return v.MachinePolicyId }).(pulumi.StringOutput)
}

// The name of this resource.
func (o KubernetesClusterDeploymentTargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) OperatingSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput { return v.OperatingSystem }).(pulumi.StringOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) ProxyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringPtrOutput { return v.ProxyId }).(pulumi.StringPtrOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) RunningInContainer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.BoolPtrOutput { return v.RunningInContainer }).(pulumi.BoolPtrOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) ShellName() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput { return v.ShellName }).(pulumi.StringOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) ShellVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput { return v.ShellVersion }).(pulumi.StringOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) SkipTlsVerification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.BoolPtrOutput { return v.SkipTlsVerification }).(pulumi.BoolPtrOutput)
}

// The space ID associated with this resource.
func (o KubernetesClusterDeploymentTargetOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
func (o KubernetesClusterDeploymentTargetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A summary elaborating on the status of this resource.
func (o KubernetesClusterDeploymentTargetOutput) StatusSummary() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput { return v.StatusSummary }).(pulumi.StringOutput)
}

// A list of tenant tags associated with this resource.
func (o KubernetesClusterDeploymentTargetOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringArrayOutput { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
func (o KubernetesClusterDeploymentTargetOutput) TenantedDeploymentParticipation() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput {
		return v.TenantedDeploymentParticipation
	}).(pulumi.StringOutput)
}

// A list of tenant IDs associated with this resource.
func (o KubernetesClusterDeploymentTargetOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringArrayOutput { return v.Tenants }).(pulumi.StringArrayOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o KubernetesClusterDeploymentTargetOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesClusterDeploymentTarget) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

type KubernetesClusterDeploymentTargetArrayOutput struct{ *pulumi.OutputState }

func (KubernetesClusterDeploymentTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubernetesClusterDeploymentTarget)(nil)).Elem()
}

func (o KubernetesClusterDeploymentTargetArrayOutput) ToKubernetesClusterDeploymentTargetArrayOutput() KubernetesClusterDeploymentTargetArrayOutput {
	return o
}

func (o KubernetesClusterDeploymentTargetArrayOutput) ToKubernetesClusterDeploymentTargetArrayOutputWithContext(ctx context.Context) KubernetesClusterDeploymentTargetArrayOutput {
	return o
}

func (o KubernetesClusterDeploymentTargetArrayOutput) Index(i pulumi.IntInput) KubernetesClusterDeploymentTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubernetesClusterDeploymentTarget {
		return vs[0].([]*KubernetesClusterDeploymentTarget)[vs[1].(int)]
	}).(KubernetesClusterDeploymentTargetOutput)
}

type KubernetesClusterDeploymentTargetMapOutput struct{ *pulumi.OutputState }

func (KubernetesClusterDeploymentTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubernetesClusterDeploymentTarget)(nil)).Elem()
}

func (o KubernetesClusterDeploymentTargetMapOutput) ToKubernetesClusterDeploymentTargetMapOutput() KubernetesClusterDeploymentTargetMapOutput {
	return o
}

func (o KubernetesClusterDeploymentTargetMapOutput) ToKubernetesClusterDeploymentTargetMapOutputWithContext(ctx context.Context) KubernetesClusterDeploymentTargetMapOutput {
	return o
}

func (o KubernetesClusterDeploymentTargetMapOutput) MapIndex(k pulumi.StringInput) KubernetesClusterDeploymentTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubernetesClusterDeploymentTarget {
		return vs[0].(map[string]*KubernetesClusterDeploymentTarget)[vs[1].(string)]
	}).(KubernetesClusterDeploymentTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesClusterDeploymentTargetInput)(nil)).Elem(), &KubernetesClusterDeploymentTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesClusterDeploymentTargetArrayInput)(nil)).Elem(), KubernetesClusterDeploymentTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesClusterDeploymentTargetMapInput)(nil)).Elem(), KubernetesClusterDeploymentTargetMap{})
	pulumi.RegisterOutputType(KubernetesClusterDeploymentTargetOutput{})
	pulumi.RegisterOutputType(KubernetesClusterDeploymentTargetArrayOutput{})
	pulumi.RegisterOutputType(KubernetesClusterDeploymentTargetMapOutput{})
}
