// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetSpacesSpaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The description of this space.
        /// </summary>
        [Input("description", required: true)]
        public string Description { get; set; } = null!;

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// Specifies if this space is the default space in Octopus.
        /// </summary>
        [Input("isDefault", required: true)]
        public bool IsDefault { get; set; }

        /// <summary>
        /// Specifies the status of the task queue for this space.
        /// </summary>
        [Input("isTaskQueueStopped", required: true)]
        public bool IsTaskQueueStopped { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The unique slug of this space.
        /// </summary>
        [Input("slug", required: true)]
        public string Slug { get; set; } = null!;

        [Input("spaceManagersTeamMembers", required: true)]
        private List<string>? _spaceManagersTeamMembers;

        /// <summary>
        /// A list of user IDs designated to be managers of this space.
        /// </summary>
        public List<string> SpaceManagersTeamMembers
        {
            get => _spaceManagersTeamMembers ?? (_spaceManagersTeamMembers = new List<string>());
            set => _spaceManagersTeamMembers = value;
        }

        [Input("spaceManagersTeams", required: true)]
        private List<string>? _spaceManagersTeams;

        /// <summary>
        /// A list of team IDs designated to be managers of this space.
        /// </summary>
        public List<string> SpaceManagersTeams
        {
            get => _spaceManagersTeams ?? (_spaceManagersTeams = new List<string>());
            set => _spaceManagersTeams = value;
        }

        public GetSpacesSpaceArgs()
        {
        }
        public static new GetSpacesSpaceArgs Empty => new GetSpacesSpaceArgs();
    }
}
