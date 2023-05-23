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

    public sealed class GetChannelsChannelRuleInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionPackages", required: true)]
        private InputList<Inputs.GetChannelsChannelRuleActionPackageInputArgs>? _actionPackages;
        public InputList<Inputs.GetChannelsChannelRuleActionPackageInputArgs> ActionPackages
        {
            get => _actionPackages ?? (_actionPackages = new InputList<Inputs.GetChannelsChannelRuleActionPackageInputArgs>());
            set => _actionPackages = value;
        }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("tag")]
        public Input<string>? Tag { get; set; }

        [Input("versionRange")]
        public Input<string>? VersionRange { get; set; }

        public GetChannelsChannelRuleInputArgs()
        {
        }
        public static new GetChannelsChannelRuleInputArgs Empty => new GetChannelsChannelRuleInputArgs();
    }
}
