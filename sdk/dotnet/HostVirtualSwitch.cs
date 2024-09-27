// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    /// <summary>
    /// The `vsphere.HostVirtualSwitch` resource can be used to manage vSphere
    /// standard switches on an ESXi host. These switches can be used as a backing for
    /// standard port groups, which can be managed by the
    /// `vsphere.HostPortGroup` resource.
    /// 
    /// For an overview on vSphere networking concepts, see [this
    /// page][ref-vsphere-net-concepts].
    /// 
    /// [ref-vsphere-net-concepts]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.networking.doc/GUID-2B11DBB8-CB3C-4AFF-8885-EFEA0FC562F4.html
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// An existing vSwitch can be imported into this resource by its ID.
    /// 
    /// The convention of the id is a prefix, the host system [managed objectID][docs-about-morefs], and the virtual switch
    /// 
    /// name. An example would be `tf-HostVirtualSwitch:host-10:vSwitchTerraformTest`.
    /// 
    /// Import can the be done via the following command:
    /// 
    /// [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
    /// 
    /// ```sh
    /// $ pulumi import vsphere:index/hostVirtualSwitch:HostVirtualSwitch switch tf-HostVirtualSwitch:host-10:vSwitchTerraformTest
    /// ```
    /// 
    /// The above would import the vSwitch named `vSwitchTerraformTest` that is located in the `host-10`
    /// 
    /// vSphere host.
    /// </summary>
    [VSphereResourceType("vsphere:index/hostVirtualSwitch:HostVirtualSwitch")]
    public partial class HostVirtualSwitch : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of active network adapters used for load balancing.
        /// </summary>
        [Output("activeNics")]
        public Output<ImmutableArray<string>> ActiveNics { get; private set; } = null!;

        /// <summary>
        /// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
        /// that of its own.
        /// </summary>
        [Output("allowForgedTransmits")]
        public Output<bool?> AllowForgedTransmits { get; private set; } = null!;

        /// <summary>
        /// Controls whether or not the Media Access Control (MAC) address can be changed.
        /// </summary>
        [Output("allowMacChanges")]
        public Output<bool?> AllowMacChanges { get; private set; } = null!;

        /// <summary>
        /// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
        /// </summary>
        [Output("allowPromiscuous")]
        public Output<bool?> AllowPromiscuous { get; private set; } = null!;

        /// <summary>
        /// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
        /// </summary>
        [Output("beaconInterval")]
        public Output<int?> BeaconInterval { get; private set; } = null!;

        /// <summary>
        /// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
        /// only.
        /// </summary>
        [Output("checkBeacon")]
        public Output<bool?> CheckBeacon { get; private set; } = null!;

        /// <summary>
        /// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
        /// </summary>
        [Output("failback")]
        public Output<bool?> Failback { get; private set; } = null!;

        /// <summary>
        /// The managed object ID of
        /// the host to set the virtual switch up on. Forces a new resource if changed.
        /// </summary>
        [Output("hostSystemId")]
        public Output<string> HostSystemId { get; private set; } = null!;

        /// <summary>
        /// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
        /// </summary>
        [Output("linkDiscoveryOperation")]
        public Output<string?> LinkDiscoveryOperation { get; private set; } = null!;

        /// <summary>
        /// The discovery protocol type. Valid values are cdp and lldp.
        /// </summary>
        [Output("linkDiscoveryProtocol")]
        public Output<string?> LinkDiscoveryProtocol { get; private set; } = null!;

        /// <summary>
        /// The maximum transmission unit (MTU) for the virtual
        /// switch. Default: `1500`.
        /// </summary>
        [Output("mtu")]
        public Output<int?> Mtu { get; private set; } = null!;

        /// <summary>
        /// The name of the virtual switch. Forces a new resource if
        /// changed.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The list of network adapters to bind to this virtual switch.
        /// </summary>
        [Output("networkAdapters")]
        public Output<ImmutableArray<string>> NetworkAdapters { get; private set; } = null!;

        /// <summary>
        /// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
        /// </summary>
        [Output("notifySwitches")]
        public Output<bool?> NotifySwitches { get; private set; } = null!;

        /// <summary>
        /// The number of ports to create with this
        /// virtual switch. Default: `128`.
        /// 
        /// &gt; **NOTE:** Changing the port count requires a reboot of the host. This provider
        /// will not restart the host for you.
        /// </summary>
        [Output("numberOfPorts")]
        public Output<int?> NumberOfPorts { get; private set; } = null!;

        /// <summary>
        /// The average bandwidth in bits per second if traffic shaping is enabled.
        /// </summary>
        [Output("shapingAverageBandwidth")]
        public Output<int?> ShapingAverageBandwidth { get; private set; } = null!;

        /// <summary>
        /// The maximum burst size allowed in bytes if traffic shaping is enabled.
        /// </summary>
        [Output("shapingBurstSize")]
        public Output<int?> ShapingBurstSize { get; private set; } = null!;

        /// <summary>
        /// Enable traffic shaping on this virtual switch or port group.
        /// </summary>
        [Output("shapingEnabled")]
        public Output<bool?> ShapingEnabled { get; private set; } = null!;

        /// <summary>
        /// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
        /// </summary>
        [Output("shapingPeakBandwidth")]
        public Output<int?> ShapingPeakBandwidth { get; private set; } = null!;

        /// <summary>
        /// List of standby network adapters used for failover.
        /// </summary>
        [Output("standbyNics")]
        public Output<ImmutableArray<string>> StandbyNics { get; private set; } = null!;

        /// <summary>
        /// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
        /// failover_explicit.
        /// </summary>
        [Output("teamingPolicy")]
        public Output<string?> TeamingPolicy { get; private set; } = null!;


        /// <summary>
        /// Create a HostVirtualSwitch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HostVirtualSwitch(string name, HostVirtualSwitchArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/hostVirtualSwitch:HostVirtualSwitch", name, args ?? new HostVirtualSwitchArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HostVirtualSwitch(string name, Input<string> id, HostVirtualSwitchState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/hostVirtualSwitch:HostVirtualSwitch", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HostVirtualSwitch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HostVirtualSwitch Get(string name, Input<string> id, HostVirtualSwitchState? state = null, CustomResourceOptions? options = null)
        {
            return new HostVirtualSwitch(name, id, state, options);
        }
    }

    public sealed class HostVirtualSwitchArgs : global::Pulumi.ResourceArgs
    {
        [Input("activeNics", required: true)]
        private InputList<string>? _activeNics;

        /// <summary>
        /// List of active network adapters used for load balancing.
        /// </summary>
        public InputList<string> ActiveNics
        {
            get => _activeNics ?? (_activeNics = new InputList<string>());
            set => _activeNics = value;
        }

        /// <summary>
        /// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
        /// that of its own.
        /// </summary>
        [Input("allowForgedTransmits")]
        public Input<bool>? AllowForgedTransmits { get; set; }

        /// <summary>
        /// Controls whether or not the Media Access Control (MAC) address can be changed.
        /// </summary>
        [Input("allowMacChanges")]
        public Input<bool>? AllowMacChanges { get; set; }

        /// <summary>
        /// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
        /// </summary>
        [Input("allowPromiscuous")]
        public Input<bool>? AllowPromiscuous { get; set; }

        /// <summary>
        /// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
        /// </summary>
        [Input("beaconInterval")]
        public Input<int>? BeaconInterval { get; set; }

        /// <summary>
        /// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
        /// only.
        /// </summary>
        [Input("checkBeacon")]
        public Input<bool>? CheckBeacon { get; set; }

        /// <summary>
        /// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
        /// </summary>
        [Input("failback")]
        public Input<bool>? Failback { get; set; }

        /// <summary>
        /// The managed object ID of
        /// the host to set the virtual switch up on. Forces a new resource if changed.
        /// </summary>
        [Input("hostSystemId", required: true)]
        public Input<string> HostSystemId { get; set; } = null!;

        /// <summary>
        /// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
        /// </summary>
        [Input("linkDiscoveryOperation")]
        public Input<string>? LinkDiscoveryOperation { get; set; }

        /// <summary>
        /// The discovery protocol type. Valid values are cdp and lldp.
        /// </summary>
        [Input("linkDiscoveryProtocol")]
        public Input<string>? LinkDiscoveryProtocol { get; set; }

        /// <summary>
        /// The maximum transmission unit (MTU) for the virtual
        /// switch. Default: `1500`.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// The name of the virtual switch. Forces a new resource if
        /// changed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkAdapters", required: true)]
        private InputList<string>? _networkAdapters;

        /// <summary>
        /// The list of network adapters to bind to this virtual switch.
        /// </summary>
        public InputList<string> NetworkAdapters
        {
            get => _networkAdapters ?? (_networkAdapters = new InputList<string>());
            set => _networkAdapters = value;
        }

        /// <summary>
        /// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
        /// </summary>
        [Input("notifySwitches")]
        public Input<bool>? NotifySwitches { get; set; }

        /// <summary>
        /// The number of ports to create with this
        /// virtual switch. Default: `128`.
        /// 
        /// &gt; **NOTE:** Changing the port count requires a reboot of the host. This provider
        /// will not restart the host for you.
        /// </summary>
        [Input("numberOfPorts")]
        public Input<int>? NumberOfPorts { get; set; }

        /// <summary>
        /// The average bandwidth in bits per second if traffic shaping is enabled.
        /// </summary>
        [Input("shapingAverageBandwidth")]
        public Input<int>? ShapingAverageBandwidth { get; set; }

        /// <summary>
        /// The maximum burst size allowed in bytes if traffic shaping is enabled.
        /// </summary>
        [Input("shapingBurstSize")]
        public Input<int>? ShapingBurstSize { get; set; }

        /// <summary>
        /// Enable traffic shaping on this virtual switch or port group.
        /// </summary>
        [Input("shapingEnabled")]
        public Input<bool>? ShapingEnabled { get; set; }

        /// <summary>
        /// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
        /// </summary>
        [Input("shapingPeakBandwidth")]
        public Input<int>? ShapingPeakBandwidth { get; set; }

        [Input("standbyNics")]
        private InputList<string>? _standbyNics;

        /// <summary>
        /// List of standby network adapters used for failover.
        /// </summary>
        public InputList<string> StandbyNics
        {
            get => _standbyNics ?? (_standbyNics = new InputList<string>());
            set => _standbyNics = value;
        }

        /// <summary>
        /// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
        /// failover_explicit.
        /// </summary>
        [Input("teamingPolicy")]
        public Input<string>? TeamingPolicy { get; set; }

        public HostVirtualSwitchArgs()
        {
        }
        public static new HostVirtualSwitchArgs Empty => new HostVirtualSwitchArgs();
    }

    public sealed class HostVirtualSwitchState : global::Pulumi.ResourceArgs
    {
        [Input("activeNics")]
        private InputList<string>? _activeNics;

        /// <summary>
        /// List of active network adapters used for load balancing.
        /// </summary>
        public InputList<string> ActiveNics
        {
            get => _activeNics ?? (_activeNics = new InputList<string>());
            set => _activeNics = value;
        }

        /// <summary>
        /// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
        /// that of its own.
        /// </summary>
        [Input("allowForgedTransmits")]
        public Input<bool>? AllowForgedTransmits { get; set; }

        /// <summary>
        /// Controls whether or not the Media Access Control (MAC) address can be changed.
        /// </summary>
        [Input("allowMacChanges")]
        public Input<bool>? AllowMacChanges { get; set; }

        /// <summary>
        /// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
        /// </summary>
        [Input("allowPromiscuous")]
        public Input<bool>? AllowPromiscuous { get; set; }

        /// <summary>
        /// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
        /// </summary>
        [Input("beaconInterval")]
        public Input<int>? BeaconInterval { get; set; }

        /// <summary>
        /// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
        /// only.
        /// </summary>
        [Input("checkBeacon")]
        public Input<bool>? CheckBeacon { get; set; }

        /// <summary>
        /// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
        /// </summary>
        [Input("failback")]
        public Input<bool>? Failback { get; set; }

        /// <summary>
        /// The managed object ID of
        /// the host to set the virtual switch up on. Forces a new resource if changed.
        /// </summary>
        [Input("hostSystemId")]
        public Input<string>? HostSystemId { get; set; }

        /// <summary>
        /// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
        /// </summary>
        [Input("linkDiscoveryOperation")]
        public Input<string>? LinkDiscoveryOperation { get; set; }

        /// <summary>
        /// The discovery protocol type. Valid values are cdp and lldp.
        /// </summary>
        [Input("linkDiscoveryProtocol")]
        public Input<string>? LinkDiscoveryProtocol { get; set; }

        /// <summary>
        /// The maximum transmission unit (MTU) for the virtual
        /// switch. Default: `1500`.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// The name of the virtual switch. Forces a new resource if
        /// changed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkAdapters")]
        private InputList<string>? _networkAdapters;

        /// <summary>
        /// The list of network adapters to bind to this virtual switch.
        /// </summary>
        public InputList<string> NetworkAdapters
        {
            get => _networkAdapters ?? (_networkAdapters = new InputList<string>());
            set => _networkAdapters = value;
        }

        /// <summary>
        /// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
        /// </summary>
        [Input("notifySwitches")]
        public Input<bool>? NotifySwitches { get; set; }

        /// <summary>
        /// The number of ports to create with this
        /// virtual switch. Default: `128`.
        /// 
        /// &gt; **NOTE:** Changing the port count requires a reboot of the host. This provider
        /// will not restart the host for you.
        /// </summary>
        [Input("numberOfPorts")]
        public Input<int>? NumberOfPorts { get; set; }

        /// <summary>
        /// The average bandwidth in bits per second if traffic shaping is enabled.
        /// </summary>
        [Input("shapingAverageBandwidth")]
        public Input<int>? ShapingAverageBandwidth { get; set; }

        /// <summary>
        /// The maximum burst size allowed in bytes if traffic shaping is enabled.
        /// </summary>
        [Input("shapingBurstSize")]
        public Input<int>? ShapingBurstSize { get; set; }

        /// <summary>
        /// Enable traffic shaping on this virtual switch or port group.
        /// </summary>
        [Input("shapingEnabled")]
        public Input<bool>? ShapingEnabled { get; set; }

        /// <summary>
        /// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
        /// </summary>
        [Input("shapingPeakBandwidth")]
        public Input<int>? ShapingPeakBandwidth { get; set; }

        [Input("standbyNics")]
        private InputList<string>? _standbyNics;

        /// <summary>
        /// List of standby network adapters used for failover.
        /// </summary>
        public InputList<string> StandbyNics
        {
            get => _standbyNics ?? (_standbyNics = new InputList<string>());
            set => _standbyNics = value;
        }

        /// <summary>
        /// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
        /// failover_explicit.
        /// </summary>
        [Input("teamingPolicy")]
        public Input<string>? TeamingPolicy { get; set; }

        public HostVirtualSwitchState()
        {
        }
        public static new HostVirtualSwitchState Empty => new HostVirtualSwitchState();
    }
}
