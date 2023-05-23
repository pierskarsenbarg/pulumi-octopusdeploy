// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides information about existing library variable sets.
 */
export function getLibraryVariableSets(args?: GetLibraryVariableSetsArgs, opts?: pulumi.InvokeOptions): Promise<GetLibraryVariableSetsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("octopusdeploy:index/getLibraryVariableSets:getLibraryVariableSets", {
        "contentType": args.contentType,
        "ids": args.ids,
        "libraryVariableSets": args.libraryVariableSets,
        "partialName": args.partialName,
        "skip": args.skip,
        "take": args.take,
    }, opts);
}

/**
 * A collection of arguments for invoking getLibraryVariableSets.
 */
export interface GetLibraryVariableSetsArgs {
    /**
     * A filter to search by content type.
     */
    contentType?: string;
    /**
     * A filter to search by a list of IDs.
     */
    ids?: string[];
    /**
     * A list of library variable sets that match the filter(s).
     */
    libraryVariableSets?: inputs.GetLibraryVariableSetsLibraryVariableSet[];
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
}

/**
 * A collection of values returned by getLibraryVariableSets.
 */
export interface GetLibraryVariableSetsResult {
    /**
     * A filter to search by content type.
     */
    readonly contentType?: string;
    /**
     * An auto-generated identifier that includes the timestamp when this data source was last modified.
     */
    readonly id: string;
    /**
     * A filter to search by a list of IDs.
     */
    readonly ids?: string[];
    /**
     * A list of library variable sets that match the filter(s).
     */
    readonly libraryVariableSets: outputs.GetLibraryVariableSetsLibraryVariableSet[];
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
}

export function getLibraryVariableSetsOutput(args?: GetLibraryVariableSetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLibraryVariableSetsResult> {
    return pulumi.output(args).apply(a => getLibraryVariableSets(a, opts))
}

/**
 * A collection of arguments for invoking getLibraryVariableSets.
 */
export interface GetLibraryVariableSetsOutputArgs {
    /**
     * A filter to search by content type.
     */
    contentType?: pulumi.Input<string>;
    /**
     * A filter to search by a list of IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of library variable sets that match the filter(s).
     */
    libraryVariableSets?: pulumi.Input<pulumi.Input<inputs.GetLibraryVariableSetsLibraryVariableSetArgs>[]>;
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
}
