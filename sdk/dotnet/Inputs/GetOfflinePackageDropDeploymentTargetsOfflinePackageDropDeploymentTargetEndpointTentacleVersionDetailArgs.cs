// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetOfflinePackageDropDeploymentTargetsOfflinePackageDropDeploymentTargetEndpointTentacleVersionDetailInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("upgradeLocked", required: true)]
        public Input<bool> UpgradeLocked { get; set; } = null!;

        [Input("upgradeRequired", required: true)]
        public Input<bool> UpgradeRequired { get; set; } = null!;

        [Input("upgradeSuggested", required: true)]
        public Input<bool> UpgradeSuggested { get; set; } = null!;

        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public GetOfflinePackageDropDeploymentTargetsOfflinePackageDropDeploymentTargetEndpointTentacleVersionDetailInputArgs()
        {
        }
        public static new GetOfflinePackageDropDeploymentTargetsOfflinePackageDropDeploymentTargetEndpointTentacleVersionDetailInputArgs Empty => new GetOfflinePackageDropDeploymentTargetsOfflinePackageDropDeploymentTargetEndpointTentacleVersionDetailInputArgs();
    }
}
