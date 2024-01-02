// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRolePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRolePlainArgs Empty = new GetRolePlainArgs();

    /**
     * The description of the role.
     * 
     */
    @Import(name="description")
    private @Nullable String description;

    /**
     * @return The description of the role.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The label of the role.
     * 
     */
    @Import(name="label", required=true)
    private String label;

    /**
     * @return The label of the role.
     * 
     */
    public String label() {
        return this.label;
    }

    @Import(name="name")
    private @Nullable String name;

    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The privileges associated with the role.
     * 
     */
    @Import(name="rolePrivileges")
    private @Nullable List<String> rolePrivileges;

    /**
     * @return The privileges associated with the role.
     * 
     */
    public Optional<List<String>> rolePrivileges() {
        return Optional.ofNullable(this.rolePrivileges);
    }

    private GetRolePlainArgs() {}

    private GetRolePlainArgs(GetRolePlainArgs $) {
        this.description = $.description;
        this.label = $.label;
        this.name = $.name;
        this.rolePrivileges = $.rolePrivileges;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRolePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRolePlainArgs $;

        public Builder() {
            $ = new GetRolePlainArgs();
        }

        public Builder(GetRolePlainArgs defaults) {
            $ = new GetRolePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the role.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        /**
         * @param label The label of the role.
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            $.label = label;
            return this;
        }

        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param rolePrivileges The privileges associated with the role.
         * 
         * @return builder
         * 
         */
        public Builder rolePrivileges(@Nullable List<String> rolePrivileges) {
            $.rolePrivileges = rolePrivileges;
            return this;
        }

        /**
         * @param rolePrivileges The privileges associated with the role.
         * 
         * @return builder
         * 
         */
        public Builder rolePrivileges(String... rolePrivileges) {
            return rolePrivileges(List.of(rolePrivileges));
        }

        public GetRolePlainArgs build() {
            if ($.label == null) {
                throw new MissingRequiredPropertyException("GetRolePlainArgs", "label");
            }
            return $;
        }
    }

}
