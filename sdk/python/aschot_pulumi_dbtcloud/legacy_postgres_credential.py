# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['LegacyPostgresCredentialArgs', 'LegacyPostgresCredential']

@pulumi.input_type
class LegacyPostgresCredentialArgs:
    def __init__(__self__, *,
                 default_schema: pulumi.Input[str],
                 project_id: pulumi.Input[int],
                 type: pulumi.Input[str],
                 username: pulumi.Input[str],
                 is_active: Optional[pulumi.Input[bool]] = None,
                 num_threads: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 target_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LegacyPostgresCredential resource.
        :param pulumi.Input[str] default_schema: Default schema name
        :param pulumi.Input[int] project_id: Project ID to create the Postgres/Redshift/AlloyDB credential in
        :param pulumi.Input[str] type: Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
        :param pulumi.Input[str] username: Username for Postgres/Redshift/AlloyDB
        :param pulumi.Input[bool] is_active: Whether the Postgres/Redshift/AlloyDB credential is active
        :param pulumi.Input[int] num_threads: Number of threads to use
        :param pulumi.Input[str] password: Password for Postgres/Redshift/AlloyDB
        :param pulumi.Input[str] target_name: Default schema name
        """
        pulumi.set(__self__, "default_schema", default_schema)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "username", username)
        if is_active is not None:
            pulumi.set(__self__, "is_active", is_active)
        if num_threads is not None:
            pulumi.set(__self__, "num_threads", num_threads)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if target_name is not None:
            pulumi.set(__self__, "target_name", target_name)

    @property
    @pulumi.getter(name="defaultSchema")
    def default_schema(self) -> pulumi.Input[str]:
        """
        Default schema name
        """
        return pulumi.get(self, "default_schema")

    @default_schema.setter
    def default_schema(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_schema", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[int]:
        """
        Project ID to create the Postgres/Redshift/AlloyDB credential in
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        Username for Postgres/Redshift/AlloyDB
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the Postgres/Redshift/AlloyDB credential is active
        """
        return pulumi.get(self, "is_active")

    @is_active.setter
    def is_active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_active", value)

    @property
    @pulumi.getter(name="numThreads")
    def num_threads(self) -> Optional[pulumi.Input[int]]:
        """
        Number of threads to use
        """
        return pulumi.get(self, "num_threads")

    @num_threads.setter
    def num_threads(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "num_threads", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password for Postgres/Redshift/AlloyDB
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="targetName")
    def target_name(self) -> Optional[pulumi.Input[str]]:
        """
        Default schema name
        """
        return pulumi.get(self, "target_name")

    @target_name.setter
    def target_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_name", value)


@pulumi.input_type
class _LegacyPostgresCredentialState:
    def __init__(__self__, *,
                 credential_id: Optional[pulumi.Input[int]] = None,
                 default_schema: Optional[pulumi.Input[str]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 num_threads: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 target_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LegacyPostgresCredential resources.
        :param pulumi.Input[int] credential_id: The system Postgres/Redshift/AlloyDB credential ID
        :param pulumi.Input[str] default_schema: Default schema name
        :param pulumi.Input[bool] is_active: Whether the Postgres/Redshift/AlloyDB credential is active
        :param pulumi.Input[int] num_threads: Number of threads to use
        :param pulumi.Input[str] password: Password for Postgres/Redshift/AlloyDB
        :param pulumi.Input[int] project_id: Project ID to create the Postgres/Redshift/AlloyDB credential in
        :param pulumi.Input[str] target_name: Default schema name
        :param pulumi.Input[str] type: Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
        :param pulumi.Input[str] username: Username for Postgres/Redshift/AlloyDB
        """
        if credential_id is not None:
            pulumi.set(__self__, "credential_id", credential_id)
        if default_schema is not None:
            pulumi.set(__self__, "default_schema", default_schema)
        if is_active is not None:
            pulumi.set(__self__, "is_active", is_active)
        if num_threads is not None:
            pulumi.set(__self__, "num_threads", num_threads)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if target_name is not None:
            pulumi.set(__self__, "target_name", target_name)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="credentialId")
    def credential_id(self) -> Optional[pulumi.Input[int]]:
        """
        The system Postgres/Redshift/AlloyDB credential ID
        """
        return pulumi.get(self, "credential_id")

    @credential_id.setter
    def credential_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "credential_id", value)

    @property
    @pulumi.getter(name="defaultSchema")
    def default_schema(self) -> Optional[pulumi.Input[str]]:
        """
        Default schema name
        """
        return pulumi.get(self, "default_schema")

    @default_schema.setter
    def default_schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_schema", value)

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the Postgres/Redshift/AlloyDB credential is active
        """
        return pulumi.get(self, "is_active")

    @is_active.setter
    def is_active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_active", value)

    @property
    @pulumi.getter(name="numThreads")
    def num_threads(self) -> Optional[pulumi.Input[int]]:
        """
        Number of threads to use
        """
        return pulumi.get(self, "num_threads")

    @num_threads.setter
    def num_threads(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "num_threads", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password for Postgres/Redshift/AlloyDB
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Project ID to create the Postgres/Redshift/AlloyDB credential in
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="targetName")
    def target_name(self) -> Optional[pulumi.Input[str]]:
        """
        Default schema name
        """
        return pulumi.get(self, "target_name")

    @target_name.setter
    def target_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        Username for Postgres/Redshift/AlloyDB
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


warnings.warn("""Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource""", DeprecationWarning)


class LegacyPostgresCredential(pulumi.CustomResource):
    warnings.warn("""Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_schema: Optional[pulumi.Input[str]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 num_threads: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 target_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a LegacyPostgresCredential resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_schema: Default schema name
        :param pulumi.Input[bool] is_active: Whether the Postgres/Redshift/AlloyDB credential is active
        :param pulumi.Input[int] num_threads: Number of threads to use
        :param pulumi.Input[str] password: Password for Postgres/Redshift/AlloyDB
        :param pulumi.Input[int] project_id: Project ID to create the Postgres/Redshift/AlloyDB credential in
        :param pulumi.Input[str] target_name: Default schema name
        :param pulumi.Input[str] type: Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
        :param pulumi.Input[str] username: Username for Postgres/Redshift/AlloyDB
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LegacyPostgresCredentialArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LegacyPostgresCredential resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LegacyPostgresCredentialArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LegacyPostgresCredentialArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_schema: Optional[pulumi.Input[str]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 num_threads: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 target_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""LegacyPostgresCredential is deprecated: Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LegacyPostgresCredentialArgs.__new__(LegacyPostgresCredentialArgs)

            if default_schema is None and not opts.urn:
                raise TypeError("Missing required property 'default_schema'")
            __props__.__dict__["default_schema"] = default_schema
            __props__.__dict__["is_active"] = is_active
            __props__.__dict__["num_threads"] = num_threads
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["target_name"] = target_name
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
            __props__.__dict__["credential_id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(LegacyPostgresCredential, __self__).__init__(
            'dbtcloud:index/legacyPostgresCredential:LegacyPostgresCredential',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            credential_id: Optional[pulumi.Input[int]] = None,
            default_schema: Optional[pulumi.Input[str]] = None,
            is_active: Optional[pulumi.Input[bool]] = None,
            num_threads: Optional[pulumi.Input[int]] = None,
            password: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            target_name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'LegacyPostgresCredential':
        """
        Get an existing LegacyPostgresCredential resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] credential_id: The system Postgres/Redshift/AlloyDB credential ID
        :param pulumi.Input[str] default_schema: Default schema name
        :param pulumi.Input[bool] is_active: Whether the Postgres/Redshift/AlloyDB credential is active
        :param pulumi.Input[int] num_threads: Number of threads to use
        :param pulumi.Input[str] password: Password for Postgres/Redshift/AlloyDB
        :param pulumi.Input[int] project_id: Project ID to create the Postgres/Redshift/AlloyDB credential in
        :param pulumi.Input[str] target_name: Default schema name
        :param pulumi.Input[str] type: Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
        :param pulumi.Input[str] username: Username for Postgres/Redshift/AlloyDB
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LegacyPostgresCredentialState.__new__(_LegacyPostgresCredentialState)

        __props__.__dict__["credential_id"] = credential_id
        __props__.__dict__["default_schema"] = default_schema
        __props__.__dict__["is_active"] = is_active
        __props__.__dict__["num_threads"] = num_threads
        __props__.__dict__["password"] = password
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["target_name"] = target_name
        __props__.__dict__["type"] = type
        __props__.__dict__["username"] = username
        return LegacyPostgresCredential(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="credentialId")
    def credential_id(self) -> pulumi.Output[int]:
        """
        The system Postgres/Redshift/AlloyDB credential ID
        """
        return pulumi.get(self, "credential_id")

    @property
    @pulumi.getter(name="defaultSchema")
    def default_schema(self) -> pulumi.Output[str]:
        """
        Default schema name
        """
        return pulumi.get(self, "default_schema")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the Postgres/Redshift/AlloyDB credential is active
        """
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter(name="numThreads")
    def num_threads(self) -> pulumi.Output[Optional[int]]:
        """
        Number of threads to use
        """
        return pulumi.get(self, "num_threads")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Password for Postgres/Redshift/AlloyDB
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[int]:
        """
        Project ID to create the Postgres/Redshift/AlloyDB credential in
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="targetName")
    def target_name(self) -> pulumi.Output[Optional[str]]:
        """
        Default schema name
        """
        return pulumi.get(self, "target_name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        Username for Postgres/Redshift/AlloyDB
        """
        return pulumi.get(self, "username")

