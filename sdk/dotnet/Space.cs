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
    /// This resource manages spaces in Octopus Deploy.
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
    ///     var example = new Octopusdeploy.Space("example", new()
    ///     {
    ///         Description = "A space for the development team.",
    ///         IsDefault = false,
    ///         IsTaskQueueStopped = false,
    ///         SpaceManagersTeamMembers = new[]
    ///         {
    ///             "Users-123",
    ///             "Users-321",
    ///         },
    ///         SpaceManagersTeams = new[]
    ///         {
    ///             "teams-everyone",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import octopusdeploy:index/space:Space [options] octopusdeploy_space.&lt;name&gt; &lt;space-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/space:Space")]
    public partial class Space : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of this space.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies if this space is the default space in Octopus.
        /// </summary>
        [Output("isDefault")]
        public Output<bool?> IsDefault { get; private set; } = null!;

        /// <summary>
        /// Specifies the status of the task queue for this space.
        /// </summary>
        [Output("isTaskQueueStopped")]
        public Output<bool?> IsTaskQueueStopped { get; private set; } = null!;

        /// <summary>
        /// The name of this resource, no more than 20 characters long
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The unique slug of this space.
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        /// <summary>
        /// A list of user IDs designated to be managers of this space.
        /// </summary>
        [Output("spaceManagersTeamMembers")]
        public Output<ImmutableArray<string>> SpaceManagersTeamMembers { get; private set; } = null!;

        /// <summary>
        /// A list of team IDs designated to be managers of this space.
        /// </summary>
        [Output("spaceManagersTeams")]
        public Output<ImmutableArray<string>> SpaceManagersTeams { get; private set; } = null!;


        /// <summary>
        /// Create a Space resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Space(string name, SpaceArgs? args = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/space:Space", name, args ?? new SpaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Space(string name, Input<string> id, SpaceState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/space:Space", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Space resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Space Get(string name, Input<string> id, SpaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Space(name, id, state, options);
        }
    }

    public sealed class SpaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this space.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies if this space is the default space in Octopus.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// Specifies the status of the task queue for this space.
        /// </summary>
        [Input("isTaskQueueStopped")]
        public Input<bool>? IsTaskQueueStopped { get; set; }

        /// <summary>
        /// The name of this resource, no more than 20 characters long
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique slug of this space.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        [Input("spaceManagersTeamMembers")]
        private InputList<string>? _spaceManagersTeamMembers;

        /// <summary>
        /// A list of user IDs designated to be managers of this space.
        /// </summary>
        public InputList<string> SpaceManagersTeamMembers
        {
            get => _spaceManagersTeamMembers ?? (_spaceManagersTeamMembers = new InputList<string>());
            set => _spaceManagersTeamMembers = value;
        }

        [Input("spaceManagersTeams")]
        private InputList<string>? _spaceManagersTeams;

        /// <summary>
        /// A list of team IDs designated to be managers of this space.
        /// </summary>
        public InputList<string> SpaceManagersTeams
        {
            get => _spaceManagersTeams ?? (_spaceManagersTeams = new InputList<string>());
            set => _spaceManagersTeams = value;
        }

        public SpaceArgs()
        {
        }
        public static new SpaceArgs Empty => new SpaceArgs();
    }

    public sealed class SpaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this space.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies if this space is the default space in Octopus.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// Specifies the status of the task queue for this space.
        /// </summary>
        [Input("isTaskQueueStopped")]
        public Input<bool>? IsTaskQueueStopped { get; set; }

        /// <summary>
        /// The name of this resource, no more than 20 characters long
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique slug of this space.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        [Input("spaceManagersTeamMembers")]
        private InputList<string>? _spaceManagersTeamMembers;

        /// <summary>
        /// A list of user IDs designated to be managers of this space.
        /// </summary>
        public InputList<string> SpaceManagersTeamMembers
        {
            get => _spaceManagersTeamMembers ?? (_spaceManagersTeamMembers = new InputList<string>());
            set => _spaceManagersTeamMembers = value;
        }

        [Input("spaceManagersTeams")]
        private InputList<string>? _spaceManagersTeams;

        /// <summary>
        /// A list of team IDs designated to be managers of this space.
        /// </summary>
        public InputList<string> SpaceManagersTeams
        {
            get => _spaceManagersTeams ?? (_spaceManagersTeams = new InputList<string>());
            set => _spaceManagersTeams = value;
        }

        public SpaceState()
        {
        }
        public static new SpaceState Empty => new SpaceState();
    }
}
