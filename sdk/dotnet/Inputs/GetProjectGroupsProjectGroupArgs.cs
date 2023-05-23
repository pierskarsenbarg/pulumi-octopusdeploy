// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetProjectGroupsProjectGroupInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this project group.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the retention policy associated with this project group.
        /// </summary>
        [Input("retentionPolicyId", required: true)]
        public Input<string> RetentionPolicyId { get; set; } = null!;

        /// <summary>
        /// The space ID associated with this project group.
        /// </summary>
        [Input("spaceId", required: true)]
        public Input<string> SpaceId { get; set; } = null!;

        public GetProjectGroupsProjectGroupInputArgs()
        {
        }
        public static new GetProjectGroupsProjectGroupInputArgs Empty => new GetProjectGroupsProjectGroupInputArgs();
    }
}