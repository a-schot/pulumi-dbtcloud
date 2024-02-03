using System.Collections.Generic;
using Pulumi;
using ASchot.Pulumi.Dbtcloud;

return await Deployment.RunAsync(() =>
{
    var project = new Project("cs-project", new ProjectArgs { Name = "cs-project" });

    return new Dictionary<string, object?>
    {
        ["ProjectName"] = project.Name
    };
});
