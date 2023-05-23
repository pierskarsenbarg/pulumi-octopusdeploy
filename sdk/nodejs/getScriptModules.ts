// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides information about existing script modules.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = pulumi.output(octopusdeploy.getScriptModules({
 *     ids: [
 *         "LibraryVariableSets-123",
 *         "LibraryVariableSets-321",
 *     ],
 *     partialName: "Defau",
 *     skip: 5,
 *     take: 100,
 * }));
 * ```
 */
export function getScriptModules(args?: GetScriptModulesArgs, opts?: pulumi.InvokeOptions): Promise<GetScriptModulesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("octopusdeploy:index/getScriptModules:getScriptModules", {
        "ids": args.ids,
        "partialName": args.partialName,
        "scriptModules": args.scriptModules,
        "skip": args.skip,
        "take": args.take,
    }, opts);
}

/**
 * A collection of arguments for invoking getScriptModules.
 */
export interface GetScriptModulesArgs {
    /**
     * A filter to search by a list of IDs.
     */
    ids?: string[];
    /**
     * A filter to search by the partial match of a name.
     */
    partialName?: string;
    /**
     * A list of script modules that match the filter(s).
     */
    scriptModules?: inputs.GetScriptModulesScriptModule[];
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
 * A collection of values returned by getScriptModules.
 */
export interface GetScriptModulesResult {
    /**
     * An auto-generated identifier that includes the timestamp when this data source was last modified.
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
     * A list of script modules that match the filter(s).
     */
    readonly scriptModules: outputs.GetScriptModulesScriptModule[];
    /**
     * A filter to specify the number of items to skip in the response.
     */
    readonly skip?: number;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    readonly take?: number;
}

export function getScriptModulesOutput(args?: GetScriptModulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScriptModulesResult> {
    return pulumi.output(args).apply(a => getScriptModules(a, opts))
}

/**
 * A collection of arguments for invoking getScriptModules.
 */
export interface GetScriptModulesOutputArgs {
    /**
     * A filter to search by a list of IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to search by the partial match of a name.
     */
    partialName?: pulumi.Input<string>;
    /**
     * A list of script modules that match the filter(s).
     */
    scriptModules?: pulumi.Input<pulumi.Input<inputs.GetScriptModulesScriptModuleArgs>[]>;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: pulumi.Input<number>;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: pulumi.Input<number>;
}