// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.VirtualDiskArgs;
import com.pulumi.vsphere.inputs.VirtualDiskState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `vsphere.VirtualDisk` resource can be used to create virtual disks outside
 * of any given `vsphere.VirtualMachine`
 * resource. These disks can be attached to a virtual machine by creating a disk
 * block with the `attach` parameter.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vsphere.VsphereFunctions;
 * import com.pulumi.vsphere.inputs.GetDatacenterArgs;
 * import com.pulumi.vsphere.VirtualDisk;
 * import com.pulumi.vsphere.VirtualDiskArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var datacenter = VsphereFunctions.getDatacenter(GetDatacenterArgs.builder()
 *             .name("dc-01")
 *             .build());
 * 
 *         final var datastore = VsphereFunctions.getDatacenter(GetDatacenterArgs.builder()
 *             .name("datastore-01")
 *             .build());
 * 
 *         var virtualDisk = new VirtualDisk("virtualDisk", VirtualDiskArgs.builder()
 *             .size(40)
 *             .type("thin")
 *             .vmdkPath("/foo/foo.vmdk")
 *             .createDirectories(true)
 *             .datacenter(datacenter.applyValue(getDatacenterResult -> getDatacenterResult.name()))
 *             .datastore(datastoreVsphereDatastore.name())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * An existing virtual disk can be imported into this resource
 * 
 * via supplying the full datastore path to the virtual disk. An example is below:
 * 
 * ```sh
 * $ pulumi import vsphere:index/virtualDisk:VirtualDisk virtual_disk \
 * ```
 * 
 *   &#39;{&#34;virtual_disk_path&#34;: &#34;/dc-01/[datastore-01]foo/bar.vmdk&#34;, \ &#34;create_directories&#34;: &#34;true&#34;}&#39;
 * 
 * The above would import the virtual disk located at `foo/bar.vmdk` in the `datastore-01`
 * 
 * datastore of the `dc-01` datacenter with `create_directories` set as `true`.
 * 
 */
@ResourceType(type="vsphere:index/virtualDisk:VirtualDisk")
public class VirtualDisk extends com.pulumi.resources.CustomResource {
    /**
     * The adapter type for this virtual disk. Can be
     * one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
     * 
     * &gt; **NOTE:** `adapter_type` is **deprecated**: it does not dictate the type of
     * controller that the virtual disk will be attached to on the virtual machine.
     * Please see the `scsi_type` parameter
     * in the `vsphere.VirtualMachine` resource for information on how to control
     * disk controller types. This parameter will be removed in future versions of the
     * vSphere provider.
     * 
     * @deprecated
     * this attribute has no effect on controller types - please use scsi_type in vsphere.VirtualMachine instead
     * 
     */
    @Deprecated /* this attribute has no effect on controller types - please use scsi_type in vsphere.VirtualMachine instead */
    @Export(name="adapterType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> adapterType;

    /**
     * @return The adapter type for this virtual disk. Can be
     * one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
     * 
     * &gt; **NOTE:** `adapter_type` is **deprecated**: it does not dictate the type of
     * controller that the virtual disk will be attached to on the virtual machine.
     * Please see the `scsi_type` parameter
     * in the `vsphere.VirtualMachine` resource for information on how to control
     * disk controller types. This parameter will be removed in future versions of the
     * vSphere provider.
     * 
     */
    public Output<Optional<String>> adapterType() {
        return Codegen.optional(this.adapterType);
    }
    /**
     * Tells the resource to create any
     * directories that are a part of the `vmdk_path` parameter if they are missing.
     * Default: `false`.
     * 
     * &gt; **NOTE:** Any directory created as part of the operation when
     * `create_directories` is enabled will not be deleted when the resource is
     * destroyed.
     * 
     */
    @Export(name="createDirectories", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> createDirectories;

    /**
     * @return Tells the resource to create any
     * directories that are a part of the `vmdk_path` parameter if they are missing.
     * Default: `false`.
     * 
     * &gt; **NOTE:** Any directory created as part of the operation when
     * `create_directories` is enabled will not be deleted when the resource is
     * destroyed.
     * 
     */
    public Output<Optional<Boolean>> createDirectories() {
        return Codegen.optional(this.createDirectories);
    }
    /**
     * The name of the datacenter in which to create the
     * disk. Can be omitted when ESXi or if there is only one datacenter in
     * your infrastructure.
     * 
     */
    @Export(name="datacenter", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> datacenter;

    /**
     * @return The name of the datacenter in which to create the
     * disk. Can be omitted when ESXi or if there is only one datacenter in
     * your infrastructure.
     * 
     */
    public Output<Optional<String>> datacenter() {
        return Codegen.optional(this.datacenter);
    }
    /**
     * The name of the datastore in which to create the
     * disk.
     * 
     */
    @Export(name="datastore", refs={String.class}, tree="[0]")
    private Output<String> datastore;

    /**
     * @return The name of the datastore in which to create the
     * disk.
     * 
     */
    public Output<String> datastore() {
        return this.datastore;
    }
    /**
     * Size of the disk (in GB). Decreasing the size of a disk is not possible.
     * If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
     * created.
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output<Integer> size;

    /**
     * @return Size of the disk (in GB). Decreasing the size of a disk is not possible.
     * If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
     * created.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * The type of disk to create. Can be one of
     * `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
     * information on what each kind of disk provisioning policy means, click
     * [here][docs-vmware-vm-disk-provisioning].
     * 
     * [docs-vmware-vm-disk-provisioning]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vm-administration/GUID-4C0F4D73-82F2-4B81-8AA7-1DD752A8A5AC.html
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The type of disk to create. Can be one of
     * `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
     * information on what each kind of disk provisioning policy means, click
     * [here][docs-vmware-vm-disk-provisioning].
     * 
     * [docs-vmware-vm-disk-provisioning]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vm-administration/GUID-4C0F4D73-82F2-4B81-8AA7-1DD752A8A5AC.html
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * The path, including filename, of the virtual disk to
     * be created.  This needs to end in `.vmdk`.
     * 
     */
    @Export(name="vmdkPath", refs={String.class}, tree="[0]")
    private Output<String> vmdkPath;

    /**
     * @return The path, including filename, of the virtual disk to
     * be created.  This needs to end in `.vmdk`.
     * 
     */
    public Output<String> vmdkPath() {
        return this.vmdkPath;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VirtualDisk(java.lang.String name) {
        this(name, VirtualDiskArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VirtualDisk(java.lang.String name, VirtualDiskArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VirtualDisk(java.lang.String name, VirtualDiskArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/virtualDisk:VirtualDisk", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private VirtualDisk(java.lang.String name, Output<java.lang.String> id, @Nullable VirtualDiskState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/virtualDisk:VirtualDisk", name, state, makeResourceOptions(options, id), false);
    }

    private static VirtualDiskArgs makeArgs(VirtualDiskArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VirtualDiskArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static VirtualDisk get(java.lang.String name, Output<java.lang.String> id, @Nullable VirtualDiskState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VirtualDisk(name, id, state, options);
    }
}
