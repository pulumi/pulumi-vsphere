// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGuestOsCustomizationSpecLinuxOption {
    private String domain;
    private String hostName;
    private Boolean hwClockUtc;
    private String scriptText;
    private String timeZone;

    private GetGuestOsCustomizationSpecLinuxOption() {}
    public String domain() {
        return this.domain;
    }
    public String hostName() {
        return this.hostName;
    }
    public Boolean hwClockUtc() {
        return this.hwClockUtc;
    }
    public String scriptText() {
        return this.scriptText;
    }
    public String timeZone() {
        return this.timeZone;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGuestOsCustomizationSpecLinuxOption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String domain;
        private String hostName;
        private Boolean hwClockUtc;
        private String scriptText;
        private String timeZone;
        public Builder() {}
        public Builder(GetGuestOsCustomizationSpecLinuxOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.domain = defaults.domain;
    	      this.hostName = defaults.hostName;
    	      this.hwClockUtc = defaults.hwClockUtc;
    	      this.scriptText = defaults.scriptText;
    	      this.timeZone = defaults.timeZone;
        }

        @CustomType.Setter
        public Builder domain(String domain) {
            this.domain = Objects.requireNonNull(domain);
            return this;
        }
        @CustomType.Setter
        public Builder hostName(String hostName) {
            this.hostName = Objects.requireNonNull(hostName);
            return this;
        }
        @CustomType.Setter
        public Builder hwClockUtc(Boolean hwClockUtc) {
            this.hwClockUtc = Objects.requireNonNull(hwClockUtc);
            return this;
        }
        @CustomType.Setter
        public Builder scriptText(String scriptText) {
            this.scriptText = Objects.requireNonNull(scriptText);
            return this;
        }
        @CustomType.Setter
        public Builder timeZone(String timeZone) {
            this.timeZone = Objects.requireNonNull(timeZone);
            return this;
        }
        public GetGuestOsCustomizationSpecLinuxOption build() {
            final var _resultValue = new GetGuestOsCustomizationSpecLinuxOption();
            _resultValue.domain = domain;
            _resultValue.hostName = hostName;
            _resultValue.hwClockUtc = hwClockUtc;
            _resultValue.scriptText = scriptText;
            _resultValue.timeZone = timeZone;
            return _resultValue;
        }
    }
}
