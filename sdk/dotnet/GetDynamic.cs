// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetDynamic
    {
        /// <summary>
        /// [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
        /// 
        /// The `vsphere..getDynamic` data source can be used to get the [managed object 
        ///   reference ID][docs-about-morefs] of any tagged managed object in vCenter
        ///   by providing a list of tag IDs and an optional regular expression to filter
        ///   objects by name.
        ///    
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDynamicResult> InvokeAsync(GetDynamicArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDynamicResult>("vsphere:index/getDynamic:getDynamic", args ?? new GetDynamicArgs(), options.WithVersion());
    }


    public sealed class GetDynamicArgs : Pulumi.InvokeArgs
    {
        [Input("filters", required: true)]
        private List<string>? _filters;

        /// <summary>
        /// A list of tag IDs that must be present on an object to
        /// be a match.
        /// </summary>
        public List<string> Filters
        {
            get => _filters ?? (_filters = new List<string>());
            set => _filters = value;
        }

        /// <summary>
        /// A regular expression that will be used to match
        /// the object's name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// The managed object type the returned object must match.
        /// For a full list, click [here](https://code.vmware.com/apis/196/vsphere).
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetDynamicArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDynamicResult
    {
        public readonly ImmutableArray<string> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameRegex;
        public readonly string? Type;

        [OutputConstructor]
        private GetDynamicResult(
            ImmutableArray<string> filters,

            string id,

            string? nameRegex,

            string? type)
        {
            Filters = filters;
            Id = id;
            NameRegex = nameRegex;
            Type = type;
        }
    }
}
