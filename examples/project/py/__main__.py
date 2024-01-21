"""A Python Pulumi program"""

import pulumi
import aschot_pulumi_dbtcloud as dbtcloud

project = dbtcloud.Project("py-project")

pulumi.export("project_name", project.name)
