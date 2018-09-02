package vsphere

import (
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-vsphere/vsphere"
)

const (
	vspherePkg = "vsphere"
	vsphereMod = "index"
)

// vsphereMember manufactures a type token for the Digital Ocean package and the given module and type.
func vsphereMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(vspherePkg + ":" + mod + ":" + mem)
}

// vsphereType manufactures a type token for the Digital Ocean package and the given module and type.
func vsphereType(mod string, typ string) tokens.Type {
	return tokens.Type(vsphereMember(mod, typ))
}

// vsphereDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the Digital Ocean package and names the file by simply lower casing the data
// source's first character.
func vsphereDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return vsphereMember(mod+"/"+fn, res)
}

// vsphereResource manufactures a standard resource token given a module and resource name.
// package and names the file by simply lower casing the resource's first character.
func vsphereResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return vsphereType(mod+"/"+fn, res)
}

// stringValue gets a string value from a property map if present, else ""
func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
	val, ok := vars[prop]
	if ok && val.IsString() {
		return val.StringValue()
	}
	return ""
}

// preConfigureCallback validates that AWS credentials can be succesfully discovered. This emulates the credentials
// configuration subset of `github.com/terraform-providers/terraform-provider-aws/aws.providerConfigure`.  We do this
// before passing control to the TF provider to ensure we can report actionable errors.

//InsecureFlag:  true,
//User:          stringValue(vars, "user"),
//Password:      stringValue(vars, "password"),
//VSphereServer: stringValue(vars, "server"),
//}

//fmt.Println(config)
////if stringValue(vars, "token") == "" {
////return errors.New("unable to discover VSphere credentials" +
////"- see https://pulumi.io/install/vsphere.html for details on configuration")
////}

//return nil
//}

