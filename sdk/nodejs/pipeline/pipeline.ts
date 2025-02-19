// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This resource allows you to create and manage pipelines for repositories.
 *
 * More information on pipelines can be found in the [documentation](https://buildkite.com/docs/pipelines).
 *
 * ## Import
 *
 * import a pipeline resource using the pipelines GraphQL ID
 *
 *  GraphQL ID for a pipeline can be found on its settings page
 *
 * ```sh
 * $ pulumi import buildkite:Pipeline/pipeline:Pipeline pipeline UGlwZWxpbmUtLS00MzVjYWQ1OC1lODFkLTQ1YWYtODYzNy1iMWNmODA3MDIzOGQ=
 * ```
 */
export class Pipeline extends pulumi.CustomResource {
    /**
     * Get an existing Pipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PipelineState, opts?: pulumi.CustomResourceOptions): Pipeline {
        return new Pipeline(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'buildkite:Pipeline/pipeline:Pipeline';

    /**
     * Returns true if the given object is an instance of Pipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pipeline.__pulumiType;
    }

    /**
     * Whether rebuilds are allowed for this pipeline.
     */
    public readonly allowRebuilds!: pulumi.Output<boolean>;
    /**
     * The badge URL showing build state.
     */
    public /*out*/ readonly badgeUrl!: pulumi.Output<string>;
    /**
     * Configure the pipeline to only build on this branch conditional.
     */
    public readonly branchConfiguration!: pulumi.Output<string | undefined>;
    /**
     * Whether to cancel builds when a new commit is pushed to a matching branch.
     */
    public readonly cancelIntermediateBuilds!: pulumi.Output<boolean>;
    /**
     * Filter the `cancel_intermediate_builds` setting based on this branch condition.
     */
    public readonly cancelIntermediateBuildsBranchFilter!: pulumi.Output<string>;
    /**
     * Attach this pipeline to the given cluster GraphQL ID.
     */
    public readonly clusterId!: pulumi.Output<string | undefined>;
    /**
     * A color hex code to represent this pipeline.
     */
    public readonly color!: pulumi.Output<string | undefined>;
    /**
     * Default branch of the pipeline.
     */
    public readonly defaultBranch!: pulumi.Output<string>;
    /**
     * The GraphQL ID of the team to use as the default owner of the pipeline.
     */
    public readonly defaultTeamId!: pulumi.Output<string | undefined>;
    /**
     * Set pipeline wide timeout for command steps.
     */
    public readonly defaultTimeoutInMinutes!: pulumi.Output<number>;
    /**
     * Description for the pipeline. Can include emoji 🙌.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * An emoji that represents this pipeline.
     */
    public readonly emoji!: pulumi.Output<string | undefined>;
    /**
     * Set pipeline wide maximum timeout for command steps.
     */
    public readonly maximumTimeoutInMinutes!: pulumi.Output<number>;
    /**
     * Name to give the pipeline.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The GraphQL ID of the pipeline template applied to this pipeline.
     */
    public readonly pipelineTemplateId!: pulumi.Output<string | undefined>;
    /**
     * Control settings depending on the VCS provider used in `repository`.
     */
    public readonly providerSettings!: pulumi.Output<outputs.Pipeline.PipelineProviderSettings | undefined>;
    /**
     * URL to the repository this pipeline is configured for.
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * Whether to skip queued builds if a new commit is pushed to a matching branch.
     */
    public readonly skipIntermediateBuilds!: pulumi.Output<boolean>;
    /**
     * Filter the `skip_intermediate_builds` setting based on this branch condition.
     */
    public readonly skipIntermediateBuildsBranchFilter!: pulumi.Output<string>;
    /**
     * The slug generated for the pipeline.
     */
    public /*out*/ readonly slug!: pulumi.Output<string>;
    /**
     * The YAML steps to configure for the pipeline. Defaults to `buildkite-agent pipeline upload`.
     */
    public readonly steps!: pulumi.Output<string>;
    /**
     * Tags to attribute to the pipeline. Useful for searching by in the UI.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * The UUID of the pipeline.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;
    /**
     * The webhook URL used to trigger builds from VCS providers.
     */
    public /*out*/ readonly webhookUrl!: pulumi.Output<string>;

    /**
     * Create a Pipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PipelineArgs | PipelineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PipelineState | undefined;
            resourceInputs["allowRebuilds"] = state ? state.allowRebuilds : undefined;
            resourceInputs["badgeUrl"] = state ? state.badgeUrl : undefined;
            resourceInputs["branchConfiguration"] = state ? state.branchConfiguration : undefined;
            resourceInputs["cancelIntermediateBuilds"] = state ? state.cancelIntermediateBuilds : undefined;
            resourceInputs["cancelIntermediateBuildsBranchFilter"] = state ? state.cancelIntermediateBuildsBranchFilter : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["defaultBranch"] = state ? state.defaultBranch : undefined;
            resourceInputs["defaultTeamId"] = state ? state.defaultTeamId : undefined;
            resourceInputs["defaultTimeoutInMinutes"] = state ? state.defaultTimeoutInMinutes : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["emoji"] = state ? state.emoji : undefined;
            resourceInputs["maximumTimeoutInMinutes"] = state ? state.maximumTimeoutInMinutes : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pipelineTemplateId"] = state ? state.pipelineTemplateId : undefined;
            resourceInputs["providerSettings"] = state ? state.providerSettings : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["skipIntermediateBuilds"] = state ? state.skipIntermediateBuilds : undefined;
            resourceInputs["skipIntermediateBuildsBranchFilter"] = state ? state.skipIntermediateBuildsBranchFilter : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["steps"] = state ? state.steps : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["webhookUrl"] = state ? state.webhookUrl : undefined;
        } else {
            const args = argsOrState as PipelineArgs | undefined;
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            resourceInputs["allowRebuilds"] = args ? args.allowRebuilds : undefined;
            resourceInputs["branchConfiguration"] = args ? args.branchConfiguration : undefined;
            resourceInputs["cancelIntermediateBuilds"] = args ? args.cancelIntermediateBuilds : undefined;
            resourceInputs["cancelIntermediateBuildsBranchFilter"] = args ? args.cancelIntermediateBuildsBranchFilter : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["defaultBranch"] = args ? args.defaultBranch : undefined;
            resourceInputs["defaultTeamId"] = args ? args.defaultTeamId : undefined;
            resourceInputs["defaultTimeoutInMinutes"] = args ? args.defaultTimeoutInMinutes : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["emoji"] = args ? args.emoji : undefined;
            resourceInputs["maximumTimeoutInMinutes"] = args ? args.maximumTimeoutInMinutes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pipelineTemplateId"] = args ? args.pipelineTemplateId : undefined;
            resourceInputs["providerSettings"] = args ? args.providerSettings : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["skipIntermediateBuilds"] = args ? args.skipIntermediateBuilds : undefined;
            resourceInputs["skipIntermediateBuildsBranchFilter"] = args ? args.skipIntermediateBuildsBranchFilter : undefined;
            resourceInputs["steps"] = args ? args.steps : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["badgeUrl"] = undefined /*out*/;
            resourceInputs["slug"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
            resourceInputs["webhookUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pipeline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pipeline resources.
 */
export interface PipelineState {
    /**
     * Whether rebuilds are allowed for this pipeline.
     */
    allowRebuilds?: pulumi.Input<boolean>;
    /**
     * The badge URL showing build state.
     */
    badgeUrl?: pulumi.Input<string>;
    /**
     * Configure the pipeline to only build on this branch conditional.
     */
    branchConfiguration?: pulumi.Input<string>;
    /**
     * Whether to cancel builds when a new commit is pushed to a matching branch.
     */
    cancelIntermediateBuilds?: pulumi.Input<boolean>;
    /**
     * Filter the `cancel_intermediate_builds` setting based on this branch condition.
     */
    cancelIntermediateBuildsBranchFilter?: pulumi.Input<string>;
    /**
     * Attach this pipeline to the given cluster GraphQL ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * A color hex code to represent this pipeline.
     */
    color?: pulumi.Input<string>;
    /**
     * Default branch of the pipeline.
     */
    defaultBranch?: pulumi.Input<string>;
    /**
     * The GraphQL ID of the team to use as the default owner of the pipeline.
     */
    defaultTeamId?: pulumi.Input<string>;
    /**
     * Set pipeline wide timeout for command steps.
     */
    defaultTimeoutInMinutes?: pulumi.Input<number>;
    /**
     * Description for the pipeline. Can include emoji 🙌.
     */
    description?: pulumi.Input<string>;
    /**
     * An emoji that represents this pipeline.
     */
    emoji?: pulumi.Input<string>;
    /**
     * Set pipeline wide maximum timeout for command steps.
     */
    maximumTimeoutInMinutes?: pulumi.Input<number>;
    /**
     * Name to give the pipeline.
     */
    name?: pulumi.Input<string>;
    /**
     * The GraphQL ID of the pipeline template applied to this pipeline.
     */
    pipelineTemplateId?: pulumi.Input<string>;
    /**
     * Control settings depending on the VCS provider used in `repository`.
     */
    providerSettings?: pulumi.Input<inputs.Pipeline.PipelineProviderSettings>;
    /**
     * URL to the repository this pipeline is configured for.
     */
    repository?: pulumi.Input<string>;
    /**
     * Whether to skip queued builds if a new commit is pushed to a matching branch.
     */
    skipIntermediateBuilds?: pulumi.Input<boolean>;
    /**
     * Filter the `skip_intermediate_builds` setting based on this branch condition.
     */
    skipIntermediateBuildsBranchFilter?: pulumi.Input<string>;
    /**
     * The slug generated for the pipeline.
     */
    slug?: pulumi.Input<string>;
    /**
     * The YAML steps to configure for the pipeline. Defaults to `buildkite-agent pipeline upload`.
     */
    steps?: pulumi.Input<string>;
    /**
     * Tags to attribute to the pipeline. Useful for searching by in the UI.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The UUID of the pipeline.
     */
    uuid?: pulumi.Input<string>;
    /**
     * The webhook URL used to trigger builds from VCS providers.
     */
    webhookUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pipeline resource.
 */
export interface PipelineArgs {
    /**
     * Whether rebuilds are allowed for this pipeline.
     */
    allowRebuilds?: pulumi.Input<boolean>;
    /**
     * Configure the pipeline to only build on this branch conditional.
     */
    branchConfiguration?: pulumi.Input<string>;
    /**
     * Whether to cancel builds when a new commit is pushed to a matching branch.
     */
    cancelIntermediateBuilds?: pulumi.Input<boolean>;
    /**
     * Filter the `cancel_intermediate_builds` setting based on this branch condition.
     */
    cancelIntermediateBuildsBranchFilter?: pulumi.Input<string>;
    /**
     * Attach this pipeline to the given cluster GraphQL ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * A color hex code to represent this pipeline.
     */
    color?: pulumi.Input<string>;
    /**
     * Default branch of the pipeline.
     */
    defaultBranch?: pulumi.Input<string>;
    /**
     * The GraphQL ID of the team to use as the default owner of the pipeline.
     */
    defaultTeamId?: pulumi.Input<string>;
    /**
     * Set pipeline wide timeout for command steps.
     */
    defaultTimeoutInMinutes?: pulumi.Input<number>;
    /**
     * Description for the pipeline. Can include emoji 🙌.
     */
    description?: pulumi.Input<string>;
    /**
     * An emoji that represents this pipeline.
     */
    emoji?: pulumi.Input<string>;
    /**
     * Set pipeline wide maximum timeout for command steps.
     */
    maximumTimeoutInMinutes?: pulumi.Input<number>;
    /**
     * Name to give the pipeline.
     */
    name?: pulumi.Input<string>;
    /**
     * The GraphQL ID of the pipeline template applied to this pipeline.
     */
    pipelineTemplateId?: pulumi.Input<string>;
    /**
     * Control settings depending on the VCS provider used in `repository`.
     */
    providerSettings?: pulumi.Input<inputs.Pipeline.PipelineProviderSettings>;
    /**
     * URL to the repository this pipeline is configured for.
     */
    repository: pulumi.Input<string>;
    /**
     * Whether to skip queued builds if a new commit is pushed to a matching branch.
     */
    skipIntermediateBuilds?: pulumi.Input<boolean>;
    /**
     * Filter the `skip_intermediate_builds` setting based on this branch condition.
     */
    skipIntermediateBuildsBranchFilter?: pulumi.Input<string>;
    /**
     * The YAML steps to configure for the pipeline. Defaults to `buildkite-agent pipeline upload`.
     */
    steps?: pulumi.Input<string>;
    /**
     * Tags to attribute to the pipeline. Useful for searching by in the UI.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
