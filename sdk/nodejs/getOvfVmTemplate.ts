// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.getOvfVmTemplate` data source can be used to submit an OVF to
 * vSphere and extract its hardware settings in a form that can be then used as
 * inputs for a `vsphere.VirtualMachine` resource.
 *
 * ## Remote OVF/OVA Source
 *
 * data "vsphere_ovf_vm_template" "ovfRemote" {
 *   name              = "Nested-ESXi-7.0-Terraform-Deploy-1"
 *   disk_provisioning = "thin"
 *   resource_pool_id  = data.vsphere_resource_pool.default.id
 *   datastore_id      = data.vsphere_datastore.datastore.id
 *   host_system_id    = data.vsphere_host.host.id
 *   remote_ovf_url    = "https://download3.vmware.com/software/vmw-tools/nested-esxi/Nested_ESXi7.0u3_Appliance_Template_v1.ova"
 *   ovf_network_map = {
 *     "VM Network" : data.vsphere_network.network.id
 *   }
 * }
 *
 * ## Local OVF/OVA Source
 *
 * data "vsphere_ovf_vm_template" "ovfLocal" {
 *   name              = "Nested-ESXi-7.0-Terraform-Deploy-2"
 *   disk_provisioning = "thin"
 *   resource_pool_id  = data.vsphere_resource_pool.default.id
 *   datastore_id      = data.vsphere_datastore.datastore.id
 *   host_system_id    = data.vsphere_host.host.id
 *   local_ovf_path    = "/Volume/Storage/OVA/Nested_ESXi7.0u3_Appliance_Template_v1.ova"
 *   ovf_network_map = {
 *     "VM Network" : data.vsphere_network.network.id
 *   }
 * }
 *
 * ## Deployment of VM from Remote OVF
 *
 * resource "vsphere_virtual_machine" "vmFromRemoteOvf" {
 *   name                 = "Nested-ESXi-7.0-Terraform-Deploy-1"
 *   datacenter_id        = data.vsphere_datacenter.datacenter.id
 *   datastore_id         = data.vsphere_datastore.datastore.id
 *   host_system_id       = data.vsphere_host.host.id
 *   resource_pool_id     = data.vsphere_resource_pool.default.id
 *   num_cpus             = data.vsphere_ovf_vm_template.ovfRemote.num_cpus
 *   num_cores_per_socket = data.vsphere_ovf_vm_template.ovfRemote.num_cores_per_socket
 *   memory               = data.vsphere_ovf_vm_template.ovfRemote.memory
 *   guest_id             = data.vsphere_ovf_vm_template.ovfRemote.guest_id
 *   firmware             = data.vsphere_ovf_vm_template.ovfRemote.firmware
 *   scsi_type            = data.vsphere_ovf_vm_template.ovfRemote.scsi_type
 *   nested_hv_enabled    = data.vsphere_ovf_vm_template.ovfRemote.nested_hv_enabled
 *   dynamic "network_interface" {
 *     for_each = data.vsphere_ovf_vm_template.ovfRemote.ovf_network_map
 *     content {
 *       network_id = network_interface.value
 *     }
 *   }
 *   wait_for_guest_net_timeout = 0
 *   wait_for_guest_ip_timeout  = 0
 *
 *   ovf_deploy {
 *     allow_unverified_ssl_cert = false
 *     remote_ovf_url            = data.vsphere_ovf_vm_template.ovfRemote.remote_ovf_url
 *     disk_provisioning         = data.vsphere_ovf_vm_template.ovfRemote.disk_provisioning
 *     ovf_network_map           = data.vsphere_ovf_vm_template.ovfRemote.ovf_network_map
 *   }
 *
 *   vapp {
 *     properties = {
 *       "guestinfo.hostname"  = "nested-esxi-01.example.com",
 *       "guestinfo.ipaddress" = "172.16.11.101",
 *       "guestinfo.netmask"   = "255.255.255.0",
 *       "guestinfo.gateway"   = "172.16.11.1",
 *       "guestinfo.dns"       = "172.16.11.4",
 *       "guestinfo.domain"    = "example.com",
 *       "guestinfo.ntp"       = "ntp.example.com",
 *       "guestinfo.password"  = "VMware1!",
 *       "guestinfo.ssh"       = "True"
 *     }
 *   }
 *
 *   lifecycle {
 *     ignore_changes = [
 *       annotation,
 *       disk[0].io_share_count,
 *       disk[1].io_share_count,
 *       disk[2].io_share_count,
 *       vapp[0].properties,
 *     ]
 *   }
 * }
 */
