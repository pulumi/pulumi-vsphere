// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getVmfsDisks` data source can be used to discover the storage
// devices available on an ESXi host. This data source can be combined with the
// `VmfsDatastore` resource to create VMFS
// datastores based off a set of discovered disks.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vsphere/sdk/v3/go/vsphere"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "dc1"
// 		datacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := "esxi1"
// 		host, err := vsphere.LookupHost(ctx, &vsphere.LookupHostArgs{
// 			DatacenterId: datacenter.Id,
// 			Name:         &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt2 := "mpx.vmhba1:C0:T[12]:L0"
// 		opt3 := true
// 		_, err = vsphere.GetVmfsDisks(ctx, &vsphere.GetVmfsDisksArgs{
// 			Filter:       &opt2,
// 			HostSystemId: host.Id,
// 			Rescan:       &opt3,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetVmfsDisks(ctx *pulumi.Context, args *GetVmfsDisksArgs, opts ...pulumi.InvokeOption) (*GetVmfsDisksResult, error) {
	var rv GetVmfsDisksResult
	err := ctx.Invoke("vsphere:index/getVmfsDisks:getVmfsDisks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVmfsDisks.
type GetVmfsDisksArgs struct {
	// A regular expression to filter the disks against. Only
	// disks with canonical names that match will be included.
	Filter *string `pulumi:"filter"`
	// The managed object ID of
	// the host to look for disks on.
	HostSystemId string `pulumi:"hostSystemId"`
	// Whether or not to rescan storage adapters before
	// searching for disks. This may lengthen the time it takes to perform the
	// search. Default: `false`.
	Rescan *bool `pulumi:"rescan"`
}

// A collection of values returned by getVmfsDisks.
type GetVmfsDisksResult struct {
	// A lexicographically sorted list of devices discovered by the
	// operation, matching the supplied `filter`, if provided.
	Disks        []string `pulumi:"disks"`
	Filter       *string  `pulumi:"filter"`
	HostSystemId string   `pulumi:"hostSystemId"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Rescan *bool  `pulumi:"rescan"`
}
