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

    public sealed class KubernetesClusterDeploymentTargetAwsAccountAuthenticationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        [Input("assumeRole")]
        public Input<bool>? AssumeRole { get; set; }

        [Input("assumeRoleExternalId")]
        public Input<string>? AssumeRoleExternalId { get; set; }

        [Input("assumeRoleSessionDuration")]
        public Input<int>? AssumeRoleSessionDuration { get; set; }

        [Input("assumedRoleArn")]
        public Input<string>? AssumedRoleArn { get; set; }

        [Input("assumedRoleSession")]
        public Input<string>? AssumedRoleSession { get; set; }

        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        [Input("useInstanceRole")]
        public Input<bool>? UseInstanceRole { get; set; }

        public KubernetesClusterDeploymentTargetAwsAccountAuthenticationGetArgs()
        {
        }
        public static new KubernetesClusterDeploymentTargetAwsAccountAuthenticationGetArgs Empty => new KubernetesClusterDeploymentTargetAwsAccountAuthenticationGetArgs();
    }
}
