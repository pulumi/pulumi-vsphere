// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.Folder` resource can be used to manage vSphere inventory folders.
// The resource supports creating folders of the 5 major types - datacenter
// folders, host and cluster folders, virtual machine folders, datastore folders,
// and network folders.
//
// Paths are always relative to the specific type of folder you are creating.
// Subfolders are discovered by parsing the relative path specified in `path`, so
// `foo/bar` will create a folder named `bar` in the parent folder `foo`, as long
// as that folder exists.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/folder.html.markdown.
type Folder struct {
	s *pulumi.ResourceState
}

// NewFolder registers a new resource with the given unique name, arguments, and options.
func NewFolder(ctx *pulumi.Context,
	name string, args *FolderArgs, opts ...pulumi.ResourceOpt) (*Folder, error) {
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["customAttributes"] = nil
		inputs["datacenterId"] = nil
		inputs["path"] = nil
		inputs["tags"] = nil
		inputs["type"] = nil
	} else {
		inputs["customAttributes"] = args.CustomAttributes
		inputs["datacenterId"] = args.DatacenterId
		inputs["path"] = args.Path
		inputs["tags"] = args.Tags
		inputs["type"] = args.Type
	}
	s, err := ctx.RegisterResource("vsphere:index/folder:Folder", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Folder{s: s}, nil
}

// GetFolder gets an existing Folder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolder(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FolderState, opts ...pulumi.ResourceOpt) (*Folder, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["customAttributes"] = state.CustomAttributes
		inputs["datacenterId"] = state.DatacenterId
		inputs["path"] = state.Path
		inputs["tags"] = state.Tags
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("vsphere:index/folder:Folder", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Folder{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Folder) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Folder) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Map of custom attribute ids to attribute
// value strings to set for folder. See [here][docs-setting-custom-attributes]
// for a reference on how to set values for custom attributes.
func (r *Folder) CustomAttributes() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["customAttributes"])
}

// The ID of the datacenter the folder will be created in.
// Required for all folder types except for datacenter folders. Forces a new
// resource if changed.
func (r *Folder) DatacenterId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["datacenterId"])
}

// The path of the folder and any parents, relative to the datacenter and folder type being defined.
func (r *Folder) Path() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["path"])
}

// The IDs of any tags to attach to this resource. See
// [here][docs-applying-tags] for a reference on how to apply tags.
func (r *Folder) Tags() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["tags"])
}

// The type of folder to create. Allowed options are
// `datacenter` for datacenter folders, `host` for host and cluster folders,
// `vm` for virtual machine folders, `datastore` for datastore folders, and
// `network` for network folders. Forces a new resource if changed.
func (r *Folder) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering Folder resources.
type FolderState struct {
	// Map of custom attribute ids to attribute
	// value strings to set for folder. See [here][docs-setting-custom-attributes]
	// for a reference on how to set values for custom attributes.
	CustomAttributes interface{}
	// The ID of the datacenter the folder will be created in.
	// Required for all folder types except for datacenter folders. Forces a new
	// resource if changed.
	DatacenterId interface{}
	// The path of the folder and any parents, relative to the datacenter and folder type being defined.
	Path interface{}
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
	Tags interface{}
	// The type of folder to create. Allowed options are
	// `datacenter` for datacenter folders, `host` for host and cluster folders,
	// `vm` for virtual machine folders, `datastore` for datastore folders, and
	// `network` for network folders. Forces a new resource if changed.
	Type interface{}
}

// The set of arguments for constructing a Folder resource.
type FolderArgs struct {
	// Map of custom attribute ids to attribute
	// value strings to set for folder. See [here][docs-setting-custom-attributes]
	// for a reference on how to set values for custom attributes.
	CustomAttributes interface{}
	// The ID of the datacenter the folder will be created in.
	// Required for all folder types except for datacenter folders. Forces a new
	// resource if changed.
	DatacenterId interface{}
	// The path of the folder and any parents, relative to the datacenter and folder type being defined.
	Path interface{}
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
	Tags interface{}
	// The type of folder to create. Allowed options are
	// `datacenter` for datacenter folders, `host` for host and cluster folders,
	// `vm` for virtual machine folders, `datastore` for datastore folders, and
	// `network` for network folders. Forces a new resource if changed.
	Type interface{}
}
