# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call

class InstallPluginCommand(install):
    def run(self):
        install.run(self)
        check_call(['pulumi', 'plugin', 'install', 'resource', 'vsphere', '-1536091003-g5a44e2b'])

def readme():
    with open('README.rst') as f:
        return f.read()

setup(name='pulumi_vsphere',
      version='',
      description='A Pulumi package for creating vsphere resources',
      long_description=readme(),
      cmdclass={
          'install': InstallPluginCommand,
      },
      keywords='pulumi vsphere',
      url='https://pulumi.io',
      project_urls={
          'Repository': 'https://github.com/pulumi/pulumi-vsphere'
      },
      license='MIT',
      packages=find_packages(),
      install_requires=[
          'pulumi'
      ],
      zip_safe=False)
