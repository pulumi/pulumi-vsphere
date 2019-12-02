// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vsphere
{
    /// <summary>
    /// Provides a VMware vSphere license resource. This can be used to add and remove license keys.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/license.html.markdown.
    /// </summary>
    public partial class License : Pulumi.CustomResource
    {
        /// <summary>
        /// The product edition of the license key.
        /// </summary>
        [Output("editionKey")]
        public Output<string> EditionKey { get; private set; } = null!;

        /// <summary>
        /// A map of key/value pairs to be attached as labels (tags) to the license key.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, object>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The license key to add.
        /// </summary>
        [Output("licenseKey")]
        public Output<string> LicenseKey { get; private set; } = null!;

        /// <summary>
        /// The display name for the license.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Total number of units (example: CPUs) contained in the license.
        /// </summary>
        [Output("total")]
        public Output<int> Total { get; private set; } = null!;

        /// <summary>
        /// The number of units (example: CPUs) assigned to this license.
        /// </summary>
        [Output("used")]
        public Output<int> Used { get; private set; } = null!;


        /// <summary>
        /// Create a License resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public License(string name, LicenseArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/license:License", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private License(string name, Input<string> id, LicenseState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/license:License", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing License resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static License Get(string name, Input<string> id, LicenseState? state = null, CustomResourceOptions? options = null)
        {
            return new License(name, id, state, options);
        }
    }

    public sealed class LicenseArgs : Pulumi.ResourceArgs
    {
        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// A map of key/value pairs to be attached as labels (tags) to the license key.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// The license key to add.
        /// </summary>
        [Input("licenseKey", required: true)]
        public Input<string> LicenseKey { get; set; } = null!;

        public LicenseArgs()
        {
        }
    }

    public sealed class LicenseState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The product edition of the license key.
        /// </summary>
        [Input("editionKey")]
        public Input<string>? EditionKey { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// A map of key/value pairs to be attached as labels (tags) to the license key.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// The license key to add.
        /// </summary>
        [Input("licenseKey")]
        public Input<string>? LicenseKey { get; set; }

        /// <summary>
        /// The display name for the license.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Total number of units (example: CPUs) contained in the license.
        /// </summary>
        [Input("total")]
        public Input<int>? Total { get; set; }

        /// <summary>
        /// The number of units (example: CPUs) assigned to this license.
        /// </summary>
        [Input("used")]
        public Input<int>? Used { get; set; }

        public LicenseState()
        {
        }
    }
}
