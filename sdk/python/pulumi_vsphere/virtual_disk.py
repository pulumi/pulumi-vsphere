# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from . import utilities

class VirtualDisk(pulumi.CustomResource):
    """
    The `vsphere_virtual_disk` resource can be used to create virtual disks outside
    of any given [`vsphere_virtual_machine`][docs-vsphere-virtual-machine]
    resource. These disks can be attached to a virtual machine by creating a disk
    block with the [`attach`][docs-vsphere-virtual-machine-disk-attach] parameter.
    
    [docs-vsphere-virtual-machine]: /docs/providers/vsphere/r/virtual_machine.html
    [docs-vsphere-virtual-machine-disk-attach]: /docs/providers/vsphere/r/virtual_machine.html#attach
    """
    def __init__(__self__, __name__, __opts__=None, adapter_type=None, create_directories=None, datacenter=None, datastore=None, size=None, type=None, vmdk_path=None):
        """Create a VirtualDisk resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['adapterType'] = adapter_type

        __props__['createDirectories'] = create_directories

        __props__['datacenter'] = datacenter

        if not datastore:
            raise TypeError('Missing required property datastore')
        __props__['datastore'] = datastore

        if not size:
            raise TypeError('Missing required property size')
        __props__['size'] = size

        __props__['type'] = type

        if not vmdk_path:
            raise TypeError('Missing required property vmdk_path')
        __props__['vmdkPath'] = vmdk_path

        super(VirtualDisk, __self__).__init__(
            'vsphere:index/virtualDisk:VirtualDisk',
            __name__,
            __props__,
            __opts__)

