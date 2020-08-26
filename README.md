[![Actions Status](https://github.com/pulumi/pulumi-vsphere/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-vsphere/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fvsphere.svg)](https://www.npmjs.com/package/@pulumi/vsphere)
[![Python version](https://badge.fury.io/py/pulumi-vsphere.svg)](https://pypi.org/project/pulumi-vsphere)
[![NuGet version](https://badge.fury.io/nu/pulumi.vsphere.svg)](https://badge.fury.io/nu/pulumi.vsphere)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-vsphere/sdk/v2/go)](https://pkg.go.dev/github.com/pulumi/pulumi-vsphere/sdk/v2/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-vsphere/blob/master/LICENSE)

# VSphere provider

The VSphere resource provider for Pulumi lets you use VSphere resources in your infrastructure 
programs. To use this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/vsphere

or `yarn`:

    $ yarn add @pulumi/vsphere

### Python

To use from Python, install using `pip`:

    $ pip install pulumi-vsphere

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-vsphere/sdk/v2/go/...
    
### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Vsphere   
 
## Configuration

The following configuration points are available:

- `vsphere:user` - (Required) This is the username for vSphere API operations. Can also be specified with the `VSPHERE_USER`
  environment variable.
- `vsphere:password` - (Required) This is the password for vSphere API operations. Can also be specified with the 
  `VSPHERE_PASSWORD` environment variable.
- `vsphere:vsphereServer` - (Required) This is the vCenter server name for vSphere API operations. Can also be specified
  with the `VSPHERE_SERVER` environment variable.
- `vsphere:allowUnverifiedSsl` - (Optional) Boolean that can be set to true to disable SSL certificate verification. 
  This should be used with care as it could allow an attacker to intercept your auth token. If omitted, default value is
  `false`. Can also be specified with the `VSPHERE_ALLOW_UNVERIFIED_SSL` environment variable.
- `vsphere:vimKeepAlive` - (Optional) Keep alive interval in minutes for the VIM session. Standard session timeout in 
  vSphere is 30 minutes. This defaults to `10` minutes to ensure that operations that take a longer than 30 minutes 
  without API interaction do not result in a session timeout. Can also be specified with the `VSPHERE_VIM_KEEP_ALIVE`
  environment variable.
- `vsphere:persistSession` - (Optional) Persist the SOAP and REST client sessions to disk. Default: false. Can also be 
  specified by the `VSPHERE_PERSIST_SESSION` environment variable.
- `vsphere:vimSessionPath` - (Optional) The direcotry to save the VIM SOAP API session to. Default: `${HOME}/.govmomi/sessions`.
  Can also be specified by the `VSPHERE_VIM_SESSION_PATH` environment variable.
- `vsphere:restSessionPath` - (Optional) The directory to save the REST API session (used for tags) to. Default: `${HOME}/.govmomi/rest_sessions`. 
  Can also be specified by the `VSPHERE_REST_SESSION_PATH` environment variable.
- `vsphere:clientDebug` - (Optional) When `true`, the provider logs SOAP calls made to the vSphere API to disk. The log 
  files are logged to `${HOME}/.govmomi`. Can also be specified with the `VSPHERE_CLIENT_DEBUG` environment variable.
- `vsphere:clientDebugPath` - (Optional) Override the default log path. Can also be specified with the 
  `VSPHERE_CLIENT_DEBUG_PATH` environment variable.
- `vsphere:clientDebugPathRun` - (Optional) A specific subdirectory in `client_debug_path` to use for debugging calls for
  this particular provider configuration. All data in this directory is removed at the start of the provider run. Can also
  be specified with the `VSPHERE_CLIENT_DEBUG_PATH_RUN` environment variable.

## Reference

For further information, please visit [the vSphere provider docs](https://www.pulumi.com/docs/intro/cloud-providers/vsphere) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/vsphere).
