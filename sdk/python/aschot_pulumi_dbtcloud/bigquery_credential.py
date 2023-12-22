# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BigqueryCredentialArgs', 'BigqueryCredential']

@pulumi.input_type
class BigqueryCredentialArgs:
    def __init__(__self__, *,
                 dataset: pulumi.Input[str],
                 num_threads: pulumi.Input[int],
                 project_id: pulumi.Input[int],
                 is_active: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a BigqueryCredential resource.
        :param pulumi.Input[str] dataset: Default dataset name
        :param pulumi.Input[int] num_threads: Number of threads to use
        :param pulumi.Input[int] project_id: Project ID to create the BigQuery credential in
        :param pulumi.Input[bool] is_active: Whether the BigQuery credential is active
        """
        pulumi.set(__self__, "dataset", dataset)
        pulumi.set(__self__, "num_threads", num_threads)
        pulumi.set(__self__, "project_id", project_id)
        if is_active is not None:
            pulumi.set(__self__, "is_active", is_active)

    @property
    @pulumi.getter
    def dataset(self) -> pulumi.Input[str]:
        """
        Default dataset name
        """
        return pulumi.get(self, "dataset")

    @dataset.setter
    def dataset(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset", value)

    @property
    @pulumi.getter(name="numThreads")
    def num_threads(self) -> pulumi.Input[int]:
        """
        Number of threads to use
        """
        return pulumi.get(self, "num_threads")

    @num_threads.setter
    def num_threads(self, value: pulumi.Input[int]):
        pulumi.set(self, "num_threads", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[int]:
        """
        Project ID to create the BigQuery credential in
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the BigQuery credential is active
        """
        return pulumi.get(self, "is_active")

    @is_active.setter
    def is_active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_active", value)


@pulumi.input_type
class _BigqueryCredentialState:
    def __init__(__self__, *,
                 credential_id: Optional[pulumi.Input[int]] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 num_threads: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering BigqueryCredential resources.
        :param pulumi.Input[int] credential_id: The system BigQuery credential ID
        :param pulumi.Input[str] dataset: Default dataset name
        :param pulumi.Input[bool] is_active: Whether the BigQuery credential is active
        :param pulumi.Input[int] num_threads: Number of threads to use
        :param pulumi.Input[int] project_id: Project ID to create the BigQuery credential in
        """
        if credential_id is not None:
            pulumi.set(__self__, "credential_id", credential_id)
        if dataset is not None:
            pulumi.set(__self__, "dataset", dataset)
        if is_active is not None:
            pulumi.set(__self__, "is_active", is_active)
        if num_threads is not None:
            pulumi.set(__self__, "num_threads", num_threads)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="credentialId")
    def credential_id(self) -> Optional[pulumi.Input[int]]:
        """
        The system BigQuery credential ID
        """
        return pulumi.get(self, "credential_id")

    @credential_id.setter
    def credential_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "credential_id", value)

    @property
    @pulumi.getter
    def dataset(self) -> Optional[pulumi.Input[str]]:
        """
        Default dataset name
        """
        return pulumi.get(self, "dataset")

    @dataset.setter
    def dataset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dataset", value)

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the BigQuery credential is active
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
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Project ID to create the BigQuery credential in
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)


class BigqueryCredential(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 num_threads: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import aschot_pulumi_dbtcloud as dbtcloud

        # NOTE for customers using the LEGACY dbt_cloud provider:
        # use dbt_cloud_bigquery_credential instead of dbtcloud_bigquery_credential for the legacy resource names
        # legacy names will be removed from 0.3 onwards
        my_credential = dbtcloud.BigqueryCredential("myCredential",
            project_id=dbtcloud_project["dbt_project"]["id"],
            dataset="my_bq_dataset",
            num_threads=16)
        ```

        ## Import

        ```sh
         $ pulumi import dbtcloud:index/bigqueryCredential:BigqueryCredential my_credential "project_id:credential_id"
        ```

        ```sh
         $ pulumi import dbtcloud:index/bigqueryCredential:BigqueryCredential my_credential 12345:5678
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Default dataset name
        :param pulumi.Input[bool] is_active: Whether the BigQuery credential is active
        :param pulumi.Input[int] num_threads: Number of threads to use
        :param pulumi.Input[int] project_id: Project ID to create the BigQuery credential in
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BigqueryCredentialArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import aschot_pulumi_dbtcloud as dbtcloud

        # NOTE for customers using the LEGACY dbt_cloud provider:
        # use dbt_cloud_bigquery_credential instead of dbtcloud_bigquery_credential for the legacy resource names
        # legacy names will be removed from 0.3 onwards
        my_credential = dbtcloud.BigqueryCredential("myCredential",
            project_id=dbtcloud_project["dbt_project"]["id"],
            dataset="my_bq_dataset",
            num_threads=16)
        ```

        ## Import

        ```sh
         $ pulumi import dbtcloud:index/bigqueryCredential:BigqueryCredential my_credential "project_id:credential_id"
        ```

        ```sh
         $ pulumi import dbtcloud:index/bigqueryCredential:BigqueryCredential my_credential 12345:5678
        ```

        :param str resource_name: The name of the resource.
        :param BigqueryCredentialArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BigqueryCredentialArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 num_threads: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BigqueryCredentialArgs.__new__(BigqueryCredentialArgs)

            if dataset is None and not opts.urn:
                raise TypeError("Missing required property 'dataset'")
            __props__.__dict__["dataset"] = dataset
            __props__.__dict__["is_active"] = is_active
            if num_threads is None and not opts.urn:
                raise TypeError("Missing required property 'num_threads'")
            __props__.__dict__["num_threads"] = num_threads
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["credential_id"] = None
        super(BigqueryCredential, __self__).__init__(
            'dbtcloud:index/bigqueryCredential:BigqueryCredential',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            credential_id: Optional[pulumi.Input[int]] = None,
            dataset: Optional[pulumi.Input[str]] = None,
            is_active: Optional[pulumi.Input[bool]] = None,
            num_threads: Optional[pulumi.Input[int]] = None,
            project_id: Optional[pulumi.Input[int]] = None) -> 'BigqueryCredential':
        """
        Get an existing BigqueryCredential resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] credential_id: The system BigQuery credential ID
        :param pulumi.Input[str] dataset: Default dataset name
        :param pulumi.Input[bool] is_active: Whether the BigQuery credential is active
        :param pulumi.Input[int] num_threads: Number of threads to use
        :param pulumi.Input[int] project_id: Project ID to create the BigQuery credential in
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BigqueryCredentialState.__new__(_BigqueryCredentialState)

        __props__.__dict__["credential_id"] = credential_id
        __props__.__dict__["dataset"] = dataset
        __props__.__dict__["is_active"] = is_active
        __props__.__dict__["num_threads"] = num_threads
        __props__.__dict__["project_id"] = project_id
        return BigqueryCredential(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="credentialId")
    def credential_id(self) -> pulumi.Output[int]:
        """
        The system BigQuery credential ID
        """
        return pulumi.get(self, "credential_id")

    @property
    @pulumi.getter
    def dataset(self) -> pulumi.Output[str]:
        """
        Default dataset name
        """
        return pulumi.get(self, "dataset")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the BigQuery credential is active
        """
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter(name="numThreads")
    def num_threads(self) -> pulumi.Output[int]:
        """
        Number of threads to use
        """
        return pulumi.get(self, "num_threads")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[int]:
        """
        Project ID to create the BigQuery credential in
        """
        return pulumi.get(self, "project_id")
