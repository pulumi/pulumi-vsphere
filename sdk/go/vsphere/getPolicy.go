// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.getPolicy` data source can be used to discover the UUID of a
// vSphere storage policy. This can then be used with resources or data sources that
// require a storage policy.
// 
// > **NOTE:** Storage policy support is unsupported on direct ESXi connections and
// requires vCenter 6.0 or higher.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/storage_policy.html.markdown.
func LookupPolicy(ctx *pulumi.Context, args *GetPolicyArgs) (*GetPolicyResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("vsphere:index/getPolicy:getPolicy", inputs)
	if err != nil {
		return nil, err
	}
	return &GetPolicyResult{
		Name: outputs["name"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getPolicy.
type GetPolicyArgs struct {
	// The name of the storage policy.
	Name interface{}
}

// A collection of values returned by getPolicy.
type GetPolicyResult struct {
	Name interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
