module github.com/pulumi/pulumi-vsphere/provider/v3

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/hashicorp/terraform-provider-vsphere v1.18.2
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0
	github.com/pulumi/pulumi/pkg/v3 v3.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.0.0
)

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-provider-vsphere => github.com/pulumi/terraform-provider-vsphere v1.18.2-0.20210319141625-af6d6b961182
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
