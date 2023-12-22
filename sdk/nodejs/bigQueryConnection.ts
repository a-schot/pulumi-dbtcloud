// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource to create BigQuery connections in dbt Cloud. Can be set to use OAuth for developers.
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import dbtcloud:index/bigQueryConnection:BigQueryConnection my_connection "project_id:connection_id"
 * ```
 *
 * ```sh
 *  $ pulumi import dbtcloud:index/bigQueryConnection:BigQueryConnection my_connection 12345:6789
 * ```
 */
export class BigQueryConnection extends pulumi.CustomResource {
    /**
     * Get an existing BigQueryConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BigQueryConnectionState, opts?: pulumi.CustomResourceOptions): BigQueryConnection {
        return new BigQueryConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'dbtcloud:index/bigQueryConnection:BigQueryConnection';

    /**
     * Returns true if the given object is an instance of BigQueryConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BigQueryConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BigQueryConnection.__pulumiType;
    }

    /**
     * The Application ID for BQ OAuth
     */
    public readonly applicationId!: pulumi.Output<string | undefined>;
    /**
     * The Application Secret for BQ OAuth
     */
    public readonly applicationSecret!: pulumi.Output<string | undefined>;
    /**
     * Auth Provider X509 Cert URL for the Service Account
     */
    public readonly authProviderX509CertUrl!: pulumi.Output<string>;
    /**
     * Auth URI for the Service Account
     */
    public readonly authUri!: pulumi.Output<string>;
    /**
     * Service Account email
     */
    public readonly clientEmail!: pulumi.Output<string>;
    /**
     * Client ID of the Service Account
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * Client X509 Cert URL for the Service Account
     */
    public readonly clientX509CertUrl!: pulumi.Output<string>;
    /**
     * Connection Identifier
     */
    public /*out*/ readonly connectionId!: pulumi.Output<number>;
    /**
     * Dataproc cluster name for PySpark workloads
     */
    public readonly dataprocClusterName!: pulumi.Output<string | undefined>;
    /**
     * Google Cloud region for PySpark workloads on Dataproc
     */
    public readonly dataprocRegion!: pulumi.Output<string | undefined>;
    /**
     * Project to bill for query execution
     */
    public readonly executionProject!: pulumi.Output<string | undefined>;
    /**
     * GCP project ID
     */
    public readonly gcpProjectId!: pulumi.Output<string>;
    /**
     * URI for a Google Cloud Storage bucket to host Python code executed via Datapro
     */
    public readonly gcsBucket!: pulumi.Output<string | undefined>;
    /**
     * Whether the connection is active
     */
    public readonly isActive!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the connection is configured for OAuth or not
     */
    public /*out*/ readonly isConfiguredForOauth!: pulumi.Output<boolean>;
    /**
     * Location to create new Datasets in
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Max number of bytes that can be billed for a given BigQuery query
     */
    public readonly maximumBytesBilled!: pulumi.Output<number | undefined>;
    /**
     * Connection name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The priority with which to execute BigQuery queries (batch or interactive)
     */
    public readonly priority!: pulumi.Output<string | undefined>;
    /**
     * Private key of the Service Account
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * Private key ID of the Service Account
     */
    public readonly privateKeyId!: pulumi.Output<string>;
    /**
     * Project ID to create the connection in
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * Number of retries for queries
     */
    public readonly retries!: pulumi.Output<number | undefined>;
    /**
     * Timeout in seconds for queries
     */
    public readonly timeoutSeconds!: pulumi.Output<number>;
    /**
     * Token URI for the Service Account
     */
    public readonly tokenUri!: pulumi.Output<string>;
    /**
     * The type of connection
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a BigQueryConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BigQueryConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BigQueryConnectionArgs | BigQueryConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BigQueryConnectionState | undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["applicationSecret"] = state ? state.applicationSecret : undefined;
            resourceInputs["authProviderX509CertUrl"] = state ? state.authProviderX509CertUrl : undefined;
            resourceInputs["authUri"] = state ? state.authUri : undefined;
            resourceInputs["clientEmail"] = state ? state.clientEmail : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientX509CertUrl"] = state ? state.clientX509CertUrl : undefined;
            resourceInputs["connectionId"] = state ? state.connectionId : undefined;
            resourceInputs["dataprocClusterName"] = state ? state.dataprocClusterName : undefined;
            resourceInputs["dataprocRegion"] = state ? state.dataprocRegion : undefined;
            resourceInputs["executionProject"] = state ? state.executionProject : undefined;
            resourceInputs["gcpProjectId"] = state ? state.gcpProjectId : undefined;
            resourceInputs["gcsBucket"] = state ? state.gcsBucket : undefined;
            resourceInputs["isActive"] = state ? state.isActive : undefined;
            resourceInputs["isConfiguredForOauth"] = state ? state.isConfiguredForOauth : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["maximumBytesBilled"] = state ? state.maximumBytesBilled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["privateKeyId"] = state ? state.privateKeyId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["retries"] = state ? state.retries : undefined;
            resourceInputs["timeoutSeconds"] = state ? state.timeoutSeconds : undefined;
            resourceInputs["tokenUri"] = state ? state.tokenUri : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as BigQueryConnectionArgs | undefined;
            if ((!args || args.authProviderX509CertUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authProviderX509CertUrl'");
            }
            if ((!args || args.authUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authUri'");
            }
            if ((!args || args.clientEmail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientEmail'");
            }
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clientX509CertUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientX509CertUrl'");
            }
            if ((!args || args.gcpProjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gcpProjectId'");
            }
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            if ((!args || args.privateKeyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKeyId'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.timeoutSeconds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeoutSeconds'");
            }
            if ((!args || args.tokenUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tokenUri'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["applicationId"] = args?.applicationId ? pulumi.secret(args.applicationId) : undefined;
            resourceInputs["applicationSecret"] = args?.applicationSecret ? pulumi.secret(args.applicationSecret) : undefined;
            resourceInputs["authProviderX509CertUrl"] = args ? args.authProviderX509CertUrl : undefined;
            resourceInputs["authUri"] = args ? args.authUri : undefined;
            resourceInputs["clientEmail"] = args ? args.clientEmail : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientX509CertUrl"] = args ? args.clientX509CertUrl : undefined;
            resourceInputs["dataprocClusterName"] = args ? args.dataprocClusterName : undefined;
            resourceInputs["dataprocRegion"] = args ? args.dataprocRegion : undefined;
            resourceInputs["executionProject"] = args ? args.executionProject : undefined;
            resourceInputs["gcpProjectId"] = args ? args.gcpProjectId : undefined;
            resourceInputs["gcsBucket"] = args ? args.gcsBucket : undefined;
            resourceInputs["isActive"] = args ? args.isActive : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["maximumBytesBilled"] = args ? args.maximumBytesBilled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["privateKeyId"] = args ? args.privateKeyId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["retries"] = args ? args.retries : undefined;
            resourceInputs["timeoutSeconds"] = args ? args.timeoutSeconds : undefined;
            resourceInputs["tokenUri"] = args ? args.tokenUri : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["connectionId"] = undefined /*out*/;
            resourceInputs["isConfiguredForOauth"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["applicationId", "applicationSecret", "privateKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(BigQueryConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BigQueryConnection resources.
 */
export interface BigQueryConnectionState {
    /**
     * The Application ID for BQ OAuth
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The Application Secret for BQ OAuth
     */
    applicationSecret?: pulumi.Input<string>;
    /**
     * Auth Provider X509 Cert URL for the Service Account
     */
    authProviderX509CertUrl?: pulumi.Input<string>;
    /**
     * Auth URI for the Service Account
     */
    authUri?: pulumi.Input<string>;
    /**
     * Service Account email
     */
    clientEmail?: pulumi.Input<string>;
    /**
     * Client ID of the Service Account
     */
    clientId?: pulumi.Input<string>;
    /**
     * Client X509 Cert URL for the Service Account
     */
    clientX509CertUrl?: pulumi.Input<string>;
    /**
     * Connection Identifier
     */
    connectionId?: pulumi.Input<number>;
    /**
     * Dataproc cluster name for PySpark workloads
     */
    dataprocClusterName?: pulumi.Input<string>;
    /**
     * Google Cloud region for PySpark workloads on Dataproc
     */
    dataprocRegion?: pulumi.Input<string>;
    /**
     * Project to bill for query execution
     */
    executionProject?: pulumi.Input<string>;
    /**
     * GCP project ID
     */
    gcpProjectId?: pulumi.Input<string>;
    /**
     * URI for a Google Cloud Storage bucket to host Python code executed via Datapro
     */
    gcsBucket?: pulumi.Input<string>;
    /**
     * Whether the connection is active
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * Whether the connection is configured for OAuth or not
     */
    isConfiguredForOauth?: pulumi.Input<boolean>;
    /**
     * Location to create new Datasets in
     */
    location?: pulumi.Input<string>;
    /**
     * Max number of bytes that can be billed for a given BigQuery query
     */
    maximumBytesBilled?: pulumi.Input<number>;
    /**
     * Connection name
     */
    name?: pulumi.Input<string>;
    /**
     * The priority with which to execute BigQuery queries (batch or interactive)
     */
    priority?: pulumi.Input<string>;
    /**
     * Private key of the Service Account
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Private key ID of the Service Account
     */
    privateKeyId?: pulumi.Input<string>;
    /**
     * Project ID to create the connection in
     */
    projectId?: pulumi.Input<number>;
    /**
     * Number of retries for queries
     */
    retries?: pulumi.Input<number>;
    /**
     * Timeout in seconds for queries
     */
    timeoutSeconds?: pulumi.Input<number>;
    /**
     * Token URI for the Service Account
     */
    tokenUri?: pulumi.Input<string>;
    /**
     * The type of connection
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BigQueryConnection resource.
 */
export interface BigQueryConnectionArgs {
    /**
     * The Application ID for BQ OAuth
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The Application Secret for BQ OAuth
     */
    applicationSecret?: pulumi.Input<string>;
    /**
     * Auth Provider X509 Cert URL for the Service Account
     */
    authProviderX509CertUrl: pulumi.Input<string>;
    /**
     * Auth URI for the Service Account
     */
    authUri: pulumi.Input<string>;
    /**
     * Service Account email
     */
    clientEmail: pulumi.Input<string>;
    /**
     * Client ID of the Service Account
     */
    clientId: pulumi.Input<string>;
    /**
     * Client X509 Cert URL for the Service Account
     */
    clientX509CertUrl: pulumi.Input<string>;
    /**
     * Dataproc cluster name for PySpark workloads
     */
    dataprocClusterName?: pulumi.Input<string>;
    /**
     * Google Cloud region for PySpark workloads on Dataproc
     */
    dataprocRegion?: pulumi.Input<string>;
    /**
     * Project to bill for query execution
     */
    executionProject?: pulumi.Input<string>;
    /**
     * GCP project ID
     */
    gcpProjectId: pulumi.Input<string>;
    /**
     * URI for a Google Cloud Storage bucket to host Python code executed via Datapro
     */
    gcsBucket?: pulumi.Input<string>;
    /**
     * Whether the connection is active
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * Location to create new Datasets in
     */
    location?: pulumi.Input<string>;
    /**
     * Max number of bytes that can be billed for a given BigQuery query
     */
    maximumBytesBilled?: pulumi.Input<number>;
    /**
     * Connection name
     */
    name?: pulumi.Input<string>;
    /**
     * The priority with which to execute BigQuery queries (batch or interactive)
     */
    priority?: pulumi.Input<string>;
    /**
     * Private key of the Service Account
     */
    privateKey: pulumi.Input<string>;
    /**
     * Private key ID of the Service Account
     */
    privateKeyId: pulumi.Input<string>;
    /**
     * Project ID to create the connection in
     */
    projectId: pulumi.Input<number>;
    /**
     * Number of retries for queries
     */
    retries?: pulumi.Input<number>;
    /**
     * Timeout in seconds for queries
     */
    timeoutSeconds: pulumi.Input<number>;
    /**
     * Token URI for the Service Account
     */
    tokenUri: pulumi.Input<string>;
    /**
     * The type of connection
     */
    type: pulumi.Input<string>;
}