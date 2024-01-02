// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetFolderArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFolderArgs Empty = new GetFolderArgs();

    /**
     * The absolute path of the folder. For example, given a
     * default datacenter of `default-dc`, a folder of type `vm`, and a folder name
     * of `test-folder`, the resulting path would be
     * `/default-dc/vm/test-folder`. The valid folder types to be used in
     * the path are: `vm`, `host`, `datacenter`, `datastore`, or `network`.
     * 
     */
    @Import(name="path", required=true)
    private Output<String> path;

    /**
     * @return The absolute path of the folder. For example, given a
     * default datacenter of `default-dc`, a folder of type `vm`, and a folder name
     * of `test-folder`, the resulting path would be
     * `/default-dc/vm/test-folder`. The valid folder types to be used in
     * the path are: `vm`, `host`, `datacenter`, `datastore`, or `network`.
     * 
     */
    public Output<String> path() {
        return this.path;
    }

    private GetFolderArgs() {}

    private GetFolderArgs(GetFolderArgs $) {
        this.path = $.path;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFolderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFolderArgs $;

        public Builder() {
            $ = new GetFolderArgs();
        }

        public Builder(GetFolderArgs defaults) {
            $ = new GetFolderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param path The absolute path of the folder. For example, given a
         * default datacenter of `default-dc`, a folder of type `vm`, and a folder name
         * of `test-folder`, the resulting path would be
         * `/default-dc/vm/test-folder`. The valid folder types to be used in
         * the path are: `vm`, `host`, `datacenter`, `datastore`, or `network`.
         * 
         * @return builder
         * 
         */
        public Builder path(Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The absolute path of the folder. For example, given a
         * default datacenter of `default-dc`, a folder of type `vm`, and a folder name
         * of `test-folder`, the resulting path would be
         * `/default-dc/vm/test-folder`. The valid folder types to be used in
         * the path are: `vm`, `host`, `datacenter`, `datastore`, or `network`.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        public GetFolderArgs build() {
            if ($.path == null) {
                throw new MissingRequiredPropertyException("GetFolderArgs", "path");
            }
            return $;
        }
    }

}
