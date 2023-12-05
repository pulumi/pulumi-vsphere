// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.vsphere.outputs.GetGuestOsCustomizationSpecLinuxOption;
import com.pulumi.vsphere.outputs.GetGuestOsCustomizationSpecNetworkInterface;
import com.pulumi.vsphere.outputs.GetGuestOsCustomizationSpecWindowsOption;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetGuestOsCustomizationSpec {
    private List<String> dnsServerLists;
    private List<String> dnsSuffixLists;
    private List<GetGuestOsCustomizationSpecLinuxOption> linuxOptions;
    private List<GetGuestOsCustomizationSpecNetworkInterface> networkInterfaces;
    private List<GetGuestOsCustomizationSpecWindowsOption> windowsOptions;
    private String windowsSysprepText;

    private GetGuestOsCustomizationSpec() {}
    public List<String> dnsServerLists() {
        return this.dnsServerLists;
    }
    public List<String> dnsSuffixLists() {
        return this.dnsSuffixLists;
    }
    public List<GetGuestOsCustomizationSpecLinuxOption> linuxOptions() {
        return this.linuxOptions;
    }
    public List<GetGuestOsCustomizationSpecNetworkInterface> networkInterfaces() {
        return this.networkInterfaces;
    }
    public List<GetGuestOsCustomizationSpecWindowsOption> windowsOptions() {
        return this.windowsOptions;
    }
    public String windowsSysprepText() {
        return this.windowsSysprepText;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGuestOsCustomizationSpec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> dnsServerLists;
        private List<String> dnsSuffixLists;
        private List<GetGuestOsCustomizationSpecLinuxOption> linuxOptions;
        private List<GetGuestOsCustomizationSpecNetworkInterface> networkInterfaces;
        private List<GetGuestOsCustomizationSpecWindowsOption> windowsOptions;
        private String windowsSysprepText;
        public Builder() {}
        public Builder(GetGuestOsCustomizationSpec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dnsServerLists = defaults.dnsServerLists;
    	      this.dnsSuffixLists = defaults.dnsSuffixLists;
    	      this.linuxOptions = defaults.linuxOptions;
    	      this.networkInterfaces = defaults.networkInterfaces;
    	      this.windowsOptions = defaults.windowsOptions;
    	      this.windowsSysprepText = defaults.windowsSysprepText;
        }

        @CustomType.Setter
        public Builder dnsServerLists(List<String> dnsServerLists) {
            this.dnsServerLists = Objects.requireNonNull(dnsServerLists);
            return this;
        }
        public Builder dnsServerLists(String... dnsServerLists) {
            return dnsServerLists(List.of(dnsServerLists));
        }
        @CustomType.Setter
        public Builder dnsSuffixLists(List<String> dnsSuffixLists) {
            this.dnsSuffixLists = Objects.requireNonNull(dnsSuffixLists);
            return this;
        }
        public Builder dnsSuffixLists(String... dnsSuffixLists) {
            return dnsSuffixLists(List.of(dnsSuffixLists));
        }
        @CustomType.Setter
        public Builder linuxOptions(List<GetGuestOsCustomizationSpecLinuxOption> linuxOptions) {
            this.linuxOptions = Objects.requireNonNull(linuxOptions);
            return this;
        }
        public Builder linuxOptions(GetGuestOsCustomizationSpecLinuxOption... linuxOptions) {
            return linuxOptions(List.of(linuxOptions));
        }
        @CustomType.Setter
        public Builder networkInterfaces(List<GetGuestOsCustomizationSpecNetworkInterface> networkInterfaces) {
            this.networkInterfaces = Objects.requireNonNull(networkInterfaces);
            return this;
        }
        public Builder networkInterfaces(GetGuestOsCustomizationSpecNetworkInterface... networkInterfaces) {
            return networkInterfaces(List.of(networkInterfaces));
        }
        @CustomType.Setter
        public Builder windowsOptions(List<GetGuestOsCustomizationSpecWindowsOption> windowsOptions) {
            this.windowsOptions = Objects.requireNonNull(windowsOptions);
            return this;
        }
        public Builder windowsOptions(GetGuestOsCustomizationSpecWindowsOption... windowsOptions) {
            return windowsOptions(List.of(windowsOptions));
        }
        @CustomType.Setter
        public Builder windowsSysprepText(String windowsSysprepText) {
            this.windowsSysprepText = Objects.requireNonNull(windowsSysprepText);
            return this;
        }
        public GetGuestOsCustomizationSpec build() {
            final var o = new GetGuestOsCustomizationSpec();
            o.dnsServerLists = dnsServerLists;
            o.dnsSuffixLists = dnsSuffixLists;
            o.linuxOptions = linuxOptions;
            o.networkInterfaces = networkInterfaces;
            o.windowsOptions = windowsOptions;
            o.windowsSysprepText = windowsSysprepText;
            return o;
        }
    }
}