// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy
{
    /// <summary>
    /// This resource manages SSH connection deployment targets in Octopus Deploy.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Octopusdeploy = PiersKarsenbarg.Octopusdeploy;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Octopusdeploy.SshConnectionDeploymentTarget("example", new()
    ///     {
    ///         Fingerprint = "[fingerprint]",
    ///         Host = "[host]",
    ///         Port = 22,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import octopusdeploy:index/sshConnectionDeploymentTarget:SshConnectionDeploymentTarget [options] octopusdeploy_ssh_connection_deployment_target.&lt;name&gt; &lt;account-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/sshConnectionDeploymentTarget:SshConnectionDeploymentTarget")]
    public partial class SshConnectionDeploymentTarget : global::Pulumi.CustomResource
    {
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        [Output("dotNetCorePlatform")]
        public Output<string?> DotNetCorePlatform { get; private set; } = null!;

        [Output("endpoints")]
        public Output<ImmutableArray<Outputs.SshConnectionDeploymentTargetEndpoint>> Endpoints { get; private set; } = null!;

        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        [Output("environments")]
        public Output<ImmutableArray<string>> Environments { get; private set; } = null!;

        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        [Output("hasLatestCalamari")]
        public Output<bool> HasLatestCalamari { get; private set; } = null!;

        /// <summary>
        /// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        [Output("healthStatus")]
        public Output<string> HealthStatus { get; private set; } = null!;

        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        [Output("isDisabled")]
        public Output<bool> IsDisabled { get; private set; } = null!;

        [Output("isInProcess")]
        public Output<bool> IsInProcess { get; private set; } = null!;

        [Output("machinePolicyId")]
        public Output<string> MachinePolicyId { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("operatingSystem")]
        public Output<string> OperatingSystem { get; private set; } = null!;

        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        [Output("proxyId")]
        public Output<string?> ProxyId { get; private set; } = null!;

        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        [Output("shellName")]
        public Output<string> ShellName { get; private set; } = null!;

        [Output("shellVersion")]
        public Output<string> ShellVersion { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

        /// <summary>
        /// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A summary elaborating on the status of this resource.
        /// </summary>
        [Output("statusSummary")]
        public Output<string> StatusSummary { get; private set; } = null!;

        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        [Output("tenantTags")]
        public Output<ImmutableArray<string>> TenantTags { get; private set; } = null!;

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Output("tenantedDeploymentParticipation")]
        public Output<string> TenantedDeploymentParticipation { get; private set; } = null!;

        /// <summary>
        /// A list of tenant IDs associated with this resource.
        /// </summary>
        [Output("tenants")]
        public Output<ImmutableArray<string>> Tenants { get; private set; } = null!;

        [Output("thumbprint")]
        public Output<string> Thumbprint { get; private set; } = null!;

        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;


        /// <summary>
        /// Create a SshConnectionDeploymentTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SshConnectionDeploymentTarget(string name, SshConnectionDeploymentTargetArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/sshConnectionDeploymentTarget:SshConnectionDeploymentTarget", name, args ?? new SshConnectionDeploymentTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SshConnectionDeploymentTarget(string name, Input<string> id, SshConnectionDeploymentTargetState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/sshConnectionDeploymentTarget:SshConnectionDeploymentTarget", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SshConnectionDeploymentTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SshConnectionDeploymentTarget Get(string name, Input<string> id, SshConnectionDeploymentTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new SshConnectionDeploymentTarget(name, id, state, options);
        }
    }

    public sealed class SshConnectionDeploymentTargetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        [Input("dotNetCorePlatform")]
        public Input<string>? DotNetCorePlatform { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.SshConnectionDeploymentTargetEndpointArgs>? _endpoints;
        public InputList<Inputs.SshConnectionDeploymentTargetEndpointArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.SshConnectionDeploymentTargetEndpointArgs>());
            set => _endpoints = value;
        }

        [Input("environments", required: true)]
        private InputList<string>? _environments;

        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        [Input("fingerprint", required: true)]
        public Input<string> Fingerprint { get; set; } = null!;

        /// <summary>
        /// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        [Input("healthStatus")]
        public Input<string>? HealthStatus { get; set; }

        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        [Input("machinePolicyId")]
        public Input<string>? MachinePolicyId { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("operatingSystem")]
        public Input<string>? OperatingSystem { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("proxyId")]
        public Input<string>? ProxyId { get; set; }

        [Input("roles", required: true)]
        private InputList<string>? _roles;
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("shellName")]
        public Input<string>? ShellName { get; set; }

        [Input("shellVersion")]
        public Input<string>? ShellVersion { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// A summary elaborating on the status of this resource.
        /// </summary>
        [Input("statusSummary")]
        public Input<string>? StatusSummary { get; set; }

        [Input("tenantTags")]
        private InputList<string>? _tenantTags;

        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public InputList<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new InputList<string>());
            set => _tenantTags = value;
        }

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Input("tenantedDeploymentParticipation")]
        public Input<string>? TenantedDeploymentParticipation { get; set; }

        [Input("tenants")]
        private InputList<string>? _tenants;

        /// <summary>
        /// A list of tenant IDs associated with this resource.
        /// </summary>
        public InputList<string> Tenants
        {
            get => _tenants ?? (_tenants = new InputList<string>());
            set => _tenants = value;
        }

        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public SshConnectionDeploymentTargetArgs()
        {
        }
        public static new SshConnectionDeploymentTargetArgs Empty => new SshConnectionDeploymentTargetArgs();
    }

    public sealed class SshConnectionDeploymentTargetState : global::Pulumi.ResourceArgs
    {
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        [Input("dotNetCorePlatform")]
        public Input<string>? DotNetCorePlatform { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.SshConnectionDeploymentTargetEndpointGetArgs>? _endpoints;
        public InputList<Inputs.SshConnectionDeploymentTargetEndpointGetArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.SshConnectionDeploymentTargetEndpointGetArgs>());
            set => _endpoints = value;
        }

        [Input("environments")]
        private InputList<string>? _environments;

        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        [Input("hasLatestCalamari")]
        public Input<bool>? HasLatestCalamari { get; set; }

        /// <summary>
        /// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        [Input("healthStatus")]
        public Input<string>? HealthStatus { get; set; }

        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        [Input("isInProcess")]
        public Input<bool>? IsInProcess { get; set; }

        [Input("machinePolicyId")]
        public Input<string>? MachinePolicyId { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("operatingSystem")]
        public Input<string>? OperatingSystem { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("proxyId")]
        public Input<string>? ProxyId { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("shellName")]
        public Input<string>? ShellName { get; set; }

        [Input("shellVersion")]
        public Input<string>? ShellVersion { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// A summary elaborating on the status of this resource.
        /// </summary>
        [Input("statusSummary")]
        public Input<string>? StatusSummary { get; set; }

        [Input("tenantTags")]
        private InputList<string>? _tenantTags;

        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public InputList<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new InputList<string>());
            set => _tenantTags = value;
        }

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Input("tenantedDeploymentParticipation")]
        public Input<string>? TenantedDeploymentParticipation { get; set; }

        [Input("tenants")]
        private InputList<string>? _tenants;

        /// <summary>
        /// A list of tenant IDs associated with this resource.
        /// </summary>
        public InputList<string> Tenants
        {
            get => _tenants ?? (_tenants = new InputList<string>());
            set => _tenants = value;
        }

        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public SshConnectionDeploymentTargetState()
        {
        }
        public static new SshConnectionDeploymentTargetState Empty => new SshConnectionDeploymentTargetState();
    }
}
