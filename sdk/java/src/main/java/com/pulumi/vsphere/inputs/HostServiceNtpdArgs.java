// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HostServiceNtpdArgs extends com.pulumi.resources.ResourceArgs {

    public static final HostServiceNtpdArgs Empty = new HostServiceNtpdArgs();

    /**
     * Whether the NTP service is enabled. Default is false.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Whether the NTP service is enabled. Default is false.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="ntpServers")
    private @Nullable Output<List<String>> ntpServers;

    public Optional<Output<List<String>>> ntpServers() {
        return Optional.ofNullable(this.ntpServers);
    }

    /**
     * The policy for the NTP service. Valid values are &#39;Start and stop with host&#39;, &#39;Start and stop manually&#39;, &#39;Start and stop with port usage&#39;.
     * 
     */
    @Import(name="policy")
    private @Nullable Output<String> policy;

    /**
     * @return The policy for the NTP service. Valid values are &#39;Start and stop with host&#39;, &#39;Start and stop manually&#39;, &#39;Start and stop with port usage&#39;.
     * 
     */
    public Optional<Output<String>> policy() {
        return Optional.ofNullable(this.policy);
    }

    private HostServiceNtpdArgs() {}

    private HostServiceNtpdArgs(HostServiceNtpdArgs $) {
        this.enabled = $.enabled;
        this.ntpServers = $.ntpServers;
        this.policy = $.policy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HostServiceNtpdArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HostServiceNtpdArgs $;

        public Builder() {
            $ = new HostServiceNtpdArgs();
        }

        public Builder(HostServiceNtpdArgs defaults) {
            $ = new HostServiceNtpdArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Whether the NTP service is enabled. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Whether the NTP service is enabled. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder ntpServers(@Nullable Output<List<String>> ntpServers) {
            $.ntpServers = ntpServers;
            return this;
        }

        public Builder ntpServers(List<String> ntpServers) {
            return ntpServers(Output.of(ntpServers));
        }

        public Builder ntpServers(String... ntpServers) {
            return ntpServers(List.of(ntpServers));
        }

        /**
         * @param policy The policy for the NTP service. Valid values are &#39;Start and stop with host&#39;, &#39;Start and stop manually&#39;, &#39;Start and stop with port usage&#39;.
         * 
         * @return builder
         * 
         */
        public Builder policy(@Nullable Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy The policy for the NTP service. Valid values are &#39;Start and stop with host&#39;, &#39;Start and stop manually&#39;, &#39;Start and stop with port usage&#39;.
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        public HostServiceNtpdArgs build() {
            return $;
        }
    }

}
