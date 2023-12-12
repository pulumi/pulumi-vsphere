// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetTagCategoryResult {
    private List<String> associableTypes;
    private String cardinality;
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;

    private GetTagCategoryResult() {}
    public List<String> associableTypes() {
        return this.associableTypes;
    }
    public String cardinality() {
        return this.cardinality;
    }
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTagCategoryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> associableTypes;
        private String cardinality;
        private String description;
        private String id;
        private String name;
        public Builder() {}
        public Builder(GetTagCategoryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.associableTypes = defaults.associableTypes;
    	      this.cardinality = defaults.cardinality;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder associableTypes(List<String> associableTypes) {
            this.associableTypes = Objects.requireNonNull(associableTypes);
            return this;
        }
        public Builder associableTypes(String... associableTypes) {
            return associableTypes(List.of(associableTypes));
        }
        @CustomType.Setter
        public Builder cardinality(String cardinality) {
            this.cardinality = Objects.requireNonNull(cardinality);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public GetTagCategoryResult build() {
            final var _resultValue = new GetTagCategoryResult();
            _resultValue.associableTypes = associableTypes;
            _resultValue.cardinality = cardinality;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}
