# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetJobResult',
    'AwaitableGetJobResult',
    'get_job',
    'get_job_output',
]

@pulumi.output_type
class GetJobResult:
    """
    A collection of values returned by getJob.
    """
    def __init__(__self__, deferring_environment_id=None, deferring_job_id=None, description=None, environment_id=None, id=None, job_completion_trigger_conditions=None, job_id=None, name=None, project_id=None, self_deferring=None, timeout_seconds=None, triggers=None, triggers_on_draft_pr=None):
        if deferring_environment_id and not isinstance(deferring_environment_id, int):
            raise TypeError("Expected argument 'deferring_environment_id' to be a int")
        pulumi.set(__self__, "deferring_environment_id", deferring_environment_id)
        if deferring_job_id and not isinstance(deferring_job_id, int):
            raise TypeError("Expected argument 'deferring_job_id' to be a int")
        pulumi.set(__self__, "deferring_job_id", deferring_job_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if environment_id and not isinstance(environment_id, int):
            raise TypeError("Expected argument 'environment_id' to be a int")
        pulumi.set(__self__, "environment_id", environment_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if job_completion_trigger_conditions and not isinstance(job_completion_trigger_conditions, list):
            raise TypeError("Expected argument 'job_completion_trigger_conditions' to be a list")
        pulumi.set(__self__, "job_completion_trigger_conditions", job_completion_trigger_conditions)
        if job_id and not isinstance(job_id, int):
            raise TypeError("Expected argument 'job_id' to be a int")
        pulumi.set(__self__, "job_id", job_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if self_deferring and not isinstance(self_deferring, bool):
            raise TypeError("Expected argument 'self_deferring' to be a bool")
        pulumi.set(__self__, "self_deferring", self_deferring)
        if timeout_seconds and not isinstance(timeout_seconds, int):
            raise TypeError("Expected argument 'timeout_seconds' to be a int")
        pulumi.set(__self__, "timeout_seconds", timeout_seconds)
        if triggers and not isinstance(triggers, dict):
            raise TypeError("Expected argument 'triggers' to be a dict")
        pulumi.set(__self__, "triggers", triggers)
        if triggers_on_draft_pr and not isinstance(triggers_on_draft_pr, bool):
            raise TypeError("Expected argument 'triggers_on_draft_pr' to be a bool")
        pulumi.set(__self__, "triggers_on_draft_pr", triggers_on_draft_pr)

    @property
    @pulumi.getter(name="deferringEnvironmentId")
    def deferring_environment_id(self) -> int:
        """
        ID of the environment this job defers to
        """
        return pulumi.get(self, "deferring_environment_id")

    @property
    @pulumi.getter(name="deferringJobId")
    def deferring_job_id(self) -> int:
        """
        ID of the job this job defers to
        """
        return pulumi.get(self, "deferring_job_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Long description for the job
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> int:
        """
        ID of the environment the job is in
        """
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="jobCompletionTriggerConditions")
    def job_completion_trigger_conditions(self) -> Sequence['outputs.GetJobJobCompletionTriggerConditionResult']:
        """
        Which other job should trigger this job when it finishes, and on which conditions.
        """
        return pulumi.get(self, "job_completion_trigger_conditions")

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> int:
        """
        ID of the job
        """
        return pulumi.get(self, "job_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Given name for the job
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        """
        ID of the project the job is in
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="selfDeferring")
    def self_deferring(self) -> bool:
        """
        Whether this job defers on a previous run of itself (overrides value in deferring*job*id)
        """
        return pulumi.get(self, "self_deferring")

    @property
    @pulumi.getter(name="timeoutSeconds")
    def timeout_seconds(self) -> int:
        """
        Number of seconds before the job times out
        """
        return pulumi.get(self, "timeout_seconds")

    @property
    @pulumi.getter
    def triggers(self) -> Mapping[str, bool]:
        """
        Flags for which types of triggers to use, keys of github*webhook, git*provider*webhook, schedule, custom*branch_only
        """
        return pulumi.get(self, "triggers")

    @property
    @pulumi.getter(name="triggersOnDraftPr")
    def triggers_on_draft_pr(self) -> bool:
        """
        Whether the CI job should be automatically triggered on draft PRs
        """
        return pulumi.get(self, "triggers_on_draft_pr")


class AwaitableGetJobResult(GetJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetJobResult(
            deferring_environment_id=self.deferring_environment_id,
            deferring_job_id=self.deferring_job_id,
            description=self.description,
            environment_id=self.environment_id,
            id=self.id,
            job_completion_trigger_conditions=self.job_completion_trigger_conditions,
            job_id=self.job_id,
            name=self.name,
            project_id=self.project_id,
            self_deferring=self.self_deferring,
            timeout_seconds=self.timeout_seconds,
            triggers=self.triggers,
            triggers_on_draft_pr=self.triggers_on_draft_pr)


def get_job(job_id: Optional[int] = None,
            project_id: Optional[int] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetJobResult:
    """
    Use this data source to access information about an existing resource.

    :param int job_id: ID of the job
    :param int project_id: ID of the project the job is in
    """
    __args__ = dict()
    __args__['jobId'] = job_id
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('dbtcloud:index/getJob:getJob', __args__, opts=opts, typ=GetJobResult).value

    return AwaitableGetJobResult(
        deferring_environment_id=pulumi.get(__ret__, 'deferring_environment_id'),
        deferring_job_id=pulumi.get(__ret__, 'deferring_job_id'),
        description=pulumi.get(__ret__, 'description'),
        environment_id=pulumi.get(__ret__, 'environment_id'),
        id=pulumi.get(__ret__, 'id'),
        job_completion_trigger_conditions=pulumi.get(__ret__, 'job_completion_trigger_conditions'),
        job_id=pulumi.get(__ret__, 'job_id'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        self_deferring=pulumi.get(__ret__, 'self_deferring'),
        timeout_seconds=pulumi.get(__ret__, 'timeout_seconds'),
        triggers=pulumi.get(__ret__, 'triggers'),
        triggers_on_draft_pr=pulumi.get(__ret__, 'triggers_on_draft_pr'))


@_utilities.lift_output_func(get_job)
def get_job_output(job_id: Optional[pulumi.Input[int]] = None,
                   project_id: Optional[pulumi.Input[int]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetJobResult]:
    """
    Use this data source to access information about an existing resource.

    :param int job_id: ID of the job
    :param int project_id: ID of the project the job is in
    """
    ...
