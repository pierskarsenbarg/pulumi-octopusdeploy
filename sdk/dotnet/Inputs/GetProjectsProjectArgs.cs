// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetProjectsProjectInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowDeploymentsToNoTargets", required: true)]
        public Input<bool> AllowDeploymentsToNoTargets { get; set; } = null!;

        [Input("autoCreateRelease", required: true)]
        public Input<bool> AutoCreateRelease { get; set; } = null!;

        [Input("autoDeployReleaseOverrides", required: true)]
        private InputList<string>? _autoDeployReleaseOverrides;
        public InputList<string> AutoDeployReleaseOverrides
        {
            get => _autoDeployReleaseOverrides ?? (_autoDeployReleaseOverrides = new InputList<string>());
            set => _autoDeployReleaseOverrides = value;
        }

        /// <summary>
        /// A filter to search for cloned resources by a project ID.
        /// </summary>
        [Input("clonedFromProjectId", required: true)]
        public Input<string> ClonedFromProjectId { get; set; } = null!;

        [Input("connectivityPolicies", required: true)]
        private InputList<Inputs.GetProjectsProjectConnectivityPolicyInputArgs>? _connectivityPolicies;
        public InputList<Inputs.GetProjectsProjectConnectivityPolicyInputArgs> ConnectivityPolicies
        {
            get => _connectivityPolicies ?? (_connectivityPolicies = new InputList<Inputs.GetProjectsProjectConnectivityPolicyInputArgs>());
            set => _connectivityPolicies = value;
        }

        [Input("defaultGuidedFailureMode", required: true)]
        public Input<string> DefaultGuidedFailureMode { get; set; } = null!;

        [Input("defaultToSkipIfAlreadyInstalled", required: true)]
        public Input<bool> DefaultToSkipIfAlreadyInstalled { get; set; } = null!;

        [Input("deploymentChangesTemplate", required: true)]
        public Input<string> DeploymentChangesTemplate { get; set; } = null!;

        [Input("deploymentProcessId", required: true)]
        public Input<string> DeploymentProcessId { get; set; } = null!;

        /// <summary>
        /// The description of this project.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Treats releases of different channels to the same environment as a separate deployment dimension
        /// </summary>
        [Input("discreteChannelRelease", required: true)]
        public Input<bool> DiscreteChannelRelease { get; set; } = null!;

        [Input("gitAnonymousPersistenceSettings", required: true)]
        private InputList<Inputs.GetProjectsProjectGitAnonymousPersistenceSettingInputArgs>? _gitAnonymousPersistenceSettings;

        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        public InputList<Inputs.GetProjectsProjectGitAnonymousPersistenceSettingInputArgs> GitAnonymousPersistenceSettings
        {
            get => _gitAnonymousPersistenceSettings ?? (_gitAnonymousPersistenceSettings = new InputList<Inputs.GetProjectsProjectGitAnonymousPersistenceSettingInputArgs>());
            set => _gitAnonymousPersistenceSettings = value;
        }

        [Input("gitLibraryPersistenceSettings", required: true)]
        private InputList<Inputs.GetProjectsProjectGitLibraryPersistenceSettingInputArgs>? _gitLibraryPersistenceSettings;

        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        public InputList<Inputs.GetProjectsProjectGitLibraryPersistenceSettingInputArgs> GitLibraryPersistenceSettings
        {
            get => _gitLibraryPersistenceSettings ?? (_gitLibraryPersistenceSettings = new InputList<Inputs.GetProjectsProjectGitLibraryPersistenceSettingInputArgs>());
            set => _gitLibraryPersistenceSettings = value;
        }

        [Input("gitUsernamePasswordPersistenceSettings", required: true)]
        private InputList<Inputs.GetProjectsProjectGitUsernamePasswordPersistenceSettingInputArgs>? _gitUsernamePasswordPersistenceSettings;

        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        public InputList<Inputs.GetProjectsProjectGitUsernamePasswordPersistenceSettingInputArgs> GitUsernamePasswordPersistenceSettings
        {
            get => _gitUsernamePasswordPersistenceSettings ?? (_gitUsernamePasswordPersistenceSettings = new InputList<Inputs.GetProjectsProjectGitUsernamePasswordPersistenceSettingInputArgs>());
            set => _gitUsernamePasswordPersistenceSettings = value;
        }

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("includedLibraryVariableSets", required: true)]
        private InputList<string>? _includedLibraryVariableSets;
        public InputList<string> IncludedLibraryVariableSets
        {
            get => _includedLibraryVariableSets ?? (_includedLibraryVariableSets = new InputList<string>());
            set => _includedLibraryVariableSets = value;
        }

        [Input("isDisabled", required: true)]
        public Input<bool> IsDisabled { get; set; } = null!;

        /// <summary>
        /// Treats releases of different channels to the same environment as a separate deployment dimension
        /// </summary>
        [Input("isDiscreteChannelRelease", required: true)]
        public Input<bool> IsDiscreteChannelRelease { get; set; } = null!;

        [Input("isVersionControlled", required: true)]
        public Input<bool> IsVersionControlled { get; set; } = null!;

        [Input("jiraServiceManagementExtensionSettings", required: true)]
        private InputList<Inputs.GetProjectsProjectJiraServiceManagementExtensionSettingInputArgs>? _jiraServiceManagementExtensionSettings;

        /// <summary>
        /// Provides extension settings for the Jira Service Management (JSM) integration for this project.
        /// </summary>
        public InputList<Inputs.GetProjectsProjectJiraServiceManagementExtensionSettingInputArgs> JiraServiceManagementExtensionSettings
        {
            get => _jiraServiceManagementExtensionSettings ?? (_jiraServiceManagementExtensionSettings = new InputList<Inputs.GetProjectsProjectJiraServiceManagementExtensionSettingInputArgs>());
            set => _jiraServiceManagementExtensionSettings = value;
        }

        /// <summary>
        /// The lifecycle ID associated with this project.
        /// </summary>
        [Input("lifecycleId", required: true)]
        public Input<string> LifecycleId { get; set; } = null!;

        /// <summary>
        /// The name of the project in Octopus Deploy. This name must be unique.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The project group ID associated with this project.
        /// </summary>
        [Input("projectGroupId", required: true)]
        public Input<string> ProjectGroupId { get; set; } = null!;

        [Input("releaseCreationStrategies", required: true)]
        private InputList<Inputs.GetProjectsProjectReleaseCreationStrategyInputArgs>? _releaseCreationStrategies;
        public InputList<Inputs.GetProjectsProjectReleaseCreationStrategyInputArgs> ReleaseCreationStrategies
        {
            get => _releaseCreationStrategies ?? (_releaseCreationStrategies = new InputList<Inputs.GetProjectsProjectReleaseCreationStrategyInputArgs>());
            set => _releaseCreationStrategies = value;
        }

        [Input("releaseNotesTemplate", required: true)]
        public Input<string> ReleaseNotesTemplate { get; set; } = null!;

        [Input("servicenowExtensionSettings", required: true)]
        private InputList<Inputs.GetProjectsProjectServicenowExtensionSettingInputArgs>? _servicenowExtensionSettings;

        /// <summary>
        /// Provides extension settings for the ServiceNow integration for this project.
        /// </summary>
        public InputList<Inputs.GetProjectsProjectServicenowExtensionSettingInputArgs> ServicenowExtensionSettings
        {
            get => _servicenowExtensionSettings ?? (_servicenowExtensionSettings = new InputList<Inputs.GetProjectsProjectServicenowExtensionSettingInputArgs>());
            set => _servicenowExtensionSettings = value;
        }

        /// <summary>
        /// A human-readable, unique identifier, used to identify a project.
        /// </summary>
        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        /// <summary>
        /// The space ID associated with this project.
        /// </summary>
        [Input("spaceId", required: true)]
        public Input<string> SpaceId { get; set; } = null!;

        [Input("templates", required: true)]
        private InputList<Inputs.GetProjectsProjectTemplateInputArgs>? _templates;
        public InputList<Inputs.GetProjectsProjectTemplateInputArgs> Templates
        {
            get => _templates ?? (_templates = new InputList<Inputs.GetProjectsProjectTemplateInputArgs>());
            set => _templates = value;
        }

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Input("tenantedDeploymentParticipation", required: true)]
        public Input<string> TenantedDeploymentParticipation { get; set; } = null!;

        [Input("variableSetId", required: true)]
        public Input<string> VariableSetId { get; set; } = null!;

        [Input("versioningStrategies", required: true)]
        private InputList<Inputs.GetProjectsProjectVersioningStrategyInputArgs>? _versioningStrategies;
        public InputList<Inputs.GetProjectsProjectVersioningStrategyInputArgs> VersioningStrategies
        {
            get => _versioningStrategies ?? (_versioningStrategies = new InputList<Inputs.GetProjectsProjectVersioningStrategyInputArgs>());
            set => _versioningStrategies = value;
        }

        public GetProjectsProjectInputArgs()
        {
        }
        public static new GetProjectsProjectInputArgs Empty => new GetProjectsProjectInputArgs();
    }
}
