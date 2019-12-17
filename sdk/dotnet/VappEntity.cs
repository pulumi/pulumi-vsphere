// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    /// <summary>
    /// The `vsphere..VappEntity` resource can be used to describe the behavior of an
    /// entity (virtual machine or sub-vApp container) in a vApp container.
    /// 
    /// For more information on vSphere vApps, see [this
    /// page][ref-vsphere-vapp].
    /// 
    /// [ref-vsphere-vapp]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vm_admin.doc/GUID-2A95EBB8-1779-40FA-B4FB-4D0845750879.html
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/vapp_entity.html.markdown.
    /// </summary>
    public partial class VappEntity : Pulumi.CustomResource
    {
        /// <summary>
        /// [Managed object ID|docs-about-morefs] of the vApp
        /// container the entity is a member of.
        /// </summary>
        [Output("containerId")]
        public Output<string> ContainerId { get; private set; } = null!;

        /// <summary>
        /// A list of custom attributes to set on this resource.
        /// </summary>
        [Output("customAttributes")]
        public Output<ImmutableDictionary<string, object>?> CustomAttributes { get; private set; } = null!;

        /// <summary>
        /// How to start the entity. Valid settings are none
        /// or powerOn. If set to none, then the entity does not participate in auto-start.
        /// Default: powerOn
        /// </summary>
        [Output("startAction")]
        public Output<string?> StartAction { get; private set; } = null!;

        /// <summary>
        /// Delay in seconds before continuing with the next
        /// entity in the order of entities to be started. Default: 120
        /// </summary>
        [Output("startDelay")]
        public Output<int?> StartDelay { get; private set; } = null!;

        /// <summary>
        /// Order to start and stop target in vApp. Default: 1
        /// </summary>
        [Output("startOrder")]
        public Output<int?> StartOrder { get; private set; } = null!;

        /// <summary>
        /// Defines the stop action for the entity. Can be set
        /// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
        /// does not participate in auto-stop. Default: powerOff
        /// </summary>
        [Output("stopAction")]
        public Output<string?> StopAction { get; private set; } = null!;

        /// <summary>
        /// Delay in seconds before continuing with the next
        /// entity in the order sequence. This is only used if the stopAction is
        /// guestShutdown. Default: 120
        /// </summary>
        [Output("stopDelay")]
        public Output<int?> StopDelay { get; private set; } = null!;

        /// <summary>
        /// A list of tag IDs to apply to this object.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// [Managed object ID|docs-about-morefs] of the entity
        /// to power on or power off. This can be a virtual machine or a vApp.
        /// </summary>
        [Output("targetId")]
        public Output<string> TargetId { get; private set; } = null!;

        /// <summary>
        /// Determines if the VM should be marked as being
        /// started when VMware Tools are ready instead of waiting for `start_delay`. This
        /// property has no effect for vApps. Default: false
        /// </summary>
        [Output("waitForGuest")]
        public Output<bool?> WaitForGuest { get; private set; } = null!;


        /// <summary>
        /// Create a VappEntity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VappEntity(string name, VappEntityArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/vappEntity:VappEntity", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private VappEntity(string name, Input<string> id, VappEntityState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/vappEntity:VappEntity", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VappEntity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VappEntity Get(string name, Input<string> id, VappEntityState? state = null, CustomResourceOptions? options = null)
        {
            return new VappEntity(name, id, state, options);
        }
    }

    public sealed class VappEntityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Managed object ID|docs-about-morefs] of the vApp
        /// container the entity is a member of.
        /// </summary>
        [Input("containerId", required: true)]
        public Input<string> ContainerId { get; set; } = null!;

        [Input("customAttributes")]
        private InputMap<object>? _customAttributes;

        /// <summary>
        /// A list of custom attributes to set on this resource.
        /// </summary>
        public InputMap<object> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<object>());
            set => _customAttributes = value;
        }

        /// <summary>
        /// How to start the entity. Valid settings are none
        /// or powerOn. If set to none, then the entity does not participate in auto-start.
        /// Default: powerOn
        /// </summary>
        [Input("startAction")]
        public Input<string>? StartAction { get; set; }

        /// <summary>
        /// Delay in seconds before continuing with the next
        /// entity in the order of entities to be started. Default: 120
        /// </summary>
        [Input("startDelay")]
        public Input<int>? StartDelay { get; set; }

        /// <summary>
        /// Order to start and stop target in vApp. Default: 1
        /// </summary>
        [Input("startOrder")]
        public Input<int>? StartOrder { get; set; }

        /// <summary>
        /// Defines the stop action for the entity. Can be set
        /// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
        /// does not participate in auto-stop. Default: powerOff
        /// </summary>
        [Input("stopAction")]
        public Input<string>? StopAction { get; set; }

        /// <summary>
        /// Delay in seconds before continuing with the next
        /// entity in the order sequence. This is only used if the stopAction is
        /// guestShutdown. Default: 120
        /// </summary>
        [Input("stopDelay")]
        public Input<int>? StopDelay { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tag IDs to apply to this object.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// [Managed object ID|docs-about-morefs] of the entity
        /// to power on or power off. This can be a virtual machine or a vApp.
        /// </summary>
        [Input("targetId", required: true)]
        public Input<string> TargetId { get; set; } = null!;

        /// <summary>
        /// Determines if the VM should be marked as being
        /// started when VMware Tools are ready instead of waiting for `start_delay`. This
        /// property has no effect for vApps. Default: false
        /// </summary>
        [Input("waitForGuest")]
        public Input<bool>? WaitForGuest { get; set; }

        public VappEntityArgs()
        {
        }
    }

    public sealed class VappEntityState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Managed object ID|docs-about-morefs] of the vApp
        /// container the entity is a member of.
        /// </summary>
        [Input("containerId")]
        public Input<string>? ContainerId { get; set; }

        [Input("customAttributes")]
        private InputMap<object>? _customAttributes;

        /// <summary>
        /// A list of custom attributes to set on this resource.
        /// </summary>
        public InputMap<object> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<object>());
            set => _customAttributes = value;
        }

        /// <summary>
        /// How to start the entity. Valid settings are none
        /// or powerOn. If set to none, then the entity does not participate in auto-start.
        /// Default: powerOn
        /// </summary>
        [Input("startAction")]
        public Input<string>? StartAction { get; set; }

        /// <summary>
        /// Delay in seconds before continuing with the next
        /// entity in the order of entities to be started. Default: 120
        /// </summary>
        [Input("startDelay")]
        public Input<int>? StartDelay { get; set; }

        /// <summary>
        /// Order to start and stop target in vApp. Default: 1
        /// </summary>
        [Input("startOrder")]
        public Input<int>? StartOrder { get; set; }

        /// <summary>
        /// Defines the stop action for the entity. Can be set
        /// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
        /// does not participate in auto-stop. Default: powerOff
        /// </summary>
        [Input("stopAction")]
        public Input<string>? StopAction { get; set; }

        /// <summary>
        /// Delay in seconds before continuing with the next
        /// entity in the order sequence. This is only used if the stopAction is
        /// guestShutdown. Default: 120
        /// </summary>
        [Input("stopDelay")]
        public Input<int>? StopDelay { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tag IDs to apply to this object.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// [Managed object ID|docs-about-morefs] of the entity
        /// to power on or power off. This can be a virtual machine or a vApp.
        /// </summary>
        [Input("targetId")]
        public Input<string>? TargetId { get; set; }

        /// <summary>
        /// Determines if the VM should be marked as being
        /// started when VMware Tools are ready instead of waiting for `start_delay`. This
        /// property has no effect for vApps. Default: false
        /// </summary>
        [Input("waitForGuest")]
        public Input<bool>? WaitForGuest { get; set; }

        public VappEntityState()
        {
        }
    }
}
