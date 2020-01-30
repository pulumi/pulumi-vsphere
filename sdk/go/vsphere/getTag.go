// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/tag.html.markdown.
func LookupTag(ctx *pulumi.Context, args *GetTagArgs) (*GetTagResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["categoryId"] = args.CategoryId
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("vsphere:index/getTag:getTag", inputs)
	if err != nil {
		return nil, err
	}
	return &GetTagResult{
		CategoryId: outputs["categoryId"],
		Description: outputs["description"],
		Name: outputs["name"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getTag.
type GetTagArgs struct {
	// The ID of the tag category the tag is located in.
	CategoryId interface{}
	// The name of the tag.
	Name interface{}
}

// A collection of values returned by getTag.
type GetTagResult struct {
	CategoryId interface{}
	Description interface{}
	Name interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
