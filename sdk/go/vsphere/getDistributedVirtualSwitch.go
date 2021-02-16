// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `DistributedVirtualSwitch` data source can be used to discover
// the ID and uplink data of a of a vSphere distributed virtual switch (DVS). This
// can then be used with resources or data sources that require a DVS, such as the
// `DistributedPortGroup` resource, for which
// an example is shown below.
//
// > **NOTE:** This data source requires vCenter and is not available on direct
// ESXi connections.
//
// ## Example Usage
//
// The following example locates a DVS that is named `test-dvs`, in the
// datacenter `dc1`. It then uses this DVS to set up a
// `DistributedPortGroup` resource that uses the first uplink as a
// primary uplink and the second uplink as a secondary.
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
// 		datacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := datacenter.Id
// 		dvs, err := vsphere.LookupDistributedVirtualSwitch(ctx, &vsphere.LookupDistributedVirtualSwitchArgs{
// 			DatacenterId: &opt1,
// 			Name:         "test-dvs",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.NewDistributedPortGroup(ctx, "pg", &vsphere.DistributedPortGroupArgs{
// 			ActiveUplinks: pulumi.StringArray{
// 				pulumi.String(dvs.Uplinks[0]),
// 			},
// 			DistributedVirtualSwitchUuid: pulumi.String(dvs.Id),
// 			StandbyUplinks: pulumi.StringArray{
// 				pulumi.String(dvs.Uplinks[1]),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDistributedVirtualSwitch(ctx *pulumi.Context, args *LookupDistributedVirtualSwitchArgs, opts ...pulumi.InvokeOption) (*LookupDistributedVirtualSwitchResult, error) {
	var rv LookupDistributedVirtualSwitchResult
	err := ctx.Invoke("vsphere:index/getDistributedVirtualSwitch:getDistributedVirtualSwitch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDistributedVirtualSwitch.
type LookupDistributedVirtualSwitchArgs struct {
	// The managed object reference
	// ID of the datacenter the DVS is located in. This can be
	// omitted if the search path used in `name` is an absolute path. For default
	// datacenters, use the id attribute from an empty `Datacenter` data
	// source.
	DatacenterId *string `pulumi:"datacenterId"`
	// The name of the distributed virtual switch. This can be a
	// name or path.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDistributedVirtualSwitch.
type LookupDistributedVirtualSwitchResult struct {
	DatacenterId *string `pulumi:"datacenterId"`
	// The provider-assigned unique ID for this managed resource.
	Id      string   `pulumi:"id"`
	Name    string   `pulumi:"name"`
	Uplinks []string `pulumi:"uplinks"`
}
