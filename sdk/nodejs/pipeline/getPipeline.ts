// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to look up properties on a specific pipeline. This is particularly useful for looking up the webhook URL for each pipeline.
 *
 * More info in the Buildkite [documentation](https://buildkite.com/docs/pipelines).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as buildkite from "@pulumi/buildkite";
 *
 * const pipeline = buildkite.Pipeline.getPipeline({
 *     slug: "buildkite",
 * });
 * ```
 */
export function getPipeline(args: GetPipelineArgs, opts?: pulumi.InvokeOptions): Promise<GetPipelineResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("buildkite:Pipeline/getPipeline:getPipeline", {
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getPipeline.
 */
export interface GetPipelineArgs {
    /**
     * The slug of the pipeline.
     */
    slug: string;
}

/**
 * A collection of values returned by getPipeline.
 */
export interface GetPipelineResult {
    /**
     * The default branch to prefill when new builds are created or triggered.
     */
    readonly defaultBranch: string;
    /**
     * The description of the pipeline.
     */
    readonly description: string;
    /**
     * The GraphQL ID of the pipeline.
     */
    readonly id: string;
    /**
     * The name of the pipeline.
     */
    readonly name: string;
    /**
     * The git URL of the repository.
     */
    readonly repository: string;
    /**
     * The slug of the pipeline.
     */
    readonly slug: string;
    /**
     * The UUID of the pipeline.
     */
    readonly uuid: string;
    /**
     * The Buildkite webhook URL that triggers builds on this pipeline.
     */
    readonly webhookUrl: string;
}
/**
 * Use this data source to look up properties on a specific pipeline. This is particularly useful for looking up the webhook URL for each pipeline.
 *
 * More info in the Buildkite [documentation](https://buildkite.com/docs/pipelines).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as buildkite from "@pulumi/buildkite";
 *
 * const pipeline = buildkite.Pipeline.getPipeline({
 *     slug: "buildkite",
 * });
 * ```
 */
export function getPipelineOutput(args: GetPipelineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPipelineResult> {
    return pulumi.output(args).apply((a: any) => getPipeline(a, opts))
}

/**
 * A collection of arguments for invoking getPipeline.
 */
export interface GetPipelineOutputArgs {
    /**
     * The slug of the pipeline.
     */
    slug: pulumi.Input<string>;
}
