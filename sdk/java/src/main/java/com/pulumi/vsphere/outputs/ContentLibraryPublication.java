// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContentLibraryPublication {
    /**
     * @return Authentication method to connect ro a published content library. Must be `NONE` or `BASIC`.
     * 
     */
    private @Nullable String authenticationMethod;
    /**
     * @return Password used for authentication.
     * 
     */
    private @Nullable String password;
    /**
     * @return The URL of the published content library.
     * 
     */
    private @Nullable String publishUrl;
    /**
     * @return Publish the content library. Default `false`.
     * 
     */
    private @Nullable Boolean published;
    /**
     * @return Username used for authentication.
     * 
     */
    private @Nullable String username;

    private ContentLibraryPublication() {}
    /**
     * @return Authentication method to connect ro a published content library. Must be `NONE` or `BASIC`.
     * 
     */
    public Optional<String> authenticationMethod() {
        return Optional.ofNullable(this.authenticationMethod);
    }
    /**
     * @return Password used for authentication.
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return The URL of the published content library.
     * 
     */
    public Optional<String> publishUrl() {
        return Optional.ofNullable(this.publishUrl);
    }
    /**
     * @return Publish the content library. Default `false`.
     * 
     */
    public Optional<Boolean> published() {
        return Optional.ofNullable(this.published);
    }
    /**
     * @return Username used for authentication.
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContentLibraryPublication defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String authenticationMethod;
        private @Nullable String password;
        private @Nullable String publishUrl;
        private @Nullable Boolean published;
        private @Nullable String username;
        public Builder() {}
        public Builder(ContentLibraryPublication defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authenticationMethod = defaults.authenticationMethod;
    	      this.password = defaults.password;
    	      this.publishUrl = defaults.publishUrl;
    	      this.published = defaults.published;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder authenticationMethod(@Nullable String authenticationMethod) {

            this.authenticationMethod = authenticationMethod;
            return this;
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {

            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder publishUrl(@Nullable String publishUrl) {

            this.publishUrl = publishUrl;
            return this;
        }
        @CustomType.Setter
        public Builder published(@Nullable Boolean published) {

            this.published = published;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {

            this.username = username;
            return this;
        }
        public ContentLibraryPublication build() {
            final var _resultValue = new ContentLibraryPublication();
            _resultValue.authenticationMethod = authenticationMethod;
            _resultValue.password = password;
            _resultValue.publishUrl = publishUrl;
            _resultValue.published = published;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
