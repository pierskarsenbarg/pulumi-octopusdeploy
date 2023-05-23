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

    public sealed class ProjectConnectivityPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowDeploymentsToNoTargets")]
        public Input<bool>? AllowDeploymentsToNoTargets { get; set; }

        [Input("excludeUnhealthyTargets")]
        public Input<bool>? ExcludeUnhealthyTargets { get; set; }

        [Input("skipMachineBehavior")]
        public Input<string>? SkipMachineBehavior { get; set; }

        [Input("targetRoles")]
        private InputList<string>? _targetRoles;
        public InputList<string> TargetRoles
        {
            get => _targetRoles ?? (_targetRoles = new InputList<string>());
            set => _targetRoles = value;
        }

        public ProjectConnectivityPolicyGetArgs()
        {
        }
        public static new ProjectConnectivityPolicyGetArgs Empty => new ProjectConnectivityPolicyGetArgs();
    }
}
