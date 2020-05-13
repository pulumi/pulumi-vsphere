// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `vsphere..VirtualMachineSnapshot` resource can be used to manage snapshots
 * for a virtual machine.
 * 
 * For more information on managing snapshots and how they work in VMware, see
 * [here][ext-vm-snapshot-management].
 * 
 * [ext-vm-snapshot-management]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vm_admin.doc/GUID-CA948C69-7F58-4519-AEB1-739545EA94E5.html
 * 
 * > **NOTE:** A snapshot in VMware differs from traditional disk snapshots, and
 * can contain the actual running state of the virtual machine, data for all disks
 * that have not been set to be independent from the snapshot (including ones that
 * have been attached via the `attach`
 * parameter to the `vsphere..VirtualMachine` `disk` block), and even the
 * configuration of the virtual machine at the time of the snapshot. Virtual
 * machine, disk activity, and configuration changes post-snapshot are not
 * included in the original state. Use this resource with care! Neither VMware nor
 * HashiCorp recommends retaining snapshots for a extended period of time and does
 * NOT recommend using them as as backup feature. For more information on the
 * limitation of virtual machine snapshots, see [here][ext-vm-snap-limitations].
 * 
 * [ext-vm-snap-limitations]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vm_admin.doc/GUID-53F65726-A23B-4CF0-A7D5-48E584B88613.html
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const demo1 = new vsphere.VirtualMachineSnapshot("demo1", {
 *     consolidate: true,
 *     description: "This is Demo Snapshot",
 *     memory: true,
 *     quiesce: true,
 *     removeChildren: false,
 *     snapshotName: "Snapshot Name",
 *     virtualMachineUuid: "9aac5551-a351-4158-8c5c-15a71e8ec5c9",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/virtual_machine_snapshot.html.markdown.
 */
export class VirtualMachineSnapshot extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachineSnapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualMachineSnapshotState, opts?: pulumi.CustomResourceOptions): VirtualMachineSnapshot {
        return new VirtualMachineSnapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/virtualMachineSnapshot:VirtualMachineSnapshot';

    /**
     * Returns true if the given object is an instance of VirtualMachineSnapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualMachineSnapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualMachineSnapshot.__pulumiType;
    }

    /**
     * If set to `true`, the delta disks involved in this
     * snapshot will be consolidated into the parent when this resource is
     * destroyed.
     */
    public readonly consolidate!: pulumi.Output<boolean | undefined>;
    /**
     * A description for the snapshot.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * If set to `true`, a dump of the internal state of the
     * virtual machine is included in the snapshot.
     */
    public readonly memory!: pulumi.Output<boolean>;
    /**
     * If set to `true`, and the virtual machine is powered
     * on when the snapshot is taken, VMware Tools is used to quiesce the file
     * system in the virtual machine.
     */
    public readonly quiesce!: pulumi.Output<boolean>;
    /**
     * If set to `true`, the entire snapshot subtree
     * is removed when this resource is destroyed.
     */
    public readonly removeChildren!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the snapshot.
     */
    public readonly snapshotName!: pulumi.Output<string>;
    /**
     * The virtual machine UUID.
     */
    public readonly virtualMachineUuid!: pulumi.Output<string>;

    /**
     * Create a VirtualMachineSnapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineSnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualMachineSnapshotArgs | VirtualMachineSnapshotState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VirtualMachineSnapshotState | undefined;
            inputs["consolidate"] = state ? state.consolidate : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["memory"] = state ? state.memory : undefined;
            inputs["quiesce"] = state ? state.quiesce : undefined;
            inputs["removeChildren"] = state ? state.removeChildren : undefined;
            inputs["snapshotName"] = state ? state.snapshotName : undefined;
            inputs["virtualMachineUuid"] = state ? state.virtualMachineUuid : undefined;
        } else {
            const args = argsOrState as VirtualMachineSnapshotArgs | undefined;
            if (!args || args.description === undefined) {
                throw new Error("Missing required property 'description'");
            }
            if (!args || args.memory === undefined) {
                throw new Error("Missing required property 'memory'");
            }
            if (!args || args.quiesce === undefined) {
                throw new Error("Missing required property 'quiesce'");
            }
            if (!args || args.snapshotName === undefined) {
                throw new Error("Missing required property 'snapshotName'");
            }
            if (!args || args.virtualMachineUuid === undefined) {
                throw new Error("Missing required property 'virtualMachineUuid'");
            }
            inputs["consolidate"] = args ? args.consolidate : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["memory"] = args ? args.memory : undefined;
            inputs["quiesce"] = args ? args.quiesce : undefined;
            inputs["removeChildren"] = args ? args.removeChildren : undefined;
            inputs["snapshotName"] = args ? args.snapshotName : undefined;
            inputs["virtualMachineUuid"] = args ? args.virtualMachineUuid : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualMachineSnapshot.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualMachineSnapshot resources.
 */
export interface VirtualMachineSnapshotState {
    /**
     * If set to `true`, the delta disks involved in this
     * snapshot will be consolidated into the parent when this resource is
     * destroyed.
     */
    readonly consolidate?: pulumi.Input<boolean>;
    /**
     * A description for the snapshot.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If set to `true`, a dump of the internal state of the
     * virtual machine is included in the snapshot.
     */
    readonly memory?: pulumi.Input<boolean>;
    /**
     * If set to `true`, and the virtual machine is powered
     * on when the snapshot is taken, VMware Tools is used to quiesce the file
     * system in the virtual machine.
     */
    readonly quiesce?: pulumi.Input<boolean>;
    /**
     * If set to `true`, the entire snapshot subtree
     * is removed when this resource is destroyed.
     */
    readonly removeChildren?: pulumi.Input<boolean>;
    /**
     * The name of the snapshot.
     */
    readonly snapshotName?: pulumi.Input<string>;
    /**
     * The virtual machine UUID.
     */
    readonly virtualMachineUuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualMachineSnapshot resource.
 */
export interface VirtualMachineSnapshotArgs {
    /**
     * If set to `true`, the delta disks involved in this
     * snapshot will be consolidated into the parent when this resource is
     * destroyed.
     */
    readonly consolidate?: pulumi.Input<boolean>;
    /**
     * A description for the snapshot.
     */
    readonly description: pulumi.Input<string>;
    /**
     * If set to `true`, a dump of the internal state of the
     * virtual machine is included in the snapshot.
     */
    readonly memory: pulumi.Input<boolean>;
    /**
     * If set to `true`, and the virtual machine is powered
     * on when the snapshot is taken, VMware Tools is used to quiesce the file
     * system in the virtual machine.
     */
    readonly quiesce: pulumi.Input<boolean>;
    /**
     * If set to `true`, the entire snapshot subtree
     * is removed when this resource is destroyed.
     */
    readonly removeChildren?: pulumi.Input<boolean>;
    /**
     * The name of the snapshot.
     */
    readonly snapshotName: pulumi.Input<string>;
    /**
     * The virtual machine UUID.
     */
    readonly virtualMachineUuid: pulumi.Input<string>;
}
