// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * @deprecated Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource
 */
export class LegacyConnection extends pulumi.CustomResource {
    /**
     * Get an existing LegacyConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LegacyConnectionState, opts?: pulumi.CustomResourceOptions): LegacyConnection {
        pulumi.log.warn("LegacyConnection is deprecated: Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource")
        return new LegacyConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'dbtcloud:index/legacyConnection:LegacyConnection';

    /**
     * Returns true if the given object is an instance of LegacyConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LegacyConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LegacyConnection.__pulumiType;
    }

    /**
     * Account name for the connection
     */
    public readonly account!: pulumi.Output<string | undefined>;
    /**
     * Adapter id created for the Databricks connection
     */
    public /*out*/ readonly adapterId!: pulumi.Output<number>;
    /**
     * Whether or not the connection should allow client session keep alive
     */
    public readonly allowKeepAlive!: pulumi.Output<boolean | undefined>;
    /**
     * Whether or not the connection should allow SSO
     */
    public readonly allowSso!: pulumi.Output<boolean | undefined>;
    /**
     * Catalog name if Unity Catalog is enabled in your Databricks workspace
     */
    public readonly catalog!: pulumi.Output<string | undefined>;
    /**
     * Connection Identifier
     */
    public /*out*/ readonly connectionId!: pulumi.Output<number>;
    /**
     * Database name for the connection
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * Host name for the connection, including Databricks cluster
     */
    public readonly hostName!: pulumi.Output<string | undefined>;
    /**
     * The HTTP path of the Databricks cluster or SQL warehouse
     */
    public readonly httpPath!: pulumi.Output<string | undefined>;
    /**
     * Whether the connection is active
     */
    public readonly isActive!: pulumi.Output<boolean | undefined>;
    /**
     * Connection name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * OAuth client identifier
     */
    public readonly oauthClientId!: pulumi.Output<string | undefined>;
    /**
     * OAuth client secret
     */
    public readonly oauthClientSecret!: pulumi.Output<string | undefined>;
    /**
     * Port number to connect via
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The ID of the PrivateLink connection. This ID can be found using the `privatelink_endpoint` data source
     */
    public readonly privateLinkEndpointId!: pulumi.Output<string | undefined>;
    /**
     * Project ID to create the connection in
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * Role name for the connection
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * Whether or not tunneling should be enabled on your database connection
     */
    public readonly tunnelEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The type of connection
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Warehouse name for the connection
     */
    public readonly warehouse!: pulumi.Output<string | undefined>;

