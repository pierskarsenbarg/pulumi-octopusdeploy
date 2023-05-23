// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource manages GCP accounts in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = new octopusdeploy.GcpAccount("example", {
 *     jsonKey: "json-key",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/gcpAccount:GcpAccount [options] octopusdeploy_gcp_account.<name> <account-id>
 * ```
 */
export class GcpAccount extends pulumi.CustomResource {
    /**
     * Get an existing GcpAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GcpAccountState, opts?: pulumi.CustomResourceOptions): GcpAccount {
        return new GcpAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/gcpAccount:GcpAccount';

    /**
     * Returns true if the given object is an instance of GcpAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GcpAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GcpAccount.__pulumiType;
    }

    /**
     * A user-friendly description of this GCP account.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A list of environment IDs associated with this resource.
     */
    public readonly environments!: pulumi.Output<string[]>;
    /**
     * The JSON key associated with this GCP account.
     */
    public readonly jsonKey!: pulumi.Output<string>;
    /**
     * The name of this GCP account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The space ID associated with this resource.
     */
    public readonly spaceId!: pulumi.Output<string>;
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

    /**
     * Create a GcpAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GcpAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GcpAccountArgs | GcpAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GcpAccountState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["environments"] = state ? state.environments : undefined;
            resourceInputs["jsonKey"] = state ? state.jsonKey : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["tenantTags"] = state ? state.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = state ? state.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = state ? state.tenants : undefined;
        } else {
            const args = argsOrState as GcpAccountArgs | undefined;
            if ((!args || args.jsonKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jsonKey'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environments"] = args ? args.environments : undefined;
            resourceInputs["jsonKey"] = args ? args.jsonKey : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["tenantTags"] = args ? args.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = args ? args.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = args ? args.tenants : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GcpAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GcpAccount resources.
 */
export interface GcpAccountState {
    /**
     * A user-friendly description of this GCP account.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The JSON key associated with this GCP account.
     */
    jsonKey?: pulumi.Input<string>;
    /**
     * The name of this GCP account.
     */
    name?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
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
}

/**
 * The set of arguments for constructing a GcpAccount resource.
 */
export interface GcpAccountArgs {
    /**
     * A user-friendly description of this GCP account.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The JSON key associated with this GCP account.
     */
    jsonKey: pulumi.Input<string>;
    /**
     * The name of this GCP account.
     */
    name?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
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
}