export function getOvfVmTemplate(args: GetOvfVmTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetOvfVmTemplateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vsphere:index/getOvfVmTemplate:getOvfVmTemplate", {
        "allowUnverifiedSslCert": args.allowUnverifiedSslCert,
        "datastoreId": args.datastoreId,
        "deploymentOption": args.deploymentOption,
        "diskProvisioning": args.diskProvisioning,
        "enableHiddenProperties": args.enableHiddenProperties,
        "folder": args.folder,
        "hostSystemId": args.hostSystemId,
        "ipAllocationPolicy": args.ipAllocationPolicy,
        "ipProtocol": args.ipProtocol,
        "localOvfPath": args.localOvfPath,
        "name": args.name,
        "ovfNetworkMap": args.ovfNetworkMap,
        "remoteOvfUrl": args.remoteOvfUrl,
        "resourcePoolId": args.resourcePoolId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOvfVmTemplate.
 */
export interface GetOvfVmTemplateArgs {
    /**
     * Allow unverified SSL certificates
     * when deploying OVF/OVA from a URL.
     */
    allowUnverifiedSslCert?: boolean;
    /**
     * The ID of the virtual machine's datastore. The
     * virtual machine configuration is placed here, along with any virtual disks
     * that are created without datastores.
     */
    datastoreId?: string;
    /**
     * The key of the chosen deployment option. If
     * empty, the default option is chosen.
     */
    deploymentOption?: string;
    /**
     * The disk provisioning type. If set, all the
     * disks in the deployed OVA/OVF will have the same specified disk type. Can be
     * one of `thin`, `flat`, `thick` or `sameAsSource`.
     */
    diskProvisioning?: string;
    enableHiddenProperties?: boolean;
    /**
     * The name of the folder in which to place the virtual
     * machine.
     */
    folder?: string;
    /**
     * The ID of the ESXi host system to deploy the
     * virtual machine.
     */
    hostSystemId: string;
    /**
     * The IP allocation policy.
     */
    ipAllocationPolicy?: string;
    /**
     * The IP protocol.
     */
    ipProtocol?: string;
    /**
     * The absolute path to the OVF/OVA file on the
     * local system. When deploying from an OVF, ensure all necessary files such as
     * the `.vmdk` files are present in the same directory as the OVF.
     */
    localOvfPath?: string;
    /**
     * Name of the virtual machine to create.
     */
    name: string;
    /**
     * The mapping of name of network identifiers
     * from the OVF descriptor to network UUID in the environment.
     */
    ovfNetworkMap?: {[key: string]: string};
    /**
     * URL of the remote OVF/OVA file to be deployed.
     */
    remoteOvfUrl?: string;
    /**
     * The ID of a resource pool in which to place
     * the virtual machine.
     */
    resourcePoolId: string;
}

/**
 * A collection of values returned by getOvfVmTemplate.
 */
export interface GetOvfVmTemplateResult {
    readonly allowUnverifiedSslCert?: boolean;
    /**
     * An alternate guest operating system name.
     */
    readonly alternateGuestName: string;
    /**
     * A description of the virtual machine.
     */
    readonly annotation: string;
    /**
     * Allow CPUs to be added to the virtual machine while
     * powered on.
     */
    readonly cpuHotAddEnabled: boolean;
    /**
     * Allow CPUs to be removed from the virtual machine
     * while powered on.
     */
    readonly cpuHotRemoveEnabled: boolean;
    readonly cpuPerformanceCountersEnabled: boolean;
    readonly datastoreId?: string;
    readonly deploymentOption?: string;
    readonly diskProvisioning?: string;
    readonly enableHiddenProperties?: boolean;
    /**
     * The firmware to use on the virtual machine.
     */
    readonly firmware: string;
    readonly folder?: string;
    /**
     * The ID for the guest operating system
     */
    readonly guestId: string;
    readonly hostSystemId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ideControllerCount: number;
    readonly ipAllocationPolicy?: string;
    readonly ipProtocol?: string;
    readonly localOvfPath?: string;
    /**
     * The size of the virtual machine memory, in MB.
     */
    readonly memory: number;
    /**
     * Allow memory to be added to the virtual machine
     * while powered on.
     */
    readonly memoryHotAddEnabled: boolean;
    readonly name: string;
    /**
     * Enable nested hardware virtualization on the virtual
     * machine, facilitating nested virtualization in the guest.
     */
    readonly nestedHvEnabled: boolean;
    /**
     * The number of cores per virtual CPU in the virtual
     * machine.
     */
    readonly numCoresPerSocket: number;
    /**
     * The number of virtual CPUs to assign to the virtual machine.
     */
    readonly numCpus: number;
    readonly ovfNetworkMap?: {[key: string]: string};
    readonly remoteOvfUrl?: string;
    readonly resourcePoolId: string;
    readonly sataControllerCount: number;
    readonly scsiControllerCount: number;
    readonly scsiType: string;
    /**
     * The swap file placement policy for the virtual
     * machine.
     */
    readonly swapPlacementPolicy: string;
}

export function getOvfVmTemplateOutput(args: GetOvfVmTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOvfVmTemplateResult> {
    return pulumi.output(args).apply(a => getOvfVmTemplate(a, opts))
}

/**
 * A collection of arguments for invoking getOvfVmTemplate.
 */
export interface GetOvfVmTemplateOutputArgs {
    /**
     * Allow unverified SSL certificates
     * when deploying OVF/OVA from a URL.
     */
    allowUnverifiedSslCert?: pulumi.Input<boolean>;
    /**
     * The ID of the virtual machine's datastore. The
     * virtual machine configuration is placed here, along with any virtual disks
     * that are created without datastores.
     */
    datastoreId?: pulumi.Input<string>;
    /**
     * The key of the chosen deployment option. If
     * empty, the default option is chosen.
     */
    deploymentOption?: pulumi.Input<string>;
    /**
     * The disk provisioning type. If set, all the
     * disks in the deployed OVA/OVF will have the same specified disk type. Can be
     * one of `thin`, `flat`, `thick` or `sameAsSource`.
     */
    diskProvisioning?: pulumi.Input<string>;
    enableHiddenProperties?: pulumi.Input<boolean>;
    /**
     * The name of the folder in which to place the virtual
     * machine.
     */
    folder?: pulumi.Input<string>;
    /**
     * The ID of the ESXi host system to deploy the
     * virtual machine.
     */
    hostSystemId: pulumi.Input<string>;
    /**
     * The IP allocation policy.
     */
    ipAllocationPolicy?: pulumi.Input<string>;
    /**
     * The IP protocol.
     */
    ipProtocol?: pulumi.Input<string>;
    /**
     * The absolute path to the OVF/OVA file on the
     * local system. When deploying from an OVF, ensure all necessary files such as
     * the `.vmdk` files are present in the same directory as the OVF.
     */
    localOvfPath?: pulumi.Input<string>;
    /**
     * Name of the virtual machine to create.
     */
    name: pulumi.Input<string>;
    /**
     * The mapping of name of network identifiers
     * from the OVF descriptor to network UUID in the environment.
     */
    ovfNetworkMap?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * URL of the remote OVF/OVA file to be deployed.
     */
    remoteOvfUrl?: pulumi.Input<string>;
    /**
     * The ID of a resource pool in which to place
     * the virtual machine.
     */
    resourcePoolId: pulumi.Input<string>;
}
