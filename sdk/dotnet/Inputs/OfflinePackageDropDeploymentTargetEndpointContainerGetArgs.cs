// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class OfflinePackageDropDeploymentTargetEndpointContainerGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("feedId")]
        public Input<string>? FeedId { get; set; }

        [Input("image")]
        public Input<string>? Image { get; set; }

        public OfflinePackageDropDeploymentTargetEndpointContainerGetArgs()
        {
        }
        public static new OfflinePackageDropDeploymentTargetEndpointContainerGetArgs Empty => new OfflinePackageDropDeploymentTargetEndpointContainerGetArgs();
    }
}
