import * as pulumi from "@pulumi/pulumi";
import * as dbtcloud from "@aschot/pulumi-dbtcloud";

const project =  new dbtcloud.Project("ts-project");

export const projectname = project.name;
