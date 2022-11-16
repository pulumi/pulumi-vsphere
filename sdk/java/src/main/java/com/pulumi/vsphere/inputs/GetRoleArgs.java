// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRoleArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRoleArgs Empty = new GetRoleArgs();

    /**
     * The description of the role.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the role.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The label of the role.
     * 
     */
    @Import(name="label", required=true)
    private Output<String> label;

    /**
     * @return The label of the role.
     * 
     */
    public Output<String> label() {
        return this.label;
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The privileges associated with the role.
     * 
     */
    @Import(name="rolePrivileges")
    private @Nullable Output<List<String>> rolePrivileges;

    /**
     * @return The privileges associated with the role.
     * 
     */
    public Optional<Output<List<String>>> rolePrivileges() {
        return Optional.ofNullable(this.rolePrivileges);
    }

    private GetRoleArgs() {}

    private GetRoleArgs(GetRoleArgs $) {
        this.description = $.description;
        this.label = $.label;
        this.name = $.name;
        this.rolePrivileges = $.rolePrivileges;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRoleArgs $;

        public Builder() {
            $ = new GetRoleArgs();
        }

        public Builder(GetRoleArgs defaults) {
            $ = new GetRoleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the role.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the role.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param label The label of the role.
         * 
         * @return builder
         * 
         */
        public Builder label(Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label The label of the role.
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param rolePrivileges The privileges associated with the role.
         * 
         * @return builder
         * 
         */
        public Builder rolePrivileges(@Nullable Output<List<String>> rolePrivileges) {
            $.rolePrivileges = rolePrivileges;
            return this;
        }

        /**
         * @param rolePrivileges The privileges associated with the role.
         * 
         * @return builder
         * 
         */
        public Builder rolePrivileges(List<String> rolePrivileges) {
            return rolePrivileges(Output.of(rolePrivileges));
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

        public GetRoleArgs build() {
            $.label = Objects.requireNonNull($.label, "expected parameter 'label' to be non-null");
            return $;
        }
    }

}
