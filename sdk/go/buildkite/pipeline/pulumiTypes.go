// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pipeline

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/internal"
)

var _ = internal.GetEnvOrDefault

type PipelineProviderSettings struct {
	// Whether to create builds when branches are pushed.
	BuildBranches *bool `pulumi:"buildBranches"`
	// Whether to create builds for pull requests from third-party forks.
	BuildPullRequestForks *bool `pulumi:"buildPullRequestForks"`
	// Whether to create builds for pull requests when labels are added or removed.
	BuildPullRequestLabelsChanged *bool `pulumi:"buildPullRequestLabelsChanged"`
	// Whether to create a build when a pull request changes to "Ready for review".
	BuildPullRequestReadyForReview *bool `pulumi:"buildPullRequestReadyForReview"`
	// Whether to create builds for commits that are part of a pull request.
	BuildPullRequests *bool `pulumi:"buildPullRequests"`
	// Whether to create builds when tags are pushed.
	BuildTags *bool `pulumi:"buildTags"`
	// Automatically cancel running builds for a branch if the branch is deleted.
	CancelDeletedBranchBuilds *bool `pulumi:"cancelDeletedBranchBuilds"`
	// The condition to evaluate when deciding if a build should run. More details available in [the documentation](https://buildkite.com/docs/pipelines/conditionals#conditionals-in-pipelines).
	FilterCondition *string `pulumi:"filterCondition"`
	// Whether to filter builds to only run when the condition in `filterCondition` is true.
	FilterEnabled *bool `pulumi:"filterEnabled"`
	// Prefix branch names for third-party fork builds to ensure they don't trigger branch conditions. For example, the main branch from some-user will become some-user:main.
	PrefixPullRequestForkBranchNames *bool `pulumi:"prefixPullRequestForkBranchNames"`
	// The status to use for blocked builds. Pending can be used with [required status checks](https://help.github.com/en/articles/enabling-required-status-checks) to prevent merging pull requests with blocked builds.
	PublishBlockedAsPending *bool `pulumi:"publishBlockedAsPending"`
	// Whether to update the status of commits in Bitbucket or GitHub.
	PublishCommitStatus *bool `pulumi:"publishCommitStatus"`
	// Whether to create a separate status for each job in a build, allowing you to see the status of each job directly in Bitbucket or GitHub.
	PublishCommitStatusPerStep *bool `pulumi:"publishCommitStatusPerStep"`
	// Filter pull requests builds by the branch filter.
	PullRequestBranchFilterConfiguration *string `pulumi:"pullRequestBranchFilterConfiguration"`
	// Filter pull request builds.
	PullRequestBranchFilterEnabled *bool `pulumi:"pullRequestBranchFilterEnabled"`
	// Whether to create a separate status for pull request builds, allowing you to require a passing pull request build in your [required status checks](https://help.github.com/en/articles/enabling-required-status-checks) in GitHub.
	SeparatePullRequestStatuses *bool `pulumi:"separatePullRequestStatuses"`
	// Whether to skip creating a new build if an existing build for the commit and branch already exists. This option is only valid if the pipeline uses a GitHub repository.
	SkipBuildsForExistingCommits *bool `pulumi:"skipBuildsForExistingCommits"`
	// Whether to skip creating a new build for a pull request if an existing build for the commit and branch already exists.
	SkipPullRequestBuildsForExistingCommits *bool `pulumi:"skipPullRequestBuildsForExistingCommits"`
	// What type of event to trigger builds on. Must be one of:
	// 	- `code` will create builds when code is pushed to GitHub.
	// 	- `deployment` will create builds when a deployment is created in GitHub.
	// 	- `fork` will create builds when the GitHub repository is forked.
	// 	- `none` will not create any builds based on GitHub activity.
	//
	//     > `triggerMode` is only valid if the pipeline uses a GitHub repository.
	TriggerMode *string `pulumi:"triggerMode"`
}

