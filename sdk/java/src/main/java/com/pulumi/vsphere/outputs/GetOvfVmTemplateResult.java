// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetOvfVmTemplateResult {
    private @Nullable Boolean allowUnverifiedSslCert;
    /**
     * @return An alternate guest operating system name.
     * 
     */
    private String alternateGuestName;
    /**
     * @return A description of the virtual machine.
     * 
     */
    private String annotation;
    /**
     * @return Allow CPUs to be added to the virtual machine while
     * powered on.
     * 
     */
    private Boolean cpuHotAddEnabled;
    /**
     * @return Allow CPUs to be removed from the virtual machine
     * while powered on.
     * 
     */
    private Boolean cpuHotRemoveEnabled;
    private Boolean cpuPerformanceCountersEnabled;
    private @Nullable String datastoreId;
    private @Nullable String deploymentOption;
    private @Nullable String diskProvisioning;
    private @Nullable Boolean enableHiddenProperties;
    /**
     * @return The firmware to use on the virtual machine.
     * 
     */
    private String firmware;
    private @Nullable String folder;
    /**
     * @return The ID for the guest operating system
     * 
     */
    private String guestId;
    private String hostSystemId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer ideControllerCount;
    private @Nullable String ipAllocationPolicy;
    private @Nullable String ipProtocol;
    private @Nullable String localOvfPath;
    /**
     * @return The size of the virtual machine memory, in MB.
     * 
     */
    private Integer memory;
    /**
     * @return Allow memory to be added to the virtual machine
     * while powered on.
     * 
     */
    private Boolean memoryHotAddEnabled;
    private String name;
    /**
     * @return Enable nested hardware virtualization on the virtual
     * machine, facilitating nested virtualization in the guest.
     * 
     */
    private Boolean nestedHvEnabled;
    /**
     * @return The number of cores per virtual CPU in the virtual
     * machine.
     * 
     */
    private Integer numCoresPerSocket;
    /**
     * @return The number of virtual CPUs to assign to the virtual machine.
     * 
     */
    private Integer numCpus;
    private @Nullable Map<String,String> ovfNetworkMap;
    private @Nullable String remoteOvfUrl;
    private String resourcePoolId;
    private Integer sataControllerCount;
    private Integer scsiControllerCount;
    private String scsiType;
    /**
     * @return The swap file placement policy for the virtual
     * machine.
     * 
     */
    private String swapPlacementPolicy;

    private GetOvfVmTemplateResult() {}
    public Optional<Boolean> allowUnverifiedSslCert() {
        return Optional.ofNullable(this.allowUnverifiedSslCert);
    }
    /**
     * @return An alternate guest operating system name.
     * 
     */
    public String alternateGuestName() {
        return this.alternateGuestName;
    }
    /**
     * @return A description of the virtual machine.
     * 
     */
    public String annotation() {
        return this.annotation;
    }
    /**
     * @return Allow CPUs to be added to the virtual machine while
     * powered on.
     * 
     */
    public Boolean cpuHotAddEnabled() {
        return this.cpuHotAddEnabled;
    }
    /**
     * @return Allow CPUs to be removed from the virtual machine
     * while powered on.
     * 
     */
    public Boolean cpuHotRemoveEnabled() {
        return this.cpuHotRemoveEnabled;
    }
    public Boolean cpuPerformanceCountersEnabled() {
        return this.cpuPerformanceCountersEnabled;
    }
    public Optional<String> datastoreId() {
        return Optional.ofNullable(this.datastoreId);
    }
    public Optional<String> deploymentOption() {
        return Optional.ofNullable(this.deploymentOption);
    }
    public Optional<String> diskProvisioning() {
        return Optional.ofNullable(this.diskProvisioning);
    }
    public Optional<Boolean> enableHiddenProperties() {
        return Optional.ofNullable(this.enableHiddenProperties);
    }
    /**
     * @return The firmware to use on the virtual machine.
     * 
     */
    public String firmware() {
        return this.firmware;
    }
    public Optional<String> folder() {
        return Optional.ofNullable(this.folder);
    }
    /**
     * @return The ID for the guest operating system
     * 
     */
    public String guestId() {
        return this.guestId;
    }
    public String hostSystemId() {
        return this.hostSystemId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Integer ideControllerCount() {
        return this.ideControllerCount;
    }
    public Optional<String> ipAllocationPolicy() {
        return Optional.ofNullable(this.ipAllocationPolicy);
    }
    public Optional<String> ipProtocol() {
        return Optional.ofNullable(this.ipProtocol);
    }
    public Optional<String> localOvfPath() {
        return Optional.ofNullable(this.localOvfPath);
    }
    /**
     * @return The size of the virtual machine memory, in MB.
     * 
     */
    public Integer memory() {
        return this.memory;
    }
    /**
     * @return Allow memory to be added to the virtual machine
     * while powered on.
     * 
     */
    public Boolean memoryHotAddEnabled() {
        return this.memoryHotAddEnabled;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return Enable nested hardware virtualization on the virtual
     * machine, facilitating nested virtualization in the guest.
     * 
     */
    public Boolean nestedHvEnabled() {
        return this.nestedHvEnabled;
    }
    /**
     * @return The number of cores per virtual CPU in the virtual
     * machine.
     * 
     */
    public Integer numCoresPerSocket() {
        return this.numCoresPerSocket;
    }
    /**
     * @return The number of virtual CPUs to assign to the virtual machine.
     * 
     */
    public Integer numCpus() {
        return this.numCpus;
    }
    public Map<String,String> ovfNetworkMap() {
        return this.ovfNetworkMap == null ? Map.of() : this.ovfNetworkMap;
    }
    public Optional<String> remoteOvfUrl() {
        return Optional.ofNullable(this.remoteOvfUrl);
    }
    public String resourcePoolId() {
        return this.resourcePoolId;
    }
    public Integer sataControllerCount() {
        return this.sataControllerCount;
    }
    public Integer scsiControllerCount() {
        return this.scsiControllerCount;
    }
    public String scsiType() {
        return this.scsiType;
    }
    /**
     * @return The swap file placement policy for the virtual
     * machine.
     * 
     */
    public String swapPlacementPolicy() {
        return this.swapPlacementPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOvfVmTemplateResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean allowUnverifiedSslCert;
        private String alternateGuestName;
        private String annotation;
        private Boolean cpuHotAddEnabled;
        private Boolean cpuHotRemoveEnabled;
        private Boolean cpuPerformanceCountersEnabled;
        private @Nullable String datastoreId;
        private @Nullable String deploymentOption;
        private @Nullable String diskProvisioning;
        private @Nullable Boolean enableHiddenProperties;
        private String firmware;
        private @Nullable String folder;
        private String guestId;
        private String hostSystemId;
        private String id;
        private Integer ideControllerCount;
        private @Nullable String ipAllocationPolicy;
        private @Nullable String ipProtocol;
        private @Nullable String localOvfPath;
        private Integer memory;
        private Boolean memoryHotAddEnabled;
        private String name;
        private Boolean nestedHvEnabled;
        private Integer numCoresPerSocket;
        private Integer numCpus;
        private @Nullable Map<String,String> ovfNetworkMap;
        private @Nullable String remoteOvfUrl;
        private String resourcePoolId;
        private Integer sataControllerCount;
        private Integer scsiControllerCount;
        private String scsiType;
        private String swapPlacementPolicy;
        public Builder() {}
        public Builder(GetOvfVmTemplateResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowUnverifiedSslCert = defaults.allowUnverifiedSslCert;
    	      this.alternateGuestName = defaults.alternateGuestName;
    	      this.annotation = defaults.annotation;
    	      this.cpuHotAddEnabled = defaults.cpuHotAddEnabled;
    	      this.cpuHotRemoveEnabled = defaults.cpuHotRemoveEnabled;
    	      this.cpuPerformanceCountersEnabled = defaults.cpuPerformanceCountersEnabled;
    	      this.datastoreId = defaults.datastoreId;
    	      this.deploymentOption = defaults.deploymentOption;
    	      this.diskProvisioning = defaults.diskProvisioning;
    	      this.enableHiddenProperties = defaults.enableHiddenProperties;
    	      this.firmware = defaults.firmware;
    	      this.folder = defaults.folder;
    	      this.guestId = defaults.guestId;
    	      this.hostSystemId = defaults.hostSystemId;
    	      this.id = defaults.id;
    	      this.ideControllerCount = defaults.ideControllerCount;
    	      this.ipAllocationPolicy = defaults.ipAllocationPolicy;
    	      this.ipProtocol = defaults.ipProtocol;
    	      this.localOvfPath = defaults.localOvfPath;
    	      this.memory = defaults.memory;
    	      this.memoryHotAddEnabled = defaults.memoryHotAddEnabled;
    	      this.name = defaults.name;
    	      this.nestedHvEnabled = defaults.nestedHvEnabled;
    	      this.numCoresPerSocket = defaults.numCoresPerSocket;
    	      this.numCpus = defaults.numCpus;
    	      this.ovfNetworkMap = defaults.ovfNetworkMap;
    	      this.remoteOvfUrl = defaults.remoteOvfUrl;
    	      this.resourcePoolId = defaults.resourcePoolId;
    	      this.sataControllerCount = defaults.sataControllerCount;
    	      this.scsiControllerCount = defaults.scsiControllerCount;
    	      this.scsiType = defaults.scsiType;
    	      this.swapPlacementPolicy = defaults.swapPlacementPolicy;
        }

        @CustomType.Setter
        public Builder allowUnverifiedSslCert(@Nullable Boolean allowUnverifiedSslCert) {

            this.allowUnverifiedSslCert = allowUnverifiedSslCert;
            return this;
        }
        @CustomType.Setter
        public Builder alternateGuestName(String alternateGuestName) {
            if (alternateGuestName == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "alternateGuestName");
            }
            this.alternateGuestName = alternateGuestName;
            return this;
        }
        @CustomType.Setter
        public Builder annotation(String annotation) {
            if (annotation == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "annotation");
            }
            this.annotation = annotation;
            return this;
        }
        @CustomType.Setter
        public Builder cpuHotAddEnabled(Boolean cpuHotAddEnabled) {
            if (cpuHotAddEnabled == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "cpuHotAddEnabled");
            }
            this.cpuHotAddEnabled = cpuHotAddEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder cpuHotRemoveEnabled(Boolean cpuHotRemoveEnabled) {
            if (cpuHotRemoveEnabled == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "cpuHotRemoveEnabled");
            }
            this.cpuHotRemoveEnabled = cpuHotRemoveEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder cpuPerformanceCountersEnabled(Boolean cpuPerformanceCountersEnabled) {
            if (cpuPerformanceCountersEnabled == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "cpuPerformanceCountersEnabled");
            }
            this.cpuPerformanceCountersEnabled = cpuPerformanceCountersEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder datastoreId(@Nullable String datastoreId) {

            this.datastoreId = datastoreId;
            return this;
        }
        @CustomType.Setter
        public Builder deploymentOption(@Nullable String deploymentOption) {

            this.deploymentOption = deploymentOption;
            return this;
        }
        @CustomType.Setter
        public Builder diskProvisioning(@Nullable String diskProvisioning) {

            this.diskProvisioning = diskProvisioning;
            return this;
        }
        @CustomType.Setter
        public Builder enableHiddenProperties(@Nullable Boolean enableHiddenProperties) {

            this.enableHiddenProperties = enableHiddenProperties;
            return this;
        }
        @CustomType.Setter
        public Builder firmware(String firmware) {
            if (firmware == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "firmware");
            }
            this.firmware = firmware;
            return this;
        }
        @CustomType.Setter
        public Builder folder(@Nullable String folder) {

            this.folder = folder;
            return this;
        }
        @CustomType.Setter
        public Builder guestId(String guestId) {
            if (guestId == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "guestId");
            }
            this.guestId = guestId;
            return this;
        }
        @CustomType.Setter
        public Builder hostSystemId(String hostSystemId) {
            if (hostSystemId == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "hostSystemId");
            }
            this.hostSystemId = hostSystemId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ideControllerCount(Integer ideControllerCount) {
            if (ideControllerCount == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "ideControllerCount");
            }
            this.ideControllerCount = ideControllerCount;
            return this;
        }
        @CustomType.Setter
        public Builder ipAllocationPolicy(@Nullable String ipAllocationPolicy) {

            this.ipAllocationPolicy = ipAllocationPolicy;
            return this;
        }
        @CustomType.Setter
        public Builder ipProtocol(@Nullable String ipProtocol) {

            this.ipProtocol = ipProtocol;
            return this;
        }
        @CustomType.Setter
        public Builder localOvfPath(@Nullable String localOvfPath) {

            this.localOvfPath = localOvfPath;
            return this;
        }
        @CustomType.Setter
        public Builder memory(Integer memory) {
            if (memory == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "memory");
            }
            this.memory = memory;
            return this;
        }
        @CustomType.Setter
        public Builder memoryHotAddEnabled(Boolean memoryHotAddEnabled) {
            if (memoryHotAddEnabled == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "memoryHotAddEnabled");
            }
            this.memoryHotAddEnabled = memoryHotAddEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder nestedHvEnabled(Boolean nestedHvEnabled) {
            if (nestedHvEnabled == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "nestedHvEnabled");
            }
            this.nestedHvEnabled = nestedHvEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder numCoresPerSocket(Integer numCoresPerSocket) {
            if (numCoresPerSocket == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "numCoresPerSocket");
            }
            this.numCoresPerSocket = numCoresPerSocket;
            return this;
        }
        @CustomType.Setter
        public Builder numCpus(Integer numCpus) {
            if (numCpus == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "numCpus");
            }
            this.numCpus = numCpus;
            return this;
        }
        @CustomType.Setter
        public Builder ovfNetworkMap(@Nullable Map<String,String> ovfNetworkMap) {

            this.ovfNetworkMap = ovfNetworkMap;
            return this;
        }
        @CustomType.Setter
        public Builder remoteOvfUrl(@Nullable String remoteOvfUrl) {

            this.remoteOvfUrl = remoteOvfUrl;
            return this;
        }
        @CustomType.Setter
        public Builder resourcePoolId(String resourcePoolId) {
            if (resourcePoolId == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "resourcePoolId");
            }
            this.resourcePoolId = resourcePoolId;
            return this;
        }
        @CustomType.Setter
        public Builder sataControllerCount(Integer sataControllerCount) {
            if (sataControllerCount == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "sataControllerCount");
            }
            this.sataControllerCount = sataControllerCount;
            return this;
        }
        @CustomType.Setter
        public Builder scsiControllerCount(Integer scsiControllerCount) {
            if (scsiControllerCount == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "scsiControllerCount");
            }
            this.scsiControllerCount = scsiControllerCount;
            return this;
        }
        @CustomType.Setter
        public Builder scsiType(String scsiType) {
            if (scsiType == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "scsiType");
            }
            this.scsiType = scsiType;
            return this;
        }
        @CustomType.Setter
        public Builder swapPlacementPolicy(String swapPlacementPolicy) {
            if (swapPlacementPolicy == null) {
              throw new MissingRequiredPropertyException("GetOvfVmTemplateResult", "swapPlacementPolicy");
            }
            this.swapPlacementPolicy = swapPlacementPolicy;
            return this;
        }
        public GetOvfVmTemplateResult build() {
            final var _resultValue = new GetOvfVmTemplateResult();
            _resultValue.allowUnverifiedSslCert = allowUnverifiedSslCert;
            _resultValue.alternateGuestName = alternateGuestName;
            _resultValue.annotation = annotation;
            _resultValue.cpuHotAddEnabled = cpuHotAddEnabled;
            _resultValue.cpuHotRemoveEnabled = cpuHotRemoveEnabled;
            _resultValue.cpuPerformanceCountersEnabled = cpuPerformanceCountersEnabled;
            _resultValue.datastoreId = datastoreId;
            _resultValue.deploymentOption = deploymentOption;
            _resultValue.diskProvisioning = diskProvisioning;
            _resultValue.enableHiddenProperties = enableHiddenProperties;
            _resultValue.firmware = firmware;
            _resultValue.folder = folder;
            _resultValue.guestId = guestId;
            _resultValue.hostSystemId = hostSystemId;
            _resultValue.id = id;
            _resultValue.ideControllerCount = ideControllerCount;
            _resultValue.ipAllocationPolicy = ipAllocationPolicy;
            _resultValue.ipProtocol = ipProtocol;
            _resultValue.localOvfPath = localOvfPath;
            _resultValue.memory = memory;
            _resultValue.memoryHotAddEnabled = memoryHotAddEnabled;
            _resultValue.name = name;
            _resultValue.nestedHvEnabled = nestedHvEnabled;
            _resultValue.numCoresPerSocket = numCoresPerSocket;
            _resultValue.numCpus = numCpus;
            _resultValue.ovfNetworkMap = ovfNetworkMap;
            _resultValue.remoteOvfUrl = remoteOvfUrl;
            _resultValue.resourcePoolId = resourcePoolId;
            _resultValue.sataControllerCount = sataControllerCount;
            _resultValue.scsiControllerCount = scsiControllerCount;
            _resultValue.scsiType = scsiType;
            _resultValue.swapPlacementPolicy = swapPlacementPolicy;
            return _resultValue;
        }
    }
}
