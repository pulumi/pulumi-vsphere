// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere_vapp_container` data source can be used to discover the ID of a
 * vApp container in vSphere. This is useful to fetch the ID of a vApp container
 * that you want to use to create virtual machines in using the
 * [`vsphere_virtual_machine`][docs-virtual-machine-resource] resource. 
 * 
 * [docs-virtual-machine-resource]: /docs/providers/vsphere/r/virtual_machine.html
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const datacenter = pulumi.output(vsphere.getDatacenter({
 *     name: "dc1",
 * }));
 * const pool = datacenter.apply(datacenter => vsphere.getVappContainer({
 *     datacenterId: datacenter.id,
 *     name: "vapp-container-1",
 * }));
 * ```
 */
export function getVappContainer(args: GetVappContainerArgs, opts?: pulumi.InvokeOptions): Promise<GetVappContainerResult> {
    return pulumi.runtime.invoke("vsphere:index/getVappContainer:getVappContainer", {
        "datacenterId": args.datacenterId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getVappContainer.
 */
export interface GetVappContainerArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datacenter the vApp container is located in.
     */
    readonly datacenterId: string;
    /**
     * The name of the vApp container. This can be a name or
     * path.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getVappContainer.
 */
export interface GetVappContainerResult {
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
