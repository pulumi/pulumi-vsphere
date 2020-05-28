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
        /// The `vsphere..CustomAttribute` data source can be used to reference custom 
        /// attributes that are not managed by this provider. Its attributes are exactly the 
        /// same as the `vsphere..CustomAttribute` resource, 
        /// and, like importing, the data source takes a name to search on. The `id` and 
        /// other attributes are then populated with the data found by the search.
        /// 
        /// &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections 
        /// and require vCenter.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var attribute = Output.Create(VSphere.GetCustomAttribute.InvokeAsync(new VSphere.GetCustomAttributeArgs
        ///         {
        ///             Name = "test-attribute",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCustomAttributeResult> InvokeAsync(GetCustomAttributeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCustomAttributeResult>("vsphere:index/getCustomAttribute:getCustomAttribute", args ?? new GetCustomAttributeArgs(), options.WithVersion());
    }


    public sealed class GetCustomAttributeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the custom attribute.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetCustomAttributeArgs()
        {
        }
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
