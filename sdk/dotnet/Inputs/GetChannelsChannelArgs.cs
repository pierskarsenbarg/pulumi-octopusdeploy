// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetChannelsChannelInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this channel.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Indicates if this is the default channel for the associated project.
        /// </summary>
        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        /// <summary>
        /// The lifecycle ID associated with this channel.
        /// </summary>
        [Input("lifecycleId", required: true)]
        public Input<string> LifecycleId { get; set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The project ID associated with this channel.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("rules", required: true)]
        private InputList<Inputs.GetChannelsChannelRuleInputArgs>? _rules;

        /// <summary>
        /// A list of rules associated with this channel.
        /// </summary>
        public InputList<Inputs.GetChannelsChannelRuleInputArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.GetChannelsChannelRuleInputArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId", required: true)]
        public Input<string> SpaceId { get; set; } = null!;

        [Input("tenantTags", required: true)]
        private InputList<string>? _tenantTags;

        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public InputList<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new InputList<string>());
            set => _tenantTags = value;
        }

        public GetChannelsChannelInputArgs()
        {
        }
        public static new GetChannelsChannelInputArgs Empty => new GetChannelsChannelInputArgs();
    }
}