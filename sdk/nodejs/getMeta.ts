// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to look up the source IP addresses that Buildkite may use to send external requests,
 * including webhooks and API calls to source control systems (like GitHub Enterprise Server and BitBucket Server).
 *
 * More info in the Buildkite [documentation](https://buildkite.com/docs/apis/rest-api/meta).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as buildkite from "@pulumi/buildkite";
 *
 * const ips = buildkite.getMeta({});
 * // Create an AWS security group allowing ingress from Buildkite
 * const fromBuildkite = new aws.ec2.SecurityGroup("fromBuildkite", {ingress: [{
 *     fromPort: "*",
 *     toPort: 443,
 *     protocol: "tcp",
 *     cidrBlocks: ips.then(ips => ips.webhookIps),
 * }]});
 * ```
 */
export function getMeta(opts?: pulumi.InvokeOptions): Promise<GetMetaResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("buildkite:index/getMeta:getMeta", {
    }, opts);
}

/**
 * A collection of values returned by getMeta.
 */
export interface GetMetaResult {
    /**
     * Fixed value: `https://api.buildkite.com/v2/meta`.
     */
    readonly id: string;
    /**
     * List of IPs in CIDR format.
     */
    readonly webhookIps: string[];
}
/**
 * Use this data source to look up the source IP addresses that Buildkite may use to send external requests,
 * including webhooks and API calls to source control systems (like GitHub Enterprise Server and BitBucket Server).
 *
 * More info in the Buildkite [documentation](https://buildkite.com/docs/apis/rest-api/meta).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as buildkite from "@pulumi/buildkite";
 *
 * const ips = buildkite.getMeta({});
 * // Create an AWS security group allowing ingress from Buildkite
 * const fromBuildkite = new aws.ec2.SecurityGroup("fromBuildkite", {ingress: [{
 *     fromPort: "*",
 *     toPort: 443,
 *     protocol: "tcp",
 *     cidrBlocks: ips.then(ips => ips.webhookIps),
 * }]});
 * ```
 */
export function getMetaOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetMetaResult> {
    return pulumi.output(getMeta(opts))
}
