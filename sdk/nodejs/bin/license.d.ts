import * as pulumi from "@pulumi/pulumi";
/**
 * Provides a VMware vSphere license resource. This can be used to add and remove license keys.
 */
export declare class License extends pulumi.CustomResource {
    /**
     * Get an existing License resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LicenseState): License;
    /**
     * The product edition of the license key.
     */
    readonly editionKey: pulumi.Output<string>;
    /**
     * A map of key/value pairs to be attached as labels (tags) to the license key.
     */
    readonly labels: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The license key to add.
     */
    readonly licenseKey: pulumi.Output<string>;
    /**
     * The display name for the license.
     */
    readonly name: pulumi.Output<string>;
    /**
     * Total number of units (example: CPUs) contained in the license.
     */
    readonly total: pulumi.Output<number>;
    /**
     * The number of units (example: CPUs) assigned to this license.
     */
    readonly used: pulumi.Output<number>;
    /**
     * Create a License resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LicenseArgs, opts?: pulumi.CustomResourceOptions);
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
    readonly labels?: pulumi.Input<{
        [key: string]: any;
    }>;
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
    readonly labels?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The license key to add.
     */
    readonly licenseKey: pulumi.Input<string>;
}
