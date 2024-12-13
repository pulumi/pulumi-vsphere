// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetCustomAttribute
    {
        /// <summary>
        /// The `vsphere.CustomAttribute` data source can be used to reference custom
        /// attributes that are not managed by this provider. Its attributes are exactly the
        /// same as the `vsphere.CustomAttribute` resource,
        /// and, like importing, the data source takes a name argument for the search. The
        /// `id` and other attributes are then populated with the data found by the search.
        /// 
        /// &gt; **NOTE:** Custom attributes are unsupported on direct ESXi host connections
        /// and require vCenter Server.
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
        ///     var attribute = VSphere.GetCustomAttribute.Invoke(new()
        ///     {
        ///         Name = "test-attribute",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCustomAttributeResult> InvokeAsync(GetCustomAttributeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCustomAttributeResult>("vsphere:index/getCustomAttribute:getCustomAttribute", args ?? new GetCustomAttributeArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.CustomAttribute` data source can be used to reference custom
        /// attributes that are not managed by this provider. Its attributes are exactly the
        /// same as the `vsphere.CustomAttribute` resource,
        /// and, like importing, the data source takes a name argument for the search. The
        /// `id` and other attributes are then populated with the data found by the search.
        /// 
        /// &gt; **NOTE:** Custom attributes are unsupported on direct ESXi host connections
        /// and require vCenter Server.
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
        ///     var attribute = VSphere.GetCustomAttribute.Invoke(new()
        ///     {
        ///         Name = "test-attribute",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCustomAttributeResult> Invoke(GetCustomAttributeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomAttributeResult>("vsphere:index/getCustomAttribute:getCustomAttribute", args ?? new GetCustomAttributeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.CustomAttribute` data source can be used to reference custom
        /// attributes that are not managed by this provider. Its attributes are exactly the
        /// same as the `vsphere.CustomAttribute` resource,
        /// and, like importing, the data source takes a name argument for the search. The
        /// `id` and other attributes are then populated with the data found by the search.
        /// 
        /// &gt; **NOTE:** Custom attributes are unsupported on direct ESXi host connections
        /// and require vCenter Server.
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
        ///     var attribute = VSphere.GetCustomAttribute.Invoke(new()
        ///     {
        ///         Name = "test-attribute",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCustomAttributeResult> Invoke(GetCustomAttributeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomAttributeResult>("vsphere:index/getCustomAttribute:getCustomAttribute", args ?? new GetCustomAttributeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomAttributeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the custom attribute.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetCustomAttributeArgs()
        {
        }
        public static new GetCustomAttributeArgs Empty => new GetCustomAttributeArgs();
    }

    public sealed class GetCustomAttributeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the custom attribute.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetCustomAttributeInvokeArgs()
        {
        }
        public static new GetCustomAttributeInvokeArgs Empty => new GetCustomAttributeInvokeArgs();
    }


    [OutputType]
    public sealed class GetCustomAttributeResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ManagedObjectType;
        public readonly string Name;

        [OutputConstructor]
        private GetCustomAttributeResult(
            string id,

            string managedObjectType,

            string name)
        {
            Id = id;
            ManagedObjectType = managedObjectType;
            Name = name;
        }
    }
}
