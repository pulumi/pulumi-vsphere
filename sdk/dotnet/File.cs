// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    /// <summary>
    /// ## Example Usage
    /// </summary>
    [VSphereResourceType("vsphere:index/file:File")]
    public partial class File : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Create directories in `destination_file`
        /// path parameter on first apply if any are missing for copy operation.
        /// 
        /// &gt; **NOTE:** Any directory created as part of the `create_directories` argument
        /// will not be deleted when the resource is destroyed. New directories are not
        /// created if the `destination_file` path is changed in subsequent applies.
        /// </summary>
        [Output("createDirectories")]
        public Output<bool?> CreateDirectories { get; private set; } = null!;

        /// <summary>
        /// The name of a datacenter to which the file will be
        /// uploaded.
        /// </summary>
        [Output("datacenter")]
        public Output<string?> Datacenter { get; private set; } = null!;

        /// <summary>
        /// The name of the datastore to which to upload the
        /// file.
        /// </summary>
        [Output("datastore")]
        public Output<string> Datastore { get; private set; } = null!;

        /// <summary>
        /// The path to where the file should be uploaded
        /// or copied to on the destination `datastore` in vSphere.
        /// </summary>
        [Output("destinationFile")]
        public Output<string> DestinationFile { get; private set; } = null!;

        /// <summary>
        /// The name of a datacenter from which the file
        /// will be copied. Forces a new resource if changed.
        /// </summary>
        [Output("sourceDatacenter")]
        public Output<string?> SourceDatacenter { get; private set; } = null!;

        /// <summary>
        /// The name of the datastore from which file will
        /// be copied. Forces a new resource if changed.
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
            : base("vsphere:index/file:File", name, args ?? new FileArgs(), MakeResourceOptions(options, ""))
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

    public sealed class FileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create directories in `destination_file`
        /// path parameter on first apply if any are missing for copy operation.
        /// 
        /// &gt; **NOTE:** Any directory created as part of the `create_directories` argument
        /// will not be deleted when the resource is destroyed. New directories are not
        /// created if the `destination_file` path is changed in subsequent applies.
        /// </summary>
        [Input("createDirectories")]
        public Input<bool>? CreateDirectories { get; set; }

        /// <summary>
        /// The name of a datacenter to which the file will be
        /// uploaded.
        /// </summary>
        [Input("datacenter")]
        public Input<string>? Datacenter { get; set; }

        /// <summary>
        /// The name of the datastore to which to upload the
        /// file.
        /// </summary>
        [Input("datastore", required: true)]
        public Input<string> Datastore { get; set; } = null!;

        /// <summary>
        /// The path to where the file should be uploaded
        /// or copied to on the destination `datastore` in vSphere.
        /// </summary>
        [Input("destinationFile", required: true)]
        public Input<string> DestinationFile { get; set; } = null!;

        /// <summary>
        /// The name of a datacenter from which the file
        /// will be copied. Forces a new resource if changed.
        /// </summary>
        [Input("sourceDatacenter")]
        public Input<string>? SourceDatacenter { get; set; }

        /// <summary>
        /// The name of the datastore from which file will
        /// be copied. Forces a new resource if changed.
        /// </summary>
        [Input("sourceDatastore")]
        public Input<string>? SourceDatastore { get; set; }

        [Input("sourceFile", required: true)]
        public Input<string> SourceFile { get; set; } = null!;

        public FileArgs()
        {
        }
        public static new FileArgs Empty => new FileArgs();
    }

    public sealed class FileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create directories in `destination_file`
        /// path parameter on first apply if any are missing for copy operation.
        /// 
        /// &gt; **NOTE:** Any directory created as part of the `create_directories` argument
        /// will not be deleted when the resource is destroyed. New directories are not
        /// created if the `destination_file` path is changed in subsequent applies.
        /// </summary>
        [Input("createDirectories")]
        public Input<bool>? CreateDirectories { get; set; }

        /// <summary>
        /// The name of a datacenter to which the file will be
        /// uploaded.
        /// </summary>
        [Input("datacenter")]
        public Input<string>? Datacenter { get; set; }

        /// <summary>
        /// The name of the datastore to which to upload the
        /// file.
        /// </summary>
        [Input("datastore")]
        public Input<string>? Datastore { get; set; }

        /// <summary>
        /// The path to where the file should be uploaded
        /// or copied to on the destination `datastore` in vSphere.
        /// </summary>
        [Input("destinationFile")]
        public Input<string>? DestinationFile { get; set; }

        /// <summary>
        /// The name of a datacenter from which the file
        /// will be copied. Forces a new resource if changed.
        /// </summary>
        [Input("sourceDatacenter")]
        public Input<string>? SourceDatacenter { get; set; }

        /// <summary>
        /// The name of the datastore from which file will
        /// be copied. Forces a new resource if changed.
        /// </summary>
        [Input("sourceDatastore")]
        public Input<string>? SourceDatastore { get; set; }

        [Input("sourceFile")]
        public Input<string>? SourceFile { get; set; }

        public FileState()
        {
        }
        public static new FileState Empty => new FileState();
    }
}
