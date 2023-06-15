// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Local
{
    [LocalResourceType("local:index/sensitiveFile:SensitiveFile")]
    public partial class SensitiveFile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
        /// Conflicts with `content_base64` and `source`.
        /// Exactly one of these three arguments must be specified.
        /// </summary>
        [Output("content")]
        public Output<string?> Content { get; private set; } = null!;

        /// <summary>
        /// Sensitive Content to store in the file, expected to be binary encoded as base64 string.
        /// Conflicts with `content` and `source`.
        /// Exactly one of these three arguments must be specified.
        /// </summary>
        [Output("contentBase64")]
        public Output<string?> ContentBase64 { get; private set; } = null!;

        /// <summary>
        /// Base64 encoded SHA256 checksum of file content.
        /// </summary>
        [Output("contentBase64sha256")]
        public Output<string> ContentBase64sha256 { get; private set; } = null!;

        /// <summary>
        /// Base64 encoded SHA512 checksum of file content.
        /// </summary>
        [Output("contentBase64sha512")]
        public Output<string> ContentBase64sha512 { get; private set; } = null!;

        /// <summary>
        /// MD5 checksum of file content.
        /// </summary>
        [Output("contentMd5")]
        public Output<string> ContentMd5 { get; private set; } = null!;

        /// <summary>
        /// SHA1 checksum of file content.
        /// </summary>
        [Output("contentSha1")]
        public Output<string> ContentSha1 { get; private set; } = null!;

        /// <summary>
        /// SHA256 checksum of file content.
        /// </summary>
        [Output("contentSha256")]
        public Output<string> ContentSha256 { get; private set; } = null!;

        /// <summary>
        /// SHA512 checksum of file content.
        /// </summary>
        [Output("contentSha512")]
        public Output<string> ContentSha512 { get; private set; } = null!;

        /// <summary>
        /// Permissions to set for directories created (before umask), expressed as string in
        /// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
        /// Default value is `"0700"`.
        /// </summary>
        [Output("directoryPermission")]
        public Output<string> DirectoryPermission { get; private set; } = null!;

        /// <summary>
        /// Permissions to set for the output file (before umask), expressed as string in
        /// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
        /// Default value is `"0700"`.
        /// </summary>
        [Output("filePermission")]
        public Output<string> FilePermission { get; private set; } = null!;

        /// <summary>
        /// The path to the file that will be created.
        /// Missing parent directories will be created.
        /// If the file already exists, it will be overridden with the given content.
        /// </summary>
        [Output("filename")]
        public Output<string> Filename { get; private set; } = null!;

        /// <summary>
        /// Path to file to use as source for the one we are creating.
        /// Conflicts with `content` and `content_base64`.
        /// Exactly one of these three arguments must be specified.
        /// </summary>
        [Output("source")]
        public Output<string?> Source { get; private set; } = null!;


        /// <summary>
        /// Create a SensitiveFile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SensitiveFile(string name, SensitiveFileArgs args, CustomResourceOptions? options = null)
            : base("local:index/sensitiveFile:SensitiveFile", name, args ?? new SensitiveFileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SensitiveFile(string name, Input<string> id, SensitiveFileState? state = null, CustomResourceOptions? options = null)
            : base("local:index/sensitiveFile:SensitiveFile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "content",
                    "contentBase64",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SensitiveFile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SensitiveFile Get(string name, Input<string> id, SensitiveFileState? state = null, CustomResourceOptions? options = null)
        {
            return new SensitiveFile(name, id, state, options);
        }
    }

    public sealed class SensitiveFileArgs : global::Pulumi.ResourceArgs
    {
        [Input("content")]
        private Input<string>? _content;

        /// <summary>
        /// Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
        /// Conflicts with `content_base64` and `source`.
        /// Exactly one of these three arguments must be specified.
        /// </summary>
        public Input<string>? Content
        {
            get => _content;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _content = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("contentBase64")]
        private Input<string>? _contentBase64;

        /// <summary>
        /// Sensitive Content to store in the file, expected to be binary encoded as base64 string.
        /// Conflicts with `content` and `source`.
        /// Exactly one of these three arguments must be specified.
        /// </summary>
        public Input<string>? ContentBase64
        {
            get => _contentBase64;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _contentBase64 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Permissions to set for directories created (before umask), expressed as string in
        /// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
        /// Default value is `"0700"`.
        /// </summary>
        [Input("directoryPermission")]
        public Input<string>? DirectoryPermission { get; set; }

        /// <summary>
        /// Permissions to set for the output file (before umask), expressed as string in
        /// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
        /// Default value is `"0700"`.
        /// </summary>
        [Input("filePermission")]
        public Input<string>? FilePermission { get; set; }

        /// <summary>
        /// The path to the file that will be created.
        /// Missing parent directories will be created.
        /// If the file already exists, it will be overridden with the given content.
        /// </summary>
        [Input("filename", required: true)]
        public Input<string> Filename { get; set; } = null!;

        /// <summary>
        /// Path to file to use as source for the one we are creating.
        /// Conflicts with `content` and `content_base64`.
        /// Exactly one of these three arguments must be specified.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        public SensitiveFileArgs()
        {
        }
        public static new SensitiveFileArgs Empty => new SensitiveFileArgs();
    }

    public sealed class SensitiveFileState : global::Pulumi.ResourceArgs
    {
        [Input("content")]
        private Input<string>? _content;

        /// <summary>
        /// Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
        /// Conflicts with `content_base64` and `source`.
        /// Exactly one of these three arguments must be specified.
        /// </summary>
        public Input<string>? Content
        {
            get => _content;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _content = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("contentBase64")]
        private Input<string>? _contentBase64;

        /// <summary>
        /// Sensitive Content to store in the file, expected to be binary encoded as base64 string.
        /// Conflicts with `content` and `source`.
        /// Exactly one of these three arguments must be specified.
        /// </summary>
        public Input<string>? ContentBase64
        {
            get => _contentBase64;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _contentBase64 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Base64 encoded SHA256 checksum of file content.
        /// </summary>
        [Input("contentBase64sha256")]
        public Input<string>? ContentBase64sha256 { get; set; }

        /// <summary>
        /// Base64 encoded SHA512 checksum of file content.
        /// </summary>
        [Input("contentBase64sha512")]
        public Input<string>? ContentBase64sha512 { get; set; }

        /// <summary>
        /// MD5 checksum of file content.
        /// </summary>
        [Input("contentMd5")]
        public Input<string>? ContentMd5 { get; set; }

        /// <summary>
        /// SHA1 checksum of file content.
        /// </summary>
        [Input("contentSha1")]
        public Input<string>? ContentSha1 { get; set; }

        /// <summary>
        /// SHA256 checksum of file content.
        /// </summary>
        [Input("contentSha256")]
        public Input<string>? ContentSha256 { get; set; }

        /// <summary>
        /// SHA512 checksum of file content.
        /// </summary>
        [Input("contentSha512")]
        public Input<string>? ContentSha512 { get; set; }

        /// <summary>
        /// Permissions to set for directories created (before umask), expressed as string in
        /// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
        /// Default value is `"0700"`.
        /// </summary>
        [Input("directoryPermission")]
        public Input<string>? DirectoryPermission { get; set; }

        /// <summary>
        /// Permissions to set for the output file (before umask), expressed as string in
        /// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
        /// Default value is `"0700"`.
        /// </summary>
        [Input("filePermission")]
        public Input<string>? FilePermission { get; set; }

        /// <summary>
        /// The path to the file that will be created.
        /// Missing parent directories will be created.
        /// If the file already exists, it will be overridden with the given content.
        /// </summary>
        [Input("filename")]
        public Input<string>? Filename { get; set; }

        /// <summary>
        /// Path to file to use as source for the one we are creating.
        /// Conflicts with `content` and `content_base64`.
        /// Exactly one of these three arguments must be specified.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        public SensitiveFileState()
        {
        }
        public static new SensitiveFileState Empty => new SensitiveFileState();
    }
}
