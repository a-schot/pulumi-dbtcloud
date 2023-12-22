// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function legacyGetBigqueryConnection(args: LegacyGetBigqueryConnectionArgs, opts?: pulumi.InvokeOptions): Promise<LegacyGetBigqueryConnectionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("dbtcloud:index/legacyGetBigqueryConnection:LegacyGetBigqueryConnection", {
        "connectionId": args.connectionId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking LegacyGetBigqueryConnection.
 */
export interface LegacyGetBigqueryConnectionArgs {
    connectionId: number;
    projectId: number;
}

/**
 * A collection of values returned by LegacyGetBigqueryConnection.
 */
export interface LegacyGetBigqueryConnectionResult {
    readonly authProviderX509CertUrl: string;
    readonly authUri: string;
    readonly clientEmail: string;
    readonly clientId: string;
    readonly clientX509CertUrl: string;
    readonly connectionId: number;
    readonly dataprocClusterName: string;
    readonly dataprocRegion: string;
    readonly executionProject: string;
    readonly gcpProjectId: string;
    readonly gcsBucket: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isActive: boolean;
    readonly isConfiguredForOauth: boolean;
    readonly location: string;
    readonly maximumBytesBilled: number;
    readonly name: string;
    readonly priority: string;
    readonly privateKey: string;
    readonly privateKeyId: string;
    readonly projectId: number;
    readonly retries: number;
    readonly timeoutSeconds: number;
    readonly tokenUri: string;
    readonly type: string;
}
export function legacyGetBigqueryConnectionOutput(args: LegacyGetBigqueryConnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<LegacyGetBigqueryConnectionResult> {
    return pulumi.output(args).apply((a: any) => legacyGetBigqueryConnection(a, opts))
}

/**
 * A collection of arguments for invoking LegacyGetBigqueryConnection.
 */
export interface LegacyGetBigqueryConnectionOutputArgs {
    connectionId: pulumi.Input<number>;
    projectId: pulumi.Input<number>;
}