// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class GetProjectsProjectResult
    {
        public readonly bool AllowDeploymentsToNoTargets;
        public readonly bool AutoCreateRelease;
        public readonly ImmutableArray<string> AutoDeployReleaseOverrides;
        /// <summary>
        /// A filter to search for cloned resources by a project ID.
        /// </summary>
        public readonly string ClonedFromProjectId;
        public readonly ImmutableArray<Outputs.GetProjectsProjectConnectivityPolicyResult> ConnectivityPolicies;
        public readonly string DefaultGuidedFailureMode;
        public readonly bool DefaultToSkipIfAlreadyInstalled;
        public readonly string DeploymentChangesTemplate;
        public readonly string DeploymentProcessId;
        /// <summary>
        /// The description of this project.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Treats releases of different channels to the same environment as a separate deployment dimension
        /// </summary>
        public readonly bool DiscreteChannelRelease;
        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectsProjectGitAnonymousPersistenceSettingResult> GitAnonymousPersistenceSettings;
        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectsProjectGitLibraryPersistenceSettingResult> GitLibraryPersistenceSettings;
        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectsProjectGitUsernamePasswordPersistenceSettingResult> GitUsernamePasswordPersistenceSettings;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> IncludedLibraryVariableSets;
        public readonly bool IsDisabled;
        /// <summary>
        /// Treats releases of different channels to the same environment as a separate deployment dimension
        /// </summary>
        public readonly bool IsDiscreteChannelRelease;
        public readonly bool IsVersionControlled;
        /// <summary>
        /// Provides extension settings for the Jira Service Management (JSM) integration for this project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectsProjectJiraServiceManagementExtensionSettingResult> JiraServiceManagementExtensionSettings;
        /// <summary>
        /// The lifecycle ID associated with this project.
        /// </summary>
        public readonly string LifecycleId;
        /// <summary>
        /// The name of the project in Octopus Deploy. This name must be unique.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The project group ID associated with this project.
        /// </summary>
        public readonly string ProjectGroupId;
        public readonly ImmutableArray<Outputs.GetProjectsProjectReleaseCreationStrategyResult> ReleaseCreationStrategies;
        public readonly string ReleaseNotesTemplate;
        /// <summary>
        /// Provides extension settings for the ServiceNow integration for this project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectsProjectServicenowExtensionSettingResult> ServicenowExtensionSettings;
        /// <summary>
        /// A human-readable, unique identifier, used to identify a project.
        /// </summary>
        public readonly string Slug;
        /// <summary>
        /// The space ID associated with this project.
        /// </summary>
        public readonly string SpaceId;
        public readonly ImmutableArray<Outputs.GetProjectsProjectTemplateResult> Templates;
        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        public readonly string TenantedDeploymentParticipation;
        public readonly string VariableSetId;
        public readonly ImmutableArray<Outputs.GetProjectsProjectVersioningStrategyResult> VersioningStrategies;

        [OutputConstructor]
        private GetProjectsProjectResult(
            bool allowDeploymentsToNoTargets,

            bool autoCreateRelease,

            ImmutableArray<string> autoDeployReleaseOverrides,

            string clonedFromProjectId,

            ImmutableArray<Outputs.GetProjectsProjectConnectivityPolicyResult> connectivityPolicies,

            string defaultGuidedFailureMode,

            bool defaultToSkipIfAlreadyInstalled,

            string deploymentChangesTemplate,

            string deploymentProcessId,

            string description,

            bool discreteChannelRelease,

            ImmutableArray<Outputs.GetProjectsProjectGitAnonymousPersistenceSettingResult> gitAnonymousPersistenceSettings,

            ImmutableArray<Outputs.GetProjectsProjectGitLibraryPersistenceSettingResult> gitLibraryPersistenceSettings,

            ImmutableArray<Outputs.GetProjectsProjectGitUsernamePasswordPersistenceSettingResult> gitUsernamePasswordPersistenceSettings,

            string id,

            ImmutableArray<string> includedLibraryVariableSets,

            bool isDisabled,

            bool isDiscreteChannelRelease,

            bool isVersionControlled,

            ImmutableArray<Outputs.GetProjectsProjectJiraServiceManagementExtensionSettingResult> jiraServiceManagementExtensionSettings,

            string lifecycleId,

            string name,

            string projectGroupId,

            ImmutableArray<Outputs.GetProjectsProjectReleaseCreationStrategyResult> releaseCreationStrategies,

            string releaseNotesTemplate,

            ImmutableArray<Outputs.GetProjectsProjectServicenowExtensionSettingResult> servicenowExtensionSettings,

            string slug,

            string spaceId,

            ImmutableArray<Outputs.GetProjectsProjectTemplateResult> templates,

            string tenantedDeploymentParticipation,

            string variableSetId,

            ImmutableArray<Outputs.GetProjectsProjectVersioningStrategyResult> versioningStrategies)
        {
            AllowDeploymentsToNoTargets = allowDeploymentsToNoTargets;
            AutoCreateRelease = autoCreateRelease;
            AutoDeployReleaseOverrides = autoDeployReleaseOverrides;
            ClonedFromProjectId = clonedFromProjectId;
            ConnectivityPolicies = connectivityPolicies;
            DefaultGuidedFailureMode = defaultGuidedFailureMode;
            DefaultToSkipIfAlreadyInstalled = defaultToSkipIfAlreadyInstalled;
            DeploymentChangesTemplate = deploymentChangesTemplate;
            DeploymentProcessId = deploymentProcessId;
            Description = description;
            DiscreteChannelRelease = discreteChannelRelease;
            GitAnonymousPersistenceSettings = gitAnonymousPersistenceSettings;
            GitLibraryPersistenceSettings = gitLibraryPersistenceSettings;
            GitUsernamePasswordPersistenceSettings = gitUsernamePasswordPersistenceSettings;
            Id = id;
            IncludedLibraryVariableSets = includedLibraryVariableSets;
            IsDisabled = isDisabled;
            IsDiscreteChannelRelease = isDiscreteChannelRelease;
            IsVersionControlled = isVersionControlled;
            JiraServiceManagementExtensionSettings = jiraServiceManagementExtensionSettings;
            LifecycleId = lifecycleId;
            Name = name;
            ProjectGroupId = projectGroupId;
            ReleaseCreationStrategies = releaseCreationStrategies;
            ReleaseNotesTemplate = releaseNotesTemplate;
            ServicenowExtensionSettings = servicenowExtensionSettings;
            Slug = slug;
            SpaceId = spaceId;
            Templates = templates;
            TenantedDeploymentParticipation = tenantedDeploymentParticipation;
            VariableSetId = variableSetId;
            VersioningStrategies = versioningStrategies;
        }
    }
}