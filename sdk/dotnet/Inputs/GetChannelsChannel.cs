// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy.Inputs
{

    public sealed class GetChannelsChannelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The description of this channel.
        /// </summary>
        [Input("description", required: true)]
        public string Description { get; set; } = null!;

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// Indicates if this is the default channel for the associated project.
        /// </summary>
        [Input("isDefault", required: true)]
        public bool IsDefault { get; set; }

        /// <summary>
        /// The lifecycle ID associated with this channel.
        /// </summary>
        [Input("lifecycleId", required: true)]
        public string LifecycleId { get; set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project ID associated with this channel.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        [Input("rules", required: true)]
        private List<Inputs.GetChannelsChannelRuleArgs>? _rules;

        /// <summary>
        /// A list of rules associated with this channel.
        /// </summary>
        public List<Inputs.GetChannelsChannelRuleArgs> Rules
        {
            get => _rules ?? (_rules = new List<Inputs.GetChannelsChannelRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId", required: true)]
        public string SpaceId { get; set; } = null!;

        [Input("tenantTags", required: true)]
        private List<string>? _tenantTags;

        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public List<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new List<string>());
            set => _tenantTags = value;
        }

        public GetChannelsChannelArgs()
        {
        }
        public static new GetChannelsChannelArgs Empty => new GetChannelsChannelArgs();
    }
}
