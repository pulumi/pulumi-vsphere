// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.DatastoreCluster` data source can be used to discover the ID of a
// datastore cluster in vSphere. This is useful to fetch the ID of a datastore
// cluster that you want to use to assign datastores to using the
// [`.NasDatastore`][docs-nas-datastore-resource] or
// [`.VmfsDatastore`][docs-vmfs-datastore-resource] resources, or create
// virtual machines in using the
// [`.VirtualMachine`][docs-virtual-machine-resource] resource. 
// 
// [docs-nas-datastore-resource]: /docs/providers/vsphere/r/nas_datastore.html
// [docs-vmfs-datastore-resource]: /docs/providers/vsphere/r/vmfs_datastore.html
// [docs-virtual-machine-resource]: /docs/providers/vsphere/r/virtual_machine.html
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/datastore_cluster.html.markdown.
func LookupDatastoreCluster(ctx *pulumi.Context, args *GetDatastoreClusterArgs) (*GetDatastoreClusterResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["datacenterId"] = args.DatacenterId
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("vsphere:index/getDatastoreCluster:getDatastoreCluster", inputs)
	if err != nil {
		return nil, err
	}
	return &GetDatastoreClusterResult{
		DatacenterId: outputs["datacenterId"],
		Name: outputs["name"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getDatastoreCluster.
type GetDatastoreClusterArgs struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the datacenter the datastore cluster is located in.
	// This can be omitted if the search path used in `name` is an absolute path.
	// For default datacenters, use the id attribute from an empty
	// `.Datacenter` data source.
	DatacenterId interface{}
	// The name or absolute path to the datastore cluster.
	Name interface{}
}

// A collection of values returned by getDatastoreCluster.
type GetDatastoreClusterResult struct {
	DatacenterId interface{}
	Name interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
