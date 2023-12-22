# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['EnvironmentVariableArgs', 'EnvironmentVariable']

@pulumi.input_type
class EnvironmentVariableArgs:
    def __init__(__self__, *,
                 environment_values: pulumi.Input[Mapping[str, Any]],
                 project_id: pulumi.Input[int],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EnvironmentVariable resource.
        :param pulumi.Input[Mapping[str, Any]] environment_values: Map from environment names to respective variable value, a special key `project` should be set for the project default variable value
        :param pulumi.Input[int] project_id: Project for the variable to be created in
        :param pulumi.Input[str] name: Name for the variable, must be unique within a project, must be prefixed with 'DBT_'
        """
        pulumi.set(__self__, "environment_values", environment_values)
        pulumi.set(__self__, "project_id", project_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="environmentValues")
    def environment_values(self) -> pulumi.Input[Mapping[str, Any]]:
        """
        Map from environment names to respective variable value, a special key `project` should be set for the project default variable value
        """
        return pulumi.get(self, "environment_values")

    @environment_values.setter
    def environment_values(self, value: pulumi.Input[Mapping[str, Any]]):
        pulumi.set(self, "environment_values", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[int]:
        """
        Project for the variable to be created in
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name for the variable, must be unique within a project, must be prefixed with 'DBT_'
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _EnvironmentVariableState:
    def __init__(__self__, *,
                 environment_values: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering EnvironmentVariable resources.
        :param pulumi.Input[Mapping[str, Any]] environment_values: Map from environment names to respective variable value, a special key `project` should be set for the project default variable value
        :param pulumi.Input[str] name: Name for the variable, must be unique within a project, must be prefixed with 'DBT_'
        :param pulumi.Input[int] project_id: Project for the variable to be created in
        """
        if environment_values is not None:
            pulumi.set(__self__, "environment_values", environment_values)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="environmentValues")
    def environment_values(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map from environment names to respective variable value, a special key `project` should be set for the project default variable value
        """
        return pulumi.get(self, "environment_values")

    @environment_values.setter
    def environment_values(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "environment_values", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name for the variable, must be unique within a project, must be prefixed with 'DBT_'
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Project for the variable to be created in
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)


class EnvironmentVariable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environment_values: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        *Note*: Some upstream resources can be slow to create, so if creating a project or environment at
        the same time as the environment variables, it's recommended to use the `depends_on` meta argument.

        ## Example Usage

        ```python
        import pulumi
        import aschot_dbtcloud as dbtcloud

        # NOTE for customers using the LEGACY dbt_cloud provider:
        # use dbt_cloud_environment_variable instead of dbtcloud_environment_variable for the legacy resource names
        # legacy names will be removed from 0.3 onwards
        dbt_my_env_var = dbtcloud.EnvironmentVariable("dbtMyEnvVar",
            project_id=dbtcloud_project["dbt_project"]["id"],
            environment_values={
                "project": "my_project_level_value",
                "Dev": "my_env_level_value",
                "CI": "my_ci_override_value",
                "Prod": "my_prod_override_value",
            },
            opts=pulumi.ResourceOptions(depends_on=[
                    dbtcloud_project["dbt_project"],
                    dbtcloud_environment["dev_env"],
                    dbtcloud_environment["ci_env"],
                    dbtcloud_environment["prod_env"],
                ]))
        ```

        ## Import

        Import using a project ID and environment variable name found in the URL and UI or via the API.

        ```sh
         $ pulumi import dbtcloud:index/environmentVariable:EnvironmentVariable test_environment_variable "project_id:environment_variable_name"
        ```

        ```sh
         $ pulumi import dbtcloud:index/environmentVariable:EnvironmentVariable test_environment_variable 12345:DBT_ENV_VAR
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] environment_values: Map from environment names to respective variable value, a special key `project` should be set for the project default variable value
        :param pulumi.Input[str] name: Name for the variable, must be unique within a project, must be prefixed with 'DBT_'
        :param pulumi.Input[int] project_id: Project for the variable to be created in
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvironmentVariableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        *Note*: Some upstream resources can be slow to create, so if creating a project or environment at
        the same time as the environment variables, it's recommended to use the `depends_on` meta argument.

        ## Example Usage

        ```python
        import pulumi
        import aschot_dbtcloud as dbtcloud

        # NOTE for customers using the LEGACY dbt_cloud provider:
        # use dbt_cloud_environment_variable instead of dbtcloud_environment_variable for the legacy resource names
        # legacy names will be removed from 0.3 onwards
        dbt_my_env_var = dbtcloud.EnvironmentVariable("dbtMyEnvVar",
            project_id=dbtcloud_project["dbt_project"]["id"],
            environment_values={
                "project": "my_project_level_value",
                "Dev": "my_env_level_value",
                "CI": "my_ci_override_value",
                "Prod": "my_prod_override_value",
            },
            opts=pulumi.ResourceOptions(depends_on=[
                    dbtcloud_project["dbt_project"],
                    dbtcloud_environment["dev_env"],
                    dbtcloud_environment["ci_env"],
                    dbtcloud_environment["prod_env"],
                ]))
        ```

        ## Import

        Import using a project ID and environment variable name found in the URL and UI or via the API.

        ```sh
         $ pulumi import dbtcloud:index/environmentVariable:EnvironmentVariable test_environment_variable "project_id:environment_variable_name"
        ```

        ```sh
         $ pulumi import dbtcloud:index/environmentVariable:EnvironmentVariable test_environment_variable 12345:DBT_ENV_VAR
        ```

        :param str resource_name: The name of the resource.
        :param EnvironmentVariableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentVariableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environment_values: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvironmentVariableArgs.__new__(EnvironmentVariableArgs)

            if environment_values is None and not opts.urn:
                raise TypeError("Missing required property 'environment_values'")
            __props__.__dict__["environment_values"] = environment_values
            __props__.__dict__["name"] = name
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
        super(EnvironmentVariable, __self__).__init__(
            'dbtcloud:index/environmentVariable:EnvironmentVariable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            environment_values: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[int]] = None) -> 'EnvironmentVariable':
        """
        Get an existing EnvironmentVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] environment_values: Map from environment names to respective variable value, a special key `project` should be set for the project default variable value
        :param pulumi.Input[str] name: Name for the variable, must be unique within a project, must be prefixed with 'DBT_'
        :param pulumi.Input[int] project_id: Project for the variable to be created in
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnvironmentVariableState.__new__(_EnvironmentVariableState)

        __props__.__dict__["environment_values"] = environment_values
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        return EnvironmentVariable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="environmentValues")
    def environment_values(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        Map from environment names to respective variable value, a special key `project` should be set for the project default variable value
        """
        return pulumi.get(self, "environment_values")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name for the variable, must be unique within a project, must be prefixed with 'DBT_'
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[int]:
        """
        Project for the variable to be created in
        """
        return pulumi.get(self, "project_id")

