// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a VMware vSphere license resource. This can be used to add and remove license keys.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const licenseKey = new vsphere.License("licenseKey", {
 *     labels: {
 *         VpxClientLicenseLabel: "Hello World",
 *         Workflow: "Hello World",
 *     },
 *     licenseKey: "452CQ-2EK54-K8742-00000-00000",
 * });
 * ```
 */
export class License extends pulumi.CustomResource {
    /**
     * Get an existing License resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LicenseState, opts?: pulumi.CustomResourceOptions): License {
        return new License(name, <any>state, { ...opts, id: id });
    }

    /**
     * The product edition of the license key.
     */
    public /*out*/ readonly editionKey!: pulumi.Output<string>;
    /**
     * A map of key/value pairs to be attached as labels (tags) to the license key.
     */
    public readonly labels!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The license key to add.
     */
    public readonly licenseKey!: pulumi.Output<string>;
    /**
     * The display name for the license.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Total number of units (example: CPUs) contained in the license.
     */
    public /*out*/ readonly total!: pulumi.Output<number>;
    /**
     * The number of units (example: CPUs) assigned to this license.
     */
    public /*out*/ readonly used!: pulumi.Output<number>;

    /**
     * Create a License resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LicenseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LicenseArgs | LicenseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LicenseState | undefined;
            inputs["editionKey"] = state ? state.editionKey : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["licenseKey"] = state ? state.licenseKey : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["total"] = state ? state.total : undefined;
            inputs["used"] = state ? state.used : undefined;
        } else {
            const args = argsOrState as LicenseArgs | undefined;
            if (!args || args.licenseKey === undefined) {
                throw new Error("Missing required property 'licenseKey'");
            }
            inputs["labels"] = args ? args.labels : undefined;
            inputs["licenseKey"] = args ? args.licenseKey : undefined;
            inputs["editionKey"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["total"] = undefined /*out*/;
            inputs["used"] = undefined /*out*/;
        }
        super("vsphere:index/license:License", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering License resources.
 */
export interface LicenseState {
    /**
     * The product edition of the license key.
     */
    readonly editionKey?: pulumi.Input<string>;
    /**
     * A map of key/value pairs to be attached as labels (tags) to the license key.
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The license key to add.
     */
    readonly licenseKey?: pulumi.Input<string>;
    /**
     * The display name for the license.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Total number of units (example: CPUs) contained in the license.
     */
    readonly total?: pulumi.Input<number>;
    /**
     * The number of units (example: CPUs) assigned to this license.
     */
    readonly used?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a License resource.
 */
export interface LicenseArgs {
    /**
     * A map of key/value pairs to be attached as labels (tags) to the license key.
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The license key to add.
     */
    readonly licenseKey: pulumi.Input<string>;
}
