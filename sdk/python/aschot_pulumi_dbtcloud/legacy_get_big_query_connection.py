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
    'LegacyGetBigQueryConnectionResult',
    'AwaitableLegacyGetBigQueryConnectionResult',
    'legacy_get_big_query_connection',
    'legacy_get_big_query_connection_output',
]

@pulumi.output_type
class LegacyGetBigQueryConnectionResult:
    """
    A collection of values returned by LegacyGetBigQueryConnection.
    """
    def __init__(__self__, auth_provider_x509_cert_url=None, auth_uri=None, client_email=None, client_id=None, client_x509_cert_url=None, connection_id=None, dataproc_cluster_name=None, dataproc_region=None, execution_project=None, gcp_project_id=None, gcs_bucket=None, id=None, is_active=None, is_configured_for_oauth=None, location=None, maximum_bytes_billed=None, name=None, priority=None, private_key=None, private_key_id=None, project_id=None, retries=None, timeout_seconds=None, token_uri=None, type=None):
        if auth_provider_x509_cert_url and not isinstance(auth_provider_x509_cert_url, str):
            raise TypeError("Expected argument 'auth_provider_x509_cert_url' to be a str")
        pulumi.set(__self__, "auth_provider_x509_cert_url", auth_provider_x509_cert_url)
        if auth_uri and not isinstance(auth_uri, str):
            raise TypeError("Expected argument 'auth_uri' to be a str")
        pulumi.set(__self__, "auth_uri", auth_uri)
        if client_email and not isinstance(client_email, str):
            raise TypeError("Expected argument 'client_email' to be a str")
        pulumi.set(__self__, "client_email", client_email)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if client_x509_cert_url and not isinstance(client_x509_cert_url, str):
            raise TypeError("Expected argument 'client_x509_cert_url' to be a str")
        pulumi.set(__self__, "client_x509_cert_url", client_x509_cert_url)
        if connection_id and not isinstance(connection_id, int):
            raise TypeError("Expected argument 'connection_id' to be a int")
        pulumi.set(__self__, "connection_id", connection_id)
        if dataproc_cluster_name and not isinstance(dataproc_cluster_name, str):
            raise TypeError("Expected argument 'dataproc_cluster_name' to be a str")
        pulumi.set(__self__, "dataproc_cluster_name", dataproc_cluster_name)
        if dataproc_region and not isinstance(dataproc_region, str):
            raise TypeError("Expected argument 'dataproc_region' to be a str")
        pulumi.set(__self__, "dataproc_region", dataproc_region)
        if execution_project and not isinstance(execution_project, str):
            raise TypeError("Expected argument 'execution_project' to be a str")
        pulumi.set(__self__, "execution_project", execution_project)
        if gcp_project_id and not isinstance(gcp_project_id, str):
            raise TypeError("Expected argument 'gcp_project_id' to be a str")
        pulumi.set(__self__, "gcp_project_id", gcp_project_id)
        if gcs_bucket and not isinstance(gcs_bucket, str):
            raise TypeError("Expected argument 'gcs_bucket' to be a str")
        pulumi.set(__self__, "gcs_bucket", gcs_bucket)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_active and not isinstance(is_active, bool):
            raise TypeError("Expected argument 'is_active' to be a bool")
        pulumi.set(__self__, "is_active", is_active)
        if is_configured_for_oauth and not isinstance(is_configured_for_oauth, bool):
            raise TypeError("Expected argument 'is_configured_for_oauth' to be a bool")
        pulumi.set(__self__, "is_configured_for_oauth", is_configured_for_oauth)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if maximum_bytes_billed and not isinstance(maximum_bytes_billed, int):
            raise TypeError("Expected argument 'maximum_bytes_billed' to be a int")
        pulumi.set(__self__, "maximum_bytes_billed", maximum_bytes_billed)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if priority and not isinstance(priority, str):
            raise TypeError("Expected argument 'priority' to be a str")
        pulumi.set(__self__, "priority", priority)
        if private_key and not isinstance(private_key, str):
            raise TypeError("Expected argument 'private_key' to be a str")
        pulumi.set(__self__, "private_key", private_key)
        if private_key_id and not isinstance(private_key_id, str):
            raise TypeError("Expected argument 'private_key_id' to be a str")
        pulumi.set(__self__, "private_key_id", private_key_id)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if retries and not isinstance(retries, int):
            raise TypeError("Expected argument 'retries' to be a int")
        pulumi.set(__self__, "retries", retries)
        if timeout_seconds and not isinstance(timeout_seconds, int):
            raise TypeError("Expected argument 'timeout_seconds' to be a int")
        pulumi.set(__self__, "timeout_seconds", timeout_seconds)
        if token_uri and not isinstance(token_uri, str):
            raise TypeError("Expected argument 'token_uri' to be a str")
        pulumi.set(__self__, "token_uri", token_uri)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="authProviderX509CertUrl")
    def auth_provider_x509_cert_url(self) -> str:
        return pulumi.get(self, "auth_provider_x509_cert_url")

    @property
    @pulumi.getter(name="authUri")
    def auth_uri(self) -> str:
        return pulumi.get(self, "auth_uri")

    @property
    @pulumi.getter(name="clientEmail")
    def client_email(self) -> str:
        return pulumi.get(self, "client_email")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientX509CertUrl")
    def client_x509_cert_url(self) -> str:
        return pulumi.get(self, "client_x509_cert_url")

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> int:
        return pulumi.get(self, "connection_id")

    @property
    @pulumi.getter(name="dataprocClusterName")
    def dataproc_cluster_name(self) -> str:
        return pulumi.get(self, "dataproc_cluster_name")

    @property
    @pulumi.getter(name="dataprocRegion")
    def dataproc_region(self) -> str:
        return pulumi.get(self, "dataproc_region")

    @property
    @pulumi.getter(name="executionProject")
    def execution_project(self) -> str:
        return pulumi.get(self, "execution_project")

    @property
    @pulumi.getter(name="gcpProjectId")
    def gcp_project_id(self) -> str:
        return pulumi.get(self, "gcp_project_id")

    @property
    @pulumi.getter(name="gcsBucket")
    def gcs_bucket(self) -> str:
        return pulumi.get(self, "gcs_bucket")

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
    @pulumi.getter(name="isConfiguredForOauth")
    def is_configured_for_oauth(self) -> bool:
        return pulumi.get(self, "is_configured_for_oauth")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maximumBytesBilled")
    def maximum_bytes_billed(self) -> int:
        return pulumi.get(self, "maximum_bytes_billed")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> str:
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> str:
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="privateKeyId")
    def private_key_id(self) -> str:
        return pulumi.get(self, "private_key_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def retries(self) -> int:
        return pulumi.get(self, "retries")

    @property
    @pulumi.getter(name="timeoutSeconds")
    def timeout_seconds(self) -> int:
        return pulumi.get(self, "timeout_seconds")

    @property
    @pulumi.getter(name="tokenUri")
    def token_uri(self) -> str:
        return pulumi.get(self, "token_uri")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


class AwaitableLegacyGetBigQueryConnectionResult(LegacyGetBigQueryConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return LegacyGetBigQueryConnectionResult(
            auth_provider_x509_cert_url=self.auth_provider_x509_cert_url,
            auth_uri=self.auth_uri,
            client_email=self.client_email,
            client_id=self.client_id,
            client_x509_cert_url=self.client_x509_cert_url,
            connection_id=self.connection_id,
            dataproc_cluster_name=self.dataproc_cluster_name,
            dataproc_region=self.dataproc_region,
            execution_project=self.execution_project,
            gcp_project_id=self.gcp_project_id,
            gcs_bucket=self.gcs_bucket,
            id=self.id,
            is_active=self.is_active,
            is_configured_for_oauth=self.is_configured_for_oauth,
            location=self.location,
            maximum_bytes_billed=self.maximum_bytes_billed,
            name=self.name,
            priority=self.priority,
            private_key=self.private_key,
            private_key_id=self.private_key_id,
            project_id=self.project_id,
            retries=self.retries,
            timeout_seconds=self.timeout_seconds,
            token_uri=self.token_uri,
            type=self.type)


def legacy_get_big_query_connection(connection_id: Optional[int] = None,
                                    project_id: Optional[int] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableLegacyGetBigQueryConnectionResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['connectionId'] = connection_id
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('dbtcloud:index/legacyGetBigQueryConnection:LegacyGetBigQueryConnection', __args__, opts=opts, typ=LegacyGetBigQueryConnectionResult).value

    return AwaitableLegacyGetBigQueryConnectionResult(
        auth_provider_x509_cert_url=pulumi.get(__ret__, 'auth_provider_x509_cert_url'),
        auth_uri=pulumi.get(__ret__, 'auth_uri'),
        client_email=pulumi.get(__ret__, 'client_email'),
        client_id=pulumi.get(__ret__, 'client_id'),
        client_x509_cert_url=pulumi.get(__ret__, 'client_x509_cert_url'),
        connection_id=pulumi.get(__ret__, 'connection_id'),
        dataproc_cluster_name=pulumi.get(__ret__, 'dataproc_cluster_name'),
        dataproc_region=pulumi.get(__ret__, 'dataproc_region'),
        execution_project=pulumi.get(__ret__, 'execution_project'),
        gcp_project_id=pulumi.get(__ret__, 'gcp_project_id'),
        gcs_bucket=pulumi.get(__ret__, 'gcs_bucket'),
        id=pulumi.get(__ret__, 'id'),
        is_active=pulumi.get(__ret__, 'is_active'),
        is_configured_for_oauth=pulumi.get(__ret__, 'is_configured_for_oauth'),
        location=pulumi.get(__ret__, 'location'),
        maximum_bytes_billed=pulumi.get(__ret__, 'maximum_bytes_billed'),
        name=pulumi.get(__ret__, 'name'),
        priority=pulumi.get(__ret__, 'priority'),
        private_key=pulumi.get(__ret__, 'private_key'),
        private_key_id=pulumi.get(__ret__, 'private_key_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        retries=pulumi.get(__ret__, 'retries'),
        timeout_seconds=pulumi.get(__ret__, 'timeout_seconds'),
        token_uri=pulumi.get(__ret__, 'token_uri'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(legacy_get_big_query_connection)
def legacy_get_big_query_connection_output(connection_id: Optional[pulumi.Input[int]] = None,
                                           project_id: Optional[pulumi.Input[int]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[LegacyGetBigQueryConnectionResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
