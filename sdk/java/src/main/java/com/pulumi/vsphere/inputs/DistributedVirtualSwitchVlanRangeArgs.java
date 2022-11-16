// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;


public final class DistributedVirtualSwitchVlanRangeArgs extends com.pulumi.resources.ResourceArgs {

    public static final DistributedVirtualSwitchVlanRangeArgs Empty = new DistributedVirtualSwitchVlanRangeArgs();

    @Import(name="maxVlan", required=true)
    private Output<Integer> maxVlan;

    public Output<Integer> maxVlan() {
        return this.maxVlan;
    }

    @Import(name="minVlan", required=true)
    private Output<Integer> minVlan;

    public Output<Integer> minVlan() {
        return this.minVlan;
    }

    private DistributedVirtualSwitchVlanRangeArgs() {}

    private DistributedVirtualSwitchVlanRangeArgs(DistributedVirtualSwitchVlanRangeArgs $) {
        this.maxVlan = $.maxVlan;
        this.minVlan = $.minVlan;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DistributedVirtualSwitchVlanRangeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DistributedVirtualSwitchVlanRangeArgs $;

        public Builder() {
            $ = new DistributedVirtualSwitchVlanRangeArgs();
        }

        public Builder(DistributedVirtualSwitchVlanRangeArgs defaults) {
            $ = new DistributedVirtualSwitchVlanRangeArgs(Objects.requireNonNull(defaults));
        }

        public Builder maxVlan(Output<Integer> maxVlan) {
            $.maxVlan = maxVlan;
            return this;
        }

        public Builder maxVlan(Integer maxVlan) {
            return maxVlan(Output.of(maxVlan));
        }

        public Builder minVlan(Output<Integer> minVlan) {
            $.minVlan = minVlan;
            return this;
        }

        public Builder minVlan(Integer minVlan) {
            return minVlan(Output.of(minVlan));
        }

        public DistributedVirtualSwitchVlanRangeArgs build() {
            $.maxVlan = Objects.requireNonNull($.maxVlan, "expected parameter 'maxVlan' to be non-null");
            $.minVlan = Objects.requireNonNull($.minVlan, "expected parameter 'minVlan' to be non-null");
            return $;
        }
    }

}
