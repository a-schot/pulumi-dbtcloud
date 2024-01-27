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
 * // NOTE for customers using the LEGACY dbt_cloud provider:
 * const myCredential = new dbtcloud.BigQueryCredential("myCredential", {
 *     projectId: dbtcloud_project.dbt_project.id,
 *     dataset: "my_bq_dataset",
 *     numThreads: 16,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import dbtcloud:index/bigQueryCredential:BigQueryCredential my_credential "project_id:credential_id"
 * ```
 *
 * ```sh
 *  $ pulumi import dbtcloud:index/bigQueryCredential:BigQueryCredential my_credential 12345:5678
 * ```
 */
export class BigQueryCredential extends pulumi.CustomResource {
    /**
     * Get an existing BigQueryCredential resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BigQueryCredentialState, opts?: pulumi.CustomResourceOptions): BigQueryCredential {
        return new BigQueryCredential(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'dbtcloud:index/bigQueryCredential:BigQueryCredential';

    /**
     * Returns true if the given object is an instance of BigQueryCredential.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BigQueryCredential {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BigQueryCredential.__pulumiType;
    }

    /**
     * The system BigQuery credential ID
     */
    public /*out*/ readonly credentialId!: pulumi.Output<number>;
    /**
     * Default dataset name
     */
    public readonly dataset!: pulumi.Output<string>;
    /**
     * Whether the BigQuery credential is active
     */
    public readonly isActive!: pulumi.Output<boolean | undefined>;
    /**
     * Number of threads to use
     */
    public readonly numThreads!: pulumi.Output<number>;
    /**
     * Project ID to create the BigQuery credential in
     */
    public readonly projectId!: pulumi.Output<number>;

    /**
     * Create a BigQueryCredential resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BigQueryCredentialArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BigQueryCredentialArgs | BigQueryCredentialState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BigQueryCredentialState | undefined;
            resourceInputs["credentialId"] = state ? state.credentialId : undefined;
            resourceInputs["dataset"] = state ? state.dataset : undefined;
            resourceInputs["isActive"] = state ? state.isActive : undefined;
            resourceInputs["numThreads"] = state ? state.numThreads : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as BigQueryCredentialArgs | undefined;
            if ((!args || args.dataset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataset'");
            }
            if ((!args || args.numThreads === undefined) && !opts.urn) {
                throw new Error("Missing required property 'numThreads'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["dataset"] = args ? args.dataset : undefined;
            resourceInputs["isActive"] = args ? args.isActive : undefined;
            resourceInputs["numThreads"] = args ? args.numThreads : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["credentialId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BigQueryCredential.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BigQueryCredential resources.
 */
export interface BigQueryCredentialState {
    /**
     * The system BigQuery credential ID
     */
    credentialId?: pulumi.Input<number>;
    /**
     * Default dataset name
     */
    dataset?: pulumi.Input<string>;
    /**
     * Whether the BigQuery credential is active
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * Number of threads to use
     */
    numThreads?: pulumi.Input<number>;
    /**
     * Project ID to create the BigQuery credential in
     */
    projectId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a BigQueryCredential resource.
 */
export interface BigQueryCredentialArgs {
    /**
     * Default dataset name
     */
    dataset: pulumi.Input<string>;
    /**
     * Whether the BigQuery credential is active
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * Number of threads to use
     */
    numThreads: pulumi.Input<number>;
    /**
     * Project ID to create the BigQuery credential in
     */
    projectId: pulumi.Input<number>;
}
