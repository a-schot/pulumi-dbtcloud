// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getBigqueryConnection(args: GetBigqueryConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetBigqueryConnectionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("dbtcloud:index/getBigqueryConnection:getBigqueryConnection", {
        "connectionId": args.connectionId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getBigqueryConnection.
 */
export interface GetBigqueryConnectionArgs {
    /**
     * Connection Identifier
     */
    connectionId: number;
    /**
     * Project ID to create the connection in
     */
    projectId: number;
}

/**
 * A collection of values returned by getBigqueryConnection.
 */
export interface GetBigqueryConnectionResult {
    /**
     * Auth Provider X509 Cert URL for the Service Account
     */
    readonly authProviderX509CertUrl: string;
    /**
     * Auth URI for the Service Account
     */
    readonly authUri: string;
    /**
     * Service Account email
     */
    readonly clientEmail: string;
    /**
     * Client ID of the Service Account
     */
    readonly clientId: string;
    /**
     * Client X509 Cert URL for the Service Account
     */
    readonly clientX509CertUrl: string;
    /**
     * Connection Identifier
     */
    readonly connectionId: number;
    /**
     * Dataproc cluster name for PySpark workloads
     */
    readonly dataprocClusterName: string;
    /**
     * Google Cloud region for PySpark workloads on Dataproc
     */
    readonly dataprocRegion: string;
    /**
     * Project to bill for query execution
     */
    readonly executionProject: string;
    /**
     * GCP project ID
     */
    readonly gcpProjectId: string;
    /**
     * URI for a Google Cloud Storage bucket to host Python code executed via Datapro
     */
    readonly gcsBucket: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Whether the connection is active
     */
    readonly isActive: boolean;
    /**
     * Whether the connection is configured for OAuth or not
     */
    readonly isConfiguredForOauth: boolean;
    /**
     * Location to create new Datasets in
     */
    readonly location: string;
    /**
     * Max number of bytes that can be billed for a given BigQuery query
     */
    readonly maximumBytesBilled: number;
    /**
     * Connection name
     */
    readonly name: string;
    /**
     * The priority with which to execute BigQuery queries
     */
    readonly priority: string;
    /**
     * Private key of the Service Account
     */
    readonly privateKey: string;
    /**
     * Private key ID of the Service Account
     */
    readonly privateKeyId: string;
    /**
     * Project ID to create the connection in
     */
    readonly projectId: number;
    /**
     * Number of retries for queries
     */
    readonly retries: number;
    /**
     * Timeout in seconds for queries
     */
    readonly timeoutSeconds: number;
    /**
     * Token URI for the Service Account
     */
    readonly tokenUri: string;
    /**
     * The type of connection
     */
    readonly type: string;
}
export function getBigqueryConnectionOutput(args: GetBigqueryConnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBigqueryConnectionResult> {
    return pulumi.output(args).apply((a: any) => getBigqueryConnection(a, opts))
}

/**
 * A collection of arguments for invoking getBigqueryConnection.
 */
export interface GetBigqueryConnectionOutputArgs {
    /**
     * Connection Identifier
     */
    connectionId: pulumi.Input<number>;
    /**
     * Project ID to create the connection in
     */
    projectId: pulumi.Input<number>;
}