// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource manages Azure subscription accounts in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pierskarsenbarg/octopusdeploy";
 *
 * const example = new octopusdeploy.AzureSubscriptionAccount("example", {subscriptionId: "00000000-0000-0000-0000-000000000000"});
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/azureSubscriptionAccount:AzureSubscriptionAccount [options] octopusdeploy_azure_subscription_account.<name> <account-id>
 * ```
 */
export class AzureSubscriptionAccount extends pulumi.CustomResource {
    /**
     * Get an existing AzureSubscriptionAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AzureSubscriptionAccountState, opts?: pulumi.CustomResourceOptions): AzureSubscriptionAccount {
        return new AzureSubscriptionAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/azureSubscriptionAccount:AzureSubscriptionAccount';

    /**
     * Returns true if the given object is an instance of AzureSubscriptionAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureSubscriptionAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureSubscriptionAccount.__pulumiType;
    }

    /**
     * The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
     */
    public readonly azureEnvironment!: pulumi.Output<string>;
    public readonly certificate!: pulumi.Output<string>;
    public readonly certificateThumbprint!: pulumi.Output<string>;
    /**
     * The description of this Azure subscription account.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A list of environment IDs associated with this resource.
     */
    public readonly environments!: pulumi.Output<string[]>;
    public readonly managementEndpoint!: pulumi.Output<string>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The space ID associated with this resource.
     */
    public readonly spaceId!: pulumi.Output<string>;
    /**
     * The storage endpoint suffix associated with this Azure subscription account.
     */
    public readonly storageEndpointSuffix!: pulumi.Output<string>;
    /**
     * The subscription ID of this resource.
     */
    public readonly subscriptionId!: pulumi.Output<string>;
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
     * Create a AzureSubscriptionAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureSubscriptionAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureSubscriptionAccountArgs | AzureSubscriptionAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AzureSubscriptionAccountState | undefined;
            resourceInputs["azureEnvironment"] = state ? state.azureEnvironment : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["certificateThumbprint"] = state ? state.certificateThumbprint : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["environments"] = state ? state.environments : undefined;
            resourceInputs["managementEndpoint"] = state ? state.managementEndpoint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["storageEndpointSuffix"] = state ? state.storageEndpointSuffix : undefined;
            resourceInputs["subscriptionId"] = state ? state.subscriptionId : undefined;
            resourceInputs["tenantTags"] = state ? state.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = state ? state.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = state ? state.tenants : undefined;
        } else {
            const args = argsOrState as AzureSubscriptionAccountArgs | undefined;
            if ((!args || args.managementEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managementEndpoint'");
            }
            if ((!args || args.storageEndpointSuffix === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageEndpointSuffix'");
            }
            if ((!args || args.subscriptionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subscriptionId'");
            }
            resourceInputs["azureEnvironment"] = args ? args.azureEnvironment : undefined;
            resourceInputs["certificate"] = args?.certificate ? pulumi.secret(args.certificate) : undefined;
            resourceInputs["certificateThumbprint"] = args?.certificateThumbprint ? pulumi.secret(args.certificateThumbprint) : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environments"] = args ? args.environments : undefined;
            resourceInputs["managementEndpoint"] = args ? args.managementEndpoint : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["storageEndpointSuffix"] = args ? args.storageEndpointSuffix : undefined;
            resourceInputs["subscriptionId"] = args ? args.subscriptionId : undefined;
            resourceInputs["tenantTags"] = args ? args.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = args ? args.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = args ? args.tenants : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["certificate", "certificateThumbprint"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AzureSubscriptionAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AzureSubscriptionAccount resources.
 */
export interface AzureSubscriptionAccountState {
    /**
     * The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
     */
    azureEnvironment?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    certificateThumbprint?: pulumi.Input<string>;
    /**
     * The description of this Azure subscription account.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    managementEndpoint?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The storage endpoint suffix associated with this Azure subscription account.
     */
    storageEndpointSuffix?: pulumi.Input<string>;
    /**
     * The subscription ID of this resource.
     */
    subscriptionId?: pulumi.Input<string>;
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
 * The set of arguments for constructing a AzureSubscriptionAccount resource.
 */
export interface AzureSubscriptionAccountArgs {
    /**
     * The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
     */
    azureEnvironment?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    certificateThumbprint?: pulumi.Input<string>;
    /**
     * The description of this Azure subscription account.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    managementEndpoint: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The storage endpoint suffix associated with this Azure subscription account.
     */
    storageEndpointSuffix: pulumi.Input<string>;
    /**
     * The subscription ID of this resource.
     */
    subscriptionId: pulumi.Input<string>;
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
