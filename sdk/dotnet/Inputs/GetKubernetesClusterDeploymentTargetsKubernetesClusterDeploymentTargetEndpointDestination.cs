// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetEndpointDestinationArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationType")]
        public string? DestinationType { get; set; }

        [Input("dropFolderPath")]
        public string? DropFolderPath { get; set; }

        public GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetEndpointDestinationArgs()
        {
        }
        public static new GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetEndpointDestinationArgs Empty => new GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetEndpointDestinationArgs();
    }
}
