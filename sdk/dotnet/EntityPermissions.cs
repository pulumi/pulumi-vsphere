// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    [VSphereResourceType("vsphere:index/entityPermissions:EntityPermissions")]
    public partial class EntityPermissions : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The managed object id (uuid for some entities) on
        /// which permissions are to be created.
        /// </summary>
        [Output("entityId")]
        public Output<string> EntityId { get; private set; } = null!;

        /// <summary>
        /// The managed object type, types can be found in the
        /// managed object type section
        /// [here](https://developer.broadcom.com/xapis/vsphere-web-services-api/latest/).
        /// </summary>
        [Output("entityType")]
        public Output<string> EntityType { get; private set; } = null!;

        /// <summary>
        /// The permissions to be given on this entity. Keep
        /// the permissions sorted alphabetically on `user_or_group` for a better user
        /// experience.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<Outputs.EntityPermissionsPermission>> Permissions { get; private set; } = null!;


        /// <summary>
        /// Create a EntityPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EntityPermissions(string name, EntityPermissionsArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/entityPermissions:EntityPermissions", name, args ?? new EntityPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EntityPermissions(string name, Input<string> id, EntityPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/entityPermissions:EntityPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EntityPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EntityPermissions Get(string name, Input<string> id, EntityPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new EntityPermissions(name, id, state, options);
        }
    }

    public sealed class EntityPermissionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The managed object id (uuid for some entities) on
        /// which permissions are to be created.
        /// </summary>
        [Input("entityId", required: true)]
        public Input<string> EntityId { get; set; } = null!;

        /// <summary>
        /// The managed object type, types can be found in the
        /// managed object type section
        /// [here](https://developer.broadcom.com/xapis/vsphere-web-services-api/latest/).
        /// </summary>
        [Input("entityType", required: true)]
        public Input<string> EntityType { get; set; } = null!;

        [Input("permissions", required: true)]
        private InputList<Inputs.EntityPermissionsPermissionArgs>? _permissions;

        /// <summary>
        /// The permissions to be given on this entity. Keep
        /// the permissions sorted alphabetically on `user_or_group` for a better user
        /// experience.
        /// </summary>
        public InputList<Inputs.EntityPermissionsPermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.EntityPermissionsPermissionArgs>());
            set => _permissions = value;
        }

        public EntityPermissionsArgs()
        {
        }
        public static new EntityPermissionsArgs Empty => new EntityPermissionsArgs();
    }

    public sealed class EntityPermissionsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The managed object id (uuid for some entities) on
        /// which permissions are to be created.
        /// </summary>
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        /// <summary>
        /// The managed object type, types can be found in the
        /// managed object type section
        /// [here](https://developer.broadcom.com/xapis/vsphere-web-services-api/latest/).
        /// </summary>
        [Input("entityType")]
        public Input<string>? EntityType { get; set; }

        [Input("permissions")]
        private InputList<Inputs.EntityPermissionsPermissionGetArgs>? _permissions;

        /// <summary>
        /// The permissions to be given on this entity. Keep
        /// the permissions sorted alphabetically on `user_or_group` for a better user
        /// experience.
        /// </summary>
        public InputList<Inputs.EntityPermissionsPermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.EntityPermissionsPermissionGetArgs>());
            set => _permissions = value;
        }

        public EntityPermissionsState()
        {
        }
        public static new EntityPermissionsState Empty => new EntityPermissionsState();
    }
}
