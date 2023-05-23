// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    public static class GetPollingTentacleDeploymentTargets
    {
        /// <summary>
        /// Provides information about existing polling tentacle deployment targets.
        /// </summary>
        public static Task<GetPollingTentacleDeploymentTargetsResult> InvokeAsync(GetPollingTentacleDeploymentTargetsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPollingTentacleDeploymentTargetsResult>("octopusdeploy:index/getPollingTentacleDeploymentTargets:getPollingTentacleDeploymentTargets", args ?? new GetPollingTentacleDeploymentTargetsArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about existing polling tentacle deployment targets.
        /// </summary>
        public static Output<GetPollingTentacleDeploymentTargetsResult> Invoke(GetPollingTentacleDeploymentTargetsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPollingTentacleDeploymentTargetsResult>("octopusdeploy:index/getPollingTentacleDeploymentTargets:getPollingTentacleDeploymentTargets", args ?? new GetPollingTentacleDeploymentTargetsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPollingTentacleDeploymentTargetsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter to search by deployment ID.
        /// </summary>
        [Input("deploymentId")]
        public string? DeploymentId { get; set; }

        [Input("environments")]
        private List<string>? _environments;

        /// <summary>
        /// A filter to search by a list of environment IDs.
        /// </summary>
        public List<string> Environments
        {
            get => _environments ?? (_environments = new List<string>());
            set => _environments = value;
        }

        [Input("healthStatuses")]
        private List<string>? _healthStatuses;

        /// <summary>
        /// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        public List<string> HealthStatuses
        {
            get => _healthStatuses ?? (_healthStatuses = new List<string>());
            set => _healthStatuses = value;
        }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A filter to search by the disabled status of a resource.
        /// </summary>
        [Input("isDisabled")]
        public bool? IsDisabled { get; set; }

        /// <summary>
        /// A filter to search by name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public string? PartialName { get; set; }

        [Input("pollingTentacleDeploymentTargets")]
        private List<Inputs.GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetArgs>? _pollingTentacleDeploymentTargets;

        /// <summary>
        /// A list of polling tentacle deployment targets that match the filter(s).
        /// </summary>
        public List<Inputs.GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetArgs> PollingTentacleDeploymentTargets
        {
            get => _pollingTentacleDeploymentTargets ?? (_pollingTentacleDeploymentTargets = new List<Inputs.GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetArgs>());
            set => _pollingTentacleDeploymentTargets = value;
        }

        [Input("roles")]
        private List<string>? _roles;

        /// <summary>
        /// A filter to search by a list of role IDs.
        /// </summary>
        public List<string> Roles
        {
            get => _roles ?? (_roles = new List<string>());
            set => _roles = value;
        }

        [Input("shellNames")]
        private List<string>? _shellNames;

        /// <summary>
        /// A list of shell names to match in the query and/or search
        /// </summary>
        public List<string> ShellNames
        {
            get => _shellNames ?? (_shellNames = new List<string>());
            set => _shellNames = value;
        }

        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        [Input("skip")]
        public int? Skip { get; set; }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public int? Take { get; set; }

        [Input("tenantTags")]
        private List<string>? _tenantTags;

        /// <summary>
        /// A filter to search by a list of tenant tags.
        /// </summary>
        public List<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new List<string>());
            set => _tenantTags = value;
        }

        [Input("tenants")]
        private List<string>? _tenants;

        /// <summary>
        /// A filter to search by a list of tenant IDs.
        /// </summary>
        public List<string> Tenants
        {
            get => _tenants ?? (_tenants = new List<string>());
            set => _tenants = value;
        }

        /// <summary>
        /// The thumbprint of the deployment target to match in the query and/or search
        /// </summary>
        [Input("thumbprint")]
        public string? Thumbprint { get; set; }

        public GetPollingTentacleDeploymentTargetsArgs()
        {
        }
        public static new GetPollingTentacleDeploymentTargetsArgs Empty => new GetPollingTentacleDeploymentTargetsArgs();
    }

    public sealed class GetPollingTentacleDeploymentTargetsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter to search by deployment ID.
        /// </summary>
        [Input("deploymentId")]
        public Input<string>? DeploymentId { get; set; }

        [Input("environments")]
        private InputList<string>? _environments;

        /// <summary>
        /// A filter to search by a list of environment IDs.
        /// </summary>
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        [Input("healthStatuses")]
        private InputList<string>? _healthStatuses;

        /// <summary>
        /// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        public InputList<string> HealthStatuses
        {
            get => _healthStatuses ?? (_healthStatuses = new InputList<string>());
            set => _healthStatuses = value;
        }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A filter to search by the disabled status of a resource.
        /// </summary>
        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        /// <summary>
        /// A filter to search by name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public Input<string>? PartialName { get; set; }

        [Input("pollingTentacleDeploymentTargets")]
        private InputList<Inputs.GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetInputArgs>? _pollingTentacleDeploymentTargets;

        /// <summary>
        /// A list of polling tentacle deployment targets that match the filter(s).
        /// </summary>
        public InputList<Inputs.GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetInputArgs> PollingTentacleDeploymentTargets
        {
            get => _pollingTentacleDeploymentTargets ?? (_pollingTentacleDeploymentTargets = new InputList<Inputs.GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetInputArgs>());
            set => _pollingTentacleDeploymentTargets = value;
        }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// A filter to search by a list of role IDs.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("shellNames")]
        private InputList<string>? _shellNames;

        /// <summary>
        /// A list of shell names to match in the query and/or search
        /// </summary>
        public InputList<string> ShellNames
        {
            get => _shellNames ?? (_shellNames = new InputList<string>());
            set => _shellNames = value;
        }

        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        [Input("skip")]
        public Input<int>? Skip { get; set; }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public Input<int>? Take { get; set; }

        [Input("tenantTags")]
        private InputList<string>? _tenantTags;

        /// <summary>
        /// A filter to search by a list of tenant tags.
        /// </summary>
        public InputList<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new InputList<string>());
            set => _tenantTags = value;
        }

        [Input("tenants")]
        private InputList<string>? _tenants;

        /// <summary>
        /// A filter to search by a list of tenant IDs.
        /// </summary>
        public InputList<string> Tenants
        {
            get => _tenants ?? (_tenants = new InputList<string>());
            set => _tenants = value;
        }

        /// <summary>
        /// The thumbprint of the deployment target to match in the query and/or search
        /// </summary>
        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        public GetPollingTentacleDeploymentTargetsInvokeArgs()
        {
        }
        public static new GetPollingTentacleDeploymentTargetsInvokeArgs Empty => new GetPollingTentacleDeploymentTargetsInvokeArgs();
    }


    [OutputType]
    public sealed class GetPollingTentacleDeploymentTargetsResult
    {
        /// <summary>
        /// A filter to search by deployment ID.
        /// </summary>
        public readonly string? DeploymentId;
        /// <summary>
        /// A filter to search by a list of environment IDs.
        /// </summary>
        public readonly ImmutableArray<string> Environments;
        /// <summary>
        /// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        public readonly ImmutableArray<string> HealthStatuses;
        /// <summary>
        /// An auto-generated identifier that includes the timestamp when this data source was last modified.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A filter to search by the disabled status of a resource.
        /// </summary>
        public readonly bool? IsDisabled;
        /// <summary>
        /// A filter to search by name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        public readonly string? PartialName;
        /// <summary>
        /// A list of polling tentacle deployment targets that match the filter(s).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetResult> PollingTentacleDeploymentTargets;
        /// <summary>
        /// A filter to search by a list of role IDs.
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        /// <summary>
        /// A list of shell names to match in the query and/or search
        /// </summary>
        public readonly ImmutableArray<string> ShellNames;
        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        public readonly int? Skip;
        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        public readonly int? Take;
        /// <summary>
        /// A filter to search by a list of tenant tags.
        /// </summary>
        public readonly ImmutableArray<string> TenantTags;
        /// <summary>
        /// A filter to search by a list of tenant IDs.
        /// </summary>
        public readonly ImmutableArray<string> Tenants;
        /// <summary>
        /// The thumbprint of the deployment target to match in the query and/or search
        /// </summary>
        public readonly string? Thumbprint;

        [OutputConstructor]
        private GetPollingTentacleDeploymentTargetsResult(
            string? deploymentId,

            ImmutableArray<string> environments,

            ImmutableArray<string> healthStatuses,

            string id,

            ImmutableArray<string> ids,

            bool? isDisabled,

            string? name,

            string? partialName,

            ImmutableArray<Outputs.GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetResult> pollingTentacleDeploymentTargets,

            ImmutableArray<string> roles,

            ImmutableArray<string> shellNames,

            int? skip,

            int? take,

            ImmutableArray<string> tenantTags,

            ImmutableArray<string> tenants,

            string? thumbprint)
        {
            DeploymentId = deploymentId;
            Environments = environments;
            HealthStatuses = healthStatuses;
            Id = id;
            Ids = ids;
            IsDisabled = isDisabled;
            Name = name;
            PartialName = partialName;
            PollingTentacleDeploymentTargets = pollingTentacleDeploymentTargets;
            Roles = roles;
            ShellNames = shellNames;
            Skip = skip;
            Take = take;
            TenantTags = tenantTags;
            Tenants = tenants;
            Thumbprint = thumbprint;
        }
    }
}