// PipelineProviderSettingsInput is an input type that accepts PipelineProviderSettingsArgs and PipelineProviderSettingsOutput values.
// You can construct a concrete instance of `PipelineProviderSettingsInput` via:
//
//	PipelineProviderSettingsArgs{...}
type PipelineProviderSettingsInput interface {
	pulumi.Input

	ToPipelineProviderSettingsOutput() PipelineProviderSettingsOutput
	ToPipelineProviderSettingsOutputWithContext(context.Context) PipelineProviderSettingsOutput
}

type PipelineProviderSettingsArgs struct {
	// Whether to create builds when branches are pushed.
	BuildBranches pulumi.BoolPtrInput `pulumi:"buildBranches"`
	// Whether to create builds for pull requests from third-party forks.
	BuildPullRequestForks pulumi.BoolPtrInput `pulumi:"buildPullRequestForks"`
	// Whether to create builds for pull requests when labels are added or removed.
	BuildPullRequestLabelsChanged pulumi.BoolPtrInput `pulumi:"buildPullRequestLabelsChanged"`
	// Whether to create a build when a pull request changes to "Ready for review".
	BuildPullRequestReadyForReview pulumi.BoolPtrInput `pulumi:"buildPullRequestReadyForReview"`
	// Whether to create builds for commits that are part of a pull request.
	BuildPullRequests pulumi.BoolPtrInput `pulumi:"buildPullRequests"`
	// Whether to create builds when tags are pushed.
	BuildTags pulumi.BoolPtrInput `pulumi:"buildTags"`
	// Automatically cancel running builds for a branch if the branch is deleted.
	CancelDeletedBranchBuilds pulumi.BoolPtrInput `pulumi:"cancelDeletedBranchBuilds"`
	// The condition to evaluate when deciding if a build should run. More details available in [the documentation](https://buildkite.com/docs/pipelines/conditionals#conditionals-in-pipelines).
	FilterCondition pulumi.StringPtrInput `pulumi:"filterCondition"`
	// Whether to filter builds to only run when the condition in `filterCondition` is true.
	FilterEnabled pulumi.BoolPtrInput `pulumi:"filterEnabled"`
	// Prefix branch names for third-party fork builds to ensure they don't trigger branch conditions. For example, the main branch from some-user will become some-user:main.
	PrefixPullRequestForkBranchNames pulumi.BoolPtrInput `pulumi:"prefixPullRequestForkBranchNames"`
	// The status to use for blocked builds. Pending can be used with [required status checks](https://help.github.com/en/articles/enabling-required-status-checks) to prevent merging pull requests with blocked builds.
	PublishBlockedAsPending pulumi.BoolPtrInput `pulumi:"publishBlockedAsPending"`
	// Whether to update the status of commits in Bitbucket or GitHub.
	PublishCommitStatus pulumi.BoolPtrInput `pulumi:"publishCommitStatus"`
	// Whether to create a separate status for each job in a build, allowing you to see the status of each job directly in Bitbucket or GitHub.
	PublishCommitStatusPerStep pulumi.BoolPtrInput `pulumi:"publishCommitStatusPerStep"`
	// Filter pull requests builds by the branch filter.
	PullRequestBranchFilterConfiguration pulumi.StringPtrInput `pulumi:"pullRequestBranchFilterConfiguration"`
	// Filter pull request builds.
	PullRequestBranchFilterEnabled pulumi.BoolPtrInput `pulumi:"pullRequestBranchFilterEnabled"`
	// Whether to create a separate status for pull request builds, allowing you to require a passing pull request build in your [required status checks](https://help.github.com/en/articles/enabling-required-status-checks) in GitHub.
	SeparatePullRequestStatuses pulumi.BoolPtrInput `pulumi:"separatePullRequestStatuses"`
	// Whether to skip creating a new build if an existing build for the commit and branch already exists. This option is only valid if the pipeline uses a GitHub repository.
	SkipBuildsForExistingCommits pulumi.BoolPtrInput `pulumi:"skipBuildsForExistingCommits"`
	// Whether to skip creating a new build for a pull request if an existing build for the commit and branch already exists.
	SkipPullRequestBuildsForExistingCommits pulumi.BoolPtrInput `pulumi:"skipPullRequestBuildsForExistingCommits"`
	// What type of event to trigger builds on. Must be one of:
	// 	- `code` will create builds when code is pushed to GitHub.
	// 	- `deployment` will create builds when a deployment is created in GitHub.
	// 	- `fork` will create builds when the GitHub repository is forked.
	// 	- `none` will not create any builds based on GitHub activity.
	//
	//     > `triggerMode` is only valid if the pipeline uses a GitHub repository.
	TriggerMode pulumi.StringPtrInput `pulumi:"triggerMode"`
}

