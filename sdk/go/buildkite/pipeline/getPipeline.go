// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pipeline

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/internal"
)

// Use this data source to look up properties on a specific pipeline. This is particularly useful for looking up the webhook URL for each pipeline.
//
// More info in the Buildkite [documentation](https://buildkite.com/docs/pipelines).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/Pipeline"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Pipeline.GetPipeline(ctx, &pipeline.GetPipelineArgs{
//				Slug: "buildkite",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPipeline(ctx *pulumi.Context, args *LookupPipelineArgs, opts ...pulumi.InvokeOption) (*LookupPipelineResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPipelineResult
	err := ctx.Invoke("buildkite:Pipeline/getPipeline:getPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPipeline.
type LookupPipelineArgs struct {
	// The slug of the pipeline.
	Slug string `pulumi:"slug"`
}

// A collection of values returned by getPipeline.
type LookupPipelineResult struct {
	// The default branch to prefill when new builds are created or triggered.
	DefaultBranch string `pulumi:"defaultBranch"`
	// The description of the pipeline.
	Description string `pulumi:"description"`
	// The GraphQL ID of the pipeline.
	Id string `pulumi:"id"`
	// The name of the pipeline.
	Name string `pulumi:"name"`
	// The git URL of the repository.
	Repository string `pulumi:"repository"`
	// The slug of the pipeline.
	Slug string `pulumi:"slug"`
	// The UUID of the pipeline.
	Uuid string `pulumi:"uuid"`
	// The Buildkite webhook URL that triggers builds on this pipeline.
	WebhookUrl string `pulumi:"webhookUrl"`
}

func LookupPipelineOutput(ctx *pulumi.Context, args LookupPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPipelineResult, error) {
			args := v.(LookupPipelineArgs)
			r, err := LookupPipeline(ctx, &args, opts...)
			var s LookupPipelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPipelineResultOutput)
}

// A collection of arguments for invoking getPipeline.
type LookupPipelineOutputArgs struct {
	// The slug of the pipeline.
	Slug pulumi.StringInput `pulumi:"slug"`
}

func (LookupPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineArgs)(nil)).Elem()
}

// A collection of values returned by getPipeline.
type LookupPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineResult)(nil)).Elem()
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutput() LookupPipelineResultOutput {
	return o
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutputWithContext(ctx context.Context) LookupPipelineResultOutput {
	return o
}

// The default branch to prefill when new builds are created or triggered.
func (o LookupPipelineResultOutput) DefaultBranch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.DefaultBranch }).(pulumi.StringOutput)
}

// The description of the pipeline.
func (o LookupPipelineResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Description }).(pulumi.StringOutput)
}

// The GraphQL ID of the pipeline.
func (o LookupPipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the pipeline.
func (o LookupPipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

// The git URL of the repository.
func (o LookupPipelineResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Repository }).(pulumi.StringOutput)
}

// The slug of the pipeline.
func (o LookupPipelineResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Slug }).(pulumi.StringOutput)
}

// The UUID of the pipeline.
func (o LookupPipelineResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// The Buildkite webhook URL that triggers builds on this pipeline.
func (o LookupPipelineResultOutput) WebhookUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.WebhookUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineResultOutput{})
}
