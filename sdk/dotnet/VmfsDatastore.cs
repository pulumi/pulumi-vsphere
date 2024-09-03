// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    [VSphereResourceType("vsphere:index/vmfsDatastore:VmfsDatastore")]
    public partial class VmfsDatastore : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The connectivity status of the datastore. If this is `false`,
        /// some other computed attributes may be out of date.
        /// </summary>
        [Output("accessible")]
        public Output<bool> Accessible { get; private set; } = null!;

        /// <summary>
        /// Maximum capacity of the datastore, in megabytes.
        /// </summary>
        [Output("capacity")]
        public Output<int> Capacity { get; private set; } = null!;

        /// <summary>
        /// Map of custom attribute ids to attribute
        /// value string to set on datastore resource.
        /// 
        /// &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
        /// and require vCenter.
        /// </summary>
        [Output("customAttributes")]
        public Output<ImmutableDictionary<string, string>?> CustomAttributes { get; private set; } = null!;

        /// <summary>
        /// The managed object
        /// ID of a datastore cluster to put this datastore in.
        /// Conflicts with `folder`.
        /// </summary>
        [Output("datastoreClusterId")]
        public Output<string?> DatastoreClusterId { get; private set; } = null!;

        /// <summary>
        /// The disks to use with the datastore.
        /// </summary>
        [Output("disks")]
        public Output<ImmutableArray<string>> Disks { get; private set; } = null!;

        /// <summary>
        /// The relative path to a folder to put this datastore in.
        /// This is a path relative to the datacenter you are deploying the datastore to.
        /// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
        /// The provider will place a datastore named `test` in a datastore folder
        /// located at `/dc1/datastore/foo/bar`, with the final inventory path being
        /// `/dc1/datastore/foo/bar/test`. Conflicts with
        /// `datastore_cluster_id`.
        /// </summary>
        [Output("folder")]
        public Output<string?> Folder { get; private set; } = null!;

        /// <summary>
        /// Available space of this datastore, in megabytes.
        /// </summary>
        [Output("freeSpace")]
        public Output<int> FreeSpace { get; private set; } = null!;

        /// <summary>
        /// The managed object ID of
        /// the host to set the datastore up on. Note that this is not necessarily the
        /// only host that the datastore will be set up on - see
        /// here for more info. Forces a
        /// new resource if changed.
        /// </summary>
        [Output("hostSystemId")]
        public Output<string> HostSystemId { get; private set; } = null!;

        /// <summary>
        /// The current maintenance mode state of the datastore.
        /// </summary>
        [Output("maintenanceMode")]
        public Output<string> MaintenanceMode { get; private set; } = null!;

        /// <summary>
        /// If `true`, more than one host in the datacenter has
        /// been configured with access to the datastore.
        /// </summary>
        [Output("multipleHostAccess")]
        public Output<bool> MultipleHostAccess { get; private set; } = null!;

        /// <summary>
        /// The name of the datastore. Forces a new resource if
        /// changed.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// 
        /// &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
        /// requires vCenter 6.0 or higher.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Total additional storage space, in megabytes,
        /// potentially used by all virtual machines on this datastore.
        /// </summary>
        [Output("uncommittedSpace")]
        public Output<int> UncommittedSpace { get; private set; } = null!;

        /// <summary>
        /// The unique locator for the datastore.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a VmfsDatastore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VmfsDatastore(string name, VmfsDatastoreArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/vmfsDatastore:VmfsDatastore", name, args ?? new VmfsDatastoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VmfsDatastore(string name, Input<string> id, VmfsDatastoreState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/vmfsDatastore:VmfsDatastore", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VmfsDatastore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VmfsDatastore Get(string name, Input<string> id, VmfsDatastoreState? state = null, CustomResourceOptions? options = null)
        {
            return new VmfsDatastore(name, id, state, options);
        }
    }

    public sealed class VmfsDatastoreArgs : global::Pulumi.ResourceArgs
    {
        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// Map of custom attribute ids to attribute
        /// value string to set on datastore resource.
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
        /// The managed object
        /// ID of a datastore cluster to put this datastore in.
        /// Conflicts with `folder`.
        /// </summary>
        [Input("datastoreClusterId")]
        public Input<string>? DatastoreClusterId { get; set; }

        [Input("disks", required: true)]
        private InputList<string>? _disks;

        /// <summary>
        /// The disks to use with the datastore.
        /// </summary>
        public InputList<string> Disks
        {
            get => _disks ?? (_disks = new InputList<string>());
            set => _disks = value;
        }

        /// <summary>
        /// The relative path to a folder to put this datastore in.
        /// This is a path relative to the datacenter you are deploying the datastore to.
        /// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
        /// The provider will place a datastore named `test` in a datastore folder
        /// located at `/dc1/datastore/foo/bar`, with the final inventory path being
        /// `/dc1/datastore/foo/bar/test`. Conflicts with
        /// `datastore_cluster_id`.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// The managed object ID of
        /// the host to set the datastore up on. Note that this is not necessarily the
        /// only host that the datastore will be set up on - see
        /// here for more info. Forces a
        /// new resource if changed.
        /// </summary>
        [Input("hostSystemId", required: true)]
        public Input<string> HostSystemId { get; set; } = null!;

        /// <summary>
        /// The name of the datastore. Forces a new resource if
        /// changed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public VmfsDatastoreArgs()
        {
        }
        public static new VmfsDatastoreArgs Empty => new VmfsDatastoreArgs();
    }

    public sealed class VmfsDatastoreState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The connectivity status of the datastore. If this is `false`,
        /// some other computed attributes may be out of date.
        /// </summary>
        [Input("accessible")]
        public Input<bool>? Accessible { get; set; }

        /// <summary>
        /// Maximum capacity of the datastore, in megabytes.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// Map of custom attribute ids to attribute
        /// value string to set on datastore resource.
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
        /// The managed object
        /// ID of a datastore cluster to put this datastore in.
        /// Conflicts with `folder`.
        /// </summary>
        [Input("datastoreClusterId")]
        public Input<string>? DatastoreClusterId { get; set; }

        [Input("disks")]
        private InputList<string>? _disks;

        /// <summary>
        /// The disks to use with the datastore.
        /// </summary>
        public InputList<string> Disks
        {
            get => _disks ?? (_disks = new InputList<string>());
            set => _disks = value;
        }

        /// <summary>
        /// The relative path to a folder to put this datastore in.
        /// This is a path relative to the datacenter you are deploying the datastore to.
        /// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
        /// The provider will place a datastore named `test` in a datastore folder
        /// located at `/dc1/datastore/foo/bar`, with the final inventory path being
        /// `/dc1/datastore/foo/bar/test`. Conflicts with
        /// `datastore_cluster_id`.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// Available space of this datastore, in megabytes.
        /// </summary>
        [Input("freeSpace")]
        public Input<int>? FreeSpace { get; set; }

        /// <summary>
        /// The managed object ID of
        /// the host to set the datastore up on. Note that this is not necessarily the
        /// only host that the datastore will be set up on - see
        /// here for more info. Forces a
        /// new resource if changed.
        /// </summary>
        [Input("hostSystemId")]
        public Input<string>? HostSystemId { get; set; }

        /// <summary>
        /// The current maintenance mode state of the datastore.
        /// </summary>
        [Input("maintenanceMode")]
        public Input<string>? MaintenanceMode { get; set; }

        /// <summary>
        /// If `true`, more than one host in the datacenter has
        /// been configured with access to the datastore.
        /// </summary>
        [Input("multipleHostAccess")]
        public Input<bool>? MultipleHostAccess { get; set; }

        /// <summary>
        /// The name of the datastore. Forces a new resource if
        /// changed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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
        /// Total additional storage space, in megabytes,
        /// potentially used by all virtual machines on this datastore.
        /// </summary>
        [Input("uncommittedSpace")]
        public Input<int>? UncommittedSpace { get; set; }

        /// <summary>
        /// The unique locator for the datastore.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public VmfsDatastoreState()
        {
        }
        public static new VmfsDatastoreState Empty => new VmfsDatastoreState();
    }
}
