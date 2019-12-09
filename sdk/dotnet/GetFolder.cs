// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static partial class Invokes
    {
        /// <summary>
        /// The `vsphere..Folder` data source can be used to get the general attributes of a
        /// vSphere inventory folder. Paths are absolute and include must include the
        /// datacenter.  
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/folder.html.markdown.
        /// </summary>
        public static Task<GetFolderResult> GetFolder(GetFolderArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFolderResult>("vsphere:index/getFolder:getFolder", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetFolderArgs : Pulumi.ResourceArgs
    {
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public GetFolderArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetFolderResult
    {
        public readonly string Path;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetFolderResult(
            string path,
            string id)
        {
            Path = path;
            Id = id;
        }
    }
}
