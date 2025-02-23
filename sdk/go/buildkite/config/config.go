// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/internal"
)

var _ = internal.GetEnvOrDefault

func GetApiToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "buildkite:apiToken")
}

// Enable this to archive pipelines when destroying the resource. This is opposed to completely deleting pipelines.
func GetArchivePipelineOnDelete(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "buildkite:archivePipelineOnDelete")
}

// Base URL for the GraphQL API to use. If not provided, the value is taken from the `BUILDKITE_GRAPHQL_URL` environment
// variable.
func GetGraphqlUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "buildkite:graphqlUrl")
}

// The Buildkite organization slug. This can be found on the [settings](https://buildkite.com/organizations/~/settings)
// page. If not provided, the value is taken from the `BUILDKITE_ORGANIZATION_SLUG` environment variable.
func GetOrganization(ctx *pulumi.Context) string {
	return config.Get(ctx, "buildkite:organization")
}

// Base URL for the REST API to use. If not provided, the value is taken from the `BUILDKITE_REST_URL` environment
// variable.
func GetRestUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "buildkite:restUrl")
}
func GetTimeouts(ctx *pulumi.Context) string {
	return config.Get(ctx, "buildkite:timeouts")
}
