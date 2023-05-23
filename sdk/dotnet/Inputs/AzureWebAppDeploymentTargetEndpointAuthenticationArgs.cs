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

    public sealed class AzureWebAppDeploymentTargetEndpointAuthenticationArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        [Input("adminLogin")]
        public Input<string>? AdminLogin { get; set; }

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

        [Input("authenticationType")]
        public Input<string>? AuthenticationType { get; set; }

        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        [Input("clusterResourceGroup")]
        public Input<string>? ClusterResourceGroup { get; set; }

        [Input("impersonateServiceAccount")]
        public Input<bool>? ImpersonateServiceAccount { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("serviceAccountEmails")]
        public Input<string>? ServiceAccountEmails { get; set; }

        [Input("useInstanceRole")]
        public Input<bool>? UseInstanceRole { get; set; }

        [Input("useVmServiceAccount")]
        public Input<bool>? UseVmServiceAccount { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AzureWebAppDeploymentTargetEndpointAuthenticationArgs()
        {
        }
        public static new AzureWebAppDeploymentTargetEndpointAuthenticationArgs Empty => new AzureWebAppDeploymentTargetEndpointAuthenticationArgs();
    }
}
