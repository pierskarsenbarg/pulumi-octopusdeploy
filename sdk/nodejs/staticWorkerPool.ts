// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource manages static worker pools in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = new octopusdeploy.StaticWorkerPool("example", {
 *     description: "Description for the static worker pool.",
 *     isDefault: true,
 *     sortOrder: 5,
 * });
 * ```
 */
export class StaticWorkerPool extends pulumi.CustomResource {
    /**
     * Get an existing StaticWorkerPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StaticWorkerPoolState, opts?: pulumi.CustomResourceOptions): StaticWorkerPool {
        return new StaticWorkerPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/staticWorkerPool:StaticWorkerPool';

    /**
     * Returns true if the given object is an instance of StaticWorkerPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StaticWorkerPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StaticWorkerPool.__pulumiType;
    }

    public /*out*/ readonly canAddWorkers!: pulumi.Output<boolean>;
    /**
     * The description of this static worker pool.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly isDefault!: pulumi.Output<boolean | undefined>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The order number to sort a dynamic worker pool.
     */
    public readonly sortOrder!: pulumi.Output<number>;
    /**
     * The space ID associated with this resource.
     */
    public readonly spaceId!: pulumi.Output<string>;

    /**
     * Create a StaticWorkerPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StaticWorkerPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StaticWorkerPoolArgs | StaticWorkerPoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StaticWorkerPoolState | undefined;
            resourceInputs["canAddWorkers"] = state ? state.canAddWorkers : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sortOrder"] = state ? state.sortOrder : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
        } else {
            const args = argsOrState as StaticWorkerPoolArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isDefault"] = args ? args.isDefault : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sortOrder"] = args ? args.sortOrder : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["canAddWorkers"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StaticWorkerPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StaticWorkerPool resources.
 */
export interface StaticWorkerPoolState {
    canAddWorkers?: pulumi.Input<boolean>;
    /**
     * The description of this static worker pool.
     */
    description?: pulumi.Input<string>;
    isDefault?: pulumi.Input<boolean>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The order number to sort a dynamic worker pool.
     */
    sortOrder?: pulumi.Input<number>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StaticWorkerPool resource.
 */
export interface StaticWorkerPoolArgs {
    /**
     * The description of this static worker pool.
     */
    description?: pulumi.Input<string>;
    isDefault?: pulumi.Input<boolean>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The order number to sort a dynamic worker pool.
     */
    sortOrder?: pulumi.Input<number>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
}
