// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class SshConnectionDeploymentTargetEndpointArgs : global::Pulumi.ResourceArgs
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

        [Input("authentication")]
        public Input<Inputs.SshConnectionDeploymentTargetEndpointAuthenticationArgs>? Authentication { get; set; }

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

        [Input("containers")]
        private InputList<Inputs.SshConnectionDeploymentTargetEndpointContainerArgs>? _containers;
        public InputList<Inputs.SshConnectionDeploymentTargetEndpointContainerArgs> Containers
        {
            get => _containers ?? (_containers = new InputList<Inputs.SshConnectionDeploymentTargetEndpointContainerArgs>());
            set => _containers = value;
        }

        [Input("defaultWorkerPoolId")]
        public Input<string>? DefaultWorkerPoolId { get; set; }

        [Input("destinations")]
        private InputList<Inputs.SshConnectionDeploymentTargetEndpointDestinationArgs>? _destinations;
        public InputList<Inputs.SshConnectionDeploymentTargetEndpointDestinationArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.SshConnectionDeploymentTargetEndpointDestinationArgs>());
            set => _destinations = value;
        }

        [Input("dotNetCorePlatform")]
        public Input<string>? DotNetCorePlatform { get; set; }

        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

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

        [Input("tentacleVersionDetails")]
        private InputList<Inputs.SshConnectionDeploymentTargetEndpointTentacleVersionDetailArgs>? _tentacleVersionDetails;
        public InputList<Inputs.SshConnectionDeploymentTargetEndpointTentacleVersionDetailArgs> TentacleVersionDetails
        {
            get => _tentacleVersionDetails ?? (_tentacleVersionDetails = new InputList<Inputs.SshConnectionDeploymentTargetEndpointTentacleVersionDetailArgs>());
            set => _tentacleVersionDetails = value;
        }

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

        public SshConnectionDeploymentTargetEndpointArgs()
        {
        }
        public static new SshConnectionDeploymentTargetEndpointArgs Empty => new SshConnectionDeploymentTargetEndpointArgs();
    }
}