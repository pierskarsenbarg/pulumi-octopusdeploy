// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetMachinePoliciesMachinePolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("connectionConnectTimeout", required: true)]
        public int ConnectionConnectTimeout { get; set; }

        [Input("connectionRetryCountLimit", required: true)]
        public int ConnectionRetryCountLimit { get; set; }

        [Input("connectionRetrySleepInterval", required: true)]
        public int ConnectionRetrySleepInterval { get; set; }

        [Input("connectionRetryTimeLimit", required: true)]
        public int ConnectionRetryTimeLimit { get; set; }

        /// <summary>
        /// The description of this machine policy.
        /// </summary>
        [Input("description", required: true)]
        public string Description { get; set; } = null!;

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("isDefault", required: true)]
        public bool IsDefault { get; set; }

        [Input("machineCleanupPolicies", required: true)]
        private List<Inputs.GetMachinePoliciesMachinePolicyMachineCleanupPolicyArgs>? _machineCleanupPolicies;
        public List<Inputs.GetMachinePoliciesMachinePolicyMachineCleanupPolicyArgs> MachineCleanupPolicies
        {
            get => _machineCleanupPolicies ?? (_machineCleanupPolicies = new List<Inputs.GetMachinePoliciesMachinePolicyMachineCleanupPolicyArgs>());
            set => _machineCleanupPolicies = value;
        }

        [Input("machineConnectivityPolicies", required: true)]
        private List<Inputs.GetMachinePoliciesMachinePolicyMachineConnectivityPolicyArgs>? _machineConnectivityPolicies;
        public List<Inputs.GetMachinePoliciesMachinePolicyMachineConnectivityPolicyArgs> MachineConnectivityPolicies
        {
            get => _machineConnectivityPolicies ?? (_machineConnectivityPolicies = new List<Inputs.GetMachinePoliciesMachinePolicyMachineConnectivityPolicyArgs>());
            set => _machineConnectivityPolicies = value;
        }

        [Input("machineHealthCheckPolicies", required: true)]
        private List<Inputs.GetMachinePoliciesMachinePolicyMachineHealthCheckPolicyArgs>? _machineHealthCheckPolicies;
        public List<Inputs.GetMachinePoliciesMachinePolicyMachineHealthCheckPolicyArgs> MachineHealthCheckPolicies
        {
            get => _machineHealthCheckPolicies ?? (_machineHealthCheckPolicies = new List<Inputs.GetMachinePoliciesMachinePolicyMachineHealthCheckPolicyArgs>());
            set => _machineHealthCheckPolicies = value;
        }

        [Input("machineUpdatePolicies", required: true)]
        private List<Inputs.GetMachinePoliciesMachinePolicyMachineUpdatePolicyArgs>? _machineUpdatePolicies;
        public List<Inputs.GetMachinePoliciesMachinePolicyMachineUpdatePolicyArgs> MachineUpdatePolicies
        {
            get => _machineUpdatePolicies ?? (_machineUpdatePolicies = new List<Inputs.GetMachinePoliciesMachinePolicyMachineUpdatePolicyArgs>());
            set => _machineUpdatePolicies = value;
        }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("pollingRequestMaximumMessageProcessingTimeout", required: true)]
        public int PollingRequestMaximumMessageProcessingTimeout { get; set; }

        [Input("pollingRequestQueueTimeout", required: true)]
        public int PollingRequestQueueTimeout { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId", required: true)]
        public string SpaceId { get; set; } = null!;

        public GetMachinePoliciesMachinePolicyArgs()
        {
        }
        public static new GetMachinePoliciesMachinePolicyArgs Empty => new GetMachinePoliciesMachinePolicyArgs();
    }
}
