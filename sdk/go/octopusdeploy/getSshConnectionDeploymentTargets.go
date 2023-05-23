// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about existing SSH connection deployment targets.
func GetSshConnectionDeploymentTargets(ctx *pulumi.Context, args *GetSshConnectionDeploymentTargetsArgs, opts ...pulumi.InvokeOption) (*GetSshConnectionDeploymentTargetsResult, error) {
	var rv GetSshConnectionDeploymentTargetsResult
	err := ctx.Invoke("octopusdeploy:index/getSshConnectionDeploymentTargets:getSshConnectionDeploymentTargets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSshConnectionDeploymentTargets.
type GetSshConnectionDeploymentTargetsArgs struct {
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
	// A list of SSH connection deployment targets that match the filter(s).
	SshConnectionDeploymentTargets []GetSshConnectionDeploymentTargetsSshConnectionDeploymentTarget `pulumi:"sshConnectionDeploymentTargets"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
	// A filter to search by a list of tenant tags.
	TenantTags []string `pulumi:"tenantTags"`
	// A filter to search by a list of tenant IDs.
	Tenants []string `pulumi:"tenants"`
	// The thumbprint of the deployment target to match in the query and/or search
	Thumbprint *string `pulumi:"thumbprint"`
}

// A collection of values returned by getSshConnectionDeploymentTargets.
type GetSshConnectionDeploymentTargetsResult struct {
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
	// A list of SSH connection deployment targets that match the filter(s).
	SshConnectionDeploymentTargets []GetSshConnectionDeploymentTargetsSshConnectionDeploymentTarget `pulumi:"sshConnectionDeploymentTargets"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
	// A filter to search by a list of tenant tags.
	TenantTags []string `pulumi:"tenantTags"`
	// A filter to search by a list of tenant IDs.
	Tenants []string `pulumi:"tenants"`
	// The thumbprint of the deployment target to match in the query and/or search
	Thumbprint *string `pulumi:"thumbprint"`
}

func GetSshConnectionDeploymentTargetsOutput(ctx *pulumi.Context, args GetSshConnectionDeploymentTargetsOutputArgs, opts ...pulumi.InvokeOption) GetSshConnectionDeploymentTargetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSshConnectionDeploymentTargetsResult, error) {
			args := v.(GetSshConnectionDeploymentTargetsArgs)
			r, err := GetSshConnectionDeploymentTargets(ctx, &args, opts...)
			var s GetSshConnectionDeploymentTargetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSshConnectionDeploymentTargetsResultOutput)
}

// A collection of arguments for invoking getSshConnectionDeploymentTargets.
type GetSshConnectionDeploymentTargetsOutputArgs struct {
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
	// A list of SSH connection deployment targets that match the filter(s).
	SshConnectionDeploymentTargets GetSshConnectionDeploymentTargetsSshConnectionDeploymentTargetArrayInput `pulumi:"sshConnectionDeploymentTargets"`
	// A filter to specify the number of items to take (or return) in the response.
	Take pulumi.IntPtrInput `pulumi:"take"`
	// A filter to search by a list of tenant tags.
	TenantTags pulumi.StringArrayInput `pulumi:"tenantTags"`
	// A filter to search by a list of tenant IDs.
	Tenants pulumi.StringArrayInput `pulumi:"tenants"`
	// The thumbprint of the deployment target to match in the query and/or search
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (GetSshConnectionDeploymentTargetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSshConnectionDeploymentTargetsArgs)(nil)).Elem()
}

// A collection of values returned by getSshConnectionDeploymentTargets.
type GetSshConnectionDeploymentTargetsResultOutput struct{ *pulumi.OutputState }

func (GetSshConnectionDeploymentTargetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSshConnectionDeploymentTargetsResult)(nil)).Elem()
}

func (o GetSshConnectionDeploymentTargetsResultOutput) ToGetSshConnectionDeploymentTargetsResultOutput() GetSshConnectionDeploymentTargetsResultOutput {
	return o
}

func (o GetSshConnectionDeploymentTargetsResultOutput) ToGetSshConnectionDeploymentTargetsResultOutputWithContext(ctx context.Context) GetSshConnectionDeploymentTargetsResultOutput {
	return o
}

// A filter to search by deployment ID.
func (o GetSshConnectionDeploymentTargetsResultOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) *string { return v.DeploymentId }).(pulumi.StringPtrOutput)
}

// A filter to search by a list of environment IDs.
func (o GetSshConnectionDeploymentTargetsResultOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) []string { return v.Environments }).(pulumi.StringArrayOutput)
}

// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
func (o GetSshConnectionDeploymentTargetsResultOutput) HealthStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) []string { return v.HealthStatuses }).(pulumi.StringArrayOutput)
}

// An auto-generated identifier that includes the timestamp when this data source was last modified.
func (o GetSshConnectionDeploymentTargetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A filter to search by a list of IDs.
func (o GetSshConnectionDeploymentTargetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A filter to search by the disabled status of a resource.
func (o GetSshConnectionDeploymentTargetsResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// A filter to search by name.
func (o GetSshConnectionDeploymentTargetsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A filter to search by the partial match of a name.
func (o GetSshConnectionDeploymentTargetsResultOutput) PartialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) *string { return v.PartialName }).(pulumi.StringPtrOutput)
}

// A filter to search by a list of role IDs.
func (o GetSshConnectionDeploymentTargetsResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// A list of shell names to match in the query and/or search
func (o GetSshConnectionDeploymentTargetsResultOutput) ShellNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) []string { return v.ShellNames }).(pulumi.StringArrayOutput)
}

// A filter to specify the number of items to skip in the response.
func (o GetSshConnectionDeploymentTargetsResultOutput) Skip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) *int { return v.Skip }).(pulumi.IntPtrOutput)
}

// A list of SSH connection deployment targets that match the filter(s).
func (o GetSshConnectionDeploymentTargetsResultOutput) SshConnectionDeploymentTargets() GetSshConnectionDeploymentTargetsSshConnectionDeploymentTargetArrayOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) []GetSshConnectionDeploymentTargetsSshConnectionDeploymentTarget {
		return v.SshConnectionDeploymentTargets
	}).(GetSshConnectionDeploymentTargetsSshConnectionDeploymentTargetArrayOutput)
}

// A filter to specify the number of items to take (or return) in the response.
func (o GetSshConnectionDeploymentTargetsResultOutput) Take() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) *int { return v.Take }).(pulumi.IntPtrOutput)
}

// A filter to search by a list of tenant tags.
func (o GetSshConnectionDeploymentTargetsResultOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) []string { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// A filter to search by a list of tenant IDs.
func (o GetSshConnectionDeploymentTargetsResultOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) []string { return v.Tenants }).(pulumi.StringArrayOutput)
}

// The thumbprint of the deployment target to match in the query and/or search
func (o GetSshConnectionDeploymentTargetsResultOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSshConnectionDeploymentTargetsResult) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSshConnectionDeploymentTargetsResultOutput{})
}
