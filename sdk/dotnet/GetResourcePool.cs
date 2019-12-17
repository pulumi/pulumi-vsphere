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
        /// The `vsphere..ResourcePool` data source can be used to discover the ID of a
        /// resource pool in vSphere. This is useful to fetch the ID of a resource pool
        /// that you want to use to create virtual machines in using the
        /// [`vsphere..VirtualMachine`][docs-virtual-machine-resource] resource. 
        /// 
        /// [docs-virtual-machine-resource]: /docs/providers/vsphere/r/virtual_machine.html
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/resource_pool.html.markdown.
        /// </summary>
        public static Task<GetResourcePoolResult> GetResourcePool(GetResourcePoolArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourcePoolResult>("vsphere:index/getResourcePool:getResourcePool", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetResourcePoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [managed object reference
        /// ID][docs-about-morefs] of the datacenter the resource pool is located in.
        /// This can be omitted if the search path used in `name` is an absolute path.
        /// For default datacenters, use the id attribute from an empty
        /// `vsphere..Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// The name of the resource pool. This can be a name or
        /// path. This is required when using vCenter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetResourcePoolArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetResourcePoolResult
    {
        public readonly string? DatacenterId;
        public readonly string? Name;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetResourcePoolResult(
            string? datacenterId,
            string? name,
            string id)
        {
            DatacenterId = datacenterId;
            Name = name;
            Id = id;
        }
    }
}
