// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("vsphere");
/**
 * If set, VMware vSphere client will permit unverifiable SSL certificates.
 * 
 */
    public Optional<Boolean> allowUnverifiedSsl() {
        return Codegen.booleanProp("allowUnverifiedSsl").config(config).env("VSPHERE_ALLOW_UNVERIFIED_SSL").get();
    }
/**
 * API timeout in minutes (Default: 5)
 * 
 */
    public Optional<Integer> apiTimeout() {
        return Codegen.integerProp("apiTimeout").config(config).get();
    }
/**
 * govmomi debug
 * 
 */
    public Optional<Boolean> clientDebug() {
        return Codegen.booleanProp("clientDebug").config(config).env("VSPHERE_CLIENT_DEBUG").get();
    }
/**
 * govmomi debug path for debug
 * 
 */
    public Optional<String> clientDebugPath() {
        return Codegen.stringProp("clientDebugPath").config(config).env("VSPHERE_CLIENT_DEBUG_PATH").get();
    }
/**
 * govmomi debug path for a single run
 * 
 */
    public Optional<String> clientDebugPathRun() {
        return Codegen.stringProp("clientDebugPathRun").config(config).env("VSPHERE_CLIENT_DEBUG_PATH_RUN").get();
    }
/**
 * The user password for vSphere API operations.
 * 
 */
    public String password() {
        return Codegen.stringProp("password").config(config).require();
    }
/**
 * Persist vSphere client sessions to disk
 * 
 */
    public Optional<Boolean> persistSession() {
        return Codegen.booleanProp("persistSession").config(config).env("VSPHERE_PERSIST_SESSION").get();
    }
/**
 * The directory to save vSphere REST API sessions to
 * 
 */
    public Optional<String> restSessionPath() {
        return Codegen.stringProp("restSessionPath").config(config).env("VSPHERE_REST_SESSION_PATH").get();
    }
/**
 * The user name for vSphere API operations.
 * 
 */
    public String user() {
        return Codegen.stringProp("user").config(config).require();
    }
    public Optional<String> vcenterServer() {
        return Codegen.stringProp("vcenterServer").config(config).get();
    }
/**
 * Keep alive interval for the VIM session in minutes
 * 
 */
    public Optional<Integer> vimKeepAlive() {
        return Codegen.integerProp("vimKeepAlive").config(config).env("VSPHERE_VIM_KEEP_ALIVE").get();
    }
/**
 * The directory to save vSphere SOAP API sessions to
 * 
 */
    public Optional<String> vimSessionPath() {
        return Codegen.stringProp("vimSessionPath").config(config).env("VSPHERE_VIM_SESSION_PATH").get();
    }
/**
 * The vSphere Server name for vSphere API operations.
 * 
 */
    public Optional<String> vsphereServer() {
        return Codegen.stringProp("vsphereServer").config(config).get();
    }
}
