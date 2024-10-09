// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetHostThumbprintPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetHostThumbprintPlainArgs Empty = new GetHostThumbprintPlainArgs();

    /**
     * The address of the ESXi host to retrieve the thumbprint
     * from.
     * 
     */
    @Import(name="address", required=true)
    private String address;

    /**
     * @return The address of the ESXi host to retrieve the thumbprint
     * from.
     * 
     */
    public String address() {
        return this.address;
    }

    /**
     * Disables SSL certificate verification. Default: `false`
     * 
     */
    @Import(name="insecure")
    private @Nullable Boolean insecure;

    /**
     * @return Disables SSL certificate verification. Default: `false`
     * 
     */
    public Optional<Boolean> insecure() {
        return Optional.ofNullable(this.insecure);
    }

    /**
     * The port to use connecting to the ESXi host. Default: 443
     * 
     */
    @Import(name="port")
    private @Nullable String port;

    /**
     * @return The port to use connecting to the ESXi host. Default: 443
     * 
     */
    public Optional<String> port() {
        return Optional.ofNullable(this.port);
    }

    private GetHostThumbprintPlainArgs() {}

    private GetHostThumbprintPlainArgs(GetHostThumbprintPlainArgs $) {
        this.address = $.address;
        this.insecure = $.insecure;
        this.port = $.port;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetHostThumbprintPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetHostThumbprintPlainArgs $;

        public Builder() {
            $ = new GetHostThumbprintPlainArgs();
        }

        public Builder(GetHostThumbprintPlainArgs defaults) {
            $ = new GetHostThumbprintPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address The address of the ESXi host to retrieve the thumbprint
         * from.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            $.address = address;
            return this;
        }

        /**
         * @param insecure Disables SSL certificate verification. Default: `false`
         * 
         * @return builder
         * 
         */
        public Builder insecure(@Nullable Boolean insecure) {
            $.insecure = insecure;
            return this;
        }

        /**
         * @param port The port to use connecting to the ESXi host. Default: 443
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable String port) {
            $.port = port;
            return this;
        }

        public GetHostThumbprintPlainArgs build() {
            if ($.address == null) {
                throw new MissingRequiredPropertyException("GetHostThumbprintPlainArgs", "address");
            }
            return $;
        }
    }

}
