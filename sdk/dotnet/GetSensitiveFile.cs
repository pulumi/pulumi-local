// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Local
{
    public static class GetSensitiveFile
    {
        /// <summary>
        /// Reads a file that contains sensitive data, from the local filesystem.
        /// 
        /// The attributes exposed by this data source are marked as
        /// sensitive.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// using Local = Pulumi.Local;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Local.GetSensitiveFile.Invoke(new()
        ///     {
        ///         Filename = $"{path.Module}/foo.bar",
        ///     });
        /// 
        ///     var sharedZip = new Aws.S3.BucketObjectv2("sharedZip", new()
        ///     {
        ///         Bucket = "my-bucket",
        ///         Key = "my-key",
        ///         Content = foo.Apply(getSensitiveFileResult =&gt; getSensitiveFileResult.Content),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSensitiveFileResult> InvokeAsync(GetSensitiveFileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSensitiveFileResult>("local:index/getSensitiveFile:getSensitiveFile", args ?? new GetSensitiveFileArgs(), options.WithDefaults());

        /// <summary>
        /// Reads a file that contains sensitive data, from the local filesystem.
        /// 
        /// The attributes exposed by this data source are marked as
        /// sensitive.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// using Local = Pulumi.Local;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Local.GetSensitiveFile.Invoke(new()
        ///     {
        ///         Filename = $"{path.Module}/foo.bar",
        ///     });
        /// 
        ///     var sharedZip = new Aws.S3.BucketObjectv2("sharedZip", new()
        ///     {
        ///         Bucket = "my-bucket",
        ///         Key = "my-key",
        ///         Content = foo.Apply(getSensitiveFileResult =&gt; getSensitiveFileResult.Content),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSensitiveFileResult> Invoke(GetSensitiveFileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSensitiveFileResult>("local:index/getSensitiveFile:getSensitiveFile", args ?? new GetSensitiveFileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSensitiveFileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Path to the file that will be read. The data source will return an error if the file does not exist.
        /// </summary>
        [Input("filename", required: true)]
        public string Filename { get; set; } = null!;

        public GetSensitiveFileArgs()
        {
        }
        public static new GetSensitiveFileArgs Empty => new GetSensitiveFileArgs();
    }

    public sealed class GetSensitiveFileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Path to the file that will be read. The data source will return an error if the file does not exist.
        /// </summary>
        [Input("filename", required: true)]
        public Input<string> Filename { get; set; } = null!;

        public GetSensitiveFileInvokeArgs()
        {
        }
        public static new GetSensitiveFileInvokeArgs Empty => new GetSensitiveFileInvokeArgs();
    }


    [OutputType]
    public sealed class GetSensitiveFileResult
    {
        /// <summary>
        /// Raw content of the file that was read, as UTF-8 encoded string. Files that do not contain UTF-8 text will have invalid UTF-8 sequences in `content`
        /// replaced with the Unicode replacement character.
        /// </summary>
        public readonly string Content;
        /// <summary>
        /// Base64 encoded version of the file content (use this when dealing with binary data).
        /// </summary>
        public readonly string ContentBase64;
        /// <summary>
        /// Base64 encoded SHA256 checksum of file content.
        /// </summary>
        public readonly string ContentBase64sha256;
        /// <summary>
        /// Base64 encoded SHA512 checksum of file content.
        /// </summary>
        public readonly string ContentBase64sha512;
        /// <summary>
        /// MD5 checksum of file content.
        /// </summary>
        public readonly string ContentMd5;
        /// <summary>
        /// SHA1 checksum of file content.
        /// </summary>
        public readonly string ContentSha1;
        /// <summary>
        /// SHA256 checksum of file content.
        /// </summary>
        public readonly string ContentSha256;
        /// <summary>
        /// SHA512 checksum of file content.
        /// </summary>
        public readonly string ContentSha512;
        /// <summary>
        /// Path to the file that will be read. The data source will return an error if the file does not exist.
        /// </summary>
        public readonly string Filename;
        /// <summary>
        /// The hexadecimal encoding of the SHA1 checksum of the file content.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSensitiveFileResult(
            string content,

            string contentBase64,

            string contentBase64sha256,

            string contentBase64sha512,

            string contentMd5,

            string contentSha1,

            string contentSha256,

            string contentSha512,

            string filename,

            string id)
        {
            Content = content;
            ContentBase64 = contentBase64;
            ContentBase64sha256 = contentBase64sha256;
            ContentBase64sha512 = contentBase64sha512;
            ContentMd5 = contentMd5;
            ContentSha1 = contentSha1;
            ContentSha256 = contentSha256;
            ContentSha512 = contentSha512;
            Filename = filename;
            Id = id;
        }
    }
}
