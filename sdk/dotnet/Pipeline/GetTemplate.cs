// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Buildkite.Pipeline
{
    public static class GetTemplate
    {
        /// <summary>
        /// Use this data source to retrieve a pipeline template by its ID or name.
        /// 
        /// More information on pipeline templates can be found in the [documentation](https://buildkite.com/docs/pipelines/templates).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Buildkite = Pulumi.Buildkite;
        /// using Buildkite = Pulumiverse.Buildkite;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var repository = "git@github.com:my-org/my-repo.git";
        /// 
        ///     var devTemplate = Buildkite.Pipeline.GetTemplate.Invoke(new()
        ///     {
        ///         Id = buildkite_pipeline_template.Template_dev.Id,
        ///     });
        /// 
        ///     var frontendTemplate = Buildkite.Pipeline.GetTemplate.Invoke(new()
        ///     {
        ///         Name = "Frontend app template",
        ///     });
        /// 
        ///     var apiv2Dev = new Buildkite.Pipeline.Pipeline("apiv2Dev", new()
        ///     {
        ///         Repository = repository,
        ///         PipelineTemplateId = devTemplate.Apply(getTemplateResult =&gt; getTemplateResult.Id),
        ///     });
        /// 
        ///     var frontend = new Buildkite.Pipeline.Pipeline("frontend", new()
        ///     {
        ///         Repository = repository,
        ///         PipelineTemplateId = frontendTemplate.Apply(getTemplateResult =&gt; getTemplateResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTemplateResult> InvokeAsync(GetTemplateArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTemplateResult>("buildkite:Pipeline/getTemplate:getTemplate", args ?? new GetTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve a pipeline template by its ID or name.
        /// 
        /// More information on pipeline templates can be found in the [documentation](https://buildkite.com/docs/pipelines/templates).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Buildkite = Pulumi.Buildkite;
        /// using Buildkite = Pulumiverse.Buildkite;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var repository = "git@github.com:my-org/my-repo.git";
        /// 
        ///     var devTemplate = Buildkite.Pipeline.GetTemplate.Invoke(new()
        ///     {
        ///         Id = buildkite_pipeline_template.Template_dev.Id,
        ///     });
        /// 
        ///     var frontendTemplate = Buildkite.Pipeline.GetTemplate.Invoke(new()
        ///     {
        ///         Name = "Frontend app template",
        ///     });
        /// 
        ///     var apiv2Dev = new Buildkite.Pipeline.Pipeline("apiv2Dev", new()
        ///     {
        ///         Repository = repository,
        ///         PipelineTemplateId = devTemplate.Apply(getTemplateResult =&gt; getTemplateResult.Id),
        ///     });
        /// 
        ///     var frontend = new Buildkite.Pipeline.Pipeline("frontend", new()
        ///     {
        ///         Repository = repository,
        ///         PipelineTemplateId = frontendTemplate.Apply(getTemplateResult =&gt; getTemplateResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTemplateResult> Invoke(GetTemplateInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTemplateResult>("buildkite:Pipeline/getTemplate:getTemplate", args ?? new GetTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTemplateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The GraphQL ID of the pipeline template.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the pipeline template.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetTemplateArgs()
        {
        }
        public static new GetTemplateArgs Empty => new GetTemplateArgs();
    }

    public sealed class GetTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The GraphQL ID of the pipeline template.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the pipeline template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetTemplateInvokeArgs()
        {
        }
        public static new GetTemplateInvokeArgs Empty => new GetTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetTemplateResult
    {
        /// <summary>
        /// If the pipeline template is available for assignment by non admin users.
        /// </summary>
        public readonly bool Available;
        /// <summary>
        /// The YAML step configuration for the pipeline template.
        /// </summary>
        public readonly string Configuration;
        /// <summary>
        /// The description for the pipeline template.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The GraphQL ID of the pipeline template.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the pipeline template.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The UUID of the pipeline template.
        /// </summary>
        public readonly string Uuid;

        [OutputConstructor]
        private GetTemplateResult(
            bool available,

            string configuration,

            string description,

            string id,

            string name,

            string uuid)
        {
            Available = available;
            Configuration = configuration;
            Description = description;
            Id = id;
            Name = name;
            Uuid = uuid;
        }
    }
}
