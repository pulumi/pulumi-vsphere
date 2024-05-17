// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.vsphere.inputs.SupervisorEgressCidrArgs;
import com.pulumi.vsphere.inputs.SupervisorIngressCidrArgs;
import com.pulumi.vsphere.inputs.SupervisorManagementNetworkArgs;
import com.pulumi.vsphere.inputs.SupervisorNamespaceArgs;
import com.pulumi.vsphere.inputs.SupervisorPodCidrArgs;
import com.pulumi.vsphere.inputs.SupervisorServiceCidrArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SupervisorState extends com.pulumi.resources.ResourceArgs {

    public static final SupervisorState Empty = new SupervisorState();

    /**
     * The identifier of the compute cluster.
     * 
     */
    @Import(name="cluster")
    private @Nullable Output<String> cluster;

    /**
     * @return The identifier of the compute cluster.
     * 
     */
    public Optional<Output<String>> cluster() {
        return Optional.ofNullable(this.cluster);
    }

    /**
     * The identifier of the subscribed content library.
     * 
     */
    @Import(name="contentLibrary")
    private @Nullable Output<String> contentLibrary;

    /**
     * @return The identifier of the subscribed content library.
     * 
     */
    public Optional<Output<String>> contentLibrary() {
        return Optional.ofNullable(this.contentLibrary);
    }

    /**
     * The UUID of the distributed switch.
     * 
     */
    @Import(name="dvsUuid")
    private @Nullable Output<String> dvsUuid;

    /**
     * @return The UUID of the distributed switch.
     * 
     */
    public Optional<Output<String>> dvsUuid() {
        return Optional.ofNullable(this.dvsUuid);
    }

    /**
     * The identifier of the NSX Edge Cluster.
     * 
     */
    @Import(name="edgeCluster")
    private @Nullable Output<String> edgeCluster;

    /**
     * @return The identifier of the NSX Edge Cluster.
     * 
     */
    public Optional<Output<String>> edgeCluster() {
        return Optional.ofNullable(this.edgeCluster);
    }

    /**
     * CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs.
     * 
     */
    @Import(name="egressCidrs")
    private @Nullable Output<List<SupervisorEgressCidrArgs>> egressCidrs;

    /**
     * @return CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs.
     * 
     */
    public Optional<Output<List<SupervisorEgressCidrArgs>>> egressCidrs() {
        return Optional.ofNullable(this.egressCidrs);
    }

    /**
     * CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer.
     * 
     */
    @Import(name="ingressCidrs")
    private @Nullable Output<List<SupervisorIngressCidrArgs>> ingressCidrs;

    /**
     * @return CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer.
     * 
     */
    public Optional<Output<List<SupervisorIngressCidrArgs>>> ingressCidrs() {
        return Optional.ofNullable(this.ingressCidrs);
    }

    /**
     * The list of addresses of the primary DNS servers.
     * 
     */
    @Import(name="mainDns")
    private @Nullable Output<List<String>> mainDns;

    /**
     * @return The list of addresses of the primary DNS servers.
     * 
     */
    public Optional<Output<List<String>>> mainDns() {
        return Optional.ofNullable(this.mainDns);
    }

    /**
     * The configuration for the management network which the control plane VMs will be connected to.
     * * * `network` - ID of the network. (e.g. a distributed port group).
     * * * `starting_address` - Starting address of the management network range.
     * * * `subnet_mask` - Subnet mask.
     * * * `gateway` - Gateway IP address.
     * * * `address_count` - Number of addresses to allocate. Starts from `starting_address`
     * 
     */
    @Import(name="managementNetwork")
    private @Nullable Output<SupervisorManagementNetworkArgs> managementNetwork;

    /**
     * @return The configuration for the management network which the control plane VMs will be connected to.
     * * * `network` - ID of the network. (e.g. a distributed port group).
     * * * `starting_address` - Starting address of the management network range.
     * * * `subnet_mask` - Subnet mask.
     * * * `gateway` - Gateway IP address.
     * * * `address_count` - Number of addresses to allocate. Starts from `starting_address`
     * 
     */
    public Optional<Output<SupervisorManagementNetworkArgs>> managementNetwork() {
        return Optional.ofNullable(this.managementNetwork);
    }

    /**
     * The list of namespaces to create in the Supervisor cluster
     * 
     */
    @Import(name="namespaces")
    private @Nullable Output<List<SupervisorNamespaceArgs>> namespaces;

    /**
     * @return The list of namespaces to create in the Supervisor cluster
     * 
     */
    public Optional<Output<List<SupervisorNamespaceArgs>>> namespaces() {
        return Optional.ofNullable(this.namespaces);
    }

    /**
     * CIDR blocks from which Kubernetes allocates pod IP addresses. Minimum subnet size is 23.
     * 
     */
    @Import(name="podCidrs")
    private @Nullable Output<List<SupervisorPodCidrArgs>> podCidrs;

    /**
     * @return CIDR blocks from which Kubernetes allocates pod IP addresses. Minimum subnet size is 23.
     * 
     */
    public Optional<Output<List<SupervisorPodCidrArgs>>> podCidrs() {
        return Optional.ofNullable(this.podCidrs);
    }

    /**
     * List of DNS search domains.
     * 
     */
    @Import(name="searchDomains")
    private @Nullable Output<String> searchDomains;

    /**
     * @return List of DNS search domains.
     * 
     */
    public Optional<Output<String>> searchDomains() {
        return Optional.ofNullable(this.searchDomains);
    }

    /**
     * CIDR block from which Kubernetes allocates service cluster IP addresses.
     * 
     */
    @Import(name="serviceCidr")
    private @Nullable Output<SupervisorServiceCidrArgs> serviceCidr;

    /**
     * @return CIDR block from which Kubernetes allocates service cluster IP addresses.
     * 
     */
    public Optional<Output<SupervisorServiceCidrArgs>> serviceCidr() {
        return Optional.ofNullable(this.serviceCidr);
    }

    /**
     * The size of the Kubernetes API server.
     * 
     */
    @Import(name="sizingHint")
    private @Nullable Output<String> sizingHint;

    /**
     * @return The size of the Kubernetes API server.
     * 
     */
    public Optional<Output<String>> sizingHint() {
        return Optional.ofNullable(this.sizingHint);
    }

    /**
     * The name of the storage policy.
     * 
     */
    @Import(name="storagePolicy")
    private @Nullable Output<String> storagePolicy;

    /**
     * @return The name of the storage policy.
     * 
     */
    public Optional<Output<String>> storagePolicy() {
        return Optional.ofNullable(this.storagePolicy);
    }

    /**
     * The list of addresses of the DNS servers to use for the worker nodes.
     * 
     */
    @Import(name="workerDns")
    private @Nullable Output<List<String>> workerDns;

    /**
     * @return The list of addresses of the DNS servers to use for the worker nodes.
     * 
     */
    public Optional<Output<List<String>>> workerDns() {
        return Optional.ofNullable(this.workerDns);
    }

    private SupervisorState() {}

    private SupervisorState(SupervisorState $) {
        this.cluster = $.cluster;
        this.contentLibrary = $.contentLibrary;
        this.dvsUuid = $.dvsUuid;
        this.edgeCluster = $.edgeCluster;
        this.egressCidrs = $.egressCidrs;
        this.ingressCidrs = $.ingressCidrs;
        this.mainDns = $.mainDns;
        this.managementNetwork = $.managementNetwork;
        this.namespaces = $.namespaces;
        this.podCidrs = $.podCidrs;
        this.searchDomains = $.searchDomains;
        this.serviceCidr = $.serviceCidr;
        this.sizingHint = $.sizingHint;
        this.storagePolicy = $.storagePolicy;
        this.workerDns = $.workerDns;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SupervisorState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SupervisorState $;

        public Builder() {
            $ = new SupervisorState();
        }

        public Builder(SupervisorState defaults) {
            $ = new SupervisorState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cluster The identifier of the compute cluster.
         * 
         * @return builder
         * 
         */
        public Builder cluster(@Nullable Output<String> cluster) {
            $.cluster = cluster;
            return this;
        }

        /**
         * @param cluster The identifier of the compute cluster.
         * 
         * @return builder
         * 
         */
        public Builder cluster(String cluster) {
            return cluster(Output.of(cluster));
        }

        /**
         * @param contentLibrary The identifier of the subscribed content library.
         * 
         * @return builder
         * 
         */
        public Builder contentLibrary(@Nullable Output<String> contentLibrary) {
            $.contentLibrary = contentLibrary;
            return this;
        }

        /**
         * @param contentLibrary The identifier of the subscribed content library.
         * 
         * @return builder
         * 
         */
        public Builder contentLibrary(String contentLibrary) {
            return contentLibrary(Output.of(contentLibrary));
        }

        /**
         * @param dvsUuid The UUID of the distributed switch.
         * 
         * @return builder
         * 
         */
        public Builder dvsUuid(@Nullable Output<String> dvsUuid) {
            $.dvsUuid = dvsUuid;
            return this;
        }

        /**
         * @param dvsUuid The UUID of the distributed switch.
         * 
         * @return builder
         * 
         */
        public Builder dvsUuid(String dvsUuid) {
            return dvsUuid(Output.of(dvsUuid));
        }

        /**
         * @param edgeCluster The identifier of the NSX Edge Cluster.
         * 
         * @return builder
         * 
         */
        public Builder edgeCluster(@Nullable Output<String> edgeCluster) {
            $.edgeCluster = edgeCluster;
            return this;
        }

        /**
         * @param edgeCluster The identifier of the NSX Edge Cluster.
         * 
         * @return builder
         * 
         */
        public Builder edgeCluster(String edgeCluster) {
            return edgeCluster(Output.of(edgeCluster));
        }

        /**
         * @param egressCidrs CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs.
         * 
         * @return builder
         * 
         */
        public Builder egressCidrs(@Nullable Output<List<SupervisorEgressCidrArgs>> egressCidrs) {
            $.egressCidrs = egressCidrs;
            return this;
        }

        /**
         * @param egressCidrs CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs.
         * 
         * @return builder
         * 
         */
        public Builder egressCidrs(List<SupervisorEgressCidrArgs> egressCidrs) {
            return egressCidrs(Output.of(egressCidrs));
        }

        /**
         * @param egressCidrs CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs.
         * 
         * @return builder
         * 
         */
        public Builder egressCidrs(SupervisorEgressCidrArgs... egressCidrs) {
            return egressCidrs(List.of(egressCidrs));
        }

        /**
         * @param ingressCidrs CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer.
         * 
         * @return builder
         * 
         */
        public Builder ingressCidrs(@Nullable Output<List<SupervisorIngressCidrArgs>> ingressCidrs) {
            $.ingressCidrs = ingressCidrs;
            return this;
        }

        /**
         * @param ingressCidrs CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer.
         * 
         * @return builder
         * 
         */
        public Builder ingressCidrs(List<SupervisorIngressCidrArgs> ingressCidrs) {
            return ingressCidrs(Output.of(ingressCidrs));
        }

        /**
         * @param ingressCidrs CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer.
         * 
         * @return builder
         * 
         */
        public Builder ingressCidrs(SupervisorIngressCidrArgs... ingressCidrs) {
            return ingressCidrs(List.of(ingressCidrs));
        }

        /**
         * @param mainDns The list of addresses of the primary DNS servers.
         * 
         * @return builder
         * 
         */
        public Builder mainDns(@Nullable Output<List<String>> mainDns) {
            $.mainDns = mainDns;
            return this;
        }

        /**
         * @param mainDns The list of addresses of the primary DNS servers.
         * 
         * @return builder
         * 
         */
        public Builder mainDns(List<String> mainDns) {
            return mainDns(Output.of(mainDns));
        }

        /**
         * @param mainDns The list of addresses of the primary DNS servers.
         * 
         * @return builder
         * 
         */
        public Builder mainDns(String... mainDns) {
            return mainDns(List.of(mainDns));
        }

        /**
         * @param managementNetwork The configuration for the management network which the control plane VMs will be connected to.
         * * * `network` - ID of the network. (e.g. a distributed port group).
         * * * `starting_address` - Starting address of the management network range.
         * * * `subnet_mask` - Subnet mask.
         * * * `gateway` - Gateway IP address.
         * * * `address_count` - Number of addresses to allocate. Starts from `starting_address`
         * 
         * @return builder
         * 
         */
        public Builder managementNetwork(@Nullable Output<SupervisorManagementNetworkArgs> managementNetwork) {
            $.managementNetwork = managementNetwork;
            return this;
        }

        /**
         * @param managementNetwork The configuration for the management network which the control plane VMs will be connected to.
         * * * `network` - ID of the network. (e.g. a distributed port group).
         * * * `starting_address` - Starting address of the management network range.
         * * * `subnet_mask` - Subnet mask.
         * * * `gateway` - Gateway IP address.
         * * * `address_count` - Number of addresses to allocate. Starts from `starting_address`
         * 
         * @return builder
         * 
         */
        public Builder managementNetwork(SupervisorManagementNetworkArgs managementNetwork) {
            return managementNetwork(Output.of(managementNetwork));
        }

        /**
         * @param namespaces The list of namespaces to create in the Supervisor cluster
         * 
         * @return builder
         * 
         */
        public Builder namespaces(@Nullable Output<List<SupervisorNamespaceArgs>> namespaces) {
            $.namespaces = namespaces;
            return this;
        }

        /**
         * @param namespaces The list of namespaces to create in the Supervisor cluster
         * 
         * @return builder
         * 
         */
        public Builder namespaces(List<SupervisorNamespaceArgs> namespaces) {
            return namespaces(Output.of(namespaces));
        }

        /**
         * @param namespaces The list of namespaces to create in the Supervisor cluster
         * 
         * @return builder
         * 
         */
        public Builder namespaces(SupervisorNamespaceArgs... namespaces) {
            return namespaces(List.of(namespaces));
        }

        /**
         * @param podCidrs CIDR blocks from which Kubernetes allocates pod IP addresses. Minimum subnet size is 23.
         * 
         * @return builder
         * 
         */
        public Builder podCidrs(@Nullable Output<List<SupervisorPodCidrArgs>> podCidrs) {
            $.podCidrs = podCidrs;
            return this;
        }

        /**
         * @param podCidrs CIDR blocks from which Kubernetes allocates pod IP addresses. Minimum subnet size is 23.
         * 
         * @return builder
         * 
         */
        public Builder podCidrs(List<SupervisorPodCidrArgs> podCidrs) {
            return podCidrs(Output.of(podCidrs));
        }

        /**
         * @param podCidrs CIDR blocks from which Kubernetes allocates pod IP addresses. Minimum subnet size is 23.
         * 
         * @return builder
         * 
         */
        public Builder podCidrs(SupervisorPodCidrArgs... podCidrs) {
            return podCidrs(List.of(podCidrs));
        }

        /**
         * @param searchDomains List of DNS search domains.
         * 
         * @return builder
         * 
         */
        public Builder searchDomains(@Nullable Output<String> searchDomains) {
            $.searchDomains = searchDomains;
            return this;
        }

        /**
         * @param searchDomains List of DNS search domains.
         * 
         * @return builder
         * 
         */
        public Builder searchDomains(String searchDomains) {
            return searchDomains(Output.of(searchDomains));
        }

        /**
         * @param serviceCidr CIDR block from which Kubernetes allocates service cluster IP addresses.
         * 
         * @return builder
         * 
         */
        public Builder serviceCidr(@Nullable Output<SupervisorServiceCidrArgs> serviceCidr) {
            $.serviceCidr = serviceCidr;
            return this;
        }

        /**
         * @param serviceCidr CIDR block from which Kubernetes allocates service cluster IP addresses.
         * 
         * @return builder
         * 
         */
        public Builder serviceCidr(SupervisorServiceCidrArgs serviceCidr) {
            return serviceCidr(Output.of(serviceCidr));
        }

        /**
         * @param sizingHint The size of the Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder sizingHint(@Nullable Output<String> sizingHint) {
            $.sizingHint = sizingHint;
            return this;
        }

        /**
         * @param sizingHint The size of the Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder sizingHint(String sizingHint) {
            return sizingHint(Output.of(sizingHint));
        }

        /**
         * @param storagePolicy The name of the storage policy.
         * 
         * @return builder
         * 
         */
        public Builder storagePolicy(@Nullable Output<String> storagePolicy) {
            $.storagePolicy = storagePolicy;
            return this;
        }

        /**
         * @param storagePolicy The name of the storage policy.
         * 
         * @return builder
         * 
         */
        public Builder storagePolicy(String storagePolicy) {
            return storagePolicy(Output.of(storagePolicy));
        }

        /**
         * @param workerDns The list of addresses of the DNS servers to use for the worker nodes.
         * 
         * @return builder
         * 
         */
        public Builder workerDns(@Nullable Output<List<String>> workerDns) {
            $.workerDns = workerDns;
            return this;
        }

        /**
         * @param workerDns The list of addresses of the DNS servers to use for the worker nodes.
         * 
         * @return builder
         * 
         */
        public Builder workerDns(List<String> workerDns) {
            return workerDns(Output.of(workerDns));
        }

        /**
         * @param workerDns The list of addresses of the DNS servers to use for the worker nodes.
         * 
         * @return builder
         * 
         */
        public Builder workerDns(String... workerDns) {
            return workerDns(List.of(workerDns));
        }

        public SupervisorState build() {
            return $;
        }
    }

}