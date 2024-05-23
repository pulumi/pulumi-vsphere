// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.FileArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.FileState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### Uploading a File
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vsphere.File;
 * import com.pulumi.vsphere.FileArgs;
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
 *         var ubuntuVmdkUpload = new File("ubuntuVmdkUpload", FileArgs.builder()
 *             .datacenter("dc-01")
 *             .datastore("datastore-01")
 *             .sourceFile("/my/src/path/custom_ubuntu.vmdk")
 *             .destinationFile("/my/dst/path/custom_ubuntu.vmdk")
 *             .createDirectories(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Copying a File
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vsphere.File;
 * import com.pulumi.vsphere.FileArgs;
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
 *         var ubuntuCopy = new File("ubuntuCopy", FileArgs.builder()
 *             .sourceDatacenter("dc-01")
 *             .datacenter("dc-01")
 *             .sourceDatastore("datastore-01")
 *             .datastore("datastore-01")
 *             .sourceFile("/my/src/path/custom_ubuntu.vmdk")
 *             .destinationFile("/my/dst/path/custom_ubuntu.vmdk")
 *             .createDirectories(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vsphere:index/file:File")
public class File extends com.pulumi.resources.CustomResource {
    /**
     * Create directories in `destination_file`
     * path parameter on first apply if any are missing for copy operation.
     * 
     * &gt; **NOTE:** Any directory created as part of the `create_directories` argument
     * will not be deleted when the resource is destroyed. New directories are not
     * created if the `destination_file` path is changed in subsequent applies.
     * 
     */
    @Export(name="createDirectories", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> createDirectories;

    /**
     * @return Create directories in `destination_file`
     * path parameter on first apply if any are missing for copy operation.
     * 
     * &gt; **NOTE:** Any directory created as part of the `create_directories` argument
     * will not be deleted when the resource is destroyed. New directories are not
     * created if the `destination_file` path is changed in subsequent applies.
     * 
     */
    public Output<Optional<Boolean>> createDirectories() {
        return Codegen.optional(this.createDirectories);
    }
    /**
     * The name of a datacenter to which the file will be
     * uploaded.
     * 
     */
    @Export(name="datacenter", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> datacenter;

    /**
     * @return The name of a datacenter to which the file will be
     * uploaded.
     * 
     */
    public Output<Optional<String>> datacenter() {
        return Codegen.optional(this.datacenter);
    }
    /**
     * The name of the datastore to which to upload the
     * file.
     * 
     */
    @Export(name="datastore", refs={String.class}, tree="[0]")
    private Output<String> datastore;

    /**
     * @return The name of the datastore to which to upload the
     * file.
     * 
     */
    public Output<String> datastore() {
        return this.datastore;
    }
    /**
     * The path to where the file should be uploaded
     * or copied to on the destination `datastore` in vSphere.
     * 
     */
    @Export(name="destinationFile", refs={String.class}, tree="[0]")
    private Output<String> destinationFile;

    /**
     * @return The path to where the file should be uploaded
     * or copied to on the destination `datastore` in vSphere.
     * 
     */
    public Output<String> destinationFile() {
        return this.destinationFile;
    }
    /**
     * The name of a datacenter from which the file
     * will be copied. Forces a new resource if changed.
     * 
     */
    @Export(name="sourceDatacenter", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceDatacenter;

    /**
     * @return The name of a datacenter from which the file
     * will be copied. Forces a new resource if changed.
     * 
     */
    public Output<Optional<String>> sourceDatacenter() {
        return Codegen.optional(this.sourceDatacenter);
    }
    /**
     * The name of the datastore from which file will
     * be copied. Forces a new resource if changed.
     * 
     */
    @Export(name="sourceDatastore", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceDatastore;

    /**
     * @return The name of the datastore from which file will
     * be copied. Forces a new resource if changed.
     * 
     */
    public Output<Optional<String>> sourceDatastore() {
        return Codegen.optional(this.sourceDatastore);
    }
    @Export(name="sourceFile", refs={String.class}, tree="[0]")
    private Output<String> sourceFile;

    public Output<String> sourceFile() {
        return this.sourceFile;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public File(String name) {
        this(name, FileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public File(String name, FileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public File(String name, FileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/file:File", name, args == null ? FileArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private File(String name, Output<String> id, @Nullable FileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/file:File", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static File get(String name, Output<String> id, @Nullable FileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new File(name, id, state, options);
    }
}
