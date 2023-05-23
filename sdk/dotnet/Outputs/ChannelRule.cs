// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class ChannelRule
    {
        public readonly ImmutableArray<Outputs.ChannelRuleActionPackage> ActionPackages;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string? Id;
        public readonly string? Tag;
        public readonly string? VersionRange;

        [OutputConstructor]
        private ChannelRule(
            ImmutableArray<Outputs.ChannelRuleActionPackage> actionPackages,

            string? id,

            string? tag,

            string? versionRange)
        {
            ActionPackages = actionPackages;
            Id = id;
            Tag = tag;
            VersionRange = versionRange;
        }
    }
}
