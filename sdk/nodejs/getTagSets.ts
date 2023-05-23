// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides information about existing tag sets.
 */
export function getTagSets(args?: GetTagSetsArgs, opts?: pulumi.InvokeOptions): Promise<GetTagSetsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("octopusdeploy:index/getTagSets:getTagSets", {
        "ids": args.ids,
        "partialName": args.partialName,
        "skip": args.skip,
        "tagSets": args.tagSets,
        "take": args.take,
    }, opts);
}

/**
 * A collection of arguments for invoking getTagSets.
 */
export interface GetTagSetsArgs {
    /**
     * A filter to search by a list of IDs.
     */
    ids?: string[];
    /**
     * A filter to search by the partial match of a name.
     */
    partialName?: string;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: number;
    /**
     * A list of tag sets that match the filter(s).
     */
    tagSets?: inputs.GetTagSetsTagSet[];
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: number;
}

/**
 * A collection of values returned by getTagSets.
 */
export interface GetTagSetsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A filter to search by a list of IDs.
     */
    readonly ids?: string[];
    /**
     * A filter to search by the partial match of a name.
     */
    readonly partialName?: string;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    readonly skip?: number;
    /**
     * A list of tag sets that match the filter(s).
     */
    readonly tagSets: outputs.GetTagSetsTagSet[];
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    readonly take?: number;
}

export function getTagSetsOutput(args?: GetTagSetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTagSetsResult> {
    return pulumi.output(args).apply(a => getTagSets(a, opts))
}

/**
 * A collection of arguments for invoking getTagSets.
 */
export interface GetTagSetsOutputArgs {
    /**
     * A filter to search by a list of IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to search by the partial match of a name.
     */
    partialName?: pulumi.Input<string>;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: pulumi.Input<number>;
    /**
     * A list of tag sets that match the filter(s).
     */
    tagSets?: pulumi.Input<pulumi.Input<inputs.GetTagSetsTagSetArgs>[]>;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: pulumi.Input<number>;
}
