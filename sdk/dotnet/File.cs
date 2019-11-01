// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vsphere
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/file.html.markdown.
    /// </summary>
    public partial class File : Pulumi.CustomResource
    {
        /// <summary>
        /// Create directories in `destination_file`
        /// path parameter if any missing for copy operation.
        /// </summary>
        [Output("createDirectories")]
        public Output<bool?> CreateDirectories { get; private set; } = null!;

        /// <summary>
        /// The name of a datacenter in which the file will be
        /// uploaded to.
        /// </summary>
        [Output("datacenter")]
        public Output<string?> Datacenter { get; private set; } = null!;

        /// <summary>
        /// The name of the datastore in which to upload the
        /// file to.
        /// </summary>
        [Output("datastore")]
        public Output<string> Datastore { get; private set; } = null!;

        /// <summary>
        /// The path to where the file should be uploaded
        /// or copied to on vSphere.
        /// </summary>
        [Output("destinationFile")]
        public Output<string> DestinationFile { get; private set; } = null!;

        /// <summary>
        /// The name of a datacenter in which the file
        /// will be copied from. Forces a new resource if changed.
        /// </summary>
        [Output("sourceDatacenter")]
        public Output<string?> SourceDatacenter { get; private set; } = null!;

        /// <summary>
        /// The name of the datastore in which file will
        /// be copied from. Forces a new resource if changed.
        /// </summary>
        [Output("sourceDatastore")]
        public Output<string?> SourceDatastore { get; private set; } = null!;

        [Output("sourceFile")]
        public Output<string> SourceFile { get; private set; } = null!;


        /// <summary>
        /// Create a File resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public File(string name, FileArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/file:File", name, args, MakeResourceOptions(options, ""))
        {
        }

        private File(string name, Input<string> id, FileState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/file:File", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing File resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static File Get(string name, Input<string> id, FileState? state = null, CustomResourceOptions? options = null)
        {
            return new File(name, id, state, options);
        }
    }

    public sealed class FileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create directories in `destination_file`
        /// path parameter if any missing for copy operation.
        /// </summary>
        [Input("createDirectories")]
        public Input<bool>? CreateDirectories { get; set; }

        /// <summary>
        /// The name of a datacenter in which the file will be
        /// uploaded to.
        /// </summary>
        [Input("datacenter")]
        public Input<string>? Datacenter { get; set; }

        /// <summary>
        /// The name of the datastore in which to upload the
        /// file to.
        /// </summary>
        [Input("datastore", required: true)]
        public Input<string> Datastore { get; set; } = null!;

        /// <summary>
        /// The path to where the file should be uploaded
        /// or copied to on vSphere.
        /// </summary>
        [Input("destinationFile", required: true)]
        public Input<string> DestinationFile { get; set; } = null!;

        /// <summary>
        /// The name of a datacenter in which the file
        /// will be copied from. Forces a new resource if changed.
        /// </summary>
        [Input("sourceDatacenter")]
        public Input<string>? SourceDatacenter { get; set; }

        /// <summary>
        /// The name of the datastore in which file will
        /// be copied from. Forces a new resource if changed.
        /// </summary>
        [Input("sourceDatastore")]
        public Input<string>? SourceDatastore { get; set; }

        [Input("sourceFile", required: true)]
        public Input<string> SourceFile { get; set; } = null!;

        public FileArgs()
        {
        }
    }

    public sealed class FileState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create directories in `destination_file`
        /// path parameter if any missing for copy operation.
        /// </summary>
        [Input("createDirectories")]
        public Input<bool>? CreateDirectories { get; set; }

        /// <summary>
        /// The name of a datacenter in which the file will be
        /// uploaded to.
        /// </summary>
        [Input("datacenter")]
        public Input<string>? Datacenter { get; set; }

        /// <summary>
        /// The name of the datastore in which to upload the
        /// file to.
        /// </summary>
        [Input("datastore")]
        public Input<string>? Datastore { get; set; }

        /// <summary>
        /// The path to where the file should be uploaded
        /// or copied to on vSphere.
        /// </summary>
        [Input("destinationFile")]
        public Input<string>? DestinationFile { get; set; }

        /// <summary>
        /// The name of a datacenter in which the file
        /// will be copied from. Forces a new resource if changed.
        /// </summary>
        [Input("sourceDatacenter")]
        public Input<string>? SourceDatacenter { get; set; }

        /// <summary>
        /// The name of the datastore in which file will
        /// be copied from. Forces a new resource if changed.
        /// </summary>
        [Input("sourceDatastore")]
        public Input<string>? SourceDatastore { get; set; }

        [Input("sourceFile")]
        public Input<string>? SourceFile { get; set; }

        public FileState()
        {
        }
    }
}
