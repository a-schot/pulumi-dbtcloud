// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ASchot.Pulumi.Dbtcloud
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Dbtcloud = ASchot.Pulumi.Dbtcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // dbt Cloud allows us to create internal and external notifications
    ///     // an internal notification will send emails to the user mentioned in `user_id`
    ///     //
    ///     // NOTE: If internal notification settings already exist for a user, currently you MUST import
    ///     // those first into the state file before you can create a new internal notification for that user.
    ///     // Failure to do so, will result in the user losing access to existing notifications and dbt
    ///     // support will need to be contacted to restore access.
    ///     // cmd: terraform import dbtcloud_notification.prod_job_internal_notification &lt;user_id&gt;
    ///     var prodJobInternalNotification = new Dbtcloud.Notification("prodJobInternalNotification", new()
    ///     {
    ///         UserId = 100,
    ///         OnSuccesses = new[]
    ///         {
    ///             dbtcloud_job.Prod_job.Id,
    ///         },
    ///         OnFailures = new[]
    ///         {
    ///             12345,
    ///         },
    ///         NotificationType = 1,
    ///     });
    /// 
    ///     // we can also send "external" email notifications to emails to related to dbt Cloud users
    ///     var prodJobExternalNotification = new Dbtcloud.Notification("prodJobExternalNotification", new()
    ///     {
    ///         UserId = 100,
    ///         OnFailures = new[]
    ///         {
    ///             23456,
    ///             56788,
    ///         },
    ///         OnCancels = new[]
    ///         {
    ///             dbtcloud_job.Prod_job.Id,
    ///         },
    ///         NotificationType = 4,
    ///         ExternalEmail = "my_email@mail.com",
    ///     });
    /// 
    ///     // and finally, we can set up Slack notifications
    ///     var prodJobSlackNotifications = new Dbtcloud.Notification("prodJobSlackNotifications", new()
    ///     {
    ///         UserId = 100,
    ///         OnFailures = new[]
    ///         {
    ///             23456,
    ///             56788,
    ///         },
    ///         OnCancels = new[]
    ///         {
    ///             dbtcloud_job.Prod_job.Id,
    ///         },
    ///         NotificationType = 2,
    ///         SlackChannelId = "C12345ABCDE",
    ///         SlackChannelName = "#my-awesome-channel",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Import using a notification ID
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/notification:Notification my_notification "notification_id"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/notification:Notification my_notification 12345
    /// ```
    /// </summary>
    [DbtcloudResourceType("dbtcloud:index/notification:Notification")]
    public partial class Notification : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The external email to receive the notification
        /// </summary>
        [Output("externalEmail")]
        public Output<string?> ExternalEmail { get; private set; } = null!;

        /// <summary>
        /// Type of notification (1 = dbt Cloud user email (default): does not require an external_email ; 2 = Slack channel: requires `slack_channel_id` and `slack_channel_name` ; 4 = external email: requires setting an `external_email`)
        /// </summary>
        [Output("notificationType")]
        public Output<int?> NotificationType { get; private set; } = null!;

        /// <summary>
        /// List of job IDs to trigger the webhook on cancel
        /// </summary>
        [Output("onCancels")]
        public Output<ImmutableArray<int>> OnCancels { get; private set; } = null!;

        /// <summary>
        /// List of job IDs to trigger the webhook on failure
        /// </summary>
        [Output("onFailures")]
        public Output<ImmutableArray<int>> OnFailures { get; private set; } = null!;

        /// <summary>
        /// List of job IDs to trigger the webhook on success
        /// </summary>
        [Output("onSuccesses")]
        public Output<ImmutableArray<int>> OnSuccesses { get; private set; } = null!;

        /// <summary>
        /// The ID of the Slack channel to receive the notification. It can be found at the bottom of the Slack channel settings
        /// </summary>
        [Output("slackChannelId")]
        public Output<string?> SlackChannelId { get; private set; } = null!;

        /// <summary>
        /// The name of the slack channel
        /// </summary>
        [Output("slackChannelName")]
        public Output<string?> SlackChannelName { get; private set; } = null!;

        /// <summary>
        /// State of the notification (1 = active (default), 2 = inactive)
        /// </summary>
        [Output("state")]
        public Output<int?> State { get; private set; } = null!;

        /// <summary>
        /// Internal dbt Cloud User ID. Must be the user_id for an existing user even if the notification is an external one
        /// </summary>
        [Output("userId")]
        public Output<int> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a Notification resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Notification(string name, NotificationArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/notification:Notification", name, args ?? new NotificationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Notification(string name, Input<string> id, NotificationState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/notification:Notification", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github::api.github.com/a-schot/pulumi-dbtcloud",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Notification resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Notification Get(string name, Input<string> id, NotificationState? state = null, CustomResourceOptions? options = null)
        {
            return new Notification(name, id, state, options);
        }
    }

    public sealed class NotificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The external email to receive the notification
        /// </summary>
        [Input("externalEmail")]
        public Input<string>? ExternalEmail { get; set; }

        /// <summary>
        /// Type of notification (1 = dbt Cloud user email (default): does not require an external_email ; 2 = Slack channel: requires `slack_channel_id` and `slack_channel_name` ; 4 = external email: requires setting an `external_email`)
        /// </summary>
        [Input("notificationType")]
        public Input<int>? NotificationType { get; set; }

        [Input("onCancels")]
        private InputList<int>? _onCancels;

        /// <summary>
        /// List of job IDs to trigger the webhook on cancel
        /// </summary>
        public InputList<int> OnCancels
        {
            get => _onCancels ?? (_onCancels = new InputList<int>());
            set => _onCancels = value;
        }

        [Input("onFailures")]
        private InputList<int>? _onFailures;

        /// <summary>
        /// List of job IDs to trigger the webhook on failure
        /// </summary>
        public InputList<int> OnFailures
        {
            get => _onFailures ?? (_onFailures = new InputList<int>());
            set => _onFailures = value;
        }

        [Input("onSuccesses")]
        private InputList<int>? _onSuccesses;

        /// <summary>
        /// List of job IDs to trigger the webhook on success
        /// </summary>
        public InputList<int> OnSuccesses
        {
            get => _onSuccesses ?? (_onSuccesses = new InputList<int>());
            set => _onSuccesses = value;
        }

        /// <summary>
        /// The ID of the Slack channel to receive the notification. It can be found at the bottom of the Slack channel settings
        /// </summary>
        [Input("slackChannelId")]
        public Input<string>? SlackChannelId { get; set; }

        /// <summary>
        /// The name of the slack channel
        /// </summary>
        [Input("slackChannelName")]
        public Input<string>? SlackChannelName { get; set; }

        /// <summary>
        /// State of the notification (1 = active (default), 2 = inactive)
        /// </summary>
        [Input("state")]
        public Input<int>? State { get; set; }

        /// <summary>
        /// Internal dbt Cloud User ID. Must be the user_id for an existing user even if the notification is an external one
        /// </summary>
        [Input("userId", required: true)]
        public Input<int> UserId { get; set; } = null!;

        public NotificationArgs()
        {
        }
        public static new NotificationArgs Empty => new NotificationArgs();
    }

    public sealed class NotificationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The external email to receive the notification
        /// </summary>
        [Input("externalEmail")]
        public Input<string>? ExternalEmail { get; set; }

        /// <summary>
        /// Type of notification (1 = dbt Cloud user email (default): does not require an external_email ; 2 = Slack channel: requires `slack_channel_id` and `slack_channel_name` ; 4 = external email: requires setting an `external_email`)
        /// </summary>
        [Input("notificationType")]
        public Input<int>? NotificationType { get; set; }

        [Input("onCancels")]
        private InputList<int>? _onCancels;

        /// <summary>
        /// List of job IDs to trigger the webhook on cancel
        /// </summary>
        public InputList<int> OnCancels
        {
            get => _onCancels ?? (_onCancels = new InputList<int>());
            set => _onCancels = value;
        }

        [Input("onFailures")]
        private InputList<int>? _onFailures;

        /// <summary>
        /// List of job IDs to trigger the webhook on failure
        /// </summary>
        public InputList<int> OnFailures
        {
            get => _onFailures ?? (_onFailures = new InputList<int>());
            set => _onFailures = value;
        }

        [Input("onSuccesses")]
        private InputList<int>? _onSuccesses;

        /// <summary>
        /// List of job IDs to trigger the webhook on success
        /// </summary>
        public InputList<int> OnSuccesses
        {
            get => _onSuccesses ?? (_onSuccesses = new InputList<int>());
            set => _onSuccesses = value;
        }

        /// <summary>
        /// The ID of the Slack channel to receive the notification. It can be found at the bottom of the Slack channel settings
        /// </summary>
        [Input("slackChannelId")]
        public Input<string>? SlackChannelId { get; set; }

        /// <summary>
        /// The name of the slack channel
        /// </summary>
        [Input("slackChannelName")]
        public Input<string>? SlackChannelName { get; set; }

        /// <summary>
        /// State of the notification (1 = active (default), 2 = inactive)
        /// </summary>
        [Input("state")]
        public Input<int>? State { get; set; }

        /// <summary>
        /// Internal dbt Cloud User ID. Must be the user_id for an existing user even if the notification is an external one
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public NotificationState()
        {
        }
        public static new NotificationState Empty => new NotificationState();
    }
}
