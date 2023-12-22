// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getDatabricksCredential(args: GetDatabricksCredentialArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabricksCredentialResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("dbtcloud:index/getDatabricksCredential:getDatabricksCredential", {
        "credentialId": args.credentialId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabricksCredential.
 */
export interface GetDatabricksCredentialArgs {
    /**
     * Credential ID
     */
    credentialId: number;
    /**
     * Project ID
     */
    projectId: number;
}

/**
 * A collection of values returned by getDatabricksCredential.
 */
export interface GetDatabricksCredentialResult {
    /**
     * Databricks adapter ID for the credential
     */
    readonly adapterId: number;
    /**
     * The catalog where to create models
     */
    readonly catalog: string;
    /**
     * Credential ID
     */
    readonly credentialId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Number of threads to use
     */
    readonly numThreads: number;
    /**
     * Project ID
     */
    readonly projectId: number;
    /**
     * The schema where to create models
     */
    readonly schema: string;
    /**
     * Target name
     */
    readonly targetName: string;
}
export function getDatabricksCredentialOutput(args: GetDatabricksCredentialOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabricksCredentialResult> {
    return pulumi.output(args).apply((a: any) => getDatabricksCredential(a, opts))
}

/**
 * A collection of arguments for invoking getDatabricksCredential.
 */
export interface GetDatabricksCredentialOutputArgs {
    /**
     * Credential ID
     */
    credentialId: pulumi.Input<number>;
    /**
     * Project ID
     */
    projectId: pulumi.Input<number>;
}
