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
    public sealed class PollingTentacleDeploymentTargetTentacleVersionDetail
    {
        public readonly bool? UpgradeLocked;
        public readonly bool? UpgradeRequired;
        public readonly bool? UpgradeSuggested;
        public readonly string? Version;

        [OutputConstructor]
        private PollingTentacleDeploymentTargetTentacleVersionDetail(
            bool? upgradeLocked,

            bool? upgradeRequired,

            bool? upgradeSuggested,

            string? version)
        {
            UpgradeLocked = upgradeLocked;
            UpgradeRequired = upgradeRequired;
            UpgradeSuggested = upgradeSuggested;
            Version = version;
        }
    }
}
