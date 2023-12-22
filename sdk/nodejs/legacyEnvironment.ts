// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class LegacyEnvironment extends pulumi.CustomResource {
    /**
     * Get an existing LegacyEnvironment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LegacyEnvironmentState, opts?: pulumi.CustomResourceOptions): LegacyEnvironment {
        return new LegacyEnvironment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'dbtcloud:index/legacyEnvironment:LegacyEnvironment';

    /**
     * Returns true if the given object is an instance of LegacyEnvironment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LegacyEnvironment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LegacyEnvironment.__pulumiType;
    }

    /**
     * Credential ID to create the environment with. A credential is not required for development environments but is required
     * for deployment environments
     */
    public readonly credentialId!: pulumi.Output<number | undefined>;
    /**
     * Which custom branch to use in this environment
     */
    public readonly customBranch!: pulumi.Output<string | undefined>;
    /**
     * Version number of dbt to use in this environment. It needs to be in the format `major.minor.0-latest` or
     * `major.minor.0-pre`, e.g. `1.5.0-latest`
     */
    public readonly dbtVersion!: pulumi.Output<string>;
    /**
     * The type of environment. Only valid for environments of type 'deployment' and for now can only be empty or set to
     * 'production'
     */
    public readonly deploymentType!: pulumi.Output<string | undefined>;
    /**
     * Environment ID within the project
     */
    public /*out*/ readonly environmentId!: pulumi.Output<number>;
    /**
     * ID of the extended attributes for the environment
     */
    public readonly extendedAttributesId!: pulumi.Output<number | undefined>;
    /**
     * Whether the environment is active
     */
    public readonly isActive!: pulumi.Output<boolean | undefined>;
    /**
     * Environment name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Project ID to create the environment in
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * The type of environment (must be either development or deployment)
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Whether to use a custom git branch in this environment
     */
    public readonly useCustomBranch!: pulumi.Output<boolean | undefined>;

    /**
     * Create a LegacyEnvironment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LegacyEnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LegacyEnvironmentArgs | LegacyEnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LegacyEnvironmentState | undefined;
            resourceInputs["credentialId"] = state ? state.credentialId : undefined;
            resourceInputs["customBranch"] = state ? state.customBranch : undefined;
            resourceInputs["dbtVersion"] = state ? state.dbtVersion : undefined;
            resourceInputs["deploymentType"] = state ? state.deploymentType : undefined;
            resourceInputs["environmentId"] = state ? state.environmentId : undefined;
            resourceInputs["extendedAttributesId"] = state ? state.extendedAttributesId : undefined;
            resourceInputs["isActive"] = state ? state.isActive : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["useCustomBranch"] = state ? state.useCustomBranch : undefined;
        } else {
            const args = argsOrState as LegacyEnvironmentArgs | undefined;
            if ((!args || args.dbtVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbtVersion'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["credentialId"] = args ? args.credentialId : undefined;
            resourceInputs["customBranch"] = args ? args.customBranch : undefined;
            resourceInputs["dbtVersion"] = args ? args.dbtVersion : undefined;
            resourceInputs["deploymentType"] = args ? args.deploymentType : undefined;
            resourceInputs["extendedAttributesId"] = args ? args.extendedAttributesId : undefined;
            resourceInputs["isActive"] = args ? args.isActive : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["useCustomBranch"] = args ? args.useCustomBranch : undefined;
            resourceInputs["environmentId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LegacyEnvironment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LegacyEnvironment resources.
 */
export interface LegacyEnvironmentState {
    /**
     * Credential ID to create the environment with. A credential is not required for development environments but is required
     * for deployment environments
     */
    credentialId?: pulumi.Input<number>;
    /**
     * Which custom branch to use in this environment
     */
    customBranch?: pulumi.Input<string>;
    /**
     * Version number of dbt to use in this environment. It needs to be in the format `major.minor.0-latest` or
     * `major.minor.0-pre`, e.g. `1.5.0-latest`
     */
    dbtVersion?: pulumi.Input<string>;
    /**
     * The type of environment. Only valid for environments of type 'deployment' and for now can only be empty or set to
     * 'production'
     */
    deploymentType?: pulumi.Input<string>;
    /**
     * Environment ID within the project
     */
    environmentId?: pulumi.Input<number>;
    /**
     * ID of the extended attributes for the environment
     */
    extendedAttributesId?: pulumi.Input<number>;
    /**
     * Whether the environment is active
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * Environment name
     */
    name?: pulumi.Input<string>;
    /**
     * Project ID to create the environment in
     */
    projectId?: pulumi.Input<number>;
    /**
     * The type of environment (must be either development or deployment)
     */
    type?: pulumi.Input<string>;
    /**
     * Whether to use a custom git branch in this environment
     */
    useCustomBranch?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a LegacyEnvironment resource.
 */
export interface LegacyEnvironmentArgs {
    /**
     * Credential ID to create the environment with. A credential is not required for development environments but is required
     * for deployment environments
     */
    credentialId?: pulumi.Input<number>;
    /**
     * Which custom branch to use in this environment
     */
    customBranch?: pulumi.Input<string>;
    /**
     * Version number of dbt to use in this environment. It needs to be in the format `major.minor.0-latest` or
     * `major.minor.0-pre`, e.g. `1.5.0-latest`
     */
    dbtVersion: pulumi.Input<string>;
    /**
     * The type of environment. Only valid for environments of type 'deployment' and for now can only be empty or set to
     * 'production'
     */
    deploymentType?: pulumi.Input<string>;
    /**
     * ID of the extended attributes for the environment
     */
    extendedAttributesId?: pulumi.Input<number>;
    /**
     * Whether the environment is active
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * Environment name
     */
    name?: pulumi.Input<string>;
    /**
     * Project ID to create the environment in
     */
    projectId: pulumi.Input<number>;
    /**
     * The type of environment (must be either development or deployment)
     */
    type: pulumi.Input<string>;
    /**
     * Whether to use a custom git branch in this environment
     */
    useCustomBranch?: pulumi.Input<boolean>;
}