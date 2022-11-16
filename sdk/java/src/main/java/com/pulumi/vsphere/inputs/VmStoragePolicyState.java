// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.vsphere.inputs.VmStoragePolicyTagRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VmStoragePolicyState extends com.pulumi.resources.ResourceArgs {

    public static final VmStoragePolicyState Empty = new VmStoragePolicyState();

    /**
     * Description of the storage policy.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the storage policy.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the storage policy.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the storage policy.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * List of tag rules. The tag category and tags to be associated to this storage policy.
     * 
     */
    @Import(name="tagRules")
    private @Nullable Output<List<VmStoragePolicyTagRuleArgs>> tagRules;

    /**
     * @return List of tag rules. The tag category and tags to be associated to this storage policy.
     * 
     */
    public Optional<Output<List<VmStoragePolicyTagRuleArgs>>> tagRules() {
        return Optional.ofNullable(this.tagRules);
    }

    private VmStoragePolicyState() {}

    private VmStoragePolicyState(VmStoragePolicyState $) {
        this.description = $.description;
        this.name = $.name;
        this.tagRules = $.tagRules;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VmStoragePolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VmStoragePolicyState $;

        public Builder() {
            $ = new VmStoragePolicyState();
        }

        public Builder(VmStoragePolicyState defaults) {
            $ = new VmStoragePolicyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description of the storage policy.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the storage policy.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The name of the storage policy.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the storage policy.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tagRules List of tag rules. The tag category and tags to be associated to this storage policy.
         * 
         * @return builder
         * 
         */
        public Builder tagRules(@Nullable Output<List<VmStoragePolicyTagRuleArgs>> tagRules) {
            $.tagRules = tagRules;
            return this;
        }

        /**
         * @param tagRules List of tag rules. The tag category and tags to be associated to this storage policy.
         * 
         * @return builder
         * 
         */
        public Builder tagRules(List<VmStoragePolicyTagRuleArgs> tagRules) {
            return tagRules(Output.of(tagRules));
        }

        /**
         * @param tagRules List of tag rules. The tag category and tags to be associated to this storage policy.
         * 
         * @return builder
         * 
         */
        public Builder tagRules(VmStoragePolicyTagRuleArgs... tagRules) {
            return tagRules(List.of(tagRules));
        }

        public VmStoragePolicyState build() {
            return $;
        }
    }

}
