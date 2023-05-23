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
    /// This resource manages lifecycles in Octopus Deploy.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Octopusdeploy = Pulumi.Octopusdeploy;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Octopusdeploy.Lifecycle("example", new()
    ///     {
    ///         Description = "This is the default lifecycle.",
    ///         Phases = new[]
    ///         {
    ///             new Octopusdeploy.Inputs.LifecyclePhaseArgs
    ///             {
    ///                 AutomaticDeploymentTargets = new[]
    ///                 {
    ///                     "Environments-321",
    ///                 },
    ///                 Name = "foo",
    ///                 ReleaseRetentionPolicy = new Octopusdeploy.Inputs.LifecyclePhaseReleaseRetentionPolicyArgs
    ///                 {
    ///                     QuantityToKeep = 1,
    ///                     ShouldKeepForever = true,
    ///                     Unit = "Days",
    ///                 },
    ///                 TentacleRetentionPolicy = new Octopusdeploy.Inputs.LifecyclePhaseTentacleRetentionPolicyArgs
    ///                 {
    ///                     QuantityToKeep = 30,
    ///                     ShouldKeepForever = false,
    ///                     Unit = "Items",
    ///                 },
    ///             },
    ///             new Octopusdeploy.Inputs.LifecyclePhaseArgs
    ///             {
    ///                 IsOptionalPhase = true,
    ///                 Name = "bar",
    ///                 OptionalDeploymentTargets = new[]
    ///                 {
    ///                     "Environments-321",
    ///                 },
    ///             },
    ///         },
    ///         ReleaseRetentionPolicy = new Octopusdeploy.Inputs.LifecycleReleaseRetentionPolicyArgs
    ///         {
    ///             QuantityToKeep = 1,
    ///             ShouldKeepForever = true,
    ///             Unit = "Days",
    ///         },
    ///         TentacleRetentionPolicy = new Octopusdeploy.Inputs.LifecycleTentacleRetentionPolicyArgs
    ///         {
    ///             QuantityToKeep = 30,
    ///             ShouldKeepForever = false,
    ///             Unit = "Items",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import octopusdeploy:index/lifecycle:Lifecycle [options] octopusdeploy_lifecycle.&lt;name&gt; &lt;lifecycle-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/lifecycle:Lifecycle")]
    public partial class Lifecycle : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of this lifecycle.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("phases")]
        public Output<ImmutableArray<Outputs.LifecyclePhase>> Phases { get; private set; } = null!;

        [Output("releaseRetentionPolicy")]
        public Output<Outputs.LifecycleReleaseRetentionPolicy> ReleaseRetentionPolicy { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

        [Output("tentacleRetentionPolicy")]
        public Output<Outputs.LifecycleTentacleRetentionPolicy> TentacleRetentionPolicy { get; private set; } = null!;


        /// <summary>
        /// Create a Lifecycle resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Lifecycle(string name, LifecycleArgs? args = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/lifecycle:Lifecycle", name, args ?? new LifecycleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Lifecycle(string name, Input<string> id, LifecycleState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/lifecycle:Lifecycle", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Lifecycle resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Lifecycle Get(string name, Input<string> id, LifecycleState? state = null, CustomResourceOptions? options = null)
        {
            return new Lifecycle(name, id, state, options);
        }
    }

    public sealed class LifecycleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this lifecycle.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("phases")]
        private InputList<Inputs.LifecyclePhaseArgs>? _phases;
        public InputList<Inputs.LifecyclePhaseArgs> Phases
        {
            get => _phases ?? (_phases = new InputList<Inputs.LifecyclePhaseArgs>());
            set => _phases = value;
        }

        [Input("releaseRetentionPolicy")]
        public Input<Inputs.LifecycleReleaseRetentionPolicyArgs>? ReleaseRetentionPolicy { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        [Input("tentacleRetentionPolicy")]
        public Input<Inputs.LifecycleTentacleRetentionPolicyArgs>? TentacleRetentionPolicy { get; set; }

        public LifecycleArgs()
        {
        }
        public static new LifecycleArgs Empty => new LifecycleArgs();
    }

    public sealed class LifecycleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this lifecycle.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("phases")]
        private InputList<Inputs.LifecyclePhaseGetArgs>? _phases;
        public InputList<Inputs.LifecyclePhaseGetArgs> Phases
        {
            get => _phases ?? (_phases = new InputList<Inputs.LifecyclePhaseGetArgs>());
            set => _phases = value;
        }

        [Input("releaseRetentionPolicy")]
        public Input<Inputs.LifecycleReleaseRetentionPolicyGetArgs>? ReleaseRetentionPolicy { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        [Input("tentacleRetentionPolicy")]
        public Input<Inputs.LifecycleTentacleRetentionPolicyGetArgs>? TentacleRetentionPolicy { get; set; }

        public LifecycleState()
        {
        }
        public static new LifecycleState Empty => new LifecycleState();
    }
}
