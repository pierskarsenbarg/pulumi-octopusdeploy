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
    public sealed class LifecyclePhaseReleaseRetentionPolicy
    {
        /// <summary>
        /// The number of days/releases to keep. The default value is `30`. If `0` then all are kept.
        /// </summary>
        public readonly int? QuantityToKeep;
        /// <summary>
        /// Indicates if items should never be deleted. The default value is `false`.
        /// </summary>
        public readonly bool? ShouldKeepForever;
        /// <summary>
        /// The unit of quantity to keep. Valid units are `Days` or `Items`. The default value is `Days`.
        /// </summary>
        public readonly string? Unit;

        [OutputConstructor]
        private LifecyclePhaseReleaseRetentionPolicy(
            int? quantityToKeep,

            bool? shouldKeepForever,

            string? unit)
        {
            QuantityToKeep = quantityToKeep;
            ShouldKeepForever = shouldKeepForever;
            Unit = unit;
        }
    }
}
