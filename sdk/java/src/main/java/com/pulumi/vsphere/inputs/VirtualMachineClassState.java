// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualMachineClassState extends com.pulumi.resources.ResourceArgs {

    public static final VirtualMachineClassState Empty = new VirtualMachineClassState();

    /**
     * The percentage of the available CPU capacity which will be reserved.
     * 
     */
    @Import(name="cpuReservation")
    private @Nullable Output<Integer> cpuReservation;

    /**
     * @return The percentage of the available CPU capacity which will be reserved.
     * 
     */
    public Optional<Output<Integer>> cpuReservation() {
        return Optional.ofNullable(this.cpuReservation);
    }

    /**
     * The number of CPUs.
     * 
     */
    @Import(name="cpus")
    private @Nullable Output<Integer> cpus;

    /**
     * @return The number of CPUs.
     * 
     */
    public Optional<Output<Integer>> cpus() {
        return Optional.ofNullable(this.cpus);
    }

    /**
     * The amount of memory in MB.
     * 
     */
    @Import(name="memory")
    private @Nullable Output<Integer> memory;

    /**
     * @return The amount of memory in MB.
     * 
     */
    public Optional<Output<Integer>> memory() {
        return Optional.ofNullable(this.memory);
    }

    /**
     * The percentage of memory reservation.
     * 
     */
    @Import(name="memoryReservation")
    private @Nullable Output<Integer> memoryReservation;

    /**
     * @return The percentage of memory reservation.
     * 
     */
    public Optional<Output<Integer>> memoryReservation() {
        return Optional.ofNullable(this.memoryReservation);
    }

    /**
     * The name for the class.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name for the class.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The identifiers of the vGPU devices for the class. If this is set memory reservation needs to be 100.
     * 
     */
    @Import(name="vgpuDevices")
    private @Nullable Output<List<String>> vgpuDevices;

    /**
     * @return The identifiers of the vGPU devices for the class. If this is set memory reservation needs to be 100.
     * 
     */
    public Optional<Output<List<String>>> vgpuDevices() {
        return Optional.ofNullable(this.vgpuDevices);
    }

    private VirtualMachineClassState() {}

    private VirtualMachineClassState(VirtualMachineClassState $) {
        this.cpuReservation = $.cpuReservation;
        this.cpus = $.cpus;
        this.memory = $.memory;
        this.memoryReservation = $.memoryReservation;
        this.name = $.name;
        this.vgpuDevices = $.vgpuDevices;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualMachineClassState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualMachineClassState $;

        public Builder() {
            $ = new VirtualMachineClassState();
        }

        public Builder(VirtualMachineClassState defaults) {
            $ = new VirtualMachineClassState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cpuReservation The percentage of the available CPU capacity which will be reserved.
         * 
         * @return builder
         * 
         */
        public Builder cpuReservation(@Nullable Output<Integer> cpuReservation) {
            $.cpuReservation = cpuReservation;
            return this;
        }

        /**
         * @param cpuReservation The percentage of the available CPU capacity which will be reserved.
         * 
         * @return builder
         * 
         */
        public Builder cpuReservation(Integer cpuReservation) {
            return cpuReservation(Output.of(cpuReservation));
        }

        /**
         * @param cpus The number of CPUs.
         * 
         * @return builder
         * 
         */
        public Builder cpus(@Nullable Output<Integer> cpus) {
            $.cpus = cpus;
            return this;
        }

        /**
         * @param cpus The number of CPUs.
         * 
         * @return builder
         * 
         */
        public Builder cpus(Integer cpus) {
            return cpus(Output.of(cpus));
        }

        /**
         * @param memory The amount of memory in MB.
         * 
         * @return builder
         * 
         */
        public Builder memory(@Nullable Output<Integer> memory) {
            $.memory = memory;
            return this;
        }

        /**
         * @param memory The amount of memory in MB.
         * 
         * @return builder
         * 
         */
        public Builder memory(Integer memory) {
            return memory(Output.of(memory));
        }

        /**
         * @param memoryReservation The percentage of memory reservation.
         * 
         * @return builder
         * 
         */
        public Builder memoryReservation(@Nullable Output<Integer> memoryReservation) {
            $.memoryReservation = memoryReservation;
            return this;
        }

        /**
         * @param memoryReservation The percentage of memory reservation.
         * 
         * @return builder
         * 
         */
        public Builder memoryReservation(Integer memoryReservation) {
            return memoryReservation(Output.of(memoryReservation));
        }

        /**
         * @param name The name for the class.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name for the class.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param vgpuDevices The identifiers of the vGPU devices for the class. If this is set memory reservation needs to be 100.
         * 
         * @return builder
         * 
         */
        public Builder vgpuDevices(@Nullable Output<List<String>> vgpuDevices) {
            $.vgpuDevices = vgpuDevices;
            return this;
        }

        /**
         * @param vgpuDevices The identifiers of the vGPU devices for the class. If this is set memory reservation needs to be 100.
         * 
         * @return builder
         * 
         */
        public Builder vgpuDevices(List<String> vgpuDevices) {
            return vgpuDevices(Output.of(vgpuDevices));
        }

        /**
         * @param vgpuDevices The identifiers of the vGPU devices for the class. If this is set memory reservation needs to be 100.
         * 
         * @return builder
         * 
         */
        public Builder vgpuDevices(String... vgpuDevices) {
            return vgpuDevices(List.of(vgpuDevices));
        }

        public VirtualMachineClassState build() {
            return $;
        }
    }

}
