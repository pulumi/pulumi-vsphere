CHANGELOG
=========

## HEAD (Unreleased)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0
* Upgrade to v1.18.2 of the vSphere Terraform Provider

---

## 2.3.2 (2020-05-14)
* Upgrade to v1.18.1 of the vSphere Terraform Provider

## 2.3.1 (2020-05-11)
* Upgrade to pulumi-terraform-bridge v2.3.1

## 2.3.0 (2020-05-05)
* Upgrade to v1.18.0 of the vSphere Terraform Provider

## 2.2.1 (2020-04-29)
* Upgrade to v1.17.4 of the vSphere Terraform Provider

## 2.2.0 (2020-04-28)
* Regenerate datasource examples to be async
* Upgrade to pulumi-terraform-bridge v2.1.0

## 2.1.1 (2020-04-24)
* Upgrade to v1.17.3 of the vSphere Terraform Provider

## 2.1.0 (2020-04-20)
* Upgrade to v1.17.2 of the vSphere Terraform Provider

## 2.0.0 (2020-04-17)
* Update to pulumi-terraform-bridge v2.0.0
* Update to pulumi  v2.0.0

## 1.9.0 (2020-04-14)
* Upgrade to v1.17.1 of the vSphere Terraform Provider

## 1.8.0 (2020-04-01)
* Update to pulumi-terraform-bridge v1.8.4
* Switch to go mod layout

## 1.7.0 (2020-03-24)
* Upgrade to v1.17.0 of the vSphere Terraform Provider

## 1.6.0 (2020-03-23)
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## 1.5.0 (2020-03-06)
* Upgrade to v1.16.2 of the vSphere Terraform Provider

## 1.4.1 (2020-02-07)
* Upgrade to v1.16.1 of the vSphere Terraform Provider

## 1.4.0 (2020-02-06)
* Upgrade to v1.16.0 of the vSphere Terraform Provider

## 1.3.0 (2019-12-19)
* Namespace names in .NET SDK are adjusted to PascalCase
([#51](https://github.com/pulumi/pulumi-vsphere/pull/51)).
* Upgrade to v1.14.0 of the vSphere Terraform Provider

## 1.2.0 (2019-12-04)
* Upgrade to support go 1.13.x
* Upgrade to pulumi-terraform@dotnet

## 1.1.0 (2019-11-13)
* Remove README.rst from the Python package and replace it with README.md

## 1.0.0 (2019-11-13)
* Add support for DotNet SDK Generation

## 0.17.12 (2019-10-02)
* Regenerate SDK against tf2pulumi 0.6.0
* Upgrade to v1.13.0 of the vSphere Terraform Provider

## 0.17.11 (2019-09-05)
* Upgrade to Pulumi v1.0.0

## 0.17.10 (2019-08-29)
* Upgrade pulumi-terraform to 3f206601e7

## 0.17.9 (2019-08-20)
* Depend on latest pulumi package

## 0.17.8 (2019-08-09)
* Update to pulumi-terraform@9db2fc93cd

## 0.17.7 (2019-07-09)
* Fix detailed diffs with nested computed values.

## 0.17.6 (2019-07-08)
* Communicate detailed information about the difference between a resource's desired and actual state during a Pulumi update

## 0.17.5 (2019-06-21)
* Update to pulumi-terraform@3635bed3a5 which stops maps containing `.` being treated as nested maps.

## 0.17.4 (2019-06-20)
* Add TypeScript type guards for each resource class ([7ace3e9b5f](https://github.com/pulumi/pulumi-terraform/commit/7ace3e9b5f2dcd4692b029ba4b80360582d7949a))
* Update to v1.12.0 of the vSphere Terraform Provider

## 0.17.3 (2019-06-05)
* Update to v1.11.0 of the vSphere Terraform provider

## 0.17.2 (2019-02-22)
* Update the version of the Pulumi-Terraform bridge. No user facing changes should be visible.

## 0.17.1 (2019-03-20)
* Update to v1.10.0 of the vSphere Terraform provider
* Fix a bug where setting a property value back to a default value results in no change
* Project numeric types in Python as `float` instead of `int`, fixing some crashes

## 0.17.0 (2019-03-06)
* Updated to the latest version of the `pulumi` SDK

## 0.16.4 (2019-02-11)
* Add support for the `deleteBeforeReplace` resource option and improved delete-before-replace behaviour introduced in Pulumi v0.16.14

## 0.16.3 (2019-01-19)
* Update to v1.9.1 of the vSphere Terraform provider
* Add documentation comments to the Node.js SDK

## 0.16.2 (2018-12-05)
* Update to v1.9.0 of the vSphere Terraform provider
* Fix an issue where the Pulumi CLI would warn of an incorrect plugin version for vSphere

## 0.16.1 (2018-11-13)
* Support Python 3.6 and higher, instead of only Python 3.7

## 0.16.0 (2018-10-16)
* Initial release of the vSphere Provider
