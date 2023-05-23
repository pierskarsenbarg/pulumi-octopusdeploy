// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource manages an AWS Elastic Container Registry in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = new octopusdeploy.AwsElasticContainerRegistry("example", {
 *     accessKey: "access-key",
 *     region: "us-east-1",
 *     secretKey: "secret-key",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry [options] octopusdeploy_aws_elastic_container_registry.<name> <feed-id>
 * ```
 */
export class AwsElasticContainerRegistry extends pulumi.CustomResource {
    /**
     * Get an existing AwsElasticContainerRegistry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AwsElasticContainerRegistryState, opts?: pulumi.CustomResourceOptions): AwsElasticContainerRegistry {
        return new AwsElasticContainerRegistry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry';

    /**
     * Returns true if the given object is an instance of AwsElasticContainerRegistry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AwsElasticContainerRegistry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AwsElasticContainerRegistry.__pulumiType;
    }

    /**
     * The AWS access key to use when authenticating against Amazon Web Services.
     */
    public readonly accessKey!: pulumi.Output<string>;
    /**
     * A short, memorable, unique name for this feed. Example: ACME Builds.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly packageAcquisitionLocationOptions!: pulumi.Output<string[]>;
    /**
     * The AWS region where the registry resides.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The AWS secret key to use when authenticating against Amazon Web Services.
     */
    public readonly secretKey!: pulumi.Output<string>;
    /**
     * The space ID associated with this feed.
     */
    public readonly spaceId!: pulumi.Output<string>;

    /**
     * Create a AwsElasticContainerRegistry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AwsElasticContainerRegistryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AwsElasticContainerRegistryArgs | AwsElasticContainerRegistryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AwsElasticContainerRegistryState | undefined;
            resourceInputs["accessKey"] = state ? state.accessKey : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["packageAcquisitionLocationOptions"] = state ? state.packageAcquisitionLocationOptions : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
        } else {
            const args = argsOrState as AwsElasticContainerRegistryArgs | undefined;
            if ((!args || args.accessKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessKey'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.secretKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretKey'");
            }
            resourceInputs["accessKey"] = args ? args.accessKey : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["packageAcquisitionLocationOptions"] = args ? args.packageAcquisitionLocationOptions : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretKey"] = args?.secretKey ? pulumi.secret(args.secretKey) : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AwsElasticContainerRegistry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AwsElasticContainerRegistry resources.
 */
export interface AwsElasticContainerRegistryState {
    /**
     * The AWS access key to use when authenticating against Amazon Web Services.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * A short, memorable, unique name for this feed. Example: ACME Builds.
     */
    name?: pulumi.Input<string>;
    packageAcquisitionLocationOptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The AWS region where the registry resides.
     */
    region?: pulumi.Input<string>;
    /**
     * The AWS secret key to use when authenticating against Amazon Web Services.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * The space ID associated with this feed.
     */
    spaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AwsElasticContainerRegistry resource.
 */
export interface AwsElasticContainerRegistryArgs {
    /**
     * The AWS access key to use when authenticating against Amazon Web Services.
     */
    accessKey: pulumi.Input<string>;
    /**
     * A short, memorable, unique name for this feed. Example: ACME Builds.
     */
    name?: pulumi.Input<string>;
    packageAcquisitionLocationOptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The AWS region where the registry resides.
     */
    region: pulumi.Input<string>;
    /**
     * The AWS secret key to use when authenticating against Amazon Web Services.
     */
    secretKey: pulumi.Input<string>;
    /**
     * The space ID associated with this feed.
     */
    spaceId?: pulumi.Input<string>;
}
