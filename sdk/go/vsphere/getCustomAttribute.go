// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/custom_attribute.html.markdown.
func LookupCustomAttribute(ctx *pulumi.Context, args *GetCustomAttributeArgs) (*GetCustomAttributeResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("vsphere:index/getCustomAttribute:getCustomAttribute", inputs)
	if err != nil {
		return nil, err
	}
	return &GetCustomAttributeResult{
		ManagedObjectType: outputs["managedObjectType"],
		Name:              outputs["name"],
		Id:                outputs["id"],
	}, nil
}

// A collection of arguments for invoking getCustomAttribute.
type GetCustomAttributeArgs struct {
	// The name of the custom attribute.
	Name interface{}
}

// A collection of values returned by getCustomAttribute.
type GetCustomAttributeResult struct {
	ManagedObjectType interface{}
	Name              interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
