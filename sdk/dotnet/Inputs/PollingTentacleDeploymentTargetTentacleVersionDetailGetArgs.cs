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

    public sealed class PollingTentacleDeploymentTargetTentacleVersionDetailGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("upgradeLocked")]
        public Input<bool>? UpgradeLocked { get; set; }

        [Input("upgradeRequired")]
        public Input<bool>? UpgradeRequired { get; set; }

        [Input("upgradeSuggested")]
        public Input<bool>? UpgradeSuggested { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public PollingTentacleDeploymentTargetTentacleVersionDetailGetArgs()
        {
        }
        public static new PollingTentacleDeploymentTargetTentacleVersionDetailGetArgs Empty => new PollingTentacleDeploymentTargetTentacleVersionDetailGetArgs();
    }
}
