// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package testsuite

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/internal"
)

// Manage team access to a test suite.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/TestSuite"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := TestSuite.NewTestSuite(ctx, "main", &TestSuite.TestSuiteArgs{
//				DefaultBranch: pulumi.String("main"),
//				TeamOwnerId:   pulumi.String("VGVhbU1lbWJlci0tLTVlZDEyMmY2LTM2NjQtNDI1MS04YzMwLTc4NjRiMDdiZDQ4Zg=="),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = TestSuite.NewTeam(ctx, "mainEveryone", &TestSuite.TeamArgs{
//				TestSuiteId: main.ID(),
//				TeamId:      pulumi.String("VGVhbS0tLWU1YjQyMDQyLTUzN2QtNDZjNi04MjY0LTliZjFkMzkyYjZkNQ=="),
//				AccessLevel: pulumi.String("MANAGE_AND_READ"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// import a test suite team resource using the GraphQL ID
//
// #
//
//	you can use this query to find the ID:
//
//	query getTeamSuiteIds {
//
//	organization(slug: "ORGANIZATION_SLUG") {
//
//	suites(first: 1, search:"SUITE_SEARCH_TERM") {
//
//	edges {
//
//	node {
//
//	id
//
//	name
//
//	teams(first: 10){
//
//	edges {
//
//	node {
//
//	id
//
//	accessLevel
//
//	team{
//
//	name
//
//	}
//
//	}
//
//	}
//
//	}
//
//	}
//
//	}
//
//	}
//
//	}
//
//	}
//
// ```sh
// $ pulumi import buildkite:TestSuite/team:Team main_everyone VGVhbvDf4eRef20tMzIxMGEfYTctNzEF5g00M8f5s6E2YjYtODNlOGNlZgD6HcBi
// ```
type Team struct {
	pulumi.CustomResourceState

	// The access level the team has on the test suite. Either `READ_ONLY` or `MANAGE_AND_READ`.
	AccessLevel pulumi.StringOutput `pulumi:"accessLevel"`
	// The GraphQL ID of the team.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
	// The GraphQL ID of the test suite.
	TestSuiteId pulumi.StringOutput `pulumi:"testSuiteId"`
	// The UUID of the test suite-team relationship.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewTeam registers a new resource with the given unique name, arguments, and options.
func NewTeam(ctx *pulumi.Context,
	name string, args *TeamArgs, opts ...pulumi.ResourceOption) (*Team, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessLevel == nil {
		return nil, errors.New("invalid value for required argument 'AccessLevel'")
	}
	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	if args.TestSuiteId == nil {
		return nil, errors.New("invalid value for required argument 'TestSuiteId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Team
	err := ctx.RegisterResource("buildkite:TestSuite/team:Team", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeam gets an existing Team resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamState, opts ...pulumi.ResourceOption) (*Team, error) {
	var resource Team
	err := ctx.ReadResource("buildkite:TestSuite/team:Team", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Team resources.
type teamState struct {
	// The access level the team has on the test suite. Either `READ_ONLY` or `MANAGE_AND_READ`.
	AccessLevel *string `pulumi:"accessLevel"`
	// The GraphQL ID of the team.
	TeamId *string `pulumi:"teamId"`
	// The GraphQL ID of the test suite.
	TestSuiteId *string `pulumi:"testSuiteId"`
	// The UUID of the test suite-team relationship.
	Uuid *string `pulumi:"uuid"`
}

type TeamState struct {
	// The access level the team has on the test suite. Either `READ_ONLY` or `MANAGE_AND_READ`.
	AccessLevel pulumi.StringPtrInput
	// The GraphQL ID of the team.
	TeamId pulumi.StringPtrInput
	// The GraphQL ID of the test suite.
	TestSuiteId pulumi.StringPtrInput
	// The UUID of the test suite-team relationship.
	Uuid pulumi.StringPtrInput
}

func (TeamState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamState)(nil)).Elem()
}

type teamArgs struct {
	// The access level the team has on the test suite. Either `READ_ONLY` or `MANAGE_AND_READ`.
	AccessLevel string `pulumi:"accessLevel"`
	// The GraphQL ID of the team.
	TeamId string `pulumi:"teamId"`
	// The GraphQL ID of the test suite.
	TestSuiteId string `pulumi:"testSuiteId"`
}

// The set of arguments for constructing a Team resource.
type TeamArgs struct {
	// The access level the team has on the test suite. Either `READ_ONLY` or `MANAGE_AND_READ`.
	AccessLevel pulumi.StringInput
	// The GraphQL ID of the team.
	TeamId pulumi.StringInput
	// The GraphQL ID of the test suite.
	TestSuiteId pulumi.StringInput
}

func (TeamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamArgs)(nil)).Elem()
}

type TeamInput interface {
	pulumi.Input

	ToTeamOutput() TeamOutput
	ToTeamOutputWithContext(ctx context.Context) TeamOutput
}

func (*Team) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (i *Team) ToTeamOutput() TeamOutput {
	return i.ToTeamOutputWithContext(context.Background())
}

func (i *Team) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamOutput)
}

