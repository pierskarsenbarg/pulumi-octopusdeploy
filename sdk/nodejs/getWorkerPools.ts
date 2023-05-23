// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides information about existing worker pools.
 */
export function getWorkerPools(args?: GetWorkerPoolsArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkerPoolsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("octopusdeploy:index/getWorkerPools:getWorkerPools", {
        "ids": args.ids,
        "name": args.name,
        "partialName": args.partialName,
        "skip": args.skip,
        "take": args.take,
        "workerPools": args.workerPools,
    }, opts);
}

/**
 * A collection of arguments for invoking getWorkerPools.
 */
export interface GetWorkerPoolsArgs {
    /**
     * A filter to search by a list of IDs.
     */
    ids?: string[];
    /**
     * A filter to search by name.
     */
    name?: string;
    /**
     * A filter to search by the partial match of a name.
     */
    partialName?: string;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: number;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: number;
    /**
     * A list of worker pools that match the filter(s).
     */
    workerPools?: inputs.GetWorkerPoolsWorkerPool[];
}

/**
 * A collection of values returned by getWorkerPools.
 */
export interface GetWorkerPoolsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A filter to search by a list of IDs.
     */
    readonly ids?: string[];
    /**
     * A filter to search by name.
     */
    readonly name?: string;
    /**
     * A filter to search by the partial match of a name.
     */
    readonly partialName?: string;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    readonly skip?: number;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    readonly take?: number;
    /**
     * A list of worker pools that match the filter(s).
     */
    readonly workerPools: outputs.GetWorkerPoolsWorkerPool[];
}

export function getWorkerPoolsOutput(args?: GetWorkerPoolsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWorkerPoolsResult> {
    return pulumi.output(args).apply(a => getWorkerPools(a, opts))
}

/**
 * A collection of arguments for invoking getWorkerPools.
 */
export interface GetWorkerPoolsOutputArgs {
    /**
     * A filter to search by a list of IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to search by name.
     */
    name?: pulumi.Input<string>;
    /**
     * A filter to search by the partial match of a name.
     */
    partialName?: pulumi.Input<string>;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: pulumi.Input<number>;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: pulumi.Input<number>;
    /**
     * A list of worker pools that match the filter(s).
     */
    workerPools?: pulumi.Input<pulumi.Input<inputs.GetWorkerPoolsWorkerPoolArgs>[]>;
}