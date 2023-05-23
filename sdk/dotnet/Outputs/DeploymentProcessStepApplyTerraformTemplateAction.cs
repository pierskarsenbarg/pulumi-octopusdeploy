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
    public sealed class DeploymentProcessStepApplyTerraformTemplateAction
    {
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionActionTemplate? ActionTemplate;
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionAdvancedOptions AdvancedOptions;
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionAwsAccount? AwsAccount;
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionAzureAccount? AzureAccount;
        public readonly bool? CanBeUsedForProjectVersioning;
        public readonly ImmutableArray<string> Channels;
        public readonly string? Condition;
        public readonly ImmutableArray<Outputs.DeploymentProcessStepApplyTerraformTemplateActionContainer> Containers;
        public readonly ImmutableArray<string> Environments;
        public readonly ImmutableArray<string> ExcludedEnvironments;
        public readonly ImmutableArray<string> Features;
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionGoogleCloudAccount? GoogleCloudAccount;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string? Id;
        public readonly string? InlineTemplate;
        public readonly bool? IsDisabled;
        public readonly bool? IsRequired;
        public readonly string Name;
        public readonly string? Notes;
        public readonly ImmutableArray<Outputs.DeploymentProcessStepApplyTerraformTemplateActionPackage> Packages;
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionPrimaryPackage? PrimaryPackage;
        public readonly ImmutableDictionary<string, string>? Properties;
        public readonly bool? RunOnServer;
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionTemplate? Template;
        public readonly string? TemplateParameters;
        public readonly ImmutableArray<string> TenantTags;

        [OutputConstructor]
        private DeploymentProcessStepApplyTerraformTemplateAction(
            Outputs.DeploymentProcessStepApplyTerraformTemplateActionActionTemplate? actionTemplate,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionAdvancedOptions advancedOptions,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionAwsAccount? awsAccount,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionAzureAccount? azureAccount,

            bool? canBeUsedForProjectVersioning,

            ImmutableArray<string> channels,

            string? condition,

            ImmutableArray<Outputs.DeploymentProcessStepApplyTerraformTemplateActionContainer> containers,

            ImmutableArray<string> environments,

            ImmutableArray<string> excludedEnvironments,

            ImmutableArray<string> features,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionGoogleCloudAccount? googleCloudAccount,

            string? id,

            string? inlineTemplate,

            bool? isDisabled,

            bool? isRequired,

            string name,

            string? notes,

            ImmutableArray<Outputs.DeploymentProcessStepApplyTerraformTemplateActionPackage> packages,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionPrimaryPackage? primaryPackage,

            ImmutableDictionary<string, string>? properties,

            bool? runOnServer,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionTemplate? template,

            string? templateParameters,

            ImmutableArray<string> tenantTags)
        {
            ActionTemplate = actionTemplate;
            AdvancedOptions = advancedOptions;
            AwsAccount = awsAccount;
            AzureAccount = azureAccount;
            CanBeUsedForProjectVersioning = canBeUsedForProjectVersioning;
            Channels = channels;
            Condition = condition;
            Containers = containers;
            Environments = environments;
            ExcludedEnvironments = excludedEnvironments;
            Features = features;
            GoogleCloudAccount = googleCloudAccount;
            Id = id;
            InlineTemplate = inlineTemplate;
            IsDisabled = isDisabled;
            IsRequired = isRequired;
            Name = name;
            Notes = notes;
            Packages = packages;
            PrimaryPackage = primaryPackage;
            Properties = properties;
            RunOnServer = runOnServer;
            Template = template;
            TemplateParameters = templateParameters;
            TenantTags = tenantTags;
        }
    }
}
