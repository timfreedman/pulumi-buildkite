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

// A test suite is a collection of tests. A run is to a suite what a build is to a Pipeline.Use this resource to manage [Test Suites](https://buildkite.com/docs/test-analytics) on Buildkite.
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
//			_, err := TestSuite.NewTestSuite(ctx, "main", &TestSuite.TestSuiteArgs{
//				DefaultBranch: pulumi.String("main"),
//				TeamOwnerId:   pulumi.String("VGVhbvDf4eRef20tMzIxMGEfYTctNzEF5g00M8f5s6E2YjYtODNlOGNlZgD6HcBi"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type TestSuite struct {
	pulumi.CustomResourceState

	// The API token to use to send test run data to the API.
	ApiToken pulumi.StringOutput `pulumi:"apiToken"`
	// The default branch for the repository this test suite is for.
	DefaultBranch pulumi.StringOutput `pulumi:"defaultBranch"`
	// The name to give the test suite.
	Name pulumi.StringOutput `pulumi:"name"`
	// The generated slug of the test suite.
	Slug pulumi.StringOutput `pulumi:"slug"`
	// The GraphQL ID of the team to mark as the owner/admin of the test suite.
	TeamOwnerId pulumi.StringOutput `pulumi:"teamOwnerId"`
	// The UUID of the test suite.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewTestSuite registers a new resource with the given unique name, arguments, and options.
func NewTestSuite(ctx *pulumi.Context,
	name string, args *TestSuiteArgs, opts ...pulumi.ResourceOption) (*TestSuite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultBranch == nil {
		return nil, errors.New("invalid value for required argument 'DefaultBranch'")
	}
	if args.TeamOwnerId == nil {
		return nil, errors.New("invalid value for required argument 'TeamOwnerId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"apiToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TestSuite
	err := ctx.RegisterResource("buildkite:TestSuite/testSuite:TestSuite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTestSuite gets an existing TestSuite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTestSuite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TestSuiteState, opts ...pulumi.ResourceOption) (*TestSuite, error) {
	var resource TestSuite
	err := ctx.ReadResource("buildkite:TestSuite/testSuite:TestSuite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TestSuite resources.
type testSuiteState struct {
	// The API token to use to send test run data to the API.
	ApiToken *string `pulumi:"apiToken"`
	// The default branch for the repository this test suite is for.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// The name to give the test suite.
	Name *string `pulumi:"name"`
	// The generated slug of the test suite.
	Slug *string `pulumi:"slug"`
	// The GraphQL ID of the team to mark as the owner/admin of the test suite.
	TeamOwnerId *string `pulumi:"teamOwnerId"`
	// The UUID of the test suite.
	Uuid *string `pulumi:"uuid"`
}

type TestSuiteState struct {
	// The API token to use to send test run data to the API.
	ApiToken pulumi.StringPtrInput
	// The default branch for the repository this test suite is for.
	DefaultBranch pulumi.StringPtrInput
	// The name to give the test suite.
	Name pulumi.StringPtrInput
	// The generated slug of the test suite.
	Slug pulumi.StringPtrInput
	// The GraphQL ID of the team to mark as the owner/admin of the test suite.
	TeamOwnerId pulumi.StringPtrInput
	// The UUID of the test suite.
	Uuid pulumi.StringPtrInput
}

func (TestSuiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*testSuiteState)(nil)).Elem()
}

type testSuiteArgs struct {
	// The default branch for the repository this test suite is for.
	DefaultBranch string `pulumi:"defaultBranch"`
	// The name to give the test suite.
	Name *string `pulumi:"name"`
	// The GraphQL ID of the team to mark as the owner/admin of the test suite.
	TeamOwnerId string `pulumi:"teamOwnerId"`
}

// The set of arguments for constructing a TestSuite resource.
type TestSuiteArgs struct {
	// The default branch for the repository this test suite is for.
	DefaultBranch pulumi.StringInput
	// The name to give the test suite.
	Name pulumi.StringPtrInput
	// The GraphQL ID of the team to mark as the owner/admin of the test suite.
	TeamOwnerId pulumi.StringInput
}

func (TestSuiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*testSuiteArgs)(nil)).Elem()
}

type TestSuiteInput interface {
	pulumi.Input

	ToTestSuiteOutput() TestSuiteOutput
	ToTestSuiteOutputWithContext(ctx context.Context) TestSuiteOutput
}

func (*TestSuite) ElementType() reflect.Type {
	return reflect.TypeOf((**TestSuite)(nil)).Elem()
}

func (i *TestSuite) ToTestSuiteOutput() TestSuiteOutput {
	return i.ToTestSuiteOutputWithContext(context.Background())
}

func (i *TestSuite) ToTestSuiteOutputWithContext(ctx context.Context) TestSuiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestSuiteOutput)
}

