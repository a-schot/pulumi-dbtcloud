# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GroupGroupPermissionArgs',
    'ServiceTokenServiceTokenPermissionArgs',
]

@pulumi.input_type
class GroupGroupPermissionArgs:
    def __init__(__self__, *,
                 all_projects: pulumi.Input[bool],
                 permission_set: pulumi.Input[str],
                 project_id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[bool] all_projects: Whether or not to apply this permission to all projects for this group
        :param pulumi.Input[str] permission_set: Set of permissions to apply
        :param pulumi.Input[int] project_id: Project ID to apply this permission to for this group
        """
        pulumi.set(__self__, "all_projects", all_projects)
        pulumi.set(__self__, "permission_set", permission_set)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="allProjects")
    def all_projects(self) -> pulumi.Input[bool]:
        """
        Whether or not to apply this permission to all projects for this group
        """
        return pulumi.get(self, "all_projects")

    @all_projects.setter
    def all_projects(self, value: pulumi.Input[bool]):
        pulumi.set(self, "all_projects", value)

    @property
    @pulumi.getter(name="permissionSet")
    def permission_set(self) -> pulumi.Input[str]:
        """
        Set of permissions to apply
        """
        return pulumi.get(self, "permission_set")

    @permission_set.setter
    def permission_set(self, value: pulumi.Input[str]):
        pulumi.set(self, "permission_set", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Project ID to apply this permission to for this group
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class ServiceTokenServiceTokenPermissionArgs:
    def __init__(__self__, *,
                 all_projects: pulumi.Input[bool],
                 permission_set: pulumi.Input[str],
                 project_id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[bool] all_projects: Whether or not to apply this permission to all projects for this service token
        :param pulumi.Input[str] permission_set: Set of permissions to apply
        :param pulumi.Input[int] project_id: Project ID to apply this permission to for this service token
        """
        pulumi.set(__self__, "all_projects", all_projects)
        pulumi.set(__self__, "permission_set", permission_set)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="allProjects")
    def all_projects(self) -> pulumi.Input[bool]:
        """
        Whether or not to apply this permission to all projects for this service token
        """
        return pulumi.get(self, "all_projects")

    @all_projects.setter
    def all_projects(self, value: pulumi.Input[bool]):
        pulumi.set(self, "all_projects", value)

    @property
    @pulumi.getter(name="permissionSet")
    def permission_set(self) -> pulumi.Input[str]:
        """
        Set of permissions to apply
        """
        return pulumi.get(self, "permission_set")

    @permission_set.setter
    def permission_set(self, value: pulumi.Input[str]):
        pulumi.set(self, "permission_set", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Project ID to apply this permission to for this service token
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)


