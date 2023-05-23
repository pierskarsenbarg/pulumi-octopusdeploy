// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetAzureWebAppDeploymentTargetsAzureWebAppDeploymentTargetEndpointAuthenticationArgs : global::Pulumi.InvokeArgs
    {
        [Input("accountId")]
        public string? AccountId { get; set; }

        [Input("adminLogin")]
        public string? AdminLogin { get; set; }

        [Input("assumeRole")]
        public bool? AssumeRole { get; set; }

        [Input("assumeRoleExternalId")]
        public string? AssumeRoleExternalId { get; set; }

        [Input("assumeRoleSessionDuration")]
        public int? AssumeRoleSessionDuration { get; set; }

        [Input("assumedRoleArn")]
        public string? AssumedRoleArn { get; set; }

        [Input("assumedRoleSession")]
        public string? AssumedRoleSession { get; set; }

        [Input("authenticationType")]
        public string? AuthenticationType { get; set; }

        [Input("clientCertificate")]
        public string? ClientCertificate { get; set; }

        [Input("clusterName")]
        public string? ClusterName { get; set; }

        [Input("clusterResourceGroup")]
        public string? ClusterResourceGroup { get; set; }

        [Input("impersonateServiceAccount")]
        public bool? ImpersonateServiceAccount { get; set; }

        [Input("project")]
        public string? Project { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("serviceAccountEmails")]
        public string? ServiceAccountEmails { get; set; }

        [Input("useInstanceRole")]
        public bool? UseInstanceRole { get; set; }

        [Input("useVmServiceAccount")]
        public bool? UseVmServiceAccount { get; set; }

        [Input("zone")]
        public string? Zone { get; set; }

        public GetAzureWebAppDeploymentTargetsAzureWebAppDeploymentTargetEndpointAuthenticationArgs()
        {
        }
        public static new GetAzureWebAppDeploymentTargetsAzureWebAppDeploymentTargetEndpointAuthenticationArgs Empty => new GetAzureWebAppDeploymentTargetsAzureWebAppDeploymentTargetEndpointAuthenticationArgs();
    }
}
