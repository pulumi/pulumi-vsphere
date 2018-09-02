/**
 * If set, VMware vSphere client will permit unverifiable SSL certificates.
 */
export declare let allowUnverifiedSsl: boolean | undefined;
/**
 * govmomi debug
 */
export declare let clientDebug: boolean | undefined;
/**
 * govmomi debug path for debug
 */
export declare let clientDebugPath: string | undefined;
/**
 * govmomi debug path for a single run
 */
export declare let clientDebugPathRun: string | undefined;
/**
 * The user password for vSphere API operations.
 */
export declare let password: string;
/**
 * Persist vSphere client sessions to disk
 */
export declare let persistSession: boolean | undefined;
/**
 * The directory to save vSphere REST API sessions to
 */
export declare let restSessionPath: string | undefined;
/**
 * The user name for vSphere API operations.
 */
export declare let user: string;
export declare let vcenterServer: string | undefined;
/**
 * The directory to save vSphere SOAP API sessions to
 */
export declare let vimSessionPath: string | undefined;
/**
 * The vSphere Server name for vSphere API operations.
 */
export declare let vsphereServer: string | undefined;
