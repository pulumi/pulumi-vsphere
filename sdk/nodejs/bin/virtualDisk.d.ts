import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_virtual_disk` resource can be used to create virtual disks outside
 * of any given [`vsphere_virtual_machine`][docs-vsphere-virtual-machine]
 * resource. These disks can be attached to a virtual machine by creating a disk
 * block with the [`attach`][docs-vsphere-virtual-machine-disk-attach] parameter.
 *
 * [docs-vsphere-virtual-machine]: /docs/providers/vsphere/r/virtual_machine.html
 * [docs-vsphere-virtual-machine-disk-attach]: /docs/providers/vsphere/r/virtual_machine.html#attach
 */
export declare class VirtualDisk extends pulumi.CustomResource {
    /**
     * Get an existing VirtualDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualDiskState): VirtualDisk;
    /**
     * The adapter type for this virtual disk. Can be
     * one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
     */
    readonly adapterType: pulumi.Output<string | undefined>;
    /**
     * Tells the resource to create any
     * directories that are a part of the `vmdk_path` parameter if they are missing.
     * Default: `false`.
     */
    readonly createDirectories: pulumi.Output<boolean | undefined>;
    /**
     * The name of the datacenter in which to create the
     * disk. Can be omitted when when ESXi or if there is only one datacenter in
     * your infrastructure.
     */
    readonly datacenter: pulumi.Output<string | undefined>;
    /**
     * The name of the datastore in which to create the
     * disk.
     */
    readonly datastore: pulumi.Output<string>;
    /**
     * Size of the disk (in GB).
     */
    readonly size: pulumi.Output<number>;
    /**
     * The type of disk to create. Can be one of
     * `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
     * information on what each kind of disk provisioning policy means, click
     * [here][docs-vmware-vm-disk-provisioning].
     */
    readonly type: pulumi.Output<string | undefined>;
    /**
     * The path, including filename, of the virtual disk to
     * be created.  This needs to end in `.vmdk`.
     */
    readonly vmdkPath: pulumi.Output<string>;
    /**
     * Create a VirtualDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualDiskArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering VirtualDisk resources.
 */
export interface VirtualDiskState {
    /**
     * The adapter type for this virtual disk. Can be
     * one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
     */
    readonly adapterType?: pulumi.Input<string>;
    /**
     * Tells the resource to create any
     * directories that are a part of the `vmdk_path` parameter if they are missing.
     * Default: `false`.
     */
    readonly createDirectories?: pulumi.Input<boolean>;
    /**
     * The name of the datacenter in which to create the
     * disk. Can be omitted when when ESXi or if there is only one datacenter in
     * your infrastructure.
     */
    readonly datacenter?: pulumi.Input<string>;
    /**
     * The name of the datastore in which to create the
     * disk.
     */
    readonly datastore?: pulumi.Input<string>;
    /**
     * Size of the disk (in GB).
     */
    readonly size?: pulumi.Input<number>;
    /**
     * The type of disk to create. Can be one of
     * `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
     * information on what each kind of disk provisioning policy means, click
     * [here][docs-vmware-vm-disk-provisioning].
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The path, including filename, of the virtual disk to
     * be created.  This needs to end in `.vmdk`.
     */
    readonly vmdkPath?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a VirtualDisk resource.
 */
export interface VirtualDiskArgs {
    /**
     * The adapter type for this virtual disk. Can be
     * one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
     */
    readonly adapterType?: pulumi.Input<string>;
    /**
     * Tells the resource to create any
     * directories that are a part of the `vmdk_path` parameter if they are missing.
     * Default: `false`.
     */
    readonly createDirectories?: pulumi.Input<boolean>;
    /**
     * The name of the datacenter in which to create the
     * disk. Can be omitted when when ESXi or if there is only one datacenter in
     * your infrastructure.
     */
    readonly datacenter?: pulumi.Input<string>;
    /**
     * The name of the datastore in which to create the
     * disk.
     */
    readonly datastore: pulumi.Input<string>;
    /**
     * Size of the disk (in GB).
     */
    readonly size: pulumi.Input<number>;
    /**
     * The type of disk to create. Can be one of
     * `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
     * information on what each kind of disk provisioning policy means, click
     * [here][docs-vmware-vm-disk-provisioning].
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The path, including filename, of the virtual disk to
     * be created.  This needs to end in `.vmdk`.
     */
    readonly vmdkPath: pulumi.Input<string>;
}
