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
    /// This resource manages project groups in Octopus Deploy.
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
    ///     var example = new Octopusdeploy.ProjectGroup("example", new()
    ///     {
    ///         Description = "The development project group.",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import octopusdeploy:index/projectGroup:ProjectGroup [options] octopusdeploy_project_group.&lt;name&gt; &lt;project_group-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/projectGroup:ProjectGroup")]
    public partial class ProjectGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of this project group.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the retention policy associated with this project group.
        /// </summary>
        [Output("retentionPolicyId")]
        public Output<string> RetentionPolicyId { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this project group.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectGroup(string name, ProjectGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/projectGroup:ProjectGroup", name, args ?? new ProjectGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectGroup(string name, Input<string> id, ProjectGroupState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/projectGroup:ProjectGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectGroup Get(string name, Input<string> id, ProjectGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectGroup(name, id, state, options);
        }
    }

    public sealed class ProjectGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this project group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the retention policy associated with this project group.
        /// </summary>
        [Input("retentionPolicyId")]
        public Input<string>? RetentionPolicyId { get; set; }

        /// <summary>
        /// The space ID associated with this project group.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        public ProjectGroupArgs()
        {
        }
        public static new ProjectGroupArgs Empty => new ProjectGroupArgs();
    }

    public sealed class ProjectGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this project group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the retention policy associated with this project group.
        /// </summary>
        [Input("retentionPolicyId")]
        public Input<string>? RetentionPolicyId { get; set; }

        /// <summary>
        /// The space ID associated with this project group.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        public ProjectGroupState()
        {
        }
        public static new ProjectGroupState Empty => new ProjectGroupState();
    }
}
