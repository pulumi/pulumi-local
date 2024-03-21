// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.local;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.local.FileArgs;
import com.pulumi.local.Utilities;
import com.pulumi.local.inputs.FileState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.local.File;
 * import com.pulumi.local.FileArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var foo = new File(&#34;foo&#34;, FileArgs.builder()        
 *             .content(&#34;foo!&#34;)
 *             .filename(String.format(&#34;%s/foo.bar&#34;, path.module()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="local:index/file:File")
public class File extends com.pulumi.resources.CustomResource {
    /**
     * Content to store in the file, expected to be a UTF-8 encoded string.
     * Conflicts with `sensitive_content`, `content_base64` and `source`.
     * Exactly one of these four arguments must be specified.
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> content;

    /**
     * @return Content to store in the file, expected to be a UTF-8 encoded string.
     * Conflicts with `sensitive_content`, `content_base64` and `source`.
     * Exactly one of these four arguments must be specified.
     * 
     */
    public Output<Optional<String>> content() {
        return Codegen.optional(this.content);
    }
    /**
     * Content to store in the file, expected to be binary encoded as base64 string.
     * Conflicts with `content`, `sensitive_content` and `source`.
     * Exactly one of these four arguments must be specified.
     * 
     */
    @Export(name="contentBase64", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> contentBase64;

    /**
     * @return Content to store in the file, expected to be binary encoded as base64 string.
     * Conflicts with `content`, `sensitive_content` and `source`.
     * Exactly one of these four arguments must be specified.
     * 
     */
    public Output<Optional<String>> contentBase64() {
        return Codegen.optional(this.contentBase64);
    }
    /**
     * Base64 encoded SHA256 checksum of file content.
     * 
     */
    @Export(name="contentBase64sha256", refs={String.class}, tree="[0]")
    private Output<String> contentBase64sha256;

    /**
     * @return Base64 encoded SHA256 checksum of file content.
     * 
     */
    public Output<String> contentBase64sha256() {
        return this.contentBase64sha256;
    }
    /**
     * Base64 encoded SHA512 checksum of file content.
     * 
     */
    @Export(name="contentBase64sha512", refs={String.class}, tree="[0]")
    private Output<String> contentBase64sha512;

    /**
     * @return Base64 encoded SHA512 checksum of file content.
     * 
     */
    public Output<String> contentBase64sha512() {
        return this.contentBase64sha512;
    }
    /**
     * MD5 checksum of file content.
     * 
     */
    @Export(name="contentMd5", refs={String.class}, tree="[0]")
    private Output<String> contentMd5;

    /**
     * @return MD5 checksum of file content.
     * 
     */
    public Output<String> contentMd5() {
        return this.contentMd5;
    }
    /**
     * SHA1 checksum of file content.
     * 
     */
    @Export(name="contentSha1", refs={String.class}, tree="[0]")
    private Output<String> contentSha1;

    /**
     * @return SHA1 checksum of file content.
     * 
     */
    public Output<String> contentSha1() {
        return this.contentSha1;
    }
    /**
     * SHA256 checksum of file content.
     * 
     */
    @Export(name="contentSha256", refs={String.class}, tree="[0]")
    private Output<String> contentSha256;

    /**
     * @return SHA256 checksum of file content.
     * 
     */
    public Output<String> contentSha256() {
        return this.contentSha256;
    }
    /**
     * SHA512 checksum of file content.
     * 
     */
    @Export(name="contentSha512", refs={String.class}, tree="[0]")
    private Output<String> contentSha512;

    /**
     * @return SHA512 checksum of file content.
     * 
     */
    public Output<String> contentSha512() {
        return this.contentSha512;
    }
    /**
     * Permissions to set for directories created (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `&#34;0777&#34;`.
     * 
     */
    @Export(name="directoryPermission", refs={String.class}, tree="[0]")
    private Output<String> directoryPermission;

    /**
     * @return Permissions to set for directories created (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `&#34;0777&#34;`.
     * 
     */
    public Output<String> directoryPermission() {
        return this.directoryPermission;
    }
    /**
     * Permissions to set for the output file (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `&#34;0777&#34;`.
     * 
     */
    @Export(name="filePermission", refs={String.class}, tree="[0]")
    private Output<String> filePermission;

    /**
     * @return Permissions to set for the output file (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `&#34;0777&#34;`.
     * 
     */
    public Output<String> filePermission() {
        return this.filePermission;
    }
    /**
     * The path to the file that will be created.
     * Missing parent directories will be created.
     * If the file already exists, it will be overridden with the given content.
     * 
     */
    @Export(name="filename", refs={String.class}, tree="[0]")
    private Output<String> filename;

    /**
     * @return The path to the file that will be created.
     * Missing parent directories will be created.
     * If the file already exists, it will be overridden with the given content.
     * 
     */
    public Output<String> filename() {
        return this.filename;
    }
    /**
     * Sensitive content to store in the file, expected to be an UTF-8 encoded string.
     * Will not be displayed in diffs.
     * Conflicts with `content`, `content_base64` and `source`.
     * Exactly one of these four arguments must be specified.
     * If in need to use *sensitive* content, please use the `local.SensitiveFile`
     * resource instead.
     * 
     * @deprecated
     * Use the `local.SensitiveFile` resource instead
     * 
     */
    @Deprecated /* Use the `local.SensitiveFile` resource instead */
    @Export(name="sensitiveContent", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sensitiveContent;

    /**
     * @return Sensitive content to store in the file, expected to be an UTF-8 encoded string.
     * Will not be displayed in diffs.
     * Conflicts with `content`, `content_base64` and `source`.
     * Exactly one of these four arguments must be specified.
     * If in need to use *sensitive* content, please use the `local.SensitiveFile`
     * resource instead.
     * 
     */
    public Output<Optional<String>> sensitiveContent() {
        return Codegen.optional(this.sensitiveContent);
    }
    /**
     * Path to file to use as source for the one we are creating.
     * Conflicts with `content`, `sensitive_content` and `content_base64`.
     * Exactly one of these four arguments must be specified.
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> source;

    /**
     * @return Path to file to use as source for the one we are creating.
     * Conflicts with `content`, `sensitive_content` and `content_base64`.
     * Exactly one of these four arguments must be specified.
     * 
     */
    public Output<Optional<String>> source() {
        return Codegen.optional(this.source);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public File(String name) {
        this(name, FileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public File(String name, FileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public File(String name, FileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("local:index/file:File", name, args == null ? FileArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private File(String name, Output<String> id, @Nullable FileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("local:index/file:File", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "sensitiveContent"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static File get(String name, Output<String> id, @Nullable FileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new File(name, id, state, options);
    }
}
