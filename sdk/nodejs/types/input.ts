// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GroupGroupPermission {
    /**
     * Whether or not to apply this permission to all projects for this group
     */
    allProjects: pulumi.Input<boolean>;
    /**
     * Set of permissions to apply
     */
    permissionSet: pulumi.Input<string>;
    /**
     * Project ID to apply this permission to for this group
     */
    projectId?: pulumi.Input<number>;
}

export interface LegacyGroupGroupPermission {
    allProjects: pulumi.Input<boolean>;
    permissionSet: pulumi.Input<string>;
    projectId?: pulumi.Input<number>;
}

export interface LegacyServiceTokenServiceTokenPermission {
    allProjects: pulumi.Input<boolean>;
    permissionSet: pulumi.Input<string>;
    projectId?: pulumi.Input<number>;
}

export interface ServiceTokenServiceTokenPermission {
    /**
     * Whether or not to apply this permission to all projects for this service token
     */
    allProjects: pulumi.Input<boolean>;
    /**
     * Set of permissions to apply
     */
    permissionSet: pulumi.Input<string>;
    /**
     * Project ID to apply this permission to for this service token
     */
    projectId?: pulumi.Input<number>;
}
