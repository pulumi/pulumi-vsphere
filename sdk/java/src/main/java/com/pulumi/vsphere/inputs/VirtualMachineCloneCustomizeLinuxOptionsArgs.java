// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualMachineCloneCustomizeLinuxOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final VirtualMachineCloneCustomizeLinuxOptionsArgs Empty = new VirtualMachineCloneCustomizeLinuxOptionsArgs();

    @Import(name="domain", required=true)
    private Output<String> domain;

    public Output<String> domain() {
        return this.domain;
    }

    @Import(name="hostName", required=true)
    private Output<String> hostName;

    public Output<String> hostName() {
        return this.hostName;
    }

    @Import(name="hwClockUtc")
    private @Nullable Output<Boolean> hwClockUtc;

    public Optional<Output<Boolean>> hwClockUtc() {
        return Optional.ofNullable(this.hwClockUtc);
    }

    @Import(name="scriptText")
    private @Nullable Output<String> scriptText;

    public Optional<Output<String>> scriptText() {
        return Optional.ofNullable(this.scriptText);
    }

    @Import(name="timeZone")
    private @Nullable Output<String> timeZone;

    public Optional<Output<String>> timeZone() {
        return Optional.ofNullable(this.timeZone);
    }

    private VirtualMachineCloneCustomizeLinuxOptionsArgs() {}

    private VirtualMachineCloneCustomizeLinuxOptionsArgs(VirtualMachineCloneCustomizeLinuxOptionsArgs $) {
        this.domain = $.domain;
        this.hostName = $.hostName;
        this.hwClockUtc = $.hwClockUtc;
        this.scriptText = $.scriptText;
        this.timeZone = $.timeZone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualMachineCloneCustomizeLinuxOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualMachineCloneCustomizeLinuxOptionsArgs $;

        public Builder() {
            $ = new VirtualMachineCloneCustomizeLinuxOptionsArgs();
        }

        public Builder(VirtualMachineCloneCustomizeLinuxOptionsArgs defaults) {
            $ = new VirtualMachineCloneCustomizeLinuxOptionsArgs(Objects.requireNonNull(defaults));
        }

        public Builder domain(Output<String> domain) {
            $.domain = domain;
            return this;
        }

        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        public Builder hostName(Output<String> hostName) {
            $.hostName = hostName;
            return this;
        }

        public Builder hostName(String hostName) {
            return hostName(Output.of(hostName));
        }

        public Builder hwClockUtc(@Nullable Output<Boolean> hwClockUtc) {
            $.hwClockUtc = hwClockUtc;
            return this;
        }

        public Builder hwClockUtc(Boolean hwClockUtc) {
            return hwClockUtc(Output.of(hwClockUtc));
        }

        public Builder scriptText(@Nullable Output<String> scriptText) {
            $.scriptText = scriptText;
            return this;
        }

        public Builder scriptText(String scriptText) {
            return scriptText(Output.of(scriptText));
        }

        public Builder timeZone(@Nullable Output<String> timeZone) {
            $.timeZone = timeZone;
            return this;
        }

        public Builder timeZone(String timeZone) {
            return timeZone(Output.of(timeZone));
        }

        public VirtualMachineCloneCustomizeLinuxOptionsArgs build() {
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            $.hostName = Objects.requireNonNull($.hostName, "expected parameter 'hostName' to be non-null");
            return $;
        }
    }

}
