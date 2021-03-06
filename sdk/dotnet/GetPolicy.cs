// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetPolicy
    {
        /// <summary>
        /// The `vsphere.getPolicy` data source can be used to discover the UUID of a
        /// vSphere storage policy. This can then be used with resources or data sources that
        /// require a storage policy.
        /// 
        /// &gt; **NOTE:** Storage policy support is unsupported on direct ESXi connections and
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
        ///         var policy = Output.Create(VSphere.GetPolicy.InvokeAsync(new VSphere.GetPolicyArgs
        ///         {
        ///             Name = "policy1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPolicyResult> InvokeAsync(GetPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyResult>("vsphere:index/getPolicy:getPolicy", args ?? new GetPolicyArgs(), options.WithVersion());
    }


    public sealed class GetPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the storage policy.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetPolicyResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
