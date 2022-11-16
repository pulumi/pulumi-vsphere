// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetContentLibraryItemArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetContentLibraryItemArgs Empty = new GetContentLibraryItemArgs();

    /**
     * The ID of the content library in which the item exists.
     * 
     */
    @Import(name="libraryId", required=true)
    private Output<String> libraryId;

    /**
     * @return The ID of the content library in which the item exists.
     * 
     */
    public Output<String> libraryId() {
        return this.libraryId;
    }

    /**
     * The name of the content library item.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the content library item.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The type for the content library item. One of `ovf`, `vm-template`, or `iso`
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type for the content library item. One of `ovf`, `vm-template`, or `iso`
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private GetContentLibraryItemArgs() {}

    private GetContentLibraryItemArgs(GetContentLibraryItemArgs $) {
        this.libraryId = $.libraryId;
        this.name = $.name;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetContentLibraryItemArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetContentLibraryItemArgs $;

        public Builder() {
            $ = new GetContentLibraryItemArgs();
        }

        public Builder(GetContentLibraryItemArgs defaults) {
            $ = new GetContentLibraryItemArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param libraryId The ID of the content library in which the item exists.
         * 
         * @return builder
         * 
         */
        public Builder libraryId(Output<String> libraryId) {
            $.libraryId = libraryId;
            return this;
        }

        /**
         * @param libraryId The ID of the content library in which the item exists.
         * 
         * @return builder
         * 
         */
        public Builder libraryId(String libraryId) {
            return libraryId(Output.of(libraryId));
        }

        /**
         * @param name The name of the content library item.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the content library item.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param type The type for the content library item. One of `ovf`, `vm-template`, or `iso`
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type for the content library item. One of `ovf`, `vm-template`, or `iso`
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public GetContentLibraryItemArgs build() {
            $.libraryId = Objects.requireNonNull($.libraryId, "expected parameter 'libraryId' to be non-null");
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}