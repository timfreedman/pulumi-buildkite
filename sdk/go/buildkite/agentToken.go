// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package buildkite

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: agentToken
//
// This resource allows you to create and manage agent tokens.
//
// Buildkite Documentation: https://buildkite.com/docs/agent/v3/tokens
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/grapl-security/pulumi-buildkite/sdk/go/buildkite"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := buildkite.NewAgentToken(ctx, "fleet", &buildkite.AgentTokenArgs{
// 			Description: pulumi.String("token used by build fleet"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Tokens can be imported using the `GraphQL ID` (not UUID), e.g.
//
// ```sh
//  $ pulumi import buildkite:index/agentToken:AgentToken fleet QWdlbnRUb2tlbi0tLTQzNWNhZDU4LWU4MWQtNDVhZi04NjM3LWIxY2Y4MDcwMjM4ZA==
// ```
type AgentToken struct {
	pulumi.CustomResourceState

	// This is the description of the agent token.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The value of the created agent token.
	Token pulumi.StringOutput `pulumi:"token"`
	// The UUID of the token.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewAgentToken registers a new resource with the given unique name, arguments, and options.
func NewAgentToken(ctx *pulumi.Context,
	name string, args *AgentTokenArgs, opts ...pulumi.ResourceOption) (*AgentToken, error) {
	if args == nil {
		args = &AgentTokenArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource AgentToken
	err := ctx.RegisterResource("buildkite:index/agentToken:AgentToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentToken gets an existing AgentToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentTokenState, opts ...pulumi.ResourceOption) (*AgentToken, error) {
	var resource AgentToken
	err := ctx.ReadResource("buildkite:index/agentToken:AgentToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentToken resources.
type agentTokenState struct {
	// This is the description of the agent token.
	Description *string `pulumi:"description"`
	// The value of the created agent token.
	Token *string `pulumi:"token"`
	// The UUID of the token.
	Uuid *string `pulumi:"uuid"`
}

type AgentTokenState struct {
	// This is the description of the agent token.
	Description pulumi.StringPtrInput
	// The value of the created agent token.
	Token pulumi.StringPtrInput
	// The UUID of the token.
	Uuid pulumi.StringPtrInput
}

func (AgentTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentTokenState)(nil)).Elem()
}

type agentTokenArgs struct {
	// This is the description of the agent token.
	Description *string `pulumi:"description"`
}

// The set of arguments for constructing a AgentToken resource.
type AgentTokenArgs struct {
	// This is the description of the agent token.
	Description pulumi.StringPtrInput
}

func (AgentTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentTokenArgs)(nil)).Elem()
}

type AgentTokenInput interface {
	pulumi.Input

	ToAgentTokenOutput() AgentTokenOutput
	ToAgentTokenOutputWithContext(ctx context.Context) AgentTokenOutput
}

func (*AgentToken) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentToken)(nil)).Elem()
}

func (i *AgentToken) ToAgentTokenOutput() AgentTokenOutput {
	return i.ToAgentTokenOutputWithContext(context.Background())
}

func (i *AgentToken) ToAgentTokenOutputWithContext(ctx context.Context) AgentTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentTokenOutput)
}

// AgentTokenArrayInput is an input type that accepts AgentTokenArray and AgentTokenArrayOutput values.
// You can construct a concrete instance of `AgentTokenArrayInput` via:
//
//          AgentTokenArray{ AgentTokenArgs{...} }
type AgentTokenArrayInput interface {
	pulumi.Input

	ToAgentTokenArrayOutput() AgentTokenArrayOutput
	ToAgentTokenArrayOutputWithContext(context.Context) AgentTokenArrayOutput
}

type AgentTokenArray []AgentTokenInput

func (AgentTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AgentToken)(nil)).Elem()
}

func (i AgentTokenArray) ToAgentTokenArrayOutput() AgentTokenArrayOutput {
	return i.ToAgentTokenArrayOutputWithContext(context.Background())
}

func (i AgentTokenArray) ToAgentTokenArrayOutputWithContext(ctx context.Context) AgentTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentTokenArrayOutput)
}

// AgentTokenMapInput is an input type that accepts AgentTokenMap and AgentTokenMapOutput values.
// You can construct a concrete instance of `AgentTokenMapInput` via:
//
//          AgentTokenMap{ "key": AgentTokenArgs{...} }
type AgentTokenMapInput interface {
	pulumi.Input

	ToAgentTokenMapOutput() AgentTokenMapOutput
	ToAgentTokenMapOutputWithContext(context.Context) AgentTokenMapOutput
}

type AgentTokenMap map[string]AgentTokenInput

func (AgentTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AgentToken)(nil)).Elem()
}

func (i AgentTokenMap) ToAgentTokenMapOutput() AgentTokenMapOutput {
	return i.ToAgentTokenMapOutputWithContext(context.Background())
}

func (i AgentTokenMap) ToAgentTokenMapOutputWithContext(ctx context.Context) AgentTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentTokenMapOutput)
}

type AgentTokenOutput struct{ *pulumi.OutputState }

func (AgentTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentToken)(nil)).Elem()
}

func (o AgentTokenOutput) ToAgentTokenOutput() AgentTokenOutput {
	return o
}

func (o AgentTokenOutput) ToAgentTokenOutputWithContext(ctx context.Context) AgentTokenOutput {
	return o
}

// This is the description of the agent token.
func (o AgentTokenOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentToken) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The value of the created agent token.
func (o AgentTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// The UUID of the token.
func (o AgentTokenOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentToken) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type AgentTokenArrayOutput struct{ *pulumi.OutputState }

func (AgentTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AgentToken)(nil)).Elem()
}

func (o AgentTokenArrayOutput) ToAgentTokenArrayOutput() AgentTokenArrayOutput {
	return o
}

func (o AgentTokenArrayOutput) ToAgentTokenArrayOutputWithContext(ctx context.Context) AgentTokenArrayOutput {
	return o
}

func (o AgentTokenArrayOutput) Index(i pulumi.IntInput) AgentTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AgentToken {
		return vs[0].([]*AgentToken)[vs[1].(int)]
	}).(AgentTokenOutput)
}

type AgentTokenMapOutput struct{ *pulumi.OutputState }

func (AgentTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AgentToken)(nil)).Elem()
}

func (o AgentTokenMapOutput) ToAgentTokenMapOutput() AgentTokenMapOutput {
	return o
}

func (o AgentTokenMapOutput) ToAgentTokenMapOutputWithContext(ctx context.Context) AgentTokenMapOutput {
	return o
}

func (o AgentTokenMapOutput) MapIndex(k pulumi.StringInput) AgentTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AgentToken {
		return vs[0].(map[string]*AgentToken)[vs[1].(string)]
	}).(AgentTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AgentTokenInput)(nil)).Elem(), &AgentToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*AgentTokenArrayInput)(nil)).Elem(), AgentTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AgentTokenMapInput)(nil)).Elem(), AgentTokenMap{})
	pulumi.RegisterOutputType(AgentTokenOutput{})
	pulumi.RegisterOutputType(AgentTokenArrayOutput{})
	pulumi.RegisterOutputType(AgentTokenMapOutput{})
}
