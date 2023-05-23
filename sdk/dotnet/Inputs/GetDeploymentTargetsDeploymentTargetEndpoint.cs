// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetDeploymentTargetsDeploymentTargetEndpointArgs : global::Pulumi.InvokeArgs
    {
        [Input("aadClientCredentialSecret")]
        public string? AadClientCredentialSecret { get; set; }

        [Input("aadCredentialType")]
        public string? AadCredentialType { get; set; }

        [Input("aadUserCredentialUsername")]
        public string? AadUserCredentialUsername { get; set; }

        [Input("accountId")]
        public string? AccountId { get; set; }

        [Input("applicationsDirectory")]
        public string? ApplicationsDirectory { get; set; }

        [Input("authentication", required: true)]
        public Inputs.GetDeploymentTargetsDeploymentTargetEndpointAuthenticationArgs Authentication { get; set; } = null!;

        [Input("certificateSignatureAlgorithm")]
        public string? CertificateSignatureAlgorithm { get; set; }

        [Input("certificateStoreLocation")]
        public string? CertificateStoreLocation { get; set; }

        [Input("certificateStoreName")]
        public string? CertificateStoreName { get; set; }

        [Input("clientCertificateVariable")]
        public string? ClientCertificateVariable { get; set; }

        [Input("cloudServiceName")]
        public string? CloudServiceName { get; set; }

        [Input("clusterCertificate")]
        public string? ClusterCertificate { get; set; }

        [Input("clusterUrl")]
        public string? ClusterUrl { get; set; }

        [Input("communicationStyle", required: true)]
        public string CommunicationStyle { get; set; } = null!;

        [Input("connectionEndpoint")]
        public string? ConnectionEndpoint { get; set; }

        [Input("containers", required: true)]
        private List<Inputs.GetDeploymentTargetsDeploymentTargetEndpointContainerArgs>? _containers;
        public List<Inputs.GetDeploymentTargetsDeploymentTargetEndpointContainerArgs> Containers
        {
            get => _containers ?? (_containers = new List<Inputs.GetDeploymentTargetsDeploymentTargetEndpointContainerArgs>());
            set => _containers = value;
        }

        [Input("defaultWorkerPoolId")]
        public string? DefaultWorkerPoolId { get; set; }

        [Input("destinations", required: true)]
        private List<Inputs.GetDeploymentTargetsDeploymentTargetEndpointDestinationArgs>? _destinations;
        public List<Inputs.GetDeploymentTargetsDeploymentTargetEndpointDestinationArgs> Destinations
        {
            get => _destinations ?? (_destinations = new List<Inputs.GetDeploymentTargetsDeploymentTargetEndpointDestinationArgs>());
            set => _destinations = value;
        }

        [Input("dotNetCorePlatform")]
        public string? DotNetCorePlatform { get; set; }

        [Input("fingerprint")]
        public string? Fingerprint { get; set; }

        [Input("host")]
        public string? Host { get; set; }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("namespace")]
        public string? Namespace { get; set; }

        [Input("port")]
        public int? Port { get; set; }

        [Input("proxyId")]
        public string? ProxyId { get; set; }

        [Input("resourceGroupName")]
        public string? ResourceGroupName { get; set; }

        [Input("runningInContainer")]
        public bool? RunningInContainer { get; set; }

        [Input("securityMode")]
        public string? SecurityMode { get; set; }

        [Input("serverCertificateThumbprint")]
        public string? ServerCertificateThumbprint { get; set; }

        [Input("skipTlsVerification")]
        public bool? SkipTlsVerification { get; set; }

        [Input("slot")]
        public string? Slot { get; set; }

        [Input("storageAccountName")]
        public string? StorageAccountName { get; set; }

        [Input("swapIfPossible")]
        public bool? SwapIfPossible { get; set; }

        [Input("tentacleVersionDetails", required: true)]
        private List<Inputs.GetDeploymentTargetsDeploymentTargetEndpointTentacleVersionDetailArgs>? _tentacleVersionDetails;
        public List<Inputs.GetDeploymentTargetsDeploymentTargetEndpointTentacleVersionDetailArgs> TentacleVersionDetails
        {
            get => _tentacleVersionDetails ?? (_tentacleVersionDetails = new List<Inputs.GetDeploymentTargetsDeploymentTargetEndpointTentacleVersionDetailArgs>());
            set => _tentacleVersionDetails = value;
        }

        /// <summary>
        /// The thumbprint of the deployment target to match in the query and/or search
        /// </summary>
        [Input("thumbprint")]
        public string? Thumbprint { get; set; }

        [Input("uri")]
        public string? Uri { get; set; }

        [Input("useCurrentInstanceCount")]
        public bool? UseCurrentInstanceCount { get; set; }

        [Input("webAppName")]
        public string? WebAppName { get; set; }

        [Input("webAppSlotName")]
        public string? WebAppSlotName { get; set; }

        [Input("workingDirectory")]
        public string? WorkingDirectory { get; set; }

        public GetDeploymentTargetsDeploymentTargetEndpointArgs()
        {
        }
        public static new GetDeploymentTargetsDeploymentTargetEndpointArgs Empty => new GetDeploymentTargetsDeploymentTargetEndpointArgs();
    }
}