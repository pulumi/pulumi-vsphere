// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class EntityPermissionsPermissionArgs extends com.pulumi.resources.ResourceArgs {

    public static final EntityPermissionsPermissionArgs Empty = new EntityPermissionsPermissionArgs();

    /**
     * Whether user_or_group field refers to a user or a group. True for a group and false for a user.
     * 
     */
    @Import(name="isGroup", required=true)
    private Output<Boolean> isGroup;

    /**
     * @return Whether user_or_group field refers to a user or a group. True for a group and false for a user.
     * 
     */
    public Output<Boolean> isGroup() {
        return this.isGroup;
    }

    /**
     * Whether or not this permission propagates down the hierarchy to sub-entities.
     * 
     */
    @Import(name="propagate", required=true)
    private Output<Boolean> propagate;

    /**
     * @return Whether or not this permission propagates down the hierarchy to sub-entities.
     * 
     */
    public Output<Boolean> propagate() {
        return this.propagate;
    }

    /**
     * The role id of the role to be given to the user on the specified entity.
     * 
     */
    @Import(name="roleId", required=true)
    private Output<String> roleId;

    /**
     * @return The role id of the role to be given to the user on the specified entity.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }

    /**
     * The user/group getting the permission.
     * 
     */
    @Import(name="userOrGroup", required=true)
    private Output<String> userOrGroup;

    /**
     * @return The user/group getting the permission.
     * 
     */
    public Output<String> userOrGroup() {
        return this.userOrGroup;
    }

    private EntityPermissionsPermissionArgs() {}

    private EntityPermissionsPermissionArgs(EntityPermissionsPermissionArgs $) {
        this.isGroup = $.isGroup;
        this.propagate = $.propagate;
        this.roleId = $.roleId;
        this.userOrGroup = $.userOrGroup;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EntityPermissionsPermissionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EntityPermissionsPermissionArgs $;

        public Builder() {
            $ = new EntityPermissionsPermissionArgs();
        }

        public Builder(EntityPermissionsPermissionArgs defaults) {
            $ = new EntityPermissionsPermissionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param isGroup Whether user_or_group field refers to a user or a group. True for a group and false for a user.
         * 
         * @return builder
         * 
         */
        public Builder isGroup(Output<Boolean> isGroup) {
            $.isGroup = isGroup;
            return this;
        }

        /**
         * @param isGroup Whether user_or_group field refers to a user or a group. True for a group and false for a user.
         * 
         * @return builder
         * 
         */
        public Builder isGroup(Boolean isGroup) {
            return isGroup(Output.of(isGroup));
        }

        /**
         * @param propagate Whether or not this permission propagates down the hierarchy to sub-entities.
         * 
         * @return builder
         * 
         */
        public Builder propagate(Output<Boolean> propagate) {
            $.propagate = propagate;
            return this;
        }

        /**
         * @param propagate Whether or not this permission propagates down the hierarchy to sub-entities.
         * 
         * @return builder
         * 
         */
        public Builder propagate(Boolean propagate) {
            return propagate(Output.of(propagate));
        }

        /**
         * @param roleId The role id of the role to be given to the user on the specified entity.
         * 
         * @return builder
         * 
         */
        public Builder roleId(Output<String> roleId) {
            $.roleId = roleId;
            return this;
        }

        /**
         * @param roleId The role id of the role to be given to the user on the specified entity.
         * 
         * @return builder
         * 
         */
        public Builder roleId(String roleId) {
            return roleId(Output.of(roleId));
        }

        /**
         * @param userOrGroup The user/group getting the permission.
         * 
         * @return builder
         * 
         */
        public Builder userOrGroup(Output<String> userOrGroup) {
            $.userOrGroup = userOrGroup;
            return this;
        }

        /**
         * @param userOrGroup The user/group getting the permission.
         * 
         * @return builder
         * 
         */
        public Builder userOrGroup(String userOrGroup) {
            return userOrGroup(Output.of(userOrGroup));
        }

        public EntityPermissionsPermissionArgs build() {
            if ($.isGroup == null) {
                throw new MissingRequiredPropertyException("EntityPermissionsPermissionArgs", "isGroup");
            }
            if ($.propagate == null) {
                throw new MissingRequiredPropertyException("EntityPermissionsPermissionArgs", "propagate");
            }
            if ($.roleId == null) {
                throw new MissingRequiredPropertyException("EntityPermissionsPermissionArgs", "roleId");
            }
            if ($.userOrGroup == null) {
                throw new MissingRequiredPropertyException("EntityPermissionsPermissionArgs", "userOrGroup");
            }
            return $;
        }
    }

}
