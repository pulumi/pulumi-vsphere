// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetGuestOsCustomizationSpecWindowsOption {
    /**
     * @return The new administrator password for this virtual machine.
     * 
     */
    private String adminPassword;
    /**
     * @return Specifies whether or not the guest operating system automatically logs on as Administrator.
     * 
     */
    private Boolean autoLogon;
    /**
     * @return Specifies how many times the guest operating system should auto-logon the Administrator account when `auto_logon` is `true`.
     * 
     */
    private Integer autoLogonCount;
    /**
     * @return The hostname for this virtual machine.
     * 
     */
    private String computerName;
    /**
     * @return The user account used to join this virtual machine to the Active Directory domain.
     * 
     */
    private @Nullable String domainAdminPassword;
    /**
     * @return The user account of the domain administrator used to join this virtual machine to the domain.
     * 
     */
    private String domainAdminUser;
    /**
     * @return The MachineObjectOU which specifies the full LDAP path name of the OU to which the virtual machine belongs.
     * 
     */
    private String domainOu;
    /**
     * @return The full name of the user of this virtual machine.
     * 
     */
    private String fullName;
    /**
     * @return The Active Directory domain for the virtual machine to join.
     * 
     */
    private String joinDomain;
    /**
     * @return The organization name this virtual machine is being installed for.
     * 
     */
    private String organizationName;
    /**
     * @return The product key for this virtual machine.
     * 
     */
    private String productKey;
    /**
     * @return A list of commands to run at first user logon, after guest customization.
     * 
     */
    private List<String> runOnceCommandLists;
    /**
     * @return The new time zone for the virtual machine. This is a sysprep-dictated timezone code.
     * 
     */
    private Integer timeZone;
    /**
     * @return The workgroup for this virtual machine if not joining an Active Directory domain.
     * 
     */
    private String workgroup;

    private GetGuestOsCustomizationSpecWindowsOption() {}
    /**
     * @return The new administrator password for this virtual machine.
     * 
     */
    public String adminPassword() {
        return this.adminPassword;
    }
    /**
     * @return Specifies whether or not the guest operating system automatically logs on as Administrator.
     * 
     */
    public Boolean autoLogon() {
        return this.autoLogon;
    }
    /**
     * @return Specifies how many times the guest operating system should auto-logon the Administrator account when `auto_logon` is `true`.
     * 
     */
    public Integer autoLogonCount() {
        return this.autoLogonCount;
    }
    /**
     * @return The hostname for this virtual machine.
     * 
     */
    public String computerName() {
        return this.computerName;
    }
    /**
     * @return The user account used to join this virtual machine to the Active Directory domain.
     * 
     */
    public Optional<String> domainAdminPassword() {
        return Optional.ofNullable(this.domainAdminPassword);
    }
    /**
     * @return The user account of the domain administrator used to join this virtual machine to the domain.
     * 
     */
    public String domainAdminUser() {
        return this.domainAdminUser;
    }
    /**
     * @return The MachineObjectOU which specifies the full LDAP path name of the OU to which the virtual machine belongs.
     * 
     */
    public String domainOu() {
        return this.domainOu;
    }
    /**
     * @return The full name of the user of this virtual machine.
     * 
     */
    public String fullName() {
        return this.fullName;
    }
    /**
     * @return The Active Directory domain for the virtual machine to join.
     * 
     */
    public String joinDomain() {
        return this.joinDomain;
    }
    /**
     * @return The organization name this virtual machine is being installed for.
     * 
     */
    public String organizationName() {
        return this.organizationName;
    }
    /**
     * @return The product key for this virtual machine.
     * 
     */
    public String productKey() {
        return this.productKey;
    }
    /**
     * @return A list of commands to run at first user logon, after guest customization.
     * 
     */
    public List<String> runOnceCommandLists() {
        return this.runOnceCommandLists;
    }
    /**
     * @return The new time zone for the virtual machine. This is a sysprep-dictated timezone code.
     * 
     */
    public Integer timeZone() {
        return this.timeZone;
    }
    /**
     * @return The workgroup for this virtual machine if not joining an Active Directory domain.
     * 
     */
    public String workgroup() {
        return this.workgroup;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGuestOsCustomizationSpecWindowsOption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String adminPassword;
        private Boolean autoLogon;
        private Integer autoLogonCount;
        private String computerName;
        private @Nullable String domainAdminPassword;
        private String domainAdminUser;
        private String domainOu;
        private String fullName;
        private String joinDomain;
        private String organizationName;
        private String productKey;
        private List<String> runOnceCommandLists;
        private Integer timeZone;
        private String workgroup;
        public Builder() {}
        public Builder(GetGuestOsCustomizationSpecWindowsOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminPassword = defaults.adminPassword;
    	      this.autoLogon = defaults.autoLogon;
    	      this.autoLogonCount = defaults.autoLogonCount;
    	      this.computerName = defaults.computerName;
    	      this.domainAdminPassword = defaults.domainAdminPassword;
    	      this.domainAdminUser = defaults.domainAdminUser;
    	      this.domainOu = defaults.domainOu;
    	      this.fullName = defaults.fullName;
    	      this.joinDomain = defaults.joinDomain;
    	      this.organizationName = defaults.organizationName;
    	      this.productKey = defaults.productKey;
    	      this.runOnceCommandLists = defaults.runOnceCommandLists;
    	      this.timeZone = defaults.timeZone;
    	      this.workgroup = defaults.workgroup;
        }

        @CustomType.Setter
        public Builder adminPassword(String adminPassword) {
            if (adminPassword == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "adminPassword");
            }
            this.adminPassword = adminPassword;
            return this;
        }
        @CustomType.Setter
        public Builder autoLogon(Boolean autoLogon) {
            if (autoLogon == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "autoLogon");
            }
            this.autoLogon = autoLogon;
            return this;
        }
        @CustomType.Setter
        public Builder autoLogonCount(Integer autoLogonCount) {
            if (autoLogonCount == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "autoLogonCount");
            }
            this.autoLogonCount = autoLogonCount;
            return this;
        }
        @CustomType.Setter
        public Builder computerName(String computerName) {
            if (computerName == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "computerName");
            }
            this.computerName = computerName;
            return this;
        }
        @CustomType.Setter
        public Builder domainAdminPassword(@Nullable String domainAdminPassword) {

            this.domainAdminPassword = domainAdminPassword;
            return this;
        }
        @CustomType.Setter
        public Builder domainAdminUser(String domainAdminUser) {
            if (domainAdminUser == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "domainAdminUser");
            }
            this.domainAdminUser = domainAdminUser;
            return this;
        }
        @CustomType.Setter
        public Builder domainOu(String domainOu) {
            if (domainOu == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "domainOu");
            }
            this.domainOu = domainOu;
            return this;
        }
        @CustomType.Setter
        public Builder fullName(String fullName) {
            if (fullName == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "fullName");
            }
            this.fullName = fullName;
            return this;
        }
        @CustomType.Setter
        public Builder joinDomain(String joinDomain) {
            if (joinDomain == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "joinDomain");
            }
            this.joinDomain = joinDomain;
            return this;
        }
        @CustomType.Setter
        public Builder organizationName(String organizationName) {
            if (organizationName == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "organizationName");
            }
            this.organizationName = organizationName;
            return this;
        }
        @CustomType.Setter
        public Builder productKey(String productKey) {
            if (productKey == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "productKey");
            }
            this.productKey = productKey;
            return this;
        }
        @CustomType.Setter
        public Builder runOnceCommandLists(List<String> runOnceCommandLists) {
            if (runOnceCommandLists == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "runOnceCommandLists");
            }
            this.runOnceCommandLists = runOnceCommandLists;
            return this;
        }
        public Builder runOnceCommandLists(String... runOnceCommandLists) {
            return runOnceCommandLists(List.of(runOnceCommandLists));
        }
        @CustomType.Setter
        public Builder timeZone(Integer timeZone) {
            if (timeZone == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "timeZone");
            }
            this.timeZone = timeZone;
            return this;
        }
        @CustomType.Setter
        public Builder workgroup(String workgroup) {
            if (workgroup == null) {
              throw new MissingRequiredPropertyException("GetGuestOsCustomizationSpecWindowsOption", "workgroup");
            }
            this.workgroup = workgroup;
            return this;
        }
        public GetGuestOsCustomizationSpecWindowsOption build() {
            final var _resultValue = new GetGuestOsCustomizationSpecWindowsOption();
            _resultValue.adminPassword = adminPassword;
            _resultValue.autoLogon = autoLogon;
            _resultValue.autoLogonCount = autoLogonCount;
            _resultValue.computerName = computerName;
            _resultValue.domainAdminPassword = domainAdminPassword;
            _resultValue.domainAdminUser = domainAdminUser;
            _resultValue.domainOu = domainOu;
            _resultValue.fullName = fullName;
            _resultValue.joinDomain = joinDomain;
            _resultValue.organizationName = organizationName;
            _resultValue.productKey = productKey;
            _resultValue.runOnceCommandLists = runOnceCommandLists;
            _resultValue.timeZone = timeZone;
            _resultValue.workgroup = workgroup;
            return _resultValue;
        }
    }
}
