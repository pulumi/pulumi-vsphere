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
    /// Provides a VMware vSphere host resource. This represents an ESXi host that
    /// can be used either as a member of a cluster or as a standalone host.
    /// 
    /// ## Example Usage
    /// ### Create a standalone host
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using VSphere = Pulumi.VSphere;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var datacenter = VSphere.GetDatacenter.Invoke(new()
    ///     {
    ///         Name = "dc-01",
    ///     });
    /// 
    ///     var thumbprint = VSphere.GetHostThumbprint.Invoke(new()
    ///     {
    ///         Address = "esx-01.example.com",
    ///         Insecure = true,
    ///     });
    /// 
    ///     var esx_01 = new VSphere.Host("esx-01", new()
    ///     {
    ///         Hostname = "esx-01.example.com",
    ///         Username = "root",
    ///         Password = "password",
    ///         License = "00000-00000-00000-00000-00000",
    ///         Thumbprint = thumbprint.Apply(getHostThumbprintResult =&gt; getHostThumbprintResult.Id),
    ///         Datacenter = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// ### Create host in a compute cluster
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using VSphere = Pulumi.VSphere;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var datacenter = VSphere.GetDatacenter.Invoke(new()
    ///     {
    ///         Name = "dc-01",
    ///     });
    /// 
    ///     var cluster = VSphere.GetComputeCluster.Invoke(new()
    ///     {
    ///         Name = "cluster-01",
    ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
    ///     });
    /// 
    ///     var thumbprint = VSphere.GetHostThumbprint.Invoke(new()
    ///     {
    ///         Address = "esx-01.example.com",
    ///         Insecure = true,
    ///     });
    /// 
    ///     var esx_01 = new VSphere.Host("esx-01", new()
    ///     {
    ///         Hostname = "esx-01.example.com",
    ///         Username = "root",
    ///         Password = "password",
    ///         License = "00000-00000-00000-00000-00000",
    ///         Thumbprint = thumbprint.Apply(getHostThumbprintResult =&gt; getHostThumbprintResult.Id),
    ///         Cluster = cluster.Apply(getComputeClusterResult =&gt; getComputeClusterResult.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// ## Importing
    /// 
    /// An existing host can be [imported][docs-import] into this resource by supplying
    /// the host's ID. An example is below:
    /// 
    /// [docs-import]: /docs/import/index.html
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    /// });
    /// ```
    /// 
    /// The above would import the host with ID `host-123`.
    /// </summary>
    [VSphereResourceType("vsphere:index/host:Host")]
    public partial class Host : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Compute Cluster this host should
        /// be added to. This should not be set if `datacenter` is set. Conflicts with:
        /// `cluster_managed`.
        /// </summary>
        [Output("cluster")]
        public Output<string?> Cluster { get; private set; } = null!;

        /// <summary>
        /// Can be set to `true` if compute cluster
        /// membership will be managed through the `compute_cluster` resource rather
        /// than the`host` resource. Conflicts with: `cluster`.
        /// </summary>
        [Output("clusterManaged")]
        public Output<bool?> ClusterManaged { get; private set; } = null!;

        /// <summary>
        /// If set to false then the host will be disconnected.
        /// Default is `false`.
        /// </summary>
        [Output("connected")]
        public Output<bool?> Connected { get; private set; } = null!;

        /// <summary>
        /// A map of custom attribute IDs and string
        /// values to apply to the resource. Please refer to the
        /// `vsphere_custom_attributes` resource for more information on applying
        /// tags to resources.
        /// 
        /// &gt; **NOTE:** Custom attributes are not supported on direct ESXi host
        /// connections and require vCenter Server.
        /// 
        /// [docs-host-thumbprint-data-source]: /docs/providers/vsphere/d/host_thumbprint.html
        /// </summary>
        [Output("customAttributes")]
        public Output<ImmutableDictionary<string, string>?> CustomAttributes { get; private set; } = null!;

        /// <summary>
        /// The ID of the datacenter this host should
        /// be added to. This should not be set if `cluster` is set.
        /// </summary>
        [Output("datacenter")]
        public Output<string?> Datacenter { get; private set; } = null!;

        /// <summary>
        /// If set to `true` then it will force the host to be added,
        /// even if the host is already connected to a different vCenter Server instance.
        /// Default is `false`.
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// FQDN or IP address of the host to be added.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// The license key that will be applied to the host.
        /// The license key is expected to be present in vSphere.
        /// </summary>
        [Output("license")]
        public Output<string?> License { get; private set; } = null!;

        /// <summary>
        /// Set the lockdown state of the host. Valid options are
        /// `disabled`, `normal`, and `strict`. Default is `disabled`.
        /// </summary>
        [Output("lockdown")]
        public Output<string?> Lockdown { get; private set; } = null!;

        /// <summary>
        /// Set the management state of the host.
        /// Default is `false`.
        /// </summary>
        [Output("maintenance")]
        public Output<bool?> Maintenance { get; private set; } = null!;

        /// <summary>
        /// Password that will be used by vSphere to authenticate
        /// to the host.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The IDs of any tags to attach to this resource. Please
        /// refer to the `vsphere.Tag` resource for more information on applying
        /// tags to resources.
        /// 
        /// &gt; **NOTE:** Tagging support is not supported on direct ESXi host
        /// connections and require vCenter Server.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Host's certificate SHA-1 thumbprint. If not set the
        /// CA that signed the host's certificate should be trusted. If the CA is not
        /// trusted and no thumbprint is set then the operation will fail. See data source
        /// [`vsphere.getHostThumbprint`][docs-host-thumbprint-data-source].
        /// </summary>
        [Output("thumbprint")]
        public Output<string?> Thumbprint { get; private set; } = null!;

        /// <summary>
        /// Username that will be used by vSphere to authenticate
        /// to the host.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Host resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Host(string name, HostArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/host:Host", name, args ?? new HostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Host(string name, Input<string> id, HostState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/host:Host", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Host resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Host Get(string name, Input<string> id, HostState? state = null, CustomResourceOptions? options = null)
        {
            return new Host(name, id, state, options);
        }
    }

    public sealed class HostArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Compute Cluster this host should
        /// be added to. This should not be set if `datacenter` is set. Conflicts with:
        /// `cluster_managed`.
        /// </summary>
        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        /// <summary>
        /// Can be set to `true` if compute cluster
        /// membership will be managed through the `compute_cluster` resource rather
        /// than the`host` resource. Conflicts with: `cluster`.
        /// </summary>
        [Input("clusterManaged")]
        public Input<bool>? ClusterManaged { get; set; }

        /// <summary>
        /// If set to false then the host will be disconnected.
        /// Default is `false`.
        /// </summary>
        [Input("connected")]
        public Input<bool>? Connected { get; set; }

        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// A map of custom attribute IDs and string
        /// values to apply to the resource. Please refer to the
        /// `vsphere_custom_attributes` resource for more information on applying
        /// tags to resources.
        /// 
        /// &gt; **NOTE:** Custom attributes are not supported on direct ESXi host
        /// connections and require vCenter Server.
        /// 
        /// [docs-host-thumbprint-data-source]: /docs/providers/vsphere/d/host_thumbprint.html
        /// </summary>
        public InputMap<string> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<string>());
            set => _customAttributes = value;
        }

        /// <summary>
        /// The ID of the datacenter this host should
        /// be added to. This should not be set if `cluster` is set.
        /// </summary>
        [Input("datacenter")]
        public Input<string>? Datacenter { get; set; }

        /// <summary>
        /// If set to `true` then it will force the host to be added,
        /// even if the host is already connected to a different vCenter Server instance.
        /// Default is `false`.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// FQDN or IP address of the host to be added.
        /// </summary>
        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        /// <summary>
        /// The license key that will be applied to the host.
        /// The license key is expected to be present in vSphere.
        /// </summary>
        [Input("license")]
        public Input<string>? License { get; set; }

        /// <summary>
        /// Set the lockdown state of the host. Valid options are
        /// `disabled`, `normal`, and `strict`. Default is `disabled`.
        /// </summary>
        [Input("lockdown")]
        public Input<string>? Lockdown { get; set; }

        /// <summary>
        /// Set the management state of the host.
        /// Default is `false`.
        /// </summary>
        [Input("maintenance")]
        public Input<bool>? Maintenance { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Password that will be used by vSphere to authenticate
        /// to the host.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The IDs of any tags to attach to this resource. Please
        /// refer to the `vsphere.Tag` resource for more information on applying
        /// tags to resources.
        /// 
        /// &gt; **NOTE:** Tagging support is not supported on direct ESXi host
        /// connections and require vCenter Server.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Host's certificate SHA-1 thumbprint. If not set the
        /// CA that signed the host's certificate should be trusted. If the CA is not
        /// trusted and no thumbprint is set then the operation will fail. See data source
        /// [`vsphere.getHostThumbprint`][docs-host-thumbprint-data-source].
        /// </summary>
        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        /// <summary>
        /// Username that will be used by vSphere to authenticate
        /// to the host.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public HostArgs()
        {
        }
        public static new HostArgs Empty => new HostArgs();
    }

    public sealed class HostState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Compute Cluster this host should
        /// be added to. This should not be set if `datacenter` is set. Conflicts with:
        /// `cluster_managed`.
        /// </summary>
        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        /// <summary>
        /// Can be set to `true` if compute cluster
        /// membership will be managed through the `compute_cluster` resource rather
        /// than the`host` resource. Conflicts with: `cluster`.
        /// </summary>
        [Input("clusterManaged")]
        public Input<bool>? ClusterManaged { get; set; }

        /// <summary>
        /// If set to false then the host will be disconnected.
        /// Default is `false`.
        /// </summary>
        [Input("connected")]
        public Input<bool>? Connected { get; set; }

        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// A map of custom attribute IDs and string
        /// values to apply to the resource. Please refer to the
        /// `vsphere_custom_attributes` resource for more information on applying
        /// tags to resources.
        /// 
        /// &gt; **NOTE:** Custom attributes are not supported on direct ESXi host
        /// connections and require vCenter Server.
        /// 
        /// [docs-host-thumbprint-data-source]: /docs/providers/vsphere/d/host_thumbprint.html
        /// </summary>
        public InputMap<string> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<string>());
            set => _customAttributes = value;
        }

        /// <summary>
        /// The ID of the datacenter this host should
        /// be added to. This should not be set if `cluster` is set.
        /// </summary>
        [Input("datacenter")]
        public Input<string>? Datacenter { get; set; }

        /// <summary>
        /// If set to `true` then it will force the host to be added,
        /// even if the host is already connected to a different vCenter Server instance.
        /// Default is `false`.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// FQDN or IP address of the host to be added.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The license key that will be applied to the host.
        /// The license key is expected to be present in vSphere.
        /// </summary>
        [Input("license")]
        public Input<string>? License { get; set; }

        /// <summary>
        /// Set the lockdown state of the host. Valid options are
        /// `disabled`, `normal`, and `strict`. Default is `disabled`.
        /// </summary>
        [Input("lockdown")]
        public Input<string>? Lockdown { get; set; }

        /// <summary>
        /// Set the management state of the host.
        /// Default is `false`.
        /// </summary>
        [Input("maintenance")]
        public Input<bool>? Maintenance { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password that will be used by vSphere to authenticate
        /// to the host.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The IDs of any tags to attach to this resource. Please
        /// refer to the `vsphere.Tag` resource for more information on applying
        /// tags to resources.
        /// 
        /// &gt; **NOTE:** Tagging support is not supported on direct ESXi host
        /// connections and require vCenter Server.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Host's certificate SHA-1 thumbprint. If not set the
        /// CA that signed the host's certificate should be trusted. If the CA is not
        /// trusted and no thumbprint is set then the operation will fail. See data source
        /// [`vsphere.getHostThumbprint`][docs-host-thumbprint-data-source].
        /// </summary>
        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        /// <summary>
        /// Username that will be used by vSphere to authenticate
        /// to the host.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public HostState()
        {
        }
        public static new HostState Empty => new HostState();
    }
}
