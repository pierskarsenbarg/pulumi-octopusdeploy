// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetAuthenticationArgs : global::Pulumi.InvokeArgs
    {
        [Input("accountId")]
        public string? AccountId { get; set; }

        public GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetAuthenticationArgs()
        {
        }
        public static new GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetAuthenticationArgs Empty => new GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetAuthenticationArgs();
    }
}
