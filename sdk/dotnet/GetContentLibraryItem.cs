// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetContentLibraryItem
    {
        /// <summary>
        /// The `vsphere.ContentLibraryItem` data source can be used to discover the ID
        /// of a content library item.
        /// 
        /// &gt; **NOTE:** This resource requires vCenter Server and is not available on
        /// direct ESXi host connections.
        /// </summary>
        public static Task<GetContentLibraryItemResult> InvokeAsync(GetContentLibraryItemArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContentLibraryItemResult>("vsphere:index/getContentLibraryItem:getContentLibraryItem", args ?? new GetContentLibraryItemArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.ContentLibraryItem` data source can be used to discover the ID
        /// of a content library item.
        /// 
        /// &gt; **NOTE:** This resource requires vCenter Server and is not available on
        /// direct ESXi host connections.
        /// </summary>
        public static Output<GetContentLibraryItemResult> Invoke(GetContentLibraryItemInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContentLibraryItemResult>("vsphere:index/getContentLibraryItem:getContentLibraryItem", args ?? new GetContentLibraryItemInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContentLibraryItemArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the content library in which the item exists.
        /// </summary>
        [Input("libraryId", required: true)]
        public string LibraryId { get; set; } = null!;

        /// <summary>
        /// The name of the content library item.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The type for the content library item. One of `ovf`, `vm-template`, or `iso`
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetContentLibraryItemArgs()
        {
        }
        public static new GetContentLibraryItemArgs Empty => new GetContentLibraryItemArgs();
    }

    public sealed class GetContentLibraryItemInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the content library in which the item exists.
        /// </summary>
        [Input("libraryId", required: true)]
        public Input<string> LibraryId { get; set; } = null!;

        /// <summary>
        /// The name of the content library item.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The type for the content library item. One of `ovf`, `vm-template`, or `iso`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetContentLibraryItemInvokeArgs()
        {
        }
        public static new GetContentLibraryItemInvokeArgs Empty => new GetContentLibraryItemInvokeArgs();
    }


    [OutputType]
    public sealed class GetContentLibraryItemResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LibraryId;
        public readonly string Name;
        public readonly string Type;

        [OutputConstructor]
        private GetContentLibraryItemResult(
            string id,

            string libraryId,

            string name,

            string type)
        {
            Id = id;
            LibraryId = libraryId;
            Name = name;
            Type = type;
        }
    }
}
