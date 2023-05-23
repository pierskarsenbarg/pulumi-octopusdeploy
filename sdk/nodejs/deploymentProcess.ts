// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This resource manages deployment processes in Octopus Deploy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * // basic deployment process with 2 run a script steps
 * const example = new octopusdeploy.DeploymentProcess("example", {
 *     projectId: "Projects-123",
 *     steps: [
 *         {
 *             condition: "Success",
 *             name: "Hello world (using PowerShell)",
 *             packageRequirement: "LetOctopusDecide",
 *             runScriptActions: [{
 *                 canBeUsedForProjectVersioning: false,
 *                 condition: "Success",
 *                 isDisabled: false,
 *                 isRequired: true,
 *                 name: "Hello world (using PowerShell)",
 *                 runOnServer: true,
 *                 scriptBody: `  Write-Host 'Hello world, using PowerShell'
 *   #TODO: Experiment with steps of your own :)
 *   Write-Host '[Learn more about the types of steps available in Octopus](https://g.octopushq.com/OnboardingAddStepsLearnMore)'
 * `,
 *             }],
 *             startTrigger: "StartAfterPrevious",
 *         },
 *         {
 *             condition: "Success",
 *             name: "Hello world (using Bash)",
 *             packageRequirement: "LetOctopusDecide",
 *             runScriptActions: [{
 *                 canBeUsedForProjectVersioning: false,
 *                 condition: "Success",
 *                 isDisabled: false,
 *                 isRequired: true,
 *                 name: "Hello world (using Bash)",
 *                 runOnServer: true,
 *                 scriptBody: `  echo 'Hello world, using Bash'
 *   #TODO: Experiment with steps of your own :)
 *   echo '[Learn more about the types of steps available in Octopus](https://g.octopushq.com/OnboardingAddStepsLearnMore)'
 * `,
 *             }],
 *             startTrigger: "StartWithPrevious",
 *         },
 *     ],
 * });
 * // basic deployment process with 2 run a script steps as child steps
 * const childStepExample = new octopusdeploy.DeploymentProcess("child_step_example", {
 *     projectId: "Projects-123",
 *     steps: [{
 *         condition: "Success",
 *         name: "Hello world (using PowerShell)",
 *         packageRequirement: "LetOctopusDecide",
 *         runScriptActions: [
 *             {
 *                 canBeUsedForProjectVersioning: false,
 *                 condition: "Success",
 *                 isDisabled: false,
 *                 isRequired: true,
 *                 name: "Hello world (using PowerShell)",
 *                 scriptBody: `  Write-Host 'Hello world, using PowerShell'
 *   #TODO: Experiment with steps of your own :)
 *   Write-Host '[Learn more about the types of steps available in Octopus](https://g.octopushq.com/OnboardingAddStepsLearnMore)'
 * `,
 *             },
 *             {
 *                 canBeUsedForProjectVersioning: false,
 *                 condition: "Success",
 *                 isDisabled: false,
 *                 isRequired: true,
 *                 name: "Hello world (using Bash)",
 *                 scriptBody: `  echo 'Hello world, using Bash'
 *   #TODO: Experiment with steps of your own :)
 *   echo '[Learn more about the types of steps available in Octopus](https://g.octopushq.com/OnboardingAddStepsLearnMore)'
 * `,
 *             },
 *         ],
 *         startTrigger: "StartAfterPrevious",
 *         targetRoles: ["hello-world"],
 *     }],
 * });
 * // rolling deployment process with a step with 2 run a script steps as child steps deploying to 2 targets in parallel
 * const childStepRollingDeploymentExample = new octopusdeploy.DeploymentProcess("child_step_rolling_deployment_example", {
 *     projectId: "Projects-123",
 *     steps: [{
 *         condition: "Success",
 *         name: "Hello world (using PowerShell)",
 *         packageRequirement: "LetOctopusDecide",
 *         runScriptActions: [
 *             {
 *                 canBeUsedForProjectVersioning: false,
 *                 condition: "Success",
 *                 isDisabled: false,
 *                 isRequired: true,
 *                 name: "Hello world (using PowerShell)",
 *                 scriptBody: `  Write-Host 'Hello world, using PowerShell'
 *   #TODO: Experiment with steps of your own :)
 *   Write-Host '[Learn more about the types of steps available in Octopus](https://g.octopushq.com/OnboardingAddStepsLearnMore)'
 * `,
 *             },
 *             {
 *                 canBeUsedForProjectVersioning: false,
 *                 condition: "Success",
 *                 isDisabled: false,
 *                 isRequired: true,
 *                 name: "Hello world (using Bash)",
 *                 scriptBody: `  echo 'Hello world, using Bash'
 *   #TODO: Experiment with steps of your own :)
 *   echo '[Learn more about the types of steps available in Octopus](https://g.octopushq.com/OnboardingAddStepsLearnMore)'
 * `,
 *             },
 *         ],
 *         startTrigger: "StartAfterPrevious",
 *         targetRoles: ["hello-world"],
 *         windowSize: "2",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/deploymentProcess:DeploymentProcess [options] octopusdeploy_deployment_process.<name> <deployment-process-id>
 * ```
 */
export class DeploymentProcess extends pulumi.CustomResource {
    /**
     * Get an existing DeploymentProcess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeploymentProcessState, opts?: pulumi.CustomResourceOptions): DeploymentProcess {
        return new DeploymentProcess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/deploymentProcess:DeploymentProcess';

    /**
     * Returns true if the given object is an instance of DeploymentProcess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeploymentProcess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeploymentProcess.__pulumiType;
    }

    /**
     * The branch name associated with this deployment process (i.e. `main`). This value is optional and only applies to associated projects that are stored in version control.
     */
    public readonly branch!: pulumi.Output<string>;
    public readonly lastSnapshotId!: pulumi.Output<string | undefined>;
    /**
     * The project ID associated with this deployment process.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The space ID associated with this resource.
     */
    public readonly spaceId!: pulumi.Output<string>;
    public readonly steps!: pulumi.Output<outputs.DeploymentProcessStep[] | undefined>;
    /**
     * The version number of this deployment process.
     */
    public readonly version!: pulumi.Output<number>;

    /**
     * Create a DeploymentProcess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentProcessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeploymentProcessArgs | DeploymentProcessState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeploymentProcessState | undefined;
            resourceInputs["branch"] = state ? state.branch : undefined;
            resourceInputs["lastSnapshotId"] = state ? state.lastSnapshotId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["steps"] = state ? state.steps : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as DeploymentProcessArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["branch"] = args ? args.branch : undefined;
            resourceInputs["lastSnapshotId"] = args ? args.lastSnapshotId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["steps"] = args ? args.steps : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeploymentProcess.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeploymentProcess resources.
 */
export interface DeploymentProcessState {
    /**
     * The branch name associated with this deployment process (i.e. `main`). This value is optional and only applies to associated projects that are stored in version control.
     */
    branch?: pulumi.Input<string>;
    lastSnapshotId?: pulumi.Input<string>;
    /**
     * The project ID associated with this deployment process.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    steps?: pulumi.Input<pulumi.Input<inputs.DeploymentProcessStep>[]>;
    /**
     * The version number of this deployment process.
     */
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a DeploymentProcess resource.
 */
export interface DeploymentProcessArgs {
    /**
     * The branch name associated with this deployment process (i.e. `main`). This value is optional and only applies to associated projects that are stored in version control.
     */
    branch?: pulumi.Input<string>;
    lastSnapshotId?: pulumi.Input<string>;
    /**
     * The project ID associated with this deployment process.
     */
    projectId: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    steps?: pulumi.Input<pulumi.Input<inputs.DeploymentProcessStep>[]>;
    /**
     * The version number of this deployment process.
     */
    version?: pulumi.Input<number>;
}
