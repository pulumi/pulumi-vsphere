// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetHost
    {
        /// <summary>
        /// The `vsphere..Host` data source can be used to discover the ID of a vSphere
        /// host. This can then be used with resources or data sources that require a host
        /// managed object reference ID.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetHostResult> InvokeAsync(GetHostArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHostResult>("vsphere:index/getHost:getHost", args ?? new GetHostArgs(), options.WithVersion());
    }


    public sealed class GetHostArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The [managed object reference
        /// ID][docs-about-morefs] of a datacenter.
        /// </summary>
        [Input("datacenterId", required: true)]
        public string DatacenterId { get; set; } = null!;

        /// <summary>
        /// The name of the host. This can be a name or path. Can be
        /// omitted if there is only one host in your inventory.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetHostArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHostResult
    {
        public readonly string DatacenterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        /// <summary>
        /// The [managed object ID][docs-about-morefs] of the host's
        /// root resource pool.
        /// </summary>
        public readonly string ResourcePoolId;

        [OutputConstructor]
        private GetHostResult(
            string datacenterId,

            string id,

            string? name,

            string resourcePoolId)
        {
            DatacenterId = datacenterId;
            Id = id;
            Name = name;
            ResourcePoolId = resourcePoolId;
        }
    }
}
