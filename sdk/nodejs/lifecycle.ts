// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource manages lifecycles in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pierskarsenbarg/octopusdeploy";
 *
 * const example = new octopusdeploy.Lifecycle("example", {
 *     description: "This is the default lifecycle.",
 *     phases: [
 *         {
 *             automaticDeploymentTargets: ["Environments-321"],
 *             name: "foo",
 *             releaseRetentionPolicy: {
 *                 quantityToKeep: 1,
 *                 shouldKeepForever: true,
 *                 unit: "Days",
 *             },
 *             tentacleRetentionPolicy: {
 *                 quantityToKeep: 30,
 *                 shouldKeepForever: false,
 *                 unit: "Items",
 *             },
 *         },
 *         {
 *             isOptionalPhase: true,
 *             name: "bar",
 *             optionalDeploymentTargets: ["Environments-321"],
 *         },
 *     ],
 *     releaseRetentionPolicy: {
 *         quantityToKeep: 1,
 *         shouldKeepForever: true,
 *         unit: "Days",
 *     },
 *     tentacleRetentionPolicy: {
 *         quantityToKeep: 30,
 *         shouldKeepForever: false,
 *         unit: "Items",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/lifecycle:Lifecycle [options] octopusdeploy_lifecycle.<name> <lifecycle-id>
 * ```
 */
export class Lifecycle extends pulumi.CustomResource {
    /**
     * Get an existing Lifecycle resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LifecycleState, opts?: pulumi.CustomResourceOptions): Lifecycle {
        return new Lifecycle(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/lifecycle:Lifecycle';

    /**
     * Returns true if the given object is an instance of Lifecycle.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Lifecycle {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Lifecycle.__pulumiType;
    }

    /**
     * The description of this lifecycle.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly phases!: pulumi.Output<outputs.LifecyclePhase[]>;
    public readonly releaseRetentionPolicy!: pulumi.Output<outputs.LifecycleReleaseRetentionPolicy>;
    /**
     * The space ID associated with this resource.
     */
    public readonly spaceId!: pulumi.Output<string>;
    public readonly tentacleRetentionPolicy!: pulumi.Output<outputs.LifecycleTentacleRetentionPolicy>;

    /**
     * Create a Lifecycle resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LifecycleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LifecycleArgs | LifecycleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LifecycleState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["phases"] = state ? state.phases : undefined;
            resourceInputs["releaseRetentionPolicy"] = state ? state.releaseRetentionPolicy : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["tentacleRetentionPolicy"] = state ? state.tentacleRetentionPolicy : undefined;
        } else {
            const args = argsOrState as LifecycleArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["phases"] = args ? args.phases : undefined;
            resourceInputs["releaseRetentionPolicy"] = args ? args.releaseRetentionPolicy : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["tentacleRetentionPolicy"] = args ? args.tentacleRetentionPolicy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Lifecycle.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Lifecycle resources.
 */
export interface LifecycleState {
    /**
     * The description of this lifecycle.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    phases?: pulumi.Input<pulumi.Input<inputs.LifecyclePhase>[]>;
    releaseRetentionPolicy?: pulumi.Input<inputs.LifecycleReleaseRetentionPolicy>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    tentacleRetentionPolicy?: pulumi.Input<inputs.LifecycleTentacleRetentionPolicy>;
}

/**
 * The set of arguments for constructing a Lifecycle resource.
 */
export interface LifecycleArgs {
    /**
     * The description of this lifecycle.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    phases?: pulumi.Input<pulumi.Input<inputs.LifecyclePhase>[]>;
    releaseRetentionPolicy?: pulumi.Input<inputs.LifecycleReleaseRetentionPolicy>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    tentacleRetentionPolicy?: pulumi.Input<inputs.LifecycleTentacleRetentionPolicy>;
}
