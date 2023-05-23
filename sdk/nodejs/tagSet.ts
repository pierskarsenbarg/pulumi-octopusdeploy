// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource manages tag sets in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = new octopusdeploy.TagSet("example", {description: "Provides tenants with access to certain early access programs."});
 * // tags are distinct resources and associated with tag sets through tag_set_id
 * const alpha = new octopusdeploy.Tag("alpha", {
 *     color: "#00FF00",
 *     tagSetId: example.id,
 * });
 * const beta = new octopusdeploy.Tag("beta", {
 *     color: "#FF0000",
 *     tagSetId: example.id,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/tagSet:TagSet [options] octopusdeploy_tag_set.<name> <tag-set-id>
 * ```
 */
export class TagSet extends pulumi.CustomResource {
    /**
     * Get an existing TagSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagSetState, opts?: pulumi.CustomResourceOptions): TagSet {
        return new TagSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/tagSet:TagSet';

    /**
     * Returns true if the given object is an instance of TagSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagSet.__pulumiType;
    }

    /**
     * The description of this tag set.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The sort order associated with this resource.
     */
    public readonly sortOrder!: pulumi.Output<number>;
    /**
     * The space ID associated with this resource.
     */
    public readonly spaceId!: pulumi.Output<string>;

    /**
     * Create a TagSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TagSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagSetArgs | TagSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagSetState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sortOrder"] = state ? state.sortOrder : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
        } else {
            const args = argsOrState as TagSetArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sortOrder"] = args ? args.sortOrder : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TagSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TagSet resources.
 */
export interface TagSetState {
    /**
     * The description of this tag set.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The sort order associated with this resource.
     */
    sortOrder?: pulumi.Input<number>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TagSet resource.
 */
export interface TagSetArgs {
    /**
     * The description of this tag set.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The sort order associated with this resource.
     */
    sortOrder?: pulumi.Input<number>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
}
