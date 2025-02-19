// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/internal"
)

// Use this data source to look up the organization settings.
func LookupOrganization(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupOrganizationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrganizationResult
	err := ctx.Invoke("buildkite:Organization/getOrganization:getOrganization", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getOrganization.
type LookupOrganizationResult struct {
	// List of IP addresses in CIDR format that are allowed to access the Buildkite API for this organization.
	AllowedApiIpAddresses []string `pulumi:"allowedApiIpAddresses"`
	// The GraphQL ID of the organization.
	Id string `pulumi:"id"`
	// The UUID of the organization.
	Uuid string `pulumi:"uuid"`
}

func LookupOrganizationOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) LookupOrganizationResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (LookupOrganizationResult, error) {
		r, err := LookupOrganization(ctx, opts...)
		var s LookupOrganizationResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(LookupOrganizationResultOutput)
}

// A collection of values returned by getOrganization.
type LookupOrganizationResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationResult)(nil)).Elem()
}

func (o LookupOrganizationResultOutput) ToLookupOrganizationResultOutput() LookupOrganizationResultOutput {
	return o
}

func (o LookupOrganizationResultOutput) ToLookupOrganizationResultOutputWithContext(ctx context.Context) LookupOrganizationResultOutput {
	return o
}

// List of IP addresses in CIDR format that are allowed to access the Buildkite API for this organization.
func (o LookupOrganizationResultOutput) AllowedApiIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrganizationResult) []string { return v.AllowedApiIpAddresses }).(pulumi.StringArrayOutput)
}

// The GraphQL ID of the organization.
func (o LookupOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The UUID of the organization.
func (o LookupOrganizationResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationResultOutput{})
}
