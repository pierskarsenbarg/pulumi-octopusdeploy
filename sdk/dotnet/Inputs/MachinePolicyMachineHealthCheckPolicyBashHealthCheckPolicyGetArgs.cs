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

    public sealed class MachinePolicyMachineHealthCheckPolicyBashHealthCheckPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("runType")]
        public Input<string>? RunType { get; set; }

        [Input("scriptBody")]
        public Input<string>? ScriptBody { get; set; }

        public MachinePolicyMachineHealthCheckPolicyBashHealthCheckPolicyGetArgs()
        {
        }
        public static new MachinePolicyMachineHealthCheckPolicyBashHealthCheckPolicyGetArgs Empty => new MachinePolicyMachineHealthCheckPolicyBashHealthCheckPolicyGetArgs();
    }
}
