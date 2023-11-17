// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetHostPciDevicePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetHostPciDevicePlainArgs Empty = new GetHostPciDevicePlainArgs();

    /**
     * The hexadecimal PCI device class ID
     * 
     * &gt; **NOTE:** `name_regex`, `vendor_id`, and `class_id` can all be used together.
     * 
     */
    @Import(name="classId")
    private @Nullable String classId;

    /**
     * @return The hexadecimal PCI device class ID
     * 
     * &gt; **NOTE:** `name_regex`, `vendor_id`, and `class_id` can all be used together.
     * 
     */
    public Optional<String> classId() {
        return Optional.ofNullable(this.classId);
    }

    /**
     * The [managed object reference ID][docs-about-morefs] of a host.
     * 
     */
    @Import(name="hostId", required=true)
    private String hostId;

    /**
     * @return The [managed object reference ID][docs-about-morefs] of a host.
     * 
     */
    public String hostId() {
        return this.hostId;
    }

    /**
     * A regular expression that will be used to match the
     * host PCI device name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regular expression that will be used to match the
     * host PCI device name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * The hexadecimal PCI device vendor ID.
     * 
     */
    @Import(name="vendorId")
    private @Nullable String vendorId;

    /**
     * @return The hexadecimal PCI device vendor ID.
     * 
     */
    public Optional<String> vendorId() {
        return Optional.ofNullable(this.vendorId);
    }

    private GetHostPciDevicePlainArgs() {}

    private GetHostPciDevicePlainArgs(GetHostPciDevicePlainArgs $) {
        this.classId = $.classId;
        this.hostId = $.hostId;
        this.nameRegex = $.nameRegex;
        this.vendorId = $.vendorId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetHostPciDevicePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetHostPciDevicePlainArgs $;

        public Builder() {
            $ = new GetHostPciDevicePlainArgs();
        }

        public Builder(GetHostPciDevicePlainArgs defaults) {
            $ = new GetHostPciDevicePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param classId The hexadecimal PCI device class ID
         * 
         * &gt; **NOTE:** `name_regex`, `vendor_id`, and `class_id` can all be used together.
         * 
         * @return builder
         * 
         */
        public Builder classId(@Nullable String classId) {
            $.classId = classId;
            return this;
        }

        /**
         * @param hostId The [managed object reference ID][docs-about-morefs] of a host.
         * 
         * @return builder
         * 
         */
        public Builder hostId(String hostId) {
            $.hostId = hostId;
            return this;
        }

        /**
         * @param nameRegex A regular expression that will be used to match the
         * host PCI device name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param vendorId The hexadecimal PCI device vendor ID.
         * 
         * @return builder
         * 
         */
        public Builder vendorId(@Nullable String vendorId) {
            $.vendorId = vendorId;
            return this;
        }

        public GetHostPciDevicePlainArgs build() {
            $.hostId = Objects.requireNonNull($.hostId, "expected parameter 'hostId' to be non-null");
            return $;
        }
    }

}
