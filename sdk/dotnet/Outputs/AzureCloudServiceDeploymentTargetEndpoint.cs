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
    public sealed class AzureCloudServiceDeploymentTargetEndpoint
    {
        public readonly string? AadClientCredentialSecret;
        public readonly string? AadCredentialType;
        public readonly string? AadUserCredentialUsername;
        public readonly string? AccountId;
        public readonly string? ApplicationsDirectory;
        public readonly Outputs.AzureCloudServiceDeploymentTargetEndpointAuthentication? Authentication;
        public readonly string? CertificateSignatureAlgorithm;
        public readonly string? CertificateStoreLocation;
        public readonly string? CertificateStoreName;
        public readonly string? ClientCertificateVariable;
        public readonly string? CloudServiceName;
        public readonly string? ClusterCertificate;
        public readonly string? ClusterUrl;
        public readonly string CommunicationStyle;
        public readonly string? ConnectionEndpoint;
        public readonly ImmutableArray<Outputs.AzureCloudServiceDeploymentTargetEndpointContainer> Containers;
        public readonly string? DefaultWorkerPoolId;
        public readonly ImmutableArray<Outputs.AzureCloudServiceDeploymentTargetEndpointDestination> Destinations;
        public readonly string? DotNetCorePlatform;
        public readonly string? Fingerprint;
        public readonly string? Host;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string? Id;
        public readonly string? Namespace;
        public readonly int? Port;
        public readonly string? ProxyId;
        public readonly string? ResourceGroupName;
        public readonly bool? RunningInContainer;
        public readonly string? SecurityMode;
        public readonly string? ServerCertificateThumbprint;
        public readonly bool? SkipTlsVerification;
        public readonly string? Slot;
        public readonly string? StorageAccountName;
        public readonly bool? SwapIfPossible;
        public readonly ImmutableArray<Outputs.AzureCloudServiceDeploymentTargetEndpointTentacleVersionDetail> TentacleVersionDetails;
        public readonly string? Thumbprint;
        public readonly string? Uri;
        public readonly bool? UseCurrentInstanceCount;
        public readonly string? WebAppName;
        public readonly string? WebAppSlotName;
        public readonly string? WorkingDirectory;

        [OutputConstructor]
        private AzureCloudServiceDeploymentTargetEndpoint(
            string? aadClientCredentialSecret,

            string? aadCredentialType,

            string? aadUserCredentialUsername,

            string? accountId,

            string? applicationsDirectory,

            Outputs.AzureCloudServiceDeploymentTargetEndpointAuthentication? authentication,

            string? certificateSignatureAlgorithm,

            string? certificateStoreLocation,

            string? certificateStoreName,

            string? clientCertificateVariable,

            string? cloudServiceName,

            string? clusterCertificate,

            string? clusterUrl,

            string communicationStyle,

            string? connectionEndpoint,

            ImmutableArray<Outputs.AzureCloudServiceDeploymentTargetEndpointContainer> containers,

            string? defaultWorkerPoolId,

            ImmutableArray<Outputs.AzureCloudServiceDeploymentTargetEndpointDestination> destinations,

            string? dotNetCorePlatform,

            string? fingerprint,

            string? host,

            string? id,

            string? @namespace,

            int? port,

            string? proxyId,

            string? resourceGroupName,

            bool? runningInContainer,

            string? securityMode,

            string? serverCertificateThumbprint,

            bool? skipTlsVerification,

            string? slot,

            string? storageAccountName,

            bool? swapIfPossible,

            ImmutableArray<Outputs.AzureCloudServiceDeploymentTargetEndpointTentacleVersionDetail> tentacleVersionDetails,

            string? thumbprint,

            string? uri,

            bool? useCurrentInstanceCount,

            string? webAppName,

            string? webAppSlotName,

            string? workingDirectory)
        {
            AadClientCredentialSecret = aadClientCredentialSecret;
            AadCredentialType = aadCredentialType;
            AadUserCredentialUsername = aadUserCredentialUsername;
            AccountId = accountId;
            ApplicationsDirectory = applicationsDirectory;
            Authentication = authentication;
            CertificateSignatureAlgorithm = certificateSignatureAlgorithm;
            CertificateStoreLocation = certificateStoreLocation;
            CertificateStoreName = certificateStoreName;
            ClientCertificateVariable = clientCertificateVariable;
            CloudServiceName = cloudServiceName;
            ClusterCertificate = clusterCertificate;
            ClusterUrl = clusterUrl;
            CommunicationStyle = communicationStyle;
            ConnectionEndpoint = connectionEndpoint;
            Containers = containers;
            DefaultWorkerPoolId = defaultWorkerPoolId;
            Destinations = destinations;
            DotNetCorePlatform = dotNetCorePlatform;
            Fingerprint = fingerprint;
            Host = host;
            Id = id;
            Namespace = @namespace;
            Port = port;
            ProxyId = proxyId;
            ResourceGroupName = resourceGroupName;
            RunningInContainer = runningInContainer;
            SecurityMode = securityMode;
            ServerCertificateThumbprint = serverCertificateThumbprint;
            SkipTlsVerification = skipTlsVerification;
            Slot = slot;
            StorageAccountName = storageAccountName;
            SwapIfPossible = swapIfPossible;
            TentacleVersionDetails = tentacleVersionDetails;
            Thumbprint = thumbprint;
            Uri = uri;
            UseCurrentInstanceCount = useCurrentInstanceCount;
            WebAppName = webAppName;
            WebAppSlotName = webAppSlotName;
            WorkingDirectory = workingDirectory;
        }
    }
}
