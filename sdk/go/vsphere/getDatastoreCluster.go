// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `DatastoreCluster` data source can be used to discover the ID of a
// datastore cluster in vSphere. This is useful to fetch the ID of a datastore
// cluster that you want to use to assign datastores to using the
// `NasDatastore` or
// `VmfsDatastore` resources, or create
// virtual machines in using the
// `VirtualMachine` resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vsphere/sdk/v2/go/vsphere"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "dc1"
// 		_, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := data.Vsphere_datacenter.Dc.Id
// 		_, err = vsphere.LookupDatastoreCluster(ctx, &vsphere.LookupDatastoreClusterArgs{
// 			DatacenterId: &opt1,
// 			Name:         "datastore-cluster1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDatastoreCluster(ctx *pulumi.Context, args *LookupDatastoreClusterArgs, opts ...pulumi.InvokeOption) (*LookupDatastoreClusterResult, error) {
	var rv LookupDatastoreClusterResult
	err := ctx.Invoke("vsphere:index/getDatastoreCluster:getDatastoreCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatastoreCluster.
type LookupDatastoreClusterArgs struct {
	// The managed object reference
	// ID of the datacenter the datastore cluster is located in.
	// This can be omitted if the search path used in `name` is an absolute path.
	// For default datacenters, use the id attribute from an empty
	// `Datacenter` data source.
	DatacenterId *string `pulumi:"datacenterId"`
	// The name or absolute path to the datastore cluster.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDatastoreCluster.
type LookupDatastoreClusterResult struct {
	DatacenterId *string `pulumi:"datacenterId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}
