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


public final class GetDynamicPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDynamicPlainArgs Empty = new GetDynamicPlainArgs();

    /**
     * A list of tag IDs that must be present on an object to
     * be a match.
     * 
     */
    @Import(name="filters", required=true)
    private List<String> filters;

    /**
     * @return A list of tag IDs that must be present on an object to
     * be a match.
     * 
     */
    public List<String> filters() {
        return this.filters;
    }

    /**
     * A regular expression that will be used to match the
     * object&#39;s name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regular expression that will be used to match the
     * object&#39;s name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * The managed object type the returned object must match.
     * The managed object types can be found in the managed object type section
     * [here](https://developer.broadcom.com/xapis/vsphere-web-services-api/latest/).
     * 
     */
    @Import(name="type")
    private @Nullable String type;

    /**
     * @return The managed object type the returned object must match.
     * The managed object types can be found in the managed object type section
     * [here](https://developer.broadcom.com/xapis/vsphere-web-services-api/latest/).
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    private GetDynamicPlainArgs() {}

    private GetDynamicPlainArgs(GetDynamicPlainArgs $) {
        this.filters = $.filters;
        this.nameRegex = $.nameRegex;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDynamicPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDynamicPlainArgs $;

        public Builder() {
            $ = new GetDynamicPlainArgs();
        }

        public Builder(GetDynamicPlainArgs defaults) {
            $ = new GetDynamicPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters A list of tag IDs that must be present on an object to
         * be a match.
         * 
         * @return builder
         * 
         */
        public Builder filters(List<String> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters A list of tag IDs that must be present on an object to
         * be a match.
         * 
         * @return builder
         * 
         */
        public Builder filters(String... filters) {
            return filters(List.of(filters));
        }

        /**
         * @param nameRegex A regular expression that will be used to match the
         * object&#39;s name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param type The managed object type the returned object must match.
         * The managed object types can be found in the managed object type section
         * [here](https://developer.broadcom.com/xapis/vsphere-web-services-api/latest/).
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        public GetDynamicPlainArgs build() {
            if ($.filters == null) {
                throw new MissingRequiredPropertyException("GetDynamicPlainArgs", "filters");
            }
            return $;
        }
    }

}
