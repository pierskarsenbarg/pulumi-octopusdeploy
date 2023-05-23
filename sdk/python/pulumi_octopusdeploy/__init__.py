# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .aws_account import *
from .aws_elastic_container_registry import *
from .azure_cloud_service_deployment_target import *
from .azure_service_fabric_cluster_deployment_target import *
from .azure_service_principal import *
from .azure_subscription_account import *
from .azure_web_app_deployment_target import *
from .certificate import *
from .channel import *
from .cloud_region_deployment_target import *
from .deployment_process import *
from .docker_container_registry import *
from .dynamic_worker_pool import *
from .environment import *
from .gcp_account import *
from .get_accounts import *
from .get_azure_cloud_service_deployment_targets import *
from .get_azure_service_fabric_cluster_deployment_targets import *
from .get_azure_web_app_deployment_targets import *
from .get_certificates import *
from .get_channels import *
from .get_cloud_region_deployment_targets import *
from .get_deployment_targets import *
from .get_environments import *
from .get_feeds import *
from .get_git_credentials import *
from .get_kubernetes_cluster_deployment_targets import *
from .get_library_variable_sets import *
from .get_lifecycles import *
from .get_listening_tentacle_deployment_targets import *
from .get_machine import *
from .get_machine_policies import *
from .get_offline_package_drop_deployment_targets import *
from .get_polling_tentacle_deployment_targets import *
from .get_project_groups import *
from .get_projects import *
from .get_script_modules import *
from .get_space import *
from .get_spaces import *
from .get_ssh_connection_deployment_targets import *
from .get_tag_sets import *
from .get_teams import *
from .get_tenants import *
from .get_user_roles import *
from .get_users import *
from .get_variables import *
from .get_worker_pools import *
from .git_credentials import *
from .github_repository_feed import *
from .helm_feed import *
from .kubernetes_cluster_deployment_target import *
from .library_variable_set import *
from .lifecycle import *
from .listening_tentacle_deployment_target import *
from .machine_policy import *
from .maven_feed import *
from .nuget_feed import *
from .offline_package_drop_deployment_target import *
from .polling_tentacle_deployment_target import *
from .project import *
from .project_deployment_target_trigger import *
from .project_group import *
from .provider import *
from .runbook import *
from .runbook_process import *
from .scoped_user_role import *
from .script_module import *
from .space import *
from .ssh_connection_deployment_target import *
from .ssh_key_account import *
from .static_worker_pool import *
from .tag import *
from .tag_set import *
from .team import *
from .tenant import *
from .tenant_common_variable import *
from .tenant_project_variable import *
from .token_account import *
from .user import *
from .user_role import *
from .username_password_account import *
from .variable import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_octopusdeploy.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_octopusdeploy.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "octopusdeploy",
  "mod": "index/awsAccount",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/awsAccount:AwsAccount": "AwsAccount"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/awsElasticContainerRegistry",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry": "AwsElasticContainerRegistry"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/azureCloudServiceDeploymentTarget",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/azureCloudServiceDeploymentTarget:AzureCloudServiceDeploymentTarget": "AzureCloudServiceDeploymentTarget"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/azureServiceFabricClusterDeploymentTarget",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/azureServiceFabricClusterDeploymentTarget:AzureServiceFabricClusterDeploymentTarget": "AzureServiceFabricClusterDeploymentTarget"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/azureServicePrincipal",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/azureServicePrincipal:AzureServicePrincipal": "AzureServicePrincipal"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/azureSubscriptionAccount",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/azureSubscriptionAccount:AzureSubscriptionAccount": "AzureSubscriptionAccount"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/azureWebAppDeploymentTarget",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/azureWebAppDeploymentTarget:AzureWebAppDeploymentTarget": "AzureWebAppDeploymentTarget"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/certificate",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/certificate:Certificate": "Certificate"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/channel",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/channel:Channel": "Channel"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/cloudRegionDeploymentTarget",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/cloudRegionDeploymentTarget:CloudRegionDeploymentTarget": "CloudRegionDeploymentTarget"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/deploymentProcess",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/deploymentProcess:DeploymentProcess": "DeploymentProcess"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/dockerContainerRegistry",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/dockerContainerRegistry:DockerContainerRegistry": "DockerContainerRegistry"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/dynamicWorkerPool",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/dynamicWorkerPool:DynamicWorkerPool": "DynamicWorkerPool"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/environment",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/environment:Environment": "Environment"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/gcpAccount",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/gcpAccount:GcpAccount": "GcpAccount"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/gitCredentials",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/gitCredentials:GitCredentials": "GitCredentials"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/githubRepositoryFeed",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/githubRepositoryFeed:GithubRepositoryFeed": "GithubRepositoryFeed"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/helmFeed",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/helmFeed:HelmFeed": "HelmFeed"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/kubernetesClusterDeploymentTarget",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/kubernetesClusterDeploymentTarget:KubernetesClusterDeploymentTarget": "KubernetesClusterDeploymentTarget"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/libraryVariableSet",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/libraryVariableSet:LibraryVariableSet": "LibraryVariableSet"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/lifecycle",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/lifecycle:Lifecycle": "Lifecycle"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/listeningTentacleDeploymentTarget",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/listeningTentacleDeploymentTarget:ListeningTentacleDeploymentTarget": "ListeningTentacleDeploymentTarget"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/machinePolicy",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/machinePolicy:MachinePolicy": "MachinePolicy"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/mavenFeed",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/mavenFeed:MavenFeed": "MavenFeed"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/nugetFeed",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/nugetFeed:NugetFeed": "NugetFeed"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/offlinePackageDropDeploymentTarget",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/offlinePackageDropDeploymentTarget:OfflinePackageDropDeploymentTarget": "OfflinePackageDropDeploymentTarget"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/pollingTentacleDeploymentTarget",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/pollingTentacleDeploymentTarget:PollingTentacleDeploymentTarget": "PollingTentacleDeploymentTarget"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/project",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/project:Project": "Project"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/projectDeploymentTargetTrigger",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/projectDeploymentTargetTrigger:ProjectDeploymentTargetTrigger": "ProjectDeploymentTargetTrigger"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/projectGroup",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/projectGroup:ProjectGroup": "ProjectGroup"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/runbook",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/runbook:Runbook": "Runbook"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/runbookProcess",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/runbookProcess:RunbookProcess": "RunbookProcess"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/scopedUserRole",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/scopedUserRole:ScopedUserRole": "ScopedUserRole"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/scriptModule",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/scriptModule:ScriptModule": "ScriptModule"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/space",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/space:Space": "Space"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/sshConnectionDeploymentTarget",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/sshConnectionDeploymentTarget:SshConnectionDeploymentTarget": "SshConnectionDeploymentTarget"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/sshKeyAccount",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/sshKeyAccount:SshKeyAccount": "SshKeyAccount"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/staticWorkerPool",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/staticWorkerPool:StaticWorkerPool": "StaticWorkerPool"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/tag",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/tag:Tag": "Tag"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/tagSet",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/tagSet:TagSet": "TagSet"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/team",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/team:Team": "Team"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/tenant",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/tenant:Tenant": "Tenant"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/tenantCommonVariable",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/tenantCommonVariable:TenantCommonVariable": "TenantCommonVariable"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/tenantProjectVariable",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/tenantProjectVariable:TenantProjectVariable": "TenantProjectVariable"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/tokenAccount",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/tokenAccount:TokenAccount": "TokenAccount"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/user",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/user:User": "User"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/userRole",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/userRole:UserRole": "UserRole"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/usernamePasswordAccount",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/usernamePasswordAccount:UsernamePasswordAccount": "UsernamePasswordAccount"
  }
 },
 {
  "pkg": "octopusdeploy",
  "mod": "index/variable",
  "fqn": "pulumi_octopusdeploy",
  "classes": {
   "octopusdeploy:index/variable:Variable": "Variable"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "octopusdeploy",
  "token": "pulumi:providers:octopusdeploy",
  "fqn": "pulumi_octopusdeploy",
  "class": "Provider"
 }
]
"""
)
