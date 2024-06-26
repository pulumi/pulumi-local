// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export class File extends pulumi.CustomResource {
    /**
     * Get an existing File resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FileState, opts?: pulumi.CustomResourceOptions): File {
        return new File(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'local:index/file:File';

    /**
     * Returns true if the given object is an instance of File.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is File {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === File.__pulumiType;
    }

    /**
     * Content to store in the file, expected to be a UTF-8 encoded string.
     * Conflicts with `sensitiveContent`, `contentBase64` and `source`.
     * Exactly one of these four arguments must be specified.
     */
    public readonly content!: pulumi.Output<string | undefined>;
    /**
     * Content to store in the file, expected to be binary encoded as base64 string.
     * Conflicts with `content`, `sensitiveContent` and `source`.
     * Exactly one of these four arguments must be specified.
     */
    public readonly contentBase64!: pulumi.Output<string | undefined>;
    /**
     * Base64 encoded SHA256 checksum of file content.
     */
    public /*out*/ readonly contentBase64sha256!: pulumi.Output<string>;
    /**
     * Base64 encoded SHA512 checksum of file content.
     */
    public /*out*/ readonly contentBase64sha512!: pulumi.Output<string>;
    /**
     * MD5 checksum of file content.
     */
    public /*out*/ readonly contentMd5!: pulumi.Output<string>;
    /**
     * SHA1 checksum of file content.
     */
    public /*out*/ readonly contentSha1!: pulumi.Output<string>;
    /**
     * SHA256 checksum of file content.
     */
    public /*out*/ readonly contentSha256!: pulumi.Output<string>;
    /**
     * SHA512 checksum of file content.
     */
    public /*out*/ readonly contentSha512!: pulumi.Output<string>;
    /**
     * Permissions to set for directories created (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `"0777"`.
     */
    public readonly directoryPermission!: pulumi.Output<string>;
    /**
     * Permissions to set for the output file (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `"0777"`.
     */
    public readonly filePermission!: pulumi.Output<string>;
    /**
     * The path to the file that will be created.
     * Missing parent directories will be created.
     * If the file already exists, it will be overridden with the given content.
     */
    public readonly filename!: pulumi.Output<string>;
    /**
     * Sensitive content to store in the file, expected to be an UTF-8 encoded string.
     * Will not be displayed in diffs.
     * Conflicts with `content`, `contentBase64` and `source`.
     * Exactly one of these four arguments must be specified.
     * If in need to use *sensitive* content, please use the `local.SensitiveFile`
     * resource instead.
     *
     * @deprecated Use the `local.SensitiveFile` resource instead
     */
    public readonly sensitiveContent!: pulumi.Output<string | undefined>;
    /**
     * Path to file to use as source for the one we are creating.
     * Conflicts with `content`, `sensitiveContent` and `contentBase64`.
     * Exactly one of these four arguments must be specified.
     */
    public readonly source!: pulumi.Output<string | undefined>;

    /**
     * Create a File resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FileArgs | FileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FileState | undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["contentBase64"] = state ? state.contentBase64 : undefined;
            resourceInputs["contentBase64sha256"] = state ? state.contentBase64sha256 : undefined;
            resourceInputs["contentBase64sha512"] = state ? state.contentBase64sha512 : undefined;
            resourceInputs["contentMd5"] = state ? state.contentMd5 : undefined;
            resourceInputs["contentSha1"] = state ? state.contentSha1 : undefined;
            resourceInputs["contentSha256"] = state ? state.contentSha256 : undefined;
            resourceInputs["contentSha512"] = state ? state.contentSha512 : undefined;
            resourceInputs["directoryPermission"] = state ? state.directoryPermission : undefined;
            resourceInputs["filePermission"] = state ? state.filePermission : undefined;
            resourceInputs["filename"] = state ? state.filename : undefined;
            resourceInputs["sensitiveContent"] = state ? state.sensitiveContent : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
        } else {
            const args = argsOrState as FileArgs | undefined;
            if ((!args || args.filename === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filename'");
            }
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["contentBase64"] = args ? args.contentBase64 : undefined;
            resourceInputs["directoryPermission"] = args ? args.directoryPermission : undefined;
            resourceInputs["filePermission"] = args ? args.filePermission : undefined;
            resourceInputs["filename"] = args ? args.filename : undefined;
            resourceInputs["sensitiveContent"] = args?.sensitiveContent ? pulumi.secret(args.sensitiveContent) : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["contentBase64sha256"] = undefined /*out*/;
            resourceInputs["contentBase64sha512"] = undefined /*out*/;
            resourceInputs["contentMd5"] = undefined /*out*/;
            resourceInputs["contentSha1"] = undefined /*out*/;
            resourceInputs["contentSha256"] = undefined /*out*/;
            resourceInputs["contentSha512"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["sensitiveContent"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(File.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering File resources.
 */
export interface FileState {
    /**
     * Content to store in the file, expected to be a UTF-8 encoded string.
     * Conflicts with `sensitiveContent`, `contentBase64` and `source`.
     * Exactly one of these four arguments must be specified.
     */
    content?: pulumi.Input<string>;
    /**
     * Content to store in the file, expected to be binary encoded as base64 string.
     * Conflicts with `content`, `sensitiveContent` and `source`.
     * Exactly one of these four arguments must be specified.
     */
    contentBase64?: pulumi.Input<string>;
    /**
     * Base64 encoded SHA256 checksum of file content.
     */
    contentBase64sha256?: pulumi.Input<string>;
    /**
     * Base64 encoded SHA512 checksum of file content.
     */
    contentBase64sha512?: pulumi.Input<string>;
    /**
     * MD5 checksum of file content.
     */
    contentMd5?: pulumi.Input<string>;
    /**
     * SHA1 checksum of file content.
     */
    contentSha1?: pulumi.Input<string>;
    /**
     * SHA256 checksum of file content.
     */
    contentSha256?: pulumi.Input<string>;
    /**
     * SHA512 checksum of file content.
     */
    contentSha512?: pulumi.Input<string>;
    /**
     * Permissions to set for directories created (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `"0777"`.
     */
    directoryPermission?: pulumi.Input<string>;
    /**
     * Permissions to set for the output file (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `"0777"`.
     */
    filePermission?: pulumi.Input<string>;
    /**
     * The path to the file that will be created.
     * Missing parent directories will be created.
     * If the file already exists, it will be overridden with the given content.
     */
    filename?: pulumi.Input<string>;
    /**
     * Sensitive content to store in the file, expected to be an UTF-8 encoded string.
     * Will not be displayed in diffs.
     * Conflicts with `content`, `contentBase64` and `source`.
     * Exactly one of these four arguments must be specified.
     * If in need to use *sensitive* content, please use the `local.SensitiveFile`
     * resource instead.
     *
     * @deprecated Use the `local.SensitiveFile` resource instead
     */
    sensitiveContent?: pulumi.Input<string>;
    /**
     * Path to file to use as source for the one we are creating.
     * Conflicts with `content`, `sensitiveContent` and `contentBase64`.
     * Exactly one of these four arguments must be specified.
     */
    source?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a File resource.
 */
export interface FileArgs {
    /**
     * Content to store in the file, expected to be a UTF-8 encoded string.
     * Conflicts with `sensitiveContent`, `contentBase64` and `source`.
     * Exactly one of these four arguments must be specified.
     */
    content?: pulumi.Input<string>;
    /**
     * Content to store in the file, expected to be binary encoded as base64 string.
     * Conflicts with `content`, `sensitiveContent` and `source`.
     * Exactly one of these four arguments must be specified.
     */
    contentBase64?: pulumi.Input<string>;
    /**
     * Permissions to set for directories created (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `"0777"`.
     */
    directoryPermission?: pulumi.Input<string>;
    /**
     * Permissions to set for the output file (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `"0777"`.
     */
    filePermission?: pulumi.Input<string>;
    /**
     * The path to the file that will be created.
     * Missing parent directories will be created.
     * If the file already exists, it will be overridden with the given content.
     */
    filename: pulumi.Input<string>;
    /**
     * Sensitive content to store in the file, expected to be an UTF-8 encoded string.
     * Will not be displayed in diffs.
     * Conflicts with `content`, `contentBase64` and `source`.
     * Exactly one of these four arguments must be specified.
     * If in need to use *sensitive* content, please use the `local.SensitiveFile`
     * resource instead.
     *
     * @deprecated Use the `local.SensitiveFile` resource instead
     */
    sensitiveContent?: pulumi.Input<string>;
    /**
     * Path to file to use as source for the one we are creating.
     * Conflicts with `content`, `sensitiveContent` and `contentBase64`.
     * Exactly one of these four arguments must be specified.
     */
    source?: pulumi.Input<string>;
}
