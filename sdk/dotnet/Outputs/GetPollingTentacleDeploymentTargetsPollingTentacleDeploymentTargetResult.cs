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
    public sealed class GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetResult
    {
        public readonly string CertificateSignatureAlgorithm;
        public readonly ImmutableArray<Outputs.GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetEndpointResult> Endpoints;
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
        public readonly string TentacleUrl;
        public readonly ImmutableArray<Outputs.GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetTentacleVersionDetailResult> TentacleVersionDetails;
        /// <summary>
        /// The thumbprint of the deployment target to match in the query and/or search
        /// </summary>
        public readonly string Thumbprint;
        public readonly string Uri;

        [OutputConstructor]
        private GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetResult(
            string certificateSignatureAlgorithm,

            ImmutableArray<Outputs.GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetEndpointResult> endpoints,

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

            string shellName,

            string shellVersion,

            string spaceId,

            string status,

            string statusSummary,

            ImmutableArray<string> tenantTags,

            string tenantedDeploymentParticipation,

            ImmutableArray<string> tenants,

            string tentacleUrl,

            ImmutableArray<Outputs.GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetTentacleVersionDetailResult> tentacleVersionDetails,

            string thumbprint,

            string uri)
        {
            CertificateSignatureAlgorithm = certificateSignatureAlgorithm;
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
            ShellName = shellName;
            ShellVersion = shellVersion;
            SpaceId = spaceId;
            Status = status;
            StatusSummary = statusSummary;
            TenantTags = tenantTags;
            TenantedDeploymentParticipation = tenantedDeploymentParticipation;
            Tenants = tenants;
            TentacleUrl = tentacleUrl;
            TentacleVersionDetails = tentacleVersionDetails;
            Thumbprint = thumbprint;
            Uri = uri;
        }
    }
}
