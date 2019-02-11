# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

__config__ = pulumi.Config('vsphere')

allow_unverified_ssl = __config__.get('allowUnverifiedSsl')
"""
If set, VMware vSphere client will permit unverifiable SSL certificates.
"""

client_debug = __config__.get('clientDebug')
"""
govmomi debug
"""

client_debug_path = __config__.get('clientDebugPath')
"""
govmomi debug path for debug
"""

client_debug_path_run = __config__.get('clientDebugPathRun')
"""
govmomi debug path for a single run
"""

password = __config__.require('password')
"""
The user password for vSphere API operations.
"""

persist_session = __config__.get('persistSession')
"""
Persist vSphere client sessions to disk
"""

rest_session_path = __config__.get('restSessionPath')
"""
The directory to save vSphere REST API sessions to
"""

user = __config__.require('user')
"""
The user name for vSphere API operations.
"""

vcenter_server = __config__.get('vcenterServer')

vim_session_path = __config__.get('vimSessionPath')
"""
The directory to save vSphere SOAP API sessions to
"""

vsphere_server = __config__.get('vsphereServer')
"""
The vSphere Server name for vSphere API operations.
"""

