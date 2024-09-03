# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VmfsDatastoreArgs', 'VmfsDatastore']

@pulumi.input_type
class VmfsDatastoreArgs:
    def __init__(__self__, *,
                 disks: pulumi.Input[Sequence[pulumi.Input[str]]],
                 host_system_id: pulumi.Input[str],
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 datastore_cluster_id: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VmfsDatastore resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disks: The disks to use with the datastore.
        :param pulumi.Input[str] host_system_id: The managed object ID of
               the host to set the datastore up on. Note that this is not necessarily the
               only host that the datastore will be set up on - see
               here for more info. Forces a
               new resource if changed.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: Map of custom attribute ids to attribute
               value string to set on datastore resource.
               
               > **NOTE:** Custom attributes are unsupported on direct ESXi connections
               and require vCenter.
        :param pulumi.Input[str] datastore_cluster_id: The managed object
               ID of a datastore cluster to put this datastore in.
               Conflicts with `folder`.
        :param pulumi.Input[str] folder: The relative path to a folder to put this datastore in.
               This is a path relative to the datacenter you are deploying the datastore to.
               Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
               The provider will place a datastore named `test` in a datastore folder
               located at `/dc1/datastore/foo/bar`, with the final inventory path being
               `/dc1/datastore/foo/bar/test`. Conflicts with
               `datastore_cluster_id`.
        :param pulumi.Input[str] name: The name of the datastore. Forces a new resource if
               changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The IDs of any tags to attach to this resource.
               
               > **NOTE:** Tagging support is unsupported on direct ESXi connections and
               requires vCenter 6.0 or higher.
        """
        pulumi.set(__self__, "disks", disks)
        pulumi.set(__self__, "host_system_id", host_system_id)
        if custom_attributes is not None:
            pulumi.set(__self__, "custom_attributes", custom_attributes)
        if datastore_cluster_id is not None:
            pulumi.set(__self__, "datastore_cluster_id", datastore_cluster_id)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def disks(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The disks to use with the datastore.
        """
        return pulumi.get(self, "disks")

    @disks.setter
    def disks(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "disks", value)

    @property
    @pulumi.getter(name="hostSystemId")
    def host_system_id(self) -> pulumi.Input[str]:
        """
        The managed object ID of
        the host to set the datastore up on. Note that this is not necessarily the
        only host that the datastore will be set up on - see
        here for more info. Forces a
        new resource if changed.
        """
        return pulumi.get(self, "host_system_id")

    @host_system_id.setter
    def host_system_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "host_system_id", value)

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of custom attribute ids to attribute
        value string to set on datastore resource.

        > **NOTE:** Custom attributes are unsupported on direct ESXi connections
        and require vCenter.
        """
        return pulumi.get(self, "custom_attributes")

    @custom_attributes.setter
    def custom_attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_attributes", value)

    @property
    @pulumi.getter(name="datastoreClusterId")
    def datastore_cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed object
        ID of a datastore cluster to put this datastore in.
        Conflicts with `folder`.
        """
        return pulumi.get(self, "datastore_cluster_id")

    @datastore_cluster_id.setter
    def datastore_cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datastore_cluster_id", value)

    @property
    @pulumi.getter
    def folder(self) -> Optional[pulumi.Input[str]]:
        """
        The relative path to a folder to put this datastore in.
        This is a path relative to the datacenter you are deploying the datastore to.
        Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
        The provider will place a datastore named `test` in a datastore folder
        located at `/dc1/datastore/foo/bar`, with the final inventory path being
        `/dc1/datastore/foo/bar/test`. Conflicts with
        `datastore_cluster_id`.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the datastore. Forces a new resource if
        changed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IDs of any tags to attach to this resource.

        > **NOTE:** Tagging support is unsupported on direct ESXi connections and
        requires vCenter 6.0 or higher.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _VmfsDatastoreState:
    def __init__(__self__, *,
                 accessible: Optional[pulumi.Input[bool]] = None,
                 capacity: Optional[pulumi.Input[int]] = None,
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 datastore_cluster_id: Optional[pulumi.Input[str]] = None,
                 disks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 free_space: Optional[pulumi.Input[int]] = None,
                 host_system_id: Optional[pulumi.Input[str]] = None,
                 maintenance_mode: Optional[pulumi.Input[str]] = None,
                 multiple_host_access: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 uncommitted_space: Optional[pulumi.Input[int]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VmfsDatastore resources.
        :param pulumi.Input[bool] accessible: The connectivity status of the datastore. If this is `false`,
               some other computed attributes may be out of date.
        :param pulumi.Input[int] capacity: Maximum capacity of the datastore, in megabytes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: Map of custom attribute ids to attribute
               value string to set on datastore resource.
               
               > **NOTE:** Custom attributes are unsupported on direct ESXi connections
               and require vCenter.
        :param pulumi.Input[str] datastore_cluster_id: The managed object
               ID of a datastore cluster to put this datastore in.
               Conflicts with `folder`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disks: The disks to use with the datastore.
        :param pulumi.Input[str] folder: The relative path to a folder to put this datastore in.
               This is a path relative to the datacenter you are deploying the datastore to.
               Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
               The provider will place a datastore named `test` in a datastore folder
               located at `/dc1/datastore/foo/bar`, with the final inventory path being
               `/dc1/datastore/foo/bar/test`. Conflicts with
               `datastore_cluster_id`.
        :param pulumi.Input[int] free_space: Available space of this datastore, in megabytes.
        :param pulumi.Input[str] host_system_id: The managed object ID of
               the host to set the datastore up on. Note that this is not necessarily the
               only host that the datastore will be set up on - see
               here for more info. Forces a
               new resource if changed.
        :param pulumi.Input[str] maintenance_mode: The current maintenance mode state of the datastore.
        :param pulumi.Input[bool] multiple_host_access: If `true`, more than one host in the datacenter has
               been configured with access to the datastore.
        :param pulumi.Input[str] name: The name of the datastore. Forces a new resource if
               changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The IDs of any tags to attach to this resource.
               
               > **NOTE:** Tagging support is unsupported on direct ESXi connections and
               requires vCenter 6.0 or higher.
        :param pulumi.Input[int] uncommitted_space: Total additional storage space, in megabytes,
               potentially used by all virtual machines on this datastore.
        :param pulumi.Input[str] url: The unique locator for the datastore.
        """
        if accessible is not None:
            pulumi.set(__self__, "accessible", accessible)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if custom_attributes is not None:
            pulumi.set(__self__, "custom_attributes", custom_attributes)
        if datastore_cluster_id is not None:
            pulumi.set(__self__, "datastore_cluster_id", datastore_cluster_id)
        if disks is not None:
            pulumi.set(__self__, "disks", disks)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)
        if free_space is not None:
            pulumi.set(__self__, "free_space", free_space)
        if host_system_id is not None:
            pulumi.set(__self__, "host_system_id", host_system_id)
        if maintenance_mode is not None:
            pulumi.set(__self__, "maintenance_mode", maintenance_mode)
        if multiple_host_access is not None:
            pulumi.set(__self__, "multiple_host_access", multiple_host_access)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if uncommitted_space is not None:
            pulumi.set(__self__, "uncommitted_space", uncommitted_space)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def accessible(self) -> Optional[pulumi.Input[bool]]:
        """
        The connectivity status of the datastore. If this is `false`,
        some other computed attributes may be out of date.
        """
        return pulumi.get(self, "accessible")

    @accessible.setter
    def accessible(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "accessible", value)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum capacity of the datastore, in megabytes.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of custom attribute ids to attribute
        value string to set on datastore resource.

        > **NOTE:** Custom attributes are unsupported on direct ESXi connections
        and require vCenter.
        """
        return pulumi.get(self, "custom_attributes")

    @custom_attributes.setter
    def custom_attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_attributes", value)

    @property
    @pulumi.getter(name="datastoreClusterId")
    def datastore_cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed object
        ID of a datastore cluster to put this datastore in.
        Conflicts with `folder`.
        """
        return pulumi.get(self, "datastore_cluster_id")

    @datastore_cluster_id.setter
    def datastore_cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datastore_cluster_id", value)

    @property
    @pulumi.getter
    def disks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The disks to use with the datastore.
        """
        return pulumi.get(self, "disks")

    @disks.setter
    def disks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "disks", value)

    @property
    @pulumi.getter
    def folder(self) -> Optional[pulumi.Input[str]]:
        """
        The relative path to a folder to put this datastore in.
        This is a path relative to the datacenter you are deploying the datastore to.
        Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
        The provider will place a datastore named `test` in a datastore folder
        located at `/dc1/datastore/foo/bar`, with the final inventory path being
        `/dc1/datastore/foo/bar/test`. Conflicts with
        `datastore_cluster_id`.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter(name="freeSpace")
    def free_space(self) -> Optional[pulumi.Input[int]]:
        """
        Available space of this datastore, in megabytes.
        """
        return pulumi.get(self, "free_space")

    @free_space.setter
    def free_space(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "free_space", value)

    @property
    @pulumi.getter(name="hostSystemId")
    def host_system_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed object ID of
        the host to set the datastore up on. Note that this is not necessarily the
        only host that the datastore will be set up on - see
        here for more info. Forces a
        new resource if changed.
        """
        return pulumi.get(self, "host_system_id")

    @host_system_id.setter
    def host_system_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_system_id", value)

    @property
    @pulumi.getter(name="maintenanceMode")
    def maintenance_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The current maintenance mode state of the datastore.
        """
        return pulumi.get(self, "maintenance_mode")

    @maintenance_mode.setter
    def maintenance_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "maintenance_mode", value)

    @property
    @pulumi.getter(name="multipleHostAccess")
    def multiple_host_access(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, more than one host in the datacenter has
        been configured with access to the datastore.
        """
        return pulumi.get(self, "multiple_host_access")

    @multiple_host_access.setter
    def multiple_host_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multiple_host_access", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the datastore. Forces a new resource if
        changed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IDs of any tags to attach to this resource.

        > **NOTE:** Tagging support is unsupported on direct ESXi connections and
        requires vCenter 6.0 or higher.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="uncommittedSpace")
    def uncommitted_space(self) -> Optional[pulumi.Input[int]]:
        """
        Total additional storage space, in megabytes,
        potentially used by all virtual machines on this datastore.
        """
        return pulumi.get(self, "uncommitted_space")

    @uncommitted_space.setter
    def uncommitted_space(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "uncommitted_space", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The unique locator for the datastore.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class VmfsDatastore(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 datastore_cluster_id: Optional[pulumi.Input[str]] = None,
                 disks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 host_system_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a VmfsDatastore resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: Map of custom attribute ids to attribute
               value string to set on datastore resource.
               
               > **NOTE:** Custom attributes are unsupported on direct ESXi connections
               and require vCenter.
        :param pulumi.Input[str] datastore_cluster_id: The managed object
               ID of a datastore cluster to put this datastore in.
               Conflicts with `folder`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disks: The disks to use with the datastore.
        :param pulumi.Input[str] folder: The relative path to a folder to put this datastore in.
               This is a path relative to the datacenter you are deploying the datastore to.
               Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
               The provider will place a datastore named `test` in a datastore folder
               located at `/dc1/datastore/foo/bar`, with the final inventory path being
               `/dc1/datastore/foo/bar/test`. Conflicts with
               `datastore_cluster_id`.
        :param pulumi.Input[str] host_system_id: The managed object ID of
               the host to set the datastore up on. Note that this is not necessarily the
               only host that the datastore will be set up on - see
               here for more info. Forces a
               new resource if changed.
        :param pulumi.Input[str] name: The name of the datastore. Forces a new resource if
               changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The IDs of any tags to attach to this resource.
               
               > **NOTE:** Tagging support is unsupported on direct ESXi connections and
               requires vCenter 6.0 or higher.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VmfsDatastoreArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VmfsDatastore resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VmfsDatastoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VmfsDatastoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 datastore_cluster_id: Optional[pulumi.Input[str]] = None,
                 disks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 host_system_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VmfsDatastoreArgs.__new__(VmfsDatastoreArgs)

            __props__.__dict__["custom_attributes"] = custom_attributes
            __props__.__dict__["datastore_cluster_id"] = datastore_cluster_id
            if disks is None and not opts.urn:
                raise TypeError("Missing required property 'disks'")
            __props__.__dict__["disks"] = disks
            __props__.__dict__["folder"] = folder
            if host_system_id is None and not opts.urn:
                raise TypeError("Missing required property 'host_system_id'")
            __props__.__dict__["host_system_id"] = host_system_id
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["accessible"] = None
            __props__.__dict__["capacity"] = None
            __props__.__dict__["free_space"] = None
            __props__.__dict__["maintenance_mode"] = None
            __props__.__dict__["multiple_host_access"] = None
            __props__.__dict__["uncommitted_space"] = None
            __props__.__dict__["url"] = None
        super(VmfsDatastore, __self__).__init__(
            'vsphere:index/vmfsDatastore:VmfsDatastore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accessible: Optional[pulumi.Input[bool]] = None,
            capacity: Optional[pulumi.Input[int]] = None,
            custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            datastore_cluster_id: Optional[pulumi.Input[str]] = None,
            disks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            folder: Optional[pulumi.Input[str]] = None,
            free_space: Optional[pulumi.Input[int]] = None,
            host_system_id: Optional[pulumi.Input[str]] = None,
            maintenance_mode: Optional[pulumi.Input[str]] = None,
            multiple_host_access: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            uncommitted_space: Optional[pulumi.Input[int]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'VmfsDatastore':
        """
        Get an existing VmfsDatastore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] accessible: The connectivity status of the datastore. If this is `false`,
               some other computed attributes may be out of date.
        :param pulumi.Input[int] capacity: Maximum capacity of the datastore, in megabytes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: Map of custom attribute ids to attribute
               value string to set on datastore resource.
               
               > **NOTE:** Custom attributes are unsupported on direct ESXi connections
               and require vCenter.
        :param pulumi.Input[str] datastore_cluster_id: The managed object
               ID of a datastore cluster to put this datastore in.
               Conflicts with `folder`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disks: The disks to use with the datastore.
        :param pulumi.Input[str] folder: The relative path to a folder to put this datastore in.
               This is a path relative to the datacenter you are deploying the datastore to.
               Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
               The provider will place a datastore named `test` in a datastore folder
               located at `/dc1/datastore/foo/bar`, with the final inventory path being
               `/dc1/datastore/foo/bar/test`. Conflicts with
               `datastore_cluster_id`.
        :param pulumi.Input[int] free_space: Available space of this datastore, in megabytes.
        :param pulumi.Input[str] host_system_id: The managed object ID of
               the host to set the datastore up on. Note that this is not necessarily the
               only host that the datastore will be set up on - see
               here for more info. Forces a
               new resource if changed.
        :param pulumi.Input[str] maintenance_mode: The current maintenance mode state of the datastore.
        :param pulumi.Input[bool] multiple_host_access: If `true`, more than one host in the datacenter has
               been configured with access to the datastore.
        :param pulumi.Input[str] name: The name of the datastore. Forces a new resource if
               changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The IDs of any tags to attach to this resource.
               
               > **NOTE:** Tagging support is unsupported on direct ESXi connections and
               requires vCenter 6.0 or higher.
        :param pulumi.Input[int] uncommitted_space: Total additional storage space, in megabytes,
               potentially used by all virtual machines on this datastore.
        :param pulumi.Input[str] url: The unique locator for the datastore.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VmfsDatastoreState.__new__(_VmfsDatastoreState)

        __props__.__dict__["accessible"] = accessible
        __props__.__dict__["capacity"] = capacity
        __props__.__dict__["custom_attributes"] = custom_attributes
        __props__.__dict__["datastore_cluster_id"] = datastore_cluster_id
        __props__.__dict__["disks"] = disks
        __props__.__dict__["folder"] = folder
        __props__.__dict__["free_space"] = free_space
        __props__.__dict__["host_system_id"] = host_system_id
        __props__.__dict__["maintenance_mode"] = maintenance_mode
        __props__.__dict__["multiple_host_access"] = multiple_host_access
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["uncommitted_space"] = uncommitted_space
        __props__.__dict__["url"] = url
        return VmfsDatastore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accessible(self) -> pulumi.Output[bool]:
        """
        The connectivity status of the datastore. If this is `false`,
        some other computed attributes may be out of date.
        """
        return pulumi.get(self, "accessible")

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Output[int]:
        """
        Maximum capacity of the datastore, in megabytes.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of custom attribute ids to attribute
        value string to set on datastore resource.

        > **NOTE:** Custom attributes are unsupported on direct ESXi connections
        and require vCenter.
        """
        return pulumi.get(self, "custom_attributes")

    @property
    @pulumi.getter(name="datastoreClusterId")
    def datastore_cluster_id(self) -> pulumi.Output[Optional[str]]:
        """
        The managed object
        ID of a datastore cluster to put this datastore in.
        Conflicts with `folder`.
        """
        return pulumi.get(self, "datastore_cluster_id")

    @property
    @pulumi.getter
    def disks(self) -> pulumi.Output[Sequence[str]]:
        """
        The disks to use with the datastore.
        """
        return pulumi.get(self, "disks")

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Output[Optional[str]]:
        """
        The relative path to a folder to put this datastore in.
        This is a path relative to the datacenter you are deploying the datastore to.
        Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
        The provider will place a datastore named `test` in a datastore folder
        located at `/dc1/datastore/foo/bar`, with the final inventory path being
        `/dc1/datastore/foo/bar/test`. Conflicts with
        `datastore_cluster_id`.
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter(name="freeSpace")
    def free_space(self) -> pulumi.Output[int]:
        """
        Available space of this datastore, in megabytes.
        """
        return pulumi.get(self, "free_space")

    @property
    @pulumi.getter(name="hostSystemId")
    def host_system_id(self) -> pulumi.Output[str]:
        """
        The managed object ID of
        the host to set the datastore up on. Note that this is not necessarily the
        only host that the datastore will be set up on - see
        here for more info. Forces a
        new resource if changed.
        """
        return pulumi.get(self, "host_system_id")

    @property
    @pulumi.getter(name="maintenanceMode")
    def maintenance_mode(self) -> pulumi.Output[str]:
        """
        The current maintenance mode state of the datastore.
        """
        return pulumi.get(self, "maintenance_mode")

    @property
    @pulumi.getter(name="multipleHostAccess")
    def multiple_host_access(self) -> pulumi.Output[bool]:
        """
        If `true`, more than one host in the datacenter has
        been configured with access to the datastore.
        """
        return pulumi.get(self, "multiple_host_access")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the datastore. Forces a new resource if
        changed.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IDs of any tags to attach to this resource.

        > **NOTE:** Tagging support is unsupported on direct ESXi connections and
        requires vCenter 6.0 or higher.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="uncommittedSpace")
    def uncommitted_space(self) -> pulumi.Output[int]:
        """
        Total additional storage space, in megabytes,
        potentially used by all virtual machines on this datastore.
        """
        return pulumi.get(self, "uncommitted_space")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The unique locator for the datastore.
        """
        return pulumi.get(self, "url")

