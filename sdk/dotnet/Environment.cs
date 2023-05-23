// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    /// <summary>
    /// This resource manages environments in Octopus Deploy.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Octopusdeploy = Pulumi.Octopusdeploy;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Octopusdeploy.Environment("example", new()
    ///     {
    ///         AllowDynamicInfrastructure = false,
    ///         Description = "An environment for the development team.",
    ///         JiraExtensionSettings = new Octopusdeploy.Inputs.EnvironmentJiraExtensionSettingsArgs
    ///         {
    ///             EnvironmentType = "unmapped",
    ///         },
    ///         JiraServiceManagementExtensionSettings = new Octopusdeploy.Inputs.EnvironmentJiraServiceManagementExtensionSettingsArgs
    ///         {
    ///             IsEnabled = false,
    ///         },
    ///         ServicenowExtensionSettings = new Octopusdeploy.Inputs.EnvironmentServicenowExtensionSettingsArgs
    ///         {
    ///             IsEnabled = false,
    ///         },
    ///         UseGuidedFailure = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import octopusdeploy:index/environment:Environment [options] octopusdeploy_environment.&lt;name&gt; &lt;environment-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/environment:Environment")]
    public partial class Environment : global::Pulumi.CustomResource
    {
        [Output("allowDynamicInfrastructure")]
        public Output<bool?> AllowDynamicInfrastructure { get; private set; } = null!;

        /// <summary>
        /// The description of this environment.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Provides extension settings for the Jira integration for this environment.
        /// </summary>
        [Output("jiraExtensionSettings")]
        public Output<Outputs.EnvironmentJiraExtensionSettings?> JiraExtensionSettings { get; private set; } = null!;

        /// <summary>
        /// Provides extension settings for the Jira Service Management (JSM) integration for this environment.
        /// </summary>
        [Output("jiraServiceManagementExtensionSettings")]
        public Output<Outputs.EnvironmentJiraServiceManagementExtensionSettings?> JiraServiceManagementExtensionSettings { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provides extension settings for the ServiceNow integration for this environment.
        /// </summary>
        [Output("servicenowExtensionSettings")]
        public Output<Outputs.EnvironmentServicenowExtensionSettings?> ServicenowExtensionSettings { get; private set; } = null!;

        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        /// <summary>
        /// The order number to sort an environment.
        /// </summary>
        [Output("sortOrder")]
        public Output<int> SortOrder { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this environment.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

        [Output("useGuidedFailure")]
        public Output<bool?> UseGuidedFailure { get; private set; } = null!;


        /// <summary>
        /// Create a Environment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Environment(string name, EnvironmentArgs? args = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/environment:Environment", name, args ?? new EnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Environment(string name, Input<string> id, EnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/environment:Environment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Environment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Environment Get(string name, Input<string> id, EnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Environment(name, id, state, options);
        }
    }

    public sealed class EnvironmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowDynamicInfrastructure")]
        public Input<bool>? AllowDynamicInfrastructure { get; set; }

        /// <summary>
        /// The description of this environment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Provides extension settings for the Jira integration for this environment.
        /// </summary>
        [Input("jiraExtensionSettings")]
        public Input<Inputs.EnvironmentJiraExtensionSettingsArgs>? JiraExtensionSettings { get; set; }

        /// <summary>
        /// Provides extension settings for the Jira Service Management (JSM) integration for this environment.
        /// </summary>
        [Input("jiraServiceManagementExtensionSettings")]
        public Input<Inputs.EnvironmentJiraServiceManagementExtensionSettingsArgs>? JiraServiceManagementExtensionSettings { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Provides extension settings for the ServiceNow integration for this environment.
        /// </summary>
        [Input("servicenowExtensionSettings")]
        public Input<Inputs.EnvironmentServicenowExtensionSettingsArgs>? ServicenowExtensionSettings { get; set; }

        /// <summary>
        /// The order number to sort an environment.
        /// </summary>
        [Input("sortOrder")]
        public Input<int>? SortOrder { get; set; }

        /// <summary>
        /// The space ID associated with this environment.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        [Input("useGuidedFailure")]
        public Input<bool>? UseGuidedFailure { get; set; }

        public EnvironmentArgs()
        {
        }
        public static new EnvironmentArgs Empty => new EnvironmentArgs();
    }

    public sealed class EnvironmentState : global::Pulumi.ResourceArgs
    {
        [Input("allowDynamicInfrastructure")]
        public Input<bool>? AllowDynamicInfrastructure { get; set; }

        /// <summary>
        /// The description of this environment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Provides extension settings for the Jira integration for this environment.
        /// </summary>
        [Input("jiraExtensionSettings")]
        public Input<Inputs.EnvironmentJiraExtensionSettingsGetArgs>? JiraExtensionSettings { get; set; }

        /// <summary>
        /// Provides extension settings for the Jira Service Management (JSM) integration for this environment.
        /// </summary>
        [Input("jiraServiceManagementExtensionSettings")]
        public Input<Inputs.EnvironmentJiraServiceManagementExtensionSettingsGetArgs>? JiraServiceManagementExtensionSettings { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Provides extension settings for the ServiceNow integration for this environment.
        /// </summary>
        [Input("servicenowExtensionSettings")]
        public Input<Inputs.EnvironmentServicenowExtensionSettingsGetArgs>? ServicenowExtensionSettings { get; set; }

        [Input("slug")]
        public Input<string>? Slug { get; set; }

        /// <summary>
        /// The order number to sort an environment.
        /// </summary>
        [Input("sortOrder")]
        public Input<int>? SortOrder { get; set; }

        /// <summary>
        /// The space ID associated with this environment.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        [Input("useGuidedFailure")]
        public Input<bool>? UseGuidedFailure { get; set; }

        public EnvironmentState()
        {
        }
        public static new EnvironmentState Empty => new EnvironmentState();
    }
}