// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VappEntityArgs extends com.pulumi.resources.ResourceArgs {

    public static final VappEntityArgs Empty = new VappEntityArgs();

    /**
     * Managed object ID of the vApp
     * container the entity is a member of.
     * 
     */
    @Import(name="containerId", required=true)
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
    @Import(name="customAttributes")
    private @Nullable Output<Map<String,String>> customAttributes;

    /**
     * @return A list of custom attributes to set on this resource.
     * 
     */
    public Optional<Output<Map<String,String>>> customAttributes() {
        return Optional.ofNullable(this.customAttributes);
    }

    /**
     * How to start the entity. Valid settings are none
     * or powerOn. If set to none, then the entity does not participate in auto-start.
     * Default: powerOn
     * 
     */
    @Import(name="startAction")
    private @Nullable Output<String> startAction;

    /**
     * @return How to start the entity. Valid settings are none
     * or powerOn. If set to none, then the entity does not participate in auto-start.
     * Default: powerOn
     * 
     */
    public Optional<Output<String>> startAction() {
        return Optional.ofNullable(this.startAction);
    }

    /**
     * Delay in seconds before continuing with the next
     * entity in the order of entities to be started. Default: 120
     * 
     */
    @Import(name="startDelay")
    private @Nullable Output<Integer> startDelay;

    /**
     * @return Delay in seconds before continuing with the next
     * entity in the order of entities to be started. Default: 120
     * 
     */
    public Optional<Output<Integer>> startDelay() {
        return Optional.ofNullable(this.startDelay);
    }

    /**
     * Order to start and stop target in vApp. Default: 1
     * 
     */
    @Import(name="startOrder")
    private @Nullable Output<Integer> startOrder;

    /**
     * @return Order to start and stop target in vApp. Default: 1
     * 
     */
    public Optional<Output<Integer>> startOrder() {
        return Optional.ofNullable(this.startOrder);
    }

    /**
     * Defines the stop action for the entity. Can be set
     * to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
     * does not participate in auto-stop. Default: powerOff
     * 
     */
    @Import(name="stopAction")
    private @Nullable Output<String> stopAction;

    /**
     * @return Defines the stop action for the entity. Can be set
     * to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
     * does not participate in auto-stop. Default: powerOff
     * 
     */
    public Optional<Output<String>> stopAction() {
        return Optional.ofNullable(this.stopAction);
    }

    /**
     * Delay in seconds before continuing with the next
     * entity in the order sequence. This is only used if the stopAction is
     * guestShutdown. Default: 120
     * 
     */
    @Import(name="stopDelay")
    private @Nullable Output<Integer> stopDelay;

    /**
     * @return Delay in seconds before continuing with the next
     * entity in the order sequence. This is only used if the stopAction is
     * guestShutdown. Default: 120
     * 
     */
    public Optional<Output<Integer>> stopDelay() {
        return Optional.ofNullable(this.stopDelay);
    }

    /**
     * A list of tag IDs to apply to this object.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A list of tag IDs to apply to this object.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Managed object ID of the entity
     * to power on or power off. This can be a virtual machine or a vApp.
     * 
     */
    @Import(name="targetId", required=true)
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
    @Import(name="waitForGuest")
    private @Nullable Output<Boolean> waitForGuest;

    /**
     * @return Determines if the VM should be marked as being
     * started when VMware Tools are ready instead of waiting for `start_delay`. This
     * property has no effect for vApps. Default: false
     * 
     */
    public Optional<Output<Boolean>> waitForGuest() {
        return Optional.ofNullable(this.waitForGuest);
    }

    private VappEntityArgs() {}

    private VappEntityArgs(VappEntityArgs $) {
        this.containerId = $.containerId;
        this.customAttributes = $.customAttributes;
        this.startAction = $.startAction;
        this.startDelay = $.startDelay;
        this.startOrder = $.startOrder;
        this.stopAction = $.stopAction;
        this.stopDelay = $.stopDelay;
        this.tags = $.tags;
        this.targetId = $.targetId;
        this.waitForGuest = $.waitForGuest;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VappEntityArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VappEntityArgs $;

        public Builder() {
            $ = new VappEntityArgs();
        }

        public Builder(VappEntityArgs defaults) {
            $ = new VappEntityArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param containerId Managed object ID of the vApp
         * container the entity is a member of.
         * 
         * @return builder
         * 
         */
        public Builder containerId(Output<String> containerId) {
            $.containerId = containerId;
            return this;
        }

        /**
         * @param containerId Managed object ID of the vApp
         * container the entity is a member of.
         * 
         * @return builder
         * 
         */
        public Builder containerId(String containerId) {
            return containerId(Output.of(containerId));
        }

        /**
         * @param customAttributes A list of custom attributes to set on this resource.
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(@Nullable Output<Map<String,String>> customAttributes) {
            $.customAttributes = customAttributes;
            return this;
        }

        /**
         * @param customAttributes A list of custom attributes to set on this resource.
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(Map<String,String> customAttributes) {
            return customAttributes(Output.of(customAttributes));
        }

        /**
         * @param startAction How to start the entity. Valid settings are none
         * or powerOn. If set to none, then the entity does not participate in auto-start.
         * Default: powerOn
         * 
         * @return builder
         * 
         */
        public Builder startAction(@Nullable Output<String> startAction) {
            $.startAction = startAction;
            return this;
        }

        /**
         * @param startAction How to start the entity. Valid settings are none
         * or powerOn. If set to none, then the entity does not participate in auto-start.
         * Default: powerOn
         * 
         * @return builder
         * 
         */
        public Builder startAction(String startAction) {
            return startAction(Output.of(startAction));
        }

        /**
         * @param startDelay Delay in seconds before continuing with the next
         * entity in the order of entities to be started. Default: 120
         * 
         * @return builder
         * 
         */
        public Builder startDelay(@Nullable Output<Integer> startDelay) {
            $.startDelay = startDelay;
            return this;
        }

        /**
         * @param startDelay Delay in seconds before continuing with the next
         * entity in the order of entities to be started. Default: 120
         * 
         * @return builder
         * 
         */
        public Builder startDelay(Integer startDelay) {
            return startDelay(Output.of(startDelay));
        }

        /**
         * @param startOrder Order to start and stop target in vApp. Default: 1
         * 
         * @return builder
         * 
         */
        public Builder startOrder(@Nullable Output<Integer> startOrder) {
            $.startOrder = startOrder;
            return this;
        }

        /**
         * @param startOrder Order to start and stop target in vApp. Default: 1
         * 
         * @return builder
         * 
         */
        public Builder startOrder(Integer startOrder) {
            return startOrder(Output.of(startOrder));
        }

        /**
         * @param stopAction Defines the stop action for the entity. Can be set
         * to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
         * does not participate in auto-stop. Default: powerOff
         * 
         * @return builder
         * 
         */
        public Builder stopAction(@Nullable Output<String> stopAction) {
            $.stopAction = stopAction;
            return this;
        }

        /**
         * @param stopAction Defines the stop action for the entity. Can be set
         * to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
         * does not participate in auto-stop. Default: powerOff
         * 
         * @return builder
         * 
         */
        public Builder stopAction(String stopAction) {
            return stopAction(Output.of(stopAction));
        }

        /**
         * @param stopDelay Delay in seconds before continuing with the next
         * entity in the order sequence. This is only used if the stopAction is
         * guestShutdown. Default: 120
         * 
         * @return builder
         * 
         */
        public Builder stopDelay(@Nullable Output<Integer> stopDelay) {
            $.stopDelay = stopDelay;
            return this;
        }

        /**
         * @param stopDelay Delay in seconds before continuing with the next
         * entity in the order sequence. This is only used if the stopAction is
         * guestShutdown. Default: 120
         * 
         * @return builder
         * 
         */
        public Builder stopDelay(Integer stopDelay) {
            return stopDelay(Output.of(stopDelay));
        }

        /**
         * @param tags A list of tag IDs to apply to this object.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A list of tag IDs to apply to this object.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A list of tag IDs to apply to this object.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param targetId Managed object ID of the entity
         * to power on or power off. This can be a virtual machine or a vApp.
         * 
         * @return builder
         * 
         */
        public Builder targetId(Output<String> targetId) {
            $.targetId = targetId;
            return this;
        }

        /**
         * @param targetId Managed object ID of the entity
         * to power on or power off. This can be a virtual machine or a vApp.
         * 
         * @return builder
         * 
         */
        public Builder targetId(String targetId) {
            return targetId(Output.of(targetId));
        }

        /**
         * @param waitForGuest Determines if the VM should be marked as being
         * started when VMware Tools are ready instead of waiting for `start_delay`. This
         * property has no effect for vApps. Default: false
         * 
         * @return builder
         * 
         */
        public Builder waitForGuest(@Nullable Output<Boolean> waitForGuest) {
            $.waitForGuest = waitForGuest;
            return this;
        }

        /**
         * @param waitForGuest Determines if the VM should be marked as being
         * started when VMware Tools are ready instead of waiting for `start_delay`. This
         * property has no effect for vApps. Default: false
         * 
         * @return builder
         * 
         */
        public Builder waitForGuest(Boolean waitForGuest) {
            return waitForGuest(Output.of(waitForGuest));
        }

        public VappEntityArgs build() {
            if ($.containerId == null) {
                throw new MissingRequiredPropertyException("VappEntityArgs", "containerId");
            }
            if ($.targetId == null) {
                throw new MissingRequiredPropertyException("VappEntityArgs", "targetId");
            }
            return $;
        }
    }

}
