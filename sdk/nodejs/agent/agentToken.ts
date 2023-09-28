// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource allows you to create and manage non-clustered agent tokens.
 * You can find out more about clusters in the Buildkite [documentation](https://buildkite.com/docs/agent/v3/tokens).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as buildkite from "@pulumiverse/buildkite";
 *
 * // create a default token
 * const _default = new buildkite.agent.AgentToken("default", {description: "Default token"});
 * ```
 */
export class AgentToken extends pulumi.CustomResource {
    /**
     * Get an existing AgentToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AgentTokenState, opts?: pulumi.CustomResourceOptions): AgentToken {
        return new AgentToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'buildkite:Agent/agentToken:AgentToken';

    /**
     * Returns true if the given object is an instance of AgentToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AgentToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AgentToken.__pulumiType;
    }

    /**
     * The description of the agent token. Used to help identify its use.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The token value used by an agent to register with the API.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * The UUID of the agent token.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;

    /**
     * Create a AgentToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AgentTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AgentTokenArgs | AgentTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AgentTokenState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as AgentTokenArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AgentToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AgentToken resources.
 */
export interface AgentTokenState {
    /**
     * The description of the agent token. Used to help identify its use.
     */
    description?: pulumi.Input<string>;
    /**
     * The token value used by an agent to register with the API.
     */
    token?: pulumi.Input<string>;
    /**
     * The UUID of the agent token.
     */
    uuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AgentToken resource.
 */
export interface AgentTokenArgs {
    /**
     * The description of the agent token. Used to help identify its use.
     */
    description?: pulumi.Input<string>;
}
