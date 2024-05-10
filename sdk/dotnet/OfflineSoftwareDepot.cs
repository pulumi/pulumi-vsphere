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
    /// Provides a VMware vSphere offline software depot resource.
    /// 
    /// ## Example Usage
    /// </summary>
    [VSphereResourceType("vsphere:index/offlineSoftwareDepot:OfflineSoftwareDepot")]
    public partial class OfflineSoftwareDepot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The list of custom components in the depot.
        /// </summary>
        [Output("components")]
        public Output<ImmutableArray<Outputs.OfflineSoftwareDepotComponent>> Components { get; private set; } = null!;

        /// <summary>
        /// The URL where the depot source is hosted.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;


        /// <summary>
        /// Create a OfflineSoftwareDepot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OfflineSoftwareDepot(string name, OfflineSoftwareDepotArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/offlineSoftwareDepot:OfflineSoftwareDepot", name, args ?? new OfflineSoftwareDepotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OfflineSoftwareDepot(string name, Input<string> id, OfflineSoftwareDepotState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/offlineSoftwareDepot:OfflineSoftwareDepot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OfflineSoftwareDepot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OfflineSoftwareDepot Get(string name, Input<string> id, OfflineSoftwareDepotState? state = null, CustomResourceOptions? options = null)
        {
            return new OfflineSoftwareDepot(name, id, state, options);
        }
    }

    public sealed class OfflineSoftwareDepotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL where the depot source is hosted.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        public OfflineSoftwareDepotArgs()
        {
        }
        public static new OfflineSoftwareDepotArgs Empty => new OfflineSoftwareDepotArgs();
    }

    public sealed class OfflineSoftwareDepotState : global::Pulumi.ResourceArgs
    {
        [Input("components")]
        private InputList<Inputs.OfflineSoftwareDepotComponentGetArgs>? _components;

        /// <summary>
        /// The list of custom components in the depot.
        /// </summary>
        public InputList<Inputs.OfflineSoftwareDepotComponentGetArgs> Components
        {
            get => _components ?? (_components = new InputList<Inputs.OfflineSoftwareDepotComponentGetArgs>());
            set => _components = value;
        }

        /// <summary>
        /// The URL where the depot source is hosted.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        public OfflineSoftwareDepotState()
        {
        }
        public static new OfflineSoftwareDepotState Empty => new OfflineSoftwareDepotState();
    }
}
