// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource manages a GitHub repository feed in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = new octopusdeploy.GithubRepositoryFeed("example", {
 *     downloadAttempts: 1,
 *     downloadRetryBackoffSeconds: 30,
 *     feedUri: "https://api.github.com",
 *     password: "test-password",
 *     username: "test-username",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/githubRepositoryFeed:GithubRepositoryFeed [options] octopusdeploy_github_repository_feed.<name> <feed-id>
 * ```
 */
export class GithubRepositoryFeed extends pulumi.CustomResource {
    /**
     * Get an existing GithubRepositoryFeed resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GithubRepositoryFeedState, opts?: pulumi.CustomResourceOptions): GithubRepositoryFeed {
        return new GithubRepositoryFeed(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/githubRepositoryFeed:GithubRepositoryFeed';

    /**
     * Returns true if the given object is an instance of GithubRepositoryFeed.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GithubRepositoryFeed {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GithubRepositoryFeed.__pulumiType;
    }

    /**
     * The number of times a deployment should attempt to download a package from this feed before failing.
     */
    public readonly downloadAttempts!: pulumi.Output<number | undefined>;
    /**
     * The number of seconds to apply as a linear back off between download attempts.
     */
    public readonly downloadRetryBackoffSeconds!: pulumi.Output<number | undefined>;
    public readonly feedUri!: pulumi.Output<string>;
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
     * Create a GithubRepositoryFeed resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GithubRepositoryFeedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GithubRepositoryFeedArgs | GithubRepositoryFeedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GithubRepositoryFeedState | undefined;
            resourceInputs["downloadAttempts"] = state ? state.downloadAttempts : undefined;
            resourceInputs["downloadRetryBackoffSeconds"] = state ? state.downloadRetryBackoffSeconds : undefined;
            resourceInputs["feedUri"] = state ? state.feedUri : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["packageAcquisitionLocationOptions"] = state ? state.packageAcquisitionLocationOptions : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as GithubRepositoryFeedArgs | undefined;
            if ((!args || args.feedUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'feedUri'");
            }
            resourceInputs["downloadAttempts"] = args ? args.downloadAttempts : undefined;
            resourceInputs["downloadRetryBackoffSeconds"] = args ? args.downloadRetryBackoffSeconds : undefined;
            resourceInputs["feedUri"] = args ? args.feedUri : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["packageAcquisitionLocationOptions"] = args ? args.packageAcquisitionLocationOptions : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GithubRepositoryFeed.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GithubRepositoryFeed resources.
 */
export interface GithubRepositoryFeedState {
    /**
     * The number of times a deployment should attempt to download a package from this feed before failing.
     */
    downloadAttempts?: pulumi.Input<number>;
    /**
     * The number of seconds to apply as a linear back off between download attempts.
     */
    downloadRetryBackoffSeconds?: pulumi.Input<number>;
    feedUri?: pulumi.Input<string>;
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
 * The set of arguments for constructing a GithubRepositoryFeed resource.
 */
export interface GithubRepositoryFeedArgs {
    /**
     * The number of times a deployment should attempt to download a package from this feed before failing.
     */
    downloadAttempts?: pulumi.Input<number>;
    /**
     * The number of seconds to apply as a linear back off between download attempts.
     */
    downloadRetryBackoffSeconds?: pulumi.Input<number>;
    feedUri: pulumi.Input<string>;
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