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
        /// The `vsphere..VappContainer` data source can be used to discover the ID of a
        /// vApp container in vSphere. This is useful to fetch the ID of a vApp container
        /// that you want to use to create virtual machines in using the
        /// [`vsphere..VirtualMachine`][docs-virtual-machine-resource] resource. 
        /// 
        /// [docs-virtual-machine-resource]: /docs/providers/vsphere/r/virtual_machine.html
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/vapp_container.html.markdown.
        /// </summary>
        public static Task<GetVappContainerResult> GetVappContainer(GetVappContainerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVappContainerResult>("vsphere:index/getVappContainer:getVappContainer", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetVappContainerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [managed object reference
        /// ID][docs-about-morefs] of the datacenter the vApp container is located in.
        /// </summary>
        [Input("datacenterId", required: true)]
        public Input<string> DatacenterId { get; set; } = null!;

        /// <summary>
        /// The name of the vApp container. This can be a name or
        /// path.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetVappContainerArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetVappContainerResult
    {
        public readonly string DatacenterId;
        public readonly string Name;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetVappContainerResult(
            string datacenterId,
            string name,
            string id)
        {
            DatacenterId = datacenterId;
            Name = name;
            Id = id;
        }
    }
}
