// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetEndpointContainerInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("feedId")]
        public Input<string>? FeedId { get; set; }

        [Input("image")]
        public Input<string>? Image { get; set; }

        public GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetEndpointContainerInputArgs()
        {
        }
        public static new GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetEndpointContainerInputArgs Empty => new GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetEndpointContainerInputArgs();
    }
}
