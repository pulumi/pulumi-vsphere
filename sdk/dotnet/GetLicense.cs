// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetLicense
    {
        /// <summary>
        /// The `vsphere.License` data source can be used to get the general attributes of
        /// a license keys from a vCenter Server instance.
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
        ///         var license = Output.Create(VSphere.GetLicense.InvokeAsync(new VSphere.GetLicenseArgs
        ///         {
        ///             LicenseKey = "00000-00000-00000-00000-00000",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLicenseResult> InvokeAsync(GetLicenseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLicenseResult>("vsphere:index/getLicense:getLicense", args ?? new GetLicenseArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.License` data source can be used to get the general attributes of
        /// a license keys from a vCenter Server instance.
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
        ///         var license = Output.Create(VSphere.GetLicense.InvokeAsync(new VSphere.GetLicenseArgs
        ///         {
        ///             LicenseKey = "00000-00000-00000-00000-00000",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetLicenseResult> Invoke(GetLicenseInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLicenseResult>("vsphere:index/getLicense:getLicense", args ?? new GetLicenseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLicenseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The license key.
        /// </summary>
        [Input("licenseKey", required: true)]
        public string LicenseKey { get; set; } = null!;

        public GetLicenseArgs()
        {
        }
    }

    public sealed class GetLicenseInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The license key.
        /// </summary>
        [Input("licenseKey", required: true)]
        public Input<string> LicenseKey { get; set; } = null!;

        public GetLicenseInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLicenseResult
    {
        /// <summary>
        /// The product edition of the license key.
        /// </summary>
        public readonly string EditionKey;
        public readonly string Id;
        /// <summary>
        /// A map of key/value pairs attached as labels (tags) to the license key.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string LicenseKey;
        /// <summary>
        /// The display name for the license.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Total number of units (example: CPUs) contained in the license.
        /// </summary>
        public readonly int Total;
        /// <summary>
        /// The number of units (example: CPUs) assigned to this license.
        /// </summary>
        public readonly int Used;

        [OutputConstructor]
        private GetLicenseResult(
            string editionKey,

            string id,

            ImmutableDictionary<string, string> labels,

            string licenseKey,

            string name,

            int total,

            int used)
        {
            EditionKey = editionKey;
            Id = id;
            Labels = labels;
            LicenseKey = licenseKey;
            Name = name;
            Total = total;
            Used = used;
        }
    }
}
