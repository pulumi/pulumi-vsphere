// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static partial class Invokes
    {
        [Obsolete("Use GetTag.InvokeAsync() instead")]
        public static Task<GetTagResult> GetTag(GetTagArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagResult>("vsphere:index/getTag:getTag", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetTag
    {
        public static Task<GetTagResult> InvokeAsync(GetTagArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagResult>("vsphere:index/getTag:getTag", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetTagArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the tag category the tag is located in.
        /// </summary>
        [Input("categoryId", required: true)]
        public string CategoryId { get; set; } = null!;

        /// <summary>
        /// The name of the tag.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetTagArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetTagResult
    {
        public readonly string CategoryId;
        public readonly string Description;
        public readonly string Name;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetTagResult(
            string categoryId,
            string description,
            string name,
            string id)
        {
            CategoryId = categoryId;
            Description = description;
            Name = name;
            Id = id;
        }
    }
}
