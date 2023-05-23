// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetEnvironmentsEnvironmentArgs : global::Pulumi.InvokeArgs
    {
        [Input("allowDynamicInfrastructure", required: true)]
        public bool AllowDynamicInfrastructure { get; set; }

        /// <summary>
        /// The description of this environment.
        /// </summary>
        [Input("description", required: true)]
        public string Description { get; set; } = null!;

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("jiraExtensionSettings", required: true)]
        private List<Inputs.GetEnvironmentsEnvironmentJiraExtensionSettingArgs>? _jiraExtensionSettings;

        /// <summary>
        /// Provides extension settings for the Jira integration for this environment.
        /// </summary>
        public List<Inputs.GetEnvironmentsEnvironmentJiraExtensionSettingArgs> JiraExtensionSettings
        {
            get => _jiraExtensionSettings ?? (_jiraExtensionSettings = new List<Inputs.GetEnvironmentsEnvironmentJiraExtensionSettingArgs>());
            set => _jiraExtensionSettings = value;
        }

        [Input("jiraServiceManagementExtensionSettings", required: true)]
        private List<Inputs.GetEnvironmentsEnvironmentJiraServiceManagementExtensionSettingArgs>? _jiraServiceManagementExtensionSettings;

        /// <summary>
        /// Provides extension settings for the Jira Service Management (JSM) integration for this environment.
        /// </summary>
        public List<Inputs.GetEnvironmentsEnvironmentJiraServiceManagementExtensionSettingArgs> JiraServiceManagementExtensionSettings
        {
            get => _jiraServiceManagementExtensionSettings ?? (_jiraServiceManagementExtensionSettings = new List<Inputs.GetEnvironmentsEnvironmentJiraServiceManagementExtensionSettingArgs>());
            set => _jiraServiceManagementExtensionSettings = value;
        }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("servicenowExtensionSettings", required: true)]
        private List<Inputs.GetEnvironmentsEnvironmentServicenowExtensionSettingArgs>? _servicenowExtensionSettings;

        /// <summary>
        /// Provides extension settings for the ServiceNow integration for this environment.
        /// </summary>
        public List<Inputs.GetEnvironmentsEnvironmentServicenowExtensionSettingArgs> ServicenowExtensionSettings
        {
            get => _servicenowExtensionSettings ?? (_servicenowExtensionSettings = new List<Inputs.GetEnvironmentsEnvironmentServicenowExtensionSettingArgs>());
            set => _servicenowExtensionSettings = value;
        }

        [Input("slug", required: true)]
        public string Slug { get; set; } = null!;

        /// <summary>
        /// The order number to sort an environment.
        /// </summary>
        [Input("sortOrder", required: true)]
        public int SortOrder { get; set; }

        /// <summary>
        /// The space ID associated with this environment.
        /// </summary>
        [Input("spaceId", required: true)]
        public string SpaceId { get; set; } = null!;

        [Input("useGuidedFailure", required: true)]
        public bool UseGuidedFailure { get; set; }

        public GetEnvironmentsEnvironmentArgs()
        {
        }
        public static new GetEnvironmentsEnvironmentArgs Empty => new GetEnvironmentsEnvironmentArgs();
    }
}