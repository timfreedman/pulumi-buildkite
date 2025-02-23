// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package buildkite

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/internal"
)

var _ = internal.GetEnvOrDefault

type ProviderTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	Create *string `pulumi:"create"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
	Delete *string `pulumi:"delete"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
	Read *string `pulumi:"read"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	Update *string `pulumi:"update"`
}

// ProviderTimeoutsInput is an input type that accepts ProviderTimeoutsArgs and ProviderTimeoutsOutput values.
// You can construct a concrete instance of `ProviderTimeoutsInput` via:
//
//	ProviderTimeoutsArgs{...}
type ProviderTimeoutsInput interface {
	pulumi.Input

	ToProviderTimeoutsOutput() ProviderTimeoutsOutput
	ToProviderTimeoutsOutputWithContext(context.Context) ProviderTimeoutsOutput
}

type ProviderTimeoutsArgs struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	Create pulumi.StringPtrInput `pulumi:"create"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
	Delete pulumi.StringPtrInput `pulumi:"delete"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
	Read pulumi.StringPtrInput `pulumi:"read"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	Update pulumi.StringPtrInput `pulumi:"update"`
}

func (ProviderTimeoutsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderTimeouts)(nil)).Elem()
}

func (i ProviderTimeoutsArgs) ToProviderTimeoutsOutput() ProviderTimeoutsOutput {
	return i.ToProviderTimeoutsOutputWithContext(context.Background())
}

func (i ProviderTimeoutsArgs) ToProviderTimeoutsOutputWithContext(ctx context.Context) ProviderTimeoutsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderTimeoutsOutput)
}

func (i ProviderTimeoutsArgs) ToProviderTimeoutsPtrOutput() ProviderTimeoutsPtrOutput {
	return i.ToProviderTimeoutsPtrOutputWithContext(context.Background())
}

func (i ProviderTimeoutsArgs) ToProviderTimeoutsPtrOutputWithContext(ctx context.Context) ProviderTimeoutsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderTimeoutsOutput).ToProviderTimeoutsPtrOutputWithContext(ctx)
}

// ProviderTimeoutsPtrInput is an input type that accepts ProviderTimeoutsArgs, ProviderTimeoutsPtr and ProviderTimeoutsPtrOutput values.
// You can construct a concrete instance of `ProviderTimeoutsPtrInput` via:
//
//	        ProviderTimeoutsArgs{...}
//
//	or:
//
//	        nil
type ProviderTimeoutsPtrInput interface {
	pulumi.Input

	ToProviderTimeoutsPtrOutput() ProviderTimeoutsPtrOutput
	ToProviderTimeoutsPtrOutputWithContext(context.Context) ProviderTimeoutsPtrOutput
}

type providerTimeoutsPtrType ProviderTimeoutsArgs

func ProviderTimeoutsPtr(v *ProviderTimeoutsArgs) ProviderTimeoutsPtrInput {
	return (*providerTimeoutsPtrType)(v)
}

func (*providerTimeoutsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderTimeouts)(nil)).Elem()
}

func (i *providerTimeoutsPtrType) ToProviderTimeoutsPtrOutput() ProviderTimeoutsPtrOutput {
	return i.ToProviderTimeoutsPtrOutputWithContext(context.Background())
}

func (i *providerTimeoutsPtrType) ToProviderTimeoutsPtrOutputWithContext(ctx context.Context) ProviderTimeoutsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderTimeoutsPtrOutput)
}

type ProviderTimeoutsOutput struct{ *pulumi.OutputState }

func (ProviderTimeoutsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderTimeouts)(nil)).Elem()
}

func (o ProviderTimeoutsOutput) ToProviderTimeoutsOutput() ProviderTimeoutsOutput {
	return o
}

func (o ProviderTimeoutsOutput) ToProviderTimeoutsOutputWithContext(ctx context.Context) ProviderTimeoutsOutput {
	return o
}

func (o ProviderTimeoutsOutput) ToProviderTimeoutsPtrOutput() ProviderTimeoutsPtrOutput {
	return o.ToProviderTimeoutsPtrOutputWithContext(context.Background())
}

func (o ProviderTimeoutsOutput) ToProviderTimeoutsPtrOutputWithContext(ctx context.Context) ProviderTimeoutsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProviderTimeouts) *ProviderTimeouts {
		return &v
	}).(ProviderTimeoutsPtrOutput)
}

// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
func (o ProviderTimeoutsOutput) Create() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderTimeouts) *string { return v.Create }).(pulumi.StringPtrOutput)
}

// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
func (o ProviderTimeoutsOutput) Delete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderTimeouts) *string { return v.Delete }).(pulumi.StringPtrOutput)
}

// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
func (o ProviderTimeoutsOutput) Read() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderTimeouts) *string { return v.Read }).(pulumi.StringPtrOutput)
}

// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
func (o ProviderTimeoutsOutput) Update() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderTimeouts) *string { return v.Update }).(pulumi.StringPtrOutput)
}

type ProviderTimeoutsPtrOutput struct{ *pulumi.OutputState }

func (ProviderTimeoutsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderTimeouts)(nil)).Elem()
}

func (o ProviderTimeoutsPtrOutput) ToProviderTimeoutsPtrOutput() ProviderTimeoutsPtrOutput {
	return o
}

func (o ProviderTimeoutsPtrOutput) ToProviderTimeoutsPtrOutputWithContext(ctx context.Context) ProviderTimeoutsPtrOutput {
	return o
}

func (o ProviderTimeoutsPtrOutput) Elem() ProviderTimeoutsOutput {
	return o.ApplyT(func(v *ProviderTimeouts) ProviderTimeouts {
		if v != nil {
			return *v
		}
		var ret ProviderTimeouts
		return ret
	}).(ProviderTimeoutsOutput)
}

// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
func (o ProviderTimeoutsPtrOutput) Create() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Create
	}).(pulumi.StringPtrOutput)
}

// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
func (o ProviderTimeoutsPtrOutput) Delete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(pulumi.StringPtrOutput)
}

// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
func (o ProviderTimeoutsPtrOutput) Read() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Read
	}).(pulumi.StringPtrOutput)
}

// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
func (o ProviderTimeoutsPtrOutput) Update() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Update
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderTimeoutsInput)(nil)).Elem(), ProviderTimeoutsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderTimeoutsPtrInput)(nil)).Elem(), ProviderTimeoutsArgs{})
	pulumi.RegisterOutputType(ProviderTimeoutsOutput{})
	pulumi.RegisterOutputType(ProviderTimeoutsPtrOutput{})
}
