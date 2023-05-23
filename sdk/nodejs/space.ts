// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource manages spaces in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = new octopusdeploy.Space("example", {
 *     description: "A space for the development team.",
 *     isDefault: false,
 *     isTaskQueueStopped: false,
 *     spaceManagersTeamMembers: [
 *         "Users-123",
 *         "Users-321",
 *     ],
 *     spaceManagersTeams: ["teams-everyone"],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/space:Space [options] octopusdeploy_space.<name> <space-id>
 * ```
 */
export class Space extends pulumi.CustomResource {
    /**
     * Get an existing Space resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SpaceState, opts?: pulumi.CustomResourceOptions): Space {
        return new Space(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/space:Space';

    /**
     * Returns true if the given object is an instance of Space.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Space {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Space.__pulumiType;
    }

    /**
     * The description of this space.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies if this space is the default space in Octopus.
     */
    public readonly isDefault!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the status of the task queue for this space.
     */
    public readonly isTaskQueueStopped!: pulumi.Output<boolean | undefined>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The unique slug of this space.
     */
    public readonly slug!: pulumi.Output<string>;
    /**
     * A list of user IDs designated to be managers of this space.
     */
    public readonly spaceManagersTeamMembers!: pulumi.Output<string[]>;
    /**
     * A list of team IDs designated to be managers of this space.
     */
    public readonly spaceManagersTeams!: pulumi.Output<string[]>;

    /**
     * Create a Space resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SpaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SpaceArgs | SpaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SpaceState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["isTaskQueueStopped"] = state ? state.isTaskQueueStopped : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["spaceManagersTeamMembers"] = state ? state.spaceManagersTeamMembers : undefined;
            resourceInputs["spaceManagersTeams"] = state ? state.spaceManagersTeams : undefined;
        } else {
            const args = argsOrState as SpaceArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isDefault"] = args ? args.isDefault : undefined;
            resourceInputs["isTaskQueueStopped"] = args ? args.isTaskQueueStopped : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["spaceManagersTeamMembers"] = args ? args.spaceManagersTeamMembers : undefined;
            resourceInputs["spaceManagersTeams"] = args ? args.spaceManagersTeams : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Space.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Space resources.
 */
export interface SpaceState {
    /**
     * The description of this space.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies if this space is the default space in Octopus.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * Specifies the status of the task queue for this space.
     */
    isTaskQueueStopped?: pulumi.Input<boolean>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The unique slug of this space.
     */
    slug?: pulumi.Input<string>;
    /**
     * A list of user IDs designated to be managers of this space.
     */
    spaceManagersTeamMembers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of team IDs designated to be managers of this space.
     */
    spaceManagersTeams?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Space resource.
 */
export interface SpaceArgs {
    /**
     * The description of this space.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies if this space is the default space in Octopus.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * Specifies the status of the task queue for this space.
     */
    isTaskQueueStopped?: pulumi.Input<boolean>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The unique slug of this space.
     */
    slug?: pulumi.Input<string>;
    /**
     * A list of user IDs designated to be managers of this space.
     */
    spaceManagersTeamMembers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of team IDs designated to be managers of this space.
     */
    spaceManagersTeams?: pulumi.Input<pulumi.Input<string>[]>;
}
