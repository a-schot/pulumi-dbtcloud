# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FabricCredentialArgs', 'FabricCredential']

@pulumi.input_type
class FabricCredentialArgs:
    def __init__(__self__, *,
                 adapter_id: pulumi.Input[int],
                 project_id: pulumi.Input[int],
                 schema: pulumi.Input[str],
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 schema_authorization: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FabricCredential resource.
        :param pulumi.Input[int] adapter_id: Fabric adapter ID for the credential
        :param pulumi.Input[int] project_id: Project ID to create the Fabric credential in
        :param pulumi.Input[str] schema: The schema where to create the dbt models
        :param pulumi.Input[str] client_id: The client ID of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        :param pulumi.Input[str] client_secret: The client secret of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        :param pulumi.Input[str] password: The password for the account to connect to. Only used when connection with AD user/pass
        :param pulumi.Input[str] schema_authorization: Optionally set this to the principal who should own the schemas created by dbt
        :param pulumi.Input[str] tenant_id: The tenant ID of the Azure Active Directory instance. This is only used when connecting to Azure SQL with a service principal.
        :param pulumi.Input[str] user: The username of the Fabric account to connect to. Only used when connection with AD user/pass
        """
        pulumi.set(__self__, "adapter_id", adapter_id)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "schema", schema)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if schema_authorization is not None:
            pulumi.set(__self__, "schema_authorization", schema_authorization)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter(name="adapterId")
    def adapter_id(self) -> pulumi.Input[int]:
        """
        Fabric adapter ID for the credential
        """
        return pulumi.get(self, "adapter_id")

    @adapter_id.setter
    def adapter_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "adapter_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[int]:
        """
        Project ID to create the Fabric credential in
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Input[str]:
        """
        The schema where to create the dbt models
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: pulumi.Input[str]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The client ID of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        """
        The client secret of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password for the account to connect to. Only used when connection with AD user/pass
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="schemaAuthorization")
    def schema_authorization(self) -> Optional[pulumi.Input[str]]:
        """
        Optionally set this to the principal who should own the schemas created by dbt
        """
        return pulumi.get(self, "schema_authorization")

    @schema_authorization.setter
    def schema_authorization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema_authorization", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The tenant ID of the Azure Active Directory instance. This is only used when connecting to Azure SQL with a service principal.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        The username of the Fabric account to connect to. Only used when connection with AD user/pass
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


