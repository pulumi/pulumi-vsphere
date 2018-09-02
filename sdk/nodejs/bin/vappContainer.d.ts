import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_vapp_container` resource can be used to create and manage
 * vApps.
 *
 * For more information on vSphere vApps, see [this
 * page][ref-vsphere-vapp].
 *
 * [ref-vsphere-vapp]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vm_admin.doc/GUID-2A95EBB8-1779-40FA-B4FB-4D0845750879.html
 */
export declare class VappContainer extends pulumi.CustomResource {
    /**
     * Get an existing VappContainer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VappContainerState): VappContainer;
    /**
     * Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     */
    readonly cpuExpandable: pulumi.Output<boolean | undefined>;
    /**
     * The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited.
     * Default: `-1`
     */
    readonly cpuLimit: pulumi.Output<number | undefined>;
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
     */
    readonly cpuReservation: pulumi.Output<number | undefined>;
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `cpu_shares` will be
     * ignored.  Default: `normal`
     */
    readonly cpuShareLevel: pulumi.Output<string | undefined>;
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `cpu_share_level` must be `custom`.
     */
    readonly cpuShares: pulumi.Output<number>;
    /**
     * A list of custom attributes to set on this resource.
     */
    readonly customAttributes: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     */
    readonly memoryExpandable: pulumi.Output<boolean | undefined>;
    /**
     * The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited.
     * Default: `-1`
     */
    readonly memoryLimit: pulumi.Output<number | undefined>;
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
     */
    readonly memoryReservation: pulumi.Output<number | undefined>;
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `memory_shares` will be
     * ignored.  Default: `normal`
     */
    readonly memoryShareLevel: pulumi.Output<string | undefined>;
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `memory_share_level` must be `custom`.
     */
    readonly memoryShares: pulumi.Output<number>;
    /**
     * The name of the vApp container.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the vApp container's parent folder.
     */
    readonly parentFolderId: pulumi.Output<string | undefined>;
    /**
     * The [managed object ID][docs-about-morefs]
     * of the parent resource pool. This can be the root resource pool for a cluster
     * or standalone host, or a resource pool itself. When moving a vApp container
     * from one parent resource pool to another, both must share a common root
     * resource pool or the move will fail.
     */
    readonly parentResourcePoolId: pulumi.Output<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags: pulumi.Output<string[] | undefined>;
    /**
     * Create a VappContainer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VappContainerArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering VappContainer resources.
 */
export interface VappContainerState {
    /**
     * Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     */
    readonly cpuExpandable?: pulumi.Input<boolean>;
    /**
     * The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited.
     * Default: `-1`
     */
    readonly cpuLimit?: pulumi.Input<number>;
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
     */
    readonly cpuReservation?: pulumi.Input<number>;
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `cpu_shares` will be
     * ignored.  Default: `normal`
     */
    readonly cpuShareLevel?: pulumi.Input<string>;
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `cpu_share_level` must be `custom`.
     */
    readonly cpuShares?: pulumi.Input<number>;
    /**
     * A list of custom attributes to set on this resource.
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     */
    readonly memoryExpandable?: pulumi.Input<boolean>;
    /**
     * The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited.
     * Default: `-1`
     */
    readonly memoryLimit?: pulumi.Input<number>;
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
     */
    readonly memoryReservation?: pulumi.Input<number>;
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `memory_shares` will be
     * ignored.  Default: `normal`
     */
    readonly memoryShareLevel?: pulumi.Input<string>;
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `memory_share_level` must be `custom`.
     */
    readonly memoryShares?: pulumi.Input<number>;
    /**
     * The name of the vApp container.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the vApp container's parent folder.
     */
    readonly parentFolderId?: pulumi.Input<string>;
    /**
     * The [managed object ID][docs-about-morefs]
     * of the parent resource pool. This can be the root resource pool for a cluster
     * or standalone host, or a resource pool itself. When moving a vApp container
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
 * The set of arguments for constructing a VappContainer resource.
 */
export interface VappContainerArgs {
    /**
     * Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     */
    readonly cpuExpandable?: pulumi.Input<boolean>;
    /**
     * The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited.
     * Default: `-1`
     */
    readonly cpuLimit?: pulumi.Input<number>;
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
     */
    readonly cpuReservation?: pulumi.Input<number>;
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `cpu_shares` will be
     * ignored.  Default: `normal`
     */
    readonly cpuShareLevel?: pulumi.Input<string>;
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `cpu_share_level` must be `custom`.
     */
    readonly cpuShares?: pulumi.Input<number>;
    /**
     * A list of custom attributes to set on this resource.
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     */
    readonly memoryExpandable?: pulumi.Input<boolean>;
    /**
     * The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited.
     * Default: `-1`
     */
    readonly memoryLimit?: pulumi.Input<number>;
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
     */
    readonly memoryReservation?: pulumi.Input<number>;
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `memory_shares` will be
     * ignored.  Default: `normal`
     */
    readonly memoryShareLevel?: pulumi.Input<string>;
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `memory_share_level` must be `custom`.
     */
    readonly memoryShares?: pulumi.Input<number>;
    /**
     * The name of the vApp container.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the vApp container's parent folder.
     */
    readonly parentFolderId?: pulumi.Input<string>;
    /**
     * The [managed object ID][docs-about-morefs]
     * of the parent resource pool. This can be the root resource pool for a cluster
     * or standalone host, or a resource pool itself. When moving a vApp container
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
