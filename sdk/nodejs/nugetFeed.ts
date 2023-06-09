// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource manages a NuGet feed in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pierskarsenbarg/octopusdeploy";
 *
 * const example = new octopusdeploy.NugetFeed("example", {
 *     downloadAttempts: 1,
 *     downloadRetryBackoffSeconds: 30,
 *     feedUri: "https://api.nuget.org/v3/index.json",
 *     isEnhancedMode: true,
 *     password: "test-password",
 *     username: "test-username",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/nugetFeed:NugetFeed [options] octopusdeploy_nuget_feed.<name> <feed-id>
 * ```
 */
export class NugetFeed extends pulumi.CustomResource {
    /**
     * Get an existing NugetFeed resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NugetFeedState, opts?: pulumi.CustomResourceOptions): NugetFeed {
        return new NugetFeed(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/nugetFeed:NugetFeed';

    /**
     * Returns true if the given object is an instance of NugetFeed.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NugetFeed {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NugetFeed.__pulumiType;
    }

    /**
     * The number of times a deployment should attempt to download a package from this feed before failing.
     */
    public readonly downloadAttempts!: pulumi.Output<number | undefined>;
    /**
     * The number of seconds to apply as a linear back off between download attempts.
     */
    public readonly downloadRetryBackoffSeconds!: pulumi.Output<number | undefined>;
    /**
     * The feed URI can be a URL or a folder path.
     */
    public readonly feedUri!: pulumi.Output<string>;
    /**
     * This will improve performance of the NuGet feed but may not be supported by some older feeds. Disable if the operation, Create Release does not return the latest version for a package.
     */
    public readonly isEnhancedMode!: pulumi.Output<boolean | undefined>;
    /**
     * A short, memorable, unique name for this feed. Example: ACME Builds.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly packageAcquisitionLocationOptions!: pulumi.Output<string[]>;
    /**
     * The password associated with this resource.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The space ID associated with this resource.
     */
    public readonly spaceId!: pulumi.Output<string>;
    /**
     * The username associated with this resource.
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a NugetFeed resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NugetFeedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NugetFeedArgs | NugetFeedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NugetFeedState | undefined;
            resourceInputs["downloadAttempts"] = state ? state.downloadAttempts : undefined;
            resourceInputs["downloadRetryBackoffSeconds"] = state ? state.downloadRetryBackoffSeconds : undefined;
            resourceInputs["feedUri"] = state ? state.feedUri : undefined;
            resourceInputs["isEnhancedMode"] = state ? state.isEnhancedMode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["packageAcquisitionLocationOptions"] = state ? state.packageAcquisitionLocationOptions : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as NugetFeedArgs | undefined;
            if ((!args || args.feedUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'feedUri'");
            }
            resourceInputs["downloadAttempts"] = args ? args.downloadAttempts : undefined;
            resourceInputs["downloadRetryBackoffSeconds"] = args ? args.downloadRetryBackoffSeconds : undefined;
            resourceInputs["feedUri"] = args ? args.feedUri : undefined;
            resourceInputs["isEnhancedMode"] = args ? args.isEnhancedMode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["packageAcquisitionLocationOptions"] = args ? args.packageAcquisitionLocationOptions : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["username"] = args?.username ? pulumi.secret(args.username) : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password", "username"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(NugetFeed.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NugetFeed resources.
 */
export interface NugetFeedState {
    /**
     * The number of times a deployment should attempt to download a package from this feed before failing.
     */
    downloadAttempts?: pulumi.Input<number>;
    /**
     * The number of seconds to apply as a linear back off between download attempts.
     */
    downloadRetryBackoffSeconds?: pulumi.Input<number>;
    /**
     * The feed URI can be a URL or a folder path.
     */
    feedUri?: pulumi.Input<string>;
    /**
     * This will improve performance of the NuGet feed but may not be supported by some older feeds. Disable if the operation, Create Release does not return the latest version for a package.
     */
    isEnhancedMode?: pulumi.Input<boolean>;
    /**
     * A short, memorable, unique name for this feed. Example: ACME Builds.
     */
    name?: pulumi.Input<string>;
    packageAcquisitionLocationOptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The password associated with this resource.
     */
    password?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The username associated with this resource.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NugetFeed resource.
 */
export interface NugetFeedArgs {
    /**
     * The number of times a deployment should attempt to download a package from this feed before failing.
     */
    downloadAttempts?: pulumi.Input<number>;
    /**
     * The number of seconds to apply as a linear back off between download attempts.
     */
    downloadRetryBackoffSeconds?: pulumi.Input<number>;
    /**
     * The feed URI can be a URL or a folder path.
     */
    feedUri: pulumi.Input<string>;
    /**
     * This will improve performance of the NuGet feed but may not be supported by some older feeds. Disable if the operation, Create Release does not return the latest version for a package.
     */
    isEnhancedMode?: pulumi.Input<boolean>;
    /**
     * A short, memorable, unique name for this feed. Example: ACME Builds.
     */
    name?: pulumi.Input<string>;
    packageAcquisitionLocationOptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The password associated with this resource.
     */
    password?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The username associated with this resource.
     */
    username?: pulumi.Input<string>;
}
