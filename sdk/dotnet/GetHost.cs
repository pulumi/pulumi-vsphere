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
        /// The `vsphere..Host` data source can be used to discover the ID of a vSphere
        /// host. This can then be used with resources or data sources that require a host
        /// managed object reference ID.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/host.html.markdown.
        /// </summary>
        public static Task<GetHostResult> GetHost(GetHostArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHostResult>("vsphere:index/getHost:getHost", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetHostArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [managed object reference
        /// ID][docs-about-morefs] of a datacenter.
        /// </summary>
        [Input("datacenterId", required: true)]
        public Input<string> DatacenterId { get; set; } = null!;

        /// <summary>
        /// The name of the host. This can be a name or path. Can be
        /// omitted if there is only one host in your inventory.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetHostArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetHostResult
    {
        public readonly string DatacenterId;
        public readonly string? Name;
        /// <summary>
        /// The [managed object ID][docs-about-morefs] of the host's
        /// root resource pool.
        /// </summary>
        public readonly string ResourcePoolId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetHostResult(
            string datacenterId,
            string? name,
            string resourcePoolId,
            string id)
        {
            DatacenterId = datacenterId;
            Name = name;
            ResourcePoolId = resourcePoolId;
            Id = id;
        }
    }
}
