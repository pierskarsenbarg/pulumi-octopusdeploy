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

    public sealed class MachinePolicyMachineConnectivityPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("machineConnectivityBehavior")]
        public Input<string>? MachineConnectivityBehavior { get; set; }

        public MachinePolicyMachineConnectivityPolicyGetArgs()
        {
        }
        public static new MachinePolicyMachineConnectivityPolicyGetArgs Empty => new MachinePolicyMachineConnectivityPolicyGetArgs();
    }
}
