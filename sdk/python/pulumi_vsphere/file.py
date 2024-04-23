# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FileArgs', 'File']

@pulumi.input_type
class FileArgs:
    def __init__(__self__, *,
                 datastore: pulumi.Input[str],
                 destination_file: pulumi.Input[str],
                 source_file: pulumi.Input[str],
                 create_directories: Optional[pulumi.Input[bool]] = None,
                 datacenter: Optional[pulumi.Input[str]] = None,
                 source_datacenter: Optional[pulumi.Input[str]] = None,
                 source_datastore: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a File resource.
        :param pulumi.Input[str] datastore: The name of the datastore to which to upload the
               file.
        :param pulumi.Input[str] destination_file: The path to where the file should be uploaded
               or copied to on the destination `datastore` in vSphere.
        :param pulumi.Input[bool] create_directories: Create directories in `destination_file`
               path parameter on first apply if any are missing for copy operation.
               
               > **NOTE:** Any directory created as part of the `create_directories` argument
               will not be deleted when the resource is destroyed. New directories are not
               created if the `destination_file` path is changed in subsequent applies.
        :param pulumi.Input[str] datacenter: The name of a datacenter to which the file will be
               uploaded.
        :param pulumi.Input[str] source_datacenter: The name of a datacenter from which the file
               will be copied. Forces a new resource if changed.
        :param pulumi.Input[str] source_datastore: The name of the datastore from which file will
               be copied. Forces a new resource if changed.
        """
        pulumi.set(__self__, "datastore", datastore)
        pulumi.set(__self__, "destination_file", destination_file)
        pulumi.set(__self__, "source_file", source_file)
        if create_directories is not None:
            pulumi.set(__self__, "create_directories", create_directories)
        if datacenter is not None:
            pulumi.set(__self__, "datacenter", datacenter)
        if source_datacenter is not None:
            pulumi.set(__self__, "source_datacenter", source_datacenter)
        if source_datastore is not None:
            pulumi.set(__self__, "source_datastore", source_datastore)

    @property
    @pulumi.getter
    def datastore(self) -> pulumi.Input[str]:
        """
        The name of the datastore to which to upload the
        file.
        """
        return pulumi.get(self, "datastore")

    @datastore.setter
    def datastore(self, value: pulumi.Input[str]):
        pulumi.set(self, "datastore", value)

    @property
    @pulumi.getter(name="destinationFile")
    def destination_file(self) -> pulumi.Input[str]:
        """
        The path to where the file should be uploaded
        or copied to on the destination `datastore` in vSphere.
        """
        return pulumi.get(self, "destination_file")

    @destination_file.setter
    def destination_file(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_file", value)

    @property
    @pulumi.getter(name="sourceFile")
    def source_file(self) -> pulumi.Input[str]:
        return pulumi.get(self, "source_file")

    @source_file.setter
    def source_file(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_file", value)

    @property
    @pulumi.getter(name="createDirectories")
    def create_directories(self) -> Optional[pulumi.Input[bool]]:
        """
        Create directories in `destination_file`
        path parameter on first apply if any are missing for copy operation.

        > **NOTE:** Any directory created as part of the `create_directories` argument
        will not be deleted when the resource is destroyed. New directories are not
        created if the `destination_file` path is changed in subsequent applies.
        """
        return pulumi.get(self, "create_directories")

    @create_directories.setter
    def create_directories(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_directories", value)

    @property
    @pulumi.getter
    def datacenter(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a datacenter to which the file will be
        uploaded.
        """
        return pulumi.get(self, "datacenter")

    @datacenter.setter
    def datacenter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datacenter", value)

    @property
    @pulumi.getter(name="sourceDatacenter")
    def source_datacenter(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a datacenter from which the file
        will be copied. Forces a new resource if changed.
        """
        return pulumi.get(self, "source_datacenter")

    @source_datacenter.setter
    def source_datacenter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_datacenter", value)

    @property
    @pulumi.getter(name="sourceDatastore")
    def source_datastore(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the datastore from which file will
        be copied. Forces a new resource if changed.
        """
        return pulumi.get(self, "source_datastore")

    @source_datastore.setter
    def source_datastore(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_datastore", value)


@pulumi.input_type
class _FileState:
    def __init__(__self__, *,
                 create_directories: Optional[pulumi.Input[bool]] = None,
                 datacenter: Optional[pulumi.Input[str]] = None,
                 datastore: Optional[pulumi.Input[str]] = None,
                 destination_file: Optional[pulumi.Input[str]] = None,
                 source_datacenter: Optional[pulumi.Input[str]] = None,
                 source_datastore: Optional[pulumi.Input[str]] = None,
                 source_file: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering File resources.
        :param pulumi.Input[bool] create_directories: Create directories in `destination_file`
               path parameter on first apply if any are missing for copy operation.
               
               > **NOTE:** Any directory created as part of the `create_directories` argument
               will not be deleted when the resource is destroyed. New directories are not
               created if the `destination_file` path is changed in subsequent applies.
        :param pulumi.Input[str] datacenter: The name of a datacenter to which the file will be
               uploaded.
        :param pulumi.Input[str] datastore: The name of the datastore to which to upload the
               file.
        :param pulumi.Input[str] destination_file: The path to where the file should be uploaded
               or copied to on the destination `datastore` in vSphere.
        :param pulumi.Input[str] source_datacenter: The name of a datacenter from which the file
               will be copied. Forces a new resource if changed.
        :param pulumi.Input[str] source_datastore: The name of the datastore from which file will
               be copied. Forces a new resource if changed.
        """
        if create_directories is not None:
            pulumi.set(__self__, "create_directories", create_directories)
        if datacenter is not None:
            pulumi.set(__self__, "datacenter", datacenter)
        if datastore is not None:
            pulumi.set(__self__, "datastore", datastore)
        if destination_file is not None:
            pulumi.set(__self__, "destination_file", destination_file)
        if source_datacenter is not None:
            pulumi.set(__self__, "source_datacenter", source_datacenter)
        if source_datastore is not None:
            pulumi.set(__self__, "source_datastore", source_datastore)
        if source_file is not None:
            pulumi.set(__self__, "source_file", source_file)

    @property
    @pulumi.getter(name="createDirectories")
    def create_directories(self) -> Optional[pulumi.Input[bool]]:
        """
        Create directories in `destination_file`
        path parameter on first apply if any are missing for copy operation.

        > **NOTE:** Any directory created as part of the `create_directories` argument
        will not be deleted when the resource is destroyed. New directories are not
        created if the `destination_file` path is changed in subsequent applies.
        """
        return pulumi.get(self, "create_directories")

    @create_directories.setter
    def create_directories(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_directories", value)

    @property
    @pulumi.getter
    def datacenter(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a datacenter to which the file will be
        uploaded.
        """
        return pulumi.get(self, "datacenter")

    @datacenter.setter
    def datacenter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datacenter", value)

    @property
    @pulumi.getter
    def datastore(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the datastore to which to upload the
        file.
        """
        return pulumi.get(self, "datastore")

    @datastore.setter
    def datastore(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datastore", value)

    @property
    @pulumi.getter(name="destinationFile")
    def destination_file(self) -> Optional[pulumi.Input[str]]:
        """
        The path to where the file should be uploaded
        or copied to on the destination `datastore` in vSphere.
        """
        return pulumi.get(self, "destination_file")

    @destination_file.setter
    def destination_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_file", value)

    @property
    @pulumi.getter(name="sourceDatacenter")
    def source_datacenter(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a datacenter from which the file
        will be copied. Forces a new resource if changed.
        """
        return pulumi.get(self, "source_datacenter")

    @source_datacenter.setter
    def source_datacenter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_datacenter", value)

    @property
    @pulumi.getter(name="sourceDatastore")
    def source_datastore(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the datastore from which file will
        be copied. Forces a new resource if changed.
        """
        return pulumi.get(self, "source_datastore")

    @source_datastore.setter
    def source_datastore(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_datastore", value)

    @property
    @pulumi.getter(name="sourceFile")
    def source_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_file")

    @source_file.setter
    def source_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_file", value)


class File(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_directories: Optional[pulumi.Input[bool]] = None,
                 datacenter: Optional[pulumi.Input[str]] = None,
                 datastore: Optional[pulumi.Input[str]] = None,
                 destination_file: Optional[pulumi.Input[str]] = None,
                 source_datacenter: Optional[pulumi.Input[str]] = None,
                 source_datastore: Optional[pulumi.Input[str]] = None,
                 source_file: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ### Uploading a File

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        ubuntu_vmdk_upload = vsphere.File("ubuntu_vmdk_upload",
            datacenter="dc-01",
            datastore="datastore-01",
            source_file="/my/src/path/custom_ubuntu.vmdk",
            destination_file="/my/dst/path/custom_ubuntu.vmdk",
            create_directories=True)
        ```

        ### Copying a File

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        ubuntu_copy = vsphere.File("ubuntu_copy",
            source_datacenter="dc-01",
            datacenter="dc-01",
            source_datastore="datastore-01",
            datastore="datastore-01",
            source_file="/my/src/path/custom_ubuntu.vmdk",
            destination_file="/my/dst/path/custom_ubuntu.vmdk",
            create_directories=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] create_directories: Create directories in `destination_file`
               path parameter on first apply if any are missing for copy operation.
               
               > **NOTE:** Any directory created as part of the `create_directories` argument
               will not be deleted when the resource is destroyed. New directories are not
               created if the `destination_file` path is changed in subsequent applies.
        :param pulumi.Input[str] datacenter: The name of a datacenter to which the file will be
               uploaded.
        :param pulumi.Input[str] datastore: The name of the datastore to which to upload the
               file.
        :param pulumi.Input[str] destination_file: The path to where the file should be uploaded
               or copied to on the destination `datastore` in vSphere.
        :param pulumi.Input[str] source_datacenter: The name of a datacenter from which the file
               will be copied. Forces a new resource if changed.
        :param pulumi.Input[str] source_datastore: The name of the datastore from which file will
               be copied. Forces a new resource if changed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ### Uploading a File

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        ubuntu_vmdk_upload = vsphere.File("ubuntu_vmdk_upload",
            datacenter="dc-01",
            datastore="datastore-01",
            source_file="/my/src/path/custom_ubuntu.vmdk",
            destination_file="/my/dst/path/custom_ubuntu.vmdk",
            create_directories=True)
        ```

        ### Copying a File

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        ubuntu_copy = vsphere.File("ubuntu_copy",
            source_datacenter="dc-01",
            datacenter="dc-01",
            source_datastore="datastore-01",
            datastore="datastore-01",
            source_file="/my/src/path/custom_ubuntu.vmdk",
            destination_file="/my/dst/path/custom_ubuntu.vmdk",
            create_directories=True)
        ```

        :param str resource_name: The name of the resource.
        :param FileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_directories: Optional[pulumi.Input[bool]] = None,
                 datacenter: Optional[pulumi.Input[str]] = None,
                 datastore: Optional[pulumi.Input[str]] = None,
                 destination_file: Optional[pulumi.Input[str]] = None,
                 source_datacenter: Optional[pulumi.Input[str]] = None,
                 source_datastore: Optional[pulumi.Input[str]] = None,
                 source_file: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FileArgs.__new__(FileArgs)

            __props__.__dict__["create_directories"] = create_directories
            __props__.__dict__["datacenter"] = datacenter
            if datastore is None and not opts.urn:
                raise TypeError("Missing required property 'datastore'")
            __props__.__dict__["datastore"] = datastore
            if destination_file is None and not opts.urn:
                raise TypeError("Missing required property 'destination_file'")
            __props__.__dict__["destination_file"] = destination_file
            __props__.__dict__["source_datacenter"] = source_datacenter
            __props__.__dict__["source_datastore"] = source_datastore
            if source_file is None and not opts.urn:
                raise TypeError("Missing required property 'source_file'")
            __props__.__dict__["source_file"] = source_file
        super(File, __self__).__init__(
            'vsphere:index/file:File',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_directories: Optional[pulumi.Input[bool]] = None,
            datacenter: Optional[pulumi.Input[str]] = None,
            datastore: Optional[pulumi.Input[str]] = None,
            destination_file: Optional[pulumi.Input[str]] = None,
            source_datacenter: Optional[pulumi.Input[str]] = None,
            source_datastore: Optional[pulumi.Input[str]] = None,
            source_file: Optional[pulumi.Input[str]] = None) -> 'File':
        """
        Get an existing File resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] create_directories: Create directories in `destination_file`
               path parameter on first apply if any are missing for copy operation.
               
               > **NOTE:** Any directory created as part of the `create_directories` argument
               will not be deleted when the resource is destroyed. New directories are not
               created if the `destination_file` path is changed in subsequent applies.
        :param pulumi.Input[str] datacenter: The name of a datacenter to which the file will be
               uploaded.
        :param pulumi.Input[str] datastore: The name of the datastore to which to upload the
               file.
        :param pulumi.Input[str] destination_file: The path to where the file should be uploaded
               or copied to on the destination `datastore` in vSphere.
        :param pulumi.Input[str] source_datacenter: The name of a datacenter from which the file
               will be copied. Forces a new resource if changed.
        :param pulumi.Input[str] source_datastore: The name of the datastore from which file will
               be copied. Forces a new resource if changed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FileState.__new__(_FileState)

        __props__.__dict__["create_directories"] = create_directories
        __props__.__dict__["datacenter"] = datacenter
        __props__.__dict__["datastore"] = datastore
        __props__.__dict__["destination_file"] = destination_file
        __props__.__dict__["source_datacenter"] = source_datacenter
        __props__.__dict__["source_datastore"] = source_datastore
        __props__.__dict__["source_file"] = source_file
        return File(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createDirectories")
    def create_directories(self) -> pulumi.Output[Optional[bool]]:
        """
        Create directories in `destination_file`
        path parameter on first apply if any are missing for copy operation.

        > **NOTE:** Any directory created as part of the `create_directories` argument
        will not be deleted when the resource is destroyed. New directories are not
        created if the `destination_file` path is changed in subsequent applies.
        """
        return pulumi.get(self, "create_directories")

    @property
    @pulumi.getter
    def datacenter(self) -> pulumi.Output[Optional[str]]:
        """
        The name of a datacenter to which the file will be
        uploaded.
        """
        return pulumi.get(self, "datacenter")

    @property
    @pulumi.getter
    def datastore(self) -> pulumi.Output[str]:
        """
        The name of the datastore to which to upload the
        file.
        """
        return pulumi.get(self, "datastore")

    @property
    @pulumi.getter(name="destinationFile")
    def destination_file(self) -> pulumi.Output[str]:
        """
        The path to where the file should be uploaded
        or copied to on the destination `datastore` in vSphere.
        """
        return pulumi.get(self, "destination_file")

    @property
    @pulumi.getter(name="sourceDatacenter")
    def source_datacenter(self) -> pulumi.Output[Optional[str]]:
        """
        The name of a datacenter from which the file
        will be copied. Forces a new resource if changed.
        """
        return pulumi.get(self, "source_datacenter")

    @property
    @pulumi.getter(name="sourceDatastore")
    def source_datastore(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the datastore from which file will
        be copied. Forces a new resource if changed.
        """
        return pulumi.get(self, "source_datastore")

    @property
    @pulumi.getter(name="sourceFile")
    def source_file(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_file")

