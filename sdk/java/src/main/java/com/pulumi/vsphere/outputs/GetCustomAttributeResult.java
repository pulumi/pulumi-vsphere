// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCustomAttributeResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String managedObjectType;
    private String name;

    private GetCustomAttributeResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String managedObjectType() {
        return this.managedObjectType;
    }
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCustomAttributeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String managedObjectType;
        private String name;
        public Builder() {}
        public Builder(GetCustomAttributeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.managedObjectType = defaults.managedObjectType;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetCustomAttributeResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder managedObjectType(String managedObjectType) {
            if (managedObjectType == null) {
              throw new MissingRequiredPropertyException("GetCustomAttributeResult", "managedObjectType");
            }
            this.managedObjectType = managedObjectType;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetCustomAttributeResult", "name");
            }
            this.name = name;
            return this;
        }
        public GetCustomAttributeResult build() {
            final var _resultValue = new GetCustomAttributeResult();
            _resultValue.id = id;
            _resultValue.managedObjectType = managedObjectType;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}
