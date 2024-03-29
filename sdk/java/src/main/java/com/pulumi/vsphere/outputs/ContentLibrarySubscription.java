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
public final class ContentLibrarySubscription {
    /**
     * @return Authentication method to connect ro a published content library. Must be `NONE` or `BASIC`.
     * 
     */
    private @Nullable String authenticationMethod;
    /**
     * @return Enable automatic synchronization with the published library. Default `false`.
     * 
     */
    private @Nullable Boolean automaticSync;
    /**
     * @return Download the library from a content only when needed. Default `true`.
     * 
     */
    private @Nullable Boolean onDemand;
    /**
     * @return Password used for authentication.
     * 
     */
    private @Nullable String password;
    /**
     * @return URL of the published content library.
     * 
     */
    private @Nullable String subscriptionUrl;
    /**
     * @return Username used for authentication.
     * 
     */
    private @Nullable String username;

    private ContentLibrarySubscription() {}
    /**
     * @return Authentication method to connect ro a published content library. Must be `NONE` or `BASIC`.
     * 
     */
    public Optional<String> authenticationMethod() {
        return Optional.ofNullable(this.authenticationMethod);
    }
    /**
     * @return Enable automatic synchronization with the published library. Default `false`.
     * 
     */
    public Optional<Boolean> automaticSync() {
        return Optional.ofNullable(this.automaticSync);
    }
    /**
     * @return Download the library from a content only when needed. Default `true`.
     * 
     */
    public Optional<Boolean> onDemand() {
        return Optional.ofNullable(this.onDemand);
    }
    /**
     * @return Password used for authentication.
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return URL of the published content library.
     * 
     */
    public Optional<String> subscriptionUrl() {
        return Optional.ofNullable(this.subscriptionUrl);
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

    public static Builder builder(ContentLibrarySubscription defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String authenticationMethod;
        private @Nullable Boolean automaticSync;
        private @Nullable Boolean onDemand;
        private @Nullable String password;
        private @Nullable String subscriptionUrl;
        private @Nullable String username;
        public Builder() {}
        public Builder(ContentLibrarySubscription defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authenticationMethod = defaults.authenticationMethod;
    	      this.automaticSync = defaults.automaticSync;
    	      this.onDemand = defaults.onDemand;
    	      this.password = defaults.password;
    	      this.subscriptionUrl = defaults.subscriptionUrl;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder authenticationMethod(@Nullable String authenticationMethod) {

            this.authenticationMethod = authenticationMethod;
            return this;
        }
        @CustomType.Setter
        public Builder automaticSync(@Nullable Boolean automaticSync) {

            this.automaticSync = automaticSync;
            return this;
        }
        @CustomType.Setter
        public Builder onDemand(@Nullable Boolean onDemand) {

            this.onDemand = onDemand;
            return this;
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {

            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder subscriptionUrl(@Nullable String subscriptionUrl) {

            this.subscriptionUrl = subscriptionUrl;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {

            this.username = username;
            return this;
        }
        public ContentLibrarySubscription build() {
            final var _resultValue = new ContentLibrarySubscription();
            _resultValue.authenticationMethod = authenticationMethod;
            _resultValue.automaticSync = automaticSync;
            _resultValue.onDemand = onDemand;
            _resultValue.password = password;
            _resultValue.subscriptionUrl = subscriptionUrl;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
