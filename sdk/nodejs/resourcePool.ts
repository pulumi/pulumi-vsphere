// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere..ResourcePool` resource can be used to create and manage
 * resource pools in standalone hosts or on compute clusters.
 * 
 * For more information on vSphere resource pools, see [this
 * page][ref-vsphere-resource_pools].
 * 
 * [ref-vsphere-resource_pools]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-60077B40-66FF-4625-934A-641703ED7601.html
 * 
 * ## Example Usage
 * 
 * The following example sets up a resource pool in a compute cluster which uses
 * the default settings for CPU and memory reservations, shares, and limits. The
 * compute cluster needs to already exist in vSphere.  
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const config = new pulumi.Config();
 * const cluster = config.get("cluster") || "cluster1";
 * const datacenter = config.get("datacenter") || "dc1";
 * 
 * const dc = vsphere.getDatacenter({
 *     name: datacenter,
 * });
 * const computeCluster = vsphere.getComputeCluster({
 *     datacenterId: dc.id,
 *     name: cluster,
 * });
 * const resourcePool = new vsphere.ResourcePool("resourcePool", {
 *     parentResourcePoolId: computeCluster.resourcePoolId,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/resource_pool.html.markdown.
 */
export class ResourcePool extends pulumi.CustomResource {
    /**
     * Get an existing ResourcePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourcePoolState, opts?: pulumi.CustomResourceOptions): ResourcePool {
        return new ResourcePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/resourcePool:ResourcePool';

    /**
     * Returns true if the given object is an instance of ResourcePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourcePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourcePool.__pulumiType;
    }

    /**
     * Determines if the reservation on a resource
     * pool can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     */
    public readonly cpuExpandable!: pulumi.Output<boolean | undefined>;
    /**
     * The CPU utilization of a resource pool will not exceed
     * this limit, even if there are available resources. Set to `-1` for unlimited.
     * Default: `-1`
     */
    public readonly cpuLimit!: pulumi.Output<number | undefined>;
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the resource pool. Default: `0`
     */
    public readonly cpuReservation!: pulumi.Output<number | undefined>;
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `cpuShares` will be
     * ignored.  Default: `normal`
     */
    public readonly cpuShareLevel!: pulumi.Output<string | undefined>;
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `cpuShareLevel` must be `custom`.
     */
    public readonly cpuShares!: pulumi.Output<number>;
    /**
     * A list of custom attributes to set on this resource.
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Determines if the reservation on a resource
     * pool can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     */
    public readonly memoryExpandable!: pulumi.Output<boolean | undefined>;
    /**
     * The CPU utilization of a resource pool will not exceed
     * this limit, even if there are available resources. Set to `-1` for unlimited.
     * Default: `-1`
     */
    public readonly memoryLimit!: pulumi.Output<number | undefined>;
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the resource pool. Default: `0`
     */
    public readonly memoryReservation!: pulumi.Output<number | undefined>;
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `memoryShares` will be
     * ignored.  Default: `normal`
     */
    public readonly memoryShareLevel!: pulumi.Output<string | undefined>;
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `memoryShareLevel` must be `custom`.
     */
    public readonly memoryShares!: pulumi.Output<number>;
    /**
     * The name of the resource pool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The [managed object ID][docs-about-morefs]
     * of the parent resource pool. This can be the root resource pool for a cluster
     * or standalone host, or a resource pool itself. When moving a resource pool
     * from one parent resource pool to another, both must share a common root
     * resource pool or the move will fail.
     */
    public readonly parentResourcePoolId!: pulumi.Output<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ResourcePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourcePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourcePoolArgs | ResourcePoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ResourcePoolState | undefined;
            inputs["cpuExpandable"] = state ? state.cpuExpandable : undefined;
            inputs["cpuLimit"] = state ? state.cpuLimit : undefined;
            inputs["cpuReservation"] = state ? state.cpuReservation : undefined;
            inputs["cpuShareLevel"] = state ? state.cpuShareLevel : undefined;
            inputs["cpuShares"] = state ? state.cpuShares : undefined;
            inputs["customAttributes"] = state ? state.customAttributes : undefined;
            inputs["memoryExpandable"] = state ? state.memoryExpandable : undefined;
            inputs["memoryLimit"] = state ? state.memoryLimit : undefined;
            inputs["memoryReservation"] = state ? state.memoryReservation : undefined;
            inputs["memoryShareLevel"] = state ? state.memoryShareLevel : undefined;
            inputs["memoryShares"] = state ? state.memoryShares : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parentResourcePoolId"] = state ? state.parentResourcePoolId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ResourcePoolArgs | undefined;
            if (!args || args.parentResourcePoolId === undefined) {
                throw new Error("Missing required property 'parentResourcePoolId'");
            }
            inputs["cpuExpandable"] = args ? args.cpuExpandable : undefined;
            inputs["cpuLimit"] = args ? args.cpuLimit : undefined;
            inputs["cpuReservation"] = args ? args.cpuReservation : undefined;
            inputs["cpuShareLevel"] = args ? args.cpuShareLevel : undefined;
            inputs["cpuShares"] = args ? args.cpuShares : undefined;
            inputs["customAttributes"] = args ? args.customAttributes : undefined;
            inputs["memoryExpandable"] = args ? args.memoryExpandable : undefined;
            inputs["memoryLimit"] = args ? args.memoryLimit : undefined;
            inputs["memoryReservation"] = args ? args.memoryReservation : undefined;
            inputs["memoryShareLevel"] = args ? args.memoryShareLevel : undefined;
            inputs["memoryShares"] = args ? args.memoryShares : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parentResourcePoolId"] = args ? args.parentResourcePoolId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ResourcePool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourcePool resources.
 */
export interface ResourcePoolState {
    /**
     * Determines if the reservation on a resource
     * pool can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     */
    readonly cpuExpandable?: pulumi.Input<boolean>;
    /**
     * The CPU utilization of a resource pool will not exceed
     * this limit, even if there are available resources. Set to `-1` for unlimited.
     * Default: `-1`
     */
    readonly cpuLimit?: pulumi.Input<number>;
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the resource pool. Default: `0`
     */
    readonly cpuReservation?: pulumi.Input<number>;
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `cpuShares` will be
     * ignored.  Default: `normal`
     */
    readonly cpuShareLevel?: pulumi.Input<string>;
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `cpuShareLevel` must be `custom`.
     */
    readonly cpuShares?: pulumi.Input<number>;
    /**
     * A list of custom attributes to set on this resource.
     */
    readonly customAttributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * Determines if the reservation on a resource
     * pool can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     */
    readonly memoryExpandable?: pulumi.Input<boolean>;
    /**
     * The CPU utilization of a resource pool will not exceed
     * this limit, even if there are available resources. Set to `-1` for unlimited.
     * Default: `-1`
     */
    readonly memoryLimit?: pulumi.Input<number>;
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the resource pool. Default: `0`
     */
    readonly memoryReservation?: pulumi.Input<number>;
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `memoryShares` will be
     * ignored.  Default: `normal`
     */
    readonly memoryShareLevel?: pulumi.Input<string>;
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `memoryShareLevel` must be `custom`.
     */
    readonly memoryShares?: pulumi.Input<number>;
    /**
     * The name of the resource pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The [managed object ID][docs-about-morefs]
     * of the parent resource pool. This can be the root resource pool for a cluster
     * or standalone host, or a resource pool itself. When moving a resource pool
     * from one parent resource pool to another, both must share a common root
     * resource pool or the move will fail.
     */
    readonly parentResourcePoolId?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ResourcePool resource.
 */
export interface ResourcePoolArgs {
    /**
     * Determines if the reservation on a resource
     * pool can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     */
    readonly cpuExpandable?: pulumi.Input<boolean>;
    /**
     * The CPU utilization of a resource pool will not exceed
     * this limit, even if there are available resources. Set to `-1` for unlimited.
     * Default: `-1`
     */
    readonly cpuLimit?: pulumi.Input<number>;
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the resource pool. Default: `0`
     */
    readonly cpuReservation?: pulumi.Input<number>;
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `cpuShares` will be
     * ignored.  Default: `normal`
     */
    readonly cpuShareLevel?: pulumi.Input<string>;
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `cpuShareLevel` must be `custom`.
     */
    readonly cpuShares?: pulumi.Input<number>;
    /**
     * A list of custom attributes to set on this resource.
     */
    readonly customAttributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * Determines if the reservation on a resource
     * pool can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     */
    readonly memoryExpandable?: pulumi.Input<boolean>;
    /**
     * The CPU utilization of a resource pool will not exceed
     * this limit, even if there are available resources. Set to `-1` for unlimited.
     * Default: `-1`
     */
    readonly memoryLimit?: pulumi.Input<number>;
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the resource pool. Default: `0`
     */
    readonly memoryReservation?: pulumi.Input<number>;
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `memoryShares` will be
     * ignored.  Default: `normal`
     */
    readonly memoryShareLevel?: pulumi.Input<string>;
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `memoryShareLevel` must be `custom`.
     */
    readonly memoryShares?: pulumi.Input<number>;
    /**
     * The name of the resource pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The [managed object ID][docs-about-morefs]
     * of the parent resource pool. This can be the root resource pool for a cluster
     * or standalone host, or a resource pool itself. When moving a resource pool
     * from one parent resource pool to another, both must share a common root
     * resource pool or the move will fail.
     */
    readonly parentResourcePoolId: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}
