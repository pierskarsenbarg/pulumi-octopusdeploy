// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetArgs : global::Pulumi.InvokeArgs
    {
        [Input("aadClientCredentialSecret", required: true)]
        public string AadClientCredentialSecret { get; set; } = null!;

        [Input("aadCredentialType", required: true)]
        public string AadCredentialType { get; set; } = null!;

        [Input("aadUserCredentialPassword", required: true)]
        private string? _aadUserCredentialPassword;
        public string? AadUserCredentialPassword
        {
            get => _aadUserCredentialPassword;
            set => _aadUserCredentialPassword = value;
        }

        [Input("aadUserCredentialUsername", required: true)]
        public string AadUserCredentialUsername { get; set; } = null!;

        [Input("certificateStoreLocation", required: true)]
        public string CertificateStoreLocation { get; set; } = null!;

        [Input("certificateStoreName", required: true)]
        public string CertificateStoreName { get; set; } = null!;

        [Input("clientCertificateVariable", required: true)]
        public string ClientCertificateVariable { get; set; } = null!;

        [Input("connectionEndpoint", required: true)]
        public string ConnectionEndpoint { get; set; } = null!;

        [Input("endpoints", required: true)]
        private List<Inputs.GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetEndpointArgs>? _endpoints;
        public List<Inputs.GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetEndpointArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new List<Inputs.GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetEndpointArgs>());
            set => _endpoints = value;
        }

        [Input("environments", required: true)]
        private List<string>? _environments;

        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        public List<string> Environments
        {
            get => _environments ?? (_environments = new List<string>());
            set => _environments = value;
        }

        [Input("hasLatestCalamari", required: true)]
        public bool HasLatestCalamari { get; set; }

        /// <summary>
        /// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        [Input("healthStatus", required: true)]
        public string HealthStatus { get; set; } = null!;

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// A filter to search by the disabled status of a resource.
        /// </summary>
        [Input("isDisabled", required: true)]
        public bool IsDisabled { get; set; }

        [Input("isInProcess", required: true)]
        public bool IsInProcess { get; set; }

        [Input("machinePolicyId", required: true)]
        public string MachinePolicyId { get; set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("operatingSystem", required: true)]
        public string OperatingSystem { get; set; } = null!;

        [Input("roles", required: true)]
        private List<string>? _roles;

        /// <summary>
        /// A filter to search by a list of role IDs.
        /// </summary>
        public List<string> Roles
        {
            get => _roles ?? (_roles = new List<string>());
            set => _roles = value;
        }

        [Input("securityMode", required: true)]
        public string SecurityMode { get; set; } = null!;

        [Input("serverCertificateThumbprint", required: true)]
        public string ServerCertificateThumbprint { get; set; } = null!;

        [Input("shellName", required: true)]
        public string ShellName { get; set; } = null!;

        [Input("shellVersion", required: true)]
        public string ShellVersion { get; set; } = null!;

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId", required: true)]
        public string SpaceId { get; set; } = null!;

        /// <summary>
        /// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
        /// </summary>
        [Input("status", required: true)]
        public string Status { get; set; } = null!;

        /// <summary>
        /// A summary elaborating on the status of this resource.
        /// </summary>
        [Input("statusSummary", required: true)]
        public string StatusSummary { get; set; } = null!;

        [Input("tenantTags", required: true)]
        private List<string>? _tenantTags;

        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public List<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new List<string>());
            set => _tenantTags = value;
        }

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Input("tenantedDeploymentParticipation", required: true)]
        public string TenantedDeploymentParticipation { get; set; } = null!;

        [Input("tenants", required: true)]
        private List<string>? _tenants;

        /// <summary>
        /// A list of tenant IDs associated with this resource.
        /// </summary>
        public List<string> Tenants
        {
            get => _tenants ?? (_tenants = new List<string>());
            set => _tenants = value;
        }

        /// <summary>
        /// The thumbprint of the deployment target to match in the query and/or search
        /// </summary>
        [Input("thumbprint", required: true)]
        public string Thumbprint { get; set; } = null!;

        [Input("uri", required: true)]
        public string Uri { get; set; } = null!;

        public GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetArgs()
        {
        }
        public static new GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetArgs Empty => new GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetArgs();
    }
}
