// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a VMware vSphere host resource. This represents an ESXi host that
 * can be used either as a member of a cluster or as a standalone host.
 *
 * ## Example Usage
 *
 * ### Create a standalone host
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const thumbprint = vsphere.getHostThumbprint({
 *     address: "esxi-01.example.com",
 *     insecure: true,
 * });
 * const esx_01 = new vsphere.Host("esx-01", {
 *     hostname: "esxi-01.example.com",
 *     username: "root",
 *     password: "password",
 *     license: "00000-00000-00000-00000-00000",
 *     thumbprint: thumbprint.then(thumbprint => thumbprint.id),
 *     datacenter: datacenter.then(datacenter => datacenter.id),
 * });
 * ```
 *
 * ### Create host in a compute cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const cluster = datacenter.then(datacenter => vsphere.getComputeCluster({
 *     name: "cluster-01",
 *     datacenterId: datacenter.id,
 * }));
 * const thumbprint = vsphere.getHostThumbprint({
 *     address: "esxi-01.example.com",
 *     insecure: true,
 * });
 * const esx_01 = new vsphere.Host("esx-01", {
 *     hostname: "esxi-01.example.com",
 *     username: "root",
 *     password: "password",
 *     license: "00000-00000-00000-00000-00000",
 *     thumbprint: thumbprint.then(thumbprint => thumbprint.id),
 *     cluster: cluster.then(cluster => cluster.id),
 *     services: [{
 *         ntpd: {
 *             enabled: true,
 *             policy: "on",
 *             ntpServers: ["pool.ntp.org"],
 *         },
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * An existing host can be imported into this resource by supplying
 *
 * the host's ID.
 *
 * [docs-import]: /docs/import/index.html
 *
 * Obtain the host's ID using the data source. For example:
 *
 * hcl
 *
 * data "vsphere_datacenter" "datacenter" {
 *
 *   name = "dc-01"
 *
 * }
 *
 * data "vsphere_host" "host" {
 *
 *   name          = "esxi-01.example.com"
 *
 *   datacenter_id = data.vsphere_datacenter.datacenter.id
 *
 * }
 *
 * output "host_id" {
 *
 *   value = data.vsphere_host.host.id
 *
 * }
 *
 * Next, create a resource configuration, For example:
 *
 * hcl
 *
 * data "vsphere_datacenter" "datacenter" {
 *
 *   name = "dc-01"
 *
 * }
 *
 * data "vsphere_host_thumbprint" "thumbprint" {
 *
 *   address  = "esxi-01.example.com"
 *
 *   insecure = true
 *
 * }
 *
 * resource "vsphere_host" "esx-01" {
 *
 *   hostname   = "esxi-01.example.com"
 *
 *   username   = "root"
 *
 *   password   = "password"
 *
 *   thumbprint = data.vsphere_host_thumbprint.thumbprint.id
 *
 *   datacenter = data.vsphere_datacenter.datacenter.id
 *
 * }
 *
 * hcl
 *
 * resource "vsphere_host" "esx-01" {
 *
 *   hostname   = "esxi-01.example.com"
 *
 *   username   = "root"
 *
 *   password   = "password"
 *
 *   license    = "00000-00000-00000-00000-00000"
 *
 *   thumbprint = data.vsphere_host_thumbprint.thumbprint.id
 *
 *   cluster    = data.vsphere_compute_cluster.cluster.id
 *
 *   services {
 *
 *     ntpd {
 *     
 *       enabled     = true
 *     
 *       policy      = "on"
 *     
 *       ntp_servers = ["pool.ntp.org"]
 *     
 *     }
 *
 *   }
 *
 * }
 *
 * ```sh
 * $ pulumi import vsphere:index/host:Host esx-01 host-123
 * ```
 *
 * The above would import the host `esxi-01.example.com` with the host ID `host-123`.
 */
export class Host extends pulumi.CustomResource {
    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostState, opts?: pulumi.CustomResourceOptions): Host {
        return new Host(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/host:Host';

    /**
     * Returns true if the given object is an instance of Host.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Host {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Host.__pulumiType;
    }

    /**
     * The ID of the Compute Cluster this host should
     * be added to. This should not be set if `datacenter` is set. Conflicts with:
     * `clusterManaged`.
     */
    public readonly cluster!: pulumi.Output<string | undefined>;
    /**
     * Can be set to `true` if compute cluster
     * membership will be managed through the `computeCluster` resource rather
     * than the`host` resource. Conflicts with: `cluster`.
     */
    public readonly clusterManaged!: pulumi.Output<boolean | undefined>;
    /**
     * If set to false then the host will be disconnected.
     * Default is `false`.
     */
    public readonly connected!: pulumi.Output<boolean | undefined>;
    /**
     * A map of custom attribute IDs and string
     * values to apply to the resource. Please refer to the
     * `vsphereCustomAttributes` resource for more information on applying
     * tags to resources.
     *
     * > **NOTE:** Custom attributes are not supported on direct ESXi host
     * connections and require vCenter Server.
     *
     * [docs-host-thumbprint-data-source]: /docs/providers/vsphere/d/host_thumbprint.html
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the datacenter this host should
     * be added to. This should not be set if `cluster` is set.
     */
    public readonly datacenter!: pulumi.Output<string | undefined>;
    /**
     * If set to `true` then it will force the host to be added,
     * even if the host is already connected to a different vCenter Server instance.
     * Default is `false`.
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    /**
     * FQDN or IP address of the host to be added.
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * The license key that will be applied to the host.
     * The license key is expected to be present in vSphere.
     */
    public readonly license!: pulumi.Output<string | undefined>;
    /**
     * Set the lockdown state of the host. Valid options are
     * `disabled`, `normal`, and `strict`. Default is `disabled`.
     */
    public readonly lockdown!: pulumi.Output<string | undefined>;
    /**
     * Set the management state of the host.
     * Default is `false`.
     */
    public readonly maintenance!: pulumi.Output<boolean | undefined>;
    /**
     * Password that will be used by vSphere to authenticate
     * to the host.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Set Services on host, the settings to be set are based on service being set as part of import.
     */
    public readonly services!: pulumi.Output<outputs.HostService[] | undefined>;
    /**
     * The IDs of any tags to attach to this resource. Please
     * refer to the `vsphere.Tag` resource for more information on applying
     * tags to resources.
     *
     * > **NOTE:** Tagging support is not supported on direct ESXi host
     * connections and require vCenter Server.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Host's certificate SHA-1 thumbprint. If not set the
     * CA that signed the host's certificate should be trusted. If the CA is not
     * trusted and no thumbprint is set then the operation will fail. See data source
     * [`vsphere.getHostThumbprint`][docs-host-thumbprint-data-source].
     */
    public readonly thumbprint!: pulumi.Output<string | undefined>;
    /**
     * Username that will be used by vSphere to authenticate
     * to the host.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a Host resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostArgs | HostState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostState | undefined;
            resourceInputs["cluster"] = state ? state.cluster : undefined;
            resourceInputs["clusterManaged"] = state ? state.clusterManaged : undefined;
            resourceInputs["connected"] = state ? state.connected : undefined;
            resourceInputs["customAttributes"] = state ? state.customAttributes : undefined;
            resourceInputs["datacenter"] = state ? state.datacenter : undefined;
            resourceInputs["force"] = state ? state.force : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["license"] = state ? state.license : undefined;
            resourceInputs["lockdown"] = state ? state.lockdown : undefined;
            resourceInputs["maintenance"] = state ? state.maintenance : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["thumbprint"] = state ? state.thumbprint : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as HostArgs | undefined;
            if ((!args || args.hostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostname'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["cluster"] = args ? args.cluster : undefined;
            resourceInputs["clusterManaged"] = args ? args.clusterManaged : undefined;
            resourceInputs["connected"] = args ? args.connected : undefined;
            resourceInputs["customAttributes"] = args ? args.customAttributes : undefined;
            resourceInputs["datacenter"] = args ? args.datacenter : undefined;
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["license"] = args ? args.license : undefined;
            resourceInputs["lockdown"] = args ? args.lockdown : undefined;
            resourceInputs["maintenance"] = args ? args.maintenance : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["thumbprint"] = args ? args.thumbprint : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Host.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Host resources.
 */
export interface HostState {
    /**
     * The ID of the Compute Cluster this host should
     * be added to. This should not be set if `datacenter` is set. Conflicts with:
     * `clusterManaged`.
     */
    cluster?: pulumi.Input<string>;
    /**
     * Can be set to `true` if compute cluster
     * membership will be managed through the `computeCluster` resource rather
     * than the`host` resource. Conflicts with: `cluster`.
     */
    clusterManaged?: pulumi.Input<boolean>;
    /**
     * If set to false then the host will be disconnected.
     * Default is `false`.
     */
    connected?: pulumi.Input<boolean>;
    /**
     * A map of custom attribute IDs and string
     * values to apply to the resource. Please refer to the
     * `vsphereCustomAttributes` resource for more information on applying
     * tags to resources.
     *
     * > **NOTE:** Custom attributes are not supported on direct ESXi host
     * connections and require vCenter Server.
     *
     * [docs-host-thumbprint-data-source]: /docs/providers/vsphere/d/host_thumbprint.html
     */
    customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the datacenter this host should
     * be added to. This should not be set if `cluster` is set.
     */
    datacenter?: pulumi.Input<string>;
    /**
     * If set to `true` then it will force the host to be added,
     * even if the host is already connected to a different vCenter Server instance.
     * Default is `false`.
     */
    force?: pulumi.Input<boolean>;
    /**
     * FQDN or IP address of the host to be added.
     */
    hostname?: pulumi.Input<string>;
    /**
     * The license key that will be applied to the host.
     * The license key is expected to be present in vSphere.
     */
    license?: pulumi.Input<string>;
    /**
     * Set the lockdown state of the host. Valid options are
     * `disabled`, `normal`, and `strict`. Default is `disabled`.
     */
    lockdown?: pulumi.Input<string>;
    /**
     * Set the management state of the host.
     * Default is `false`.
     */
    maintenance?: pulumi.Input<boolean>;
    /**
     * Password that will be used by vSphere to authenticate
     * to the host.
     */
    password?: pulumi.Input<string>;
    /**
     * Set Services on host, the settings to be set are based on service being set as part of import.
     */
    services?: pulumi.Input<pulumi.Input<inputs.HostService>[]>;
    /**
     * The IDs of any tags to attach to this resource. Please
     * refer to the `vsphere.Tag` resource for more information on applying
     * tags to resources.
     *
     * > **NOTE:** Tagging support is not supported on direct ESXi host
     * connections and require vCenter Server.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Host's certificate SHA-1 thumbprint. If not set the
     * CA that signed the host's certificate should be trusted. If the CA is not
     * trusted and no thumbprint is set then the operation will fail. See data source
     * [`vsphere.getHostThumbprint`][docs-host-thumbprint-data-source].
     */
    thumbprint?: pulumi.Input<string>;
    /**
     * Username that will be used by vSphere to authenticate
     * to the host.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Host resource.
 */
export interface HostArgs {
    /**
     * The ID of the Compute Cluster this host should
     * be added to. This should not be set if `datacenter` is set. Conflicts with:
     * `clusterManaged`.
     */
    cluster?: pulumi.Input<string>;
    /**
     * Can be set to `true` if compute cluster
     * membership will be managed through the `computeCluster` resource rather
     * than the`host` resource. Conflicts with: `cluster`.
     */
    clusterManaged?: pulumi.Input<boolean>;
    /**
     * If set to false then the host will be disconnected.
     * Default is `false`.
     */
    connected?: pulumi.Input<boolean>;
    /**
     * A map of custom attribute IDs and string
     * values to apply to the resource. Please refer to the
     * `vsphereCustomAttributes` resource for more information on applying
     * tags to resources.
     *
     * > **NOTE:** Custom attributes are not supported on direct ESXi host
     * connections and require vCenter Server.
     *
     * [docs-host-thumbprint-data-source]: /docs/providers/vsphere/d/host_thumbprint.html
     */
    customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the datacenter this host should
     * be added to. This should not be set if `cluster` is set.
     */
    datacenter?: pulumi.Input<string>;
    /**
     * If set to `true` then it will force the host to be added,
     * even if the host is already connected to a different vCenter Server instance.
     * Default is `false`.
     */
    force?: pulumi.Input<boolean>;
    /**
     * FQDN or IP address of the host to be added.
     */
    hostname: pulumi.Input<string>;
    /**
     * The license key that will be applied to the host.
     * The license key is expected to be present in vSphere.
     */
    license?: pulumi.Input<string>;
    /**
     * Set the lockdown state of the host. Valid options are
     * `disabled`, `normal`, and `strict`. Default is `disabled`.
     */
    lockdown?: pulumi.Input<string>;
    /**
     * Set the management state of the host.
     * Default is `false`.
     */
    maintenance?: pulumi.Input<boolean>;
    /**
     * Password that will be used by vSphere to authenticate
     * to the host.
     */
    password: pulumi.Input<string>;
    /**
     * Set Services on host, the settings to be set are based on service being set as part of import.
     */
    services?: pulumi.Input<pulumi.Input<inputs.HostService>[]>;
    /**
     * The IDs of any tags to attach to this resource. Please
     * refer to the `vsphere.Tag` resource for more information on applying
     * tags to resources.
     *
     * > **NOTE:** Tagging support is not supported on direct ESXi host
     * connections and require vCenter Server.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Host's certificate SHA-1 thumbprint. If not set the
     * CA that signed the host's certificate should be trusted. If the CA is not
     * trusted and no thumbprint is set then the operation will fail. See data source
     * [`vsphere.getHostThumbprint`][docs-host-thumbprint-data-source].
     */
    thumbprint?: pulumi.Input<string>;
    /**
     * Username that will be used by vSphere to authenticate
     * to the host.
     */
    username: pulumi.Input<string>;
}
