// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package team

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/internal"
)

// Use this data source to retrieve a team by slug or id. You can find out more about teams in the Buildkite
// [documentation](https://buildkite.com/docs/pipelines/permissions).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/Team"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Team.GetTeam(ctx, &team.GetTeamArgs{
//				Id: pulumi.StringRef("Everyone"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTeam(ctx *pulumi.Context, args *LookupTeamArgs, opts ...pulumi.InvokeOption) (*LookupTeamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTeamResult
	err := ctx.Invoke("buildkite:Team/getTeam:getTeam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTeam.
type LookupTeamArgs struct {
	// The ID of the team to find. Use either ID or slug.
	Id *string `pulumi:"id"`
	// The slug of the team to find. Use either ID or slug.
	Slug *string `pulumi:"slug"`
}

// A collection of values returned by getTeam.
type LookupTeamResult struct {
	// The default member role of the team.
	DefaultMemberRole string `pulumi:"defaultMemberRole"`
	// Whether the team is the default team.
	DefaultTeam bool `pulumi:"defaultTeam"`
	// The description of the team.
	Description string `pulumi:"description"`
	// The ID of the team to find. Use either ID or slug.
	Id string `pulumi:"id"`
	// Whether members can create pipelines.
	MembersCanCreatePipelines bool `pulumi:"membersCanCreatePipelines"`
	// The name of the team.
	Name string `pulumi:"name"`
	// The privacy setting of the team.
	Privacy string `pulumi:"privacy"`
	// The slug of the team to find. Use either ID or slug.
	Slug string `pulumi:"slug"`
	// The UUID of the team.
	Uuid string `pulumi:"uuid"`
}

func LookupTeamOutput(ctx *pulumi.Context, args LookupTeamOutputArgs, opts ...pulumi.InvokeOption) LookupTeamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTeamResult, error) {
			args := v.(LookupTeamArgs)
			r, err := LookupTeam(ctx, &args, opts...)
			var s LookupTeamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTeamResultOutput)
}

// A collection of arguments for invoking getTeam.
type LookupTeamOutputArgs struct {
	// The ID of the team to find. Use either ID or slug.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The slug of the team to find. Use either ID or slug.
	Slug pulumi.StringPtrInput `pulumi:"slug"`
}

func (LookupTeamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTeamArgs)(nil)).Elem()
}

// A collection of values returned by getTeam.
type LookupTeamResultOutput struct{ *pulumi.OutputState }

func (LookupTeamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTeamResult)(nil)).Elem()
}

func (o LookupTeamResultOutput) ToLookupTeamResultOutput() LookupTeamResultOutput {
	return o
}

func (o LookupTeamResultOutput) ToLookupTeamResultOutputWithContext(ctx context.Context) LookupTeamResultOutput {
	return o
}

// The default member role of the team.
func (o LookupTeamResultOutput) DefaultMemberRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.DefaultMemberRole }).(pulumi.StringOutput)
}

// Whether the team is the default team.
func (o LookupTeamResultOutput) DefaultTeam() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTeamResult) bool { return v.DefaultTeam }).(pulumi.BoolOutput)
}

// The description of the team.
func (o LookupTeamResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Description }).(pulumi.StringOutput)
}

// The ID of the team to find. Use either ID or slug.
func (o LookupTeamResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether members can create pipelines.
func (o LookupTeamResultOutput) MembersCanCreatePipelines() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTeamResult) bool { return v.MembersCanCreatePipelines }).(pulumi.BoolOutput)
}

// The name of the team.
func (o LookupTeamResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Name }).(pulumi.StringOutput)
}

// The privacy setting of the team.
func (o LookupTeamResultOutput) Privacy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Privacy }).(pulumi.StringOutput)
}

// The slug of the team to find. Use either ID or slug.
func (o LookupTeamResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Slug }).(pulumi.StringOutput)
}

// The UUID of the team.
func (o LookupTeamResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTeamResultOutput{})
}
