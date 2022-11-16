// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOvfVmTemplatePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOvfVmTemplatePlainArgs Empty = new GetOvfVmTemplatePlainArgs();

    /**
     * Allow unverified SSL certificates
     * when deploying OVF/OVA from a URL.
     * 
     */
    @Import(name="allowUnverifiedSslCert")
    private @Nullable Boolean allowUnverifiedSslCert;

    /**
     * @return Allow unverified SSL certificates
     * when deploying OVF/OVA from a URL.
     * 
     */
    public Optional<Boolean> allowUnverifiedSslCert() {
        return Optional.ofNullable(this.allowUnverifiedSslCert);
    }

    /**
     * The ID of the virtual machine&#39;s datastore. The
     * virtual machine configuration is placed here, along with any virtual disks
     * that are created without datastores.
     * 
     */
    @Import(name="datastoreId")
    private @Nullable String datastoreId;

    /**
     * @return The ID of the virtual machine&#39;s datastore. The
     * virtual machine configuration is placed here, along with any virtual disks
     * that are created without datastores.
     * 
     */
    public Optional<String> datastoreId() {
        return Optional.ofNullable(this.datastoreId);
    }

    /**
     * The key of the chosen deployment option. If
     * empty, the default option is chosen.
     * 
     */
    @Import(name="deploymentOption")
    private @Nullable String deploymentOption;

    /**
     * @return The key of the chosen deployment option. If
     * empty, the default option is chosen.
     * 
     */
    public Optional<String> deploymentOption() {
        return Optional.ofNullable(this.deploymentOption);
    }

    /**
     * The disk provisioning type. If set, all the
     * disks in the deployed OVA/OVF will have the same specified disk type. Can be
     * one of `thin`, `flat`, `thick` or `sameAsSource`.
     * 
     */
    @Import(name="diskProvisioning")
    private @Nullable String diskProvisioning;

    /**
     * @return The disk provisioning type. If set, all the
     * disks in the deployed OVA/OVF will have the same specified disk type. Can be
     * one of `thin`, `flat`, `thick` or `sameAsSource`.
     * 
     */
    public Optional<String> diskProvisioning() {
        return Optional.ofNullable(this.diskProvisioning);
    }

    @Import(name="enableHiddenProperties")
    private @Nullable Boolean enableHiddenProperties;

    public Optional<Boolean> enableHiddenProperties() {
        return Optional.ofNullable(this.enableHiddenProperties);
    }

    /**
     * The name of the folder in which to place the virtual
     * machine.
     * 
     */
    @Import(name="folder")
    private @Nullable String folder;

    /**
     * @return The name of the folder in which to place the virtual
     * machine.
     * 
     */
    public Optional<String> folder() {
        return Optional.ofNullable(this.folder);
    }

    /**
     * The ID of the ESXi host system to deploy the
     * virtual machine.
     * 
     */
    @Import(name="hostSystemId", required=true)
    private String hostSystemId;

    /**
     * @return The ID of the ESXi host system to deploy the
     * virtual machine.
     * 
     */
    public String hostSystemId() {
        return this.hostSystemId;
    }

    /**
     * The IP allocation policy.
     * 
     */
    @Import(name="ipAllocationPolicy")
    private @Nullable String ipAllocationPolicy;

    /**
     * @return The IP allocation policy.
     * 
     */
    public Optional<String> ipAllocationPolicy() {
        return Optional.ofNullable(this.ipAllocationPolicy);
    }

    /**
     * The IP protocol.
     * 
     */
    @Import(name="ipProtocol")
    private @Nullable String ipProtocol;

    /**
     * @return The IP protocol.
     * 
     */
    public Optional<String> ipProtocol() {
        return Optional.ofNullable(this.ipProtocol);
    }

    /**
     * The absolute path to the OVF/OVA file on the
     * local system. When deploying from an OVF, ensure all necessary files such as
     * the `.vmdk` files are present in the same directory as the OVF.
     * 
     */
    @Import(name="localOvfPath")
    private @Nullable String localOvfPath;

    /**
     * @return The absolute path to the OVF/OVA file on the
     * local system. When deploying from an OVF, ensure all necessary files such as
     * the `.vmdk` files are present in the same directory as the OVF.
     * 
     */
    public Optional<String> localOvfPath() {
        return Optional.ofNullable(this.localOvfPath);
    }

    /**
     * Name of the virtual machine to create.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Name of the virtual machine to create.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The mapping of name of network identifiers
     * from the OVF descriptor to network UUID in the environment.
     * 
     */
    @Import(name="ovfNetworkMap")
    private @Nullable Map<String,String> ovfNetworkMap;

    /**
     * @return The mapping of name of network identifiers
     * from the OVF descriptor to network UUID in the environment.
     * 
     */
    public Optional<Map<String,String>> ovfNetworkMap() {
        return Optional.ofNullable(this.ovfNetworkMap);
    }

    /**
     * URL of the remote OVF/OVA file to be deployed.
     * 
     */
    @Import(name="remoteOvfUrl")
    private @Nullable String remoteOvfUrl;

    /**
     * @return URL of the remote OVF/OVA file to be deployed.
     * 
     */
    public Optional<String> remoteOvfUrl() {
        return Optional.ofNullable(this.remoteOvfUrl);
    }

    /**
     * The ID of a resource pool in which to place
     * the virtual machine.
     * 
     */
    @Import(name="resourcePoolId", required=true)
    private String resourcePoolId;

    /**
     * @return The ID of a resource pool in which to place
     * the virtual machine.
     * 
     */
    public String resourcePoolId() {
        return this.resourcePoolId;
    }

    private GetOvfVmTemplatePlainArgs() {}

    private GetOvfVmTemplatePlainArgs(GetOvfVmTemplatePlainArgs $) {
        this.allowUnverifiedSslCert = $.allowUnverifiedSslCert;
        this.datastoreId = $.datastoreId;
        this.deploymentOption = $.deploymentOption;
        this.diskProvisioning = $.diskProvisioning;
        this.enableHiddenProperties = $.enableHiddenProperties;
        this.folder = $.folder;
        this.hostSystemId = $.hostSystemId;
        this.ipAllocationPolicy = $.ipAllocationPolicy;
        this.ipProtocol = $.ipProtocol;
        this.localOvfPath = $.localOvfPath;
        this.name = $.name;
        this.ovfNetworkMap = $.ovfNetworkMap;
        this.remoteOvfUrl = $.remoteOvfUrl;
        this.resourcePoolId = $.resourcePoolId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOvfVmTemplatePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOvfVmTemplatePlainArgs $;

        public Builder() {
            $ = new GetOvfVmTemplatePlainArgs();
        }

        public Builder(GetOvfVmTemplatePlainArgs defaults) {
            $ = new GetOvfVmTemplatePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowUnverifiedSslCert Allow unverified SSL certificates
         * when deploying OVF/OVA from a URL.
         * 
         * @return builder
         * 
         */
        public Builder allowUnverifiedSslCert(@Nullable Boolean allowUnverifiedSslCert) {
            $.allowUnverifiedSslCert = allowUnverifiedSslCert;
            return this;
        }

        /**
         * @param datastoreId The ID of the virtual machine&#39;s datastore. The
         * virtual machine configuration is placed here, along with any virtual disks
         * that are created without datastores.
         * 
         * @return builder
         * 
         */
        public Builder datastoreId(@Nullable String datastoreId) {
            $.datastoreId = datastoreId;
            return this;
        }

        /**
         * @param deploymentOption The key of the chosen deployment option. If
         * empty, the default option is chosen.
         * 
         * @return builder
         * 
         */
        public Builder deploymentOption(@Nullable String deploymentOption) {
            $.deploymentOption = deploymentOption;
            return this;
        }

        /**
         * @param diskProvisioning The disk provisioning type. If set, all the
         * disks in the deployed OVA/OVF will have the same specified disk type. Can be
         * one of `thin`, `flat`, `thick` or `sameAsSource`.
         * 
         * @return builder
         * 
         */
        public Builder diskProvisioning(@Nullable String diskProvisioning) {
            $.diskProvisioning = diskProvisioning;
            return this;
        }

        public Builder enableHiddenProperties(@Nullable Boolean enableHiddenProperties) {
            $.enableHiddenProperties = enableHiddenProperties;
            return this;
        }

        /**
         * @param folder The name of the folder in which to place the virtual
         * machine.
         * 
         * @return builder
         * 
         */
        public Builder folder(@Nullable String folder) {
            $.folder = folder;
            return this;
        }

        /**
         * @param hostSystemId The ID of the ESXi host system to deploy the
         * virtual machine.
         * 
         * @return builder
         * 
         */
        public Builder hostSystemId(String hostSystemId) {
            $.hostSystemId = hostSystemId;
            return this;
        }

        /**
         * @param ipAllocationPolicy The IP allocation policy.
         * 
         * @return builder
         * 
         */
        public Builder ipAllocationPolicy(@Nullable String ipAllocationPolicy) {
            $.ipAllocationPolicy = ipAllocationPolicy;
            return this;
        }

        /**
         * @param ipProtocol The IP protocol.
         * 
         * @return builder
         * 
         */
        public Builder ipProtocol(@Nullable String ipProtocol) {
            $.ipProtocol = ipProtocol;
            return this;
        }

        /**
         * @param localOvfPath The absolute path to the OVF/OVA file on the
         * local system. When deploying from an OVF, ensure all necessary files such as
         * the `.vmdk` files are present in the same directory as the OVF.
         * 
         * @return builder
         * 
         */
        public Builder localOvfPath(@Nullable String localOvfPath) {
            $.localOvfPath = localOvfPath;
            return this;
        }

        /**
         * @param name Name of the virtual machine to create.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param ovfNetworkMap The mapping of name of network identifiers
         * from the OVF descriptor to network UUID in the environment.
         * 
         * @return builder
         * 
         */
        public Builder ovfNetworkMap(@Nullable Map<String,String> ovfNetworkMap) {
            $.ovfNetworkMap = ovfNetworkMap;
            return this;
        }

        /**
         * @param remoteOvfUrl URL of the remote OVF/OVA file to be deployed.
         * 
         * @return builder
         * 
         */
        public Builder remoteOvfUrl(@Nullable String remoteOvfUrl) {
            $.remoteOvfUrl = remoteOvfUrl;
            return this;
        }

        /**
         * @param resourcePoolId The ID of a resource pool in which to place
         * the virtual machine.
         * 
         * @return builder
         * 
         */
        public Builder resourcePoolId(String resourcePoolId) {
            $.resourcePoolId = resourcePoolId;
            return this;
        }

        public GetOvfVmTemplatePlainArgs build() {
            $.hostSystemId = Objects.requireNonNull($.hostSystemId, "expected parameter 'hostSystemId' to be non-null");
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.resourcePoolId = Objects.requireNonNull($.resourcePoolId, "expected parameter 'resourcePoolId' to be non-null");
            return $;
        }
    }

}