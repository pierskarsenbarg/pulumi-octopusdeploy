// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource manages Azure service fabric cluster deployment targets in Octopus Deploy.
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/azureServiceFabricClusterDeploymentTarget:AzureServiceFabricClusterDeploymentTarget [options] octopusdeploy_azure_service_fabric_cluster_deployment_target.<name> <machine-id>
 * ```
 */
export class AzureServiceFabricClusterDeploymentTarget extends pulumi.CustomResource {
    /**
     * Get an existing AzureServiceFabricClusterDeploymentTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AzureServiceFabricClusterDeploymentTargetState, opts?: pulumi.CustomResourceOptions): AzureServiceFabricClusterDeploymentTarget {
        return new AzureServiceFabricClusterDeploymentTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/azureServiceFabricClusterDeploymentTarget:AzureServiceFabricClusterDeploymentTarget';

    /**
     * Returns true if the given object is an instance of AzureServiceFabricClusterDeploymentTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureServiceFabricClusterDeploymentTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureServiceFabricClusterDeploymentTarget.__pulumiType;
    }

    public readonly aadClientCredentialSecret!: pulumi.Output<string>;
    public readonly aadCredentialType!: pulumi.Output<string>;
    public readonly aadUserCredentialPassword!: pulumi.Output<string | undefined>;
    public readonly aadUserCredentialUsername!: pulumi.Output<string>;
    public readonly certificateStoreLocation!: pulumi.Output<string>;
    public readonly certificateStoreName!: pulumi.Output<string>;
    public readonly clientCertificateVariable!: pulumi.Output<string>;
    public readonly connectionEndpoint!: pulumi.Output<string>;
    public readonly endpoints!: pulumi.Output<outputs.AzureServiceFabricClusterDeploymentTargetEndpoint[]>;
    /**
     * A list of environment IDs associated with this resource.
     */
    public readonly environments!: pulumi.Output<string[]>;
    public /*out*/ readonly hasLatestCalamari!: pulumi.Output<boolean>;
    /**
     * Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    public readonly healthStatus!: pulumi.Output<string>;
    public readonly isDisabled!: pulumi.Output<boolean>;
    public /*out*/ readonly isInProcess!: pulumi.Output<boolean>;
    public readonly machinePolicyId!: pulumi.Output<string>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly operatingSystem!: pulumi.Output<string>;
    public readonly roles!: pulumi.Output<string[]>;
    public readonly securityMode!: pulumi.Output<string>;
    public readonly serverCertificateThumbprint!: pulumi.Output<string>;
    public readonly shellName!: pulumi.Output<string>;
    public readonly shellVersion!: pulumi.Output<string>;
    /**
     * The space ID associated with this resource.
     */
    public readonly spaceId!: pulumi.Output<string>;
    /**
     * The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * A summary elaborating on the status of this resource.
     */
    public readonly statusSummary!: pulumi.Output<string>;
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
    public readonly thumbprint!: pulumi.Output<string>;
    public readonly uri!: pulumi.Output<string>;

    /**
     * Create a AzureServiceFabricClusterDeploymentTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureServiceFabricClusterDeploymentTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureServiceFabricClusterDeploymentTargetArgs | AzureServiceFabricClusterDeploymentTargetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AzureServiceFabricClusterDeploymentTargetState | undefined;
            resourceInputs["aadClientCredentialSecret"] = state ? state.aadClientCredentialSecret : undefined;
            resourceInputs["aadCredentialType"] = state ? state.aadCredentialType : undefined;
            resourceInputs["aadUserCredentialPassword"] = state ? state.aadUserCredentialPassword : undefined;
            resourceInputs["aadUserCredentialUsername"] = state ? state.aadUserCredentialUsername : undefined;
            resourceInputs["certificateStoreLocation"] = state ? state.certificateStoreLocation : undefined;
            resourceInputs["certificateStoreName"] = state ? state.certificateStoreName : undefined;
            resourceInputs["clientCertificateVariable"] = state ? state.clientCertificateVariable : undefined;
            resourceInputs["connectionEndpoint"] = state ? state.connectionEndpoint : undefined;
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["environments"] = state ? state.environments : undefined;
            resourceInputs["hasLatestCalamari"] = state ? state.hasLatestCalamari : undefined;
            resourceInputs["healthStatus"] = state ? state.healthStatus : undefined;
            resourceInputs["isDisabled"] = state ? state.isDisabled : undefined;
            resourceInputs["isInProcess"] = state ? state.isInProcess : undefined;
            resourceInputs["machinePolicyId"] = state ? state.machinePolicyId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["operatingSystem"] = state ? state.operatingSystem : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["securityMode"] = state ? state.securityMode : undefined;
            resourceInputs["serverCertificateThumbprint"] = state ? state.serverCertificateThumbprint : undefined;
            resourceInputs["shellName"] = state ? state.shellName : undefined;
            resourceInputs["shellVersion"] = state ? state.shellVersion : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["statusSummary"] = state ? state.statusSummary : undefined;
            resourceInputs["tenantTags"] = state ? state.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = state ? state.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = state ? state.tenants : undefined;
            resourceInputs["thumbprint"] = state ? state.thumbprint : undefined;
            resourceInputs["uri"] = state ? state.uri : undefined;
        } else {
            const args = argsOrState as AzureServiceFabricClusterDeploymentTargetArgs | undefined;
            if ((!args || args.connectionEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionEndpoint'");
            }
            if ((!args || args.environments === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environments'");
            }
            if ((!args || args.roles === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roles'");
            }
            resourceInputs["aadClientCredentialSecret"] = args ? args.aadClientCredentialSecret : undefined;
            resourceInputs["aadCredentialType"] = args ? args.aadCredentialType : undefined;
            resourceInputs["aadUserCredentialPassword"] = args?.aadUserCredentialPassword ? pulumi.secret(args.aadUserCredentialPassword) : undefined;
            resourceInputs["aadUserCredentialUsername"] = args ? args.aadUserCredentialUsername : undefined;
            resourceInputs["certificateStoreLocation"] = args ? args.certificateStoreLocation : undefined;
            resourceInputs["certificateStoreName"] = args ? args.certificateStoreName : undefined;
            resourceInputs["clientCertificateVariable"] = args ? args.clientCertificateVariable : undefined;
            resourceInputs["connectionEndpoint"] = args ? args.connectionEndpoint : undefined;
            resourceInputs["endpoints"] = args ? args.endpoints : undefined;
            resourceInputs["environments"] = args ? args.environments : undefined;
            resourceInputs["healthStatus"] = args ? args.healthStatus : undefined;
            resourceInputs["isDisabled"] = args ? args.isDisabled : undefined;
            resourceInputs["machinePolicyId"] = args ? args.machinePolicyId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["operatingSystem"] = args ? args.operatingSystem : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["securityMode"] = args ? args.securityMode : undefined;
            resourceInputs["serverCertificateThumbprint"] = args ? args.serverCertificateThumbprint : undefined;
            resourceInputs["shellName"] = args ? args.shellName : undefined;
            resourceInputs["shellVersion"] = args ? args.shellVersion : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["statusSummary"] = args ? args.statusSummary : undefined;
            resourceInputs["tenantTags"] = args ? args.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = args ? args.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = args ? args.tenants : undefined;
            resourceInputs["thumbprint"] = args ? args.thumbprint : undefined;
            resourceInputs["uri"] = args ? args.uri : undefined;
            resourceInputs["hasLatestCalamari"] = undefined /*out*/;
            resourceInputs["isInProcess"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["aadUserCredentialPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AzureServiceFabricClusterDeploymentTarget.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AzureServiceFabricClusterDeploymentTarget resources.
 */
export interface AzureServiceFabricClusterDeploymentTargetState {
    aadClientCredentialSecret?: pulumi.Input<string>;
    aadCredentialType?: pulumi.Input<string>;
    aadUserCredentialPassword?: pulumi.Input<string>;
    aadUserCredentialUsername?: pulumi.Input<string>;
    certificateStoreLocation?: pulumi.Input<string>;
    certificateStoreName?: pulumi.Input<string>;
    clientCertificateVariable?: pulumi.Input<string>;
    connectionEndpoint?: pulumi.Input<string>;
    endpoints?: pulumi.Input<pulumi.Input<inputs.AzureServiceFabricClusterDeploymentTargetEndpoint>[]>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    hasLatestCalamari?: pulumi.Input<boolean>;
    /**
     * Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    healthStatus?: pulumi.Input<string>;
    isDisabled?: pulumi.Input<boolean>;
    isInProcess?: pulumi.Input<boolean>;
    machinePolicyId?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    operatingSystem?: pulumi.Input<string>;
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    securityMode?: pulumi.Input<string>;
    serverCertificateThumbprint?: pulumi.Input<string>;
    shellName?: pulumi.Input<string>;
    shellVersion?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
     */
    status?: pulumi.Input<string>;
    /**
     * A summary elaborating on the status of this resource.
     */
    statusSummary?: pulumi.Input<string>;
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
    thumbprint?: pulumi.Input<string>;
    uri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AzureServiceFabricClusterDeploymentTarget resource.
 */
export interface AzureServiceFabricClusterDeploymentTargetArgs {
    aadClientCredentialSecret?: pulumi.Input<string>;
    aadCredentialType?: pulumi.Input<string>;
    aadUserCredentialPassword?: pulumi.Input<string>;
    aadUserCredentialUsername?: pulumi.Input<string>;
    certificateStoreLocation?: pulumi.Input<string>;
    certificateStoreName?: pulumi.Input<string>;
    clientCertificateVariable?: pulumi.Input<string>;
    connectionEndpoint: pulumi.Input<string>;
    endpoints?: pulumi.Input<pulumi.Input<inputs.AzureServiceFabricClusterDeploymentTargetEndpoint>[]>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    healthStatus?: pulumi.Input<string>;
    isDisabled?: pulumi.Input<boolean>;
    machinePolicyId?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    operatingSystem?: pulumi.Input<string>;
    roles: pulumi.Input<pulumi.Input<string>[]>;
    securityMode?: pulumi.Input<string>;
    serverCertificateThumbprint?: pulumi.Input<string>;
    shellName?: pulumi.Input<string>;
    shellVersion?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
     */
    status?: pulumi.Input<string>;
    /**
     * A summary elaborating on the status of this resource.
     */
    statusSummary?: pulumi.Input<string>;
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
    thumbprint?: pulumi.Input<string>;
    uri?: pulumi.Input<string>;
}
