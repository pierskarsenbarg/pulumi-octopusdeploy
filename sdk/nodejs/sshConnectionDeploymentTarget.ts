// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This resource manages SSH connection deployment targets in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = new octopusdeploy.SshConnectionDeploymentTarget("example", {
 *     fingerprint: "[fingerprint]",
 *     host: "[host]",
 *     port: 22,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/sshConnectionDeploymentTarget:SshConnectionDeploymentTarget [options] octopusdeploy_ssh_connection_deployment_target.<name> <account-id>
 * ```
 */
export class SshConnectionDeploymentTarget extends pulumi.CustomResource {
    /**
     * Get an existing SshConnectionDeploymentTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SshConnectionDeploymentTargetState, opts?: pulumi.CustomResourceOptions): SshConnectionDeploymentTarget {
        return new SshConnectionDeploymentTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/sshConnectionDeploymentTarget:SshConnectionDeploymentTarget';

    /**
     * Returns true if the given object is an instance of SshConnectionDeploymentTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SshConnectionDeploymentTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SshConnectionDeploymentTarget.__pulumiType;
    }

    public readonly accountId!: pulumi.Output<string>;
    public readonly dotNetCorePlatform!: pulumi.Output<string | undefined>;
    public readonly endpoints!: pulumi.Output<outputs.SshConnectionDeploymentTargetEndpoint[]>;
    /**
     * A list of environment IDs associated with this resource.
     */
    public readonly environments!: pulumi.Output<string[]>;
    public readonly fingerprint!: pulumi.Output<string>;
    public /*out*/ readonly hasLatestCalamari!: pulumi.Output<boolean>;
    /**
     * Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    public readonly healthStatus!: pulumi.Output<string>;
    public readonly host!: pulumi.Output<string>;
    public readonly isDisabled!: pulumi.Output<boolean>;
    public /*out*/ readonly isInProcess!: pulumi.Output<boolean>;
    public readonly machinePolicyId!: pulumi.Output<string>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly operatingSystem!: pulumi.Output<string>;
    public readonly port!: pulumi.Output<number | undefined>;
    public readonly proxyId!: pulumi.Output<string | undefined>;
    public readonly roles!: pulumi.Output<string[]>;
    public readonly shellName!: pulumi.Output<string>;
    public readonly shellVersion!: pulumi.Output<string>;
    /**
     * The space ID associated with this resource.
     */
    public readonly spaceId!: pulumi.Output<string>;
    /**
     * The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * A summary elaborating on the status of this resource.
     */
    public readonly statusSummary!: pulumi.Output<string>;
    /**
     * A list of tenant tags associated with this resource.
     */
    public readonly tenantTags!: pulumi.Output<string[]>;
    /**
     * The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
     */
    public readonly tenantedDeploymentParticipation!: pulumi.Output<string>;
    /**
     * A list of tenant IDs associated with this resource.
     */
    public readonly tenants!: pulumi.Output<string[]>;
    public readonly thumbprint!: pulumi.Output<string>;
    public readonly uri!: pulumi.Output<string>;

    /**
     * Create a SshConnectionDeploymentTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SshConnectionDeploymentTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SshConnectionDeploymentTargetArgs | SshConnectionDeploymentTargetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SshConnectionDeploymentTargetState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["dotNetCorePlatform"] = state ? state.dotNetCorePlatform : undefined;
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["environments"] = state ? state.environments : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["hasLatestCalamari"] = state ? state.hasLatestCalamari : undefined;
            resourceInputs["healthStatus"] = state ? state.healthStatus : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["isDisabled"] = state ? state.isDisabled : undefined;
            resourceInputs["isInProcess"] = state ? state.isInProcess : undefined;
            resourceInputs["machinePolicyId"] = state ? state.machinePolicyId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["operatingSystem"] = state ? state.operatingSystem : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["proxyId"] = state ? state.proxyId : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["shellName"] = state ? state.shellName : undefined;
            resourceInputs["shellVersion"] = state ? state.shellVersion : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["statusSummary"] = state ? state.statusSummary : undefined;
            resourceInputs["tenantTags"] = state ? state.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = state ? state.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = state ? state.tenants : undefined;
            resourceInputs["thumbprint"] = state ? state.thumbprint : undefined;
            resourceInputs["uri"] = state ? state.uri : undefined;
        } else {
            const args = argsOrState as SshConnectionDeploymentTargetArgs | undefined;
            if ((!args || args.accountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountId'");
            }
            if ((!args || args.environments === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environments'");
            }
            if ((!args || args.fingerprint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fingerprint'");
            }
            if ((!args || args.host === undefined) && !opts.urn) {
                throw new Error("Missing required property 'host'");
            }
            if ((!args || args.roles === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roles'");
            }
            resourceInputs["accountId"] = args ? args.accountId : undefined;
            resourceInputs["dotNetCorePlatform"] = args ? args.dotNetCorePlatform : undefined;
            resourceInputs["endpoints"] = args ? args.endpoints : undefined;
            resourceInputs["environments"] = args ? args.environments : undefined;
            resourceInputs["fingerprint"] = args ? args.fingerprint : undefined;
            resourceInputs["healthStatus"] = args ? args.healthStatus : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["isDisabled"] = args ? args.isDisabled : undefined;
            resourceInputs["machinePolicyId"] = args ? args.machinePolicyId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["operatingSystem"] = args ? args.operatingSystem : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["proxyId"] = args ? args.proxyId : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["shellName"] = args ? args.shellName : undefined;
            resourceInputs["shellVersion"] = args ? args.shellVersion : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["statusSummary"] = args ? args.statusSummary : undefined;
            resourceInputs["tenantTags"] = args ? args.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = args ? args.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = args ? args.tenants : undefined;
            resourceInputs["thumbprint"] = args ? args.thumbprint : undefined;
            resourceInputs["uri"] = args ? args.uri : undefined;
            resourceInputs["hasLatestCalamari"] = undefined /*out*/;
            resourceInputs["isInProcess"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SshConnectionDeploymentTarget.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SshConnectionDeploymentTarget resources.
 */
export interface SshConnectionDeploymentTargetState {
    accountId?: pulumi.Input<string>;
    dotNetCorePlatform?: pulumi.Input<string>;
    endpoints?: pulumi.Input<pulumi.Input<inputs.SshConnectionDeploymentTargetEndpoint>[]>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    fingerprint?: pulumi.Input<string>;
    hasLatestCalamari?: pulumi.Input<boolean>;
    /**
     * Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    healthStatus?: pulumi.Input<string>;
    host?: pulumi.Input<string>;
    isDisabled?: pulumi.Input<boolean>;
    isInProcess?: pulumi.Input<boolean>;
    machinePolicyId?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    operatingSystem?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    proxyId?: pulumi.Input<string>;
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    shellName?: pulumi.Input<string>;
    shellVersion?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
     */
    status?: pulumi.Input<string>;
    /**
     * A summary elaborating on the status of this resource.
     */
    statusSummary?: pulumi.Input<string>;
    /**
     * A list of tenant tags associated with this resource.
     */
    tenantTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
     */
    tenantedDeploymentParticipation?: pulumi.Input<string>;
    /**
     * A list of tenant IDs associated with this resource.
     */
    tenants?: pulumi.Input<pulumi.Input<string>[]>;
    thumbprint?: pulumi.Input<string>;
    uri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SshConnectionDeploymentTarget resource.
 */
export interface SshConnectionDeploymentTargetArgs {
    accountId: pulumi.Input<string>;
    dotNetCorePlatform?: pulumi.Input<string>;
    endpoints?: pulumi.Input<pulumi.Input<inputs.SshConnectionDeploymentTargetEndpoint>[]>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments: pulumi.Input<pulumi.Input<string>[]>;
    fingerprint: pulumi.Input<string>;
    /**
     * Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    healthStatus?: pulumi.Input<string>;
    host: pulumi.Input<string>;
    isDisabled?: pulumi.Input<boolean>;
    machinePolicyId?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    operatingSystem?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    proxyId?: pulumi.Input<string>;
    roles: pulumi.Input<pulumi.Input<string>[]>;
    shellName?: pulumi.Input<string>;
    shellVersion?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
     */
    status?: pulumi.Input<string>;
    /**
     * A summary elaborating on the status of this resource.
     */
    statusSummary?: pulumi.Input<string>;
    /**
     * A list of tenant tags associated with this resource.
     */
    tenantTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
     */
    tenantedDeploymentParticipation?: pulumi.Input<string>;
    /**
     * A list of tenant IDs associated with this resource.
     */
    tenants?: pulumi.Input<pulumi.Input<string>[]>;
    thumbprint?: pulumi.Input<string>;
    uri?: pulumi.Input<string>;
}
