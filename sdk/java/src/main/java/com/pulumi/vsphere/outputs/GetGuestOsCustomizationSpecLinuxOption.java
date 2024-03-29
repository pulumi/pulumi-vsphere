// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGuestOsCustomizationSpecLinuxOption {
    /**
     * @return The domain name for this virtual machine.
     * 
     */
    private String domain;
    /**
     * @return The hostname for this virtual machine.
     * 
     */
    private String hostName;
    /**
     * @return Specifies whether or not the hardware clock should be in UTC or not.
     * 
     */
    private Boolean hwClockUtc;
    /**
     * @return The customization script to run before and or after guest customization.
     * 
     */
    private String scriptText;
    /**
     * @return Set the time zone on the guest operating system. For a list of the acceptable values for Linux customization specifications, see [List of Time Zone Database Zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) on Wikipedia.
     * 
     */
    private String timeZone;

    private GetGuestOsCustomizationSpecLinuxOption() {}
    /**
     * @return The domain name for this virtual machine.
     * 
     */
    public String domain() {
        return this.domain;
    }
    /**
     * @return The hostname for this virtual machine.
     * 
     */
    public String hostName() {
        return this.hostName;
    }
    /**
     * @return Specifies whether or not the hardware clock should be in UTC or not.
     * 
     */
    public Boolean hwClockUtc() {
        return this.hwClockUtc;
    }
    /**
     * @return The customization script to run before and or after guest customization.
     * 
     */
    public String scriptText() {
        return this.scriptText;
    }
    /**
     * @return Set the time zone on the guest operating system. For a list of the acceptable values for Linux customization specifications, see [List of Time Zone Database Zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) on Wikipedia.
     * 
     */
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
            if (domain == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecLinuxOption", "domain");
            }
            this.domain = domain;
            return this;
        }
        @CustomType.Setter
        public Builder hostName(String hostName) {
            if (hostName == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecLinuxOption", "hostName");
            }
            this.hostName = hostName;
            return this;
        }
        @CustomType.Setter
        public Builder hwClockUtc(Boolean hwClockUtc) {
            if (hwClockUtc == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecLinuxOption", "hwClockUtc");
            }
            this.hwClockUtc = hwClockUtc;
            return this;
        }
        @CustomType.Setter
        public Builder scriptText(String scriptText) {
            if (scriptText == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecLinuxOption", "scriptText");
            }
            this.scriptText = scriptText;
            return this;
        }
        @CustomType.Setter
        public Builder timeZone(String timeZone) {
            if (timeZone == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecLinuxOption", "timeZone");
            }
            this.timeZone = timeZone;
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
