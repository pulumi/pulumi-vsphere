// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LicenseState extends com.pulumi.resources.ResourceArgs {

    public static final LicenseState Empty = new LicenseState();

    /**
     * The product edition of the license key.
     * 
     */
    @Import(name="editionKey")
    private @Nullable Output<String> editionKey;

    /**
     * @return The product edition of the license key.
     * 
     */
    public Optional<Output<String>> editionKey() {
        return Optional.ofNullable(this.editionKey);
    }

    /**
     * A map of key/value pairs to be attached as labels (tags) to the license key.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return A map of key/value pairs to be attached as labels (tags) to the license key.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The license key to add.
     * 
     */
    @Import(name="licenseKey")
    private @Nullable Output<String> licenseKey;

    /**
     * @return The license key to add.
     * 
     */
    public Optional<Output<String>> licenseKey() {
        return Optional.ofNullable(this.licenseKey);
    }

    /**
     * The display name for the license.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The display name for the license.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Total number of units (example: CPUs) contained in the license.
     * 
     */
    @Import(name="total")
    private @Nullable Output<Integer> total;

    /**
     * @return Total number of units (example: CPUs) contained in the license.
     * 
     */
    public Optional<Output<Integer>> total() {
        return Optional.ofNullable(this.total);
    }

    /**
     * The number of units (example: CPUs) assigned to this license.
     * 
     */
    @Import(name="used")
    private @Nullable Output<Integer> used;

    /**
     * @return The number of units (example: CPUs) assigned to this license.
     * 
     */
    public Optional<Output<Integer>> used() {
        return Optional.ofNullable(this.used);
    }

    private LicenseState() {}

    private LicenseState(LicenseState $) {
        this.editionKey = $.editionKey;
        this.labels = $.labels;
        this.licenseKey = $.licenseKey;
        this.name = $.name;
        this.total = $.total;
        this.used = $.used;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LicenseState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LicenseState $;

        public Builder() {
            $ = new LicenseState();
        }

        public Builder(LicenseState defaults) {
            $ = new LicenseState(Objects.requireNonNull(defaults));
        }

        /**
         * @param editionKey The product edition of the license key.
         * 
         * @return builder
         * 
         */
        public Builder editionKey(@Nullable Output<String> editionKey) {
            $.editionKey = editionKey;
            return this;
        }

        /**
         * @param editionKey The product edition of the license key.
         * 
         * @return builder
         * 
         */
        public Builder editionKey(String editionKey) {
            return editionKey(Output.of(editionKey));
        }

        /**
         * @param labels A map of key/value pairs to be attached as labels (tags) to the license key.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels A map of key/value pairs to be attached as labels (tags) to the license key.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param licenseKey The license key to add.
         * 
         * @return builder
         * 
         */
        public Builder licenseKey(@Nullable Output<String> licenseKey) {
            $.licenseKey = licenseKey;
            return this;
        }

        /**
         * @param licenseKey The license key to add.
         * 
         * @return builder
         * 
         */
        public Builder licenseKey(String licenseKey) {
            return licenseKey(Output.of(licenseKey));
        }

        /**
         * @param name The display name for the license.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The display name for the license.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param total Total number of units (example: CPUs) contained in the license.
         * 
         * @return builder
         * 
         */
        public Builder total(@Nullable Output<Integer> total) {
            $.total = total;
            return this;
        }

        /**
         * @param total Total number of units (example: CPUs) contained in the license.
         * 
         * @return builder
         * 
         */
        public Builder total(Integer total) {
            return total(Output.of(total));
        }

        /**
         * @param used The number of units (example: CPUs) assigned to this license.
         * 
         * @return builder
         * 
         */
        public Builder used(@Nullable Output<Integer> used) {
            $.used = used;
            return this;
        }

        /**
         * @param used The number of units (example: CPUs) assigned to this license.
         * 
         * @return builder
         * 
         */
        public Builder used(Integer used) {
            return used(Output.of(used));
        }

        public LicenseState build() {
            return $;
        }
    }

}