module github.com/pulumi/pulumi-vsphere/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.5.0
	github.com/pulumi/pulumi/sdk/v2 v2.4.0
	github.com/terraform-providers/terraform-provider-vsphere v1.18.2
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-vsphere => github.com/pulumi/terraform-provider-vsphere v1.18.2-0.20200602092320-28748e1841c5
)
