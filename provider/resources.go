// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package octopusdeploy

import (
	"fmt"
	"path/filepath"

	"github.com/OctopusDeploy/terraform-provider-octopusdeploy/octopusdeploy"
	"github.com/pierskarsenbarg/pulumi-octopusdeploy/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "octopusdeploy"
	// modules:
	mainMod = "index" // the octopusdeploy module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(octopusdeploy.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "octopusdeploy",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Octopus Deploy",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Piers Karsenbarg",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing octopusdeploy cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "octopusdeploy", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/pierskarenbarg/pulumi-octopusdeploy",
		Repository: "https://github.com/pierskarenbarg/pulumi-octopusdeploy",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "OctopusDeploy",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
			"octopusdeploy_aws_account":                                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AwsAccount")},
			"octopusdeploy_aws_elastic_container_registry":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AwsElasticContainerRegistry")},
			"octopusdeploy_azure_cloud_service_deployment_target":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AzureCloudServiceDeploymentTarget")},
			"octopusdeploy_azure_service_fabric_cluster_deployment_target": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AzureServiceFabricClusterDeploymentTarget")},
			"octopusdeploy_azure_service_principal":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AzureServicePrincipal")},
			"octopusdeploy_azure_subscription_account":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AzureSubscriptionAccount")},
			"octopusdeploy_azure_web_app_deployment_target":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AzureWebAppDeploymentTarget")},
			"octopusdeploy_certificate":                                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate")},
			"octopusdeploy_channel":                                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Channel")},
			"octopusdeploy_cloud_region_deployment_target":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudRegionDeploymentTarget")},
			"octopusdeploy_deployment_process":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DeploymentProcess")},
			"octopusdeploy_docker_container_registry":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DockerContainerRegistry")},
			"octopusdeploy_dynamic_worker_pool":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DynamicWorkerPool")},
			"octopusdeploy_environment":                                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Environment")},
			"octopusdeploy_gcp_account":                                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "GcpAccount")},
			"octopusdeploy_git_credential":                                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "GitCredentials")},
			"octopusdeploy_github_repository_feed":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "GithubRepositoryFeed")},
			"octopusdeploy_helm_feed":                                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HelmFeed")},
			"octopusdeploy_kubernetes_cluster_deployment_target":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "KubernetesClusterDeploymentTarget")},
			"octopusdeploy_library_variable_set":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LibraryVariableSet")},
			"octopusdeploy_lifecycle":                                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Lifecycle")},
			"octopusdeploy_listening_tentacle_deployment_target":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ListeningTentacleDeploymentTarget")},
			"octopusdeploy_machine_policy":                                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MachinePolicy")},
			"octopusdeploy_maven_feed":                                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MavenFeed")},
			"octopusdeploy_nuget_feed":                                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NugetFeed")},
			"octopusdeploy_offline_package_drop_deployment_target":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OfflinePackageDropDeploymentTarget")},
			"octopusdeploy_polling_tentacle_deployment_target":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "PollingTentacleDeploymentTarget")},
			"octopusdeploy_project":                                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Project")},
			"octopusdeploy_project_deployment_target_trigger":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectDeploymentTargetTrigger")},
			"octopusdeploy_project_group":                                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectGroup")},
			"octopusdeploy_runbook":                                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Runbook")},
			"octopusdeploy_runbook_process":                                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "RunbookProcess")},
			"octopusdeploy_scoped_user_role":                               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ScopedUserRole")},
			"octopusdeploy_script_module":                                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ScriptModule")},
			"octopusdeploy_space":                                          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Space")},
			"octopusdeploy_ssh_connection_deployment_target":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SshConnectionDeploymentTarget")},
			"octopusdeploy_ssh_key_account":                                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SshKeyAccount")},
			"octopusdeploy_static_worker_pool":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StaticWorkerPool")},
			"octopusdeploy_tag":                                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Tag")},
			"octopusdeploy_tag_set":                                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TagSet")},
			"octopusdeploy_team":                                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Team")},
			"octopusdeploy_tenant":                                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Tenant")},
			"octopusdeploy_tenant_common_variable":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TenantCommonVariable")},
			"octopusdeploy_tenant_project_variable":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TenantProjectVariable")},
			"octopusdeploy_token_account":                                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TokenAccount")},
			"octopusdeploy_user":                                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "User")},
			"octopusdeploy_user_role":                                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "UserRole")},
			"octopusdeploy_username_password_account":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "UsernamePasswordAccount")},
			"octopusdeploy_variable":                                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Variable")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
			"octopusdeploy_accounts":                                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAccounts")},
			"octopusdeploy_azure_cloud_service_deployment_targets":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAzureCloudServiceDeploymentTargets")},
			"octopusdeploy_azure_service_fabric_cluster_deployment_targets": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAzureServiceFabricClusterDeploymentTargets")},
			"octopusdeploy_azure_web_app_deployment_targets":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAzureWebAppDeploymentTargets")},
			"octopusdeploy_certificates":                                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCertificates")},
			"octopusdeploy_channels":                                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getChannels")},
			"octopusdeploy_cloud_region_deployment_targets":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudRegionDeploymentTargets")},
			"octopusdeploy_deployment_targets":                              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDeploymentTargets")},
			"octopusdeploy_environments":                                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getEnvironments")},
			"octopusdeploy_feeds":                                           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFeeds")},
			"octopusdeploy_git_credentials":                                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGitCredentials")},
			"octopusdeploy_kubernetes_cluster_deployment_targets":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getKubernetesClusterDeploymentTargets")},
			"octopusdeploy_library_variable_sets":                           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLibraryVariableSets")},
			"octopusdeploy_lifecycles":                                      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLifecycles")},
			"octopusdeploy_listening_tentacle_deployment_targets":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getListeningTentacleDeploymentTargets")},
			"octopusdeploy_machine":                                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMachine")},
			"octopusdeploy_machine_policies":                                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMachinePolicies")},
			"octopusdeploy_offline_package_drop_deployment_targets":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOfflinePackageDropDeploymentTargets")},
			"octopusdeploy_polling_tentacle_deployment_targets":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPollingTentacleDeploymentTargets")},
			"octopusdeploy_project_groups":                                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getProjectGroups")},
			"octopusdeploy_projects":                                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getProjects")},
			"octopusdeploy_script_modules":                                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getScriptModules")},
			"octopusdeploy_space":                                           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSpace")},
			"octopusdeploy_spaces":                                          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSpaces")},
			"octopusdeploy_ssh_connection_deployment_targets":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSshConnectionDeploymentTargets")},
			"octopusdeploy_tag_sets":                                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getTagSets")},
			"octopusdeploy_teams":                                           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getTeams")},
			"octopusdeploy_tenants":                                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getTenants")},
			"octopusdeploy_user_roles":                                      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUserRoles")},
			"octopusdeploy_users":                                           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUsers")},
			"octopusdeploy_variables":                                       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVariables")},
			"octopusdeploy_worker_pools":                                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getWorkerPools")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			PackageName: "@pierskarsenbarg/octopusdeploy",
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PackageName: "pierskarsenbarg_pulumi_sdm",
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pierskarsenbarg/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "PiersKarsenbarg",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
