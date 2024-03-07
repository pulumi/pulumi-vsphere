// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetVmfsDisks
    {
        /// <summary>
        /// The `vsphere.getVmfsDisks` data source can be used to discover the storage
        /// devices available on an ESXi host. This data source can be combined with the
        /// `vsphere.VmfsDatastore` resource to create VMFS
        /// datastores based off a set of discovered disks.
        /// 
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var datacenter = VSphere.GetDatacenter.Invoke(new()
        ///     {
        ///         Name = "dc-01",
        ///     });
        /// 
        ///     var host = VSphere.GetHost.Invoke(new()
        ///     {
        ///         Name = "esxi-01.example.com",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        ///     var vmfsDisks = VSphere.GetVmfsDisks.Invoke(new()
        ///     {
        ///         HostSystemId = host.Apply(getHostResult =&gt; getHostResult.Id),
        ///         Rescan = true,
        ///         Filter = "mpx.vmhba1:C0:T[12]:L0",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetVmfsDisksResult> InvokeAsync(GetVmfsDisksArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVmfsDisksResult>("vsphere:index/getVmfsDisks:getVmfsDisks", args ?? new GetVmfsDisksArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.getVmfsDisks` data source can be used to discover the storage
        /// devices available on an ESXi host. This data source can be combined with the
        /// `vsphere.VmfsDatastore` resource to create VMFS
        /// datastores based off a set of discovered disks.
        /// 
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var datacenter = VSphere.GetDatacenter.Invoke(new()
        ///     {
        ///         Name = "dc-01",
        ///     });
        /// 
        ///     var host = VSphere.GetHost.Invoke(new()
        ///     {
        ///         Name = "esxi-01.example.com",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        ///     var vmfsDisks = VSphere.GetVmfsDisks.Invoke(new()
        ///     {
        ///         HostSystemId = host.Apply(getHostResult =&gt; getHostResult.Id),
        ///         Rescan = true,
        ///         Filter = "mpx.vmhba1:C0:T[12]:L0",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetVmfsDisksResult> Invoke(GetVmfsDisksInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVmfsDisksResult>("vsphere:index/getVmfsDisks:getVmfsDisks", args ?? new GetVmfsDisksInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVmfsDisksArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A regular expression to filter the disks against. Only
        /// disks with canonical names that match will be included.
        /// 
        /// &gt; **NOTE:** Using a `filter` is recommended if there is any chance the host
        /// will have any specific storage devices added to it that may affect the order of
        /// the output `disks` attribute below, which is lexicographically sorted.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// The managed object ID of
        /// the host to look for disks on.
        /// </summary>
        [Input("hostSystemId", required: true)]
        public string HostSystemId { get; set; } = null!;

        /// <summary>
        /// Whether or not to rescan storage adapters before
        /// searching for disks. This may lengthen the time it takes to perform the
        /// search. Default: `false`.
        /// </summary>
        [Input("rescan")]
        public bool? Rescan { get; set; }

        public GetVmfsDisksArgs()
        {
        }
        public static new GetVmfsDisksArgs Empty => new GetVmfsDisksArgs();
    }

    public sealed class GetVmfsDisksInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A regular expression to filter the disks against. Only
        /// disks with canonical names that match will be included.
        /// 
        /// &gt; **NOTE:** Using a `filter` is recommended if there is any chance the host
        /// will have any specific storage devices added to it that may affect the order of
        /// the output `disks` attribute below, which is lexicographically sorted.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The managed object ID of
        /// the host to look for disks on.
        /// </summary>
        [Input("hostSystemId", required: true)]
        public Input<string> HostSystemId { get; set; } = null!;

        /// <summary>
        /// Whether or not to rescan storage adapters before
        /// searching for disks. This may lengthen the time it takes to perform the
        /// search. Default: `false`.
        /// </summary>
        [Input("rescan")]
        public Input<bool>? Rescan { get; set; }

        public GetVmfsDisksInvokeArgs()
        {
        }
        public static new GetVmfsDisksInvokeArgs Empty => new GetVmfsDisksInvokeArgs();
    }


    [OutputType]
    public sealed class GetVmfsDisksResult
    {
        /// <summary>
        /// A lexicographically sorted list of devices discovered by the
        /// operation, matching the supplied `filter`, if provided.
        /// </summary>
        public readonly ImmutableArray<string> Disks;
        public readonly string? Filter;
        public readonly string HostSystemId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? Rescan;

        [OutputConstructor]
        private GetVmfsDisksResult(
            ImmutableArray<string> disks,

            string? filter,

            string hostSystemId,

            string id,

            bool? rescan)
        {
            Disks = disks;
            Filter = filter;
            HostSystemId = hostSystemId;
            Id = id;
            Rescan = rescan;
        }
    }
}
