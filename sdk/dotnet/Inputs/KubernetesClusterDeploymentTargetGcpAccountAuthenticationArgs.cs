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

    public sealed class KubernetesClusterDeploymentTargetGcpAccountAuthenticationArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        [Input("impersonateServiceAccount")]
        public Input<bool>? ImpersonateServiceAccount { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("serviceAccountEmails")]
        public Input<string>? ServiceAccountEmails { get; set; }

        [Input("useVmServiceAccount")]
        public Input<bool>? UseVmServiceAccount { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public KubernetesClusterDeploymentTargetGcpAccountAuthenticationArgs()
        {
        }
        public static new KubernetesClusterDeploymentTargetGcpAccountAuthenticationArgs Empty => new KubernetesClusterDeploymentTargetGcpAccountAuthenticationArgs();
    }
}
