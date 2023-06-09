// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy
{
    /// <summary>
    /// This resource manages runbooks in Octopus Deploy.
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/runbook:Runbook")]
    public partial class Runbook : global::Pulumi.CustomResource
    {
        [Output("connectivityPolicy")]
        public Output<Outputs.RunbookConnectivityPolicy> ConnectivityPolicy { get; private set; } = null!;

        /// <summary>
        /// Sets the runbook guided failure mode.
        /// </summary>
        [Output("defaultGuidedFailureMode")]
        public Output<string> DefaultGuidedFailureMode { get; private set; } = null!;

        /// <summary>
        /// The description of this runbook.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Determines how the runbook is scoped to environments.
        /// </summary>
        [Output("environmentScope")]
        public Output<string> EnvironmentScope { get; private set; } = null!;

        /// <summary>
        /// When environment_scope is set to "Specified", this is the list of environments the runbook can be run against.
        /// </summary>
        [Output("environments")]
        public Output<ImmutableArray<string>> Environments { get; private set; } = null!;

        /// <summary>
        /// Whether to force packages to be re-downloaded or not
        /// </summary>
        [Output("forcePackageDownload")]
        public Output<bool> ForcePackageDownload { get; private set; } = null!;

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Output("multiTenancyMode")]
        public Output<string> MultiTenancyMode { get; private set; } = null!;

        /// <summary>
        /// The name of the runbook in Octopus Deploy. This name must be unique.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project that this runbook belongs to.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The published snapshot ID.
        /// </summary>
        [Output("publishedRunbookSnapshotId")]
        public Output<string> PublishedRunbookSnapshotId { get; private set; } = null!;

        /// <summary>
        /// Sets the runbook retention policy
        /// </summary>
        [Output("retentionPolicy")]
        public Output<Outputs.RunbookRetentionPolicy> RetentionPolicy { get; private set; } = null!;

        /// <summary>
        /// The runbook process ID.
        /// </summary>
        [Output("runbookProcessId")]
        public Output<string> RunbookProcessId { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this runbook.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;


        /// <summary>
        /// Create a Runbook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Runbook(string name, RunbookArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/runbook:Runbook", name, args ?? new RunbookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Runbook(string name, Input<string> id, RunbookState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/runbook:Runbook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Runbook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Runbook Get(string name, Input<string> id, RunbookState? state = null, CustomResourceOptions? options = null)
        {
            return new Runbook(name, id, state, options);
        }
    }

    public sealed class RunbookArgs : global::Pulumi.ResourceArgs
    {
        [Input("connectivityPolicy")]
        public Input<Inputs.RunbookConnectivityPolicyArgs>? ConnectivityPolicy { get; set; }

        /// <summary>
        /// Sets the runbook guided failure mode.
        /// </summary>
        [Input("defaultGuidedFailureMode")]
        public Input<string>? DefaultGuidedFailureMode { get; set; }

        /// <summary>
        /// The description of this runbook.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Determines how the runbook is scoped to environments.
        /// </summary>
        [Input("environmentScope")]
        public Input<string>? EnvironmentScope { get; set; }

        [Input("environments")]
        private InputList<string>? _environments;

        /// <summary>
        /// When environment_scope is set to "Specified", this is the list of environments the runbook can be run against.
        /// </summary>
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        /// <summary>
        /// Whether to force packages to be re-downloaded or not
        /// </summary>
        [Input("forcePackageDownload")]
        public Input<bool>? ForcePackageDownload { get; set; }

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Input("multiTenancyMode")]
        public Input<string>? MultiTenancyMode { get; set; }

        /// <summary>
        /// The name of the runbook in Octopus Deploy. This name must be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project that this runbook belongs to.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Sets the runbook retention policy
        /// </summary>
        [Input("retentionPolicy")]
        public Input<Inputs.RunbookRetentionPolicyArgs>? RetentionPolicy { get; set; }

        /// <summary>
        /// The space ID associated with this runbook.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        public RunbookArgs()
        {
        }
        public static new RunbookArgs Empty => new RunbookArgs();
    }

    public sealed class RunbookState : global::Pulumi.ResourceArgs
    {
        [Input("connectivityPolicy")]
        public Input<Inputs.RunbookConnectivityPolicyGetArgs>? ConnectivityPolicy { get; set; }

        /// <summary>
        /// Sets the runbook guided failure mode.
        /// </summary>
        [Input("defaultGuidedFailureMode")]
        public Input<string>? DefaultGuidedFailureMode { get; set; }

        /// <summary>
        /// The description of this runbook.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Determines how the runbook is scoped to environments.
        /// </summary>
        [Input("environmentScope")]
        public Input<string>? EnvironmentScope { get; set; }

        [Input("environments")]
        private InputList<string>? _environments;

        /// <summary>
        /// When environment_scope is set to "Specified", this is the list of environments the runbook can be run against.
        /// </summary>
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        /// <summary>
        /// Whether to force packages to be re-downloaded or not
        /// </summary>
        [Input("forcePackageDownload")]
        public Input<bool>? ForcePackageDownload { get; set; }

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Input("multiTenancyMode")]
        public Input<string>? MultiTenancyMode { get; set; }

        /// <summary>
        /// The name of the runbook in Octopus Deploy. This name must be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project that this runbook belongs to.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The published snapshot ID.
        /// </summary>
        [Input("publishedRunbookSnapshotId")]
        public Input<string>? PublishedRunbookSnapshotId { get; set; }

        /// <summary>
        /// Sets the runbook retention policy
        /// </summary>
        [Input("retentionPolicy")]
        public Input<Inputs.RunbookRetentionPolicyGetArgs>? RetentionPolicy { get; set; }

        /// <summary>
        /// The runbook process ID.
        /// </summary>
        [Input("runbookProcessId")]
        public Input<string>? RunbookProcessId { get; set; }

        /// <summary>
        /// The space ID associated with this runbook.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        public RunbookState()
        {
        }
        public static new RunbookState Empty => new RunbookState();
    }
}
