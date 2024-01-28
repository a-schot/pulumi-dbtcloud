// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * @deprecated Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource
 */
export class LegacyProjectConnection extends pulumi.CustomResource {
    /**
     * Get an existing LegacyProjectConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LegacyProjectConnectionState, opts?: pulumi.CustomResourceOptions): LegacyProjectConnection {
        pulumi.log.warn("LegacyProjectConnection is deprecated: Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource")
        return new LegacyProjectConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'dbtcloud:index/legacyProjectConnection:LegacyProjectConnection';

    /**
     * Returns true if the given object is an instance of LegacyProjectConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LegacyProjectConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LegacyProjectConnection.__pulumiType;
    }

    /**
     * Connection ID
     */
    public readonly connectionId!: pulumi.Output<number>;
    /**
     * Project ID
     */
    public readonly projectId!: pulumi.Output<number>;

    /**
     * Create a LegacyProjectConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource */
    constructor(name: string, args: LegacyProjectConnectionArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource */
    constructor(name: string, argsOrState?: LegacyProjectConnectionArgs | LegacyProjectConnectionState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LegacyProjectConnection is deprecated: Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LegacyProjectConnectionState | undefined;
            resourceInputs["connectionId"] = state ? state.connectionId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as LegacyProjectConnectionArgs | undefined;
            if ((!args || args.connectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionId'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["connectionId"] = args ? args.connectionId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LegacyProjectConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LegacyProjectConnection resources.
 */
export interface LegacyProjectConnectionState {
    /**
     * Connection ID
     */
    connectionId?: pulumi.Input<number>;
    /**
     * Project ID
     */
    projectId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a LegacyProjectConnection resource.
 */
export interface LegacyProjectConnectionArgs {
    /**
     * Connection ID
     */
    connectionId: pulumi.Input<number>;
    /**
     * Project ID
     */
    projectId: pulumi.Input<number>;
}