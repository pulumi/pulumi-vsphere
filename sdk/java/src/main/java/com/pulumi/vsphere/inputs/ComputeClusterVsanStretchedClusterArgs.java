// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ComputeClusterVsanStretchedClusterArgs extends com.pulumi.resources.ResourceArgs {

    public static final ComputeClusterVsanStretchedClusterArgs Empty = new ComputeClusterVsanStretchedClusterArgs();

    /**
     * The managed object IDs of the hosts to put in the first fault domain.
     * 
     */
    @Import(name="preferredFaultDomainHostIds", required=true)
    private Output<List<String>> preferredFaultDomainHostIds;

    /**
     * @return The managed object IDs of the hosts to put in the first fault domain.
     * 
     */
    public Output<List<String>> preferredFaultDomainHostIds() {
        return this.preferredFaultDomainHostIds;
    }

    /**
     * The name of first fault domain. Default is `Preferred`.
     * 
     */
    @Import(name="preferredFaultDomainName")
    private @Nullable Output<String> preferredFaultDomainName;

    /**
     * @return The name of first fault domain. Default is `Preferred`.
     * 
     */
    public Optional<Output<String>> preferredFaultDomainName() {
        return Optional.ofNullable(this.preferredFaultDomainName);
    }

    /**
     * The managed object IDs of the hosts to put in the second fault domain.
     * 
     */
    @Import(name="secondaryFaultDomainHostIds", required=true)
    private Output<List<String>> secondaryFaultDomainHostIds;

    /**
     * @return The managed object IDs of the hosts to put in the second fault domain.
     * 
     */
    public Output<List<String>> secondaryFaultDomainHostIds() {
        return this.secondaryFaultDomainHostIds;
    }

    /**
     * The name of second fault domain. Default is `Secondary`.
     * 
     * &gt; **NOTE:** You must disable vSphere HA before you enable vSAN on the cluster.
     * You can enable or re-enable vSphere HA after vSAN is configured.
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vsphere.ComputeCluster;
     * import com.pulumi.vsphere.ComputeClusterArgs;
     * import com.pulumi.vsphere.inputs.ComputeClusterVsanDiskGroupArgs;
     * import com.pulumi.vsphere.inputs.ComputeClusterVsanFaultDomainArgs;
     * import com.pulumi.vsphere.inputs.ComputeClusterVsanStretchedClusterArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var computeCluster = new ComputeCluster(&#34;computeCluster&#34;, ComputeClusterArgs.builder()        
     *             .datacenterId(data.vsphere_datacenter().datacenter().id())
     *             .hostSystemIds(data.vsphere_host().host().stream().map(element -&gt; element.id()).collect(toList()))
     *             .drsEnabled(true)
     *             .drsAutomationLevel(&#34;fullyAutomated&#34;)
     *             .haEnabled(false)
     *             .vsanEnabled(true)
     *             .vsanEsaEnabled(true)
     *             .vsanDedupEnabled(true)
     *             .vsanCompressionEnabled(true)
     *             .vsanPerformanceEnabled(true)
     *             .vsanVerboseModeEnabled(true)
     *             .vsanNetworkDiagnosticModeEnabled(true)
     *             .vsanUnmapEnabled(true)
     *             .vsanDitEncryptionEnabled(true)
     *             .vsanDitRekeyInterval(1800)
     *             .vsanDiskGroups(ComputeClusterVsanDiskGroupArgs.builder()
     *                 .cache(data.vsphere_vmfs_disks().cache_disks()[0])
     *                 .storages(data.vsphere_vmfs_disks().storage_disks())
     *                 .build())
     *             .vsanFaultDomains(ComputeClusterVsanFaultDomainArgs.builder()
     *                 .faultDomains(                
     *                     ComputeClusterVsanFaultDomainFaultDomainArgs.builder()
     *                         .name(&#34;fd1&#34;)
     *                         .hostIds(data.vsphere_host().faultdomain1_hosts().stream().map(element -&gt; element.id()).collect(toList()))
     *                         .build(),
     *                     ComputeClusterVsanFaultDomainFaultDomainArgs.builder()
     *                         .name(&#34;fd2&#34;)
     *                         .hostIds(data.vsphere_host().faultdomain2_hosts().stream().map(element -&gt; element.id()).collect(toList()))
     *                         .build())
     *                 .build())
     *             .vsanStretchedCluster(ComputeClusterVsanStretchedClusterArgs.builder()
     *                 .preferredFaultDomainHostIds(data.vsphere_host().preferred_fault_domain_host().stream().map(element -&gt; element.id()).collect(toList()))
     *                 .secondaryFaultDomainHostIds(data.vsphere_host().secondary_fault_domain_host().stream().map(element -&gt; element.id()).collect(toList()))
     *                 .witnessNode(data.vsphere_host().witness_host().id())
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    @Import(name="secondaryFaultDomainName")
    private @Nullable Output<String> secondaryFaultDomainName;

    /**
     * @return The name of second fault domain. Default is `Secondary`.
     * 
     * &gt; **NOTE:** You must disable vSphere HA before you enable vSAN on the cluster.
     * You can enable or re-enable vSphere HA after vSAN is configured.
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vsphere.ComputeCluster;
     * import com.pulumi.vsphere.ComputeClusterArgs;
     * import com.pulumi.vsphere.inputs.ComputeClusterVsanDiskGroupArgs;
     * import com.pulumi.vsphere.inputs.ComputeClusterVsanFaultDomainArgs;
     * import com.pulumi.vsphere.inputs.ComputeClusterVsanStretchedClusterArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var computeCluster = new ComputeCluster(&#34;computeCluster&#34;, ComputeClusterArgs.builder()        
     *             .datacenterId(data.vsphere_datacenter().datacenter().id())
     *             .hostSystemIds(data.vsphere_host().host().stream().map(element -&gt; element.id()).collect(toList()))
     *             .drsEnabled(true)
     *             .drsAutomationLevel(&#34;fullyAutomated&#34;)
     *             .haEnabled(false)
     *             .vsanEnabled(true)
     *             .vsanEsaEnabled(true)
     *             .vsanDedupEnabled(true)
     *             .vsanCompressionEnabled(true)
     *             .vsanPerformanceEnabled(true)
     *             .vsanVerboseModeEnabled(true)
     *             .vsanNetworkDiagnosticModeEnabled(true)
     *             .vsanUnmapEnabled(true)
     *             .vsanDitEncryptionEnabled(true)
     *             .vsanDitRekeyInterval(1800)
     *             .vsanDiskGroups(ComputeClusterVsanDiskGroupArgs.builder()
     *                 .cache(data.vsphere_vmfs_disks().cache_disks()[0])
     *                 .storages(data.vsphere_vmfs_disks().storage_disks())
     *                 .build())
     *             .vsanFaultDomains(ComputeClusterVsanFaultDomainArgs.builder()
     *                 .faultDomains(                
     *                     ComputeClusterVsanFaultDomainFaultDomainArgs.builder()
     *                         .name(&#34;fd1&#34;)
     *                         .hostIds(data.vsphere_host().faultdomain1_hosts().stream().map(element -&gt; element.id()).collect(toList()))
     *                         .build(),
     *                     ComputeClusterVsanFaultDomainFaultDomainArgs.builder()
     *                         .name(&#34;fd2&#34;)
     *                         .hostIds(data.vsphere_host().faultdomain2_hosts().stream().map(element -&gt; element.id()).collect(toList()))
     *                         .build())
     *                 .build())
     *             .vsanStretchedCluster(ComputeClusterVsanStretchedClusterArgs.builder()
     *                 .preferredFaultDomainHostIds(data.vsphere_host().preferred_fault_domain_host().stream().map(element -&gt; element.id()).collect(toList()))
     *                 .secondaryFaultDomainHostIds(data.vsphere_host().secondary_fault_domain_host().stream().map(element -&gt; element.id()).collect(toList()))
     *                 .witnessNode(data.vsphere_host().witness_host().id())
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public Optional<Output<String>> secondaryFaultDomainName() {
        return Optional.ofNullable(this.secondaryFaultDomainName);
    }

    /**
     * The managed object IDs of the host selected as witness node when enable stretched cluster.
     * 
     */
    @Import(name="witnessNode", required=true)
    private Output<String> witnessNode;

    /**
     * @return The managed object IDs of the host selected as witness node when enable stretched cluster.
     * 
     */
    public Output<String> witnessNode() {
        return this.witnessNode;
    }

    private ComputeClusterVsanStretchedClusterArgs() {}

    private ComputeClusterVsanStretchedClusterArgs(ComputeClusterVsanStretchedClusterArgs $) {
        this.preferredFaultDomainHostIds = $.preferredFaultDomainHostIds;
        this.preferredFaultDomainName = $.preferredFaultDomainName;
        this.secondaryFaultDomainHostIds = $.secondaryFaultDomainHostIds;
        this.secondaryFaultDomainName = $.secondaryFaultDomainName;
        this.witnessNode = $.witnessNode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ComputeClusterVsanStretchedClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ComputeClusterVsanStretchedClusterArgs $;

        public Builder() {
            $ = new ComputeClusterVsanStretchedClusterArgs();
        }

        public Builder(ComputeClusterVsanStretchedClusterArgs defaults) {
            $ = new ComputeClusterVsanStretchedClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param preferredFaultDomainHostIds The managed object IDs of the hosts to put in the first fault domain.
         * 
         * @return builder
         * 
         */
        public Builder preferredFaultDomainHostIds(Output<List<String>> preferredFaultDomainHostIds) {
            $.preferredFaultDomainHostIds = preferredFaultDomainHostIds;
            return this;
        }

        /**
         * @param preferredFaultDomainHostIds The managed object IDs of the hosts to put in the first fault domain.
         * 
         * @return builder
         * 
         */
        public Builder preferredFaultDomainHostIds(List<String> preferredFaultDomainHostIds) {
            return preferredFaultDomainHostIds(Output.of(preferredFaultDomainHostIds));
        }

        /**
         * @param preferredFaultDomainHostIds The managed object IDs of the hosts to put in the first fault domain.
         * 
         * @return builder
         * 
         */
        public Builder preferredFaultDomainHostIds(String... preferredFaultDomainHostIds) {
            return preferredFaultDomainHostIds(List.of(preferredFaultDomainHostIds));
        }

        /**
         * @param preferredFaultDomainName The name of first fault domain. Default is `Preferred`.
         * 
         * @return builder
         * 
         */
        public Builder preferredFaultDomainName(@Nullable Output<String> preferredFaultDomainName) {
            $.preferredFaultDomainName = preferredFaultDomainName;
            return this;
        }

        /**
         * @param preferredFaultDomainName The name of first fault domain. Default is `Preferred`.
         * 
         * @return builder
         * 
         */
        public Builder preferredFaultDomainName(String preferredFaultDomainName) {
            return preferredFaultDomainName(Output.of(preferredFaultDomainName));
        }

        /**
         * @param secondaryFaultDomainHostIds The managed object IDs of the hosts to put in the second fault domain.
         * 
         * @return builder
         * 
         */
        public Builder secondaryFaultDomainHostIds(Output<List<String>> secondaryFaultDomainHostIds) {
            $.secondaryFaultDomainHostIds = secondaryFaultDomainHostIds;
            return this;
        }

        /**
         * @param secondaryFaultDomainHostIds The managed object IDs of the hosts to put in the second fault domain.
         * 
         * @return builder
         * 
         */
        public Builder secondaryFaultDomainHostIds(List<String> secondaryFaultDomainHostIds) {
            return secondaryFaultDomainHostIds(Output.of(secondaryFaultDomainHostIds));
        }

        /**
         * @param secondaryFaultDomainHostIds The managed object IDs of the hosts to put in the second fault domain.
         * 
         * @return builder
         * 
         */
        public Builder secondaryFaultDomainHostIds(String... secondaryFaultDomainHostIds) {
            return secondaryFaultDomainHostIds(List.of(secondaryFaultDomainHostIds));
        }

        /**
         * @param secondaryFaultDomainName The name of second fault domain. Default is `Secondary`.
         * 
         * &gt; **NOTE:** You must disable vSphere HA before you enable vSAN on the cluster.
         * You can enable or re-enable vSphere HA after vSAN is configured.
         * 
         * &lt;!--Start PulumiCodeChooser --&gt;
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
         * import com.pulumi.vsphere.ComputeCluster;
         * import com.pulumi.vsphere.ComputeClusterArgs;
         * import com.pulumi.vsphere.inputs.ComputeClusterVsanDiskGroupArgs;
         * import com.pulumi.vsphere.inputs.ComputeClusterVsanFaultDomainArgs;
         * import com.pulumi.vsphere.inputs.ComputeClusterVsanStretchedClusterArgs;
         * import java.util.List;
         * import java.util.ArrayList;
         * import java.util.Map;
         * import java.io.File;
         * import java.nio.file.Files;
         * import java.nio.file.Paths;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *         var computeCluster = new ComputeCluster(&#34;computeCluster&#34;, ComputeClusterArgs.builder()        
         *             .datacenterId(data.vsphere_datacenter().datacenter().id())
         *             .hostSystemIds(data.vsphere_host().host().stream().map(element -&gt; element.id()).collect(toList()))
         *             .drsEnabled(true)
         *             .drsAutomationLevel(&#34;fullyAutomated&#34;)
         *             .haEnabled(false)
         *             .vsanEnabled(true)
         *             .vsanEsaEnabled(true)
         *             .vsanDedupEnabled(true)
         *             .vsanCompressionEnabled(true)
         *             .vsanPerformanceEnabled(true)
         *             .vsanVerboseModeEnabled(true)
         *             .vsanNetworkDiagnosticModeEnabled(true)
         *             .vsanUnmapEnabled(true)
         *             .vsanDitEncryptionEnabled(true)
         *             .vsanDitRekeyInterval(1800)
         *             .vsanDiskGroups(ComputeClusterVsanDiskGroupArgs.builder()
         *                 .cache(data.vsphere_vmfs_disks().cache_disks()[0])
         *                 .storages(data.vsphere_vmfs_disks().storage_disks())
         *                 .build())
         *             .vsanFaultDomains(ComputeClusterVsanFaultDomainArgs.builder()
         *                 .faultDomains(                
         *                     ComputeClusterVsanFaultDomainFaultDomainArgs.builder()
         *                         .name(&#34;fd1&#34;)
         *                         .hostIds(data.vsphere_host().faultdomain1_hosts().stream().map(element -&gt; element.id()).collect(toList()))
         *                         .build(),
         *                     ComputeClusterVsanFaultDomainFaultDomainArgs.builder()
         *                         .name(&#34;fd2&#34;)
         *                         .hostIds(data.vsphere_host().faultdomain2_hosts().stream().map(element -&gt; element.id()).collect(toList()))
         *                         .build())
         *                 .build())
         *             .vsanStretchedCluster(ComputeClusterVsanStretchedClusterArgs.builder()
         *                 .preferredFaultDomainHostIds(data.vsphere_host().preferred_fault_domain_host().stream().map(element -&gt; element.id()).collect(toList()))
         *                 .secondaryFaultDomainHostIds(data.vsphere_host().secondary_fault_domain_host().stream().map(element -&gt; element.id()).collect(toList()))
         *                 .witnessNode(data.vsphere_host().witness_host().id())
         *                 .build())
         *             .build());
         * 
         *     }
         * }
         * ```
         * &lt;!--End PulumiCodeChooser --&gt;
         * 
         * @return builder
         * 
         */
        public Builder secondaryFaultDomainName(@Nullable Output<String> secondaryFaultDomainName) {
            $.secondaryFaultDomainName = secondaryFaultDomainName;
            return this;
        }

        /**
         * @param secondaryFaultDomainName The name of second fault domain. Default is `Secondary`.
         * 
         * &gt; **NOTE:** You must disable vSphere HA before you enable vSAN on the cluster.
         * You can enable or re-enable vSphere HA after vSAN is configured.
         * 
         * &lt;!--Start PulumiCodeChooser --&gt;
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
         * import com.pulumi.vsphere.ComputeCluster;
         * import com.pulumi.vsphere.ComputeClusterArgs;
         * import com.pulumi.vsphere.inputs.ComputeClusterVsanDiskGroupArgs;
         * import com.pulumi.vsphere.inputs.ComputeClusterVsanFaultDomainArgs;
         * import com.pulumi.vsphere.inputs.ComputeClusterVsanStretchedClusterArgs;
         * import java.util.List;
         * import java.util.ArrayList;
         * import java.util.Map;
         * import java.io.File;
         * import java.nio.file.Files;
         * import java.nio.file.Paths;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *         var computeCluster = new ComputeCluster(&#34;computeCluster&#34;, ComputeClusterArgs.builder()        
         *             .datacenterId(data.vsphere_datacenter().datacenter().id())
         *             .hostSystemIds(data.vsphere_host().host().stream().map(element -&gt; element.id()).collect(toList()))
         *             .drsEnabled(true)
         *             .drsAutomationLevel(&#34;fullyAutomated&#34;)
         *             .haEnabled(false)
         *             .vsanEnabled(true)
         *             .vsanEsaEnabled(true)
         *             .vsanDedupEnabled(true)
         *             .vsanCompressionEnabled(true)
         *             .vsanPerformanceEnabled(true)
         *             .vsanVerboseModeEnabled(true)
         *             .vsanNetworkDiagnosticModeEnabled(true)
         *             .vsanUnmapEnabled(true)
         *             .vsanDitEncryptionEnabled(true)
         *             .vsanDitRekeyInterval(1800)
         *             .vsanDiskGroups(ComputeClusterVsanDiskGroupArgs.builder()
         *                 .cache(data.vsphere_vmfs_disks().cache_disks()[0])
         *                 .storages(data.vsphere_vmfs_disks().storage_disks())
         *                 .build())
         *             .vsanFaultDomains(ComputeClusterVsanFaultDomainArgs.builder()
         *                 .faultDomains(                
         *                     ComputeClusterVsanFaultDomainFaultDomainArgs.builder()
         *                         .name(&#34;fd1&#34;)
         *                         .hostIds(data.vsphere_host().faultdomain1_hosts().stream().map(element -&gt; element.id()).collect(toList()))
         *                         .build(),
         *                     ComputeClusterVsanFaultDomainFaultDomainArgs.builder()
         *                         .name(&#34;fd2&#34;)
         *                         .hostIds(data.vsphere_host().faultdomain2_hosts().stream().map(element -&gt; element.id()).collect(toList()))
         *                         .build())
         *                 .build())
         *             .vsanStretchedCluster(ComputeClusterVsanStretchedClusterArgs.builder()
         *                 .preferredFaultDomainHostIds(data.vsphere_host().preferred_fault_domain_host().stream().map(element -&gt; element.id()).collect(toList()))
         *                 .secondaryFaultDomainHostIds(data.vsphere_host().secondary_fault_domain_host().stream().map(element -&gt; element.id()).collect(toList()))
         *                 .witnessNode(data.vsphere_host().witness_host().id())
         *                 .build())
         *             .build());
         * 
         *     }
         * }
         * ```
         * &lt;!--End PulumiCodeChooser --&gt;
         * 
         * @return builder
         * 
         */
        public Builder secondaryFaultDomainName(String secondaryFaultDomainName) {
            return secondaryFaultDomainName(Output.of(secondaryFaultDomainName));
        }

        /**
         * @param witnessNode The managed object IDs of the host selected as witness node when enable stretched cluster.
         * 
         * @return builder
         * 
         */
        public Builder witnessNode(Output<String> witnessNode) {
            $.witnessNode = witnessNode;
            return this;
        }

        /**
         * @param witnessNode The managed object IDs of the host selected as witness node when enable stretched cluster.
         * 
         * @return builder
         * 
         */
        public Builder witnessNode(String witnessNode) {
            return witnessNode(Output.of(witnessNode));
        }

        public ComputeClusterVsanStretchedClusterArgs build() {
            if ($.preferredFaultDomainHostIds == null) {
                throw new MissingRequiredPropertyException("ComputeClusterVsanStretchedClusterArgs", "preferredFaultDomainHostIds");
            }
            if ($.secondaryFaultDomainHostIds == null) {
                throw new MissingRequiredPropertyException("ComputeClusterVsanStretchedClusterArgs", "secondaryFaultDomainHostIds");
            }
            if ($.witnessNode == null) {
                throw new MissingRequiredPropertyException("ComputeClusterVsanStretchedClusterArgs", "witnessNode");
            }
            return $;
        }
    }

}
