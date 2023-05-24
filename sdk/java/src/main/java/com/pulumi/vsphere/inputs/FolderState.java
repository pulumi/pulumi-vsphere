// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FolderState extends com.pulumi.resources.ResourceArgs {

    public static final FolderState Empty = new FolderState();

    /**
     * Map of custom attribute ids to attribute
     * value strings to set for folder. See [here][docs-setting-custom-attributes]
     * for a reference on how to set values for custom attributes.
     * 
     * [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
     * 
     * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     * 
     */
    @Import(name="customAttributes")
    private @Nullable Output<Map<String,String>> customAttributes;

    /**
     * @return Map of custom attribute ids to attribute
     * value strings to set for folder. See [here][docs-setting-custom-attributes]
     * for a reference on how to set values for custom attributes.
     * 
     * [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
     * 
     * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     * 
     */
    public Optional<Output<Map<String,String>>> customAttributes() {
        return Optional.ofNullable(this.customAttributes);
    }

    /**
     * The ID of the datacenter the folder will be created in.
     * Required for all folder types except for datacenter folders. Forces a new
     * resource if changed.
     * 
     */
    @Import(name="datacenterId")
    private @Nullable Output<String> datacenterId;

    /**
     * @return The ID of the datacenter the folder will be created in.
     * Required for all folder types except for datacenter folders. Forces a new
     * resource if changed.
     * 
     */
    public Optional<Output<String>> datacenterId() {
        return Optional.ofNullable(this.datacenterId);
    }

    /**
     * The path of the folder to be created. This is relative to
     * the root of the type of folder you are creating, and the supplied datacenter.
     * For example, given a default datacenter of `default-dc`, a folder of type
     * `vm` (denoting a virtual machine folder), and a supplied folder of
     * `test-folder`, the resulting path would be
     * `/default-dc/vm/test-folder`.
     * 
     * &gt; **NOTE:** `path` can be modified - the resulting behavior is dependent on
     * what section of `path` you are modifying. If you are modifying the parent (so
     * any part before the last `/`), your folder will be moved to that new parent. If
     * modifying the name (the part after the last `/`), your folder will be renamed.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path of the folder to be created. This is relative to
     * the root of the type of folder you are creating, and the supplied datacenter.
     * For example, given a default datacenter of `default-dc`, a folder of type
     * `vm` (denoting a virtual machine folder), and a supplied folder of
     * `test-folder`, the resulting path would be
     * `/default-dc/vm/test-folder`.
     * 
     * &gt; **NOTE:** `path` can be modified - the resulting behavior is dependent on
     * what section of `path` you are modifying. If you are modifying the parent (so
     * any part before the last `/`), your folder will be moved to that new parent. If
     * modifying the name (the part after the last `/`), your folder will be renamed.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The IDs of any tags to attach to this resource.
     * 
     * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The IDs of any tags to attach to this resource.
     * 
     * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The type of folder to create. Allowed options are
     * `datacenter` for datacenter folders, `host` for host and cluster folders,
     * `vm` for virtual machine folders, `datastore` for datastore folders, and
     * `network` for network folders. Forces a new resource if changed.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of folder to create. Allowed options are
     * `datacenter` for datacenter folders, `host` for host and cluster folders,
     * `vm` for virtual machine folders, `datastore` for datastore folders, and
     * `network` for network folders. Forces a new resource if changed.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private FolderState() {}

    private FolderState(FolderState $) {
        this.customAttributes = $.customAttributes;
        this.datacenterId = $.datacenterId;
        this.path = $.path;
        this.tags = $.tags;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FolderState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FolderState $;

        public Builder() {
            $ = new FolderState();
        }

        public Builder(FolderState defaults) {
            $ = new FolderState(Objects.requireNonNull(defaults));
        }

        /**
         * @param customAttributes Map of custom attribute ids to attribute
         * value strings to set for folder. See [here][docs-setting-custom-attributes]
         * for a reference on how to set values for custom attributes.
         * 
         * [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
         * 
         * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
         * and require vCenter.
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(@Nullable Output<Map<String,String>> customAttributes) {
            $.customAttributes = customAttributes;
            return this;
        }

        /**
         * @param customAttributes Map of custom attribute ids to attribute
         * value strings to set for folder. See [here][docs-setting-custom-attributes]
         * for a reference on how to set values for custom attributes.
         * 
         * [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
         * 
         * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
         * and require vCenter.
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(Map<String,String> customAttributes) {
            return customAttributes(Output.of(customAttributes));
        }

        /**
         * @param datacenterId The ID of the datacenter the folder will be created in.
         * Required for all folder types except for datacenter folders. Forces a new
         * resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(@Nullable Output<String> datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        /**
         * @param datacenterId The ID of the datacenter the folder will be created in.
         * Required for all folder types except for datacenter folders. Forces a new
         * resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(String datacenterId) {
            return datacenterId(Output.of(datacenterId));
        }

        /**
         * @param path The path of the folder to be created. This is relative to
         * the root of the type of folder you are creating, and the supplied datacenter.
         * For example, given a default datacenter of `default-dc`, a folder of type
         * `vm` (denoting a virtual machine folder), and a supplied folder of
         * `test-folder`, the resulting path would be
         * `/default-dc/vm/test-folder`.
         * 
         * &gt; **NOTE:** `path` can be modified - the resulting behavior is dependent on
         * what section of `path` you are modifying. If you are modifying the parent (so
         * any part before the last `/`), your folder will be moved to that new parent. If
         * modifying the name (the part after the last `/`), your folder will be renamed.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path of the folder to be created. This is relative to
         * the root of the type of folder you are creating, and the supplied datacenter.
         * For example, given a default datacenter of `default-dc`, a folder of type
         * `vm` (denoting a virtual machine folder), and a supplied folder of
         * `test-folder`, the resulting path would be
         * `/default-dc/vm/test-folder`.
         * 
         * &gt; **NOTE:** `path` can be modified - the resulting behavior is dependent on
         * what section of `path` you are modifying. If you are modifying the parent (so
         * any part before the last `/`), your folder will be moved to that new parent. If
         * modifying the name (the part after the last `/`), your folder will be renamed.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param tags The IDs of any tags to attach to this resource.
         * 
         * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
         * requires vCenter 6.0 or higher.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The IDs of any tags to attach to this resource.
         * 
         * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
         * requires vCenter 6.0 or higher.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The IDs of any tags to attach to this resource.
         * 
         * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
         * requires vCenter 6.0 or higher.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param type The type of folder to create. Allowed options are
         * `datacenter` for datacenter folders, `host` for host and cluster folders,
         * `vm` for virtual machine folders, `datastore` for datastore folders, and
         * `network` for network folders. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of folder to create. Allowed options are
         * `datacenter` for datacenter folders, `host` for host and cluster folders,
         * `vm` for virtual machine folders, `datastore` for datastore folders, and
         * `network` for network folders. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public FolderState build() {
            return $;
        }
    }

}
