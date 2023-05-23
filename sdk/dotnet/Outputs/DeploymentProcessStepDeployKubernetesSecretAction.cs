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
    public sealed class DeploymentProcessStepDeployKubernetesSecretAction
    {
        public readonly Outputs.DeploymentProcessStepDeployKubernetesSecretActionActionTemplate? ActionTemplate;
        public readonly bool? CanBeUsedForProjectVersioning;
        public readonly ImmutableArray<string> Channels;
        public readonly string? Condition;
        public readonly ImmutableArray<Outputs.DeploymentProcessStepDeployKubernetesSecretActionContainer> Containers;
        public readonly ImmutableArray<string> Environments;
        public readonly ImmutableArray<string> ExcludedEnvironments;
        public readonly ImmutableArray<string> Features;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string? Id;
        public readonly bool? IsDisabled;
        public readonly bool? IsRequired;
        public readonly string Name;
        public readonly string? Notes;
        public readonly ImmutableArray<Outputs.DeploymentProcessStepDeployKubernetesSecretActionPackage> Packages;
        public readonly ImmutableDictionary<string, string>? Properties;
        public readonly bool? RunOnServer;
        public readonly string SecretName;
        public readonly ImmutableDictionary<string, string> SecretValues;
        public readonly ImmutableArray<string> TenantTags;

        [OutputConstructor]
        private DeploymentProcessStepDeployKubernetesSecretAction(
            Outputs.DeploymentProcessStepDeployKubernetesSecretActionActionTemplate? actionTemplate,

            bool? canBeUsedForProjectVersioning,

            ImmutableArray<string> channels,

            string? condition,

            ImmutableArray<Outputs.DeploymentProcessStepDeployKubernetesSecretActionContainer> containers,

            ImmutableArray<string> environments,

            ImmutableArray<string> excludedEnvironments,

            ImmutableArray<string> features,

            string? id,

            bool? isDisabled,

            bool? isRequired,

            string name,

            string? notes,

            ImmutableArray<Outputs.DeploymentProcessStepDeployKubernetesSecretActionPackage> packages,

            ImmutableDictionary<string, string>? properties,

            bool? runOnServer,

            string secretName,

            ImmutableDictionary<string, string> secretValues,

            ImmutableArray<string> tenantTags)
        {
            ActionTemplate = actionTemplate;
            CanBeUsedForProjectVersioning = canBeUsedForProjectVersioning;
            Channels = channels;
            Condition = condition;
            Containers = containers;
            Environments = environments;
            ExcludedEnvironments = excludedEnvironments;
            Features = features;
            Id = id;
            IsDisabled = isDisabled;
            IsRequired = isRequired;
            Name = name;
            Notes = notes;
            Packages = packages;
            Properties = properties;
            RunOnServer = runOnServer;
            SecretName = secretName;
            SecretValues = secretValues;
            TenantTags = tenantTags;
        }
    }
}
