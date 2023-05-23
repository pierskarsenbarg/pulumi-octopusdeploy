// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class ChannelRuleGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionPackages", required: true)]
        private InputList<Inputs.ChannelRuleActionPackageGetArgs>? _actionPackages;
        public InputList<Inputs.ChannelRuleActionPackageGetArgs> ActionPackages
        {
            get => _actionPackages ?? (_actionPackages = new InputList<Inputs.ChannelRuleActionPackageGetArgs>());
            set => _actionPackages = value;
        }

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("tag")]
        public Input<string>? Tag { get; set; }

        [Input("versionRange")]
        public Input<string>? VersionRange { get; set; }

        public ChannelRuleGetArgs()
        {
        }
        public static new ChannelRuleGetArgs Empty => new ChannelRuleGetArgs();
    }
}