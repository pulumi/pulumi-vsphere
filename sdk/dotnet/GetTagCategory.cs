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
        /// The `vsphere.TagCategory` data source can be used to reference tag categories
        /// that are not managed by this provider. Its attributes are the same as the
        /// `vsphere.TagCategory` resource, and, like importing,
        /// the data source uses a name and category as search criteria. The `id` and other
        /// attributes are populated with the data found by the search.
        /// 
        /// &gt; **NOTE:** Tagging is not supported on direct ESXi hosts connections and
        /// requires vCenter Server.
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
        ///     var category = VSphere.GetTagCategory.Invoke(new()
        ///     {
        ///         Name = "example-category",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetTagCategoryResult> InvokeAsync(GetTagCategoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTagCategoryResult>("vsphere:index/getTagCategory:getTagCategory", args ?? new GetTagCategoryArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.TagCategory` data source can be used to reference tag categories
        /// that are not managed by this provider. Its attributes are the same as the
        /// `vsphere.TagCategory` resource, and, like importing,
        /// the data source uses a name and category as search criteria. The `id` and other
        /// attributes are populated with the data found by the search.
        /// 
        /// &gt; **NOTE:** Tagging is not supported on direct ESXi hosts connections and
        /// requires vCenter Server.
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
        ///     var category = VSphere.GetTagCategory.Invoke(new()
        ///     {
        ///         Name = "example-category",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTagCategoryResult> Invoke(GetTagCategoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTagCategoryResult>("vsphere:index/getTagCategory:getTagCategory", args ?? new GetTagCategoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTagCategoryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the tag category.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetTagCategoryArgs()
        {
        }
        public static new GetTagCategoryArgs Empty => new GetTagCategoryArgs();
    }

    public sealed class GetTagCategoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the tag category.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetTagCategoryInvokeArgs()
        {
        }
        public static new GetTagCategoryInvokeArgs Empty => new GetTagCategoryInvokeArgs();
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
