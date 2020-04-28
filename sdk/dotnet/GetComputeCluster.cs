// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetComputeCluster
    {
        public static Task<GetComputeClusterResult> InvokeAsync(GetComputeClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetComputeClusterResult>("vsphere:index/getComputeCluster:getComputeCluster", args ?? new GetComputeClusterArgs(), options.WithVersion());
    }


    public sealed class GetComputeClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The [managed object reference
        /// ID][docs-about-morefs] of the datacenter the cluster is located in.  This can
        /// be omitted if the search path used in `name` is an absolute path.  For
        /// default datacenters, use the id attribute from an empty `vsphere..Datacenter`
        /// data source.
        /// </summary>
        [Input("datacenterId")]
        public string? DatacenterId { get; set; }

        /// <summary>
        /// The name or absolute path to the cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetComputeClusterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetComputeClusterResult
    {
        public readonly string? DatacenterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourcePoolId;

        [OutputConstructor]
        private GetComputeClusterResult(
            string? datacenterId,

            string id,

            string name,

            string resourcePoolId)
        {
            DatacenterId = datacenterId;
            Id = id;
            Name = name;
            ResourcePoolId = resourcePoolId;
        }
    }
}
