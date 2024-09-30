# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VappEntityArgs', 'VappEntity']

@pulumi.input_type
class VappEntityArgs:
    def __init__(__self__, *,
                 container_id: pulumi.Input[str],
                 target_id: pulumi.Input[str],
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 start_action: Optional[pulumi.Input[str]] = None,
                 start_delay: Optional[pulumi.Input[int]] = None,
                 start_order: Optional[pulumi.Input[int]] = None,
                 stop_action: Optional[pulumi.Input[str]] = None,
                 stop_delay: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 wait_for_guest: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a VappEntity resource.
        :param pulumi.Input[str] container_id: Managed object ID of the vApp
               container the entity is a member of.
        :param pulumi.Input[str] target_id: Managed object ID of the entity
               to power on or power off. This can be a virtual machine or a vApp.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: A list of custom attributes to set on this resource.
        :param pulumi.Input[str] start_action: How to start the entity. Valid settings are none
               or powerOn. If set to none, then the entity does not participate in auto-start.
               Default: powerOn
        :param pulumi.Input[int] start_delay: Delay in seconds before continuing with the next
               entity in the order of entities to be started. Default: 120
        :param pulumi.Input[int] start_order: Order to start and stop target in vApp. Default: 1
        :param pulumi.Input[str] stop_action: Defines the stop action for the entity. Can be set
               to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
               does not participate in auto-stop. Default: powerOff
        :param pulumi.Input[int] stop_delay: Delay in seconds before continuing with the next
               entity in the order sequence. This is only used if the stopAction is
               guestShutdown. Default: 120
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tag IDs to apply to this object.
        :param pulumi.Input[bool] wait_for_guest: Determines if the VM should be marked as being
               started when VMware Tools are ready instead of waiting for `start_delay`. This
               property has no effect for vApps. Default: false
        """
        pulumi.set(__self__, "container_id", container_id)
        pulumi.set(__self__, "target_id", target_id)
        if custom_attributes is not None:
            pulumi.set(__self__, "custom_attributes", custom_attributes)
        if start_action is not None:
            pulumi.set(__self__, "start_action", start_action)
        if start_delay is not None:
            pulumi.set(__self__, "start_delay", start_delay)
        if start_order is not None:
            pulumi.set(__self__, "start_order", start_order)
        if stop_action is not None:
            pulumi.set(__self__, "stop_action", stop_action)
        if stop_delay is not None:
            pulumi.set(__self__, "stop_delay", stop_delay)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if wait_for_guest is not None:
            pulumi.set(__self__, "wait_for_guest", wait_for_guest)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> pulumi.Input[str]:
        """
        Managed object ID of the vApp
        container the entity is a member of.
        """
        return pulumi.get(self, "container_id")

    @container_id.setter
    def container_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_id", value)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> pulumi.Input[str]:
        """
        Managed object ID of the entity
        to power on or power off. This can be a virtual machine or a vApp.
        """
        return pulumi.get(self, "target_id")

    @target_id.setter
    def target_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_id", value)

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A list of custom attributes to set on this resource.
        """
        return pulumi.get(self, "custom_attributes")

    @custom_attributes.setter
    def custom_attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_attributes", value)

    @property
    @pulumi.getter(name="startAction")
    def start_action(self) -> Optional[pulumi.Input[str]]:
        """
        How to start the entity. Valid settings are none
        or powerOn. If set to none, then the entity does not participate in auto-start.
        Default: powerOn
        """
        return pulumi.get(self, "start_action")

    @start_action.setter
    def start_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_action", value)

    @property
    @pulumi.getter(name="startDelay")
    def start_delay(self) -> Optional[pulumi.Input[int]]:
        """
        Delay in seconds before continuing with the next
        entity in the order of entities to be started. Default: 120
        """
        return pulumi.get(self, "start_delay")

    @start_delay.setter
    def start_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_delay", value)

    @property
    @pulumi.getter(name="startOrder")
    def start_order(self) -> Optional[pulumi.Input[int]]:
        """
        Order to start and stop target in vApp. Default: 1
        """
        return pulumi.get(self, "start_order")

    @start_order.setter
    def start_order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_order", value)

    @property
    @pulumi.getter(name="stopAction")
    def stop_action(self) -> Optional[pulumi.Input[str]]:
        """
        Defines the stop action for the entity. Can be set
        to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
        does not participate in auto-stop. Default: powerOff
        """
        return pulumi.get(self, "stop_action")

    @stop_action.setter
    def stop_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stop_action", value)

    @property
    @pulumi.getter(name="stopDelay")
    def stop_delay(self) -> Optional[pulumi.Input[int]]:
        """
        Delay in seconds before continuing with the next
        entity in the order sequence. This is only used if the stopAction is
        guestShutdown. Default: 120
        """
        return pulumi.get(self, "stop_delay")

    @stop_delay.setter
    def stop_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "stop_delay", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of tag IDs to apply to this object.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="waitForGuest")
    def wait_for_guest(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if the VM should be marked as being
        started when VMware Tools are ready instead of waiting for `start_delay`. This
        property has no effect for vApps. Default: false
        """
        return pulumi.get(self, "wait_for_guest")

    @wait_for_guest.setter
    def wait_for_guest(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_for_guest", value)


@pulumi.input_type
class _VappEntityState:
    def __init__(__self__, *,
                 container_id: Optional[pulumi.Input[str]] = None,
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 start_action: Optional[pulumi.Input[str]] = None,
                 start_delay: Optional[pulumi.Input[int]] = None,
                 start_order: Optional[pulumi.Input[int]] = None,
                 stop_action: Optional[pulumi.Input[str]] = None,
                 stop_delay: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 wait_for_guest: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering VappEntity resources.
        :param pulumi.Input[str] container_id: Managed object ID of the vApp
               container the entity is a member of.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: A list of custom attributes to set on this resource.
        :param pulumi.Input[str] start_action: How to start the entity. Valid settings are none
               or powerOn. If set to none, then the entity does not participate in auto-start.
               Default: powerOn
        :param pulumi.Input[int] start_delay: Delay in seconds before continuing with the next
               entity in the order of entities to be started. Default: 120
        :param pulumi.Input[int] start_order: Order to start and stop target in vApp. Default: 1
        :param pulumi.Input[str] stop_action: Defines the stop action for the entity. Can be set
               to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
               does not participate in auto-stop. Default: powerOff
        :param pulumi.Input[int] stop_delay: Delay in seconds before continuing with the next
               entity in the order sequence. This is only used if the stopAction is
               guestShutdown. Default: 120
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tag IDs to apply to this object.
        :param pulumi.Input[str] target_id: Managed object ID of the entity
               to power on or power off. This can be a virtual machine or a vApp.
        :param pulumi.Input[bool] wait_for_guest: Determines if the VM should be marked as being
               started when VMware Tools are ready instead of waiting for `start_delay`. This
               property has no effect for vApps. Default: false
        """
        if container_id is not None:
            pulumi.set(__self__, "container_id", container_id)
        if custom_attributes is not None:
            pulumi.set(__self__, "custom_attributes", custom_attributes)
        if start_action is not None:
            pulumi.set(__self__, "start_action", start_action)
        if start_delay is not None:
            pulumi.set(__self__, "start_delay", start_delay)
        if start_order is not None:
            pulumi.set(__self__, "start_order", start_order)
        if stop_action is not None:
            pulumi.set(__self__, "stop_action", stop_action)
        if stop_delay is not None:
            pulumi.set(__self__, "stop_delay", stop_delay)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if target_id is not None:
            pulumi.set(__self__, "target_id", target_id)
        if wait_for_guest is not None:
            pulumi.set(__self__, "wait_for_guest", wait_for_guest)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> Optional[pulumi.Input[str]]:
        """
        Managed object ID of the vApp
        container the entity is a member of.
        """
        return pulumi.get(self, "container_id")

    @container_id.setter
    def container_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_id", value)

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A list of custom attributes to set on this resource.
        """
        return pulumi.get(self, "custom_attributes")

    @custom_attributes.setter
    def custom_attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_attributes", value)

    @property
    @pulumi.getter(name="startAction")
    def start_action(self) -> Optional[pulumi.Input[str]]:
        """
        How to start the entity. Valid settings are none
        or powerOn. If set to none, then the entity does not participate in auto-start.
        Default: powerOn
        """
        return pulumi.get(self, "start_action")

    @start_action.setter
    def start_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_action", value)

    @property
    @pulumi.getter(name="startDelay")
    def start_delay(self) -> Optional[pulumi.Input[int]]:
        """
        Delay in seconds before continuing with the next
        entity in the order of entities to be started. Default: 120
        """
        return pulumi.get(self, "start_delay")

    @start_delay.setter
    def start_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_delay", value)

    @property
    @pulumi.getter(name="startOrder")
    def start_order(self) -> Optional[pulumi.Input[int]]:
        """
        Order to start and stop target in vApp. Default: 1
        """
        return pulumi.get(self, "start_order")

    @start_order.setter
    def start_order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_order", value)

    @property
    @pulumi.getter(name="stopAction")
    def stop_action(self) -> Optional[pulumi.Input[str]]:
        """
        Defines the stop action for the entity. Can be set
        to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
        does not participate in auto-stop. Default: powerOff
        """
        return pulumi.get(self, "stop_action")

    @stop_action.setter
    def stop_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stop_action", value)

    @property
    @pulumi.getter(name="stopDelay")
    def stop_delay(self) -> Optional[pulumi.Input[int]]:
        """
        Delay in seconds before continuing with the next
        entity in the order sequence. This is only used if the stopAction is
        guestShutdown. Default: 120
        """
        return pulumi.get(self, "stop_delay")

    @stop_delay.setter
    def stop_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "stop_delay", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of tag IDs to apply to this object.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> Optional[pulumi.Input[str]]:
        """
        Managed object ID of the entity
        to power on or power off. This can be a virtual machine or a vApp.
        """
        return pulumi.get(self, "target_id")

    @target_id.setter
    def target_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_id", value)

    @property
    @pulumi.getter(name="waitForGuest")
    def wait_for_guest(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if the VM should be marked as being
        started when VMware Tools are ready instead of waiting for `start_delay`. This
        property has no effect for vApps. Default: false
        """
        return pulumi.get(self, "wait_for_guest")

    @wait_for_guest.setter
    def wait_for_guest(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_for_guest", value)


class VappEntity(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_id: Optional[pulumi.Input[str]] = None,
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 start_action: Optional[pulumi.Input[str]] = None,
                 start_delay: Optional[pulumi.Input[int]] = None,
                 start_order: Optional[pulumi.Input[int]] = None,
                 stop_action: Optional[pulumi.Input[str]] = None,
                 stop_delay: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 wait_for_guest: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        The `VappEntity` resource can be used to describe the behavior of an
        entity (virtual machine or sub-vApp container) in a vApp container.

        For more information on vSphere vApps, see [this
        page][ref-vsphere-vapp].

        [ref-vsphere-vapp]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vm-administration/GUID-2A95EBB8-1779-40FA-B4FB-4D0845750879.html

        ## Example Usage

        The basic example below sets up a vApp container and a virtual machine in a
        compute cluster and then creates a vApp entity to change the virtual machine's
        power on behavior in the vApp container.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        config = pulumi.Config()
        datacenter = config.get("datacenter")
        if datacenter is None:
            datacenter = "dc-01"
        cluster = config.get("cluster")
        if cluster is None:
            cluster = "cluster-01"
        datacenter_get_datacenter = vsphere.get_datacenter(name=datacenter)
        compute_cluster = vsphere.get_compute_cluster(name=cluster,
            datacenter_id=datacenter_get_datacenter.id)
        network = vsphere.get_network(name="network1",
            datacenter_id=datacenter_get_datacenter.id)
        datastore = vsphere.get_datastore(name="datastore1",
            datacenter_id=datacenter_get_datacenter.id)
        vapp_container = vsphere.VappContainer("vapp_container",
            name="vapp-container-test",
            parent_resource_pool_id=compute_cluster.id)
        vm = vsphere.VirtualMachine("vm",
            name="virtual-machine-test",
            resource_pool_id=vapp_container.id,
            datastore_id=datastore.id,
            num_cpus=2,
            memory=1024,
            guest_id="ubuntu64Guest",
            disks=[{
                "label": "disk0",
                "size": 1,
            }],
            network_interfaces=[{
                "network_id": network.id,
            }])
        vapp_entity = vsphere.VappEntity("vapp_entity",
            target_id=vm.moid,
            container_id=vapp_container.id,
            start_action="none")
        ```

        ## Import

        An existing vApp entity can be imported into this resource via

        the ID of the vApp Entity.

        ```sh
        $ pulumi import vsphere:index/vappEntity:VappEntity vapp_entity vm-123:res-456
        ```

        The above would import the vApp entity that governs the behavior of the virtual

        machine with a [managed object ID][docs-about-morefs] of vm-123 in the vApp

        container with the [managed object ID][docs-about-morefs] res-456.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_id: Managed object ID of the vApp
               container the entity is a member of.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: A list of custom attributes to set on this resource.
        :param pulumi.Input[str] start_action: How to start the entity. Valid settings are none
               or powerOn. If set to none, then the entity does not participate in auto-start.
               Default: powerOn
        :param pulumi.Input[int] start_delay: Delay in seconds before continuing with the next
               entity in the order of entities to be started. Default: 120
        :param pulumi.Input[int] start_order: Order to start and stop target in vApp. Default: 1
        :param pulumi.Input[str] stop_action: Defines the stop action for the entity. Can be set
               to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
               does not participate in auto-stop. Default: powerOff
        :param pulumi.Input[int] stop_delay: Delay in seconds before continuing with the next
               entity in the order sequence. This is only used if the stopAction is
               guestShutdown. Default: 120
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tag IDs to apply to this object.
        :param pulumi.Input[str] target_id: Managed object ID of the entity
               to power on or power off. This can be a virtual machine or a vApp.
        :param pulumi.Input[bool] wait_for_guest: Determines if the VM should be marked as being
               started when VMware Tools are ready instead of waiting for `start_delay`. This
               property has no effect for vApps. Default: false
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VappEntityArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `VappEntity` resource can be used to describe the behavior of an
        entity (virtual machine or sub-vApp container) in a vApp container.

        For more information on vSphere vApps, see [this
        page][ref-vsphere-vapp].

        [ref-vsphere-vapp]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vm-administration/GUID-2A95EBB8-1779-40FA-B4FB-4D0845750879.html

        ## Example Usage

        The basic example below sets up a vApp container and a virtual machine in a
        compute cluster and then creates a vApp entity to change the virtual machine's
        power on behavior in the vApp container.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        config = pulumi.Config()
        datacenter = config.get("datacenter")
        if datacenter is None:
            datacenter = "dc-01"
        cluster = config.get("cluster")
        if cluster is None:
            cluster = "cluster-01"
        datacenter_get_datacenter = vsphere.get_datacenter(name=datacenter)
        compute_cluster = vsphere.get_compute_cluster(name=cluster,
            datacenter_id=datacenter_get_datacenter.id)
        network = vsphere.get_network(name="network1",
            datacenter_id=datacenter_get_datacenter.id)
        datastore = vsphere.get_datastore(name="datastore1",
            datacenter_id=datacenter_get_datacenter.id)
        vapp_container = vsphere.VappContainer("vapp_container",
            name="vapp-container-test",
            parent_resource_pool_id=compute_cluster.id)
        vm = vsphere.VirtualMachine("vm",
            name="virtual-machine-test",
            resource_pool_id=vapp_container.id,
            datastore_id=datastore.id,
            num_cpus=2,
            memory=1024,
            guest_id="ubuntu64Guest",
            disks=[{
                "label": "disk0",
                "size": 1,
            }],
            network_interfaces=[{
                "network_id": network.id,
            }])
        vapp_entity = vsphere.VappEntity("vapp_entity",
            target_id=vm.moid,
            container_id=vapp_container.id,
            start_action="none")
        ```

        ## Import

        An existing vApp entity can be imported into this resource via

        the ID of the vApp Entity.

        ```sh
        $ pulumi import vsphere:index/vappEntity:VappEntity vapp_entity vm-123:res-456
        ```

        The above would import the vApp entity that governs the behavior of the virtual

        machine with a [managed object ID][docs-about-morefs] of vm-123 in the vApp

        container with the [managed object ID][docs-about-morefs] res-456.

        :param str resource_name: The name of the resource.
        :param VappEntityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VappEntityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_id: Optional[pulumi.Input[str]] = None,
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 start_action: Optional[pulumi.Input[str]] = None,
                 start_delay: Optional[pulumi.Input[int]] = None,
                 start_order: Optional[pulumi.Input[int]] = None,
                 stop_action: Optional[pulumi.Input[str]] = None,
                 stop_delay: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 wait_for_guest: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VappEntityArgs.__new__(VappEntityArgs)

            if container_id is None and not opts.urn:
                raise TypeError("Missing required property 'container_id'")
            __props__.__dict__["container_id"] = container_id
            __props__.__dict__["custom_attributes"] = custom_attributes
            __props__.__dict__["start_action"] = start_action
            __props__.__dict__["start_delay"] = start_delay
            __props__.__dict__["start_order"] = start_order
            __props__.__dict__["stop_action"] = stop_action
            __props__.__dict__["stop_delay"] = stop_delay
            __props__.__dict__["tags"] = tags
            if target_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_id'")
            __props__.__dict__["target_id"] = target_id
            __props__.__dict__["wait_for_guest"] = wait_for_guest
        super(VappEntity, __self__).__init__(
            'vsphere:index/vappEntity:VappEntity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            container_id: Optional[pulumi.Input[str]] = None,
            custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            start_action: Optional[pulumi.Input[str]] = None,
            start_delay: Optional[pulumi.Input[int]] = None,
            start_order: Optional[pulumi.Input[int]] = None,
            stop_action: Optional[pulumi.Input[str]] = None,
            stop_delay: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            target_id: Optional[pulumi.Input[str]] = None,
            wait_for_guest: Optional[pulumi.Input[bool]] = None) -> 'VappEntity':
        """
        Get an existing VappEntity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_id: Managed object ID of the vApp
               container the entity is a member of.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: A list of custom attributes to set on this resource.
        :param pulumi.Input[str] start_action: How to start the entity. Valid settings are none
               or powerOn. If set to none, then the entity does not participate in auto-start.
               Default: powerOn
        :param pulumi.Input[int] start_delay: Delay in seconds before continuing with the next
               entity in the order of entities to be started. Default: 120
        :param pulumi.Input[int] start_order: Order to start and stop target in vApp. Default: 1
        :param pulumi.Input[str] stop_action: Defines the stop action for the entity. Can be set
               to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
               does not participate in auto-stop. Default: powerOff
        :param pulumi.Input[int] stop_delay: Delay in seconds before continuing with the next
               entity in the order sequence. This is only used if the stopAction is
               guestShutdown. Default: 120
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tag IDs to apply to this object.
        :param pulumi.Input[str] target_id: Managed object ID of the entity
               to power on or power off. This can be a virtual machine or a vApp.
        :param pulumi.Input[bool] wait_for_guest: Determines if the VM should be marked as being
               started when VMware Tools are ready instead of waiting for `start_delay`. This
               property has no effect for vApps. Default: false
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VappEntityState.__new__(_VappEntityState)

        __props__.__dict__["container_id"] = container_id
        __props__.__dict__["custom_attributes"] = custom_attributes
        __props__.__dict__["start_action"] = start_action
        __props__.__dict__["start_delay"] = start_delay
        __props__.__dict__["start_order"] = start_order
        __props__.__dict__["stop_action"] = stop_action
        __props__.__dict__["stop_delay"] = stop_delay
        __props__.__dict__["tags"] = tags
        __props__.__dict__["target_id"] = target_id
        __props__.__dict__["wait_for_guest"] = wait_for_guest
        return VappEntity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> pulumi.Output[str]:
        """
        Managed object ID of the vApp
        container the entity is a member of.
        """
        return pulumi.get(self, "container_id")

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A list of custom attributes to set on this resource.
        """
        return pulumi.get(self, "custom_attributes")

    @property
    @pulumi.getter(name="startAction")
    def start_action(self) -> pulumi.Output[Optional[str]]:
        """
        How to start the entity. Valid settings are none
        or powerOn. If set to none, then the entity does not participate in auto-start.
        Default: powerOn
        """
        return pulumi.get(self, "start_action")

    @property
    @pulumi.getter(name="startDelay")
    def start_delay(self) -> pulumi.Output[Optional[int]]:
        """
        Delay in seconds before continuing with the next
        entity in the order of entities to be started. Default: 120
        """
        return pulumi.get(self, "start_delay")

    @property
    @pulumi.getter(name="startOrder")
    def start_order(self) -> pulumi.Output[Optional[int]]:
        """
        Order to start and stop target in vApp. Default: 1
        """
        return pulumi.get(self, "start_order")

    @property
    @pulumi.getter(name="stopAction")
    def stop_action(self) -> pulumi.Output[Optional[str]]:
        """
        Defines the stop action for the entity. Can be set
        to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
        does not participate in auto-stop. Default: powerOff
        """
        return pulumi.get(self, "stop_action")

    @property
    @pulumi.getter(name="stopDelay")
    def stop_delay(self) -> pulumi.Output[Optional[int]]:
        """
        Delay in seconds before continuing with the next
        entity in the order sequence. This is only used if the stopAction is
        guestShutdown. Default: 120
        """
        return pulumi.get(self, "stop_delay")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of tag IDs to apply to this object.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> pulumi.Output[str]:
        """
        Managed object ID of the entity
        to power on or power off. This can be a virtual machine or a vApp.
        """
        return pulumi.get(self, "target_id")

    @property
    @pulumi.getter(name="waitForGuest")
    def wait_for_guest(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if the VM should be marked as being
        started when VMware Tools are ready instead of waiting for `start_delay`. This
        property has no effect for vApps. Default: false
        """
        return pulumi.get(self, "wait_for_guest")

