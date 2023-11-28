// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GuestOsCustomizationSpecWindowsOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GuestOsCustomizationSpecWindowsOptionsArgs Empty = new GuestOsCustomizationSpecWindowsOptionsArgs();

    @Import(name="adminPassword")
    private @Nullable Output<String> adminPassword;

    public Optional<Output<String>> adminPassword() {
        return Optional.ofNullable(this.adminPassword);
    }

    @Import(name="autoLogon")
    private @Nullable Output<Boolean> autoLogon;

    public Optional<Output<Boolean>> autoLogon() {
        return Optional.ofNullable(this.autoLogon);
    }

    @Import(name="autoLogonCount")
    private @Nullable Output<Integer> autoLogonCount;

    public Optional<Output<Integer>> autoLogonCount() {
        return Optional.ofNullable(this.autoLogonCount);
    }

    @Import(name="computerName", required=true)
    private Output<String> computerName;

    public Output<String> computerName() {
        return this.computerName;
    }

    @Import(name="domainAdminPassword")
    private @Nullable Output<String> domainAdminPassword;

    public Optional<Output<String>> domainAdminPassword() {
        return Optional.ofNullable(this.domainAdminPassword);
    }

    @Import(name="domainAdminUser")
    private @Nullable Output<String> domainAdminUser;

    public Optional<Output<String>> domainAdminUser() {
        return Optional.ofNullable(this.domainAdminUser);
    }

    @Import(name="fullName")
    private @Nullable Output<String> fullName;

    public Optional<Output<String>> fullName() {
        return Optional.ofNullable(this.fullName);
    }

    @Import(name="joinDomain")
    private @Nullable Output<String> joinDomain;

    public Optional<Output<String>> joinDomain() {
        return Optional.ofNullable(this.joinDomain);
    }

    @Import(name="organizationName")
    private @Nullable Output<String> organizationName;

    public Optional<Output<String>> organizationName() {
        return Optional.ofNullable(this.organizationName);
    }

    @Import(name="productKey")
    private @Nullable Output<String> productKey;

    public Optional<Output<String>> productKey() {
        return Optional.ofNullable(this.productKey);
    }

    @Import(name="runOnceCommandLists")
    private @Nullable Output<List<String>> runOnceCommandLists;

    public Optional<Output<List<String>>> runOnceCommandLists() {
        return Optional.ofNullable(this.runOnceCommandLists);
    }

    @Import(name="timeZone")
    private @Nullable Output<Integer> timeZone;

    public Optional<Output<Integer>> timeZone() {
        return Optional.ofNullable(this.timeZone);
    }

    @Import(name="workgroup")
    private @Nullable Output<String> workgroup;

    public Optional<Output<String>> workgroup() {
        return Optional.ofNullable(this.workgroup);
    }

    private GuestOsCustomizationSpecWindowsOptionsArgs() {}

    private GuestOsCustomizationSpecWindowsOptionsArgs(GuestOsCustomizationSpecWindowsOptionsArgs $) {
        this.adminPassword = $.adminPassword;
        this.autoLogon = $.autoLogon;
        this.autoLogonCount = $.autoLogonCount;
        this.computerName = $.computerName;
        this.domainAdminPassword = $.domainAdminPassword;
        this.domainAdminUser = $.domainAdminUser;
        this.fullName = $.fullName;
        this.joinDomain = $.joinDomain;
        this.organizationName = $.organizationName;
        this.productKey = $.productKey;
        this.runOnceCommandLists = $.runOnceCommandLists;
        this.timeZone = $.timeZone;
        this.workgroup = $.workgroup;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GuestOsCustomizationSpecWindowsOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GuestOsCustomizationSpecWindowsOptionsArgs $;

        public Builder() {
            $ = new GuestOsCustomizationSpecWindowsOptionsArgs();
        }

        public Builder(GuestOsCustomizationSpecWindowsOptionsArgs defaults) {
            $ = new GuestOsCustomizationSpecWindowsOptionsArgs(Objects.requireNonNull(defaults));
        }

        public Builder adminPassword(@Nullable Output<String> adminPassword) {
            $.adminPassword = adminPassword;
            return this;
        }

        public Builder adminPassword(String adminPassword) {
            return adminPassword(Output.of(adminPassword));
        }

        public Builder autoLogon(@Nullable Output<Boolean> autoLogon) {
            $.autoLogon = autoLogon;
            return this;
        }

        public Builder autoLogon(Boolean autoLogon) {
            return autoLogon(Output.of(autoLogon));
        }

        public Builder autoLogonCount(@Nullable Output<Integer> autoLogonCount) {
            $.autoLogonCount = autoLogonCount;
            return this;
        }

        public Builder autoLogonCount(Integer autoLogonCount) {
            return autoLogonCount(Output.of(autoLogonCount));
        }

        public Builder computerName(Output<String> computerName) {
            $.computerName = computerName;
            return this;
        }

        public Builder computerName(String computerName) {
            return computerName(Output.of(computerName));
        }

        public Builder domainAdminPassword(@Nullable Output<String> domainAdminPassword) {
            $.domainAdminPassword = domainAdminPassword;
            return this;
        }

        public Builder domainAdminPassword(String domainAdminPassword) {
            return domainAdminPassword(Output.of(domainAdminPassword));
        }

        public Builder domainAdminUser(@Nullable Output<String> domainAdminUser) {
            $.domainAdminUser = domainAdminUser;
            return this;
        }

        public Builder domainAdminUser(String domainAdminUser) {
            return domainAdminUser(Output.of(domainAdminUser));
        }

        public Builder fullName(@Nullable Output<String> fullName) {
            $.fullName = fullName;
            return this;
        }

        public Builder fullName(String fullName) {
            return fullName(Output.of(fullName));
        }

        public Builder joinDomain(@Nullable Output<String> joinDomain) {
            $.joinDomain = joinDomain;
            return this;
        }

        public Builder joinDomain(String joinDomain) {
            return joinDomain(Output.of(joinDomain));
        }

        public Builder organizationName(@Nullable Output<String> organizationName) {
            $.organizationName = organizationName;
            return this;
        }

        public Builder organizationName(String organizationName) {
            return organizationName(Output.of(organizationName));
        }

        public Builder productKey(@Nullable Output<String> productKey) {
            $.productKey = productKey;
            return this;
        }

        public Builder productKey(String productKey) {
            return productKey(Output.of(productKey));
        }

        public Builder runOnceCommandLists(@Nullable Output<List<String>> runOnceCommandLists) {
            $.runOnceCommandLists = runOnceCommandLists;
            return this;
        }

        public Builder runOnceCommandLists(List<String> runOnceCommandLists) {
            return runOnceCommandLists(Output.of(runOnceCommandLists));
        }

        public Builder runOnceCommandLists(String... runOnceCommandLists) {
            return runOnceCommandLists(List.of(runOnceCommandLists));
        }

        public Builder timeZone(@Nullable Output<Integer> timeZone) {
            $.timeZone = timeZone;
            return this;
        }

        public Builder timeZone(Integer timeZone) {
            return timeZone(Output.of(timeZone));
        }

        public Builder workgroup(@Nullable Output<String> workgroup) {
            $.workgroup = workgroup;
            return this;
        }

        public Builder workgroup(String workgroup) {
            return workgroup(Output.of(workgroup));
        }

        public GuestOsCustomizationSpecWindowsOptionsArgs build() {
            $.computerName = Objects.requireNonNull($.computerName, "expected parameter 'computerName' to be non-null");
            return $;
        }
    }

}