func Provider() tfbridge.ProviderInfo {
	p := vsphere.Provider().(*schema.Provider)
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "vsphere",
		Description: "A Pulumi package for creating vsphere resources",
		Keywords:    []string{"pulumi", "vsphere"},
		License:     "MIT",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/Smithx10/pulumi-vsphere",
		//PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"vsphere_compute_cluster":                         {Tok: vsphereResource(vsphereMod, "ComputeCluster")},
			"vsphere_compute_cluster_host_group":              {Tok: vsphereResource(vsphereMod, "ComputeClusterHostGroup")},
			"vsphere_compute_cluster_vm_affinity_rule":        {Tok: vsphereResource(vsphereMod, "ComputeClusterVmAffinityRule")},
			"vsphere_compute_cluster_vm_dependency_rule":      {Tok: vsphereResource(vsphereMod, "ComputeClusterVmDependencyRule")},
			"vsphere_compute_cluster_vm_group":                {Tok: vsphereResource(vsphereMod, "ComputeClusterVmGroup")},
			"vsphere_compute_cluster_vm_host_rule":            {Tok: vsphereResource(vsphereMod, "ComputeClusterVmHostRule")},
			"vsphere_custom_attribute":                        {Tok: vsphereResource(vsphereMod, "CustomAttribute")},
			"vsphere_datacenter":                              {Tok: vsphereResource(vsphereMod, "Datacenter")},
			"vsphere_datastore_cluster":                       {Tok: vsphereResource(vsphereMod, "DatastoreCluster")},
			"vsphere_datastore_cluster_vm_anti_affinity_rule": {Tok: vsphereResource(vsphereMod, "DatastoreClusterVmAntiAffinityRule")},
			"vsphere_distributed_port_group":                  {Tok: vsphereResource(vsphereMod, "DistributedPortGroup")},
			"vsphere_distributed_virtual_switch":              {Tok: vsphereResource(vsphereMod, "DistributedVirtualSwitch")},
			"vsphere_dpm_host_override":                       {Tok: vsphereResource(vsphereMod, "DpmHostOverride")},
			"vsphere_drs_vm_override":                         {Tok: vsphereResource(vsphereMod, "DrsVmOverride")},
			"vsphere_file":                                    {Tok: vsphereResource(vsphereMod, "File")},
			"vsphere_folder":                                  {Tok: vsphereResource(vsphereMod, "Folder")},
			"vsphere_folder_migrate":                          {Tok: vsphereResource(vsphereMod, "FolderMigrate")},
			"vsphere_ha_vm_override":                          {Tok: vsphereResource(vsphereMod, "HaVmOverride")},
			"vsphere_host_port_group":                         {Tok: vsphereResource(vsphereMod, "HostPortGroup")},
			"vsphere_host_virtual_switch":                     {Tok: vsphereResource(vsphereMod, "HostVirtualSwitch")},
			"vsphere_license":                                 {Tok: vsphereResource(vsphereMod, "License")},
			"vsphere_nas_datastore":                           {Tok: vsphereResource(vsphereMod, "NasDatastore")},
			"vsphere_resource_pool":                           {Tok: vsphereResource(vsphereMod, "ResourcePool")},
			"vsphere_storage_drs_vm_override":                 {Tok: vsphereResource(vsphereMod, "StorageDrsVmOverride")},
			"vsphere_tag":                                     {Tok: vsphereResource(vsphereMod, "Tag")},
			"vsphere_tag_category":                            {Tok: vsphereResource(vsphereMod, "TagCategory")},
			"vsphere_vapp_container":                          {Tok: vsphereResource(vsphereMod, "VappContainer")},
			"vsphere_virtual_disk":                            {Tok: vsphereResource(vsphereMod, "VirtualDisk")},
			"vsphere_virtual_machine":                         {Tok: vsphereResource(vsphereMod, "VirtualMachine")},
			"vsphere_virtual_machine_migrate":                 {Tok: vsphereResource(vsphereMod, "VirtualMachineMigrate")},
			"vsphere_virtual_machine_snapshot":                {Tok: vsphereResource(vsphereMod, "VirtualMachineSnapshot")},
			"vsphere_vmfs_datastore":                          {Tok: vsphereResource(vsphereMod, "VmfsDatastore")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"vsphere_compute_cluster":            {Tok: vsphereDataSource(vsphereMod, "getComputeCluster")},
			"vsphere_custom_attribute":           {Tok: vsphereDataSource(vsphereMod, "getCustomAttribute")},
			"vsphere_datacenter":                 {Tok: vsphereDataSource(vsphereMod, "getDatacenter")},
			"vsphere_datastore":                  {Tok: vsphereDataSource(vsphereMod, "getDatastore")},
			"vsphere_datastore_cluster":          {Tok: vsphereDataSource(vsphereMod, "getDatastoreCluster")},
			"vsphere_distributed_virtual_switch": {Tok: vsphereDataSource(vsphereMod, "getDistributedVirtualSwitch")},
			"vsphere_host":                       {Tok: vsphereDataSource(vsphereMod, "getHost")},
			"vsphere_network":                    {Tok: vsphereDataSource(vsphereMod, "getNetwork")},
			"vsphere_resource_pool":              {Tok: vsphereDataSource(vsphereMod, "getResourcePool")},
			"vsphere_tag":                        {Tok: vsphereDataSource(vsphereMod, "getTag")},
			"vsphere_tag_category":               {Tok: vsphereDataSource(vsphereMod, "getTagCategory")},
			"vsphere_virtual_machine":            {Tok: vsphereDataSource(vsphereMod, "getVirtualMachine")},
			"vsphere_vmfs_disks":                 {Tok: vsphereDataSource(vsphereMod, "getVmfsDisks")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi":    "^0.15.0",
				"builtin-modules":   "3.0.0",
				"read-package-tree": "^5.2.1",
				"resolve":           "^1.8.1",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.9.2",
			},
			Overlay: &tfbridge.OverlayInfo{
				Files:   []string{},
				Modules: map[string]*tfbridge.OverlayInfo{},
			},
		},
	}
	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameField = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameField]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameField]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameField] = tfbridge.AutoName(nameField, 255)
				}
			}
		}
	}

	return prov
}
