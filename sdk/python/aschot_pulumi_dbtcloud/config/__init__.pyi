# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

accountId: Optional[int]
"""
Account identifier for your dbt Cloud implementation. Instead of setting the parameter, you can set the environment
variable `DBT_CLOUD_ACCOUNT_ID`
"""

hostUrl: Optional[str]
"""
URL for your dbt Cloud deployment. Instead of setting the parameter, you can set the environment variable
`DBT_CLOUD_HOST_URL` - Defaults to https://cloud.getdbt.com/api
"""

token: Optional[str]
"""
API token for your dbt Cloud. Instead of setting the parameter, you can set the environment variable `DBT_CLOUD_TOKEN`
"""

