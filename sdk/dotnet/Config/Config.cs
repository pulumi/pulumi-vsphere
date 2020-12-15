// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.VSphere
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("vsphere");
        /// <summary>
        /// If set, VMware vSphere client will permit unverifiable SSL certificates.
        /// </summary>
        public static bool? AllowUnverifiedSsl { get; set; } = __config.GetBoolean("allowUnverifiedSsl") ?? Utilities.GetEnvBoolean("VSPHERE_ALLOW_UNVERIFIED_SSL");

        /// <summary>
        /// API timeout in minutes (Default: 5)
        /// </summary>
        public static int? ApiTimeout { get; set; } = __config.GetInt32("apiTimeout");

        /// <summary>
        /// govmomi debug
        /// </summary>
        public static bool? ClientDebug { get; set; } = __config.GetBoolean("clientDebug") ?? Utilities.GetEnvBoolean("VSPHERE_CLIENT_DEBUG");

        /// <summary>
        /// govmomi debug path for debug
        /// </summary>
        public static string? ClientDebugPath { get; set; } = __config.Get("clientDebugPath") ?? Utilities.GetEnv("VSPHERE_CLIENT_DEBUG_PATH");

        /// <summary>
        /// govmomi debug path for a single run
        /// </summary>
        public static string? ClientDebugPathRun { get; set; } = __config.Get("clientDebugPathRun") ?? Utilities.GetEnv("VSPHERE_CLIENT_DEBUG_PATH_RUN");

        /// <summary>
        /// The user password for vSphere API operations.
        /// </summary>
        public static string? Password { get; set; } = __config.Get("password") ?? Utilities.GetEnv("VSPHERE_PASSWORD");

        /// <summary>
        /// Persist vSphere client sessions to disk
        /// </summary>
        public static bool? PersistSession { get; set; } = __config.GetBoolean("persistSession") ?? Utilities.GetEnvBoolean("VSPHERE_PERSIST_SESSION");

        /// <summary>
        /// The directory to save vSphere REST API sessions to
        /// </summary>
        public static string? RestSessionPath { get; set; } = __config.Get("restSessionPath") ?? Utilities.GetEnv("VSPHERE_REST_SESSION_PATH");

        /// <summary>
        /// The user name for vSphere API operations.
        /// </summary>
        public static string? User { get; set; } = __config.Get("user") ?? Utilities.GetEnv("VSPHERE_USER");

        public static string? VcenterServer { get; set; } = __config.Get("vcenterServer");

        /// <summary>
        /// Keep alive interval for the VIM session in minutes
        /// </summary>
        public static int? VimKeepAlive { get; set; } = __config.GetInt32("vimKeepAlive") ?? Utilities.GetEnvInt32("VSPHERE_VIM_KEEP_ALIVE");

        /// <summary>
        /// The directory to save vSphere SOAP API sessions to
        /// </summary>
        public static string? VimSessionPath { get; set; } = __config.Get("vimSessionPath") ?? Utilities.GetEnv("VSPHERE_VIM_SESSION_PATH");

        /// <summary>
        /// The vSphere Server name for vSphere API operations.
        /// </summary>
        public static string? VsphereServer { get; set; } = __config.Get("vsphereServer") ?? Utilities.GetEnv("VSPHERE_SERVER");

    }
}
