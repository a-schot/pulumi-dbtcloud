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
    'GroupGroupPermission',
    'ServiceTokenServiceTokenPermission',
    'GetGroupUsersUserResult',
    'GetServiceTokenServiceTokenPermissionResult',
]

@pulumi.output_type
class GroupGroupPermission(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allProjects":
            suggest = "all_projects"
        elif key == "permissionSet":
            suggest = "permission_set"
        elif key == "projectId":
            suggest = "project_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GroupGroupPermission. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GroupGroupPermission.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GroupGroupPermission.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 all_projects: bool,
                 permission_set: str,
                 project_id: Optional[int] = None):
        """
        :param bool all_projects: Whether or not to apply this permission to all projects for this group
        :param str permission_set: Set of permissions to apply
        :param int project_id: Project ID to apply this permission to for this group
        """
        pulumi.set(__self__, "all_projects", all_projects)
        pulumi.set(__self__, "permission_set", permission_set)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="allProjects")
    def all_projects(self) -> bool:
        """
        Whether or not to apply this permission to all projects for this group
        """
        return pulumi.get(self, "all_projects")

    @property
    @pulumi.getter(name="permissionSet")
    def permission_set(self) -> str:
        """
        Set of permissions to apply
        """
        return pulumi.get(self, "permission_set")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[int]:
        """
        Project ID to apply this permission to for this group
        """
        return pulumi.get(self, "project_id")


@pulumi.output_type
class ServiceTokenServiceTokenPermission(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allProjects":
            suggest = "all_projects"
        elif key == "permissionSet":
            suggest = "permission_set"
        elif key == "projectId":
            suggest = "project_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServiceTokenServiceTokenPermission. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServiceTokenServiceTokenPermission.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServiceTokenServiceTokenPermission.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 all_projects: bool,
                 permission_set: str,
                 project_id: Optional[int] = None):
        """
        :param bool all_projects: Whether or not to apply this permission to all projects for this service token
        :param str permission_set: Set of permissions to apply
        :param int project_id: Project ID to apply this permission to for this service token
        """
        pulumi.set(__self__, "all_projects", all_projects)
        pulumi.set(__self__, "permission_set", permission_set)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="allProjects")
    def all_projects(self) -> bool:
        """
        Whether or not to apply this permission to all projects for this service token
        """
        return pulumi.get(self, "all_projects")

    @property
    @pulumi.getter(name="permissionSet")
    def permission_set(self) -> str:
        """
        Set of permissions to apply
        """
        return pulumi.get(self, "permission_set")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[int]:
        """
        Project ID to apply this permission to for this service token
        """
        return pulumi.get(self, "project_id")


@pulumi.output_type
class GetGroupUsersUserResult(dict):
    def __init__(__self__, *,
                 email: str,
                 id: int):
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> int:
        return pulumi.get(self, "id")


@pulumi.output_type
class GetServiceTokenServiceTokenPermissionResult(dict):
    def __init__(__self__, *,
                 all_projects: bool,
                 permission_set: str,
                 project_id: int):
        pulumi.set(__self__, "all_projects", all_projects)
        pulumi.set(__self__, "permission_set", permission_set)
        pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="allProjects")
    def all_projects(self) -> bool:
        return pulumi.get(self, "all_projects")

    @property
    @pulumi.getter(name="permissionSet")
    def permission_set(self) -> str:
        return pulumi.get(self, "permission_set")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        return pulumi.get(self, "project_id")


