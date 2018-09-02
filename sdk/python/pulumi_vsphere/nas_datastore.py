# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class NasDatastore(pulumi.CustomResource):
    """
    The `vsphere_nas_datastore` resource can be used to create and manage NAS
    datastores on an ESXi host or a set of hosts. The resource supports mounting
    NFS v3 and v4.1 shares to be used as datastores.
    
    ~> **NOTE:** Unlike [`vsphere_vmfs_datastore`][resource-vmfs-datastore], a NAS
    datastore is only mounted on the hosts you choose to mount it on. To mount on
    multiple hosts, you must specify each host that you want to add in the
    `host_system_ids` argument.
    
    [resource-vmfs-datastore]: /docs/providers/vsphere/r/vmfs_datastore.html
    """
    def __init__(__self__, __name__, __opts__=None, access_mode=None, custom_attributes=None, datastore_cluster_id=None, folder=None, host_system_ids=None, name=None, remote_hosts=None, remote_path=None, security_type=None, tags=None, type=None):
        """Create a NasDatastore resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if access_mode and not isinstance(access_mode, basestring):
            raise TypeError('Expected property access_mode to be a basestring')
        __self__.access_mode = access_mode
        """
        Access mode for the mount point. Can be one of
        `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
        that the datastore will be read-write depending on the permissions of the
        actual share. Default: `readWrite`. Forces a new resource if changed.
        """
        __props__['accessMode'] = access_mode

        if custom_attributes and not isinstance(custom_attributes, dict):
            raise TypeError('Expected property custom_attributes to be a dict')
        __self__.custom_attributes = custom_attributes
        """
        Map of custom attribute ids to attribute 
        value strings to set on datasource resource. See
        [here][docs-setting-custom-attributes] for a reference on how to set values
        for custom attributes.
        """
        __props__['customAttributes'] = custom_attributes

        if datastore_cluster_id and not isinstance(datastore_cluster_id, basestring):
            raise TypeError('Expected property datastore_cluster_id to be a basestring')
        __self__.datastore_cluster_id = datastore_cluster_id
        """
        The [managed object
        ID][docs-about-morefs] of a datastore cluster to put this datastore in.
        Conflicts with `folder`.
        """
        __props__['datastoreClusterId'] = datastore_cluster_id

        if folder and not isinstance(folder, basestring):
            raise TypeError('Expected property folder to be a basestring')
        __self__.folder = folder
        """
        The relative path to a folder to put this datastore in.
        This is a path relative to the datacenter you are deploying the datastore to.
        Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
        Terraform will place a datastore named `terraform-test` in a datastore folder
        located at `/dc1/datastore/foo/bar`, with the final inventory path being
        `/dc1/datastore/foo/bar/terraform-test`. Conflicts with
        `datastore_cluster_id`.
        """
        __props__['folder'] = folder

        if not host_system_ids:
            raise TypeError('Missing required property host_system_ids')
        elif not isinstance(host_system_ids, list):
            raise TypeError('Expected property host_system_ids to be a list')
        __self__.host_system_ids = host_system_ids
        """
        The [managed object IDs][docs-about-morefs] of
        the hosts to mount the datastore on.
        """
        __props__['hostSystemIds'] = host_system_ids

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the datastore. Forces a new resource if
        changed.
        """
        __props__['name'] = name

        if not remote_hosts:
            raise TypeError('Missing required property remote_hosts')
        elif not isinstance(remote_hosts, list):
            raise TypeError('Expected property remote_hosts to be a list')
        __self__.remote_hosts = remote_hosts
        """
        The hostnames or IP addresses of the remote
        server or servers. Only one element should be present for NFS v3 but multiple
        can be present for NFS v4.1. Forces a new resource if changed.
        """
        __props__['remoteHosts'] = remote_hosts

        if not remote_path:
            raise TypeError('Missing required property remote_path')
        elif not isinstance(remote_path, basestring):
            raise TypeError('Expected property remote_path to be a basestring')
        __self__.remote_path = remote_path
        """
        The remote path of the mount point. Forces a new
        resource if changed.
        """
        __props__['remotePath'] = remote_path

        if security_type and not isinstance(security_type, basestring):
            raise TypeError('Expected property security_type to be a basestring')
        __self__.security_type = security_type
        """
        The security type to use when using NFS v4.1.
        Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
        if changed.
        """
        __props__['securityType'] = security_type

        if tags and not isinstance(tags, list):
            raise TypeError('Expected property tags to be a list')
        __self__.tags = tags
        """
        The IDs of any tags to attach to this resource. See
        [here][docs-applying-tags] for a reference on how to apply tags.
        """
        __props__['tags'] = tags

        if type and not isinstance(type, basestring):
            raise TypeError('Expected property type to be a basestring')
        __self__.type = type
        """
        The type of NAS volume. Can be one of `NFS` (to denote
        v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
        changed.
        """
        __props__['type'] = type

        __self__.accessible = pulumi.runtime.UNKNOWN
        """
        The connectivity status of the datastore. If this is `false`,
        some other computed attributes may be out of date.
        """
        __self__.capacity = pulumi.runtime.UNKNOWN
        """
        Maximum capacity of the datastore, in megabytes.
        """
        __self__.free_space = pulumi.runtime.UNKNOWN
        """
        Available space of this datastore, in megabytes.
        """
        __self__.maintenance_mode = pulumi.runtime.UNKNOWN
        """
        The current maintenance mode state of the datastore.
        """
        __self__.multiple_host_access = pulumi.runtime.UNKNOWN
        """
        If `true`, more than one host in the datacenter has
        been configured with access to the datastore.
        """
        __self__.protocol_endpoint = pulumi.runtime.UNKNOWN
        """
        Indicates that this NAS volume is a protocol endpoint.
        This field is only populated if the host supports virtual datastores.
        """
        __self__.uncommitted_space = pulumi.runtime.UNKNOWN
        """
        Total additional storage space, in megabytes,
        potentially used by all virtual machines on this datastore.
        """
        __self__.url = pulumi.runtime.UNKNOWN
        """
        The unique locator for the datastore.
        """

        super(NasDatastore, __self__).__init__(
            'vsphere:index/nasDatastore:NasDatastore',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'accessMode' in outs:
            self.access_mode = outs['accessMode']
        if 'accessible' in outs:
            self.accessible = outs['accessible']
        if 'capacity' in outs:
            self.capacity = outs['capacity']
        if 'customAttributes' in outs:
            self.custom_attributes = outs['customAttributes']
        if 'datastoreClusterId' in outs:
            self.datastore_cluster_id = outs['datastoreClusterId']
        if 'folder' in outs:
            self.folder = outs['folder']
        if 'freeSpace' in outs:
            self.free_space = outs['freeSpace']
        if 'hostSystemIds' in outs:
            self.host_system_ids = outs['hostSystemIds']
        if 'maintenanceMode' in outs:
            self.maintenance_mode = outs['maintenanceMode']
        if 'multipleHostAccess' in outs:
            self.multiple_host_access = outs['multipleHostAccess']
        if 'name' in outs:
            self.name = outs['name']
        if 'protocolEndpoint' in outs:
            self.protocol_endpoint = outs['protocolEndpoint']
        if 'remoteHosts' in outs:
            self.remote_hosts = outs['remoteHosts']
        if 'remotePath' in outs:
            self.remote_path = outs['remotePath']
        if 'securityType' in outs:
            self.security_type = outs['securityType']
        if 'tags' in outs:
            self.tags = outs['tags']
        if 'type' in outs:
            self.type = outs['type']
        if 'uncommittedSpace' in outs:
            self.uncommitted_space = outs['uncommittedSpace']
        if 'url' in outs:
            self.url = outs['url']