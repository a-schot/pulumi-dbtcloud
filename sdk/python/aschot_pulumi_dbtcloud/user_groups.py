# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserGroupsArgs', 'UserGroups']

@pulumi.input_type
class UserGroupsArgs:
    def __init__(__self__, *,
                 group_ids: pulumi.Input[Sequence[pulumi.Input[int]]],
                 user_id: pulumi.Input[int]):
        """
        The set of arguments for constructing a UserGroups resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] group_ids: IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
        :param pulumi.Input[int] user_id: The internal ID of a dbt Cloud user
        """
        pulumi.set(__self__, "group_ids", group_ids)
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
        """
        return pulumi.get(self, "group_ids")

    @group_ids.setter
    def group_ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "group_ids", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[int]:
        """
        The internal ID of a dbt Cloud user
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "user_id", value)


@pulumi.input_type
class _UserGroupsState:
    def __init__(__self__, *,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 user_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering UserGroups resources.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] group_ids: IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
        :param pulumi.Input[int] user_id: The internal ID of a dbt Cloud user
        """
        if group_ids is not None:
            pulumi.set(__self__, "group_ids", group_ids)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
        """
        return pulumi.get(self, "group_ids")

    @group_ids.setter
    def group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "group_ids", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[int]]:
        """
        The internal ID of a dbt Cloud user
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "user_id", value)


class UserGroups(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 user_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Assigns a set of dbt Cloud groups to a given User ID.

        > If additional groups were assigned manually in dbt Cloud, they will be removed. The full list of groups need to be provided as config.

        > This resource does not currently support deletion (e.g. a deleted resource will stay as-is in dbt Cloud).
        This is intentional in order to prevent accidental deletion of all users groups assigned to a user.
        If you would like a different behavior, please open an issue on GitHub. To remove all groups for a user, set "group_ids" to the empty set "[]".

        ## Example Usage

        ```python
        import pulumi
        import aschot_pulumi_dbtcloud as dbtcloud

        # we can assign groups to users
        my_user_groups = dbtcloud.UserGroups("myUserGroups",
            user_id=dbtcloud_user["my_user"]["id"],
            group_ids=[
                1234,
                dbtcloud_group["my_group"]["id"],
                local["my_group_id"],
            ])
        # as Delete is not handled currently, by design, removing all groups from a user can be done with
        my_other_user_groups = dbtcloud.UserGroups("myOtherUserGroups",
            user_id=123456,
            group_ids=[])
        ```

        ## Import

        Import using the User ID The User ID can be retrieved from the dbt Cloud UI or with the data source dbtcloud_user

        ```sh
         $ pulumi import dbtcloud:index/userGroups:UserGroups my_user_groups "user_id"
        ```

        ```sh
         $ pulumi import dbtcloud:index/userGroups:UserGroups my_user_groups 123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] group_ids: IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
        :param pulumi.Input[int] user_id: The internal ID of a dbt Cloud user
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserGroupsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Assigns a set of dbt Cloud groups to a given User ID.

        > If additional groups were assigned manually in dbt Cloud, they will be removed. The full list of groups need to be provided as config.

        > This resource does not currently support deletion (e.g. a deleted resource will stay as-is in dbt Cloud).
        This is intentional in order to prevent accidental deletion of all users groups assigned to a user.
        If you would like a different behavior, please open an issue on GitHub. To remove all groups for a user, set "group_ids" to the empty set "[]".

        ## Example Usage

        ```python
        import pulumi
        import aschot_pulumi_dbtcloud as dbtcloud

        # we can assign groups to users
        my_user_groups = dbtcloud.UserGroups("myUserGroups",
            user_id=dbtcloud_user["my_user"]["id"],
            group_ids=[
                1234,
                dbtcloud_group["my_group"]["id"],
                local["my_group_id"],
            ])
        # as Delete is not handled currently, by design, removing all groups from a user can be done with
        my_other_user_groups = dbtcloud.UserGroups("myOtherUserGroups",
            user_id=123456,
            group_ids=[])
        ```

        ## Import

        Import using the User ID The User ID can be retrieved from the dbt Cloud UI or with the data source dbtcloud_user

        ```sh
         $ pulumi import dbtcloud:index/userGroups:UserGroups my_user_groups "user_id"
        ```

        ```sh
         $ pulumi import dbtcloud:index/userGroups:UserGroups my_user_groups 123456
        ```

        :param str resource_name: The name of the resource.
        :param UserGroupsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserGroupsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 user_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserGroupsArgs.__new__(UserGroupsArgs)

            if group_ids is None and not opts.urn:
                raise TypeError("Missing required property 'group_ids'")
            __props__.__dict__["group_ids"] = group_ids
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(UserGroups, __self__).__init__(
            'dbtcloud:index/userGroups:UserGroups',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            user_id: Optional[pulumi.Input[int]] = None) -> 'UserGroups':
        """
        Get an existing UserGroups resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] group_ids: IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
        :param pulumi.Input[int] user_id: The internal ID of a dbt Cloud user
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserGroupsState.__new__(_UserGroupsState)

        __props__.__dict__["group_ids"] = group_ids
        __props__.__dict__["user_id"] = user_id
        return UserGroups(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> pulumi.Output[Sequence[int]]:
        """
        IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
        """
        return pulumi.get(self, "group_ids")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[int]:
        """
        The internal ID of a dbt Cloud user
        """
        return pulumi.get(self, "user_id")
