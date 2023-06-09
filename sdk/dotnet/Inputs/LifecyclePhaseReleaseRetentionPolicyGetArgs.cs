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

    public sealed class LifecyclePhaseReleaseRetentionPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of days/releases to keep. The default value is `30`. If `0` then all are kept.
        /// </summary>
        [Input("quantityToKeep")]
        public Input<int>? QuantityToKeep { get; set; }

        /// <summary>
        /// Indicates if items should never be deleted. The default value is `false`.
        /// </summary>
        [Input("shouldKeepForever")]
        public Input<bool>? ShouldKeepForever { get; set; }

        /// <summary>
        /// The unit of quantity to keep. Valid units are `Days` or `Items`. The default value is `Days`.
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public LifecyclePhaseReleaseRetentionPolicyGetArgs()
        {
        }
        public static new LifecyclePhaseReleaseRetentionPolicyGetArgs Empty => new LifecyclePhaseReleaseRetentionPolicyGetArgs();
    }
}
