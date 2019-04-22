## 0.17.3 (Unreleased)

## 0.17.2 (Released April 22nd, 2019)

- Updated the version of pulumi-terraform on which the provider is based. No user-facing changes
  should be visible.

## 0.17.1 (Released March 20th, 2019)

## Improvements

- Update to v1.10.0 of the VSphere Terraform Provider.

- Fix a bug where setting a property value back to the default results in no change

- Numeric types in Python are now projected as `float` instead of `int`, fixing some crashes.

## 0.17.0 (Released March 6th, 2019)

## Improvements

- Depend on latest version of `@pulumi/pulumi`

## 0.16.4 (Released February 11th, 2019)

### Improvements

- Support for the `deleteBeforeReplace` resource option and improved
  delete-before-replace behaviour introduced in [Pulumi
  0.16.14](https://github.com/pulumi/pulumi/blob/master/CHANGELOG.md#01614-released-january-31st-2019).

## 0.16.3 (Released January 19th, 2019)

### Improvements

- Updated to the v1.9.1 version of the Terraform VSphere Provider.

- Documentation comments for the Node.js SDK now include examples

## 0.16.2 (Released December 5th, 2018)

### Improvements

- Update provider to v1.9.0

- Fix an issue where `pulumi` would issue a warning that the plugin version was incorrect.

## 0.16.1 (Released Novemeber 13th, 2018)

### Major Changes

- If you're using Pulumi with Python, this release removes Python 2.7 support in favor of Python 3.6 and greater.
