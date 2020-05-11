package vsphere

import (
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/terraform-providers/terraform-provider-vsphere/vsphere"
)

const (
	vspherePkg = "vsphere"
	vsphereMod = "index"
)

// vsphereMember manufactures a type token for the VSphere package and the given module and type.
func vsphereMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(vspherePkg + ":" + mod + ":" + mem)
}

// vsphereType manufactures a type token for the VSphere package and the given module and type.
func vsphereType(mod string, typ string) tokens.Type {
	return tokens.Type(vsphereMember(mod, typ))
}

// vsphereDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the VSphere package and names the file by simply lower casing the data
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

func Provider() tfbridge.ProviderInfo {
	p := vsphere.Provider().(*schema.Provider)
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "vsphere",
		Description: "A Pulumi package for creating vsphere resources",
		Keywords:    []string{"pulumi", "vsphere"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-vsphere",
		Config: map[string]*tfbridge.SchemaInfo{
			"user": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_USER"},
				},
			},
			"password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_PASSWORD"},
				},
			},
			"vsphere_server": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_SERVER"},
				},
			},
			"allow_unverified_ssl": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_ALLOW_UNVERIFIED_SSL"},
				},
			},
			"client_debug": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_CLIENT_DEBUG"},
				},
			},
			"client_debug_path_run": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_CLIENT_DEBUG_PATH_RUN"},
				},
			},
			"client_debug_path": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_CLIENT_DEBUG_PATH"},
				},
			},
			"persist_session": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_PERSIST_SESSION"},
				},
			},
			"vim_session_path": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_VIM_SESSION_PATH"},
				},
			},
			"rest_session_path": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_REST_SESSION_PATH"},
				},
			},
			"vim_keep_alive": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_VIM_KEEP_ALIVE"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"vsphere_compute_cluster":                         {Tok: vsphereResource(vsphereMod, "ComputeCluster")},
			"vsphere_compute_cluster_host_group":              {Tok: vsphereResource(vsphereMod, "ComputeClusterHostGroup")},
			"vsphere_compute_cluster_vm_affinity_rule":        {Tok: vsphereResource(vsphereMod, "ComputeClusterVmAffinityRule")},
			"vsphere_compute_cluster_vm_anti_affinity_rule":   {Tok: vsphereResource(vsphereMod, "ComputeClusterVmAntiAffinityRule")},
			"vsphere_compute_cluster_vm_dependency_rule":      {Tok: vsphereResource(vsphereMod, "ComputeClusterVmDependencyRule")},
			"vsphere_compute_cluster_vm_group":                {Tok: vsphereResource(vsphereMod, "ComputeClusterVmGroup")},
			"vsphere_compute_cluster_vm_host_rule":            {Tok: vsphereResource(vsphereMod, "ComputeClusterVmHostRule")},
			"vsphere_content_library":                         {Tok: vsphereResource(vsphereMod, "ContentLibrary")},
			"vsphere_content_library_item":                    {Tok: vsphereResource(vsphereMod, "ContentLibraryItem")},
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
			"vsphere_ha_vm_override":                          {Tok: vsphereResource(vsphereMod, "HaVmOverride")},
			"vsphere_host":                                    {Tok: vsphereResource(vsphereMod, "Host")},
			"vsphere_host_port_group":                         {Tok: vsphereResource(vsphereMod, "HostPortGroup")},
			"vsphere_host_virtual_switch":                     {Tok: vsphereResource(vsphereMod, "HostVirtualSwitch")},
			"vsphere_license":                                 {Tok: vsphereResource(vsphereMod, "License")},
			"vsphere_nas_datastore":                           {Tok: vsphereResource(vsphereMod, "NasDatastore")},
			"vsphere_resource_pool":                           {Tok: vsphereResource(vsphereMod, "ResourcePool")},
			"vsphere_storage_drs_vm_override":                 {Tok: vsphereResource(vsphereMod, "StorageDrsVmOverride")},
			"vsphere_tag":                                     {Tok: vsphereResource(vsphereMod, "Tag")},
			"vsphere_tag_category":                            {Tok: vsphereResource(vsphereMod, "TagCategory")},
			"vsphere_vapp_container":                          {Tok: vsphereResource(vsphereMod, "VappContainer")},
			"vsphere_vapp_entity":                             {Tok: vsphereResource(vsphereMod, "VappEntity")},
			"vsphere_virtual_disk":                            {Tok: vsphereResource(vsphereMod, "VirtualDisk")},
			"vsphere_virtual_machine":                         {Tok: vsphereResource(vsphereMod, "VirtualMachine")},
			"vsphere_virtual_machine_snapshot":                {Tok: vsphereResource(vsphereMod, "VirtualMachineSnapshot")},
			"vsphere_vmfs_datastore":                          {Tok: vsphereResource(vsphereMod, "VmfsDatastore")},
			"vsphere_vnic":                                    {Tok: vsphereResource(vsphereMod, "Vnic")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"vsphere_compute_cluster":            {Tok: vsphereDataSource(vsphereMod, "getComputeCluster")},
			"vsphere_content_library":            {Tok: vsphereDataSource(vsphereMod, "getContentLibrary")},
			"vsphere_content_library_item":       {Tok: vsphereDataSource(vsphereMod, "getContentLibraryItem")},
			"vsphere_custom_attribute":           {Tok: vsphereDataSource(vsphereMod, "getCustomAttribute")},
			"vsphere_datacenter":                 {Tok: vsphereDataSource(vsphereMod, "getDatacenter")},
			"vsphere_datastore":                  {Tok: vsphereDataSource(vsphereMod, "getDatastore")},
			"vsphere_datastore_cluster":          {Tok: vsphereDataSource(vsphereMod, "getDatastoreCluster")},
			"vsphere_distributed_virtual_switch": {Tok: vsphereDataSource(vsphereMod, "getDistributedVirtualSwitch")},
			"vsphere_folder":                     {Tok: vsphereDataSource(vsphereMod, "getFolder")},
			"vsphere_host":                       {Tok: vsphereDataSource(vsphereMod, "getHost")},
			"vsphere_network":                    {Tok: vsphereDataSource(vsphereMod, "getNetwork")},
			"vsphere_resource_pool":              {Tok: vsphereDataSource(vsphereMod, "getResourcePool")},
			"vsphere_tag":                        {Tok: vsphereDataSource(vsphereMod, "getTag")},
			"vsphere_tag_category":               {Tok: vsphereDataSource(vsphereMod, "getTagCategory")},
			"vsphere_vapp_container":             {Tok: vsphereDataSource(vsphereMod, "getVappContainer")},
			"vsphere_virtual_machine":            {Tok: vsphereDataSource(vsphereMod, "getVirtualMachine")},
			"vsphere_vmfs_disks":                 {Tok: vsphereDataSource(vsphereMod, "getVmfsDisks")},
			"vsphere_storage_policy":             {Tok: vsphereDataSource(vsphereMod, "getPolicy")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi":    "^2.0.0",
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
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.0.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				"vsphere": "VSphere",
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