func (PipelineProviderSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineProviderSettings)(nil)).Elem()
}

func (i PipelineProviderSettingsArgs) ToPipelineProviderSettingsOutput() PipelineProviderSettingsOutput {
	return i.ToPipelineProviderSettingsOutputWithContext(context.Background())
}

func (i PipelineProviderSettingsArgs) ToPipelineProviderSettingsOutputWithContext(ctx context.Context) PipelineProviderSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineProviderSettingsOutput)
}

func (i PipelineProviderSettingsArgs) ToPipelineProviderSettingsPtrOutput() PipelineProviderSettingsPtrOutput {
	return i.ToPipelineProviderSettingsPtrOutputWithContext(context.Background())
}

func (i PipelineProviderSettingsArgs) ToPipelineProviderSettingsPtrOutputWithContext(ctx context.Context) PipelineProviderSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineProviderSettingsOutput).ToPipelineProviderSettingsPtrOutputWithContext(ctx)
}

// PipelineProviderSettingsPtrInput is an input type that accepts PipelineProviderSettingsArgs, PipelineProviderSettingsPtr and PipelineProviderSettingsPtrOutput values.
// You can construct a concrete instance of `PipelineProviderSettingsPtrInput` via:
//
//	        PipelineProviderSettingsArgs{...}
//
//	or:
//
//	        nil
type PipelineProviderSettingsPtrInput interface {
	pulumi.Input

	ToPipelineProviderSettingsPtrOutput() PipelineProviderSettingsPtrOutput
	ToPipelineProviderSettingsPtrOutputWithContext(context.Context) PipelineProviderSettingsPtrOutput
}

type pipelineProviderSettingsPtrType PipelineProviderSettingsArgs

func PipelineProviderSettingsPtr(v *PipelineProviderSettingsArgs) PipelineProviderSettingsPtrInput {
	return (*pipelineProviderSettingsPtrType)(v)
}

func (*pipelineProviderSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineProviderSettings)(nil)).Elem()
}

func (i *pipelineProviderSettingsPtrType) ToPipelineProviderSettingsPtrOutput() PipelineProviderSettingsPtrOutput {
	return i.ToPipelineProviderSettingsPtrOutputWithContext(context.Background())
}

func (i *pipelineProviderSettingsPtrType) ToPipelineProviderSettingsPtrOutputWithContext(ctx context.Context) PipelineProviderSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineProviderSettingsPtrOutput)
}

type PipelineProviderSettingsOutput struct{ *pulumi.OutputState }

func (PipelineProviderSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineProviderSettings)(nil)).Elem()
}

func (o PipelineProviderSettingsOutput) ToPipelineProviderSettingsOutput() PipelineProviderSettingsOutput {
	return o
}

func (o PipelineProviderSettingsOutput) ToPipelineProviderSettingsOutputWithContext(ctx context.Context) PipelineProviderSettingsOutput {
	return o
}

func (o PipelineProviderSettingsOutput) ToPipelineProviderSettingsPtrOutput() PipelineProviderSettingsPtrOutput {
	return o.ToPipelineProviderSettingsPtrOutputWithContext(context.Background())
}

