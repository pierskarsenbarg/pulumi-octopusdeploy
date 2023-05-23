// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This resource manages polling tentacle deployment targets in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = new octopusdeploy.ListeningTentacleDeploymentTarget("example", {
 *     environments: [
 *         "Environments-123",
 *         "Environment-321",
 *     ],
 *     isDisabled: true,
 *     machinePolicyId: "MachinePolicy-123",
 *     roles: [
 *         "Development Team",
 *         "System Administrators",
 *     ],
 *     tenantedDeploymentParticipation: "Untenanted",
 *     tentacleUrl: "https://example.com:1234/",
 *     thumbprint: "<thumbprint>",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/pollingTentacleDeploymentTarget:PollingTentacleDeploymentTarget [options] octopusdeploy_listening_tentacle_deployment_target.<name> <machine-id>
 * ```
 */
export class PollingTentacleDeploymentTarget extends pulumi.CustomResource {
    /**
     * Get an existing PollingTentacleDeploymentTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PollingTentacleDeploymentTargetState, opts?: pulumi.CustomResourceOptions): PollingTentacleDeploymentTarget {
        return new PollingTentacleDeploymentTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/pollingTentacleDeploymentTarget:PollingTentacleDeploymentTarget';

    /**
     * Returns true if the given object is an instance of PollingTentacleDeploymentTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PollingTentacleDeploymentTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PollingTentacleDeploymentTarget.__pulumiType;
    }

    public readonly certificateSignatureAlgorithm!: pulumi.Output<string | undefined>;
    public readonly endpoints!: pulumi.Output<outputs.PollingTentacleDeploymentTargetEndpoint[]>;
    /**
     * A list of environment IDs associated with this resource.
     */
    public readonly environments!: pulumi.Output<string[]>;
    public /*out*/ readonly hasLatestCalamari!: pulumi.Output<boolean>;
    /**
     * Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    public readonly healthStatus!: pulumi.Output<string>;
    public readonly isDisabled!: pulumi.Output<boolean>;
    public /*out*/ readonly isInProcess!: pulumi.Output<boolean>;
    public readonly machinePolicyId!: pulumi.Output<string>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly operatingSystem!: pulumi.Output<string>;
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
    public readonly tentacleUrl!: pulumi.Output<string>;
    public readonly tentacleVersionDetails!: pulumi.Output<outputs.PollingTentacleDeploymentTargetTentacleVersionDetail[]>;
    public readonly thumbprint!: pulumi.Output<string>;
    public readonly uri!: pulumi.Output<string>;

    /**
     * Create a PollingTentacleDeploymentTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PollingTentacleDeploymentTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PollingTentacleDeploymentTargetArgs | PollingTentacleDeploymentTargetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PollingTentacleDeploymentTargetState | undefined;
            resourceInputs["certificateSignatureAlgorithm"] = state ? state.certificateSignatureAlgorithm : undefined;
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["environments"] = state ? state.environments : undefined;
            resourceInputs["hasLatestCalamari"] = state ? state.hasLatestCalamari : undefined;
            resourceInputs["healthStatus"] = state ? state.healthStatus : undefined;
            resourceInputs["isDisabled"] = state ? state.isDisabled : undefined;
            resourceInputs["isInProcess"] = state ? state.isInProcess : undefined;
            resourceInputs["machinePolicyId"] = state ? state.machinePolicyId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["operatingSystem"] = state ? state.operatingSystem : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["shellName"] = state ? state.shellName : undefined;
            resourceInputs["shellVersion"] = state ? state.shellVersion : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["statusSummary"] = state ? state.statusSummary : undefined;
            resourceInputs["tenantTags"] = state ? state.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = state ? state.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = state ? state.tenants : undefined;
            resourceInputs["tentacleUrl"] = state ? state.tentacleUrl : undefined;
            resourceInputs["tentacleVersionDetails"] = state ? state.tentacleVersionDetails : undefined;
            resourceInputs["thumbprint"] = state ? state.thumbprint : undefined;
            resourceInputs["uri"] = state ? state.uri : undefined;
        } else {
            const args = argsOrState as PollingTentacleDeploymentTargetArgs | undefined;
            if ((!args || args.environments === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environments'");
            }
            if ((!args || args.roles === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roles'");
            }
            if ((!args || args.tentacleUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tentacleUrl'");
            }
            resourceInputs["certificateSignatureAlgorithm"] = args ? args.certificateSignatureAlgorithm : undefined;
            resourceInputs["endpoints"] = args ? args.endpoints : undefined;
            resourceInputs["environments"] = args ? args.environments : undefined;
            resourceInputs["healthStatus"] = args ? args.healthStatus : undefined;
            resourceInputs["isDisabled"] = args ? args.isDisabled : undefined;
            resourceInputs["machinePolicyId"] = args ? args.machinePolicyId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["operatingSystem"] = args ? args.operatingSystem : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["shellName"] = args ? args.shellName : undefined;
            resourceInputs["shellVersion"] = args ? args.shellVersion : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["statusSummary"] = args ? args.statusSummary : undefined;
            resourceInputs["tenantTags"] = args ? args.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = args ? args.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = args ? args.tenants : undefined;
            resourceInputs["tentacleUrl"] = args ? args.tentacleUrl : undefined;
            resourceInputs["tentacleVersionDetails"] = args ? args.tentacleVersionDetails : undefined;
            resourceInputs["thumbprint"] = args ? args.thumbprint : undefined;
            resourceInputs["uri"] = args ? args.uri : undefined;
            resourceInputs["hasLatestCalamari"] = undefined /*out*/;
            resourceInputs["isInProcess"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PollingTentacleDeploymentTarget.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PollingTentacleDeploymentTarget resources.
 */
export interface PollingTentacleDeploymentTargetState {
    certificateSignatureAlgorithm?: pulumi.Input<string>;
    endpoints?: pulumi.Input<pulumi.Input<inputs.PollingTentacleDeploymentTargetEndpoint>[]>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    hasLatestCalamari?: pulumi.Input<boolean>;
    /**
     * Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    healthStatus?: pulumi.Input<string>;
    isDisabled?: pulumi.Input<boolean>;
    isInProcess?: pulumi.Input<boolean>;
    machinePolicyId?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    operatingSystem?: pulumi.Input<string>;
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
    tentacleUrl?: pulumi.Input<string>;
    tentacleVersionDetails?: pulumi.Input<pulumi.Input<inputs.PollingTentacleDeploymentTargetTentacleVersionDetail>[]>;
    thumbprint?: pulumi.Input<string>;
    uri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PollingTentacleDeploymentTarget resource.
 */
export interface PollingTentacleDeploymentTargetArgs {
    certificateSignatureAlgorithm?: pulumi.Input<string>;
    endpoints?: pulumi.Input<pulumi.Input<inputs.PollingTentacleDeploymentTargetEndpoint>[]>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    healthStatus?: pulumi.Input<string>;
    isDisabled?: pulumi.Input<boolean>;
    machinePolicyId?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    operatingSystem?: pulumi.Input<string>;
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
    tentacleUrl: pulumi.Input<string>;
    tentacleVersionDetails?: pulumi.Input<pulumi.Input<inputs.PollingTentacleDeploymentTargetTentacleVersionDetail>[]>;
    thumbprint?: pulumi.Input<string>;
    uri?: pulumi.Input<string>;
}