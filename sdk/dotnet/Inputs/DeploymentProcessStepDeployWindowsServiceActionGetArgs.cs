// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class DeploymentProcessStepDeployWindowsServiceActionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionTemplate")]
        public Input<Inputs.DeploymentProcessStepDeployWindowsServiceActionActionTemplateGetArgs>? ActionTemplate { get; set; }

        [Input("arguments")]
        public Input<string>? Arguments { get; set; }

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
        private InputList<Inputs.DeploymentProcessStepDeployWindowsServiceActionContainerGetArgs>? _containers;
        public InputList<Inputs.DeploymentProcessStepDeployWindowsServiceActionContainerGetArgs> Containers
        {
            get => _containers ?? (_containers = new InputList<Inputs.DeploymentProcessStepDeployWindowsServiceActionContainerGetArgs>());
            set => _containers = value;
        }

        [Input("createOrUpdateService")]
        public Input<bool>? CreateOrUpdateService { get; set; }

        [Input("customAccountName")]
        public Input<string>? CustomAccountName { get; set; }

        [Input("customAccountPassword")]
        public Input<string>? CustomAccountPassword { get; set; }

        [Input("dependencies")]
        public Input<string>? Dependencies { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

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

        [Input("executablePath", required: true)]
        public Input<string> ExecutablePath { get; set; } = null!;

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
        private InputList<Inputs.DeploymentProcessStepDeployWindowsServiceActionPackageGetArgs>? _packages;
        public InputList<Inputs.DeploymentProcessStepDeployWindowsServiceActionPackageGetArgs> Packages
        {
            get => _packages ?? (_packages = new InputList<Inputs.DeploymentProcessStepDeployWindowsServiceActionPackageGetArgs>());
            set => _packages = value;
        }

        [Input("primaryPackage", required: true)]
        public Input<Inputs.DeploymentProcessStepDeployWindowsServiceActionPrimaryPackageGetArgs> PrimaryPackage { get; set; } = null!;

        [Input("properties")]
        private InputMap<string>? _properties;
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        [Input("startMode")]
        public Input<string>? StartMode { get; set; }

        [Input("tenantTags")]
        private InputList<string>? _tenantTags;
        public InputList<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new InputList<string>());
            set => _tenantTags = value;
        }

        public DeploymentProcessStepDeployWindowsServiceActionGetArgs()
        {
        }
        public static new DeploymentProcessStepDeployWindowsServiceActionGetArgs Empty => new DeploymentProcessStepDeployWindowsServiceActionGetArgs();
    }
}
