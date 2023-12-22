// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as dbtcloud from "@aschot/pulumi-dbtcloud";
 *
 * const myEnvVarJobOverride = new dbtcloud.EnvironmentVariableJobOverride("myEnvVarJobOverride", {
 *     projectId: dbtcloud_project.dbt_project.id,
 *     jobDefinitionId: dbtcloud_job.daily_job.id,
 *     rawValue: "my_override_value",
 * });
 * ```
 *
 * ## Import
 *
 * Import using a project ID, a job ID and the environment variable override ID
 *
 * ```sh
 *  $ pulumi import dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride test_environment_variable_job_override "project_id:job_id:environment_variable_override_id"
 * ```
 *
 * ```sh
 *  $ pulumi import dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride test_environment_variable_job_override 12345:678:123456
 * ```
 */
export class EnvironmentVariableJobOverride extends pulumi.CustomResource {
    /**
     * Get an existing EnvironmentVariableJobOverride resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnvironmentVariableJobOverrideState, opts?: pulumi.CustomResourceOptions): EnvironmentVariableJobOverride {
        return new EnvironmentVariableJobOverride(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride';

    /**
     * Returns true if the given object is an instance of EnvironmentVariableJobOverride.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnvironmentVariableJobOverride {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnvironmentVariableJobOverride.__pulumiType;
    }

    /**
     * The ID of the environment variable job override
     */
    public /*out*/ readonly environmentVariableJobOverrideId!: pulumi.Output<number>;
    /**
     * The job ID for which the environment variable is being overridden
     */
    public readonly jobDefinitionId!: pulumi.Output<number>;
    /**
     * The environment variable name to override
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project ID for which the environment variable is being overridden
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * The value for the override of the environment variable
     */
    public readonly rawValue!: pulumi.Output<string>;

    /**
     * Create a EnvironmentVariableJobOverride resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentVariableJobOverrideArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnvironmentVariableJobOverrideArgs | EnvironmentVariableJobOverrideState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EnvironmentVariableJobOverrideState | undefined;
            resourceInputs["environmentVariableJobOverrideId"] = state ? state.environmentVariableJobOverrideId : undefined;
            resourceInputs["jobDefinitionId"] = state ? state.jobDefinitionId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["rawValue"] = state ? state.rawValue : undefined;
        } else {
            const args = argsOrState as EnvironmentVariableJobOverrideArgs | undefined;
            if ((!args || args.jobDefinitionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobDefinitionId'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.rawValue === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rawValue'");
            }
            resourceInputs["jobDefinitionId"] = args ? args.jobDefinitionId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["rawValue"] = args ? args.rawValue : undefined;
            resourceInputs["environmentVariableJobOverrideId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EnvironmentVariableJobOverride.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EnvironmentVariableJobOverride resources.
 */
export interface EnvironmentVariableJobOverrideState {
    /**
     * The ID of the environment variable job override
     */
    environmentVariableJobOverrideId?: pulumi.Input<number>;
    /**
     * The job ID for which the environment variable is being overridden
     */
    jobDefinitionId?: pulumi.Input<number>;
    /**
     * The environment variable name to override
     */
    name?: pulumi.Input<string>;
    /**
     * The project ID for which the environment variable is being overridden
     */
    projectId?: pulumi.Input<number>;
    /**
     * The value for the override of the environment variable
     */
    rawValue?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EnvironmentVariableJobOverride resource.
 */
export interface EnvironmentVariableJobOverrideArgs {
    /**
     * The job ID for which the environment variable is being overridden
     */
    jobDefinitionId: pulumi.Input<number>;
    /**
     * The environment variable name to override
     */
    name?: pulumi.Input<string>;
    /**
     * The project ID for which the environment variable is being overridden
     */
    projectId: pulumi.Input<number>;
    /**
     * The value for the override of the environment variable
     */
    rawValue: pulumi.Input<string>;
}
