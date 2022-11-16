// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ComputeClusterVsanDiskGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final ComputeClusterVsanDiskGroupArgs Empty = new ComputeClusterVsanDiskGroupArgs();

    @Import(name="cache")
    private @Nullable Output<String> cache;

    public Optional<Output<String>> cache() {
        return Optional.ofNullable(this.cache);
    }

    @Import(name="storages")
    private @Nullable Output<List<String>> storages;

    public Optional<Output<List<String>>> storages() {
        return Optional.ofNullable(this.storages);
    }

    private ComputeClusterVsanDiskGroupArgs() {}

    private ComputeClusterVsanDiskGroupArgs(ComputeClusterVsanDiskGroupArgs $) {
        this.cache = $.cache;
        this.storages = $.storages;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ComputeClusterVsanDiskGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ComputeClusterVsanDiskGroupArgs $;

        public Builder() {
            $ = new ComputeClusterVsanDiskGroupArgs();
        }

        public Builder(ComputeClusterVsanDiskGroupArgs defaults) {
            $ = new ComputeClusterVsanDiskGroupArgs(Objects.requireNonNull(defaults));
        }

        public Builder cache(@Nullable Output<String> cache) {
            $.cache = cache;
            return this;
        }

        public Builder cache(String cache) {
            return cache(Output.of(cache));
        }

        public Builder storages(@Nullable Output<List<String>> storages) {
            $.storages = storages;
            return this;
        }

        public Builder storages(List<String> storages) {
            return storages(Output.of(storages));
        }

        public Builder storages(String... storages) {
            return storages(List.of(storages));
        }

        public ComputeClusterVsanDiskGroupArgs build() {
            return $;
        }
    }

}
