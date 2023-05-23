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
    /// This resource manages static worker pools in Octopus Deploy.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Octopusdeploy = PiersKarsenbarg.Octopusdeploy;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Octopusdeploy.StaticWorkerPool("example", new()
    ///     {
    ///         Description = "Description for the static worker pool.",
    ///         IsDefault = true,
    ///         SortOrder = 5,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/staticWorkerPool:StaticWorkerPool")]
    public partial class StaticWorkerPool : global::Pulumi.CustomResource
    {
        [Output("canAddWorkers")]
        public Output<bool> CanAddWorkers { get; private set; } = null!;

        /// <summary>
        /// The description of this static worker pool.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("isDefault")]
        public Output<bool?> IsDefault { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The order number to sort a dynamic worker pool.
        /// </summary>
        [Output("sortOrder")]
        public Output<int> SortOrder { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;


        /// <summary>
        /// Create a StaticWorkerPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StaticWorkerPool(string name, StaticWorkerPoolArgs? args = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/staticWorkerPool:StaticWorkerPool", name, args ?? new StaticWorkerPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StaticWorkerPool(string name, Input<string> id, StaticWorkerPoolState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/staticWorkerPool:StaticWorkerPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StaticWorkerPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StaticWorkerPool Get(string name, Input<string> id, StaticWorkerPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new StaticWorkerPool(name, id, state, options);
        }
    }

    public sealed class StaticWorkerPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this static worker pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The order number to sort a dynamic worker pool.
        /// </summary>
        [Input("sortOrder")]
        public Input<int>? SortOrder { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        public StaticWorkerPoolArgs()
        {
        }
        public static new StaticWorkerPoolArgs Empty => new StaticWorkerPoolArgs();
    }

    public sealed class StaticWorkerPoolState : global::Pulumi.ResourceArgs
    {
        [Input("canAddWorkers")]
        public Input<bool>? CanAddWorkers { get; set; }

        /// <summary>
        /// The description of this static worker pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The order number to sort a dynamic worker pool.
        /// </summary>
        [Input("sortOrder")]
        public Input<int>? SortOrder { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        public StaticWorkerPoolState()
        {
        }
        public static new StaticWorkerPoolState Empty => new StaticWorkerPoolState();
    }
}
