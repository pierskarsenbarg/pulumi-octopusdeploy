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

    public sealed class GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("aadClientCredentialSecret")]
        public Input<string>? AadClientCredentialSecret { get; set; }

        [Input("aadCredentialType")]
        public Input<string>? AadCredentialType { get; set; }

        [Input("aadUserCredentialUsername")]
        public Input<string>? AadUserCredentialUsername { get; set; }

        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        [Input("applicationsDirectory")]
        public Input<string>? ApplicationsDirectory { get; set; }

        [Input("authentication", required: true)]
        public Input<Inputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointAuthenticationInputArgs> Authentication { get; set; } = null!;

        [Input("certificateSignatureAlgorithm")]
        public Input<string>? CertificateSignatureAlgorithm { get; set; }

        [Input("certificateStoreLocation")]
        public Input<string>? CertificateStoreLocation { get; set; }

        [Input("certificateStoreName")]
        public Input<string>? CertificateStoreName { get; set; }

        [Input("clientCertificateVariable")]
        public Input<string>? ClientCertificateVariable { get; set; }

        [Input("cloudServiceName")]
        public Input<string>? CloudServiceName { get; set; }

        [Input("clusterCertificate")]
        public Input<string>? ClusterCertificate { get; set; }

        [Input("clusterUrl")]
        public Input<string>? ClusterUrl { get; set; }

        [Input("communicationStyle", required: true)]
        public Input<string> CommunicationStyle { get; set; } = null!;

        [Input("connectionEndpoint")]
        public Input<string>? ConnectionEndpoint { get; set; }

        [Input("containers", required: true)]
        private InputList<Inputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointContainerInputArgs>? _containers;
        public InputList<Inputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointContainerInputArgs> Containers
        {
            get => _containers ?? (_containers = new InputList<Inputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointContainerInputArgs>());
            set => _containers = value;
        }

        [Input("defaultWorkerPoolId")]
        public Input<string>? DefaultWorkerPoolId { get; set; }

        [Input("destinations", required: true)]
        private InputList<Inputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointDestinationInputArgs>? _destinations;
        public InputList<Inputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointDestinationInputArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointDestinationInputArgs>());
            set => _destinations = value;
        }

        [Input("dotNetCorePlatform")]
        public Input<string>? DotNetCorePlatform { get; set; }

        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// An auto-generated identifier that includes the timestamp when this data source was last modified.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("proxyId")]
        public Input<string>? ProxyId { get; set; }

        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("runningInContainer")]
        public Input<bool>? RunningInContainer { get; set; }

        [Input("securityMode")]
        public Input<string>? SecurityMode { get; set; }

        [Input("serverCertificateThumbprint")]
        public Input<string>? ServerCertificateThumbprint { get; set; }

        [Input("skipTlsVerification")]
        public Input<bool>? SkipTlsVerification { get; set; }

        [Input("slot")]
        public Input<string>? Slot { get; set; }

        [Input("storageAccountName")]
        public Input<string>? StorageAccountName { get; set; }

        [Input("swapIfPossible")]
        public Input<bool>? SwapIfPossible { get; set; }

        [Input("tentacleVersionDetails", required: true)]
        private InputList<Inputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointTentacleVersionDetailInputArgs>? _tentacleVersionDetails;
        public InputList<Inputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointTentacleVersionDetailInputArgs> TentacleVersionDetails
        {
            get => _tentacleVersionDetails ?? (_tentacleVersionDetails = new InputList<Inputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointTentacleVersionDetailInputArgs>());
            set => _tentacleVersionDetails = value;
        }

        /// <summary>
        /// The thumbprint of the deployment target to match in the query and/or search
        /// </summary>
        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        [Input("uri")]
        public Input<string>? Uri { get; set; }

        [Input("useCurrentInstanceCount")]
        public Input<bool>? UseCurrentInstanceCount { get; set; }

        [Input("webAppName")]
        public Input<string>? WebAppName { get; set; }

        [Input("webAppSlotName")]
        public Input<string>? WebAppSlotName { get; set; }

        [Input("workingDirectory")]
        public Input<string>? WorkingDirectory { get; set; }

        public GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointInputArgs()
        {
        }
        public static new GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointInputArgs Empty => new GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointInputArgs();
    }
}
