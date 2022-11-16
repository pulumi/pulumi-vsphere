// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVmfsDisksPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVmfsDisksPlainArgs Empty = new GetVmfsDisksPlainArgs();

    /**
     * A regular expression to filter the disks against. Only
     * disks with canonical names that match will be included.
     * 
     */
    @Import(name="filter")
    private @Nullable String filter;

    /**
     * @return A regular expression to filter the disks against. Only
     * disks with canonical names that match will be included.
     * 
     */
    public Optional<String> filter() {
        return Optional.ofNullable(this.filter);
    }

    /**
     * The managed object ID of
     * the host to look for disks on.
     * 
     */
    @Import(name="hostSystemId", required=true)
    private String hostSystemId;

    /**
     * @return The managed object ID of
     * the host to look for disks on.
     * 
     */
    public String hostSystemId() {
        return this.hostSystemId;
    }

    /**
     * Whether or not to rescan storage adapters before
     * searching for disks. This may lengthen the time it takes to perform the
     * search. Default: `false`.
     * 
     */
    @Import(name="rescan")
    private @Nullable Boolean rescan;

    /**
     * @return Whether or not to rescan storage adapters before
     * searching for disks. This may lengthen the time it takes to perform the
     * search. Default: `false`.
     * 
     */
    public Optional<Boolean> rescan() {
        return Optional.ofNullable(this.rescan);
    }

    private GetVmfsDisksPlainArgs() {}

    private GetVmfsDisksPlainArgs(GetVmfsDisksPlainArgs $) {
        this.filter = $.filter;
        this.hostSystemId = $.hostSystemId;
        this.rescan = $.rescan;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVmfsDisksPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVmfsDisksPlainArgs $;

        public Builder() {
            $ = new GetVmfsDisksPlainArgs();
        }

        public Builder(GetVmfsDisksPlainArgs defaults) {
            $ = new GetVmfsDisksPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filter A regular expression to filter the disks against. Only
         * disks with canonical names that match will be included.
         * 
         * @return builder
         * 
         */
        public Builder filter(@Nullable String filter) {
            $.filter = filter;
            return this;
        }

        /**
         * @param hostSystemId The managed object ID of
         * the host to look for disks on.
         * 
         * @return builder
         * 
         */
        public Builder hostSystemId(String hostSystemId) {
            $.hostSystemId = hostSystemId;
            return this;
        }

        /**
         * @param rescan Whether or not to rescan storage adapters before
         * searching for disks. This may lengthen the time it takes to perform the
         * search. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder rescan(@Nullable Boolean rescan) {
            $.rescan = rescan;
            return this;
        }

        public GetVmfsDisksPlainArgs build() {
            $.hostSystemId = Objects.requireNonNull($.hostSystemId, "expected parameter 'hostSystemId' to be non-null");
            return $;
        }
    }

}
