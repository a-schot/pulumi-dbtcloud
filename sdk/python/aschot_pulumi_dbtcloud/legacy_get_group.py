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
    'LegacyGetGroupResult',
    'AwaitableLegacyGetGroupResult',
    'legacy_get_group',
    'legacy_get_group_output',
]

@pulumi.output_type
class LegacyGetGroupResult:
    """
    A collection of values returned by LegacyGetGroup.
    """
    def __init__(__self__, assign_by_default=None, group_id=None, id=None, is_active=None, name=None, sso_mapping_groups=None):
        if assign_by_default and not isinstance(assign_by_default, bool):
            raise TypeError("Expected argument 'assign_by_default' to be a bool")
        pulumi.set(__self__, "assign_by_default", assign_by_default)
        if group_id and not isinstance(group_id, int):
            raise TypeError("Expected argument 'group_id' to be a int")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_active and not isinstance(is_active, bool):
            raise TypeError("Expected argument 'is_active' to be a bool")
        pulumi.set(__self__, "is_active", is_active)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if sso_mapping_groups and not isinstance(sso_mapping_groups, list):
            raise TypeError("Expected argument 'sso_mapping_groups' to be a list")
        pulumi.set(__self__, "sso_mapping_groups", sso_mapping_groups)

    @property
    @pulumi.getter(name="assignByDefault")
    def assign_by_default(self) -> bool:
        return pulumi.get(self, "assign_by_default")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> int:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> bool:
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ssoMappingGroups")
    def sso_mapping_groups(self) -> Sequence[str]:
        return pulumi.get(self, "sso_mapping_groups")


class AwaitableLegacyGetGroupResult(LegacyGetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return LegacyGetGroupResult(
            assign_by_default=self.assign_by_default,
            group_id=self.group_id,
            id=self.id,
            is_active=self.is_active,
            name=self.name,
            sso_mapping_groups=self.sso_mapping_groups)


def legacy_get_group(group_id: Optional[int] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableLegacyGetGroupResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['groupId'] = group_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('dbtcloud:index/legacyGetGroup:LegacyGetGroup', __args__, opts=opts, typ=LegacyGetGroupResult).value

    return AwaitableLegacyGetGroupResult(
        assign_by_default=pulumi.get(__ret__, 'assign_by_default'),
        group_id=pulumi.get(__ret__, 'group_id'),
        id=pulumi.get(__ret__, 'id'),
        is_active=pulumi.get(__ret__, 'is_active'),
        name=pulumi.get(__ret__, 'name'),
        sso_mapping_groups=pulumi.get(__ret__, 'sso_mapping_groups'))


@_utilities.lift_output_func(legacy_get_group)
def legacy_get_group_output(group_id: Optional[pulumi.Input[int]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[LegacyGetGroupResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
