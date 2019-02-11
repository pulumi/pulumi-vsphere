# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class File(pulumi.CustomResource):
    create_directories: pulumi.Output[bool]
    """
    Create directories in `destination_file`
    path parameter if any missing for copy operation.
    """
    datacenter: pulumi.Output[str]
    """
    The name of a datacenter in which the file will be
    uploaded to.
    """
    datastore: pulumi.Output[str]
    """
    The name of the datastore in which to upload the
    file to.
    """
    destination_file: pulumi.Output[str]
    """
    The path to where the file should be uploaded
    or copied to on vSphere.
    """
    source_datacenter: pulumi.Output[str]
    """
    The name of a datacenter in which the file
    will be copied from. Forces a new resource if changed.
    """
    source_datastore: pulumi.Output[str]
    """
    The name of the datastore in which file will
    be copied from. Forces a new resource if changed.
    """
    source_file: pulumi.Output[str]
    """
    The path to the file being uploaded from the
    Terraform host to vSphere or copied within vSphere. Forces a new resource if
    changed.
    """
    def __init__(__self__, resource_name, opts=None, create_directories=None, datacenter=None, datastore=None, destination_file=None, source_datacenter=None, source_datastore=None, source_file=None, __name__=None, __opts__=None):
        """
        The `vsphere_file` resource can be used to upload files (such as virtual disk
        files) from the host machine that Terraform is running on to a target
        datastore.  The resource can also be used to copy files between datastores, or
        from one location to another on the same datastore.
        
        Updates to destination parameters such as `datacenter`, `datastore`, or
        `destination_file` will move the managed file a new destination based on the
        values of the new settings.  If any source parameter is changed, such as
        `source_datastore`, `source_datacenter` or `source_file`), the resource will be
        re-created. Depending on if destination parameters are being changed as well,
        this may result in the destination file either being overwritten or deleted at
        the old location.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] create_directories: Create directories in `destination_file`
               path parameter if any missing for copy operation.
        :param pulumi.Input[str] datacenter: The name of a datacenter in which the file will be
               uploaded to.
        :param pulumi.Input[str] datastore: The name of the datastore in which to upload the
               file to.
        :param pulumi.Input[str] destination_file: The path to where the file should be uploaded
               or copied to on vSphere.
        :param pulumi.Input[str] source_datacenter: The name of a datacenter in which the file
               will be copied from. Forces a new resource if changed.
        :param pulumi.Input[str] source_datastore: The name of the datastore in which file will
               be copied from. Forces a new resource if changed.
        :param pulumi.Input[str] source_file: The path to the file being uploaded from the
               Terraform host to vSphere or copied within vSphere. Forces a new resource if
               changed.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['create_directories'] = create_directories

        __props__['datacenter'] = datacenter

        if datastore is None:
            raise TypeError('Missing required property datastore')
        __props__['datastore'] = datastore

        if destination_file is None:
            raise TypeError('Missing required property destination_file')
        __props__['destination_file'] = destination_file

        __props__['source_datacenter'] = source_datacenter

        __props__['source_datastore'] = source_datastore

        if source_file is None:
            raise TypeError('Missing required property source_file')
        __props__['source_file'] = source_file

        super(File, __self__).__init__(
            'vsphere:index/file:File',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