@pulumi.input_type
class _FabricCredentialState:
    def __init__(__self__, *,
                 adapter_id: Optional[pulumi.Input[int]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 credential_id: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 schema_authorization: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FabricCredential resources.
        :param pulumi.Input[int] adapter_id: Fabric adapter ID for the credential
        :param pulumi.Input[str] client_id: The client ID of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        :param pulumi.Input[str] client_secret: The client secret of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        :param pulumi.Input[int] credential_id: The system Fabric credential ID
        :param pulumi.Input[str] password: The password for the account to connect to. Only used when connection with AD user/pass
        :param pulumi.Input[int] project_id: Project ID to create the Fabric credential in
        :param pulumi.Input[str] schema: The schema where to create the dbt models
        :param pulumi.Input[str] schema_authorization: Optionally set this to the principal who should own the schemas created by dbt
        :param pulumi.Input[str] tenant_id: The tenant ID of the Azure Active Directory instance. This is only used when connecting to Azure SQL with a service principal.
        :param pulumi.Input[str] user: The username of the Fabric account to connect to. Only used when connection with AD user/pass
        """
        if adapter_id is not None:
            pulumi.set(__self__, "adapter_id", adapter_id)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if credential_id is not None:
            pulumi.set(__self__, "credential_id", credential_id)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if schema_authorization is not None:
            pulumi.set(__self__, "schema_authorization", schema_authorization)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter(name="adapterId")
    def adapter_id(self) -> Optional[pulumi.Input[int]]:
        """
        Fabric adapter ID for the credential
        """
        return pulumi.get(self, "adapter_id")

    @adapter_id.setter
    def adapter_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "adapter_id", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The client ID of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        """
        The client secret of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="credentialId")
    def credential_id(self) -> Optional[pulumi.Input[int]]:
        """
        The system Fabric credential ID
        """
        return pulumi.get(self, "credential_id")

    @credential_id.setter
    def credential_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "credential_id", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password for the account to connect to. Only used when connection with AD user/pass
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Project ID to create the Fabric credential in
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[str]]:
        """
        The schema where to create the dbt models
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter(name="schemaAuthorization")
    def schema_authorization(self) -> Optional[pulumi.Input[str]]:
        """
        Optionally set this to the principal who should own the schemas created by dbt
        """
        return pulumi.get(self, "schema_authorization")

    @schema_authorization.setter
    def schema_authorization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema_authorization", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The tenant ID of the Azure Active Directory instance. This is only used when connecting to Azure SQL with a service principal.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        The username of the Fabric account to connect to. Only used when connection with AD user/pass
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


class FabricCredential(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adapter_id: Optional[pulumi.Input[int]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 schema_authorization: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Import using a project ID and credential ID found in the URL or via the API.

        ```sh
         $ pulumi import dbtcloud:index/fabricCredential:FabricCredential my_fabric_credential "project_id:credential_id"
        ```

        ```sh
         $ pulumi import dbtcloud:index/fabricCredential:FabricCredential my_fabric_credential 12345:6789
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] adapter_id: Fabric adapter ID for the credential
        :param pulumi.Input[str] client_id: The client ID of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        :param pulumi.Input[str] client_secret: The client secret of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        :param pulumi.Input[str] password: The password for the account to connect to. Only used when connection with AD user/pass
        :param pulumi.Input[int] project_id: Project ID to create the Fabric credential in
        :param pulumi.Input[str] schema: The schema where to create the dbt models
        :param pulumi.Input[str] schema_authorization: Optionally set this to the principal who should own the schemas created by dbt
        :param pulumi.Input[str] tenant_id: The tenant ID of the Azure Active Directory instance. This is only used when connecting to Azure SQL with a service principal.
        :param pulumi.Input[str] user: The username of the Fabric account to connect to. Only used when connection with AD user/pass
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FabricCredentialArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Import using a project ID and credential ID found in the URL or via the API.

        ```sh
         $ pulumi import dbtcloud:index/fabricCredential:FabricCredential my_fabric_credential "project_id:credential_id"
        ```

        ```sh
         $ pulumi import dbtcloud:index/fabricCredential:FabricCredential my_fabric_credential 12345:6789
        ```

        :param str resource_name: The name of the resource.
        :param FabricCredentialArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FabricCredentialArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adapter_id: Optional[pulumi.Input[int]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 schema_authorization: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FabricCredentialArgs.__new__(FabricCredentialArgs)

            if adapter_id is None and not opts.urn:
                raise TypeError("Missing required property 'adapter_id'")
            __props__.__dict__["adapter_id"] = adapter_id
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["client_secret"] = None if client_secret is None else pulumi.Output.secret(client_secret)
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if schema is None and not opts.urn:
                raise TypeError("Missing required property 'schema'")
            __props__.__dict__["schema"] = schema
            __props__.__dict__["schema_authorization"] = schema_authorization
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["user"] = user
            __props__.__dict__["credential_id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["clientSecret", "password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(FabricCredential, __self__).__init__(
            'dbtcloud:index/fabricCredential:FabricCredential',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            adapter_id: Optional[pulumi.Input[int]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            credential_id: Optional[pulumi.Input[int]] = None,
            password: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            schema: Optional[pulumi.Input[str]] = None,
            schema_authorization: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'FabricCredential':
        """
        Get an existing FabricCredential resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] adapter_id: Fabric adapter ID for the credential
        :param pulumi.Input[str] client_id: The client ID of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        :param pulumi.Input[str] client_secret: The client secret of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        :param pulumi.Input[int] credential_id: The system Fabric credential ID
        :param pulumi.Input[str] password: The password for the account to connect to. Only used when connection with AD user/pass
        :param pulumi.Input[int] project_id: Project ID to create the Fabric credential in
        :param pulumi.Input[str] schema: The schema where to create the dbt models
        :param pulumi.Input[str] schema_authorization: Optionally set this to the principal who should own the schemas created by dbt
        :param pulumi.Input[str] tenant_id: The tenant ID of the Azure Active Directory instance. This is only used when connecting to Azure SQL with a service principal.
        :param pulumi.Input[str] user: The username of the Fabric account to connect to. Only used when connection with AD user/pass
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FabricCredentialState.__new__(_FabricCredentialState)

        __props__.__dict__["adapter_id"] = adapter_id
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_secret"] = client_secret
        __props__.__dict__["credential_id"] = credential_id
        __props__.__dict__["password"] = password
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["schema"] = schema
        __props__.__dict__["schema_authorization"] = schema_authorization
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["user"] = user
        return FabricCredential(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adapterId")
    def adapter_id(self) -> pulumi.Output[int]:
        """
        Fabric adapter ID for the credential
        """
        return pulumi.get(self, "adapter_id")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        The client ID of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[Optional[str]]:
        """
        The client secret of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="credentialId")
    def credential_id(self) -> pulumi.Output[int]:
        """
        The system Fabric credential ID
        """
        return pulumi.get(self, "credential_id")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password for the account to connect to. Only used when connection with AD user/pass
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[int]:
        """
        Project ID to create the Fabric credential in
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[str]:
        """
        The schema where to create the dbt models
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter(name="schemaAuthorization")
    def schema_authorization(self) -> pulumi.Output[Optional[str]]:
        """
        Optionally set this to the principal who should own the schemas created by dbt
        """
        return pulumi.get(self, "schema_authorization")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        The tenant ID of the Azure Active Directory instance. This is only used when connecting to Azure SQL with a service principal.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[Optional[str]]:
        """
        The username of the Fabric account to connect to. Only used when connection with AD user/pass
        """
        return pulumi.get(self, "user")
