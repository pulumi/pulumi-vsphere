// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ComputeClusterVmHostRule extends pulumi.CustomResource {
    /**
     * Get an existing ComputeClusterVmHostRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeClusterVmHostRuleState, opts?: pulumi.CustomResourceOptions): ComputeClusterVmHostRule {
        return new ComputeClusterVmHostRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/computeClusterVmHostRule:ComputeClusterVmHostRule';

    /**
     * Returns true if the given object is an instance of ComputeClusterVmHostRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeClusterVmHostRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeClusterVmHostRule.__pulumiType;
    }

    /**
     * When this field is used, the virtual
     * machines defined in `vmGroupName` will be run on the
     * hosts defined in this host group.
     */
    public readonly affinityHostGroupName!: pulumi.Output<string | undefined>;
    /**
     * When this field is used, the
     * virtual machines defined in `vmGroupName` will _not_ be
     * run on the hosts defined in this host group.
     */
    public readonly antiAffinityHostGroupName!: pulumi.Output<string | undefined>;
    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    public readonly computeClusterId!: pulumi.Output<string>;
    /**
     * Enable this rule in the cluster. Default: `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
     *
     * > **NOTE:** One of `affinityHostGroupName` or
     * `antiAffinityHostGroupName` must be
     * defined, but not both.
     *
     * > **NOTE:** The namespace for rule names on this resource (defined by the
     * `name` argument) is shared with all rules in the cluster - consider
     * this when naming your rules.
     */
    public readonly mandatory!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the rule. This must be unique in the
     * cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the virtual machine group to use
     * with this rule.
     */
    public readonly vmGroupName!: pulumi.Output<string>;

    /**
     * Create a ComputeClusterVmHostRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeClusterVmHostRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeClusterVmHostRuleArgs | ComputeClusterVmHostRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComputeClusterVmHostRuleState | undefined;
            resourceInputs["affinityHostGroupName"] = state ? state.affinityHostGroupName : undefined;
            resourceInputs["antiAffinityHostGroupName"] = state ? state.antiAffinityHostGroupName : undefined;
            resourceInputs["computeClusterId"] = state ? state.computeClusterId : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["mandatory"] = state ? state.mandatory : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vmGroupName"] = state ? state.vmGroupName : undefined;
        } else {
            const args = argsOrState as ComputeClusterVmHostRuleArgs | undefined;
            if ((!args || args.computeClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'computeClusterId'");
            }
            if ((!args || args.vmGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vmGroupName'");
            }
            resourceInputs["affinityHostGroupName"] = args ? args.affinityHostGroupName : undefined;
            resourceInputs["antiAffinityHostGroupName"] = args ? args.antiAffinityHostGroupName : undefined;
            resourceInputs["computeClusterId"] = args ? args.computeClusterId : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["mandatory"] = args ? args.mandatory : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vmGroupName"] = args ? args.vmGroupName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ComputeClusterVmHostRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeClusterVmHostRule resources.
 */
export interface ComputeClusterVmHostRuleState {
    /**
     * When this field is used, the virtual
     * machines defined in `vmGroupName` will be run on the
     * hosts defined in this host group.
     */
    affinityHostGroupName?: pulumi.Input<string>;
    /**
     * When this field is used, the
     * virtual machines defined in `vmGroupName` will _not_ be
     * run on the hosts defined in this host group.
     */
    antiAffinityHostGroupName?: pulumi.Input<string>;
    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    computeClusterId?: pulumi.Input<string>;
    /**
     * Enable this rule in the cluster. Default: `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
     *
     * > **NOTE:** One of `affinityHostGroupName` or
     * `antiAffinityHostGroupName` must be
     * defined, but not both.
     *
     * > **NOTE:** The namespace for rule names on this resource (defined by the
     * `name` argument) is shared with all rules in the cluster - consider
     * this when naming your rules.
     */
    mandatory?: pulumi.Input<boolean>;
    /**
     * The name of the rule. This must be unique in the
     * cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the virtual machine group to use
     * with this rule.
     */
    vmGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComputeClusterVmHostRule resource.
 */
export interface ComputeClusterVmHostRuleArgs {
    /**
     * When this field is used, the virtual
     * machines defined in `vmGroupName` will be run on the
     * hosts defined in this host group.
     */
    affinityHostGroupName?: pulumi.Input<string>;
    /**
     * When this field is used, the
     * virtual machines defined in `vmGroupName` will _not_ be
     * run on the hosts defined in this host group.
     */
    antiAffinityHostGroupName?: pulumi.Input<string>;
    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    computeClusterId: pulumi.Input<string>;
    /**
     * Enable this rule in the cluster. Default: `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
     *
     * > **NOTE:** One of `affinityHostGroupName` or
     * `antiAffinityHostGroupName` must be
     * defined, but not both.
     *
     * > **NOTE:** The namespace for rule names on this resource (defined by the
     * `name` argument) is shared with all rules in the cluster - consider
     * this when naming your rules.
     */
    mandatory?: pulumi.Input<boolean>;
    /**
     * The name of the rule. This must be unique in the
     * cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the virtual machine group to use
     * with this rule.
     */
    vmGroupName: pulumi.Input<string>;
}
