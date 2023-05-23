// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This resource manages variables in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * // create an Amazon web services account variable
 * const amazonWebServicesAccountVariable = new octopusdeploy.Variable("amazon_web_services_account_variable", {
 *     ownerId: "Projects-123",
 *     type: "AmazonWebServicesAccount",
 *     value: "Accounts-123",
 * });
 * // create an Azure service principal account variable
 * const azureServicePrincipalAccountVariable = new octopusdeploy.Account("azure_service_principal_account_variable", {
 *     name: "My Azure Service Principal Account (OK to Delete)",
 *     ownerId: "Projects-123",
 *     type: "AzureAccount",
 *     value: "Accounts-123",
 * });
 * // create a Google Cloud account variable
 * const googleCloudAccountVariable = new octopusdeploy.Variable("google_cloud_account_variable", {
 *     ownerId: "Projects-123",
 *     type: "GoogleCloudAccount",
 *     value: "Accounts-123",
 * });
 * // create a Certificate variable
 * const certificateVariable = new octopusdeploy.Variable("certificate_variable", {
 *     ownerId: "Projects-123",
 *     type: "Certificate",
 *     value: "Certificates-123",
 * });
 * // create a Sensitive variable
 * const sensitiveVariable = new octopusdeploy.Variable("sensitive_variable", {
 *     isSensitive: true,
 *     ownerId: "Projects-123",
 *     sensitiveValue: "YourSecrets",
 *     type: "Sensitive",
 * });
 * // create a String variable
 * const stringVariable = new octopusdeploy.Variable("string_variable", {
 *     ownerId: "Projects-123",
 *     type: "String",
 *     value: "PlainText",
 * });
 * // create a WorkerPool variable
 * const workerpoolVariable = new octopusdeploy.Variable("workerpool_variable", {
 *     ownerId: "Projects-123",
 *     type: "WorkerPool",
 *     value: "WorkerPools-123",
 * });
 * // create a prompted variable
 * const promptedVariable = new octopusdeploy.Variable("prompted_variable", {
 *     ownerId: "Projects-123",
 *     prompt: {
 *         description: "Variable Description",
 *         isRequired: true,
 *         label: "Variable Label",
 *     },
 *     type: "String",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/variable:Variable [options] octopusdeploy_variable.<name> <variable-id>
 * ```
 */
export class Variable extends pulumi.CustomResource {
    /**
     * Get an existing Variable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VariableState, opts?: pulumi.CustomResourceOptions): Variable {
        return new Variable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/variable:Variable';

    /**
     * Returns true if the given object is an instance of Variable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Variable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Variable.__pulumiType;
    }

    /**
     * The description of this variable.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly encryptedValue!: pulumi.Output<string>;
    /**
     * Indicates whether or not this variable is considered editable.
     */
    public readonly isEditable!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether or not this resource is considered sensitive and should be kept secret.
     */
    public readonly isSensitive!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly keyFingerprint!: pulumi.Output<string>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly ownerId!: pulumi.Output<string | undefined>;
    public readonly pgpKey!: pulumi.Output<string | undefined>;
    /**
     * @deprecated This attribute is deprecated; please use owner_id instead.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    public readonly prompt!: pulumi.Output<outputs.VariablePrompt | undefined>;
    public readonly scope!: pulumi.Output<outputs.VariableScope | undefined>;
    public readonly sensitiveValue!: pulumi.Output<string | undefined>;
    /**
     * The type of variable represented by this resource. Valid types are `AmazonWebServicesAccount`, `AzureAccount`, `GoogleCloudAccount`, `Certificate`, `Sensitive`, `String`, or `WorkerPool`.
     */
    public readonly type!: pulumi.Output<string>;
    public readonly value!: pulumi.Output<string | undefined>;

    /**
     * Create a Variable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VariableArgs | VariableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VariableState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["encryptedValue"] = state ? state.encryptedValue : undefined;
            resourceInputs["isEditable"] = state ? state.isEditable : undefined;
            resourceInputs["isSensitive"] = state ? state.isSensitive : undefined;
            resourceInputs["keyFingerprint"] = state ? state.keyFingerprint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
            resourceInputs["pgpKey"] = state ? state.pgpKey : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["prompt"] = state ? state.prompt : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["sensitiveValue"] = state ? state.sensitiveValue : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as VariableArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isEditable"] = args ? args.isEditable : undefined;
            resourceInputs["isSensitive"] = args ? args.isSensitive : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ownerId"] = args ? args.ownerId : undefined;
            resourceInputs["pgpKey"] = args ? args.pgpKey : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["prompt"] = args ? args.prompt : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["sensitiveValue"] = args ? args.sensitiveValue : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["encryptedValue"] = undefined /*out*/;
            resourceInputs["keyFingerprint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Variable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Variable resources.
 */
export interface VariableState {
    /**
     * The description of this variable.
     */
    description?: pulumi.Input<string>;
    encryptedValue?: pulumi.Input<string>;
    /**
     * Indicates whether or not this variable is considered editable.
     */
    isEditable?: pulumi.Input<boolean>;
    /**
     * Indicates whether or not this resource is considered sensitive and should be kept secret.
     */
    isSensitive?: pulumi.Input<boolean>;
    keyFingerprint?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    ownerId?: pulumi.Input<string>;
    pgpKey?: pulumi.Input<string>;
    /**
     * @deprecated This attribute is deprecated; please use owner_id instead.
     */
    projectId?: pulumi.Input<string>;
    prompt?: pulumi.Input<inputs.VariablePrompt>;
    scope?: pulumi.Input<inputs.VariableScope>;
    sensitiveValue?: pulumi.Input<string>;
    /**
     * The type of variable represented by this resource. Valid types are `AmazonWebServicesAccount`, `AzureAccount`, `GoogleCloudAccount`, `Certificate`, `Sensitive`, `String`, or `WorkerPool`.
     */
    type?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Variable resource.
 */
export interface VariableArgs {
    /**
     * The description of this variable.
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates whether or not this variable is considered editable.
     */
    isEditable?: pulumi.Input<boolean>;
    /**
     * Indicates whether or not this resource is considered sensitive and should be kept secret.
     */
    isSensitive?: pulumi.Input<boolean>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    ownerId?: pulumi.Input<string>;
    pgpKey?: pulumi.Input<string>;
    /**
     * @deprecated This attribute is deprecated; please use owner_id instead.
     */
    projectId?: pulumi.Input<string>;
    prompt?: pulumi.Input<inputs.VariablePrompt>;
    scope?: pulumi.Input<inputs.VariableScope>;
    sensitiveValue?: pulumi.Input<string>;
    /**
     * The type of variable represented by this resource. Valid types are `AmazonWebServicesAccount`, `AzureAccount`, `GoogleCloudAccount`, `Certificate`, `Sensitive`, `String`, or `WorkerPool`.
     */
    type: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}
