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
    public sealed class MachinePolicyMachineHealthCheckPolicy
    {
        public readonly Outputs.MachinePolicyMachineHealthCheckPolicyBashHealthCheckPolicy BashHealthCheckPolicy;
        public readonly string? HealthCheckCron;
        public readonly string? HealthCheckCronTimezone;
        public readonly int? HealthCheckInterval;
        public readonly string? HealthCheckType;
        public readonly Outputs.MachinePolicyMachineHealthCheckPolicyPowershellHealthCheckPolicy PowershellHealthCheckPolicy;

        [OutputConstructor]
        private MachinePolicyMachineHealthCheckPolicy(
            Outputs.MachinePolicyMachineHealthCheckPolicyBashHealthCheckPolicy bashHealthCheckPolicy,

            string? healthCheckCron,

            string? healthCheckCronTimezone,

            int? healthCheckInterval,

            string? healthCheckType,

            Outputs.MachinePolicyMachineHealthCheckPolicyPowershellHealthCheckPolicy powershellHealthCheckPolicy)
        {
            BashHealthCheckPolicy = bashHealthCheckPolicy;
            HealthCheckCron = healthCheckCron;
            HealthCheckCronTimezone = healthCheckCronTimezone;
            HealthCheckInterval = healthCheckInterval;
            HealthCheckType = healthCheckType;
            PowershellHealthCheckPolicy = powershellHealthCheckPolicy;
        }
    }
}
