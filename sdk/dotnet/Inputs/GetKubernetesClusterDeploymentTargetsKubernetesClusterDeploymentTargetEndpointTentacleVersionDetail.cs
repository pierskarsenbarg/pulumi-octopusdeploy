// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetEndpointTentacleVersionDetailArgs : global::Pulumi.InvokeArgs
    {
        [Input("upgradeLocked", required: true)]
        public bool UpgradeLocked { get; set; }

        [Input("upgradeRequired", required: true)]
        public bool UpgradeRequired { get; set; }

        [Input("upgradeSuggested", required: true)]
        public bool UpgradeSuggested { get; set; }

        [Input("version", required: true)]
        public string Version { get; set; } = null!;

        public GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetEndpointTentacleVersionDetailArgs()
        {
        }
        public static new GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetEndpointTentacleVersionDetailArgs Empty => new GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetEndpointTentacleVersionDetailArgs();
    }
}