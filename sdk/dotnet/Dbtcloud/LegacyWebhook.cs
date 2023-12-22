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
    [DbtcloudResourceType("dbtcloud:index/legacyWebhook:LegacyWebhook")]
    public partial class LegacyWebhook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Webhooks Account Identifier
        /// </summary>
        [Output("accountIdentifier")]
        public Output<string> AccountIdentifier { get; private set; } = null!;

        /// <summary>
        /// Webhooks active flag
        /// </summary>
        [Output("active")]
        public Output<bool?> Active { get; private set; } = null!;

        /// <summary>
        /// Webhooks Client URL
        /// </summary>
        [Output("clientUrl")]
        public Output<string> ClientUrl { get; private set; } = null!;

        /// <summary>
        /// Webhooks Description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Webhooks Event Types
        /// </summary>
        [Output("eventTypes")]
        public Output<ImmutableArray<string>> EventTypes { get; private set; } = null!;

        /// <summary>
        /// Secret key for the webhook. Can be used to validate the authenticity of the webhook.
        /// </summary>
        [Output("hmacSecret")]
        public Output<string> HmacSecret { get; private set; } = null!;

        /// <summary>
        /// Latest HTTP status of the webhook
        /// </summary>
        [Output("httpStatusCode")]
        public Output<string> HttpStatusCode { get; private set; } = null!;

        /// <summary>
        /// List of job IDs to trigger the webhook, An empty list will trigger on all jobs
        /// </summary>
        [Output("jobIds")]
        public Output<ImmutableArray<int>> JobIds { get; private set; } = null!;

        /// <summary>
        /// Webhooks Name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Webhooks ID
        /// </summary>
        [Output("webhookId")]
        public Output<string> WebhookId { get; private set; } = null!;


        /// <summary>
        /// Create a LegacyWebhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LegacyWebhook(string name, LegacyWebhookArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacyWebhook:LegacyWebhook", name, args ?? new LegacyWebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LegacyWebhook(string name, Input<string> id, LegacyWebhookState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacyWebhook:LegacyWebhook", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/a-schot/pulumi-dbtcloud/releases/download/v${VERSION}",
                AdditionalSecretOutputs =
                {
                    "hmacSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LegacyWebhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LegacyWebhook Get(string name, Input<string> id, LegacyWebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new LegacyWebhook(name, id, state, options);
        }
    }

    public sealed class LegacyWebhookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Webhooks active flag
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Webhooks Client URL
        /// </summary>
        [Input("clientUrl", required: true)]
        public Input<string> ClientUrl { get; set; } = null!;

        /// <summary>
        /// Webhooks Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("eventTypes", required: true)]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// Webhooks Event Types
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        [Input("jobIds")]
        private InputList<int>? _jobIds;

        /// <summary>
        /// List of job IDs to trigger the webhook, An empty list will trigger on all jobs
        /// </summary>
        public InputList<int> JobIds
        {
            get => _jobIds ?? (_jobIds = new InputList<int>());
            set => _jobIds = value;
        }

        /// <summary>
        /// Webhooks Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public LegacyWebhookArgs()
        {
        }
        public static new LegacyWebhookArgs Empty => new LegacyWebhookArgs();
    }

    public sealed class LegacyWebhookState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Webhooks Account Identifier
        /// </summary>
        [Input("accountIdentifier")]
        public Input<string>? AccountIdentifier { get; set; }

        /// <summary>
        /// Webhooks active flag
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Webhooks Client URL
        /// </summary>
        [Input("clientUrl")]
        public Input<string>? ClientUrl { get; set; }

        /// <summary>
        /// Webhooks Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("eventTypes")]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// Webhooks Event Types
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        [Input("hmacSecret")]
        private Input<string>? _hmacSecret;

        /// <summary>
        /// Secret key for the webhook. Can be used to validate the authenticity of the webhook.
        /// </summary>
        public Input<string>? HmacSecret
        {
            get => _hmacSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _hmacSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Latest HTTP status of the webhook
        /// </summary>
        [Input("httpStatusCode")]
        public Input<string>? HttpStatusCode { get; set; }

        [Input("jobIds")]
        private InputList<int>? _jobIds;

        /// <summary>
        /// List of job IDs to trigger the webhook, An empty list will trigger on all jobs
        /// </summary>
        public InputList<int> JobIds
        {
            get => _jobIds ?? (_jobIds = new InputList<int>());
            set => _jobIds = value;
        }

        /// <summary>
        /// Webhooks Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Webhooks ID
        /// </summary>
        [Input("webhookId")]
        public Input<string>? WebhookId { get; set; }

        public LegacyWebhookState()
        {
        }
        public static new LegacyWebhookState Empty => new LegacyWebhookState();
    }
}
