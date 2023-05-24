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
    /// Provides a VMware vSphere vnic resource.
    /// 
    /// ## Example Usage
    /// 
    /// ### S
    /// ### Create a vnic attached to a distributed virtual switch using the vmotion TCP/IP stack
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using VSphere = Pulumi.VSphere;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var dc = VSphere.GetDatacenter.Invoke(new()
    ///     {
    ///         Name = "mydc",
    ///     });
    /// 
    ///     var h1 = VSphere.GetHost.Invoke(new()
    ///     {
    ///         Name = "esxi1.host.test",
    ///         DatacenterId = dc.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
    ///     });
    /// 
    ///     var d1 = new VSphere.DistributedVirtualSwitch("d1", new()
    ///     {
    ///         DatacenterId = dc.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
    ///         Hosts = new[]
    ///         {
    ///             new VSphere.Inputs.DistributedVirtualSwitchHostArgs
    ///             {
    ///                 HostSystemId = h1.Apply(getHostResult =&gt; getHostResult.Id),
    ///                 Devices = new[]
    ///                 {
    ///                     "vnic3",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var p1 = new VSphere.DistributedPortGroup("p1", new()
    ///     {
    ///         VlanId = 1234,
    ///         DistributedVirtualSwitchUuid = d1.Id,
    ///     });
    /// 
    ///     var v1 = new VSphere.Vnic("v1", new()
    ///     {
    ///         Host = h1.Apply(getHostResult =&gt; getHostResult.Id),
    ///         DistributedSwitchPort = d1.Id,
    ///         DistributedPortGroup = p1.Id,
    ///         Ipv4 = new VSphere.Inputs.VnicIpv4Args
    ///         {
    ///             Dhcp = true,
    ///         },
    ///         Netstack = "vmotion",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Importing
    /// 
    /// An existing vNic can be [imported][docs-import] into this resource
    /// via supplying the vNic's ID. An example is below:
    /// 
    /// [docs-import]: /docs/import/index.html
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    /// });
    /// ```
    /// 
    /// The above would import the vnic `vmk2` from host with ID `host-123`.
    /// </summary>
    [VSphereResourceType("vsphere:index/vnic:Vnic")]
    public partial class Vnic : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Key of the distributed portgroup the nic will connect to.
        /// </summary>
        [Output("distributedPortGroup")]
        public Output<string?> DistributedPortGroup { get; private set; } = null!;

        /// <summary>
        /// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
        /// </summary>
        [Output("distributedSwitchPort")]
        public Output<string?> DistributedSwitchPort { get; private set; } = null!;

        /// <summary>
        /// ESX host the interface belongs to
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
        /// </summary>
        [Output("ipv4")]
        public Output<Outputs.VnicIpv4?> Ipv4 { get; private set; } = null!;

        /// <summary>
        /// IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
        /// </summary>
        [Output("ipv6")]
        public Output<Outputs.VnicIpv6?> Ipv6 { get; private set; } = null!;

        /// <summary>
        /// MAC address of the interface.
        /// </summary>
        [Output("mac")]
        public Output<string> Mac { get; private set; } = null!;

        /// <summary>
        /// MTU of the interface.
        /// </summary>
        [Output("mtu")]
        public Output<int> Mtu { get; private set; } = null!;

        /// <summary>
        /// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
        /// </summary>
        [Output("netstack")]
        public Output<string?> Netstack { get; private set; } = null!;

        /// <summary>
        /// Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
        /// </summary>
        [Output("portgroup")]
        public Output<string?> Portgroup { get; private set; } = null!;

        /// <summary>
        /// Enabled services setting for this interface. Current possible values are 'vmotion', 'management', and 'vsan'.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<string>> Services { get; private set; } = null!;


        /// <summary>
        /// Create a Vnic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vnic(string name, VnicArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/vnic:Vnic", name, args ?? new VnicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vnic(string name, Input<string> id, VnicState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/vnic:Vnic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vnic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vnic Get(string name, Input<string> id, VnicState? state = null, CustomResourceOptions? options = null)
        {
            return new Vnic(name, id, state, options);
        }
    }

    public sealed class VnicArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key of the distributed portgroup the nic will connect to.
        /// </summary>
        [Input("distributedPortGroup")]
        public Input<string>? DistributedPortGroup { get; set; }

        /// <summary>
        /// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
        /// </summary>
        [Input("distributedSwitchPort")]
        public Input<string>? DistributedSwitchPort { get; set; }

        /// <summary>
        /// ESX host the interface belongs to
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
        /// </summary>
        [Input("ipv4")]
        public Input<Inputs.VnicIpv4Args>? Ipv4 { get; set; }

        /// <summary>
        /// IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
        /// </summary>
        [Input("ipv6")]
        public Input<Inputs.VnicIpv6Args>? Ipv6 { get; set; }

        /// <summary>
        /// MAC address of the interface.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// MTU of the interface.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
        /// </summary>
        [Input("netstack")]
        public Input<string>? Netstack { get; set; }

        /// <summary>
        /// Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
        /// </summary>
        [Input("portgroup")]
        public Input<string>? Portgroup { get; set; }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// Enabled services setting for this interface. Current possible values are 'vmotion', 'management', and 'vsan'.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        public VnicArgs()
        {
        }
        public static new VnicArgs Empty => new VnicArgs();
    }

    public sealed class VnicState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key of the distributed portgroup the nic will connect to.
        /// </summary>
        [Input("distributedPortGroup")]
        public Input<string>? DistributedPortGroup { get; set; }

        /// <summary>
        /// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
        /// </summary>
        [Input("distributedSwitchPort")]
        public Input<string>? DistributedSwitchPort { get; set; }

        /// <summary>
        /// ESX host the interface belongs to
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
        /// </summary>
        [Input("ipv4")]
        public Input<Inputs.VnicIpv4GetArgs>? Ipv4 { get; set; }

        /// <summary>
        /// IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
        /// </summary>
        [Input("ipv6")]
        public Input<Inputs.VnicIpv6GetArgs>? Ipv6 { get; set; }

        /// <summary>
        /// MAC address of the interface.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// MTU of the interface.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
        /// </summary>
        [Input("netstack")]
        public Input<string>? Netstack { get; set; }

        /// <summary>
        /// Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
        /// </summary>
        [Input("portgroup")]
        public Input<string>? Portgroup { get; set; }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// Enabled services setting for this interface. Current possible values are 'vmotion', 'management', and 'vsan'.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        public VnicState()
        {
        }
        public static new VnicState Empty => new VnicState();
    }
}
