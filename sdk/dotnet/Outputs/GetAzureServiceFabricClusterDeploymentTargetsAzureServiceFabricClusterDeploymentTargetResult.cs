// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetResult
    {
        public readonly string AadClientCredentialSecret;
        public readonly string AadCredentialType;
        public readonly string AadUserCredentialPassword;
        public readonly string AadUserCredentialUsername;
        public readonly string CertificateStoreLocation;
        public readonly string CertificateStoreName;
        public readonly string ClientCertificateVariable;
        public readonly string ConnectionEndpoint;
        public readonly ImmutableArray<Outputs.GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetEndpointResult> Endpoints;
        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        public readonly ImmutableArray<string> Environments;
        public readonly bool HasLatestCalamari;
        /// <summary>
        /// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        public readonly string HealthStatus;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A filter to search by the disabled status of a resource.
        /// </summary>
        public readonly bool IsDisabled;
        public readonly bool IsInProcess;
        public readonly string MachinePolicyId;
        /// <summary>
        /// The name of this resource.
        /// </summary>
        public readonly string Name;
        public readonly string OperatingSystem;
        /// <summary>
        /// A filter to search by a list of role IDs.
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        public readonly string SecurityMode;
        public readonly string ServerCertificateThumbprint;
        public readonly string ShellName;
        public readonly string ShellVersion;
        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        public readonly string SpaceId;
        /// <summary>
        /// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A summary elaborating on the status of this resource.
        /// </summary>
        public readonly string StatusSummary;
        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public readonly ImmutableArray<string> TenantTags;
        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        public readonly string TenantedDeploymentParticipation;
        /// <summary>
        /// A list of tenant IDs associated with this resource.
        /// </summary>
        public readonly ImmutableArray<string> Tenants;
        /// <summary>
        /// The thumbprint of the deployment target to match in the query and/or search
        /// </summary>
        public readonly string Thumbprint;
        public readonly string Uri;

        [OutputConstructor]
        private GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetResult(
            string aadClientCredentialSecret,

            string aadCredentialType,

            string aadUserCredentialPassword,

            string aadUserCredentialUsername,

            string certificateStoreLocation,

            string certificateStoreName,

            string clientCertificateVariable,

            string connectionEndpoint,

            ImmutableArray<Outputs.GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetEndpointResult> endpoints,

            ImmutableArray<string> environments,

            bool hasLatestCalamari,

            string healthStatus,

            string id,

            bool isDisabled,

            bool isInProcess,

            string machinePolicyId,

            string name,

            string operatingSystem,

            ImmutableArray<string> roles,

            string securityMode,

            string serverCertificateThumbprint,

            string shellName,

            string shellVersion,

            string spaceId,

            string status,

            string statusSummary,

            ImmutableArray<string> tenantTags,

            string tenantedDeploymentParticipation,

            ImmutableArray<string> tenants,

            string thumbprint,

            string uri)
        {
            AadClientCredentialSecret = aadClientCredentialSecret;
            AadCredentialType = aadCredentialType;
            AadUserCredentialPassword = aadUserCredentialPassword;
            AadUserCredentialUsername = aadUserCredentialUsername;
            CertificateStoreLocation = certificateStoreLocation;
            CertificateStoreName = certificateStoreName;
            ClientCertificateVariable = clientCertificateVariable;
            ConnectionEndpoint = connectionEndpoint;
            Endpoints = endpoints;
            Environments = environments;
            HasLatestCalamari = hasLatestCalamari;
            HealthStatus = healthStatus;
            Id = id;
            IsDisabled = isDisabled;
            IsInProcess = isInProcess;
            MachinePolicyId = machinePolicyId;
            Name = name;
            OperatingSystem = operatingSystem;
            Roles = roles;
            SecurityMode = securityMode;
            ServerCertificateThumbprint = serverCertificateThumbprint;
            ShellName = shellName;
            ShellVersion = shellVersion;
            SpaceId = spaceId;
            Status = status;
            StatusSummary = statusSummary;
            TenantTags = tenantTags;
            TenantedDeploymentParticipation = tenantedDeploymentParticipation;
            Tenants = tenants;
            Thumbprint = thumbprint;
            Uri = uri;
        }
    }
}
