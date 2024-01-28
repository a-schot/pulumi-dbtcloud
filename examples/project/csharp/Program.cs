using System.Collections.Generic;
using Pulumi;
using ASchot.Pulumi.Dbtcloud;

return await Deployment.RunAsync(() =>
{
    // Add your resources here
    // e.g. var resource = new Resource("name", new ResourceArgs { });
    var project = new Project("cs-project", new ProjectArgs { Name = "cs-project" });

    // Export outputs here
    return new Dictionary<string, object?>
    {
        ["ProjectName"] = project.Name
    };
});