func (o PipelineProviderSettingsOutput) ToPipelineProviderSettingsPtrOutputWithContext(ctx context.Context) PipelineProviderSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineProviderSettings) *PipelineProviderSettings {
		return &v
	}).(PipelineProviderSettingsPtrOutput)
}

// Whether to create builds when branches are pushed.
func (o PipelineProviderSettingsOutput) BuildBranches() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.BuildBranches }).(pulumi.BoolPtrOutput)
}

// Whether to create builds for pull requests from third-party forks.
func (o PipelineProviderSettingsOutput) BuildPullRequestForks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.BuildPullRequestForks }).(pulumi.BoolPtrOutput)
}

// Whether to create builds for pull requests when labels are added or removed.
func (o PipelineProviderSettingsOutput) BuildPullRequestLabelsChanged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.BuildPullRequestLabelsChanged }).(pulumi.BoolPtrOutput)
}

// Whether to create a build when a pull request changes to "Ready for review".
func (o PipelineProviderSettingsOutput) BuildPullRequestReadyForReview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.BuildPullRequestReadyForReview }).(pulumi.BoolPtrOutput)
}

// Whether to create builds for commits that are part of a pull request.
func (o PipelineProviderSettingsOutput) BuildPullRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.BuildPullRequests }).(pulumi.BoolPtrOutput)
}

// Whether to create builds when tags are pushed.
func (o PipelineProviderSettingsOutput) BuildTags() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.BuildTags }).(pulumi.BoolPtrOutput)
}

// Automatically cancel running builds for a branch if the branch is deleted.
func (o PipelineProviderSettingsOutput) CancelDeletedBranchBuilds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.CancelDeletedBranchBuilds }).(pulumi.BoolPtrOutput)
}

// The condition to evaluate when deciding if a build should run. More details available in [the documentation](https://buildkite.com/docs/pipelines/conditionals#conditionals-in-pipelines).
func (o PipelineProviderSettingsOutput) FilterCondition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *string { return v.FilterCondition }).(pulumi.StringPtrOutput)
}

// Whether to filter builds to only run when the condition in `filterCondition` is true.
func (o PipelineProviderSettingsOutput) FilterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.FilterEnabled }).(pulumi.BoolPtrOutput)
}

// Prefix branch names for third-party fork builds to ensure they don't trigger branch conditions. For example, the main branch from some-user will become some-user:main.
func (o PipelineProviderSettingsOutput) PrefixPullRequestForkBranchNames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.PrefixPullRequestForkBranchNames }).(pulumi.BoolPtrOutput)
}

// The status to use for blocked builds. Pending can be used with [required status checks](https://help.github.com/en/articles/enabling-required-status-checks) to prevent merging pull requests with blocked builds.
func (o PipelineProviderSettingsOutput) PublishBlockedAsPending() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.PublishBlockedAsPending }).(pulumi.BoolPtrOutput)
}

// Whether to update the status of commits in Bitbucket or GitHub.
func (o PipelineProviderSettingsOutput) PublishCommitStatus() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.PublishCommitStatus }).(pulumi.BoolPtrOutput)
}

// Whether to create a separate status for each job in a build, allowing you to see the status of each job directly in Bitbucket or GitHub.
func (o PipelineProviderSettingsOutput) PublishCommitStatusPerStep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.PublishCommitStatusPerStep }).(pulumi.BoolPtrOutput)
}

// Filter pull requests builds by the branch filter.
func (o PipelineProviderSettingsOutput) PullRequestBranchFilterConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *string { return v.PullRequestBranchFilterConfiguration }).(pulumi.StringPtrOutput)
}

// Filter pull request builds.
func (o PipelineProviderSettingsOutput) PullRequestBranchFilterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.PullRequestBranchFilterEnabled }).(pulumi.BoolPtrOutput)
}