// TeamArrayInput is an input type that accepts TeamArray and TeamArrayOutput values.
// You can construct a concrete instance of `TeamArrayInput` via:
//
//	TeamArray{ TeamArgs{...} }
type TeamArrayInput interface {
	pulumi.Input

	ToTeamArrayOutput() TeamArrayOutput
	ToTeamArrayOutputWithContext(context.Context) TeamArrayOutput
}

type TeamArray []TeamInput

func (TeamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (i TeamArray) ToTeamArrayOutput() TeamArrayOutput {
	return i.ToTeamArrayOutputWithContext(context.Background())
}

func (i TeamArray) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamArrayOutput)
}

// TeamMapInput is an input type that accepts TeamMap and TeamMapOutput values.
// You can construct a concrete instance of `TeamMapInput` via:
//
//	TeamMap{ "key": TeamArgs{...} }
type TeamMapInput interface {
	pulumi.Input

	ToTeamMapOutput() TeamMapOutput
	ToTeamMapOutputWithContext(context.Context) TeamMapOutput
}

type TeamMap map[string]TeamInput

func (TeamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (i TeamMap) ToTeamMapOutput() TeamMapOutput {
	return i.ToTeamMapOutputWithContext(context.Background())
}

func (i TeamMap) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMapOutput)
}

type TeamOutput struct{ *pulumi.OutputState }

func (TeamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (o TeamOutput) ToTeamOutput() TeamOutput {
	return o
}

func (o TeamOutput) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return o
}

// The access level the team has on the test suite. Either `READ_ONLY` or `MANAGE_AND_READ`.
func (o TeamOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.AccessLevel }).(pulumi.StringOutput)
}

// The GraphQL ID of the team.
func (o TeamOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

// The GraphQL ID of the test suite.
func (o TeamOutput) TestSuiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.TestSuiteId }).(pulumi.StringOutput)
}

// The UUID of the test suite-team relationship.
func (o TeamOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type TeamArrayOutput struct{ *pulumi.OutputState }

func (TeamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (o TeamArrayOutput) ToTeamArrayOutput() TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) Index(i pulumi.IntInput) TeamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Team {
		return vs[0].([]*Team)[vs[1].(int)]
	}).(TeamOutput)
}

type TeamMapOutput struct{ *pulumi.OutputState }

func (TeamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (o TeamMapOutput) ToTeamMapOutput() TeamMapOutput {
	return o
}

func (o TeamMapOutput) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return o
}

func (o TeamMapOutput) MapIndex(k pulumi.StringInput) TeamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Team {
		return vs[0].(map[string]*Team)[vs[1].(string)]
	}).(TeamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamInput)(nil)).Elem(), &Team{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamArrayInput)(nil)).Elem(), TeamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMapInput)(nil)).Elem(), TeamMap{})
	pulumi.RegisterOutputType(TeamOutput{})
	pulumi.RegisterOutputType(TeamArrayOutput{})
	pulumi.RegisterOutputType(TeamMapOutput{})
}
