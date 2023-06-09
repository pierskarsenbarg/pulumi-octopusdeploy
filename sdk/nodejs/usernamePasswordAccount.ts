// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource manages username-password accounts in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pierskarsenbarg/octopusdeploy";
 *
 * const example = new octopusdeploy.UsernamePasswordAccount("example", {
 *     password: "###########",
 *     username: "[username]",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/usernamePasswordAccount:UsernamePasswordAccount [options] octopusdeploy_username_password_account.<name> <account-id>
 * ```
 */
export class UsernamePasswordAccount extends pulumi.CustomResource {
    /**
     * Get an existing UsernamePasswordAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UsernamePasswordAccountState, opts?: pulumi.CustomResourceOptions): UsernamePasswordAccount {
        return new UsernamePasswordAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/usernamePasswordAccount:UsernamePasswordAccount';

    /**
     * Returns true if the given object is an instance of UsernamePasswordAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UsernamePasswordAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UsernamePasswordAccount.__pulumiType;
    }

    /**
     * The description of this username/password account.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A list of environment IDs associated with this resource.
     */
    public readonly environments!: pulumi.Output<string[]>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The password associated with this resource.
     */
    public readonly password!: pulumi.Output<string | undefined>;
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
     * The username associated with this resource.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a UsernamePasswordAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UsernamePasswordAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UsernamePasswordAccountArgs | UsernamePasswordAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UsernamePasswordAccountState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["environments"] = state ? state.environments : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["tenantTags"] = state ? state.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = state ? state.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = state ? state.tenants : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as UsernamePasswordAccountArgs | undefined;
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environments"] = args ? args.environments : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["tenantTags"] = args ? args.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = args ? args.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = args ? args.tenants : undefined;
            resourceInputs["username"] = args?.username ? pulumi.secret(args.username) : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password", "username"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(UsernamePasswordAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UsernamePasswordAccount resources.
 */
export interface UsernamePasswordAccountState {
    /**
     * The description of this username/password account.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The password associated with this resource.
     */
    password?: pulumi.Input<string>;
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
    /**
     * The username associated with this resource.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UsernamePasswordAccount resource.
 */
export interface UsernamePasswordAccountArgs {
    /**
     * The description of this username/password account.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The password associated with this resource.
     */
    password?: pulumi.Input<string>;
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
    /**
     * The username associated with this resource.
     */
    username: pulumi.Input<string>;
}
