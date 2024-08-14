// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class VirtualDisk extends pulumi.CustomResource {
    /**
     * Get an existing VirtualDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualDiskState, opts?: pulumi.CustomResourceOptions): VirtualDisk {
        return new VirtualDisk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/virtualDisk:VirtualDisk';

    /**
     * Returns true if the given object is an instance of VirtualDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualDisk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualDisk.__pulumiType;
    }

    /**
     * The adapter type for this virtual disk. Can be
     * one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
     *
     * > **NOTE:** `adapterType` is **deprecated**: it does not dictate the type of
     * controller that the virtual disk will be attached to on the virtual machine.
     * Please see the `scsiType` parameter
     * in the `vsphere.VirtualMachine` resource for information on how to control
     * disk controller types. This parameter will be removed in future versions of the
     * vSphere provider.
     *
     * @deprecated this attribute has no effect on controller types - please use scsiType in vsphere.VirtualMachine instead
     */
    public readonly adapterType!: pulumi.Output<string | undefined>;
    /**
     * Tells the resource to create any
     * directories that are a part of the `vmdkPath` parameter if they are missing.
     * Default: `false`.
     *
     * > **NOTE:** Any directory created as part of the operation when
     * `createDirectories` is enabled will not be deleted when the resource is
     * destroyed.
     */
    public readonly createDirectories!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the datacenter in which to create the
     * disk. Can be omitted when ESXi or if there is only one datacenter in
     * your infrastructure.
     */
    public readonly datacenter!: pulumi.Output<string | undefined>;
    /**
     * The name of the datastore in which to create the
     * disk.
     */
    public readonly datastore!: pulumi.Output<string>;
    /**
     * Size of the disk (in GB). Decreasing the size of a disk is not possible.
     * If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
     * created.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * The type of disk to create. Can be one of
     * `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
     * information on what each kind of disk provisioning policy means, click
     * [here][docs-vmware-vm-disk-provisioning].
     *
     * [docs-vmware-vm-disk-provisioning]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vm-administration/GUID-4C0F4D73-82F2-4B81-8AA7-1DD752A8A5AC.html
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The path, including filename, of the virtual disk to
     * be created.  This needs to end in `.vmdk`.
     */
    public readonly vmdkPath!: pulumi.Output<string>;

    /**
     * Create a VirtualDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualDiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualDiskArgs | VirtualDiskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualDiskState | undefined;
            resourceInputs["adapterType"] = state ? state.adapterType : undefined;
            resourceInputs["createDirectories"] = state ? state.createDirectories : undefined;
            resourceInputs["datacenter"] = state ? state.datacenter : undefined;
            resourceInputs["datastore"] = state ? state.datastore : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vmdkPath"] = state ? state.vmdkPath : undefined;
        } else {
            const args = argsOrState as VirtualDiskArgs | undefined;
            if ((!args || args.datastore === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datastore'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            if ((!args || args.vmdkPath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vmdkPath'");
            }
            resourceInputs["adapterType"] = args ? args.adapterType : undefined;
            resourceInputs["createDirectories"] = args ? args.createDirectories : undefined;
            resourceInputs["datacenter"] = args ? args.datacenter : undefined;
            resourceInputs["datastore"] = args ? args.datastore : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vmdkPath"] = args ? args.vmdkPath : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualDisk.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualDisk resources.
 */
export interface VirtualDiskState {
    /**
     * The adapter type for this virtual disk. Can be
     * one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
     *
     * > **NOTE:** `adapterType` is **deprecated**: it does not dictate the type of
     * controller that the virtual disk will be attached to on the virtual machine.
     * Please see the `scsiType` parameter
     * in the `vsphere.VirtualMachine` resource for information on how to control
     * disk controller types. This parameter will be removed in future versions of the
     * vSphere provider.
     *
     * @deprecated this attribute has no effect on controller types - please use scsiType in vsphere.VirtualMachine instead
     */
    adapterType?: pulumi.Input<string>;
    /**
     * Tells the resource to create any
     * directories that are a part of the `vmdkPath` parameter if they are missing.
     * Default: `false`.
     *
     * > **NOTE:** Any directory created as part of the operation when
     * `createDirectories` is enabled will not be deleted when the resource is
     * destroyed.
     */
    createDirectories?: pulumi.Input<boolean>;
    /**
     * The name of the datacenter in which to create the
     * disk. Can be omitted when ESXi or if there is only one datacenter in
     * your infrastructure.
     */
    datacenter?: pulumi.Input<string>;
    /**
     * The name of the datastore in which to create the
     * disk.
     */
    datastore?: pulumi.Input<string>;
    /**
     * Size of the disk (in GB). Decreasing the size of a disk is not possible.
     * If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
     * created.
     */
    size?: pulumi.Input<number>;
    /**
     * The type of disk to create. Can be one of
     * `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
     * information on what each kind of disk provisioning policy means, click
     * [here][docs-vmware-vm-disk-provisioning].
     *
     * [docs-vmware-vm-disk-provisioning]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vm-administration/GUID-4C0F4D73-82F2-4B81-8AA7-1DD752A8A5AC.html
     */
    type?: pulumi.Input<string>;
    /**
     * The path, including filename, of the virtual disk to
     * be created.  This needs to end in `.vmdk`.
     */
    vmdkPath?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualDisk resource.
 */
export interface VirtualDiskArgs {
    /**
     * The adapter type for this virtual disk. Can be
     * one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
     *
     * > **NOTE:** `adapterType` is **deprecated**: it does not dictate the type of
     * controller that the virtual disk will be attached to on the virtual machine.
     * Please see the `scsiType` parameter
     * in the `vsphere.VirtualMachine` resource for information on how to control
     * disk controller types. This parameter will be removed in future versions of the
     * vSphere provider.
     *
     * @deprecated this attribute has no effect on controller types - please use scsiType in vsphere.VirtualMachine instead
     */
    adapterType?: pulumi.Input<string>;
    /**
     * Tells the resource to create any
     * directories that are a part of the `vmdkPath` parameter if they are missing.
     * Default: `false`.
     *
     * > **NOTE:** Any directory created as part of the operation when
     * `createDirectories` is enabled will not be deleted when the resource is
     * destroyed.
     */
    createDirectories?: pulumi.Input<boolean>;
    /**
     * The name of the datacenter in which to create the
     * disk. Can be omitted when ESXi or if there is only one datacenter in
     * your infrastructure.
     */
    datacenter?: pulumi.Input<string>;
    /**
     * The name of the datastore in which to create the
     * disk.
     */
    datastore: pulumi.Input<string>;
    /**
     * Size of the disk (in GB). Decreasing the size of a disk is not possible.
     * If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
     * created.
     */
    size: pulumi.Input<number>;
    /**
     * The type of disk to create. Can be one of
     * `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
     * information on what each kind of disk provisioning policy means, click
     * [here][docs-vmware-vm-disk-provisioning].
     *
     * [docs-vmware-vm-disk-provisioning]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vm-administration/GUID-4C0F4D73-82F2-4B81-8AA7-1DD752A8A5AC.html
     */
    type?: pulumi.Input<string>;
    /**
     * The path, including filename, of the virtual disk to
     * be created.  This needs to end in `.vmdk`.
     */
    vmdkPath: pulumi.Input<string>;
}
