// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    public static class GetSpace
    {
        /// <summary>
        /// Provides information about an existing space.
        /// </summary>
        public static Task<GetSpaceResult> InvokeAsync(GetSpaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSpaceResult>("octopusdeploy:index/getSpace:getSpace", args ?? new GetSpaceArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about an existing space.
        /// </summary>
        public static Output<GetSpaceResult> Invoke(GetSpaceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSpaceResult>("octopusdeploy:index/getSpace:getSpace", args ?? new GetSpaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSpaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetSpaceArgs()
        {
        }
        public static new GetSpaceArgs Empty => new GetSpaceArgs();
    }

    public sealed class GetSpaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetSpaceInvokeArgs()
        {
        }
        public static new GetSpaceInvokeArgs Empty => new GetSpaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetSpaceResult
    {
        /// <summary>
        /// The description of this space.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies if this space is the default space in Octopus.
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// Specifies the status of the task queue for this space.
        /// </summary>
        public readonly bool IsTaskQueueStopped;
        /// <summary>
        /// The name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The unique slug of this space.
        /// </summary>
        public readonly string Slug;
        /// <summary>
        /// A list of user IDs designated to be managers of this space.
        /// </summary>
        public readonly ImmutableArray<string> SpaceManagersTeamMembers;
        /// <summary>
        /// A list of team IDs designated to be managers of this space.
        /// </summary>
        public readonly ImmutableArray<string> SpaceManagersTeams;

        [OutputConstructor]
        private GetSpaceResult(
            string description,

            string id,

            bool isDefault,

            bool isTaskQueueStopped,

            string name,

            string slug,

            ImmutableArray<string> spaceManagersTeamMembers,

            ImmutableArray<string> spaceManagersTeams)
        {
            Description = description;
            Id = id;
            IsDefault = isDefault;
            IsTaskQueueStopped = isTaskQueueStopped;
            Name = name;
            Slug = slug;
            SpaceManagersTeamMembers = spaceManagersTeamMembers;
            SpaceManagersTeams = spaceManagersTeams;
        }
    }
}
