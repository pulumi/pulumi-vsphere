// *** WARNING: this file was generated by pulumi-language-nodejs. ***
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
 *     name: "production",
 *     categoryId: "data.vsphere_tag_category.environment.id",
 * });
 * const development = vsphere.getTag({
 *     name: "development",
 *     categoryId: "data.vsphere_tag_category.environment.id",
 * });
 * const platinum = vsphere.getTag({
 *     name: "platinum",
 *     categoryId: "data.vsphere_tag_category.service_level.id",
 * });
 * const gold = vsphere.getTag({
 *     name: "platinum",
 *     categoryId: "data.vsphere_tag_category.service_level.id",
 * });
 * const silver = vsphere.getTag({
 *     name: "silver",
 *     categoryId: "data.vsphere_tag_category.service_level.id",
 * });
 * const bronze = vsphere.getTag({
 *     name: "bronze",
 *     categoryId: "data.vsphere_tag_category.service_level.id",
 * });
 * const replicated = vsphere.getTag({
 *     name: "replicated",
 *     categoryId: "data.vsphere_tag_category.replication.id",
 * });
 * const nonReplicated = vsphere.getTag({
 *     name: "non_replicated",
 *     categoryId: "data.vsphere_tag_category.replication.id",
 * });
 * const prodDatastore = new vsphere.VmfsDatastore("prod_datastore", {tags: [
 *     "data.vsphere_tag.production.id",
 *     "data.vsphere_tag.platinum.id",
 *     "data.vsphere_tag.replicated.id",
 * ]});
 * const devDatastore = new vsphere.NasDatastore("dev_datastore", {tags: [
 *     "data.vsphere_tag.development.id",
 *     "data.vsphere_tag.silver.id",
 *     "data.vsphere_tag.non_replicated.id",
 * ]});
 * ```
 *
 * Next, storage policies are created and `tagRules` are applied.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const prodPlatinumReplicated = new vsphere.VmStoragePolicy("prod_platinum_replicated", {
 *     name: "prod_platinum_replicated",
 *     description: "prod_platinum_replicated",
 *     tagRules: [
 *         {
 *             tagCategory: environment.name,
 *             tags: [production.name],
 *             includeDatastoresWithTags: true,
 *         },
 *         {
 *             tagCategory: serviceLevel.name,
 *             tags: [platinum.name],
 *             includeDatastoresWithTags: true,
 *         },
 *         {
 *             tagCategory: replication.name,
 *             tags: [replicated.name],
 *             includeDatastoresWithTags: true,
 *         },
 *     ],
 * });
 * const devSilverNonreplicated = new vsphere.VmStoragePolicy("dev_silver_nonreplicated", {
 *     name: "dev_silver_nonreplicated",
 *     description: "dev_silver_nonreplicated",
 *     tagRules: [
 *         {
 *             tagCategory: environment.name,
 *             tags: [development.name],
 *             includeDatastoresWithTags: true,
 *         },
 *         {
 *             tagCategory: serviceLevel.name,
 *             tags: [silver.name],
 *             includeDatastoresWithTags: true,
 *         },
 *         {
 *             tagCategory: replication.name,
 *             tags: [nonReplicated.name],
 *             includeDatastoresWithTags: true,
 *         },
 *     ],
 * });
 * ```
 *
 * Lastly, when creating a virtual machine resource, a storage policy can be specified to direct virtual machine placement to a datastore which matches the policy's `tagsRules`.
 *
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
 * const prodVm = new vsphere.VirtualMachine("prod_vm", {storagePolicyId: storagePolicy.prodPlatinumReplicated.id});
 * const devVm = new vsphere.VirtualMachine("dev_vm", {storagePolicyId: storagePolicy.devSilverNonreplicated.id});
 * ```
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
