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
    public sealed class GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetEndpointContainerResult
    {
        public readonly string? FeedId;
        public readonly string? Image;

        [OutputConstructor]
        private GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetEndpointContainerResult(
            string? feedId,

            string? image)
        {
            FeedId = feedId;
            Image = image;
        }
    }
}
