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
    /// This resource manages tag sets in Octopus Deploy.
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
    ///     var example = new Octopusdeploy.TagSet("example", new()
    ///     {
    ///         Description = "Provides tenants with access to certain early access programs.",
    ///     });
    /// 
    ///     // tags are distinct resources and associated with tag sets through tag_set_id
    ///     var alpha = new Octopusdeploy.Tag("alpha", new()
    ///     {
    ///         Color = "#00FF00",
    ///         TagSetId = example.Id,
    ///     });
    /// 
    ///     var beta = new Octopusdeploy.Tag("beta", new()
    ///     {
    ///         Color = "#FF0000",
    ///         TagSetId = example.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import octopusdeploy:index/tagSet:TagSet [options] octopusdeploy_tag_set.&lt;name&gt; &lt;tag-set-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/tagSet:TagSet")]
    public partial class TagSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of this tag set.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The sort order associated with this resource.
        /// </summary>
        [Output("sortOrder")]
        public Output<int> SortOrder { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;


        /// <summary>
        /// Create a TagSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TagSet(string name, TagSetArgs? args = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/tagSet:TagSet", name, args ?? new TagSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TagSet(string name, Input<string> id, TagSetState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/tagSet:TagSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TagSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TagSet Get(string name, Input<string> id, TagSetState? state = null, CustomResourceOptions? options = null)
        {
            return new TagSet(name, id, state, options);
        }
    }

    public sealed class TagSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this tag set.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The sort order associated with this resource.
        /// </summary>
        [Input("sortOrder")]
        public Input<int>? SortOrder { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        public TagSetArgs()
        {
        }
        public static new TagSetArgs Empty => new TagSetArgs();
    }

    public sealed class TagSetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this tag set.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The sort order associated with this resource.
        /// </summary>
        [Input("sortOrder")]
        public Input<int>? SortOrder { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        public TagSetState()
        {
        }
        public static new TagSetState Empty => new TagSetState();
    }
}
