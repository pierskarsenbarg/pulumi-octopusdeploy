// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class TeamUserRoleArgs : global::Pulumi.ResourceArgs
    {
        [Input("environmentIds")]
        private InputList<string>? _environmentIds;
        public InputList<string> EnvironmentIds
        {
            get => _environmentIds ?? (_environmentIds = new InputList<string>());
            set => _environmentIds = value;
        }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("projectGroupIds")]
        private InputList<string>? _projectGroupIds;
        public InputList<string> ProjectGroupIds
        {
            get => _projectGroupIds ?? (_projectGroupIds = new InputList<string>());
            set => _projectGroupIds = value;
        }

        [Input("projectIds")]
        private InputList<string>? _projectIds;
        public InputList<string> ProjectIds
        {
            get => _projectIds ?? (_projectIds = new InputList<string>());
            set => _projectIds = value;
        }

        /// <summary>
        /// The space associated with this team.
        /// </summary>
        [Input("spaceId", required: true)]
        public Input<string> SpaceId { get; set; } = null!;

        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        [Input("tenantIds")]
        private InputList<string>? _tenantIds;
        public InputList<string> TenantIds
        {
            get => _tenantIds ?? (_tenantIds = new InputList<string>());
            set => _tenantIds = value;
        }

        [Input("userRoleId", required: true)]
        public Input<string> UserRoleId { get; set; } = null!;

        public TeamUserRoleArgs()
        {
        }
        public static new TeamUserRoleArgs Empty => new TeamUserRoleArgs();
    }
}
