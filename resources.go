package vsphere

import (
	"errors"
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
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
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	if stringValue(vars, "token") == "" {
		return errors.New("unable to discover VSphere credentials" +
			"- see https://pulumi.io/install/vsphere.html for details on configuration")
	}

	return nil
}

func Provider() tfbridge.ProviderInfo {
	p := vsphere.Provider().(*schema.Provider)
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "vsphere",
		Description:          "A Pulumi package for creating vsphere resources",
		Keywords:             []string{"pulumi", "vsphere"},
		License:              "MIT",
		Homepage:             "https://pulumi.io",
		Repository:           "https://github.com/Smithx10/pulumi-vsphere",
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"vsphere_virtual_machine": {Tok: vsphereResource(vsphereMod, "VirtualMachine")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"vsphere_virtual_machine": {Tok: vsphereDataSource(vsphereMod, "getVirtualMachine")},
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
