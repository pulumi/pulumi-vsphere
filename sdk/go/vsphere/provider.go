// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the vsphere package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	if args.AllowUnverifiedSsl == nil {
		args.AllowUnverifiedSsl = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "VSPHERE_ALLOW_UNVERIFIED_SSL").(bool))
	}
	if args.ClientDebug == nil {
		args.ClientDebug = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "VSPHERE_CLIENT_DEBUG").(bool))
	}
	if args.ClientDebugPath == nil {
		args.ClientDebugPath = pulumi.StringPtr(getEnvOrDefault("", nil, "VSPHERE_CLIENT_DEBUG_PATH").(string))
	}
	if args.ClientDebugPathRun == nil {
		args.ClientDebugPathRun = pulumi.StringPtr(getEnvOrDefault("", nil, "VSPHERE_CLIENT_DEBUG_PATH_RUN").(string))
	}
	if args.Password == nil {
		args.Password = pulumi.StringPtr(getEnvOrDefault("", nil, "VSPHERE_PASSWORD").(string))
	}
	if args.PersistSession == nil {
		args.PersistSession = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "VSPHERE_PERSIST_SESSION").(bool))
	}
	if args.RestSessionPath == nil {
		args.RestSessionPath = pulumi.StringPtr(getEnvOrDefault("", nil, "VSPHERE_REST_SESSION_PATH").(string))
	}
	if args.User == nil {
		args.User = pulumi.StringPtr(getEnvOrDefault("", nil, "VSPHERE_USER").(string))
	}
	if args.VimKeepAlive == nil {
		args.VimKeepAlive = pulumi.IntPtr(getEnvOrDefault(0, parseEnvInt, "VSPHERE_VIM_KEEP_ALIVE").(int))
	}
	if args.VimSessionPath == nil {
		args.VimSessionPath = pulumi.StringPtr(getEnvOrDefault("", nil, "VSPHERE_VIM_SESSION_PATH").(string))
	}
	if args.VsphereServer == nil {
		args.VsphereServer = pulumi.StringPtr(getEnvOrDefault("", nil, "VSPHERE_SERVER").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:vsphere", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// If set, VMware vSphere client will permit unverifiable SSL certificates.
	AllowUnverifiedSsl *bool `pulumi:"allowUnverifiedSsl"`
	// API timeout in minutes (Default: 5)
	ApiTimeout *int `pulumi:"apiTimeout"`
	// govmomi debug
	ClientDebug *bool `pulumi:"clientDebug"`
	// govmomi debug path for debug
	ClientDebugPath *string `pulumi:"clientDebugPath"`
	// govmomi debug path for a single run
	ClientDebugPathRun *string `pulumi:"clientDebugPathRun"`
	// The user password for vSphere API operations.
	Password *string `pulumi:"password"`
	// Persist vSphere client sessions to disk
	PersistSession *bool `pulumi:"persistSession"`
	// The directory to save vSphere REST API sessions to
	RestSessionPath *string `pulumi:"restSessionPath"`
	// The user name for vSphere API operations.
	User *string `pulumi:"user"`
	// Deprecated: This field has been renamed to vsphere_server.
	VcenterServer *string `pulumi:"vcenterServer"`
	// Keep alive interval for the VIM session in minutes
	VimKeepAlive *int `pulumi:"vimKeepAlive"`
	// The directory to save vSphere SOAP API sessions to
	VimSessionPath *string `pulumi:"vimSessionPath"`
	// The vSphere Server name for vSphere API operations.
	VsphereServer *string `pulumi:"vsphereServer"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// If set, VMware vSphere client will permit unverifiable SSL certificates.
	AllowUnverifiedSsl pulumi.BoolPtrInput
	// API timeout in minutes (Default: 5)
	ApiTimeout pulumi.IntPtrInput
	// govmomi debug
	ClientDebug pulumi.BoolPtrInput
	// govmomi debug path for debug
	ClientDebugPath pulumi.StringPtrInput
	// govmomi debug path for a single run
	ClientDebugPathRun pulumi.StringPtrInput
	// The user password for vSphere API operations.
	Password pulumi.StringPtrInput
	// Persist vSphere client sessions to disk
	PersistSession pulumi.BoolPtrInput
	// The directory to save vSphere REST API sessions to
	RestSessionPath pulumi.StringPtrInput
	// The user name for vSphere API operations.
	User pulumi.StringPtrInput
	// Deprecated: This field has been renamed to vsphere_server.
	VcenterServer pulumi.StringPtrInput
	// Keep alive interval for the VIM session in minutes
	VimKeepAlive pulumi.IntPtrInput
	// The directory to save vSphere SOAP API sessions to
	VimSessionPath pulumi.StringPtrInput
	// The vSphere Server name for vSphere API operations.
	VsphereServer pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (Provider) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil)).Elem()
}

func (i Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct {
	*pulumi.OutputState
}

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderOutput)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProviderOutput{})
}
