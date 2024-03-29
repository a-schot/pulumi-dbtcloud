// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetGroupUsersUser {
    email: string;
    id: number;
}

export interface GetJobJobCompletionTriggerCondition {
    jobId: number;
    projectId: number;
    statuses: string[];
}

export interface GetServiceTokenServiceTokenPermission {
    allProjects: boolean;
    permissionSet: string;
    projectId: number;
}

export interface GroupGroupPermission {
    /**
     * Whether or not to apply this permission to all projects for this group
     */
    allProjects: boolean;
    /**
     * Set of permissions to apply
     */
    permissionSet: string;
    /**
     * Project ID to apply this permission to for this group
     */
    projectId?: number;
}

export interface JobJobCompletionTriggerCondition {
    /**
     * The ID of the job that would trigger this job after completion.
     */
    jobId: number;
    /**
     * The ID of the project where the trigger job is running in.
     */
    projectId: number;
    /**
     * List of statuses to trigger the job on. Possible values are `success`, `error` and `canceled`.
     */
    statuses: string[];
}

export interface LegacyGetJobJobCompletionTriggerCondition {
    jobId: number;
    projectId: number;
    statuses: string[];
}

export interface LegacyGetServiceTokenServiceTokenPermission {
    allProjects: boolean;
    permissionSet: string;
    projectId: number;
}

export interface LegacyGroupGroupPermission {
    allProjects: boolean;
    permissionSet: string;
    projectId?: number;
}

export interface LegacyJobJobCompletionTriggerCondition {
    jobId: number;
    projectId: number;
    statuses: string[];
}

export interface LegacyServiceTokenServiceTokenPermission {
    allProjects: boolean;
    permissionSet: string;
    projectId?: number;
}

export interface ServiceTokenServiceTokenPermission {
    /**
     * Whether or not to apply this permission to all projects for this service token
     */
    allProjects: boolean;
    /**
     * Set of permissions to apply
     */
    permissionSet: string;
    /**
     * Project ID to apply this permission to for this service token
     */
    projectId?: number;
}

