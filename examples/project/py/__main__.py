"""A Python Pulumi program"""

import pulumi
import aschot_pulumi_dbtcloud as dbtcloud

project = dbtcloud.Project("project", name="project")

pulumi.export(project)
