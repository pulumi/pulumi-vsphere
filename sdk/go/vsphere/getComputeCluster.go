// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `ComputeCluster` data source can be used to discover the ID of a
// cluster in vSphere. This is useful to fetch the ID of a cluster that you want
// to use for virtual machine placement via the
// `VirtualMachine` resource, allowing
// you to specify the cluster's root resource pool directly versus using the alias
// available through the `ResourcePool`
// data source.
//
// > You may also wish to see the
// `ComputeCluster` resource for further
// details about clusters or how to work with them.
func LookupComputeCluster(ctx *pulumi.Context, args *LookupComputeClusterArgs, opts ...pulumi.InvokeOption) (*LookupComputeClusterResult, error) {
	var rv LookupComputeClusterResult
	err := ctx.Invoke("vsphere:index/getComputeCluster:getComputeCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeCluster.
type LookupComputeClusterArgs struct {
	// The managed object reference
	// ID of the datacenter the cluster is located in.  This can
	// be omitted if the search path used in `name` is an absolute path.  For
	// default datacenters, use the id attribute from an empty `Datacenter`
	// data source.
	DatacenterId *string `pulumi:"datacenterId"`
	// The name or absolute path to the cluster.
	Name string `pulumi:"name"`
}

// A collection of values returned by getComputeCluster.
type LookupComputeClusterResult struct {
	DatacenterId *string `pulumi:"datacenterId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	ResourcePoolId string `pulumi:"resourcePoolId"`
}
