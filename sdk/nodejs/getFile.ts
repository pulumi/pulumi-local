// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Reads a file from the local filesystem.
 */
export function getFile(args: GetFileArgs, opts?: pulumi.InvokeOptions): Promise<GetFileResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("local:index/getFile:getFile", {
        "filename": args.filename,
    }, opts);
}

/**
 * A collection of arguments for invoking getFile.
 */
export interface GetFileArgs {
    /**
     * Path to the file that will be read. The data source will return an error if the file does not exist.
     */
    filename: string;
}

/**
 * A collection of values returned by getFile.
 */
export interface GetFileResult {
    /**
     * Raw content of the file that was read, as UTF-8 encoded string. Files that do not contain UTF-8 text will have invalid UTF-8 sequences in `content`
     * replaced with the Unicode replacement character.
     */
    readonly content: string;
    /**
     * Base64 encoded version of the file content (use this when dealing with binary data).
     */
    readonly contentBase64: string;
    /**
     * Base64 encoded SHA256 checksum of file content.
     */
    readonly contentBase64sha256: string;
    /**
     * Base64 encoded SHA512 checksum of file content.
     */
    readonly contentBase64sha512: string;
    /**
     * MD5 checksum of file content.
     */
    readonly contentMd5: string;
    /**
     * SHA1 checksum of file content.
     */
    readonly contentSha1: string;
    /**
     * SHA256 checksum of file content.
     */
    readonly contentSha256: string;
    /**
     * SHA512 checksum of file content.
     */
    readonly contentSha512: string;
    /**
     * Path to the file that will be read. The data source will return an error if the file does not exist.
     */
    readonly filename: string;
    /**
     * The hexadecimal encoding of the SHA1 checksum of the file content.
     */
    readonly id: string;
}
/**
 * Reads a file from the local filesystem.
 */
export function getFileOutput(args: GetFileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFileResult> {
    return pulumi.output(args).apply((a: any) => getFile(a, opts))
}

/**
 * A collection of arguments for invoking getFile.
 */
export interface GetFileOutputArgs {
    /**
     * Path to the file that will be read. The data source will return an error if the file does not exist.
     */
    filename: pulumi.Input<string>;
}
