// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.File` resource can be used to upload files (such as virtual disk
 * files) from the host machine that this provider is running on to a target
 * datastore.  The resource can also be used to copy files between datastores, or
 * from one location to another on the same datastore.
 *
 * Updates to destination parameters such as `datacenter`, `datastore`, or
 * `destinationFile` will move the managed file a new destination based on the
 * values of the new settings.  If any source parameter is changed, such as
 * `sourceDatastore`, `sourceDatacenter` or `sourceFile`), the resource will be
 * re-created. Depending on if destination parameters are being changed as well,
 * this may result in the destination file either being overwritten or deleted at
 * the old location.
 *
 * ## Example Usage
 *
 * ### Uploading a file
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const ubuntuDiskUpload = new vsphere.File("ubuntu_disk_upload", {
 *     datacenter: "my_datacenter",
 *     datastore: "local",
 *     destinationFile: "/my_path/disks/custom_ubuntu.vmdk",
 *     sourceFile: "/home/ubuntu/my_disks/custom_ubuntu.vmdk",
 * });
 * ```
 *
 * ### Copying a file
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const ubuntuDiskCopy = new vsphere.File("ubuntu_disk_copy", {
 *     datacenter: "my_datacenter",
 *     datastore: "local",
 *     destinationFile: "/my_path/custom_ubuntu_id.vmdk",
 *     sourceDatacenter: "my_datacenter",
 *     sourceDatastore: "local",
 *     sourceFile: "/my_path/disks/custom_ubuntu.vmdk",
 * });
 * ```
 */
export class File extends pulumi.CustomResource {
    /**
     * Get an existing File resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FileState, opts?: pulumi.CustomResourceOptions): File {
        return new File(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/file:File';

    /**
     * Returns true if the given object is an instance of File.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is File {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === File.__pulumiType;
    }

    /**
     * Create directories in `destinationFile`
     * path parameter if any missing for copy operation.
     */
    public readonly createDirectories!: pulumi.Output<boolean | undefined>;
    /**
     * The name of a datacenter in which the file will be
     * uploaded to.
     */
    public readonly datacenter!: pulumi.Output<string | undefined>;
    /**
     * The name of the datastore in which to upload the
     * file to.
     */
    public readonly datastore!: pulumi.Output<string>;
    /**
     * The path to where the file should be uploaded
     * or copied to on vSphere.
     */
    public readonly destinationFile!: pulumi.Output<string>;
    /**
     * The name of a datacenter in which the file
     * will be copied from. Forces a new resource if changed.
     */
    public readonly sourceDatacenter!: pulumi.Output<string | undefined>;
    /**
     * The name of the datastore in which file will
     * be copied from. Forces a new resource if changed.
     */
    public readonly sourceDatastore!: pulumi.Output<string | undefined>;
    /**
     * The path to the file being uploaded from the
     * host to vSphere or copied within vSphere. Forces a new resource if
     * changed.
     */
    public readonly sourceFile!: pulumi.Output<string>;

    /**
     * Create a File resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FileArgs | FileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FileState | undefined;
            inputs["createDirectories"] = state ? state.createDirectories : undefined;
            inputs["datacenter"] = state ? state.datacenter : undefined;
            inputs["datastore"] = state ? state.datastore : undefined;
            inputs["destinationFile"] = state ? state.destinationFile : undefined;
            inputs["sourceDatacenter"] = state ? state.sourceDatacenter : undefined;
            inputs["sourceDatastore"] = state ? state.sourceDatastore : undefined;
            inputs["sourceFile"] = state ? state.sourceFile : undefined;
        } else {
            const args = argsOrState as FileArgs | undefined;
            if (!args || args.datastore === undefined) {
                throw new Error("Missing required property 'datastore'");
            }
            if (!args || args.destinationFile === undefined) {
                throw new Error("Missing required property 'destinationFile'");
            }
            if (!args || args.sourceFile === undefined) {
                throw new Error("Missing required property 'sourceFile'");
            }
            inputs["createDirectories"] = args ? args.createDirectories : undefined;
            inputs["datacenter"] = args ? args.datacenter : undefined;
            inputs["datastore"] = args ? args.datastore : undefined;
            inputs["destinationFile"] = args ? args.destinationFile : undefined;
            inputs["sourceDatacenter"] = args ? args.sourceDatacenter : undefined;
            inputs["sourceDatastore"] = args ? args.sourceDatastore : undefined;
            inputs["sourceFile"] = args ? args.sourceFile : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(File.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering File resources.
 */
export interface FileState {
    /**
     * Create directories in `destinationFile`
     * path parameter if any missing for copy operation.
     */
    readonly createDirectories?: pulumi.Input<boolean>;
    /**
     * The name of a datacenter in which the file will be
     * uploaded to.
     */
    readonly datacenter?: pulumi.Input<string>;
    /**
     * The name of the datastore in which to upload the
     * file to.
     */
    readonly datastore?: pulumi.Input<string>;
    /**
     * The path to where the file should be uploaded
     * or copied to on vSphere.
     */
    readonly destinationFile?: pulumi.Input<string>;
    /**
     * The name of a datacenter in which the file
     * will be copied from. Forces a new resource if changed.
     */
    readonly sourceDatacenter?: pulumi.Input<string>;
    /**
     * The name of the datastore in which file will
     * be copied from. Forces a new resource if changed.
     */
    readonly sourceDatastore?: pulumi.Input<string>;
    /**
     * The path to the file being uploaded from the
     * host to vSphere or copied within vSphere. Forces a new resource if
     * changed.
     */
    readonly sourceFile?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a File resource.
 */
export interface FileArgs {
    /**
     * Create directories in `destinationFile`
     * path parameter if any missing for copy operation.
     */
    readonly createDirectories?: pulumi.Input<boolean>;
    /**
     * The name of a datacenter in which the file will be
     * uploaded to.
     */
    readonly datacenter?: pulumi.Input<string>;
    /**
     * The name of the datastore in which to upload the
     * file to.
     */
    readonly datastore: pulumi.Input<string>;
    /**
     * The path to where the file should be uploaded
     * or copied to on vSphere.
     */
    readonly destinationFile: pulumi.Input<string>;
    /**
     * The name of a datacenter in which the file
     * will be copied from. Forces a new resource if changed.
     */
    readonly sourceDatacenter?: pulumi.Input<string>;
    /**
     * The name of the datastore in which file will
     * be copied from. Forces a new resource if changed.
     */
    readonly sourceDatastore?: pulumi.Input<string>;
    /**
     * The path to the file being uploaded from the
     * host to vSphere or copied within vSphere. Forces a new resource if
     * changed.
     */
    readonly sourceFile: pulumi.Input<string>;
}
