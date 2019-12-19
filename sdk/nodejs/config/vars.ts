// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("vsphere");

/**
 * If set, VMware vSphere client will permit unverifiable SSL certificates.
 */
export let allowUnverifiedSsl: boolean | undefined = __config.getObject<boolean>("allowUnverifiedSsl") || utilities.getEnvBoolean("VSPHERE_ALLOW_UNVERIFIED_SSL");
/**
 * govmomi debug
 */
export let clientDebug: boolean | undefined = __config.getObject<boolean>("clientDebug") || utilities.getEnvBoolean("VSPHERE_CLIENT_DEBUG");
/**
 * govmomi debug path for debug
 */
export let clientDebugPath: string | undefined = __config.get("clientDebugPath") || utilities.getEnv("VSPHERE_CLIENT_DEBUG_PATH");
/**
 * govmomi debug path for a single run
 */
export let clientDebugPathRun: string | undefined = __config.get("clientDebugPathRun") || utilities.getEnv("VSPHERE_CLIENT_DEBUG_PATH_RUN");
/**
 * The user password for vSphere API operations.
 */
export let password: string | undefined = __config.get("password") || utilities.getEnv("VSPHERE_PASSWORD");
/**
 * Persist vSphere client sessions to disk
 */
export let persistSession: boolean | undefined = __config.getObject<boolean>("persistSession") || utilities.getEnvBoolean("VSPHERE_PERSIST_SESSION");
/**
 * The directory to save vSphere REST API sessions to
 */
export let restSessionPath: string | undefined = __config.get("restSessionPath") || utilities.getEnv("VSPHERE_REST_SESSION_PATH");
/**
 * The user name for vSphere API operations.
 */
export let user: string | undefined = __config.get("user") || utilities.getEnv("VSPHERE_USER");
export let vcenterServer: string | undefined = __config.get("vcenterServer");
/**
 * Keep alive interval for the VIM session in minutes
 */
export let vimKeepAlive: number | undefined = __config.getObject<number>("vimKeepAlive") || utilities.getEnvNumber("VSPHERE_VIM_KEEP_ALIVE");
/**
 * The directory to save vSphere SOAP API sessions to
 */
export let vimSessionPath: string | undefined = __config.get("vimSessionPath") || utilities.getEnv("VSPHERE_VIM_SESSION_PATH");
/**
 * The vSphere Server name for vSphere API operations.
 */
export let vsphereServer: string | undefined = __config.get("vsphereServer") || utilities.getEnv("VSPHERE_SERVER");
