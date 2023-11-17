// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    [VSphereResourceType("vsphere:index/folder:Folder")]
    public partial class Folder : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Map of custom attribute ids to attribute 
        /// value strings to set for folder. See [here][docs-setting-custom-attributes] 
        /// for a reference on how to set values for custom attributes.
        /// 
        /// 
        /// 
        /// &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections 
        /// and require vCenter.
        /// </summary>
        [Output("customAttributes")]
        public Output<ImmutableDictionary<string, string>?> CustomAttributes { get; private set; } = null!;

        /// <summary>
        /// The ID of the datacenter the folder will be created in.
        /// Required for all folder types except for datacenter folders. Forces a new
        /// resource if changed.
        /// </summary>
        [Output("datacenterId")]
        public Output<string?> DatacenterId { get; private set; } = null!;

        /// <summary>
        /// The path of the folder to be created. This is relative to
        /// the root of the type of folder you are creating, and the supplied datacenter.
        /// For example, given a default datacenter of `default-dc`, a folder of type
        /// `vm` (denoting a virtual machine folder), and a supplied folder of
        /// `test-folder`, the resulting path would be
        /// `/default-dc/vm/test-folder`.
        /// 
        /// &gt; **NOTE:** `path` can be modified - the resulting behavior is dependent on
        /// what section of `path` you are modifying. If you are modifying the parent (so
        /// any part before the last `/`), your folder will be moved to that new parent. If
        /// modifying the name (the part after the last `/`), your folder will be renamed.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// 
        /// &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
        /// requires vCenter 6.0 or higher.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of folder to create. Allowed options are
        /// `datacenter` for datacenter folders, `host` for host and cluster folders,
        /// `vm` for virtual machine folders, `datastore` for datastore folders, and
        /// `network` for network folders. Forces a new resource if changed.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Folder resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Folder(string name, FolderArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/folder:Folder", name, args ?? new FolderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Folder(string name, Input<string> id, FolderState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/folder:Folder", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Folder resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Folder Get(string name, Input<string> id, FolderState? state = null, CustomResourceOptions? options = null)
        {
            return new Folder(name, id, state, options);
        }
    }

    public sealed class FolderArgs : global::Pulumi.ResourceArgs
    {
        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// Map of custom attribute ids to attribute 
        /// value strings to set for folder. See [here][docs-setting-custom-attributes] 
        /// for a reference on how to set values for custom attributes.
        /// 
        /// 
        /// 
        /// &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections 
        /// and require vCenter.
        /// </summary>
        public InputMap<string> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<string>());
            set => _customAttributes = value;
        }

        /// <summary>
        /// The ID of the datacenter the folder will be created in.
        /// Required for all folder types except for datacenter folders. Forces a new
        /// resource if changed.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// The path of the folder to be created. This is relative to
        /// the root of the type of folder you are creating, and the supplied datacenter.
        /// For example, given a default datacenter of `default-dc`, a folder of type
        /// `vm` (denoting a virtual machine folder), and a supplied folder of
        /// `test-folder`, the resulting path would be
        /// `/default-dc/vm/test-folder`.
        /// 
        /// &gt; **NOTE:** `path` can be modified - the resulting behavior is dependent on
        /// what section of `path` you are modifying. If you are modifying the parent (so
        /// any part before the last `/`), your folder will be moved to that new parent. If
        /// modifying the name (the part after the last `/`), your folder will be renamed.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// 
        /// &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
        /// requires vCenter 6.0 or higher.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of folder to create. Allowed options are
        /// `datacenter` for datacenter folders, `host` for host and cluster folders,
        /// `vm` for virtual machine folders, `datastore` for datastore folders, and
        /// `network` for network folders. Forces a new resource if changed.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FolderArgs()
        {
        }
        public static new FolderArgs Empty => new FolderArgs();
    }

    public sealed class FolderState : global::Pulumi.ResourceArgs
    {
        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// Map of custom attribute ids to attribute 
        /// value strings to set for folder. See [here][docs-setting-custom-attributes] 
        /// for a reference on how to set values for custom attributes.
        /// 
        /// 
        /// 
        /// &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections 
        /// and require vCenter.
        /// </summary>
        public InputMap<string> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<string>());
            set => _customAttributes = value;
        }

        /// <summary>
        /// The ID of the datacenter the folder will be created in.
        /// Required for all folder types except for datacenter folders. Forces a new
        /// resource if changed.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// The path of the folder to be created. This is relative to
        /// the root of the type of folder you are creating, and the supplied datacenter.
        /// For example, given a default datacenter of `default-dc`, a folder of type
        /// `vm` (denoting a virtual machine folder), and a supplied folder of
        /// `test-folder`, the resulting path would be
        /// `/default-dc/vm/test-folder`.
        /// 
        /// &gt; **NOTE:** `path` can be modified - the resulting behavior is dependent on
        /// what section of `path` you are modifying. If you are modifying the parent (so
        /// any part before the last `/`), your folder will be moved to that new parent. If
        /// modifying the name (the part after the last `/`), your folder will be renamed.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// 
        /// &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
        /// requires vCenter 6.0 or higher.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of folder to create. Allowed options are
        /// `datacenter` for datacenter folders, `host` for host and cluster folders,
        /// `vm` for virtual machine folders, `datastore` for datastore folders, and
        /// `network` for network folders. Forces a new resource if changed.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FolderState()
        {
        }
        public static new FolderState Empty => new FolderState();
    }
}
