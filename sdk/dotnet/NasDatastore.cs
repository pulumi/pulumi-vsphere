// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    [VSphereResourceType("vsphere:index/nasDatastore:NasDatastore")]
    public partial class NasDatastore : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access mode for the mount point. Can be one of
        /// `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
        /// that the datastore will be read-write depending on the permissions of the
        /// actual share. Default: `readWrite`. Forces a new resource if changed.
        /// </summary>
        [Output("accessMode")]
        public Output<string?> AccessMode { get; private set; } = null!;

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
        /// value strings to set on datasource resource.
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
        /// The managed object IDs of
        /// the hosts to mount the datastore on.
        /// </summary>
        [Output("hostSystemIds")]
        public Output<ImmutableArray<string>> HostSystemIds { get; private set; } = null!;

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
        /// Indicates that this NAS volume is a protocol endpoint.
        /// This field is only populated if the host supports virtual datastores.
        /// </summary>
        [Output("protocolEndpoint")]
        public Output<bool> ProtocolEndpoint { get; private set; } = null!;

        /// <summary>
        /// The hostnames or IP addresses of the remote
        /// server or servers. Only one element should be present for NFS v3 but multiple
        /// can be present for NFS v4.1. Forces a new resource if changed.
        /// </summary>
        [Output("remoteHosts")]
        public Output<ImmutableArray<string>> RemoteHosts { get; private set; } = null!;

        /// <summary>
        /// The remote path of the mount point. Forces a new
        /// resource if changed.
        /// </summary>
        [Output("remotePath")]
        public Output<string> RemotePath { get; private set; } = null!;

        /// <summary>
        /// The security type to use when using NFS v4.1.
        /// Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
        /// if changed.
        /// </summary>
        [Output("securityType")]
        public Output<string?> SecurityType { get; private set; } = null!;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of NAS volume. Can be one of `NFS` (to denote
        /// v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
        /// changed.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

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
        /// Create a NasDatastore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NasDatastore(string name, NasDatastoreArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/nasDatastore:NasDatastore", name, args ?? new NasDatastoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NasDatastore(string name, Input<string> id, NasDatastoreState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/nasDatastore:NasDatastore", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NasDatastore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NasDatastore Get(string name, Input<string> id, NasDatastoreState? state = null, CustomResourceOptions? options = null)
        {
            return new NasDatastore(name, id, state, options);
        }
    }

    public sealed class NasDatastoreArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access mode for the mount point. Can be one of
        /// `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
        /// that the datastore will be read-write depending on the permissions of the
        /// actual share. Default: `readWrite`. Forces a new resource if changed.
        /// </summary>
        [Input("accessMode")]
        public Input<string>? AccessMode { get; set; }

        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// Map of custom attribute ids to attribute 
        /// value strings to set on datasource resource.
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

        [Input("hostSystemIds", required: true)]
        private InputList<string>? _hostSystemIds;

        /// <summary>
        /// The managed object IDs of
        /// the hosts to mount the datastore on.
        /// </summary>
        public InputList<string> HostSystemIds
        {
            get => _hostSystemIds ?? (_hostSystemIds = new InputList<string>());
            set => _hostSystemIds = value;
        }

        /// <summary>
        /// The name of the datastore. Forces a new resource if
        /// changed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("remoteHosts", required: true)]
        private InputList<string>? _remoteHosts;

        /// <summary>
        /// The hostnames or IP addresses of the remote
        /// server or servers. Only one element should be present for NFS v3 but multiple
        /// can be present for NFS v4.1. Forces a new resource if changed.
        /// </summary>
        public InputList<string> RemoteHosts
        {
            get => _remoteHosts ?? (_remoteHosts = new InputList<string>());
            set => _remoteHosts = value;
        }

        /// <summary>
        /// The remote path of the mount point. Forces a new
        /// resource if changed.
        /// </summary>
        [Input("remotePath", required: true)]
        public Input<string> RemotePath { get; set; } = null!;

        /// <summary>
        /// The security type to use when using NFS v4.1.
        /// Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
        /// if changed.
        /// </summary>
        [Input("securityType")]
        public Input<string>? SecurityType { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of NAS volume. Can be one of `NFS` (to denote
        /// v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
        /// changed.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public NasDatastoreArgs()
        {
        }
        public static new NasDatastoreArgs Empty => new NasDatastoreArgs();
    }

    public sealed class NasDatastoreState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access mode for the mount point. Can be one of
        /// `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
        /// that the datastore will be read-write depending on the permissions of the
        /// actual share. Default: `readWrite`. Forces a new resource if changed.
        /// </summary>
        [Input("accessMode")]
        public Input<string>? AccessMode { get; set; }

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
        /// value strings to set on datasource resource.
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

        [Input("hostSystemIds")]
        private InputList<string>? _hostSystemIds;

        /// <summary>
        /// The managed object IDs of
        /// the hosts to mount the datastore on.
        /// </summary>
        public InputList<string> HostSystemIds
        {
            get => _hostSystemIds ?? (_hostSystemIds = new InputList<string>());
            set => _hostSystemIds = value;
        }

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

        /// <summary>
        /// Indicates that this NAS volume is a protocol endpoint.
        /// This field is only populated if the host supports virtual datastores.
        /// </summary>
        [Input("protocolEndpoint")]
        public Input<bool>? ProtocolEndpoint { get; set; }

        [Input("remoteHosts")]
        private InputList<string>? _remoteHosts;

        /// <summary>
        /// The hostnames or IP addresses of the remote
        /// server or servers. Only one element should be present for NFS v3 but multiple
        /// can be present for NFS v4.1. Forces a new resource if changed.
        /// </summary>
        public InputList<string> RemoteHosts
        {
            get => _remoteHosts ?? (_remoteHosts = new InputList<string>());
            set => _remoteHosts = value;
        }

        /// <summary>
        /// The remote path of the mount point. Forces a new
        /// resource if changed.
        /// </summary>
        [Input("remotePath")]
        public Input<string>? RemotePath { get; set; }

        /// <summary>
        /// The security type to use when using NFS v4.1.
        /// Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
        /// if changed.
        /// </summary>
        [Input("securityType")]
        public Input<string>? SecurityType { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of NAS volume. Can be one of `NFS` (to denote
        /// v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
        /// changed.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

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

        public NasDatastoreState()
        {
        }
        public static new NasDatastoreState Empty => new NasDatastoreState();
    }
}
