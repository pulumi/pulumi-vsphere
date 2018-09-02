"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
/**
 * The `vsphere_compute_cluster_vm_affinity_rule` resource can be used to manage
 * VM affinity rules in a cluster, either created by the
 * [`vsphere_compute_cluster`][tf-vsphere-cluster-resource] resource or looked up
 * by the [`vsphere_compute_cluster`][tf-vsphere-cluster-data-source] data source.
 *
 * [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
 * [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
 *
 * This rule can be used to tell a set to virtual machines to run together on a
 * single host within a cluster. When configured, DRS will make a best effort to
 * ensure that the virtual machines run on the same host, or prevent any operation
 * that would keep that from happening, depending on the value of the
 * `mandatory` flag.
 *
 * -> Keep in mind that this rule can only be used to tell VMs to run together on
 * a _non-specific_ host - it can't be used to pin VMs to a host. For that, see
 * the
 * [`vsphere_compute_cluster_vm_host_rule`][tf-vsphere-cluster-vm-host-rule-resource]
 * resource.
 *
 * [tf-vsphere-cluster-vm-host-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_host_rule.html
 *
 * ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 *
 * ~> **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
 */
class ComputeClusterVmAffinityRule extends pulumi.CustomResource {
    /**
     * Get an existing ComputeClusterVmAffinityRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name, id, state) {
        return new ComputeClusterVmAffinityRule(name, state, { id });
    }
    constructor(name, argsOrState, opts) {
        let inputs = {};
        if (opts && opts.id) {
            const state = argsOrState;
            inputs["computeClusterId"] = state ? state.computeClusterId : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["mandatory"] = state ? state.mandatory : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["virtualMachineIds"] = state ? state.virtualMachineIds : undefined;
        }
        else {
            const args = argsOrState;
            if (!args || args.computeClusterId === undefined) {
                throw new Error("Missing required property 'computeClusterId'");
            }
            if (!args || args.virtualMachineIds === undefined) {
                throw new Error("Missing required property 'virtualMachineIds'");
            }
            inputs["computeClusterId"] = args ? args.computeClusterId : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["mandatory"] = args ? args.mandatory : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["virtualMachineIds"] = args ? args.virtualMachineIds : undefined;
        }
        super("vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule", name, inputs, opts);
    }
}
exports.ComputeClusterVmAffinityRule = ComputeClusterVmAffinityRule;
//# sourceMappingURL=computeClusterVmAffinityRule.js.map