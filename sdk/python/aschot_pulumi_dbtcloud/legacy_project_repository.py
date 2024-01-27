# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['LegacyProjectRepositoryArgs', 'LegacyProjectRepository']

@pulumi.input_type
class LegacyProjectRepositoryArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[int],
                 repository_id: pulumi.Input[int]):
        """
        The set of arguments for constructing a LegacyProjectRepository resource.
        :param pulumi.Input[int] project_id: Project ID
        :param pulumi.Input[int] repository_id: Repository ID
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "repository_id", repository_id)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[int]:
        """
        Project ID
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Input[int]:
        """
        Repository ID
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "repository_id", value)


@pulumi.input_type
class _LegacyProjectRepositoryState:
    def __init__(__self__, *,
                 project_id: Optional[pulumi.Input[int]] = None,
                 repository_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering LegacyProjectRepository resources.
        :param pulumi.Input[int] project_id: Project ID
        :param pulumi.Input[int] repository_id: Repository ID
        """
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if repository_id is not None:
            pulumi.set(__self__, "repository_id", repository_id)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Project ID
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[pulumi.Input[int]]:
        """
        Repository ID
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "repository_id", value)


warnings.warn("""Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource""", DeprecationWarning)


class LegacyProjectRepository(pulumi.CustomResource):
    warnings.warn("""Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 repository_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a LegacyProjectRepository resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] project_id: Project ID
        :param pulumi.Input[int] repository_id: Repository ID
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LegacyProjectRepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LegacyProjectRepository resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LegacyProjectRepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LegacyProjectRepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 repository_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        pulumi.log.warn("""LegacyProjectRepository is deprecated: Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LegacyProjectRepositoryArgs.__new__(LegacyProjectRepositoryArgs)

            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if repository_id is None and not opts.urn:
                raise TypeError("Missing required property 'repository_id'")
            __props__.__dict__["repository_id"] = repository_id
        super(LegacyProjectRepository, __self__).__init__(
            'dbtcloud:index/legacyProjectRepository:LegacyProjectRepository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            repository_id: Optional[pulumi.Input[int]] = None) -> 'LegacyProjectRepository':
        """
        Get an existing LegacyProjectRepository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] project_id: Project ID
        :param pulumi.Input[int] repository_id: Repository ID
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LegacyProjectRepositoryState.__new__(_LegacyProjectRepositoryState)

        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["repository_id"] = repository_id
        return LegacyProjectRepository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[int]:
        """
        Project ID
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Output[int]:
        """
        Repository ID
        """
        return pulumi.get(self, "repository_id")

