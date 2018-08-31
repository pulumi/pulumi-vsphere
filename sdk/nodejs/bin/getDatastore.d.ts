import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_datastore` data source can be used to discover the ID of a
 * datastore in vSphere. This is useful to fetch the ID of a datastore that you
 * want to use to create virtual machines in using the
 * [`vsphere_virtual_machine`][docs-virtual-machine-resource] resource.
 *
 * [docs-virtual-machine-resource]: /docs/providers/vsphere/r/virtual_machine.html
 */
export declare function getDatastore(args: GetDatastoreArgs, opts?: pulumi.InvokeOptions): Promise<GetDatastoreResult>;
/**
 * A collection of arguments for invoking getDatastore.
 */
export interface GetDatastoreArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datacenter the datastore is located in. This
     * can be omitted if the search path used in `name` is an absolute path. For
     * default datacenters, use the id attribute from an empty `vsphere_datacenter`
     * data source.
     */
    readonly datacenterId?: string;
    /**
     * The name of the datastore. This can be a name or path.
     */
    readonly name: string;
}
/**
 * A collection of values returned by getDatastore.
 */
export interface GetDatastoreResult {
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
