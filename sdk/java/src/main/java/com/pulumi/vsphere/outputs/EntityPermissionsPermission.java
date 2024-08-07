// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class EntityPermissionsPermission {
    /**
     * @return Whether `user_or_group` field refers to a user or a
     * group. True for a group and false for a user.
     * 
     */
    private Boolean isGroup;
    /**
     * @return Whether or not this permission propagates down the
     * hierarchy to sub-entities.
     * 
     */
    private Boolean propagate;
    /**
     * @return The role id of the role to be given to the user on
     * the specified entity.
     * 
     */
    private String roleId;
    /**
     * @return The user/group getting the permission.
     * 
     */
    private String userOrGroup;

    private EntityPermissionsPermission() {}
    /**
     * @return Whether `user_or_group` field refers to a user or a
     * group. True for a group and false for a user.
     * 
     */
    public Boolean isGroup() {
        return this.isGroup;
    }
    /**
     * @return Whether or not this permission propagates down the
     * hierarchy to sub-entities.
     * 
     */
    public Boolean propagate() {
        return this.propagate;
    }
    /**
     * @return The role id of the role to be given to the user on
     * the specified entity.
     * 
     */
    public String roleId() {
        return this.roleId;
    }
    /**
     * @return The user/group getting the permission.
     * 
     */
    public String userOrGroup() {
        return this.userOrGroup;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EntityPermissionsPermission defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean isGroup;
        private Boolean propagate;
        private String roleId;
        private String userOrGroup;
        public Builder() {}
        public Builder(EntityPermissionsPermission defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.isGroup = defaults.isGroup;
    	      this.propagate = defaults.propagate;
    	      this.roleId = defaults.roleId;
    	      this.userOrGroup = defaults.userOrGroup;
        }

        @CustomType.Setter
        public Builder isGroup(Boolean isGroup) {
            if (isGroup == null) {
              throw new MissingRequiredPropertyException("EntityPermissionsPermission", "isGroup");
            }
            this.isGroup = isGroup;
            return this;
        }
        @CustomType.Setter
        public Builder propagate(Boolean propagate) {
            if (propagate == null) {
              throw new MissingRequiredPropertyException("EntityPermissionsPermission", "propagate");
            }
            this.propagate = propagate;
            return this;
        }
        @CustomType.Setter
        public Builder roleId(String roleId) {
            if (roleId == null) {
              throw new MissingRequiredPropertyException("EntityPermissionsPermission", "roleId");
            }
            this.roleId = roleId;
            return this;
        }
        @CustomType.Setter
        public Builder userOrGroup(String userOrGroup) {
            if (userOrGroup == null) {
              throw new MissingRequiredPropertyException("EntityPermissionsPermission", "userOrGroup");
            }
            this.userOrGroup = userOrGroup;
            return this;
        }
        public EntityPermissionsPermission build() {
            final var _resultValue = new EntityPermissionsPermission();
            _resultValue.isGroup = isGroup;
            _resultValue.propagate = propagate;
            _resultValue.roleId = roleId;
            _resultValue.userOrGroup = userOrGroup;
            return _resultValue;
        }
    }
}
