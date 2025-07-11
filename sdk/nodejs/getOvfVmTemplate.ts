// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.getOvfVmTemplate` data source can be used to submit an OVF to
 * vSphere and extract its hardware settings in a form that can be then used as
 * inputs for a `vsphere.VirtualMachine` resource.
 */
export function getOvfVmTemplate(args: GetOvfVmTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetOvfVmTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
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
     * disks included in the OVF/OVA will have the same specified policy. Can be
     * one of `thin`, `thick`, `eagerZeroedThick`, or `sameAsSource`.
     */
    diskProvisioning?: string;
    /**
     * Allow properties with
     * `ovf:userConfigurable=false` to be set.
     */
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
     *
     * > **NOTE:** Either `localOvfPath` or `remoteOvfUrl` is required, both can
     * not be empty.
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
/**
 * The `vsphere.getOvfVmTemplate` data source can be used to submit an OVF to
 * vSphere and extract its hardware settings in a form that can be then used as
 * inputs for a `vsphere.VirtualMachine` resource.
 */
export function getOvfVmTemplateOutput(args: GetOvfVmTemplateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOvfVmTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vsphere:index/getOvfVmTemplate:getOvfVmTemplate", {
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
     * disks included in the OVF/OVA will have the same specified policy. Can be
     * one of `thin`, `thick`, `eagerZeroedThick`, or `sameAsSource`.
     */
    diskProvisioning?: pulumi.Input<string>;
    /**
     * Allow properties with
     * `ovf:userConfigurable=false` to be set.
     */
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
     *
     * > **NOTE:** Either `localOvfPath` or `remoteOvfUrl` is required, both can
     * not be empty.
     */
    remoteOvfUrl?: pulumi.Input<string>;
    /**
     * The ID of a resource pool in which to place
     * the virtual machine.
     */
    resourcePoolId: pulumi.Input<string>;
}
