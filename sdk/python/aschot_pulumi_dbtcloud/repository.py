# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RepositoryArgs', 'Repository']

@pulumi.input_type
class RepositoryArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[int],
                 remote_url: pulumi.Input[str],
                 azure_active_directory_project_id: Optional[pulumi.Input[str]] = None,
                 azure_active_directory_repository_id: Optional[pulumi.Input[str]] = None,
                 azure_bypass_webhook_registration_failure: Optional[pulumi.Input[bool]] = None,
                 fetch_deploy_key: Optional[pulumi.Input[bool]] = None,
                 git_clone_strategy: Optional[pulumi.Input[str]] = None,
                 github_installation_id: Optional[pulumi.Input[int]] = None,
                 gitlab_project_id: Optional[pulumi.Input[int]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Repository resource.
        :param pulumi.Input[int] project_id: Project ID to create the repository in
        :param pulumi.Input[str] remote_url: Git URL for the repository or <Group>/<Project> for Gitlab
        :param pulumi.Input[str] azure_active_directory_project_id: The Azure Dev Ops project ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_project` and the project name - (for ADO native integration only)
        :param pulumi.Input[str] azure_active_directory_repository_id: The Azure Dev Ops repository ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_repository` along with the ADO Project ID and the repository name - (for ADO native integration only)
        :param pulumi.Input[bool] azure_bypass_webhook_registration_failure: If set to False (the default), the connection will fail if the service user doesn't have access to set webhooks (required for auto-triggering CI jobs). If set to True, the connection will be successful but no automated CI job will be triggered - (for ADO native integration only)
        :param pulumi.Input[bool] fetch_deploy_key: Whether we should return the public deploy key - (for the `deploy_key` strategy)
        :param pulumi.Input[str] git_clone_strategy: Git clone strategy for the repository. Can be `deploy_key` (default) for cloning via SSH Deploy Key, `github_app` for GitHub native integration, `deploy_token` for the GitLab native integration and `azure_active_directory_app` for ADO native integration
        :param pulumi.Input[int] github_installation_id: Identifier for the GitHub App - (for GitHub native integration only)
        :param pulumi.Input[int] gitlab_project_id: Identifier for the Gitlab project -  (for GitLab native integration only)
        :param pulumi.Input[bool] is_active: Whether the repository is active
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "remote_url", remote_url)
        if azure_active_directory_project_id is not None:
            pulumi.set(__self__, "azure_active_directory_project_id", azure_active_directory_project_id)
        if azure_active_directory_repository_id is not None:
            pulumi.set(__self__, "azure_active_directory_repository_id", azure_active_directory_repository_id)
        if azure_bypass_webhook_registration_failure is not None:
            pulumi.set(__self__, "azure_bypass_webhook_registration_failure", azure_bypass_webhook_registration_failure)
        if fetch_deploy_key is not None:
            pulumi.set(__self__, "fetch_deploy_key", fetch_deploy_key)
        if git_clone_strategy is not None:
            pulumi.set(__self__, "git_clone_strategy", git_clone_strategy)
        if github_installation_id is not None:
            pulumi.set(__self__, "github_installation_id", github_installation_id)
        if gitlab_project_id is not None:
            pulumi.set(__self__, "gitlab_project_id", gitlab_project_id)
        if is_active is not None:
            pulumi.set(__self__, "is_active", is_active)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[int]:
        """
        Project ID to create the repository in
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="remoteUrl")
    def remote_url(self) -> pulumi.Input[str]:
        """
        Git URL for the repository or <Group>/<Project> for Gitlab
        """
        return pulumi.get(self, "remote_url")

    @remote_url.setter
    def remote_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "remote_url", value)

    @property
    @pulumi.getter(name="azureActiveDirectoryProjectId")
    def azure_active_directory_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure Dev Ops project ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_project` and the project name - (for ADO native integration only)
        """
        return pulumi.get(self, "azure_active_directory_project_id")

    @azure_active_directory_project_id.setter
    def azure_active_directory_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_active_directory_project_id", value)

    @property
    @pulumi.getter(name="azureActiveDirectoryRepositoryId")
    def azure_active_directory_repository_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure Dev Ops repository ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_repository` along with the ADO Project ID and the repository name - (for ADO native integration only)
        """
        return pulumi.get(self, "azure_active_directory_repository_id")

    @azure_active_directory_repository_id.setter
    def azure_active_directory_repository_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_active_directory_repository_id", value)

    @property
    @pulumi.getter(name="azureBypassWebhookRegistrationFailure")
    def azure_bypass_webhook_registration_failure(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to False (the default), the connection will fail if the service user doesn't have access to set webhooks (required for auto-triggering CI jobs). If set to True, the connection will be successful but no automated CI job will be triggered - (for ADO native integration only)
        """
        return pulumi.get(self, "azure_bypass_webhook_registration_failure")

    @azure_bypass_webhook_registration_failure.setter
    def azure_bypass_webhook_registration_failure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "azure_bypass_webhook_registration_failure", value)

    @property
    @pulumi.getter(name="fetchDeployKey")
    def fetch_deploy_key(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether we should return the public deploy key - (for the `deploy_key` strategy)
        """
        return pulumi.get(self, "fetch_deploy_key")

    @fetch_deploy_key.setter
    def fetch_deploy_key(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "fetch_deploy_key", value)

    @property
    @pulumi.getter(name="gitCloneStrategy")
    def git_clone_strategy(self) -> Optional[pulumi.Input[str]]:
        """
        Git clone strategy for the repository. Can be `deploy_key` (default) for cloning via SSH Deploy Key, `github_app` for GitHub native integration, `deploy_token` for the GitLab native integration and `azure_active_directory_app` for ADO native integration
        """
        return pulumi.get(self, "git_clone_strategy")

    @git_clone_strategy.setter
    def git_clone_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "git_clone_strategy", value)

    @property
    @pulumi.getter(name="githubInstallationId")
    def github_installation_id(self) -> Optional[pulumi.Input[int]]:
        """
        Identifier for the GitHub App - (for GitHub native integration only)
        """
        return pulumi.get(self, "github_installation_id")

    @github_installation_id.setter
    def github_installation_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "github_installation_id", value)

    @property
    @pulumi.getter(name="gitlabProjectId")
    def gitlab_project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Identifier for the Gitlab project -  (for GitLab native integration only)
        """
        return pulumi.get(self, "gitlab_project_id")

    @gitlab_project_id.setter
    def gitlab_project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "gitlab_project_id", value)

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the repository is active
        """
        return pulumi.get(self, "is_active")

    @is_active.setter
    def is_active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_active", value)


@pulumi.input_type
class _RepositoryState:
    def __init__(__self__, *,
                 azure_active_directory_project_id: Optional[pulumi.Input[str]] = None,
                 azure_active_directory_repository_id: Optional[pulumi.Input[str]] = None,
                 azure_bypass_webhook_registration_failure: Optional[pulumi.Input[bool]] = None,
                 deploy_key: Optional[pulumi.Input[str]] = None,
                 fetch_deploy_key: Optional[pulumi.Input[bool]] = None,
                 git_clone_strategy: Optional[pulumi.Input[str]] = None,
                 github_installation_id: Optional[pulumi.Input[int]] = None,
                 gitlab_project_id: Optional[pulumi.Input[int]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 remote_url: Optional[pulumi.Input[str]] = None,
                 repository_credentials_id: Optional[pulumi.Input[int]] = None,
                 repository_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Repository resources.
        :param pulumi.Input[str] azure_active_directory_project_id: The Azure Dev Ops project ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_project` and the project name - (for ADO native integration only)
        :param pulumi.Input[str] azure_active_directory_repository_id: The Azure Dev Ops repository ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_repository` along with the ADO Project ID and the repository name - (for ADO native integration only)
        :param pulumi.Input[bool] azure_bypass_webhook_registration_failure: If set to False (the default), the connection will fail if the service user doesn't have access to set webhooks (required for auto-triggering CI jobs). If set to True, the connection will be successful but no automated CI job will be triggered - (for ADO native integration only)
        :param pulumi.Input[str] deploy_key: Public key generated by dbt when using `deploy_key` clone strategy
        :param pulumi.Input[bool] fetch_deploy_key: Whether we should return the public deploy key - (for the `deploy_key` strategy)
        :param pulumi.Input[str] git_clone_strategy: Git clone strategy for the repository. Can be `deploy_key` (default) for cloning via SSH Deploy Key, `github_app` for GitHub native integration, `deploy_token` for the GitLab native integration and `azure_active_directory_app` for ADO native integration
        :param pulumi.Input[int] github_installation_id: Identifier for the GitHub App - (for GitHub native integration only)
        :param pulumi.Input[int] gitlab_project_id: Identifier for the Gitlab project -  (for GitLab native integration only)
        :param pulumi.Input[bool] is_active: Whether the repository is active
        :param pulumi.Input[int] project_id: Project ID to create the repository in
        :param pulumi.Input[str] remote_url: Git URL for the repository or <Group>/<Project> for Gitlab
        :param pulumi.Input[int] repository_credentials_id: Credentials ID for the repository (From the repository side not the dbt Cloud ID)
        :param pulumi.Input[int] repository_id: Repository Identifier
        """
        if azure_active_directory_project_id is not None:
            pulumi.set(__self__, "azure_active_directory_project_id", azure_active_directory_project_id)
        if azure_active_directory_repository_id is not None:
            pulumi.set(__self__, "azure_active_directory_repository_id", azure_active_directory_repository_id)
        if azure_bypass_webhook_registration_failure is not None:
            pulumi.set(__self__, "azure_bypass_webhook_registration_failure", azure_bypass_webhook_registration_failure)
        if deploy_key is not None:
            pulumi.set(__self__, "deploy_key", deploy_key)
        if fetch_deploy_key is not None:
            pulumi.set(__self__, "fetch_deploy_key", fetch_deploy_key)
        if git_clone_strategy is not None:
            pulumi.set(__self__, "git_clone_strategy", git_clone_strategy)
        if github_installation_id is not None:
            pulumi.set(__self__, "github_installation_id", github_installation_id)
        if gitlab_project_id is not None:
            pulumi.set(__self__, "gitlab_project_id", gitlab_project_id)
        if is_active is not None:
            pulumi.set(__self__, "is_active", is_active)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if remote_url is not None:
            pulumi.set(__self__, "remote_url", remote_url)
        if repository_credentials_id is not None:
            pulumi.set(__self__, "repository_credentials_id", repository_credentials_id)
        if repository_id is not None:
            pulumi.set(__self__, "repository_id", repository_id)

    @property
    @pulumi.getter(name="azureActiveDirectoryProjectId")
    def azure_active_directory_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure Dev Ops project ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_project` and the project name - (for ADO native integration only)
        """
        return pulumi.get(self, "azure_active_directory_project_id")

    @azure_active_directory_project_id.setter
    def azure_active_directory_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_active_directory_project_id", value)

    @property
    @pulumi.getter(name="azureActiveDirectoryRepositoryId")
    def azure_active_directory_repository_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure Dev Ops repository ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_repository` along with the ADO Project ID and the repository name - (for ADO native integration only)
        """
        return pulumi.get(self, "azure_active_directory_repository_id")

    @azure_active_directory_repository_id.setter
    def azure_active_directory_repository_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_active_directory_repository_id", value)

    @property
    @pulumi.getter(name="azureBypassWebhookRegistrationFailure")
    def azure_bypass_webhook_registration_failure(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to False (the default), the connection will fail if the service user doesn't have access to set webhooks (required for auto-triggering CI jobs). If set to True, the connection will be successful but no automated CI job will be triggered - (for ADO native integration only)
        """
        return pulumi.get(self, "azure_bypass_webhook_registration_failure")

    @azure_bypass_webhook_registration_failure.setter
    def azure_bypass_webhook_registration_failure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "azure_bypass_webhook_registration_failure", value)

    @property
    @pulumi.getter(name="deployKey")
    def deploy_key(self) -> Optional[pulumi.Input[str]]:
        """
        Public key generated by dbt when using `deploy_key` clone strategy
        """
        return pulumi.get(self, "deploy_key")

    @deploy_key.setter
    def deploy_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deploy_key", value)

    @property
    @pulumi.getter(name="fetchDeployKey")
    def fetch_deploy_key(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether we should return the public deploy key - (for the `deploy_key` strategy)
        """
        return pulumi.get(self, "fetch_deploy_key")

    @fetch_deploy_key.setter
    def fetch_deploy_key(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "fetch_deploy_key", value)

    @property
    @pulumi.getter(name="gitCloneStrategy")
    def git_clone_strategy(self) -> Optional[pulumi.Input[str]]:
        """
        Git clone strategy for the repository. Can be `deploy_key` (default) for cloning via SSH Deploy Key, `github_app` for GitHub native integration, `deploy_token` for the GitLab native integration and `azure_active_directory_app` for ADO native integration
        """
        return pulumi.get(self, "git_clone_strategy")

    @git_clone_strategy.setter
    def git_clone_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "git_clone_strategy", value)

    @property
    @pulumi.getter(name="githubInstallationId")
    def github_installation_id(self) -> Optional[pulumi.Input[int]]:
        """
        Identifier for the GitHub App - (for GitHub native integration only)
        """
        return pulumi.get(self, "github_installation_id")

    @github_installation_id.setter
    def github_installation_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "github_installation_id", value)

    @property
    @pulumi.getter(name="gitlabProjectId")
    def gitlab_project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Identifier for the Gitlab project -  (for GitLab native integration only)
        """
        return pulumi.get(self, "gitlab_project_id")

    @gitlab_project_id.setter
    def gitlab_project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "gitlab_project_id", value)

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the repository is active
        """
        return pulumi.get(self, "is_active")

    @is_active.setter
    def is_active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_active", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Project ID to create the repository in
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="remoteUrl")
    def remote_url(self) -> Optional[pulumi.Input[str]]:
        """
        Git URL for the repository or <Group>/<Project> for Gitlab
        """
        return pulumi.get(self, "remote_url")

    @remote_url.setter
    def remote_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_url", value)

    @property
    @pulumi.getter(name="repositoryCredentialsId")
    def repository_credentials_id(self) -> Optional[pulumi.Input[int]]:
        """
        Credentials ID for the repository (From the repository side not the dbt Cloud ID)
        """
        return pulumi.get(self, "repository_credentials_id")

    @repository_credentials_id.setter
    def repository_credentials_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "repository_credentials_id", value)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[pulumi.Input[int]]:
        """
        Repository Identifier
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "repository_id", value)


class Repository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_active_directory_project_id: Optional[pulumi.Input[str]] = None,
                 azure_active_directory_repository_id: Optional[pulumi.Input[str]] = None,
                 azure_bypass_webhook_registration_failure: Optional[pulumi.Input[bool]] = None,
                 fetch_deploy_key: Optional[pulumi.Input[bool]] = None,
                 git_clone_strategy: Optional[pulumi.Input[str]] = None,
                 github_installation_id: Optional[pulumi.Input[int]] = None,
                 gitlab_project_id: Optional[pulumi.Input[int]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 remote_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        *Note*: Some upstream resources can be slow to create, so if creating a project at
        the same time as the repository, it's recommended to use the `depends_on` meta argument.

        In order to find the `github_installation_id`, you can log in to dbt Cloud, replace `<dbt_cloud_url>` by your dbt Cloud
        URL and run the following commands in the Google Chrome console:

        Alternatively, you can go to the page `https://<dbt_cloud_url>/api/v2/integrations/github/installations/` and read the
        value of `id`  or use the `http` provider to retrieve it automatically like in the example below.

        ## Example Usage

        ## Import

        Import using a project ID and repository ID found in the URL or via the API. <break><break>```sh<break> $ pulumi import dbtcloud:index/repository:Repository test_repository "project_id:repository_id" <break>```<break><break> <break><break>```sh<break> $ pulumi import dbtcloud:index/repository:Repository test_repository 12345:6789 <break>```<break><break>

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] azure_active_directory_project_id: The Azure Dev Ops project ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_project` and the project name - (for ADO native integration only)
        :param pulumi.Input[str] azure_active_directory_repository_id: The Azure Dev Ops repository ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_repository` along with the ADO Project ID and the repository name - (for ADO native integration only)
        :param pulumi.Input[bool] azure_bypass_webhook_registration_failure: If set to False (the default), the connection will fail if the service user doesn't have access to set webhooks (required for auto-triggering CI jobs). If set to True, the connection will be successful but no automated CI job will be triggered - (for ADO native integration only)
        :param pulumi.Input[bool] fetch_deploy_key: Whether we should return the public deploy key - (for the `deploy_key` strategy)
        :param pulumi.Input[str] git_clone_strategy: Git clone strategy for the repository. Can be `deploy_key` (default) for cloning via SSH Deploy Key, `github_app` for GitHub native integration, `deploy_token` for the GitLab native integration and `azure_active_directory_app` for ADO native integration
        :param pulumi.Input[int] github_installation_id: Identifier for the GitHub App - (for GitHub native integration only)
        :param pulumi.Input[int] gitlab_project_id: Identifier for the Gitlab project -  (for GitLab native integration only)
        :param pulumi.Input[bool] is_active: Whether the repository is active
        :param pulumi.Input[int] project_id: Project ID to create the repository in
        :param pulumi.Input[str] remote_url: Git URL for the repository or <Group>/<Project> for Gitlab
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        *Note*: Some upstream resources can be slow to create, so if creating a project at
        the same time as the repository, it's recommended to use the `depends_on` meta argument.

        In order to find the `github_installation_id`, you can log in to dbt Cloud, replace `<dbt_cloud_url>` by your dbt Cloud
        URL and run the following commands in the Google Chrome console:

        Alternatively, you can go to the page `https://<dbt_cloud_url>/api/v2/integrations/github/installations/` and read the
        value of `id`  or use the `http` provider to retrieve it automatically like in the example below.

        ## Example Usage

        ## Import

        Import using a project ID and repository ID found in the URL or via the API. <break><break>```sh<break> $ pulumi import dbtcloud:index/repository:Repository test_repository "project_id:repository_id" <break>```<break><break> <break><break>```sh<break> $ pulumi import dbtcloud:index/repository:Repository test_repository 12345:6789 <break>```<break><break>

        :param str resource_name: The name of the resource.
        :param RepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_active_directory_project_id: Optional[pulumi.Input[str]] = None,
                 azure_active_directory_repository_id: Optional[pulumi.Input[str]] = None,
                 azure_bypass_webhook_registration_failure: Optional[pulumi.Input[bool]] = None,
                 fetch_deploy_key: Optional[pulumi.Input[bool]] = None,
                 git_clone_strategy: Optional[pulumi.Input[str]] = None,
                 github_installation_id: Optional[pulumi.Input[int]] = None,
                 gitlab_project_id: Optional[pulumi.Input[int]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 remote_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryArgs.__new__(RepositoryArgs)

            __props__.__dict__["azure_active_directory_project_id"] = azure_active_directory_project_id
            __props__.__dict__["azure_active_directory_repository_id"] = azure_active_directory_repository_id
            __props__.__dict__["azure_bypass_webhook_registration_failure"] = azure_bypass_webhook_registration_failure
            __props__.__dict__["fetch_deploy_key"] = fetch_deploy_key
            __props__.__dict__["git_clone_strategy"] = git_clone_strategy
            __props__.__dict__["github_installation_id"] = github_installation_id
            __props__.__dict__["gitlab_project_id"] = gitlab_project_id
            __props__.__dict__["is_active"] = is_active
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if remote_url is None and not opts.urn:
                raise TypeError("Missing required property 'remote_url'")
            __props__.__dict__["remote_url"] = remote_url
            __props__.__dict__["deploy_key"] = None
            __props__.__dict__["repository_credentials_id"] = None
            __props__.__dict__["repository_id"] = None
        super(Repository, __self__).__init__(
            'dbtcloud:index/repository:Repository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            azure_active_directory_project_id: Optional[pulumi.Input[str]] = None,
            azure_active_directory_repository_id: Optional[pulumi.Input[str]] = None,
            azure_bypass_webhook_registration_failure: Optional[pulumi.Input[bool]] = None,
            deploy_key: Optional[pulumi.Input[str]] = None,
            fetch_deploy_key: Optional[pulumi.Input[bool]] = None,
            git_clone_strategy: Optional[pulumi.Input[str]] = None,
            github_installation_id: Optional[pulumi.Input[int]] = None,
            gitlab_project_id: Optional[pulumi.Input[int]] = None,
            is_active: Optional[pulumi.Input[bool]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            remote_url: Optional[pulumi.Input[str]] = None,
            repository_credentials_id: Optional[pulumi.Input[int]] = None,
            repository_id: Optional[pulumi.Input[int]] = None) -> 'Repository':
        """
        Get an existing Repository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] azure_active_directory_project_id: The Azure Dev Ops project ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_project` and the project name - (for ADO native integration only)
        :param pulumi.Input[str] azure_active_directory_repository_id: The Azure Dev Ops repository ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_repository` along with the ADO Project ID and the repository name - (for ADO native integration only)
        :param pulumi.Input[bool] azure_bypass_webhook_registration_failure: If set to False (the default), the connection will fail if the service user doesn't have access to set webhooks (required for auto-triggering CI jobs). If set to True, the connection will be successful but no automated CI job will be triggered - (for ADO native integration only)
        :param pulumi.Input[str] deploy_key: Public key generated by dbt when using `deploy_key` clone strategy
        :param pulumi.Input[bool] fetch_deploy_key: Whether we should return the public deploy key - (for the `deploy_key` strategy)
        :param pulumi.Input[str] git_clone_strategy: Git clone strategy for the repository. Can be `deploy_key` (default) for cloning via SSH Deploy Key, `github_app` for GitHub native integration, `deploy_token` for the GitLab native integration and `azure_active_directory_app` for ADO native integration
        :param pulumi.Input[int] github_installation_id: Identifier for the GitHub App - (for GitHub native integration only)
        :param pulumi.Input[int] gitlab_project_id: Identifier for the Gitlab project -  (for GitLab native integration only)
        :param pulumi.Input[bool] is_active: Whether the repository is active
        :param pulumi.Input[int] project_id: Project ID to create the repository in
        :param pulumi.Input[str] remote_url: Git URL for the repository or <Group>/<Project> for Gitlab
        :param pulumi.Input[int] repository_credentials_id: Credentials ID for the repository (From the repository side not the dbt Cloud ID)
        :param pulumi.Input[int] repository_id: Repository Identifier
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryState.__new__(_RepositoryState)

        __props__.__dict__["azure_active_directory_project_id"] = azure_active_directory_project_id
        __props__.__dict__["azure_active_directory_repository_id"] = azure_active_directory_repository_id
        __props__.__dict__["azure_bypass_webhook_registration_failure"] = azure_bypass_webhook_registration_failure
        __props__.__dict__["deploy_key"] = deploy_key
        __props__.__dict__["fetch_deploy_key"] = fetch_deploy_key
        __props__.__dict__["git_clone_strategy"] = git_clone_strategy
        __props__.__dict__["github_installation_id"] = github_installation_id
        __props__.__dict__["gitlab_project_id"] = gitlab_project_id
        __props__.__dict__["is_active"] = is_active
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["remote_url"] = remote_url
        __props__.__dict__["repository_credentials_id"] = repository_credentials_id
        __props__.__dict__["repository_id"] = repository_id
        return Repository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureActiveDirectoryProjectId")
    def azure_active_directory_project_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Azure Dev Ops project ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_project` and the project name - (for ADO native integration only)
        """
        return pulumi.get(self, "azure_active_directory_project_id")

    @property
    @pulumi.getter(name="azureActiveDirectoryRepositoryId")
    def azure_active_directory_repository_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Azure Dev Ops repository ID. It can be retrieved via the Azure API or using the data source `get_azure_dev_ops_repository` along with the ADO Project ID and the repository name - (for ADO native integration only)
        """
        return pulumi.get(self, "azure_active_directory_repository_id")

    @property
    @pulumi.getter(name="azureBypassWebhookRegistrationFailure")
    def azure_bypass_webhook_registration_failure(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to False (the default), the connection will fail if the service user doesn't have access to set webhooks (required for auto-triggering CI jobs). If set to True, the connection will be successful but no automated CI job will be triggered - (for ADO native integration only)
        """
        return pulumi.get(self, "azure_bypass_webhook_registration_failure")

    @property
    @pulumi.getter(name="deployKey")
    def deploy_key(self) -> pulumi.Output[str]:
        """
        Public key generated by dbt when using `deploy_key` clone strategy
        """
        return pulumi.get(self, "deploy_key")

    @property
    @pulumi.getter(name="fetchDeployKey")
    def fetch_deploy_key(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether we should return the public deploy key - (for the `deploy_key` strategy)
        """
        return pulumi.get(self, "fetch_deploy_key")

    @property
    @pulumi.getter(name="gitCloneStrategy")
    def git_clone_strategy(self) -> pulumi.Output[Optional[str]]:
        """
        Git clone strategy for the repository. Can be `deploy_key` (default) for cloning via SSH Deploy Key, `github_app` for GitHub native integration, `deploy_token` for the GitLab native integration and `azure_active_directory_app` for ADO native integration
        """
        return pulumi.get(self, "git_clone_strategy")

    @property
    @pulumi.getter(name="githubInstallationId")
    def github_installation_id(self) -> pulumi.Output[Optional[int]]:
        """
        Identifier for the GitHub App - (for GitHub native integration only)
        """
        return pulumi.get(self, "github_installation_id")

    @property
    @pulumi.getter(name="gitlabProjectId")
    def gitlab_project_id(self) -> pulumi.Output[Optional[int]]:
        """
        Identifier for the Gitlab project -  (for GitLab native integration only)
        """
        return pulumi.get(self, "gitlab_project_id")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the repository is active
        """
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[int]:
        """
        Project ID to create the repository in
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="remoteUrl")
    def remote_url(self) -> pulumi.Output[str]:
        """
        Git URL for the repository or <Group>/<Project> for Gitlab
        """
        return pulumi.get(self, "remote_url")

    @property
    @pulumi.getter(name="repositoryCredentialsId")
    def repository_credentials_id(self) -> pulumi.Output[int]:
        """
        Credentials ID for the repository (From the repository side not the dbt Cloud ID)
        """
        return pulumi.get(self, "repository_credentials_id")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Output[int]:
        """
        Repository Identifier
        """
        return pulumi.get(self, "repository_id")

