// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the vsphere package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'vsphere';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * govmomi debug path for debug
     */
    public readonly clientDebugPath!: pulumi.Output<string | undefined>;
    /**
     * govmomi debug path for a single run
     */
    public readonly clientDebugPathRun!: pulumi.Output<string | undefined>;
    /**
     * The user password for vSphere API operations.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The directory to save vSphere REST API sessions to
     */
    public readonly restSessionPath!: pulumi.Output<string | undefined>;
    /**
     * The user name for vSphere API operations.
     */
    public readonly user!: pulumi.Output<string>;
    /**
     * @deprecated This field has been renamed to vsphere_server.
     */
    public readonly vcenterServer!: pulumi.Output<string | undefined>;
    /**
     * The directory to save vSphere SOAP API sessions to
     */
    public readonly vimSessionPath!: pulumi.Output<string | undefined>;
    /**
     * The vSphere Server name for vSphere API operations.
     */
    public readonly vsphereServer!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.user === undefined) && !opts.urn) {
                throw new Error("Missing required property 'user'");
            }
            resourceInputs["allowUnverifiedSsl"] = pulumi.output((args ? args.allowUnverifiedSsl : undefined) ?? utilities.getEnvBoolean("VSPHERE_ALLOW_UNVERIFIED_SSL")).apply(JSON.stringify);
            resourceInputs["apiTimeout"] = pulumi.output(args ? args.apiTimeout : undefined).apply(JSON.stringify);
            resourceInputs["clientDebug"] = pulumi.output((args ? args.clientDebug : undefined) ?? utilities.getEnvBoolean("VSPHERE_CLIENT_DEBUG")).apply(JSON.stringify);
            resourceInputs["clientDebugPath"] = (args ? args.clientDebugPath : undefined) ?? utilities.getEnv("VSPHERE_CLIENT_DEBUG_PATH");
            resourceInputs["clientDebugPathRun"] = (args ? args.clientDebugPathRun : undefined) ?? utilities.getEnv("VSPHERE_CLIENT_DEBUG_PATH_RUN");
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["persistSession"] = pulumi.output((args ? args.persistSession : undefined) ?? utilities.getEnvBoolean("VSPHERE_PERSIST_SESSION")).apply(JSON.stringify);
            resourceInputs["restSessionPath"] = (args ? args.restSessionPath : undefined) ?? utilities.getEnv("VSPHERE_REST_SESSION_PATH");
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["vcenterServer"] = args ? args.vcenterServer : undefined;
            resourceInputs["vimKeepAlive"] = pulumi.output((args ? args.vimKeepAlive : undefined) ?? utilities.getEnvNumber("VSPHERE_VIM_KEEP_ALIVE")).apply(JSON.stringify);
            resourceInputs["vimSessionPath"] = (args ? args.vimSessionPath : undefined) ?? utilities.getEnv("VSPHERE_VIM_SESSION_PATH");
            resourceInputs["vsphereServer"] = args ? args.vsphereServer : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * If set, VMware vSphere client will permit unverifiable SSL certificates.
     */
    allowUnverifiedSsl?: pulumi.Input<boolean>;
    /**
     * API timeout in minutes (Default: 5)
     */
    apiTimeout?: pulumi.Input<number>;
    /**
     * govmomi debug
     */
    clientDebug?: pulumi.Input<boolean>;
    /**
     * govmomi debug path for debug
     */
    clientDebugPath?: pulumi.Input<string>;
    /**
     * govmomi debug path for a single run
     */
    clientDebugPathRun?: pulumi.Input<string>;
    /**
     * The user password for vSphere API operations.
     */
    password: pulumi.Input<string>;
    /**
     * Persist vSphere client sessions to disk
     */
    persistSession?: pulumi.Input<boolean>;
    /**
     * The directory to save vSphere REST API sessions to
     */
    restSessionPath?: pulumi.Input<string>;
    /**
     * The user name for vSphere API operations.
     */
    user: pulumi.Input<string>;
    /**
     * @deprecated This field has been renamed to vsphere_server.
     */
    vcenterServer?: pulumi.Input<string>;
    /**
     * Keep alive interval for the VIM session in minutes
     */
    vimKeepAlive?: pulumi.Input<number>;
    /**
     * The directory to save vSphere SOAP API sessions to
     */
    vimSessionPath?: pulumi.Input<string>;
    /**
     * The vSphere Server name for vSphere API operations.
     */
    vsphereServer?: pulumi.Input<string>;
}
