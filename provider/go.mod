module github.com/pulumi/pulumi-vsphere/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/hashicorp/terraform-provider-vsphere v1.18.2
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.5.2
	github.com/pulumi/pulumi/sdk/v2 v2.5.1-0.20200630091945-bb358c4d2173
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/terraform-provider-vsphere => github.com/pulumi/terraform-provider-vsphere v1.18.2-0.20200807194018-06041a716d3e
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
