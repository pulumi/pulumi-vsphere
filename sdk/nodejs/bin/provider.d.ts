import * as pulumi from "@pulumi/pulumi";
/**
 * The provider type for the vsphere package
 */
export declare class Provider extends pulumi.ProviderResource {
    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions);
}
/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * If set, VMware vSphere client will permit unverifiable SSL certificates.
     */
    readonly allowUnverifiedSsl?: pulumi.Input<boolean>;
    /**
     * govmomi debug
     */
    readonly clientDebug?: pulumi.Input<boolean>;
    /**
     * govmomi debug path for debug
     */
    readonly clientDebugPath?: pulumi.Input<string>;
    /**
     * govmomi debug path for a single run
     */
    readonly clientDebugPathRun?: pulumi.Input<string>;
    /**
     * The user password for vSphere API operations.
     */
    readonly password: pulumi.Input<string>;
    /**
     * Persist vSphere client sessions to disk
     */
    readonly persistSession?: pulumi.Input<boolean>;
    /**
     * The directory to save vSphere REST API sessions to
     */
    readonly restSessionPath?: pulumi.Input<string>;
    /**
     * The user name for vSphere API operations.
     */
    readonly user: pulumi.Input<string>;
    readonly vcenterServer?: pulumi.Input<string>;
    /**
     * The directory to save vSphere SOAP API sessions to
     */
    readonly vimSessionPath?: pulumi.Input<string>;
    /**
     * The vSphere Server name for vSphere API operations.
     */
    readonly vsphereServer?: pulumi.Input<string>;
}
