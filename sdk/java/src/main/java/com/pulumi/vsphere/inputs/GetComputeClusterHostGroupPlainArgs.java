// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetComputeClusterHostGroupPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetComputeClusterHostGroupPlainArgs Empty = new GetComputeClusterHostGroupPlainArgs();

    /**
     * The
     * [managed object reference ID][docs-about-morefs] of the compute cluster for
     * the host group.
     * 
     * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
     * 
     */
    @Import(name="computeClusterId", required=true)
    private String computeClusterId;

    /**
     * @return The
     * [managed object reference ID][docs-about-morefs] of the compute cluster for
     * the host group.
     * 
     * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
     * 
     */
    public String computeClusterId() {
        return this.computeClusterId;
    }

    /**
     * The name of the host group.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The name of the host group.
     * 
     */
    public String name() {
        return this.name;
    }

    private GetComputeClusterHostGroupPlainArgs() {}

    private GetComputeClusterHostGroupPlainArgs(GetComputeClusterHostGroupPlainArgs $) {
        this.computeClusterId = $.computeClusterId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetComputeClusterHostGroupPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetComputeClusterHostGroupPlainArgs $;

        public Builder() {
            $ = new GetComputeClusterHostGroupPlainArgs();
        }

        public Builder(GetComputeClusterHostGroupPlainArgs defaults) {
            $ = new GetComputeClusterHostGroupPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param computeClusterId The
         * [managed object reference ID][docs-about-morefs] of the compute cluster for
         * the host group.
         * 
         * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
         * 
         * @return builder
         * 
         */
        public Builder computeClusterId(String computeClusterId) {
            $.computeClusterId = computeClusterId;
            return this;
        }

        /**
         * @param name The name of the host group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetComputeClusterHostGroupPlainArgs build() {
            if ($.computeClusterId == null) {
                throw new MissingRequiredPropertyException("GetComputeClusterHostGroupPlainArgs", "computeClusterId");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetComputeClusterHostGroupPlainArgs", "name");
            }
            return $;
        }
    }

}