// Whether to create a separate status for pull request builds, allowing you to require a passing pull request build in your [required status checks](https://help.github.com/en/articles/enabling-required-status-checks) in GitHub.
func (o PipelineProviderSettingsOutput) SeparatePullRequestStatuses() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.SeparatePullRequestStatuses }).(pulumi.BoolPtrOutput)
}

// Whether to skip creating a new build if an existing build for the commit and branch already exists. This option is only valid if the pipeline uses a GitHub repository.
func (o PipelineProviderSettingsOutput) SkipBuildsForExistingCommits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.SkipBuildsForExistingCommits }).(pulumi.BoolPtrOutput)
}

// Whether to skip creating a new build for a pull request if an existing build for the commit and branch already exists.
func (o PipelineProviderSettingsOutput) SkipPullRequestBuildsForExistingCommits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *bool { return v.SkipPullRequestBuildsForExistingCommits }).(pulumi.BoolPtrOutput)
}

// What type of event to trigger builds on. Must be one of:
//
//   - `code` will create builds when code is pushed to GitHub.
//
//   - `deployment` will create builds when a deployment is created in GitHub.
//
//   - `fork` will create builds when the GitHub repository is forked.
//
//   - `none` will not create any builds based on GitHub activity.
//
//     > `triggerMode` is only valid if the pipeline uses a GitHub repository.
func (o PipelineProviderSettingsOutput) TriggerMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineProviderSettings) *string { return v.TriggerMode }).(pulumi.StringPtrOutput)
}

type PipelineProviderSettingsPtrOutput struct{ *pulumi.OutputState }

func (PipelineProviderSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineProviderSettings)(nil)).Elem()
}

func (o PipelineProviderSettingsPtrOutput) ToPipelineProviderSettingsPtrOutput() PipelineProviderSettingsPtrOutput {
	return o
}

func (o PipelineProviderSettingsPtrOutput) ToPipelineProviderSettingsPtrOutputWithContext(ctx context.Context) PipelineProviderSettingsPtrOutput {
	return o
}

func (o PipelineProviderSettingsPtrOutput) Elem() PipelineProviderSettingsOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) PipelineProviderSettings {
		if v != nil {
			return *v
		}
		var ret PipelineProviderSettings
		return ret
	}).(PipelineProviderSettingsOutput)
}

// Whether to create builds when branches are pushed.
func (o PipelineProviderSettingsPtrOutput) BuildBranches() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.BuildBranches
	}).(pulumi.BoolPtrOutput)
}

// Whether to create builds for pull requests from third-party forks.
func (o PipelineProviderSettingsPtrOutput) BuildPullRequestForks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.BuildPullRequestForks
	}).(pulumi.BoolPtrOutput)
}

// Whether to create builds for pull requests when labels are added or removed.
func (o PipelineProviderSettingsPtrOutput) BuildPullRequestLabelsChanged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.BuildPullRequestLabelsChanged
	}).(pulumi.BoolPtrOutput)
}

// Whether to create a build when a pull request changes to "Ready for review".
func (o PipelineProviderSettingsPtrOutput) BuildPullRequestReadyForReview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.BuildPullRequestReadyForReview
	}).(pulumi.BoolPtrOutput)
}

// Whether to create builds for commits that are part of a pull request.
func (o PipelineProviderSettingsPtrOutput) BuildPullRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.BuildPullRequests
	}).(pulumi.BoolPtrOutput)
}

// Whether to create builds when tags are pushed.
func (o PipelineProviderSettingsPtrOutput) BuildTags() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.BuildTags
	}).(pulumi.BoolPtrOutput)
}

// Automatically cancel running builds for a branch if the branch is deleted.
func (o PipelineProviderSettingsPtrOutput) CancelDeletedBranchBuilds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CancelDeletedBranchBuilds
	}).(pulumi.BoolPtrOutput)
}

// The condition to evaluate when deciding if a build should run. More details available in [the documentation](https://buildkite.com/docs/pipelines/conditionals#conditionals-in-pipelines).
func (o PipelineProviderSettingsPtrOutput) FilterCondition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *string {
		if v == nil {
			return nil
		}
		return v.FilterCondition
	}).(pulumi.StringPtrOutput)
}

