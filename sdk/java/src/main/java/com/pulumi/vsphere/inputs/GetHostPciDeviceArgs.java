// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetHostPciDeviceArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetHostPciDeviceArgs Empty = new GetHostPciDeviceArgs();

    /**
     * The hexadecimal PCI device class ID
     * 
     * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
     * 
     * &gt; **NOTE:** `name_regex`, `vendor_id`, and `class_id` can all be used together.
     * 
     */
    @Import(name="classId")
    private @Nullable Output<String> classId;

    /**
     * @return The hexadecimal PCI device class ID
     * 
     * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
     * 
     * &gt; **NOTE:** `name_regex`, `vendor_id`, and `class_id` can all be used together.
     * 
     */
    public Optional<Output<String>> classId() {
        return Optional.ofNullable(this.classId);
    }

    /**
     * The [managed object reference ID][docs-about-morefs] of
     * a host.
     * 
     */
    @Import(name="hostId", required=true)
    private Output<String> hostId;

    /**
     * @return The [managed object reference ID][docs-about-morefs] of
     * a host.
     * 
     */
    public Output<String> hostId() {
        return this.hostId;
    }

    /**
     * A regular expression that will be used to match the
     * host PCI device name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regular expression that will be used to match the
     * host PCI device name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * The hexadecimal PCI device vendor ID.
     * 
     */
    @Import(name="vendorId")
    private @Nullable Output<String> vendorId;

    /**
     * @return The hexadecimal PCI device vendor ID.
     * 
     */
    public Optional<Output<String>> vendorId() {
        return Optional.ofNullable(this.vendorId);
    }

    private GetHostPciDeviceArgs() {}

    private GetHostPciDeviceArgs(GetHostPciDeviceArgs $) {
        this.classId = $.classId;
        this.hostId = $.hostId;
        this.nameRegex = $.nameRegex;
        this.vendorId = $.vendorId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetHostPciDeviceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetHostPciDeviceArgs $;

        public Builder() {
            $ = new GetHostPciDeviceArgs();
        }

        public Builder(GetHostPciDeviceArgs defaults) {
            $ = new GetHostPciDeviceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param classId The hexadecimal PCI device class ID
         * 
         * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
         * 
         * &gt; **NOTE:** `name_regex`, `vendor_id`, and `class_id` can all be used together.
         * 
         * @return builder
         * 
         */
        public Builder classId(@Nullable Output<String> classId) {
            $.classId = classId;
            return this;
        }

        /**
         * @param classId The hexadecimal PCI device class ID
         * 
         * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
         * 
         * &gt; **NOTE:** `name_regex`, `vendor_id`, and `class_id` can all be used together.
         * 
         * @return builder
         * 
         */
        public Builder classId(String classId) {
            return classId(Output.of(classId));
        }

        /**
         * @param hostId The [managed object reference ID][docs-about-morefs] of
         * a host.
         * 
         * @return builder
         * 
         */
        public Builder hostId(Output<String> hostId) {
            $.hostId = hostId;
            return this;
        }

        /**
         * @param hostId The [managed object reference ID][docs-about-morefs] of
         * a host.
         * 
         * @return builder
         * 
         */
        public Builder hostId(String hostId) {
            return hostId(Output.of(hostId));
        }

        /**
         * @param nameRegex A regular expression that will be used to match the
         * host PCI device name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regular expression that will be used to match the
         * host PCI device name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        /**
         * @param vendorId The hexadecimal PCI device vendor ID.
         * 
         * @return builder
         * 
         */
        public Builder vendorId(@Nullable Output<String> vendorId) {
            $.vendorId = vendorId;
            return this;
        }

        /**
         * @param vendorId The hexadecimal PCI device vendor ID.
         * 
         * @return builder
         * 
         */
        public Builder vendorId(String vendorId) {
            return vendorId(Output.of(vendorId));
        }

        public GetHostPciDeviceArgs build() {
            if ($.hostId == null) {
                throw new MissingRequiredPropertyException("GetHostPciDeviceArgs", "hostId");
            }
            return $;
        }
    }

}
