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
    public static class GetWebhook
    {
        public static Task<GetWebhookResult> InvokeAsync(GetWebhookArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebhookResult>("dbtcloud:index/getWebhook:getWebhook", args ?? new GetWebhookArgs(), options.WithDefaults());

        public static Output<GetWebhookResult> Invoke(GetWebhookInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebhookResult>("dbtcloud:index/getWebhook:getWebhook", args ?? new GetWebhookInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebhookArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Webhooks ID
        /// </summary>
        [Input("webhookId", required: true)]
        public string WebhookId { get; set; } = null!;

        public GetWebhookArgs()
        {
        }
        public static new GetWebhookArgs Empty => new GetWebhookArgs();
    }

    public sealed class GetWebhookInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Webhooks ID
        /// </summary>
        [Input("webhookId", required: true)]
        public Input<string> WebhookId { get; set; } = null!;

        public GetWebhookInvokeArgs()
        {
        }
        public static new GetWebhookInvokeArgs Empty => new GetWebhookInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebhookResult
    {
        /// <summary>
        /// Webhooks Account Identifier
        /// </summary>
        public readonly string AccountIdentifier;
        /// <summary>
        /// Webhooks active flag
        /// </summary>
        public readonly bool Active;
        /// <summary>
        /// Webhooks Client URL
        /// </summary>
        public readonly string ClientUrl;
        /// <summary>
        /// Webhooks Description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Webhooks Event Types
        /// </summary>
        public readonly ImmutableArray<string> EventTypes;
        /// <summary>
        /// Webhooks HTTP Status Code
        /// </summary>
        public readonly string HttpStatusCode;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of job IDs to trigger the webhook
        /// </summary>
        public readonly ImmutableArray<int> JobIds;
        /// <summary>
        /// Webhooks Name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Webhooks ID
        /// </summary>
        public readonly string WebhookId;

        [OutputConstructor]
        private GetWebhookResult(
            string accountIdentifier,

            bool active,

            string clientUrl,

            string description,

            ImmutableArray<string> eventTypes,

            string httpStatusCode,

            string id,

            ImmutableArray<int> jobIds,

            string name,

            string webhookId)
        {
            AccountIdentifier = accountIdentifier;
            Active = active;
            ClientUrl = clientUrl;
            Description = description;
            EventTypes = eventTypes;
            HttpStatusCode = httpStatusCode;
            Id = id;
            JobIds = jobIds;
            Name = name;
            WebhookId = webhookId;
        }
    }
}
