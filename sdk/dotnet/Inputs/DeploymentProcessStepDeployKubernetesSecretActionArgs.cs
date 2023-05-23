// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class DeploymentProcessStepDeployKubernetesSecretActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionTemplate")]
        public Input<Inputs.DeploymentProcessStepDeployKubernetesSecretActionActionTemplateArgs>? ActionTemplate { get; set; }

        [Input("canBeUsedForProjectVersioning")]
        public Input<bool>? CanBeUsedForProjectVersioning { get; set; }

        [Input("channels")]
        private InputList<string>? _channels;
        public InputList<string> Channels
        {
            get => _channels ?? (_channels = new InputList<string>());
            set => _channels = value;
        }

        [Input("condition")]
        public Input<string>? Condition { get; set; }

        [Input("containers")]
        private InputList<Inputs.DeploymentProcessStepDeployKubernetesSecretActionContainerArgs>? _containers;
        public InputList<Inputs.DeploymentProcessStepDeployKubernetesSecretActionContainerArgs> Containers
        {
            get => _containers ?? (_containers = new InputList<Inputs.DeploymentProcessStepDeployKubernetesSecretActionContainerArgs>());
            set => _containers = value;
        }

        [Input("environments")]
        private InputList<string>? _environments;
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        [Input("excludedEnvironments")]
        private InputList<string>? _excludedEnvironments;
        public InputList<string> ExcludedEnvironments
        {
            get => _excludedEnvironments ?? (_excludedEnvironments = new InputList<string>());
            set => _excludedEnvironments = value;
        }

        [Input("features")]
        private InputList<string>? _features;
        public InputList<string> Features
        {
            get => _features ?? (_features = new InputList<string>());
            set => _features = value;
        }

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        [Input("isRequired")]
        public Input<bool>? IsRequired { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("packages")]
        private InputList<Inputs.DeploymentProcessStepDeployKubernetesSecretActionPackageArgs>? _packages;
        public InputList<Inputs.DeploymentProcessStepDeployKubernetesSecretActionPackageArgs> Packages
        {
            get => _packages ?? (_packages = new InputList<Inputs.DeploymentProcessStepDeployKubernetesSecretActionPackageArgs>());
            set => _packages = value;
        }

        [Input("properties")]
        private InputMap<string>? _properties;
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        [Input("runOnServer")]
        public Input<bool>? RunOnServer { get; set; }

        [Input("secretName", required: true)]
        public Input<string> SecretName { get; set; } = null!;

        [Input("secretValues", required: true)]
        private InputMap<string>? _secretValues;
        public InputMap<string> SecretValues
        {
            get => _secretValues ?? (_secretValues = new InputMap<string>());
            set => _secretValues = value;
        }

        [Input("tenantTags")]
        private InputList<string>? _tenantTags;
        public InputList<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new InputList<string>());
            set => _tenantTags = value;
        }

        public DeploymentProcessStepDeployKubernetesSecretActionArgs()
        {
        }
        public static new DeploymentProcessStepDeployKubernetesSecretActionArgs Empty => new DeploymentProcessStepDeployKubernetesSecretActionArgs();
    }
}
