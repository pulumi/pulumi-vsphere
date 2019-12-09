// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static partial class Invokes
    {
        /// <summary>
        /// The `vsphere..DistributedVirtualSwitch` data source can be used to discover
        /// the ID and uplink data of a of a vSphere distributed virtual switch (DVS). This
        /// can then be used with resources or data sources that require a DVS, such as the
        /// [`vsphere..DistributedPortGroup`][distributed-port-group] resource, for which
        /// an example is shown below.
        /// 
        /// [distributed-port-group]: /docs/providers/vsphere/r/distributed_port_group.html
        /// 
        /// &gt; **NOTE:** This data source requires vCenter and is not available on direct
        /// ESXi connections.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/distributed_virtual_switch.html.markdown.
        /// </summary>
        public static Task<GetDistributedVirtualSwitchResult> GetDistributedVirtualSwitch(GetDistributedVirtualSwitchArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDistributedVirtualSwitchResult>("vsphere:index/getDistributedVirtualSwitch:getDistributedVirtualSwitch", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetDistributedVirtualSwitchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [managed object reference
        /// ID][docs-about-morefs] of the datacenter the DVS is located in. This can be
        /// omitted if the search path used in `name` is an absolute path. For default
        /// datacenters, use the id attribute from an empty `vsphere..Datacenter` data
        /// source.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// The name of the distributed virtual switch. This can be a
        /// name or path.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetDistributedVirtualSwitchArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetDistributedVirtualSwitchResult
    {
        public readonly string? DatacenterId;
        public readonly string Name;
        public readonly ImmutableArray<string> Uplinks;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDistributedVirtualSwitchResult(
            string? datacenterId,
            string name,
            ImmutableArray<string> uplinks,
            string id)
        {
            DatacenterId = datacenterId;
            Name = name;
            Uplinks = uplinks;
            Id = id;
        }
    }
}
