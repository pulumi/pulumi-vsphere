// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/vapp_entity.html.markdown.
 */
export class VappEntity extends pulumi.CustomResource {
    /**
     * Get an existing VappEntity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VappEntityState, opts?: pulumi.CustomResourceOptions): VappEntity {
        return new VappEntity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/vappEntity:VappEntity';

    /**
     * Returns true if the given object is an instance of VappEntity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VappEntity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VappEntity.__pulumiType;
    }

    /**
     * [Managed object ID|docs-about-morefs] of the vApp
     * container the entity is a member of.
     */
    public readonly containerId!: pulumi.Output<string>;
    /**
     * A list of custom attributes to set on this resource.
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * How to start the entity. Valid settings are none
     * or powerOn. If set to none, then the entity does not participate in auto-start.
     * Default: powerOn
     */
    public readonly startAction!: pulumi.Output<string | undefined>;
    /**
     * Delay in seconds before continuing with the next
     * entity in the order of entities to be started. Default: 120
     */
    public readonly startDelay!: pulumi.Output<number | undefined>;
    /**
     * Order to start and stop target in vApp. Default: 1
     */
    public readonly startOrder!: pulumi.Output<number | undefined>;
    /**
     * Defines the stop action for the entity. Can be set
     * to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
     * does not participate in auto-stop. Default: powerOff
     */
    public readonly stopAction!: pulumi.Output<string | undefined>;
    /**
     * Delay in seconds before continuing with the next
     * entity in the order sequence. This is only used if the stopAction is
     * guestShutdown. Default: 120
     */
    public readonly stopDelay!: pulumi.Output<number | undefined>;
    /**
     * A list of tag IDs to apply to this object.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * [Managed object ID|docs-about-morefs] of the entity
     * to power on or power off. This can be a virtual machine or a vApp.
     */
    public readonly targetId!: pulumi.Output<string>;
    /**
     * Determines if the VM should be marked as being
     * started when VMware Tools are ready instead of waiting for `startDelay`. This
     * property has no effect for vApps. Default: false
     */
    public readonly waitForGuest!: pulumi.Output<boolean | undefined>;

    /**
     * Create a VappEntity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VappEntityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VappEntityArgs | VappEntityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VappEntityState | undefined;
            inputs["containerId"] = state ? state.containerId : undefined;
            inputs["customAttributes"] = state ? state.customAttributes : undefined;
            inputs["startAction"] = state ? state.startAction : undefined;
            inputs["startDelay"] = state ? state.startDelay : undefined;
            inputs["startOrder"] = state ? state.startOrder : undefined;
            inputs["stopAction"] = state ? state.stopAction : undefined;
            inputs["stopDelay"] = state ? state.stopDelay : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["targetId"] = state ? state.targetId : undefined;
            inputs["waitForGuest"] = state ? state.waitForGuest : undefined;
        } else {
            const args = argsOrState as VappEntityArgs | undefined;
            if (!args || args.containerId === undefined) {
                throw new Error("Missing required property 'containerId'");
            }
            if (!args || args.targetId === undefined) {
                throw new Error("Missing required property 'targetId'");
            }
            inputs["containerId"] = args ? args.containerId : undefined;
            inputs["customAttributes"] = args ? args.customAttributes : undefined;
            inputs["startAction"] = args ? args.startAction : undefined;
            inputs["startDelay"] = args ? args.startDelay : undefined;
            inputs["startOrder"] = args ? args.startOrder : undefined;
            inputs["stopAction"] = args ? args.stopAction : undefined;
            inputs["stopDelay"] = args ? args.stopDelay : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetId"] = args ? args.targetId : undefined;
            inputs["waitForGuest"] = args ? args.waitForGuest : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VappEntity.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VappEntity resources.
 */
export interface VappEntityState {
    /**
     * [Managed object ID|docs-about-morefs] of the vApp
     * container the entity is a member of.
     */
    readonly containerId?: pulumi.Input<string>;
    /**
     * A list of custom attributes to set on this resource.
     */
    readonly customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * How to start the entity. Valid settings are none
     * or powerOn. If set to none, then the entity does not participate in auto-start.
     * Default: powerOn
     */
    readonly startAction?: pulumi.Input<string>;
    /**
     * Delay in seconds before continuing with the next
     * entity in the order of entities to be started. Default: 120
     */
    readonly startDelay?: pulumi.Input<number>;
    /**
     * Order to start and stop target in vApp. Default: 1
     */
    readonly startOrder?: pulumi.Input<number>;
    /**
     * Defines the stop action for the entity. Can be set
     * to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
     * does not participate in auto-stop. Default: powerOff
     */
    readonly stopAction?: pulumi.Input<string>;
    /**
     * Delay in seconds before continuing with the next
     * entity in the order sequence. This is only used if the stopAction is
     * guestShutdown. Default: 120
     */
    readonly stopDelay?: pulumi.Input<number>;
    /**
     * A list of tag IDs to apply to this object.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * [Managed object ID|docs-about-morefs] of the entity
     * to power on or power off. This can be a virtual machine or a vApp.
     */
    readonly targetId?: pulumi.Input<string>;
    /**
     * Determines if the VM should be marked as being
     * started when VMware Tools are ready instead of waiting for `startDelay`. This
     * property has no effect for vApps. Default: false
     */
    readonly waitForGuest?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a VappEntity resource.
 */
export interface VappEntityArgs {
    /**
     * [Managed object ID|docs-about-morefs] of the vApp
     * container the entity is a member of.
     */
    readonly containerId: pulumi.Input<string>;
    /**
     * A list of custom attributes to set on this resource.
     */
    readonly customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * How to start the entity. Valid settings are none
     * or powerOn. If set to none, then the entity does not participate in auto-start.
     * Default: powerOn
     */
    readonly startAction?: pulumi.Input<string>;
    /**
     * Delay in seconds before continuing with the next
     * entity in the order of entities to be started. Default: 120
     */
    readonly startDelay?: pulumi.Input<number>;
    /**
     * Order to start and stop target in vApp. Default: 1
     */
    readonly startOrder?: pulumi.Input<number>;
    /**
     * Defines the stop action for the entity. Can be set
     * to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
     * does not participate in auto-stop. Default: powerOff
     */
    readonly stopAction?: pulumi.Input<string>;
    /**
     * Delay in seconds before continuing with the next
     * entity in the order sequence. This is only used if the stopAction is
     * guestShutdown. Default: 120
     */
    readonly stopDelay?: pulumi.Input<number>;
    /**
     * A list of tag IDs to apply to this object.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * [Managed object ID|docs-about-morefs] of the entity
     * to power on or power off. This can be a virtual machine or a vApp.
     */
    readonly targetId: pulumi.Input<string>;
    /**
     * Determines if the VM should be marked as being
     * started when VMware Tools are ready instead of waiting for `startDelay`. This
     * property has no effect for vApps. Default: false
     */
    readonly waitForGuest?: pulumi.Input<boolean>;
}
