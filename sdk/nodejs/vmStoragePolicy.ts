// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `vsphere.VmStoragePolicy` resource can be used to create and manage storage
 * policies. Using this resource, tag based placement rules can be created to
 * place virtual machines on a datastore with matching tags. If storage requirements for the applications on the virtual machine change, you can modify the storage policy that was originally applied to the virtual machine.
 *
 * ## Example Usage
 *
 * The following example creates storage policies with `tagRules` base on sets of environment, service level, and replication attributes.
 *
 * In this example, tags are first applied to datastores.
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const environment = vsphere.getTagCategory({
 *     name: "environment",
 * });
 * const serviceLevel = vsphere.getTagCategory({
 *     name: "service_level",
 * });
 * const replication = vsphere.getTagCategory({
 *     name: "replication",
 * });
 * const production = vsphere.getTag({
 *     categoryId: "data.vsphere_tag_category.environment.id",
 *     name: "production",
 * });
 * const development = vsphere.getTag({
 *     categoryId: "data.vsphere_tag_category.environment.id",
 *     name: "development",
 * });
 * const platinum = vsphere.getTag({
 *     categoryId: "data.vsphere_tag_category.service_level.id",
 *     name: "platinum",
 * });
 * const gold = vsphere.getTag({
 *     categoryId: "data.vsphere_tag_category.service_level.id",
 *     name: "platinum",
 * });
 * const silver = vsphere.getTag({
 *     categoryId: "data.vsphere_tag_category.service_level.id",
 *     name: "silver",
 * });
 * const bronze = vsphere.getTag({
 *     categoryId: "data.vsphere_tag_category.service_level.id",
 *     name: "bronze",
 * });
 * const replicated = vsphere.getTag({
 *     categoryId: "data.vsphere_tag_category.replication.id",
 *     name: "replicated",
 * });
 * const nonReplicated = vsphere.getTag({
 *     categoryId: "data.vsphere_tag_category.replication.id",
 *     name: "non_replicated",
 * });
 * const prodDatastore = new vsphere.VmfsDatastore("prodDatastore", {tags: [
 *     "data.vsphere_tag.production.id",
 *     "data.vsphere_tag.platinum.id",
 *     "data.vsphere_tag.replicated.id",
 * ]});
 * const devDatastore = new vsphere.NasDatastore("devDatastore", {tags: [
 *     "data.vsphere_tag.development.id",
 *     "data.vsphere_tag.silver.id",
 *     "data.vsphere_tag.non_replicated.id",
 * ]});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Next, storage policies are created and `tagRules` are applied.
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const prodPlatinumReplicated = new vsphere.VmStoragePolicy("prodPlatinumReplicated", {
 *     description: "prod_platinum_replicated",
 *     tagRules: [
 *         {
 *             tagCategory: data.vsphere_tag_category.environment.name,
 *             tags: [data.vsphere_tag.production.name],
 *             includeDatastoresWithTags: true,
 *         },
 *         {
 *             tagCategory: data.vsphere_tag_category.service_level.name,
 *             tags: [data.vsphere_tag.platinum.name],
 *             includeDatastoresWithTags: true,
 *         },
 *         {
 *             tagCategory: data.vsphere_tag_category.replication.name,
 *             tags: [data.vsphere_tag.replicated.name],
 *             includeDatastoresWithTags: true,
 *         },
 *     ],
 * });
 * const devSilverNonreplicated = new vsphere.VmStoragePolicy("devSilverNonreplicated", {
 *     description: "dev_silver_nonreplicated",
 *     tagRules: [
 *         {
 *             tagCategory: data.vsphere_tag_category.environment.name,
 *             tags: [data.vsphere_tag.development.name],
 *             includeDatastoresWithTags: true,
 *         },
 *         {
 *             tagCategory: data.vsphere_tag_category.service_level.name,
 *             tags: [data.vsphere_tag.silver.name],
 *             includeDatastoresWithTags: true,
 *         },
 *         {
 *             tagCategory: data.vsphere_tag_category.replication.name,
 *             tags: [data.vsphere_tag.non_replicated.name],
 *             includeDatastoresWithTags: true,
 *         },
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Lasttly, when creating a virtual machine resource, a storage policy can be specificed to direct virtual machine placement to a datastore which matches the policy's `tagsRules`.
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const prodPlatinumReplicated = vsphere.getPolicy({
 *     name: "prod_platinum_replicated",
 * });
 * const devSilverNonreplicated = vsphere.getPolicy({
 *     name: "dev_silver_nonreplicated",
 * });
 * const prodVm = new vsphere.VirtualMachine("prodVm", {storagePolicyId: data.vsphere_storage_policy.storage_policy.prod_platinum_replicated.id});
 * // ... other configuration ...
 * const devVm = new vsphere.VirtualMachine("devVm", {storagePolicyId: data.vsphere_storage_policy.storage_policy.dev_silver_nonreplicated.id});
 * // ... other configuration ...
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class VmStoragePolicy extends pulumi.CustomResource {
    /**
     * Get an existing VmStoragePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VmStoragePolicyState, opts?: pulumi.CustomResourceOptions): VmStoragePolicy {
        return new VmStoragePolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/vmStoragePolicy:VmStoragePolicy';

    /**
     * Returns true if the given object is an instance of VmStoragePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VmStoragePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VmStoragePolicy.__pulumiType;
    }

    /**
     * Description of the storage policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the storage policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of tag rules. The tag category and tags to be associated to this storage policy.
     */
    public readonly tagRules!: pulumi.Output<outputs.VmStoragePolicyTagRule[]>;

    /**
     * Create a VmStoragePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VmStoragePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VmStoragePolicyArgs | VmStoragePolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VmStoragePolicyState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tagRules"] = state ? state.tagRules : undefined;
        } else {
            const args = argsOrState as VmStoragePolicyArgs | undefined;
            if ((!args || args.tagRules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tagRules'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tagRules"] = args ? args.tagRules : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VmStoragePolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VmStoragePolicy resources.
 */
export interface VmStoragePolicyState {
    /**
     * Description of the storage policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the storage policy.
     */
    name?: pulumi.Input<string>;
    /**
     * List of tag rules. The tag category and tags to be associated to this storage policy.
     */
    tagRules?: pulumi.Input<pulumi.Input<inputs.VmStoragePolicyTagRule>[]>;
}

/**
 * The set of arguments for constructing a VmStoragePolicy resource.
 */
export interface VmStoragePolicyArgs {
    /**
     * Description of the storage policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the storage policy.
     */
    name?: pulumi.Input<string>;
    /**
     * List of tag rules. The tag category and tags to be associated to this storage policy.
     */
    tagRules: pulumi.Input<pulumi.Input<inputs.VmStoragePolicyTagRule>[]>;
}
