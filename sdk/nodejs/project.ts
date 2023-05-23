// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource manages projects in Octopus Deploy.
 *
 * > Credentials are stored in state as plaintext. Read more about sensitive data in state.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = new octopusdeploy.Project("example", {
 *     autoCreateRelease: false,
 *     defaultGuidedFailureMode: "EnvironmentDefault",
 *     defaultToSkipIfAlreadyInstalled: false,
 *     description: "The development project.",
 *     discreteChannelRelease: false,
 *     isDisabled: false,
 *     isDiscreteChannelRelease: false,
 *     isVersionControlled: false,
 *     lifecycleId: "Lifecycles-123",
 *     projectGroupId: "ProjectGroups-123",
 *     tenantedDeploymentParticipation: "TenantedOrUntenanted",
 *     connectivityPolicy: {
 *         allowDeploymentsToNoTargets: false,
 *         excludeUnhealthyTargets: false,
 *         skipMachineBehavior: "SkipUnavailableMachines",
 *     },
 *     jiraServiceManagementExtensionSettings: {
 *         connectionId: "133d7fe602514060a48bc42ee9870f99",
 *         isEnabled: false,
 *         serviceDeskProjectName: "Test Service Desk Project (OK to Delete)",
 *     },
 *     servicenowExtensionSettings: {
 *         connectionId: "989034685e2c48c4b06a29286c9ef5cc",
 *         isEnabled: false,
 *         isStateAutomaticallyTransitioned: false,
 *         standardChangeTemplateName: "Standard Change Template Name (OK to Delete)",
 *     },
 *     templates: [{
 *         defaultValue: "example-default-value",
 *         helpText: "example-help-test",
 *         label: "example-label",
 *         name: "example-template-value",
 *         displaySettings: {
 *             "Octopus.ControlType": "SingleLineText",
 *         },
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import octopusdeploy:index/project:Project [options] octopusdeploy_project.<name> <project-id>
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * @deprecated This value is only valid for an associated connectivity policy and should not be specified here.
     */
    public readonly allowDeploymentsToNoTargets!: pulumi.Output<boolean | undefined>;
    public readonly autoCreateRelease!: pulumi.Output<boolean>;
    public readonly autoDeployReleaseOverrides!: pulumi.Output<string>;
    public readonly clonedFromProjectId!: pulumi.Output<string>;
    public readonly connectivityPolicy!: pulumi.Output<outputs.ProjectConnectivityPolicy>;
    public readonly defaultGuidedFailureMode!: pulumi.Output<string>;
    public readonly defaultToSkipIfAlreadyInstalled!: pulumi.Output<boolean>;
    public readonly deploymentChangesTemplate!: pulumi.Output<string>;
    public /*out*/ readonly deploymentProcessId!: pulumi.Output<string>;
    /**
     * The description of this project.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Treats releases of different channels to the same environment as a separate deployment dimension
     */
    public readonly discreteChannelRelease!: pulumi.Output<boolean | undefined>;
    /**
     * Provides Git-related persistence settings for a version-controlled project.
     */
    public readonly gitAnonymousPersistenceSettings!: pulumi.Output<outputs.ProjectGitAnonymousPersistenceSettings | undefined>;
    /**
     * Provides Git-related persistence settings for a version-controlled project.
     */
    public readonly gitLibraryPersistenceSettings!: pulumi.Output<outputs.ProjectGitLibraryPersistenceSettings | undefined>;
    /**
     * Provides Git-related persistence settings for a version-controlled project.
     */
    public readonly gitUsernamePasswordPersistenceSettings!: pulumi.Output<outputs.ProjectGitUsernamePasswordPersistenceSettings | undefined>;
    public readonly includedLibraryVariableSets!: pulumi.Output<string[]>;
    public readonly isDisabled!: pulumi.Output<boolean>;
    /**
     * Treats releases of different channels to the same environment as a separate deployment dimension
     */
    public readonly isDiscreteChannelRelease!: pulumi.Output<boolean>;
    public readonly isVersionControlled!: pulumi.Output<boolean>;
    /**
     * Provides extension settings for the Jira Service Management (JSM) integration for this project.
     */
    public readonly jiraServiceManagementExtensionSettings!: pulumi.Output<outputs.ProjectJiraServiceManagementExtensionSettings | undefined>;
    /**
     * The lifecycle ID associated with this project.
     */
    public readonly lifecycleId!: pulumi.Output<string>;
    /**
     * The name of the project in Octopus Deploy. This name must be unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project group ID associated with this project.
     */
    public readonly projectGroupId!: pulumi.Output<string>;
    public readonly releaseCreationStrategy!: pulumi.Output<outputs.ProjectReleaseCreationStrategy>;
    public readonly releaseNotesTemplate!: pulumi.Output<string>;
    /**
     * Provides extension settings for the ServiceNow integration for this project.
     */
    public readonly servicenowExtensionSettings!: pulumi.Output<outputs.ProjectServicenowExtensionSettings | undefined>;
    /**
     * A human-readable, unique identifier, used to identify a project.
     */
    public readonly slug!: pulumi.Output<string>;
    /**
     * The space ID associated with this project.
     */
    public readonly spaceId!: pulumi.Output<string>;
    public readonly templates!: pulumi.Output<outputs.ProjectTemplate[] | undefined>;
    /**
     * The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
     */
    public readonly tenantedDeploymentParticipation!: pulumi.Output<string>;
    public /*out*/ readonly variableSetId!: pulumi.Output<string>;
    public readonly versioningStrategies!: pulumi.Output<outputs.ProjectVersioningStrategy[]>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["allowDeploymentsToNoTargets"] = state ? state.allowDeploymentsToNoTargets : undefined;
            resourceInputs["autoCreateRelease"] = state ? state.autoCreateRelease : undefined;
            resourceInputs["autoDeployReleaseOverrides"] = state ? state.autoDeployReleaseOverrides : undefined;
            resourceInputs["clonedFromProjectId"] = state ? state.clonedFromProjectId : undefined;
            resourceInputs["connectivityPolicy"] = state ? state.connectivityPolicy : undefined;
            resourceInputs["defaultGuidedFailureMode"] = state ? state.defaultGuidedFailureMode : undefined;
            resourceInputs["defaultToSkipIfAlreadyInstalled"] = state ? state.defaultToSkipIfAlreadyInstalled : undefined;
            resourceInputs["deploymentChangesTemplate"] = state ? state.deploymentChangesTemplate : undefined;
            resourceInputs["deploymentProcessId"] = state ? state.deploymentProcessId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["discreteChannelRelease"] = state ? state.discreteChannelRelease : undefined;
            resourceInputs["gitAnonymousPersistenceSettings"] = state ? state.gitAnonymousPersistenceSettings : undefined;
            resourceInputs["gitLibraryPersistenceSettings"] = state ? state.gitLibraryPersistenceSettings : undefined;
            resourceInputs["gitUsernamePasswordPersistenceSettings"] = state ? state.gitUsernamePasswordPersistenceSettings : undefined;
            resourceInputs["includedLibraryVariableSets"] = state ? state.includedLibraryVariableSets : undefined;
            resourceInputs["isDisabled"] = state ? state.isDisabled : undefined;
            resourceInputs["isDiscreteChannelRelease"] = state ? state.isDiscreteChannelRelease : undefined;
            resourceInputs["isVersionControlled"] = state ? state.isVersionControlled : undefined;
            resourceInputs["jiraServiceManagementExtensionSettings"] = state ? state.jiraServiceManagementExtensionSettings : undefined;
            resourceInputs["lifecycleId"] = state ? state.lifecycleId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectGroupId"] = state ? state.projectGroupId : undefined;
            resourceInputs["releaseCreationStrategy"] = state ? state.releaseCreationStrategy : undefined;
            resourceInputs["releaseNotesTemplate"] = state ? state.releaseNotesTemplate : undefined;
            resourceInputs["servicenowExtensionSettings"] = state ? state.servicenowExtensionSettings : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["templates"] = state ? state.templates : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = state ? state.tenantedDeploymentParticipation : undefined;
            resourceInputs["variableSetId"] = state ? state.variableSetId : undefined;
            resourceInputs["versioningStrategies"] = state ? state.versioningStrategies : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.lifecycleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lifecycleId'");
            }
            if ((!args || args.projectGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectGroupId'");
            }
            resourceInputs["allowDeploymentsToNoTargets"] = args ? args.allowDeploymentsToNoTargets : undefined;
            resourceInputs["autoCreateRelease"] = args ? args.autoCreateRelease : undefined;
            resourceInputs["autoDeployReleaseOverrides"] = args ? args.autoDeployReleaseOverrides : undefined;
            resourceInputs["clonedFromProjectId"] = args ? args.clonedFromProjectId : undefined;
            resourceInputs["connectivityPolicy"] = args ? args.connectivityPolicy : undefined;
            resourceInputs["defaultGuidedFailureMode"] = args ? args.defaultGuidedFailureMode : undefined;
            resourceInputs["defaultToSkipIfAlreadyInstalled"] = args ? args.defaultToSkipIfAlreadyInstalled : undefined;
            resourceInputs["deploymentChangesTemplate"] = args ? args.deploymentChangesTemplate : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["discreteChannelRelease"] = args ? args.discreteChannelRelease : undefined;
            resourceInputs["gitAnonymousPersistenceSettings"] = args ? args.gitAnonymousPersistenceSettings : undefined;
            resourceInputs["gitLibraryPersistenceSettings"] = args ? args.gitLibraryPersistenceSettings : undefined;
            resourceInputs["gitUsernamePasswordPersistenceSettings"] = args ? args.gitUsernamePasswordPersistenceSettings : undefined;
            resourceInputs["includedLibraryVariableSets"] = args ? args.includedLibraryVariableSets : undefined;
            resourceInputs["isDisabled"] = args ? args.isDisabled : undefined;
            resourceInputs["isDiscreteChannelRelease"] = args ? args.isDiscreteChannelRelease : undefined;
            resourceInputs["isVersionControlled"] = args ? args.isVersionControlled : undefined;
            resourceInputs["jiraServiceManagementExtensionSettings"] = args ? args.jiraServiceManagementExtensionSettings : undefined;
            resourceInputs["lifecycleId"] = args ? args.lifecycleId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectGroupId"] = args ? args.projectGroupId : undefined;
            resourceInputs["releaseCreationStrategy"] = args ? args.releaseCreationStrategy : undefined;
            resourceInputs["releaseNotesTemplate"] = args ? args.releaseNotesTemplate : undefined;
            resourceInputs["servicenowExtensionSettings"] = args ? args.servicenowExtensionSettings : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["templates"] = args ? args.templates : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = args ? args.tenantedDeploymentParticipation : undefined;
            resourceInputs["versioningStrategies"] = args ? args.versioningStrategies : undefined;
            resourceInputs["deploymentProcessId"] = undefined /*out*/;
            resourceInputs["variableSetId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * @deprecated This value is only valid for an associated connectivity policy and should not be specified here.
     */
    allowDeploymentsToNoTargets?: pulumi.Input<boolean>;
    autoCreateRelease?: pulumi.Input<boolean>;
    autoDeployReleaseOverrides?: pulumi.Input<string>;
    clonedFromProjectId?: pulumi.Input<string>;
    connectivityPolicy?: pulumi.Input<inputs.ProjectConnectivityPolicy>;
    defaultGuidedFailureMode?: pulumi.Input<string>;
    defaultToSkipIfAlreadyInstalled?: pulumi.Input<boolean>;
    deploymentChangesTemplate?: pulumi.Input<string>;
    deploymentProcessId?: pulumi.Input<string>;
    /**
     * The description of this project.
     */
    description?: pulumi.Input<string>;
    /**
     * Treats releases of different channels to the same environment as a separate deployment dimension
     */
    discreteChannelRelease?: pulumi.Input<boolean>;
    /**
     * Provides Git-related persistence settings for a version-controlled project.
     */
    gitAnonymousPersistenceSettings?: pulumi.Input<inputs.ProjectGitAnonymousPersistenceSettings>;
    /**
     * Provides Git-related persistence settings for a version-controlled project.
     */
    gitLibraryPersistenceSettings?: pulumi.Input<inputs.ProjectGitLibraryPersistenceSettings>;
    /**
     * Provides Git-related persistence settings for a version-controlled project.
     */
    gitUsernamePasswordPersistenceSettings?: pulumi.Input<inputs.ProjectGitUsernamePasswordPersistenceSettings>;
    includedLibraryVariableSets?: pulumi.Input<pulumi.Input<string>[]>;
    isDisabled?: pulumi.Input<boolean>;
    /**
     * Treats releases of different channels to the same environment as a separate deployment dimension
     */
    isDiscreteChannelRelease?: pulumi.Input<boolean>;
    isVersionControlled?: pulumi.Input<boolean>;
    /**
     * Provides extension settings for the Jira Service Management (JSM) integration for this project.
     */
    jiraServiceManagementExtensionSettings?: pulumi.Input<inputs.ProjectJiraServiceManagementExtensionSettings>;
    /**
     * The lifecycle ID associated with this project.
     */
    lifecycleId?: pulumi.Input<string>;
    /**
     * The name of the project in Octopus Deploy. This name must be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The project group ID associated with this project.
     */
    projectGroupId?: pulumi.Input<string>;
    releaseCreationStrategy?: pulumi.Input<inputs.ProjectReleaseCreationStrategy>;
    releaseNotesTemplate?: pulumi.Input<string>;
    /**
     * Provides extension settings for the ServiceNow integration for this project.
     */
    servicenowExtensionSettings?: pulumi.Input<inputs.ProjectServicenowExtensionSettings>;
    /**
     * A human-readable, unique identifier, used to identify a project.
     */
    slug?: pulumi.Input<string>;
    /**
     * The space ID associated with this project.
     */
    spaceId?: pulumi.Input<string>;
    templates?: pulumi.Input<pulumi.Input<inputs.ProjectTemplate>[]>;
    /**
     * The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
     */
    tenantedDeploymentParticipation?: pulumi.Input<string>;
    variableSetId?: pulumi.Input<string>;
    versioningStrategies?: pulumi.Input<pulumi.Input<inputs.ProjectVersioningStrategy>[]>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * @deprecated This value is only valid for an associated connectivity policy and should not be specified here.
     */
    allowDeploymentsToNoTargets?: pulumi.Input<boolean>;
    autoCreateRelease?: pulumi.Input<boolean>;
    autoDeployReleaseOverrides?: pulumi.Input<string>;
    clonedFromProjectId?: pulumi.Input<string>;
    connectivityPolicy?: pulumi.Input<inputs.ProjectConnectivityPolicy>;
    defaultGuidedFailureMode?: pulumi.Input<string>;
    defaultToSkipIfAlreadyInstalled?: pulumi.Input<boolean>;
    deploymentChangesTemplate?: pulumi.Input<string>;
    /**
     * The description of this project.
     */
    description?: pulumi.Input<string>;
    /**
     * Treats releases of different channels to the same environment as a separate deployment dimension
     */
    discreteChannelRelease?: pulumi.Input<boolean>;
    /**
     * Provides Git-related persistence settings for a version-controlled project.
     */
    gitAnonymousPersistenceSettings?: pulumi.Input<inputs.ProjectGitAnonymousPersistenceSettings>;
    /**
     * Provides Git-related persistence settings for a version-controlled project.
     */
    gitLibraryPersistenceSettings?: pulumi.Input<inputs.ProjectGitLibraryPersistenceSettings>;
    /**
     * Provides Git-related persistence settings for a version-controlled project.
     */
    gitUsernamePasswordPersistenceSettings?: pulumi.Input<inputs.ProjectGitUsernamePasswordPersistenceSettings>;
    includedLibraryVariableSets?: pulumi.Input<pulumi.Input<string>[]>;
    isDisabled?: pulumi.Input<boolean>;
    /**
     * Treats releases of different channels to the same environment as a separate deployment dimension
     */
    isDiscreteChannelRelease?: pulumi.Input<boolean>;
    isVersionControlled?: pulumi.Input<boolean>;
    /**
     * Provides extension settings for the Jira Service Management (JSM) integration for this project.
     */
    jiraServiceManagementExtensionSettings?: pulumi.Input<inputs.ProjectJiraServiceManagementExtensionSettings>;
    /**
     * The lifecycle ID associated with this project.
     */
    lifecycleId: pulumi.Input<string>;
    /**
     * The name of the project in Octopus Deploy. This name must be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The project group ID associated with this project.
     */
    projectGroupId: pulumi.Input<string>;
    releaseCreationStrategy?: pulumi.Input<inputs.ProjectReleaseCreationStrategy>;
    releaseNotesTemplate?: pulumi.Input<string>;
    /**
     * Provides extension settings for the ServiceNow integration for this project.
     */
    servicenowExtensionSettings?: pulumi.Input<inputs.ProjectServicenowExtensionSettings>;
    /**
     * A human-readable, unique identifier, used to identify a project.
     */
    slug?: pulumi.Input<string>;
    /**
     * The space ID associated with this project.
     */
    spaceId?: pulumi.Input<string>;
    templates?: pulumi.Input<pulumi.Input<inputs.ProjectTemplate>[]>;
    /**
     * The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
     */
    tenantedDeploymentParticipation?: pulumi.Input<string>;
    versioningStrategies?: pulumi.Input<pulumi.Input<inputs.ProjectVersioningStrategy>[]>;
}
