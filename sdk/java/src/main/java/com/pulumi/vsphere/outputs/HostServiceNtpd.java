// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class HostServiceNtpd {
    /**
     * @return Whether the NTP service is enabled. Default is false.
     * 
     */
    private @Nullable Boolean enabled;
    private @Nullable List<String> ntpServers;
    /**
     * @return The policy for the NTP service. Valid values are &#39;Start and stop with host&#39;, &#39;Start and stop manually&#39;, &#39;Start and stop with port usage&#39;.
     * 
     */
    private @Nullable String policy;

    private HostServiceNtpd() {}
    /**
     * @return Whether the NTP service is enabled. Default is false.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    public List<String> ntpServers() {
        return this.ntpServers == null ? List.of() : this.ntpServers;
    }
    /**
     * @return The policy for the NTP service. Valid values are &#39;Start and stop with host&#39;, &#39;Start and stop manually&#39;, &#39;Start and stop with port usage&#39;.
     * 
     */
    public Optional<String> policy() {
        return Optional.ofNullable(this.policy);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(HostServiceNtpd defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enabled;
        private @Nullable List<String> ntpServers;
        private @Nullable String policy;
        public Builder() {}
        public Builder(HostServiceNtpd defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.ntpServers = defaults.ntpServers;
    	      this.policy = defaults.policy;
        }

        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {

            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder ntpServers(@Nullable List<String> ntpServers) {

            this.ntpServers = ntpServers;
            return this;
        }
        public Builder ntpServers(String... ntpServers) {
            return ntpServers(List.of(ntpServers));
        }
        @CustomType.Setter
        public Builder policy(@Nullable String policy) {

            this.policy = policy;
            return this;
        }
        public HostServiceNtpd build() {
            final var _resultValue = new HostServiceNtpd();
            _resultValue.enabled = enabled;
            _resultValue.ntpServers = ntpServers;
            _resultValue.policy = policy;
            return _resultValue;
        }
    }
}
