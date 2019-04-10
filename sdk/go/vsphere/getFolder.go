// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `vsphere_folder` data source can be used to get the general attributes of a
// vSphere inventory folder. Paths are absolute and include must include the
// datacenter.  
func LookupFolder(ctx *pulumi.Context, args *GetFolderArgs) (*GetFolderResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["path"] = args.Path
	}
	outputs, err := ctx.Invoke("vsphere:index/getFolder:getFolder", inputs)
	if err != nil {
		return nil, err
	}
	return &GetFolderResult{
		Path: outputs["path"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getFolder.
type GetFolderArgs struct {
	// The absolute path of the folder. For example, given a
	// default datacenter of `default-dc`, a folder of type `vm`, and a folder name
	// of `terraform-test-folder`, the resulting path would be
	// `/default-dc/vm/terraform-test-folder`. The valid folder types to be used in
	// the path are: `vm`, `host`, `datacenter`, `datastore`, or `network`.
	Path interface{}
}

// A collection of values returned by getFolder.
type GetFolderResult struct {
	Path interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
