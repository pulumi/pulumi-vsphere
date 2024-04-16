// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.DatacenterArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.DatacenterState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VMware vSphere datacenter resource. This can be used as the primary
 * container of inventory objects such as hosts and virtual machines.
 * 
 * ## Example Usage
 * 
 * ### Create datacenter on the root folder
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vsphere.Datacenter;
 * import com.pulumi.vsphere.DatacenterArgs;
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
 *         var prodDatacenter = new Datacenter(&#34;prodDatacenter&#34;, DatacenterArgs.builder()        
 *             .name(&#34;my_prod_datacenter&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Create datacenter on a subfolder
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vsphere.Datacenter;
 * import com.pulumi.vsphere.DatacenterArgs;
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
 *         var researchDatacenter = new Datacenter(&#34;researchDatacenter&#34;, DatacenterArgs.builder()        
 *             .name(&#34;my_research_datacenter&#34;)
 *             .folder(&#34;/research/&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Importing
 * 
 * An existing datacenter can be [imported][docs-import] into this resource
 * via supplying the full path to the datacenter. An example is below:
 * 
 * [docs-import]: /docs/import/index.html
 * 
 * The above would import the datacenter named `dc1`.
 * 
 */
@ResourceType(type="vsphere:index/datacenter:Datacenter")
public class Datacenter extends com.pulumi.resources.CustomResource {
    /**
     * Map of custom attribute ids to value
     * strings to set for datacenter resource. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     * 
     * [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
     * 
     * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     * 
     */
    @Export(name="customAttributes", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> customAttributes;

    /**
     * @return Map of custom attribute ids to value
     * strings to set for datacenter resource. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     * 
     * [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
     * 
     * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     * 
     */
    public Output<Optional<Map<String,String>>> customAttributes() {
        return Codegen.optional(this.customAttributes);
    }
    /**
     * The folder where the datacenter should be created.
     * Forces a new resource if changed.
     * 
     */
    @Export(name="folder", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> folder;

    /**
     * @return The folder where the datacenter should be created.
     * Forces a new resource if changed.
     * 
     */
    public Output<Optional<String>> folder() {
        return Codegen.optional(this.folder);
    }
    /**
     * Managed object ID of this datacenter.
     * 
     */
    @Export(name="moid", refs={String.class}, tree="[0]")
    private Output<String> moid;

    /**
     * @return Managed object ID of this datacenter.
     * 
     */
    public Output<String> moid() {
        return this.moid;
    }
    /**
     * The name of the datacenter. This name needs to be unique
     * within the folder. Forces a new resource if changed.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the datacenter. This name needs to be unique
     * within the folder. Forces a new resource if changed.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The IDs of any tags to attach to this resource.
     * 
     * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return The IDs of any tags to attach to this resource.
     * 
     * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Datacenter(String name) {
        this(name, DatacenterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Datacenter(String name, @Nullable DatacenterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Datacenter(String name, @Nullable DatacenterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/datacenter:Datacenter", name, args == null ? DatacenterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Datacenter(String name, Output<String> id, @Nullable DatacenterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/datacenter:Datacenter", name, state, makeResourceOptions(options, id));
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
    public static Datacenter get(String name, Output<String> id, @Nullable DatacenterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Datacenter(name, id, state, options);
    }
}
