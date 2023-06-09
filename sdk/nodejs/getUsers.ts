// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides information about existing users.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = octopusdeploy.getUsers({
 *     ids: [
 *         "Users-123",
 *         "Users-321",
 *     ],
 *     skip: 5,
 *     take: 100,
 * });
 * ```
 */
export function getUsers(args?: GetUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("octopusdeploy:index/getUsers:getUsers", {
        "filter": args.filter,
        "ids": args.ids,
        "skip": args.skip,
        "take": args.take,
        "users": args.users,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersArgs {
    /**
     * A filter with which to search.
     */
    filter?: string;
    /**
     * A filter to search by a list of IDs.
     */
    ids?: string[];
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: number;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: number;
    /**
     * A list of users that match the filter(s).
     */
    users?: inputs.GetUsersUser[];
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
    /**
     * A filter with which to search.
     */
    readonly filter?: string;
    /**
     * An auto-generated identifier that includes the timestamp when this data source was last modified.
     */
    readonly id: string;
    /**
     * A filter to search by a list of IDs.
     */
    readonly ids?: string[];
    /**
     * A filter to specify the number of items to skip in the response.
     */
    readonly skip?: number;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    readonly take?: number;
    /**
     * A list of users that match the filter(s).
     */
    readonly users: outputs.GetUsersUser[];
}
/**
 * Provides information about existing users.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = octopusdeploy.getUsers({
 *     ids: [
 *         "Users-123",
 *         "Users-321",
 *     ],
 *     skip: 5,
 *     take: 100,
 * });
 * ```
 */
export function getUsersOutput(args?: GetUsersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUsersResult> {
    return pulumi.output(args).apply((a: any) => getUsers(a, opts))
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersOutputArgs {
    /**
     * A filter with which to search.
     */
    filter?: pulumi.Input<string>;
    /**
     * A filter to search by a list of IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: pulumi.Input<number>;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: pulumi.Input<number>;
    /**
     * A list of users that match the filter(s).
     */
    users?: pulumi.Input<pulumi.Input<inputs.GetUsersUserArgs>[]>;
}
