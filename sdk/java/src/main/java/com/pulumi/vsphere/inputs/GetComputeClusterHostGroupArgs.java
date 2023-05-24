// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetComputeClusterHostGroupArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetComputeClusterHostGroupArgs Empty = new GetComputeClusterHostGroupArgs();

    /**
     * The [managed object reference ID][docs-about-morefs]
     * of the compute cluster for the host group.
     * 
     * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
     * 
     */
    @Import(name="computeClusterId", required=true)
    private Output<String> computeClusterId;

    /**
     * @return The [managed object reference ID][docs-about-morefs]
     * of the compute cluster for the host group.
     * 
     * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
     * 
     */
    public Output<String> computeClusterId() {
        return this.computeClusterId;
    }

    /**
     * The name of the host group.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the host group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetComputeClusterHostGroupArgs() {}

    private GetComputeClusterHostGroupArgs(GetComputeClusterHostGroupArgs $) {
        this.computeClusterId = $.computeClusterId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetComputeClusterHostGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetComputeClusterHostGroupArgs $;

        public Builder() {
            $ = new GetComputeClusterHostGroupArgs();
        }

        public Builder(GetComputeClusterHostGroupArgs defaults) {
            $ = new GetComputeClusterHostGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param computeClusterId The [managed object reference ID][docs-about-morefs]
         * of the compute cluster for the host group.
         * 
         * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
         * 
         * @return builder
         * 
         */
        public Builder computeClusterId(Output<String> computeClusterId) {
            $.computeClusterId = computeClusterId;
            return this;
        }

        /**
         * @param computeClusterId The [managed object reference ID][docs-about-morefs]
         * of the compute cluster for the host group.
         * 
         * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
         * 
         * @return builder
         * 
         */
        public Builder computeClusterId(String computeClusterId) {
            return computeClusterId(Output.of(computeClusterId));
        }

        /**
         * @param name The name of the host group.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the host group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetComputeClusterHostGroupArgs build() {
            $.computeClusterId = Objects.requireNonNull($.computeClusterId, "expected parameter 'computeClusterId' to be non-null");
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
