// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * @deprecated Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource
 */
export class LegacyProjectRepository extends pulumi.CustomResource {
    /**
     * Get an existing LegacyProjectRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LegacyProjectRepositoryState, opts?: pulumi.CustomResourceOptions): LegacyProjectRepository {
        pulumi.log.warn("LegacyProjectRepository is deprecated: Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource")
        return new LegacyProjectRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'dbtcloud:index/legacyProjectRepository:LegacyProjectRepository';

    /**
     * Returns true if the given object is an instance of LegacyProjectRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LegacyProjectRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LegacyProjectRepository.__pulumiType;
    }

    /**
     * Project ID
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * Repository ID
     */
    public readonly repositoryId!: pulumi.Output<number>;

    /**
     * Create a LegacyProjectRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource */
    constructor(name: string, args: LegacyProjectRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource */
    constructor(name: string, argsOrState?: LegacyProjectRepositoryArgs | LegacyProjectRepositoryState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LegacyProjectRepository is deprecated: Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LegacyProjectRepositoryState | undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["repositoryId"] = state ? state.repositoryId : undefined;
        } else {
            const args = argsOrState as LegacyProjectRepositoryArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.repositoryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repositoryId'");
            }
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["repositoryId"] = args ? args.repositoryId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LegacyProjectRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LegacyProjectRepository resources.
 */
export interface LegacyProjectRepositoryState {
    /**
     * Project ID
     */
    projectId?: pulumi.Input<number>;
    /**
     * Repository ID
     */
    repositoryId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a LegacyProjectRepository resource.
 */
export interface LegacyProjectRepositoryArgs {
    /**
     * Project ID
     */
    projectId: pulumi.Input<number>;
    /**
     * Repository ID
     */
    repositoryId: pulumi.Input<number>;
}