// TestSuiteArrayInput is an input type that accepts TestSuiteArray and TestSuiteArrayOutput values.
// You can construct a concrete instance of `TestSuiteArrayInput` via:
//
//	TestSuiteArray{ TestSuiteArgs{...} }
type TestSuiteArrayInput interface {
	pulumi.Input

	ToTestSuiteArrayOutput() TestSuiteArrayOutput
	ToTestSuiteArrayOutputWithContext(context.Context) TestSuiteArrayOutput
}

type TestSuiteArray []TestSuiteInput

func (TestSuiteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TestSuite)(nil)).Elem()
}

func (i TestSuiteArray) ToTestSuiteArrayOutput() TestSuiteArrayOutput {
	return i.ToTestSuiteArrayOutputWithContext(context.Background())
}

func (i TestSuiteArray) ToTestSuiteArrayOutputWithContext(ctx context.Context) TestSuiteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestSuiteArrayOutput)
}

// TestSuiteMapInput is an input type that accepts TestSuiteMap and TestSuiteMapOutput values.
// You can construct a concrete instance of `TestSuiteMapInput` via:
//
//	TestSuiteMap{ "key": TestSuiteArgs{...} }
type TestSuiteMapInput interface {
	pulumi.Input

	ToTestSuiteMapOutput() TestSuiteMapOutput
	ToTestSuiteMapOutputWithContext(context.Context) TestSuiteMapOutput
}

type TestSuiteMap map[string]TestSuiteInput

func (TestSuiteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TestSuite)(nil)).Elem()
}

func (i TestSuiteMap) ToTestSuiteMapOutput() TestSuiteMapOutput {
	return i.ToTestSuiteMapOutputWithContext(context.Background())
}

func (i TestSuiteMap) ToTestSuiteMapOutputWithContext(ctx context.Context) TestSuiteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestSuiteMapOutput)
}

type TestSuiteOutput struct{ *pulumi.OutputState }

func (TestSuiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TestSuite)(nil)).Elem()
}

func (o TestSuiteOutput) ToTestSuiteOutput() TestSuiteOutput {
	return o
}

func (o TestSuiteOutput) ToTestSuiteOutputWithContext(ctx context.Context) TestSuiteOutput {
	return o
}

// The API token to use to send test run data to the API.
func (o TestSuiteOutput) ApiToken() pulumi.StringOutput {
	return o.ApplyT(func(v *TestSuite) pulumi.StringOutput { return v.ApiToken }).(pulumi.StringOutput)
}

// The default branch for the repository this test suite is for.
func (o TestSuiteOutput) DefaultBranch() pulumi.StringOutput {
	return o.ApplyT(func(v *TestSuite) pulumi.StringOutput { return v.DefaultBranch }).(pulumi.StringOutput)
}

// The name to give the test suite.
func (o TestSuiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TestSuite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The generated slug of the test suite.
func (o TestSuiteOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *TestSuite) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// The GraphQL ID of the team to mark as the owner/admin of the test suite.
func (o TestSuiteOutput) TeamOwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *TestSuite) pulumi.StringOutput { return v.TeamOwnerId }).(pulumi.StringOutput)
}

// The UUID of the test suite.
func (o TestSuiteOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *TestSuite) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type TestSuiteArrayOutput struct{ *pulumi.OutputState }

func (TestSuiteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TestSuite)(nil)).Elem()
}

func (o TestSuiteArrayOutput) ToTestSuiteArrayOutput() TestSuiteArrayOutput {
	return o
}

func (o TestSuiteArrayOutput) ToTestSuiteArrayOutputWithContext(ctx context.Context) TestSuiteArrayOutput {
	return o
}

func (o TestSuiteArrayOutput) Index(i pulumi.IntInput) TestSuiteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TestSuite {
		return vs[0].([]*TestSuite)[vs[1].(int)]
	}).(TestSuiteOutput)
}

type TestSuiteMapOutput struct{ *pulumi.OutputState }

func (TestSuiteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TestSuite)(nil)).Elem()
}

func (o TestSuiteMapOutput) ToTestSuiteMapOutput() TestSuiteMapOutput {
	return o
}

func (o TestSuiteMapOutput) ToTestSuiteMapOutputWithContext(ctx context.Context) TestSuiteMapOutput {
	return o
}

func (o TestSuiteMapOutput) MapIndex(k pulumi.StringInput) TestSuiteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TestSuite {
		return vs[0].(map[string]*TestSuite)[vs[1].(string)]
	}).(TestSuiteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TestSuiteInput)(nil)).Elem(), &TestSuite{})
	pulumi.RegisterInputType(reflect.TypeOf((*TestSuiteArrayInput)(nil)).Elem(), TestSuiteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TestSuiteMapInput)(nil)).Elem(), TestSuiteMap{})
	pulumi.RegisterOutputType(TestSuiteOutput{})
	pulumi.RegisterOutputType(TestSuiteArrayOutput{})
	pulumi.RegisterOutputType(TestSuiteMapOutput{})
}
