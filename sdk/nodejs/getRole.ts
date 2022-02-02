// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.Role` data source can be used to discover the id and privileges associated
 * with a role given its name or display label in vsphere UI.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const role1 = pulumi.output(vsphere.getRole({
 *     label: "Virtual machine user (sample)",
 * }));
 * ```
 */
export function getRole(args: GetRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetRoleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vsphere:index/getRole:getRole", {
        "description": args.description,
        "label": args.label,
        "name": args.name,
        "rolePrivileges": args.rolePrivileges,
    }, opts);
}

/**
 * A collection of arguments for invoking getRole.
 */
export interface GetRoleArgs {
    /**
     * The description of the role.
     */
    description?: string;
    /**
     * The label of the role.
     */
    label: string;
    name?: string;
    /**
     * The privileges associated with the role.
     */
    rolePrivileges?: string[];
}

/**
 * A collection of values returned by getRole.
 */
export interface GetRoleResult {
    /**
     * The description of the role.
     */
    readonly description?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The display label of the role.
     */
    readonly label: string;
    readonly name?: string;
    /**
     * The privileges associated with the role.
     */
    readonly rolePrivileges?: string[];
}

export function getRoleOutput(args: GetRoleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRoleResult> {
    return pulumi.output(args).apply(a => getRole(a, opts))
}

/**
 * A collection of arguments for invoking getRole.
 */
export interface GetRoleOutputArgs {
    /**
     * The description of the role.
     */
    description?: pulumi.Input<string>;
    /**
     * The label of the role.
     */
    label: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * The privileges associated with the role.
     */
    rolePrivileges?: pulumi.Input<pulumi.Input<string>[]>;
}
