// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere..ComputeClusterHostGroup` resource can be used to manage groups
 * of hosts in a cluster, either created by the
 * `vsphere..ComputeCluster` resource or looked up
 * by the `vsphere..ComputeCluster` data source.
 * 
 * 
 * This resource mainly serves as an input to the
 * `vsphere..ComputeClusterVmHostRule`
 * resource - see the documentation for that resource for further details on how
 * to use host groups.
 * 
 * > **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 * 
 * > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const config = new pulumi.Config();
 * const datacenter = config.get("datacenter") || "dc1";
 * const hosts = config.get("hosts") || [
 *     "esxi1",
 *     "esxi2",
 *     "esxi3",
 * ];
 * 
 * const dc = pulumi.output(vsphere.getDatacenter({
 *     name: datacenter,
 * }, { async: true }));
 * const hostsHost: pulumi.Output<vsphere.GetHostResult>[] = [];
 * for (let i = 0; i < hosts.length; i++) {
 *     hostsHost.push(dc.apply(dc => vsphere.getHost({
 *         datacenterId: dc.id,
 *         name: hosts[i],
 *     }, { async: true })));
 * }
 * const computeCluster = new vsphere.ComputeCluster("computeCluster", {
 *     datacenterId: dc.id,
 *     drsAutomationLevel: "fullyAutomated",
 *     drsEnabled: true,
 *     haEnabled: true,
 *     hostSystemIds: hostsHost.map(v => v.id),
 * });
 * const clusterHostGroup = new vsphere.ComputeClusterHostGroup("clusterHostGroup", {
 *     computeClusterId: computeCluster.id,
 *     hostSystemIds: hostsHost.map(v => v.id),
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/compute_cluster_host_group.html.markdown.
 */
export class ComputeClusterHostGroup extends pulumi.CustomResource {
    /**
     * Get an existing ComputeClusterHostGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeClusterHostGroupState, opts?: pulumi.CustomResourceOptions): ComputeClusterHostGroup {
        return new ComputeClusterHostGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup';

    /**
     * Returns true if the given object is an instance of ComputeClusterHostGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeClusterHostGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeClusterHostGroup.__pulumiType;
    }

    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    public readonly computeClusterId!: pulumi.Output<string>;
    /**
     * The managed object IDs of
     * the hosts to put in the cluster.
     */
    public readonly hostSystemIds!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the host group. This must be unique in the
     * cluster. Forces a new resource if changed.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ComputeClusterHostGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeClusterHostGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeClusterHostGroupArgs | ComputeClusterHostGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ComputeClusterHostGroupState | undefined;
            inputs["computeClusterId"] = state ? state.computeClusterId : undefined;
            inputs["hostSystemIds"] = state ? state.hostSystemIds : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ComputeClusterHostGroupArgs | undefined;
            if (!args || args.computeClusterId === undefined) {
                throw new Error("Missing required property 'computeClusterId'");
            }
            inputs["computeClusterId"] = args ? args.computeClusterId : undefined;
            inputs["hostSystemIds"] = args ? args.hostSystemIds : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ComputeClusterHostGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeClusterHostGroup resources.
 */
export interface ComputeClusterHostGroupState {
    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId?: pulumi.Input<string>;
    /**
     * The managed object IDs of
     * the hosts to put in the cluster.
     */
    readonly hostSystemIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the host group. This must be unique in the
     * cluster. Forces a new resource if changed.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComputeClusterHostGroup resource.
 */
export interface ComputeClusterHostGroupArgs {
    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Input<string>;
    /**
     * The managed object IDs of
     * the hosts to put in the cluster.
     */
    readonly hostSystemIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the host group. This must be unique in the
     * cluster. Forces a new resource if changed.
     */
    readonly name?: pulumi.Input<string>;
}