    /**
     * Create a LegacyConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource */
    constructor(name: string, args: LegacyConnectionArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource */
    constructor(name: string, argsOrState?: LegacyConnectionArgs | LegacyConnectionState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LegacyConnection is deprecated: Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LegacyConnectionState | undefined;
            resourceInputs["account"] = state ? state.account : undefined;
            resourceInputs["adapterId"] = state ? state.adapterId : undefined;
            resourceInputs["allowKeepAlive"] = state ? state.allowKeepAlive : undefined;
            resourceInputs["allowSso"] = state ? state.allowSso : undefined;
            resourceInputs["catalog"] = state ? state.catalog : undefined;
            resourceInputs["connectionId"] = state ? state.connectionId : undefined;
            resourceInputs["database"] = state ? state.database : undefined;
            resourceInputs["hostName"] = state ? state.hostName : undefined;
            resourceInputs["httpPath"] = state ? state.httpPath : undefined;
            resourceInputs["isActive"] = state ? state.isActive : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["oauthClientId"] = state ? state.oauthClientId : undefined;
            resourceInputs["oauthClientSecret"] = state ? state.oauthClientSecret : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["privateLinkEndpointId"] = state ? state.privateLinkEndpointId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["tunnelEnabled"] = state ? state.tunnelEnabled : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["warehouse"] = state ? state.warehouse : undefined;
        } else {
            const args = argsOrState as LegacyConnectionArgs | undefined;
            if ((!args || args.database === undefined) && !opts.urn) {
                throw new Error("Missing required property 'database'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["account"] = args ? args.account : undefined;
            resourceInputs["allowKeepAlive"] = args ? args.allowKeepAlive : undefined;
            resourceInputs["allowSso"] = args ? args.allowSso : undefined;
            resourceInputs["catalog"] = args ? args.catalog : undefined;
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["hostName"] = args ? args.hostName : undefined;
            resourceInputs["httpPath"] = args ? args.httpPath : undefined;
            resourceInputs["isActive"] = args ? args.isActive : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oauthClientId"] = args ? args.oauthClientId : undefined;
            resourceInputs["oauthClientSecret"] = args ? args.oauthClientSecret : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["privateLinkEndpointId"] = args ? args.privateLinkEndpointId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["tunnelEnabled"] = args ? args.tunnelEnabled : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["warehouse"] = args ? args.warehouse : undefined;
            resourceInputs["adapterId"] = undefined /*out*/;
            resourceInputs["connectionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LegacyConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LegacyConnection resources.
 */
export interface LegacyConnectionState {
    /**
     * Account name for the connection
     */
    account?: pulumi.Input<string>;
    /**
     * Adapter id created for the Databricks connection
     */
    adapterId?: pulumi.Input<number>;
    /**
     * Whether or not the connection should allow client session keep alive
     */
    allowKeepAlive?: pulumi.Input<boolean>;
    /**
     * Whether or not the connection should allow SSO
     */
    allowSso?: pulumi.Input<boolean>;
    /**
     * Catalog name if Unity Catalog is enabled in your Databricks workspace
     */
    catalog?: pulumi.Input<string>;
    /**
     * Connection Identifier
     */
    connectionId?: pulumi.Input<number>;
    /**
     * Database name for the connection
     */
    database?: pulumi.Input<string>;
    /**
     * Host name for the connection, including Databricks cluster
     */
    hostName?: pulumi.Input<string>;
    /**
     * The HTTP path of the Databricks cluster or SQL warehouse
     */
    httpPath?: pulumi.Input<string>;
    /**
     * Whether the connection is active
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * Connection name
     */
    name?: pulumi.Input<string>;
    /**
     * OAuth client identifier
     */
    oauthClientId?: pulumi.Input<string>;
    /**
     * OAuth client secret
     */
    oauthClientSecret?: pulumi.Input<string>;
    /**
     * Port number to connect via
     */
    port?: pulumi.Input<number>;
    /**
     * The ID of the PrivateLink connection. This ID can be found using the `privatelink_endpoint` data source
     */
    privateLinkEndpointId?: pulumi.Input<string>;
    /**
     * Project ID to create the connection in
     */
    projectId?: pulumi.Input<number>;
    /**
     * Role name for the connection
     */
    role?: pulumi.Input<string>;
    /**
     * Whether or not tunneling should be enabled on your database connection
     */
    tunnelEnabled?: pulumi.Input<boolean>;
    /**
     * The type of connection
     */
    type?: pulumi.Input<string>;
    /**
     * Warehouse name for the connection
     */
    warehouse?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LegacyConnection resource.
 */
export interface LegacyConnectionArgs {
    /**
     * Account name for the connection
     */
    account?: pulumi.Input<string>;
    /**
     * Whether or not the connection should allow client session keep alive
     */
    allowKeepAlive?: pulumi.Input<boolean>;
    /**
     * Whether or not the connection should allow SSO
     */
    allowSso?: pulumi.Input<boolean>;
    /**
     * Catalog name if Unity Catalog is enabled in your Databricks workspace
     */
    catalog?: pulumi.Input<string>;
    /**
     * Database name for the connection
     */
    database: pulumi.Input<string>;
    /**
     * Host name for the connection, including Databricks cluster
     */
    hostName?: pulumi.Input<string>;
    /**
     * The HTTP path of the Databricks cluster or SQL warehouse
     */
    httpPath?: pulumi.Input<string>;
    /**
     * Whether the connection is active
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * Connection name
     */
    name?: pulumi.Input<string>;
    /**
     * OAuth client identifier
     */
    oauthClientId?: pulumi.Input<string>;
    /**
     * OAuth client secret
     */
    oauthClientSecret?: pulumi.Input<string>;
    /**
     * Port number to connect via
     */
    port?: pulumi.Input<number>;
    /**
     * The ID of the PrivateLink connection. This ID can be found using the `privatelink_endpoint` data source
     */
    privateLinkEndpointId?: pulumi.Input<string>;
    /**
     * Project ID to create the connection in
     */
    projectId: pulumi.Input<number>;
    /**
     * Role name for the connection
     */
    role?: pulumi.Input<string>;
    /**
     * Whether or not tunneling should be enabled on your database connection
     */
    tunnelEnabled?: pulumi.Input<boolean>;
    /**
     * The type of connection
     */
    type: pulumi.Input<string>;
    /**
     * Warehouse name for the connection
     */
    warehouse?: pulumi.Input<string>;
}
