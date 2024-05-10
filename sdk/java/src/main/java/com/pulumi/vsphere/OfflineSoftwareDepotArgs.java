// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class OfflineSoftwareDepotArgs extends com.pulumi.resources.ResourceArgs {

    public static final OfflineSoftwareDepotArgs Empty = new OfflineSoftwareDepotArgs();

    /**
     * The URL where the depot source is hosted.
     * 
     */
    @Import(name="location", required=true)
    private Output<String> location;

    /**
     * @return The URL where the depot source is hosted.
     * 
     */
    public Output<String> location() {
        return this.location;
    }

    private OfflineSoftwareDepotArgs() {}

    private OfflineSoftwareDepotArgs(OfflineSoftwareDepotArgs $) {
        this.location = $.location;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OfflineSoftwareDepotArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OfflineSoftwareDepotArgs $;

        public Builder() {
            $ = new OfflineSoftwareDepotArgs();
        }

        public Builder(OfflineSoftwareDepotArgs defaults) {
            $ = new OfflineSoftwareDepotArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param location The URL where the depot source is hosted.
         * 
         * @return builder
         * 
         */
        public Builder location(Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The URL where the depot source is hosted.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        public OfflineSoftwareDepotArgs build() {
            if ($.location == null) {
                throw new MissingRequiredPropertyException("OfflineSoftwareDepotArgs", "location");
            }
            return $;
        }
    }

}
