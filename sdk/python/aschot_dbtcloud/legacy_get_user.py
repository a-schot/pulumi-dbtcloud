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
    'LegacyGetUserResult',
    'AwaitableLegacyGetUserResult',
    'legacy_get_user',
    'legacy_get_user_output',
]

@pulumi.output_type
class LegacyGetUserResult:
    """
    A collection of values returned by LegacyGetUser.
    """
    def __init__(__self__, email=None, id=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> int:
        return pulumi.get(self, "id")


class AwaitableLegacyGetUserResult(LegacyGetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return LegacyGetUserResult(
            email=self.email,
            id=self.id)


def legacy_get_user(email: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableLegacyGetUserResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['email'] = email
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('dbtcloud:index/legacyGetUser:LegacyGetUser', __args__, opts=opts, typ=LegacyGetUserResult).value

    return AwaitableLegacyGetUserResult(
        email=pulumi.get(__ret__, 'email'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(legacy_get_user)
def legacy_get_user_output(email: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[LegacyGetUserResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
