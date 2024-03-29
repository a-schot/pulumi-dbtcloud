// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * @deprecated Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource
 */
export class LegacyProjectArtefacts extends pulumi.CustomResource {
    /**
     * Get an existing LegacyProjectArtefacts resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LegacyProjectArtefactsState, opts?: pulumi.CustomResourceOptions): LegacyProjectArtefacts {
        pulumi.log.warn("LegacyProjectArtefacts is deprecated: Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource")
        return new LegacyProjectArtefacts(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'dbtcloud:index/legacyProjectArtefacts:LegacyProjectArtefacts';

    /**
     * Returns true if the given object is an instance of LegacyProjectArtefacts.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LegacyProjectArtefacts {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LegacyProjectArtefacts.__pulumiType;
    }

    /**
     * Docs Job ID
     */
    public readonly docsJobId!: pulumi.Output<number | undefined>;
    /**
     * Freshness Job ID
     */
    public readonly freshnessJobId!: pulumi.Output<number | undefined>;
    /**
     * Project ID
     */
    public readonly projectId!: pulumi.Output<number>;

    /**
     * Create a LegacyProjectArtefacts resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource */
    constructor(name: string, args: LegacyProjectArtefactsArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource */
    constructor(name: string, argsOrState?: LegacyProjectArtefactsArgs | LegacyProjectArtefactsState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LegacyProjectArtefacts is deprecated: Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LegacyProjectArtefactsState | undefined;
            resourceInputs["docsJobId"] = state ? state.docsJobId : undefined;
            resourceInputs["freshnessJobId"] = state ? state.freshnessJobId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as LegacyProjectArtefactsArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["docsJobId"] = args ? args.docsJobId : undefined;
            resourceInputs["freshnessJobId"] = args ? args.freshnessJobId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LegacyProjectArtefacts.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LegacyProjectArtefacts resources.
 */
export interface LegacyProjectArtefactsState {
    /**
     * Docs Job ID
     */
    docsJobId?: pulumi.Input<number>;
    /**
     * Freshness Job ID
     */
    freshnessJobId?: pulumi.Input<number>;
    /**
     * Project ID
     */
    projectId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a LegacyProjectArtefacts resource.
 */
export interface LegacyProjectArtefactsArgs {
    /**
     * Docs Job ID
     */
    docsJobId?: pulumi.Input<number>;
    /**
     * Freshness Job ID
     */
    freshnessJobId?: pulumi.Input<number>;
    /**
     * Project ID
     */
    projectId: pulumi.Input<number>;
}
