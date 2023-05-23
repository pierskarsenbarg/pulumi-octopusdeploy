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
    public sealed class AzureWebAppDeploymentTargetEndpointAuthentication
    {
        public readonly string? AccountId;
        public readonly string? AdminLogin;
        public readonly bool? AssumeRole;
        public readonly string? AssumeRoleExternalId;
        public readonly int? AssumeRoleSessionDuration;
        public readonly string? AssumedRoleArn;
        public readonly string? AssumedRoleSession;
        public readonly string? AuthenticationType;
        public readonly string? ClientCertificate;
        public readonly string? ClusterName;
        public readonly string? ClusterResourceGroup;
        public readonly bool? ImpersonateServiceAccount;
        public readonly string? Project;
        public readonly string? Region;
        public readonly string? ServiceAccountEmails;
        public readonly bool? UseInstanceRole;
        public readonly bool? UseVmServiceAccount;
        public readonly string? Zone;

        [OutputConstructor]
        private AzureWebAppDeploymentTargetEndpointAuthentication(
            string? accountId,

            string? adminLogin,

            bool? assumeRole,

            string? assumeRoleExternalId,

            int? assumeRoleSessionDuration,

            string? assumedRoleArn,

            string? assumedRoleSession,

            string? authenticationType,

            string? clientCertificate,

            string? clusterName,

            string? clusterResourceGroup,

            bool? impersonateServiceAccount,

            string? project,

            string? region,

            string? serviceAccountEmails,

            bool? useInstanceRole,

            bool? useVmServiceAccount,

            string? zone)
        {
            AccountId = accountId;
            AdminLogin = adminLogin;
            AssumeRole = assumeRole;
            AssumeRoleExternalId = assumeRoleExternalId;
            AssumeRoleSessionDuration = assumeRoleSessionDuration;
            AssumedRoleArn = assumedRoleArn;
            AssumedRoleSession = assumedRoleSession;
            AuthenticationType = authenticationType;
            ClientCertificate = clientCertificate;
            ClusterName = clusterName;
            ClusterResourceGroup = clusterResourceGroup;
            ImpersonateServiceAccount = impersonateServiceAccount;
            Project = project;
            Region = region;
            ServiceAccountEmails = serviceAccountEmails;
            UseInstanceRole = useInstanceRole;
            UseVmServiceAccount = useVmServiceAccount;
            Zone = zone;
        }
    }
}
