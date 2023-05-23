// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about existing Kubernetes cluster deployment targets.
func GetKubernetesClusterDeploymentTargets(ctx *pulumi.Context, args *GetKubernetesClusterDeploymentTargetsArgs, opts ...pulumi.InvokeOption) (*GetKubernetesClusterDeploymentTargetsResult, error) {
	var rv GetKubernetesClusterDeploymentTargetsResult
	err := ctx.Invoke("octopusdeploy:index/getKubernetesClusterDeploymentTargets:getKubernetesClusterDeploymentTargets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubernetesClusterDeploymentTargets.
type GetKubernetesClusterDeploymentTargetsArgs struct {
	// A filter to search by deployment ID.
	DeploymentId *string `pulumi:"deploymentId"`
	// A filter to search by a list of environment IDs.
	Environments []string `pulumi:"environments"`
	// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatuses []string `pulumi:"healthStatuses"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A filter to search by the disabled status of a resource.
	IsDisabled *bool `pulumi:"isDisabled"`
	// A list of Kubernetes cluster deployment targets that match the filter(s).
	KubernetesClusterDeploymentTargets []GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTarget `pulumi:"kubernetesClusterDeploymentTargets"`
	// A filter to search by name.
	Name *string `pulumi:"name"`
	// A filter to search by the partial match of a name.
	PartialName *string `pulumi:"partialName"`
	// A filter to search by a list of role IDs.
	Roles []string `pulumi:"roles"`
	// A list of shell names to match in the query and/or search
	ShellNames []string `pulumi:"shellNames"`
	// A filter to specify the number of items to skip in the response.
	Skip *int `pulumi:"skip"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
	// A filter to search by a list of tenant tags.
	TenantTags []string `pulumi:"tenantTags"`
	// A filter to search by a list of tenant IDs.
	Tenants []string `pulumi:"tenants"`
	// The thumbprint of the deployment target to match in the query and/or search
	Thumbprint *string `pulumi:"thumbprint"`
}

// A collection of values returned by getKubernetesClusterDeploymentTargets.
type GetKubernetesClusterDeploymentTargetsResult struct {
	// A filter to search by deployment ID.
	DeploymentId *string `pulumi:"deploymentId"`
	// A filter to search by a list of environment IDs.
	Environments []string `pulumi:"environments"`
	// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatuses []string `pulumi:"healthStatuses"`
	// An auto-generated identifier that includes the timestamp when this data source was last modified.
	Id string `pulumi:"id"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A filter to search by the disabled status of a resource.
	IsDisabled *bool `pulumi:"isDisabled"`
	// A list of Kubernetes cluster deployment targets that match the filter(s).
	KubernetesClusterDeploymentTargets []GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTarget `pulumi:"kubernetesClusterDeploymentTargets"`
	// A filter to search by name.
	Name *string `pulumi:"name"`
	// A filter to search by the partial match of a name.
	PartialName *string `pulumi:"partialName"`
	// A filter to search by a list of role IDs.
	Roles []string `pulumi:"roles"`
	// A list of shell names to match in the query and/or search
	ShellNames []string `pulumi:"shellNames"`
	// A filter to specify the number of items to skip in the response.
	Skip *int `pulumi:"skip"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
	// A filter to search by a list of tenant tags.
	TenantTags []string `pulumi:"tenantTags"`
	// A filter to search by a list of tenant IDs.
	Tenants []string `pulumi:"tenants"`
	// The thumbprint of the deployment target to match in the query and/or search
	Thumbprint *string `pulumi:"thumbprint"`
}

func GetKubernetesClusterDeploymentTargetsOutput(ctx *pulumi.Context, args GetKubernetesClusterDeploymentTargetsOutputArgs, opts ...pulumi.InvokeOption) GetKubernetesClusterDeploymentTargetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKubernetesClusterDeploymentTargetsResult, error) {
			args := v.(GetKubernetesClusterDeploymentTargetsArgs)
			r, err := GetKubernetesClusterDeploymentTargets(ctx, &args, opts...)
			var s GetKubernetesClusterDeploymentTargetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKubernetesClusterDeploymentTargetsResultOutput)
}

// A collection of arguments for invoking getKubernetesClusterDeploymentTargets.
type GetKubernetesClusterDeploymentTargetsOutputArgs struct {
	// A filter to search by deployment ID.
	DeploymentId pulumi.StringPtrInput `pulumi:"deploymentId"`
	// A filter to search by a list of environment IDs.
	Environments pulumi.StringArrayInput `pulumi:"environments"`
	// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatuses pulumi.StringArrayInput `pulumi:"healthStatuses"`
	// A filter to search by a list of IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A filter to search by the disabled status of a resource.
	IsDisabled pulumi.BoolPtrInput `pulumi:"isDisabled"`
	// A list of Kubernetes cluster deployment targets that match the filter(s).
	KubernetesClusterDeploymentTargets GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetArrayInput `pulumi:"kubernetesClusterDeploymentTargets"`
	// A filter to search by name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A filter to search by the partial match of a name.
	PartialName pulumi.StringPtrInput `pulumi:"partialName"`
	// A filter to search by a list of role IDs.
	Roles pulumi.StringArrayInput `pulumi:"roles"`
	// A list of shell names to match in the query and/or search
	ShellNames pulumi.StringArrayInput `pulumi:"shellNames"`
	// A filter to specify the number of items to skip in the response.
	Skip pulumi.IntPtrInput `pulumi:"skip"`
	// A filter to specify the number of items to take (or return) in the response.
	Take pulumi.IntPtrInput `pulumi:"take"`
	// A filter to search by a list of tenant tags.
	TenantTags pulumi.StringArrayInput `pulumi:"tenantTags"`
	// A filter to search by a list of tenant IDs.
	Tenants pulumi.StringArrayInput `pulumi:"tenants"`
	// The thumbprint of the deployment target to match in the query and/or search
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (GetKubernetesClusterDeploymentTargetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesClusterDeploymentTargetsArgs)(nil)).Elem()
}

// A collection of values returned by getKubernetesClusterDeploymentTargets.
type GetKubernetesClusterDeploymentTargetsResultOutput struct{ *pulumi.OutputState }

func (GetKubernetesClusterDeploymentTargetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesClusterDeploymentTargetsResult)(nil)).Elem()
}

func (o GetKubernetesClusterDeploymentTargetsResultOutput) ToGetKubernetesClusterDeploymentTargetsResultOutput() GetKubernetesClusterDeploymentTargetsResultOutput {
	return o
}

func (o GetKubernetesClusterDeploymentTargetsResultOutput) ToGetKubernetesClusterDeploymentTargetsResultOutputWithContext(ctx context.Context) GetKubernetesClusterDeploymentTargetsResultOutput {
	return o
}

// A filter to search by deployment ID.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) *string { return v.DeploymentId }).(pulumi.StringPtrOutput)
}

// A filter to search by a list of environment IDs.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) []string { return v.Environments }).(pulumi.StringArrayOutput)
}

// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) HealthStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) []string { return v.HealthStatuses }).(pulumi.StringArrayOutput)
}

// An auto-generated identifier that includes the timestamp when this data source was last modified.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A filter to search by a list of IDs.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A filter to search by the disabled status of a resource.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// A list of Kubernetes cluster deployment targets that match the filter(s).
func (o GetKubernetesClusterDeploymentTargetsResultOutput) KubernetesClusterDeploymentTargets() GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetArrayOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) []GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTarget {
		return v.KubernetesClusterDeploymentTargets
	}).(GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetArrayOutput)
}

// A filter to search by name.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A filter to search by the partial match of a name.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) PartialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) *string { return v.PartialName }).(pulumi.StringPtrOutput)
}

// A filter to search by a list of role IDs.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// A list of shell names to match in the query and/or search
func (o GetKubernetesClusterDeploymentTargetsResultOutput) ShellNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) []string { return v.ShellNames }).(pulumi.StringArrayOutput)
}

// A filter to specify the number of items to skip in the response.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) Skip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) *int { return v.Skip }).(pulumi.IntPtrOutput)
}

// A filter to specify the number of items to take (or return) in the response.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) Take() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) *int { return v.Take }).(pulumi.IntPtrOutput)
}

// A filter to search by a list of tenant tags.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) []string { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// A filter to search by a list of tenant IDs.
func (o GetKubernetesClusterDeploymentTargetsResultOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) []string { return v.Tenants }).(pulumi.StringArrayOutput)
}

// The thumbprint of the deployment target to match in the query and/or search
func (o GetKubernetesClusterDeploymentTargetsResultOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesClusterDeploymentTargetsResult) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubernetesClusterDeploymentTargetsResultOutput{})
}