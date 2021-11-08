// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Role extends pulumi.CustomResource {
    /**
     * Get an existing Role resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoleState, opts?: pulumi.CustomResourceOptions): Role {
        return new Role(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/role:Role';

    /**
     * Returns true if the given object is an instance of Role.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Role {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Role.__pulumiType;
    }

    /**
     * The display label of the role.
     */
    public /*out*/ readonly label!: pulumi.Output<string>;
    /**
     * The name of the role.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The privileges to be associated with this role.
     */
    public readonly rolePrivileges!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Role resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoleArgs | RoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RoleState | undefined;
            inputs["label"] = state ? state.label : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["rolePrivileges"] = state ? state.rolePrivileges : undefined;
        } else {
            const args = argsOrState as RoleArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["rolePrivileges"] = args ? args.rolePrivileges : undefined;
            inputs["label"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Role.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Role resources.
 */
export interface RoleState {
    /**
     * The display label of the role.
     */
    label?: pulumi.Input<string>;
    /**
     * The name of the role.
     */
    name?: pulumi.Input<string>;
    /**
     * The privileges to be associated with this role.
     */
    rolePrivileges?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Role resource.
 */
export interface RoleArgs {
    /**
     * The name of the role.
     */
    name?: pulumi.Input<string>;
    /**
     * The privileges to be associated with this role.
     */
    rolePrivileges?: pulumi.Input<pulumi.Input<string>[]>;
}
