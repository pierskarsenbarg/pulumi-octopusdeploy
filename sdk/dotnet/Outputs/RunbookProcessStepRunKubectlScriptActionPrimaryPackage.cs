// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class RunbookProcessStepRunKubectlScriptActionPrimaryPackage
    {
        public readonly string? AcquisitionLocation;
        public readonly string? FeedId;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string? Id;
        public readonly string? Name;
        public readonly string PackageId;
        public readonly ImmutableDictionary<string, string>? Properties;

        [OutputConstructor]
        private RunbookProcessStepRunKubectlScriptActionPrimaryPackage(
            string? acquisitionLocation,

            string? feedId,

            string? id,

            string? name,

            string packageId,

            ImmutableDictionary<string, string>? properties)
        {
            AcquisitionLocation = acquisitionLocation;
            FeedId = feedId;
            Id = id;
            Name = name;
            PackageId = packageId;
            Properties = properties;
        }
    }
}
