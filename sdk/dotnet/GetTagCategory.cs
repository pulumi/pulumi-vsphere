// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetTagCategory
    {
        /// <summary>
        /// The `vsphere..TagCategory` data source can be used to reference tag categories
        /// that are not managed by this provider. Its attributes are exactly the same as the
        /// `vsphere..TagCategory` resource, and, like importing,
        /// the data source takes a name to search on. The `id` and other attributes are
        /// then populated with the data found by the search.
        /// 
        /// &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
        /// requires vCenter 6.0 or higher.
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
        ///         var category = Output.Create(VSphere.GetTagCategory.InvokeAsync(new VSphere.GetTagCategoryArgs
        ///         {
        ///             Name = "test-category",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTagCategoryResult> InvokeAsync(GetTagCategoryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagCategoryResult>("vsphere:index/getTagCategory:getTagCategory", args ?? new GetTagCategoryArgs(), options.WithVersion());
    }


    public sealed class GetTagCategoryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the tag category.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetTagCategoryArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTagCategoryResult
    {
        public readonly ImmutableArray<string> AssociableTypes;
        public readonly string Cardinality;
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetTagCategoryResult(
            ImmutableArray<string> associableTypes,

            string cardinality,

            string description,

            string id,

            string name)
        {
            AssociableTypes = associableTypes;
            Cardinality = cardinality;
            Description = description;
            Id = id;
            Name = name;
        }
    }
}
