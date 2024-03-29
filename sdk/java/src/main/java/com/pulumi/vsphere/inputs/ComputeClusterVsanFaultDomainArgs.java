// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.vsphere.inputs.ComputeClusterVsanFaultDomainFaultDomainArgs;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ComputeClusterVsanFaultDomainArgs extends com.pulumi.resources.ResourceArgs {

    public static final ComputeClusterVsanFaultDomainArgs Empty = new ComputeClusterVsanFaultDomainArgs();

    /**
     * The configuration for single fault domain.
     * 
     */
    @Import(name="faultDomains")
    private @Nullable Output<List<ComputeClusterVsanFaultDomainFaultDomainArgs>> faultDomains;

    /**
     * @return The configuration for single fault domain.
     * 
     */
    public Optional<Output<List<ComputeClusterVsanFaultDomainFaultDomainArgs>>> faultDomains() {
        return Optional.ofNullable(this.faultDomains);
    }

    private ComputeClusterVsanFaultDomainArgs() {}

    private ComputeClusterVsanFaultDomainArgs(ComputeClusterVsanFaultDomainArgs $) {
        this.faultDomains = $.faultDomains;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ComputeClusterVsanFaultDomainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ComputeClusterVsanFaultDomainArgs $;

        public Builder() {
            $ = new ComputeClusterVsanFaultDomainArgs();
        }

        public Builder(ComputeClusterVsanFaultDomainArgs defaults) {
            $ = new ComputeClusterVsanFaultDomainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param faultDomains The configuration for single fault domain.
         * 
         * @return builder
         * 
         */
        public Builder faultDomains(@Nullable Output<List<ComputeClusterVsanFaultDomainFaultDomainArgs>> faultDomains) {
            $.faultDomains = faultDomains;
            return this;
        }

        /**
         * @param faultDomains The configuration for single fault domain.
         * 
         * @return builder
         * 
         */
        public Builder faultDomains(List<ComputeClusterVsanFaultDomainFaultDomainArgs> faultDomains) {
            return faultDomains(Output.of(faultDomains));
        }

        /**
         * @param faultDomains The configuration for single fault domain.
         * 
         * @return builder
         * 
         */
        public Builder faultDomains(ComputeClusterVsanFaultDomainFaultDomainArgs... faultDomains) {
            return faultDomains(List.of(faultDomains));
        }

        public ComputeClusterVsanFaultDomainArgs build() {
            return $;
        }
    }

}
