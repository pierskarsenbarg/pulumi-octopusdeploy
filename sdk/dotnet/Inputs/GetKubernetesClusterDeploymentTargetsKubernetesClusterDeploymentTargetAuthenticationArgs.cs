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

    public sealed class GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetAuthenticationInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        public GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetAuthenticationInputArgs()
        {
        }
        public static new GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetAuthenticationInputArgs Empty => new GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetAuthenticationInputArgs();
    }
}
