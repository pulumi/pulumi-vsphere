// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetDatacenter
    {
        /// <summary>
        /// The `vsphere.Datacenter` data source can be used to discover the ID of a
        /// vSphere datacenter object. This can then be used with resources or data sources
        /// that require a datacenter, such as the `vsphere.Host`
        /// data source.
        /// 
        /// ## Example Usage
        /// 
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
        /// });
        /// ```
        /// </summary>
        public static Task<GetDatacenterResult> InvokeAsync(GetDatacenterArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatacenterResult>("vsphere:index/getDatacenter:getDatacenter", args ?? new GetDatacenterArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.Datacenter` data source can be used to discover the ID of a
        /// vSphere datacenter object. This can then be used with resources or data sources
        /// that require a datacenter, such as the `vsphere.Host`
        /// data source.
        /// 
        /// ## Example Usage
        /// 
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
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatacenterResult> Invoke(GetDatacenterInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatacenterResult>("vsphere:index/getDatacenter:getDatacenter", args ?? new GetDatacenterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.Datacenter` data source can be used to discover the ID of a
        /// vSphere datacenter object. This can then be used with resources or data sources
        /// that require a datacenter, such as the `vsphere.Host`
        /// data source.
        /// 
        /// ## Example Usage
        /// 
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
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatacenterResult> Invoke(GetDatacenterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatacenterResult>("vsphere:index/getDatacenter:getDatacenter", args ?? new GetDatacenterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatacenterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the datacenter. This can be a name or path.
        /// Can be omitted if there is only one datacenter in the inventory.
        /// 
        /// &gt; **NOTE:** When used with an ESXi host, this data source _always_ returns the
        /// host's "default" datacenter, which is a special datacenter name unrelated to the
        /// datacenters that exist in the vSphere inventory when managed by a vCenter Server
        /// instance. Hence, the `name` attribute is completely ignored.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetDatacenterArgs()
        {
        }
        public static new GetDatacenterArgs Empty => new GetDatacenterArgs();
    }

    public sealed class GetDatacenterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the datacenter. This can be a name or path.
        /// Can be omitted if there is only one datacenter in the inventory.
        /// 
        /// &gt; **NOTE:** When used with an ESXi host, this data source _always_ returns the
        /// host's "default" datacenter, which is a special datacenter name unrelated to the
        /// datacenters that exist in the vSphere inventory when managed by a vCenter Server
        /// instance. Hence, the `name` attribute is completely ignored.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetDatacenterInvokeArgs()
        {
        }
        public static new GetDatacenterInvokeArgs Empty => new GetDatacenterInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatacenterResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        /// <summary>
        /// List of all virtual machines included in the vSphere datacenter object.
        /// </summary>
        public readonly ImmutableArray<string> VirtualMachines;

        [OutputConstructor]
        private GetDatacenterResult(
            string id,

            string? name,

            ImmutableArray<string> virtualMachines)
        {
            Id = id;
            Name = name;
            VirtualMachines = virtualMachines;
        }
    }
}
