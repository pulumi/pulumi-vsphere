// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.VappEntityArgs;
import com.pulumi.vsphere.inputs.VappEntityState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="vsphere:index/vappEntity:VappEntity")
public class VappEntity extends com.pulumi.resources.CustomResource {
    /**
     * Managed object ID of the vApp
     * container the entity is a member of.
     * 
     */
    @Export(name="containerId", type=String.class, parameters={})
    private Output<String> containerId;

    /**
     * @return Managed object ID of the vApp
     * container the entity is a member of.
     * 
     */
    public Output<String> containerId() {
        return this.containerId;
    }
    /**
     * A list of custom attributes to set on this resource.
     * 
     */
    @Export(name="customAttributes", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> customAttributes;

    /**
     * @return A list of custom attributes to set on this resource.
     * 
     */
    public Output<Optional<Map<String,String>>> customAttributes() {
        return Codegen.optional(this.customAttributes);
    }
    /**
     * How to start the entity. Valid settings are none
     * or powerOn. If set to none, then the entity does not participate in auto-start.
     * Default: powerOn
     * 
     */
    @Export(name="startAction", type=String.class, parameters={})
    private Output</* @Nullable */ String> startAction;

    /**
     * @return How to start the entity. Valid settings are none
     * or powerOn. If set to none, then the entity does not participate in auto-start.
     * Default: powerOn
     * 
     */
    public Output<Optional<String>> startAction() {
        return Codegen.optional(this.startAction);
    }
    /**
     * Delay in seconds before continuing with the next
     * entity in the order of entities to be started. Default: 120
     * 
     */
    @Export(name="startDelay", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> startDelay;

    /**
     * @return Delay in seconds before continuing with the next
     * entity in the order of entities to be started. Default: 120
     * 
     */
    public Output<Optional<Integer>> startDelay() {
        return Codegen.optional(this.startDelay);
    }
    /**
     * Order to start and stop target in vApp. Default: 1
     * 
     */
    @Export(name="startOrder", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> startOrder;

    /**
     * @return Order to start and stop target in vApp. Default: 1
     * 
     */
    public Output<Optional<Integer>> startOrder() {
        return Codegen.optional(this.startOrder);
    }
    /**
     * Defines the stop action for the entity. Can be set
     * to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
     * does not participate in auto-stop. Default: powerOff
     * 
     */
    @Export(name="stopAction", type=String.class, parameters={})
    private Output</* @Nullable */ String> stopAction;

    /**
     * @return Defines the stop action for the entity. Can be set
     * to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
     * does not participate in auto-stop. Default: powerOff
     * 
     */
    public Output<Optional<String>> stopAction() {
        return Codegen.optional(this.stopAction);
    }
    /**
     * Delay in seconds before continuing with the next
     * entity in the order sequence. This is only used if the stopAction is
     * guestShutdown. Default: 120
     * 
     */
    @Export(name="stopDelay", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> stopDelay;

    /**
     * @return Delay in seconds before continuing with the next
     * entity in the order sequence. This is only used if the stopAction is
     * guestShutdown. Default: 120
     * 
     */
    public Output<Optional<Integer>> stopDelay() {
        return Codegen.optional(this.stopDelay);
    }
    /**
     * A list of tag IDs to apply to this object.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A list of tag IDs to apply to this object.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Managed object ID of the entity
     * to power on or power off. This can be a virtual machine or a vApp.
     * 
     */
    @Export(name="targetId", type=String.class, parameters={})
    private Output<String> targetId;

    /**
     * @return Managed object ID of the entity
     * to power on or power off. This can be a virtual machine or a vApp.
     * 
     */
    public Output<String> targetId() {
        return this.targetId;
    }
    /**
     * Determines if the VM should be marked as being
     * started when VMware Tools are ready instead of waiting for `start_delay`. This
     * property has no effect for vApps. Default: false
     * 
     */
    @Export(name="waitForGuest", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> waitForGuest;

    /**
     * @return Determines if the VM should be marked as being
     * started when VMware Tools are ready instead of waiting for `start_delay`. This
     * property has no effect for vApps. Default: false
     * 
     */
    public Output<Optional<Boolean>> waitForGuest() {
        return Codegen.optional(this.waitForGuest);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VappEntity(String name) {
        this(name, VappEntityArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VappEntity(String name, VappEntityArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VappEntity(String name, VappEntityArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/vappEntity:VappEntity", name, args == null ? VappEntityArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VappEntity(String name, Output<String> id, @Nullable VappEntityState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/vappEntity:VappEntity", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static VappEntity get(String name, Output<String> id, @Nullable VappEntityState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VappEntity(name, id, state, options);
    }
}
