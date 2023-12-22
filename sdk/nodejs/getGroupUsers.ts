// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Returns a list of users assigned to a specific dbt Cloud group
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as dbtcloud from "@pulumi/dbtcloud";
 *
 * const myGroupUsers = dbtcloud.getGroupUsers({
 *     groupId: 1234,
 * });
 * ```
 */
export function getGroupUsers(args: GetGroupUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupUsersResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("dbtcloud:index/getGroupUsers:getGroupUsers", {
        "groupId": args.groupId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupUsers.
 */
export interface GetGroupUsersArgs {
    /**
     * ID of the group
     */
    groupId: number;
}

/**
 * A collection of values returned by getGroupUsers.
 */
export interface GetGroupUsersResult {
    /**
     * ID of the group
     */
    readonly groupId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of users (map of ID and email) in the group
     */
    readonly users: outputs.GetGroupUsersUser[];
}
/**
 * Returns a list of users assigned to a specific dbt Cloud group
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as dbtcloud from "@pulumi/dbtcloud";
 *
 * const myGroupUsers = dbtcloud.getGroupUsers({
 *     groupId: 1234,
 * });
 * ```
 */
export function getGroupUsersOutput(args: GetGroupUsersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupUsersResult> {
    return pulumi.output(args).apply((a: any) => getGroupUsers(a, opts))
}

/**
 * A collection of arguments for invoking getGroupUsers.
 */
export interface GetGroupUsersOutputArgs {
    /**
     * ID of the group
     */
    groupId: pulumi.Input<number>;
}
