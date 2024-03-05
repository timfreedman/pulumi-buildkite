// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Buildkite
{
    public static class GetMeta
    {
        /// <summary>
        /// Use this data source to look up the source IP addresses that Buildkite may use to send external requests,
        /// including webhooks and API calls to source control systems (like GitHub Enterprise Server and BitBucket Server).
        /// 
        /// More info in the Buildkite [documentation](https://buildkite.com/docs/apis/rest-api/meta).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// using Buildkite = Pulumi.Buildkite;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ips = Buildkite.GetMeta.Invoke();
        /// 
        ///     // Create an AWS security group allowing ingress from Buildkite
        ///     var fromBuildkite = new Aws.Ec2.SecurityGroup("fromBuildkite", new()
        ///     {
        ///         Ingress = new[]
        ///         {
        ///             new Aws.Ec2.Inputs.SecurityGroupIngressArgs
        ///             {
        ///                 FromPort = "*",
        ///                 ToPort = 443,
        ///                 Protocol = "tcp",
        ///                 CidrBlocks = ips.Apply(getMetaResult =&gt; getMetaResult.WebhookIps),
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMetaResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMetaResult>("buildkite:index/getMeta:getMeta", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to look up the source IP addresses that Buildkite may use to send external requests,
        /// including webhooks and API calls to source control systems (like GitHub Enterprise Server and BitBucket Server).
        /// 
        /// More info in the Buildkite [documentation](https://buildkite.com/docs/apis/rest-api/meta).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// using Buildkite = Pulumi.Buildkite;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ips = Buildkite.GetMeta.Invoke();
        /// 
        ///     // Create an AWS security group allowing ingress from Buildkite
        ///     var fromBuildkite = new Aws.Ec2.SecurityGroup("fromBuildkite", new()
        ///     {
        ///         Ingress = new[]
        ///         {
        ///             new Aws.Ec2.Inputs.SecurityGroupIngressArgs
        ///             {
        ///                 FromPort = "*",
        ///                 ToPort = 443,
        ///                 Protocol = "tcp",
        ///                 CidrBlocks = ips.Apply(getMetaResult =&gt; getMetaResult.WebhookIps),
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMetaResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMetaResult>("buildkite:index/getMeta:getMeta", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetMetaResult
    {
        /// <summary>
        /// Fixed value: `https://api.buildkite.com/v2/meta`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of IPs in CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> WebhookIps;

        [OutputConstructor]
        private GetMetaResult(
            string id,

            ImmutableArray<string> webhookIps)
        {
            Id = id;
            WebhookIps = webhookIps;
        }
    }
}
