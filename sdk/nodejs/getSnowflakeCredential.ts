// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSnowflakeCredential(args: GetSnowflakeCredentialArgs, opts?: pulumi.InvokeOptions): Promise<GetSnowflakeCredentialResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("dbtcloud:index/getSnowflakeCredential:getSnowflakeCredential", {
        "credentialId": args.credentialId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnowflakeCredential.
 */
export interface GetSnowflakeCredentialArgs {
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
 * A collection of values returned by getSnowflakeCredential.
 */
export interface GetSnowflakeCredentialResult {
    /**
     * The type of Snowflake credential ('password' or 'keypair')
     */
    readonly authType: string;
    /**
     * Credential ID
     */
    readonly credentialId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Whether the Snowflake credential is active
     */
    readonly isActive: boolean;
    /**
     * Number of threads to use
     */
    readonly numThreads: number;
    /**
     * Project ID
     */
    readonly projectId: number;
    /**
     * Default schema name
     */
    readonly schema: string;
    /**
     * Username for Snowflake
     */
    readonly user: string;
}
export function getSnowflakeCredentialOutput(args: GetSnowflakeCredentialOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSnowflakeCredentialResult> {
    return pulumi.output(args).apply((a: any) => getSnowflakeCredential(a, opts))
}

/**
 * A collection of arguments for invoking getSnowflakeCredential.
 */
export interface GetSnowflakeCredentialOutputArgs {
    /**
     * Credential ID
     */
    credentialId: pulumi.Input<number>;
    /**
     * Project ID
     */
    projectId: pulumi.Input<number>;
}
