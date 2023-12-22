# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['EnvironmentVariableJobOverrideArgs', 'EnvironmentVariableJobOverride']

@pulumi.input_type
class EnvironmentVariableJobOverrideArgs:
    def __init__(__self__, *,
                 job_definition_id: pulumi.Input[int],
                 project_id: pulumi.Input[int],
                 raw_value: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EnvironmentVariableJobOverride resource.
        :param pulumi.Input[int] job_definition_id: The job ID for which the environment variable is being overridden
        :param pulumi.Input[int] project_id: The project ID for which the environment variable is being overridden
        :param pulumi.Input[str] raw_value: The value for the override of the environment variable
        :param pulumi.Input[str] name: The environment variable name to override
        """
        pulumi.set(__self__, "job_definition_id", job_definition_id)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "raw_value", raw_value)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="jobDefinitionId")
    def job_definition_id(self) -> pulumi.Input[int]:
        """
        The job ID for which the environment variable is being overridden
        """
        return pulumi.get(self, "job_definition_id")

    @job_definition_id.setter
    def job_definition_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "job_definition_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[int]:
        """
        The project ID for which the environment variable is being overridden
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="rawValue")
    def raw_value(self) -> pulumi.Input[str]:
        """
        The value for the override of the environment variable
        """
        return pulumi.get(self, "raw_value")

    @raw_value.setter
    def raw_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "raw_value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The environment variable name to override
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _EnvironmentVariableJobOverrideState:
    def __init__(__self__, *,
                 environment_variable_job_override_id: Optional[pulumi.Input[int]] = None,
                 job_definition_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 raw_value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EnvironmentVariableJobOverride resources.
        :param pulumi.Input[int] environment_variable_job_override_id: The ID of the environment variable job override
        :param pulumi.Input[int] job_definition_id: The job ID for which the environment variable is being overridden
        :param pulumi.Input[str] name: The environment variable name to override
        :param pulumi.Input[int] project_id: The project ID for which the environment variable is being overridden
        :param pulumi.Input[str] raw_value: The value for the override of the environment variable
        """
        if environment_variable_job_override_id is not None:
            pulumi.set(__self__, "environment_variable_job_override_id", environment_variable_job_override_id)
        if job_definition_id is not None:
            pulumi.set(__self__, "job_definition_id", job_definition_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if raw_value is not None:
            pulumi.set(__self__, "raw_value", raw_value)

    @property
    @pulumi.getter(name="environmentVariableJobOverrideId")
    def environment_variable_job_override_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the environment variable job override
        """
        return pulumi.get(self, "environment_variable_job_override_id")

    @environment_variable_job_override_id.setter
    def environment_variable_job_override_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "environment_variable_job_override_id", value)

    @property
    @pulumi.getter(name="jobDefinitionId")
    def job_definition_id(self) -> Optional[pulumi.Input[int]]:
        """
        The job ID for which the environment variable is being overridden
        """
        return pulumi.get(self, "job_definition_id")

    @job_definition_id.setter
    def job_definition_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "job_definition_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The environment variable name to override
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        The project ID for which the environment variable is being overridden
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="rawValue")
    def raw_value(self) -> Optional[pulumi.Input[str]]:
        """
        The value for the override of the environment variable
        """
        return pulumi.get(self, "raw_value")

    @raw_value.setter
    def raw_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "raw_value", value)


class EnvironmentVariableJobOverride(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_definition_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 raw_value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import aschot_dbtcloud as dbtcloud

        my_env_var_job_override = dbtcloud.EnvironmentVariableJobOverride("myEnvVarJobOverride",
            project_id=dbtcloud_project["dbt_project"]["id"],
            job_definition_id=dbtcloud_job["daily_job"]["id"],
            raw_value="my_override_value")
        ```

        ## Import

        Import using a project ID, a job ID and the environment variable override ID

        ```sh
         $ pulumi import dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride test_environment_variable_job_override "project_id:job_id:environment_variable_override_id"
        ```

        ```sh
         $ pulumi import dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride test_environment_variable_job_override 12345:678:123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] job_definition_id: The job ID for which the environment variable is being overridden
        :param pulumi.Input[str] name: The environment variable name to override
        :param pulumi.Input[int] project_id: The project ID for which the environment variable is being overridden
        :param pulumi.Input[str] raw_value: The value for the override of the environment variable
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvironmentVariableJobOverrideArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import aschot_dbtcloud as dbtcloud

        my_env_var_job_override = dbtcloud.EnvironmentVariableJobOverride("myEnvVarJobOverride",
            project_id=dbtcloud_project["dbt_project"]["id"],
            job_definition_id=dbtcloud_job["daily_job"]["id"],
            raw_value="my_override_value")
        ```

        ## Import

        Import using a project ID, a job ID and the environment variable override ID

        ```sh
         $ pulumi import dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride test_environment_variable_job_override "project_id:job_id:environment_variable_override_id"
        ```

        ```sh
         $ pulumi import dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride test_environment_variable_job_override 12345:678:123456
        ```

        :param str resource_name: The name of the resource.
        :param EnvironmentVariableJobOverrideArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentVariableJobOverrideArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_definition_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 raw_value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvironmentVariableJobOverrideArgs.__new__(EnvironmentVariableJobOverrideArgs)

            if job_definition_id is None and not opts.urn:
                raise TypeError("Missing required property 'job_definition_id'")
            __props__.__dict__["job_definition_id"] = job_definition_id
            __props__.__dict__["name"] = name
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if raw_value is None and not opts.urn:
                raise TypeError("Missing required property 'raw_value'")
            __props__.__dict__["raw_value"] = raw_value
            __props__.__dict__["environment_variable_job_override_id"] = None
        super(EnvironmentVariableJobOverride, __self__).__init__(
            'dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            environment_variable_job_override_id: Optional[pulumi.Input[int]] = None,
            job_definition_id: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            raw_value: Optional[pulumi.Input[str]] = None) -> 'EnvironmentVariableJobOverride':
        """
        Get an existing EnvironmentVariableJobOverride resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] environment_variable_job_override_id: The ID of the environment variable job override
        :param pulumi.Input[int] job_definition_id: The job ID for which the environment variable is being overridden
        :param pulumi.Input[str] name: The environment variable name to override
        :param pulumi.Input[int] project_id: The project ID for which the environment variable is being overridden
        :param pulumi.Input[str] raw_value: The value for the override of the environment variable
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnvironmentVariableJobOverrideState.__new__(_EnvironmentVariableJobOverrideState)

        __props__.__dict__["environment_variable_job_override_id"] = environment_variable_job_override_id
        __props__.__dict__["job_definition_id"] = job_definition_id
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["raw_value"] = raw_value
        return EnvironmentVariableJobOverride(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="environmentVariableJobOverrideId")
    def environment_variable_job_override_id(self) -> pulumi.Output[int]:
        """
        The ID of the environment variable job override
        """
        return pulumi.get(self, "environment_variable_job_override_id")

    @property
    @pulumi.getter(name="jobDefinitionId")
    def job_definition_id(self) -> pulumi.Output[int]:
        """
        The job ID for which the environment variable is being overridden
        """
        return pulumi.get(self, "job_definition_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The environment variable name to override
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[int]:
        """
        The project ID for which the environment variable is being overridden
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="rawValue")
    def raw_value(self) -> pulumi.Output[str]:
        """
        The value for the override of the environment variable
        """
        return pulumi.get(self, "raw_value")

