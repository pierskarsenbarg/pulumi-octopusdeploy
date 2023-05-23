// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource manages a Docker Container Registry in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pierskarsenbarg/octopusdeploy";
 *
 * const example = new octopusdeploy.DockerContainerRegistry("example", {
 *     feedUri: "https://index.docker.io",
 *     password: "test-password",
 *     registryPath: "testing/test-image",
 *     username: "test-username",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/dockerContainerRegistry:DockerContainerRegistry [options] octopusdeploy_docker_container_registry.<name> <feed-id>
 * ```
 */
export class DockerContainerRegistry extends pulumi.CustomResource {
    /**
     * Get an existing DockerContainerRegistry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DockerContainerRegistryState, opts?: pulumi.CustomResourceOptions): DockerContainerRegistry {
        return new DockerContainerRegistry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/dockerContainerRegistry:DockerContainerRegistry';

    /**
     * Returns true if the given object is an instance of DockerContainerRegistry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DockerContainerRegistry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DockerContainerRegistry.__pulumiType;
    }

    public readonly apiVersion!: pulumi.Output<string | undefined>;
    /**
     * The URL to a Maven repository. This URL is the same value defined in the `repositories`/`repository`/`url` element of a Maven `settings.xml` file.
     */
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
    public readonly registryPath!: pulumi.Output<string | undefined>;
    /**
     * The space ID associated with this resource.
     */
    public readonly spaceId!: pulumi.Output<string>;
    /**
     * The username associated with this resource.
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a DockerContainerRegistry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DockerContainerRegistryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DockerContainerRegistryArgs | DockerContainerRegistryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DockerContainerRegistryState | undefined;
            resourceInputs["apiVersion"] = state ? state.apiVersion : undefined;
            resourceInputs["feedUri"] = state ? state.feedUri : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["packageAcquisitionLocationOptions"] = state ? state.packageAcquisitionLocationOptions : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["registryPath"] = state ? state.registryPath : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as DockerContainerRegistryArgs | undefined;
            if ((!args || args.feedUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'feedUri'");
            }
            resourceInputs["apiVersion"] = args ? args.apiVersion : undefined;
            resourceInputs["feedUri"] = args ? args.feedUri : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["packageAcquisitionLocationOptions"] = args ? args.packageAcquisitionLocationOptions : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["registryPath"] = args ? args.registryPath : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["username"] = args?.username ? pulumi.secret(args.username) : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password", "username"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(DockerContainerRegistry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DockerContainerRegistry resources.
 */
export interface DockerContainerRegistryState {
    apiVersion?: pulumi.Input<string>;
    /**
     * The URL to a Maven repository. This URL is the same value defined in the `repositories`/`repository`/`url` element of a Maven `settings.xml` file.
     */
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
    registryPath?: pulumi.Input<string>;
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
 * The set of arguments for constructing a DockerContainerRegistry resource.
 */
export interface DockerContainerRegistryArgs {
    apiVersion?: pulumi.Input<string>;
    /**
     * The URL to a Maven repository. This URL is the same value defined in the `repositories`/`repository`/`url` element of a Maven `settings.xml` file.
     */
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
    registryPath?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The username associated with this resource.
     */
    username?: pulumi.Input<string>;
}