// Whether to filter builds to only run when the condition in `filterCondition` is true.
func (o PipelineProviderSettingsPtrOutput) FilterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.FilterEnabled
	}).(pulumi.BoolPtrOutput)
}

// Prefix branch names for third-party fork builds to ensure they don't trigger branch conditions. For example, the main branch from some-user will become some-user:main.
func (o PipelineProviderSettingsPtrOutput) PrefixPullRequestForkBranchNames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.PrefixPullRequestForkBranchNames
	}).(pulumi.BoolPtrOutput)
}

// The status to use for blocked builds. Pending can be used with [required status checks](https://help.github.com/en/articles/enabling-required-status-checks) to prevent merging pull requests with blocked builds.
func (o PipelineProviderSettingsPtrOutput) PublishBlockedAsPending() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.PublishBlockedAsPending
	}).(pulumi.BoolPtrOutput)
}

// Whether to update the status of commits in Bitbucket or GitHub.
func (o PipelineProviderSettingsPtrOutput) PublishCommitStatus() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.PublishCommitStatus
	}).(pulumi.BoolPtrOutput)
}

// Whether to create a separate status for each job in a build, allowing you to see the status of each job directly in Bitbucket or GitHub.
func (o PipelineProviderSettingsPtrOutput) PublishCommitStatusPerStep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.PublishCommitStatusPerStep
	}).(pulumi.BoolPtrOutput)
}

// Filter pull requests builds by the branch filter.
func (o PipelineProviderSettingsPtrOutput) PullRequestBranchFilterConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *string {
		if v == nil {
			return nil
		}
		return v.PullRequestBranchFilterConfiguration
	}).(pulumi.StringPtrOutput)
}

// Filter pull request builds.
func (o PipelineProviderSettingsPtrOutput) PullRequestBranchFilterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.PullRequestBranchFilterEnabled
	}).(pulumi.BoolPtrOutput)
}

// Whether to create a separate status for pull request builds, allowing you to require a passing pull request build in your [required status checks](https://help.github.com/en/articles/enabling-required-status-checks) in GitHub.
func (o PipelineProviderSettingsPtrOutput) SeparatePullRequestStatuses() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SeparatePullRequestStatuses
	}).(pulumi.BoolPtrOutput)
}

// Whether to skip creating a new build if an existing build for the commit and branch already exists. This option is only valid if the pipeline uses a GitHub repository.
func (o PipelineProviderSettingsPtrOutput) SkipBuildsForExistingCommits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SkipBuildsForExistingCommits
	}).(pulumi.BoolPtrOutput)
}

// Whether to skip creating a new build for a pull request if an existing build for the commit and branch already exists.
func (o PipelineProviderSettingsPtrOutput) SkipPullRequestBuildsForExistingCommits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SkipPullRequestBuildsForExistingCommits
	}).(pulumi.BoolPtrOutput)
}

// What type of event to trigger builds on. Must be one of:
//
//   - `code` will create builds when code is pushed to GitHub.
//
//   - `deployment` will create builds when a deployment is created in GitHub.
//
//   - `fork` will create builds when the GitHub repository is forked.
//
//   - `none` will not create any builds based on GitHub activity.
//
//     > `triggerMode` is only valid if the pipeline uses a GitHub repository.
func (o PipelineProviderSettingsPtrOutput) TriggerMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineProviderSettings) *string {
		if v == nil {
			return nil
		}
		return v.TriggerMode
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineProviderSettingsInput)(nil)).Elem(), PipelineProviderSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineProviderSettingsPtrInput)(nil)).Elem(), PipelineProviderSettingsArgs{})
	pulumi.RegisterOutputType(PipelineProviderSettingsOutput{})
	pulumi.RegisterOutputType(PipelineProviderSettingsPtrOutput{})
}
