// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides information about existing cloud region deployment targets.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = octopusdeploy.getCloudRegionDeploymentTargets({
 *     environments: [
 *         "Environments-123",
 *         "Environments-321",
 *     ],
 *     ids: ["Machines-123"],
 *     name: "Azure North America",
 *     partialName: "Azure Nor",
 *     skip: 5,
 *     take: 100,
 * });
 * ```
 */
export function getCloudRegionDeploymentTargets(args?: GetCloudRegionDeploymentTargetsArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudRegionDeploymentTargetsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("octopusdeploy:index/getCloudRegionDeploymentTargets:getCloudRegionDeploymentTargets", {
        "cloudRegionDeploymentTargets": args.cloudRegionDeploymentTargets,
        "deploymentId": args.deploymentId,
        "environments": args.environments,
        "healthStatuses": args.healthStatuses,
        "ids": args.ids,
        "isDisabled": args.isDisabled,
        "name": args.name,
        "partialName": args.partialName,
        "roles": args.roles,
        "shellNames": args.shellNames,
        "skip": args.skip,
        "take": args.take,
        "tenantTags": args.tenantTags,
        "tenants": args.tenants,
        "thumbprint": args.thumbprint,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudRegionDeploymentTargets.
 */
export interface GetCloudRegionDeploymentTargetsArgs {
    /**
     * A list of cloud region deployment targets that match the filter(s).
     */
    cloudRegionDeploymentTargets?: inputs.GetCloudRegionDeploymentTargetsCloudRegionDeploymentTarget[];
    /**
     * A filter to search by deployment ID.
     */
    deploymentId?: string;
    /**
     * A filter to search by a list of environment IDs.
     */
    environments?: string[];
    /**
     * A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    healthStatuses?: string[];
    /**
     * A filter to search by a list of IDs.
     */
    ids?: string[];
    /**
     * A filter to search by the disabled status of a resource.
     */
    isDisabled?: boolean;
    /**
     * A filter to search by name.
     */
    name?: string;
    /**
     * A filter to search by the partial match of a name.
     */
    partialName?: string;
    /**
     * A filter to search by a list of role IDs.
     */
    roles?: string[];
    /**
     * A list of shell names to match in the query and/or search
     */
    shellNames?: string[];
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: number;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: number;
    /**
     * A filter to search by a list of tenant tags.
     */
    tenantTags?: string[];
    /**
     * A filter to search by a list of tenant IDs.
     */
    tenants?: string[];
    /**
     * The thumbprint of the deployment target to match in the query and/or search
     */
    thumbprint?: string;
}

/**
 * A collection of values returned by getCloudRegionDeploymentTargets.
 */
export interface GetCloudRegionDeploymentTargetsResult {
    /**
     * A list of cloud region deployment targets that match the filter(s).
     */
    readonly cloudRegionDeploymentTargets: outputs.GetCloudRegionDeploymentTargetsCloudRegionDeploymentTarget[];
    /**
     * A filter to search by deployment ID.
     */
    readonly deploymentId?: string;
    /**
     * A filter to search by a list of environment IDs.
     */
    readonly environments?: string[];
    /**
     * A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    readonly healthStatuses?: string[];
    /**
     * An auto-generated identifier that includes the timestamp when this data source was last modified.
     */
    readonly id: string;
    /**
     * A filter to search by a list of IDs.
     */
    readonly ids?: string[];
    /**
     * A filter to search by the disabled status of a resource.
     */
    readonly isDisabled?: boolean;
    /**
     * A filter to search by name.
     */
    readonly name?: string;
    /**
     * A filter to search by the partial match of a name.
     */
    readonly partialName?: string;
    /**
     * A filter to search by a list of role IDs.
     */
    readonly roles?: string[];
    /**
     * A list of shell names to match in the query and/or search
     */
    readonly shellNames?: string[];
    /**
     * A filter to specify the number of items to skip in the response.
     */
    readonly skip?: number;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    readonly take?: number;
    /**
     * A filter to search by a list of tenant tags.
     */
    readonly tenantTags?: string[];
    /**
     * A filter to search by a list of tenant IDs.
     */
    readonly tenants?: string[];
    /**
     * The thumbprint of the deployment target to match in the query and/or search
     */
    readonly thumbprint?: string;
}
/**
 * Provides information about existing cloud region deployment targets.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = octopusdeploy.getCloudRegionDeploymentTargets({
 *     environments: [
 *         "Environments-123",
 *         "Environments-321",
 *     ],
 *     ids: ["Machines-123"],
 *     name: "Azure North America",
 *     partialName: "Azure Nor",
 *     skip: 5,
 *     take: 100,
 * });
 * ```
 */
export function getCloudRegionDeploymentTargetsOutput(args?: GetCloudRegionDeploymentTargetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudRegionDeploymentTargetsResult> {
    return pulumi.output(args).apply((a: any) => getCloudRegionDeploymentTargets(a, opts))
}

/**
 * A collection of arguments for invoking getCloudRegionDeploymentTargets.
 */
export interface GetCloudRegionDeploymentTargetsOutputArgs {
    /**
     * A list of cloud region deployment targets that match the filter(s).
     */
    cloudRegionDeploymentTargets?: pulumi.Input<pulumi.Input<inputs.GetCloudRegionDeploymentTargetsCloudRegionDeploymentTargetArgs>[]>;
    /**
     * A filter to search by deployment ID.
     */
    deploymentId?: pulumi.Input<string>;
    /**
     * A filter to search by a list of environment IDs.
     */
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    healthStatuses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to search by a list of IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to search by the disabled status of a resource.
     */
    isDisabled?: pulumi.Input<boolean>;
    /**
     * A filter to search by name.
     */
    name?: pulumi.Input<string>;
    /**
     * A filter to search by the partial match of a name.
     */
    partialName?: pulumi.Input<string>;
    /**
     * A filter to search by a list of role IDs.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of shell names to match in the query and/or search
     */
    shellNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: pulumi.Input<number>;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: pulumi.Input<number>;
    /**
     * A filter to search by a list of tenant tags.
     */
    tenantTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to search by a list of tenant IDs.
     */
    tenants?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The thumbprint of the deployment target to match in the query and/or search
     */
    thumbprint?: pulumi.Input<string>;
}
