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

    public sealed class GetMachinePoliciesMachinePolicyMachineHealthCheckPolicyBashHealthCheckPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("runType")]
        public string? RunType { get; set; }

        [Input("scriptBody")]
        public string? ScriptBody { get; set; }

        public GetMachinePoliciesMachinePolicyMachineHealthCheckPolicyBashHealthCheckPolicyArgs()
        {
        }
        public static new GetMachinePoliciesMachinePolicyMachineHealthCheckPolicyBashHealthCheckPolicyArgs Empty => new GetMachinePoliciesMachinePolicyMachineHealthCheckPolicyBashHealthCheckPolicyArgs();
    }
}
