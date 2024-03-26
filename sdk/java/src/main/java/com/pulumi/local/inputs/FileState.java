// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.local.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FileState extends com.pulumi.resources.ResourceArgs {

    public static final FileState Empty = new FileState();

    /**
     * Content to store in the file, expected to be a UTF-8 encoded string.
     * Conflicts with `sensitive_content`, `content_base64` and `source`.
     * Exactly one of these four arguments must be specified.
     * 
     */
    @Import(name="content")
    private @Nullable Output<String> content;

    /**
     * @return Content to store in the file, expected to be a UTF-8 encoded string.
     * Conflicts with `sensitive_content`, `content_base64` and `source`.
     * Exactly one of these four arguments must be specified.
     * 
     */
    public Optional<Output<String>> content() {
        return Optional.ofNullable(this.content);
    }

    /**
     * Content to store in the file, expected to be binary encoded as base64 string.
     * Conflicts with `content`, `sensitive_content` and `source`.
     * Exactly one of these four arguments must be specified.
     * 
     */
    @Import(name="contentBase64")
    private @Nullable Output<String> contentBase64;

    /**
     * @return Content to store in the file, expected to be binary encoded as base64 string.
     * Conflicts with `content`, `sensitive_content` and `source`.
     * Exactly one of these four arguments must be specified.
     * 
     */
    public Optional<Output<String>> contentBase64() {
        return Optional.ofNullable(this.contentBase64);
    }

    /**
     * Base64 encoded SHA256 checksum of file content.
     * 
     */
    @Import(name="contentBase64sha256")
    private @Nullable Output<String> contentBase64sha256;

    /**
     * @return Base64 encoded SHA256 checksum of file content.
     * 
     */
    public Optional<Output<String>> contentBase64sha256() {
        return Optional.ofNullable(this.contentBase64sha256);
    }

    /**
     * Base64 encoded SHA512 checksum of file content.
     * 
     */
    @Import(name="contentBase64sha512")
    private @Nullable Output<String> contentBase64sha512;

    /**
     * @return Base64 encoded SHA512 checksum of file content.
     * 
     */
    public Optional<Output<String>> contentBase64sha512() {
        return Optional.ofNullable(this.contentBase64sha512);
    }

    /**
     * MD5 checksum of file content.
     * 
     */
    @Import(name="contentMd5")
    private @Nullable Output<String> contentMd5;

    /**
     * @return MD5 checksum of file content.
     * 
     */
    public Optional<Output<String>> contentMd5() {
        return Optional.ofNullable(this.contentMd5);
    }

    /**
     * SHA1 checksum of file content.
     * 
     */
    @Import(name="contentSha1")
    private @Nullable Output<String> contentSha1;

    /**
     * @return SHA1 checksum of file content.
     * 
     */
    public Optional<Output<String>> contentSha1() {
        return Optional.ofNullable(this.contentSha1);
    }

    /**
     * SHA256 checksum of file content.
     * 
     */
    @Import(name="contentSha256")
    private @Nullable Output<String> contentSha256;

    /**
     * @return SHA256 checksum of file content.
     * 
     */
    public Optional<Output<String>> contentSha256() {
        return Optional.ofNullable(this.contentSha256);
    }

    /**
     * SHA512 checksum of file content.
     * 
     */
    @Import(name="contentSha512")
    private @Nullable Output<String> contentSha512;

    /**
     * @return SHA512 checksum of file content.
     * 
     */
    public Optional<Output<String>> contentSha512() {
        return Optional.ofNullable(this.contentSha512);
    }

    /**
     * Permissions to set for directories created (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `&#34;0777&#34;`.
     * 
     */
    @Import(name="directoryPermission")
    private @Nullable Output<String> directoryPermission;

    /**
     * @return Permissions to set for directories created (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `&#34;0777&#34;`.
     * 
     */
    public Optional<Output<String>> directoryPermission() {
        return Optional.ofNullable(this.directoryPermission);
    }

    /**
     * Permissions to set for the output file (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `&#34;0777&#34;`.
     * 
     */
    @Import(name="filePermission")
    private @Nullable Output<String> filePermission;

    /**
     * @return Permissions to set for the output file (before umask), expressed as string in
     * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
     * Default value is `&#34;0777&#34;`.
     * 
     */
    public Optional<Output<String>> filePermission() {
        return Optional.ofNullable(this.filePermission);
    }

    /**
     * The path to the file that will be created.
     * Missing parent directories will be created.
     * If the file already exists, it will be overridden with the given content.
     * 
     */
    @Import(name="filename")
    private @Nullable Output<String> filename;

    /**
     * @return The path to the file that will be created.
     * Missing parent directories will be created.
     * If the file already exists, it will be overridden with the given content.
     * 
     */
    public Optional<Output<String>> filename() {
        return Optional.ofNullable(this.filename);
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
     * Use the `local_sensitive_file` resource instead
     * 
     */
    @Deprecated /* Use the `local_sensitive_file` resource instead */
    @Import(name="sensitiveContent")
    private @Nullable Output<String> sensitiveContent;

    /**
     * @return Sensitive content to store in the file, expected to be an UTF-8 encoded string.
     * Will not be displayed in diffs.
     * Conflicts with `content`, `content_base64` and `source`.
     * Exactly one of these four arguments must be specified.
     * If in need to use *sensitive* content, please use the `local.SensitiveFile`
     * resource instead.
     * 
     * @deprecated
     * Use the `local_sensitive_file` resource instead
     * 
     */
    @Deprecated /* Use the `local_sensitive_file` resource instead */
    public Optional<Output<String>> sensitiveContent() {
        return Optional.ofNullable(this.sensitiveContent);
    }

    /**
     * Path to file to use as source for the one we are creating.
     * Conflicts with `content`, `sensitive_content` and `content_base64`.
     * Exactly one of these four arguments must be specified.
     * 
     */
    @Import(name="source")
    private @Nullable Output<String> source;

    /**
     * @return Path to file to use as source for the one we are creating.
     * Conflicts with `content`, `sensitive_content` and `content_base64`.
     * Exactly one of these four arguments must be specified.
     * 
     */
    public Optional<Output<String>> source() {
        return Optional.ofNullable(this.source);
    }

    private FileState() {}

    private FileState(FileState $) {
        this.content = $.content;
        this.contentBase64 = $.contentBase64;
        this.contentBase64sha256 = $.contentBase64sha256;
        this.contentBase64sha512 = $.contentBase64sha512;
        this.contentMd5 = $.contentMd5;
        this.contentSha1 = $.contentSha1;
        this.contentSha256 = $.contentSha256;
        this.contentSha512 = $.contentSha512;
        this.directoryPermission = $.directoryPermission;
        this.filePermission = $.filePermission;
        this.filename = $.filename;
        this.sensitiveContent = $.sensitiveContent;
        this.source = $.source;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FileState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FileState $;

        public Builder() {
            $ = new FileState();
        }

        public Builder(FileState defaults) {
            $ = new FileState(Objects.requireNonNull(defaults));
        }

        /**
         * @param content Content to store in the file, expected to be a UTF-8 encoded string.
         * Conflicts with `sensitive_content`, `content_base64` and `source`.
         * Exactly one of these four arguments must be specified.
         * 
         * @return builder
         * 
         */
        public Builder content(@Nullable Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content Content to store in the file, expected to be a UTF-8 encoded string.
         * Conflicts with `sensitive_content`, `content_base64` and `source`.
         * Exactly one of these four arguments must be specified.
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        /**
         * @param contentBase64 Content to store in the file, expected to be binary encoded as base64 string.
         * Conflicts with `content`, `sensitive_content` and `source`.
         * Exactly one of these four arguments must be specified.
         * 
         * @return builder
         * 
         */
        public Builder contentBase64(@Nullable Output<String> contentBase64) {
            $.contentBase64 = contentBase64;
            return this;
        }

        /**
         * @param contentBase64 Content to store in the file, expected to be binary encoded as base64 string.
         * Conflicts with `content`, `sensitive_content` and `source`.
         * Exactly one of these four arguments must be specified.
         * 
         * @return builder
         * 
         */
        public Builder contentBase64(String contentBase64) {
            return contentBase64(Output.of(contentBase64));
        }

        /**
         * @param contentBase64sha256 Base64 encoded SHA256 checksum of file content.
         * 
         * @return builder
         * 
         */
        public Builder contentBase64sha256(@Nullable Output<String> contentBase64sha256) {
            $.contentBase64sha256 = contentBase64sha256;
            return this;
        }

        /**
         * @param contentBase64sha256 Base64 encoded SHA256 checksum of file content.
         * 
         * @return builder
         * 
         */
        public Builder contentBase64sha256(String contentBase64sha256) {
            return contentBase64sha256(Output.of(contentBase64sha256));
        }

        /**
         * @param contentBase64sha512 Base64 encoded SHA512 checksum of file content.
         * 
         * @return builder
         * 
         */
        public Builder contentBase64sha512(@Nullable Output<String> contentBase64sha512) {
            $.contentBase64sha512 = contentBase64sha512;
            return this;
        }

        /**
         * @param contentBase64sha512 Base64 encoded SHA512 checksum of file content.
         * 
         * @return builder
         * 
         */
        public Builder contentBase64sha512(String contentBase64sha512) {
            return contentBase64sha512(Output.of(contentBase64sha512));
        }

        /**
         * @param contentMd5 MD5 checksum of file content.
         * 
         * @return builder
         * 
         */
        public Builder contentMd5(@Nullable Output<String> contentMd5) {
            $.contentMd5 = contentMd5;
            return this;
        }

        /**
         * @param contentMd5 MD5 checksum of file content.
         * 
         * @return builder
         * 
         */
        public Builder contentMd5(String contentMd5) {
            return contentMd5(Output.of(contentMd5));
        }

        /**
         * @param contentSha1 SHA1 checksum of file content.
         * 
         * @return builder
         * 
         */
        public Builder contentSha1(@Nullable Output<String> contentSha1) {
            $.contentSha1 = contentSha1;
            return this;
        }

        /**
         * @param contentSha1 SHA1 checksum of file content.
         * 
         * @return builder
         * 
         */
        public Builder contentSha1(String contentSha1) {
            return contentSha1(Output.of(contentSha1));
        }

        /**
         * @param contentSha256 SHA256 checksum of file content.
         * 
         * @return builder
         * 
         */
        public Builder contentSha256(@Nullable Output<String> contentSha256) {
            $.contentSha256 = contentSha256;
            return this;
        }

        /**
         * @param contentSha256 SHA256 checksum of file content.
         * 
         * @return builder
         * 
         */
        public Builder contentSha256(String contentSha256) {
            return contentSha256(Output.of(contentSha256));
        }

        /**
         * @param contentSha512 SHA512 checksum of file content.
         * 
         * @return builder
         * 
         */
        public Builder contentSha512(@Nullable Output<String> contentSha512) {
            $.contentSha512 = contentSha512;
            return this;
        }

        /**
         * @param contentSha512 SHA512 checksum of file content.
         * 
         * @return builder
         * 
         */
        public Builder contentSha512(String contentSha512) {
            return contentSha512(Output.of(contentSha512));
        }

        /**
         * @param directoryPermission Permissions to set for directories created (before umask), expressed as string in
         * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
         * Default value is `&#34;0777&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder directoryPermission(@Nullable Output<String> directoryPermission) {
            $.directoryPermission = directoryPermission;
            return this;
        }

        /**
         * @param directoryPermission Permissions to set for directories created (before umask), expressed as string in
         * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
         * Default value is `&#34;0777&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder directoryPermission(String directoryPermission) {
            return directoryPermission(Output.of(directoryPermission));
        }

        /**
         * @param filePermission Permissions to set for the output file (before umask), expressed as string in
         * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
         * Default value is `&#34;0777&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder filePermission(@Nullable Output<String> filePermission) {
            $.filePermission = filePermission;
            return this;
        }

        /**
         * @param filePermission Permissions to set for the output file (before umask), expressed as string in
         * [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
         * Default value is `&#34;0777&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder filePermission(String filePermission) {
            return filePermission(Output.of(filePermission));
        }

        /**
         * @param filename The path to the file that will be created.
         * Missing parent directories will be created.
         * If the file already exists, it will be overridden with the given content.
         * 
         * @return builder
         * 
         */
        public Builder filename(@Nullable Output<String> filename) {
            $.filename = filename;
            return this;
        }

        /**
         * @param filename The path to the file that will be created.
         * Missing parent directories will be created.
         * If the file already exists, it will be overridden with the given content.
         * 
         * @return builder
         * 
         */
        public Builder filename(String filename) {
            return filename(Output.of(filename));
        }

        /**
         * @param sensitiveContent Sensitive content to store in the file, expected to be an UTF-8 encoded string.
         * Will not be displayed in diffs.
         * Conflicts with `content`, `content_base64` and `source`.
         * Exactly one of these four arguments must be specified.
         * If in need to use *sensitive* content, please use the `local.SensitiveFile`
         * resource instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Use the `local_sensitive_file` resource instead
         * 
         */
        @Deprecated /* Use the `local_sensitive_file` resource instead */
        public Builder sensitiveContent(@Nullable Output<String> sensitiveContent) {
            $.sensitiveContent = sensitiveContent;
            return this;
        }

        /**
         * @param sensitiveContent Sensitive content to store in the file, expected to be an UTF-8 encoded string.
         * Will not be displayed in diffs.
         * Conflicts with `content`, `content_base64` and `source`.
         * Exactly one of these four arguments must be specified.
         * If in need to use *sensitive* content, please use the `local.SensitiveFile`
         * resource instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Use the `local_sensitive_file` resource instead
         * 
         */
        @Deprecated /* Use the `local_sensitive_file` resource instead */
        public Builder sensitiveContent(String sensitiveContent) {
            return sensitiveContent(Output.of(sensitiveContent));
        }

        /**
         * @param source Path to file to use as source for the one we are creating.
         * Conflicts with `content`, `sensitive_content` and `content_base64`.
         * Exactly one of these four arguments must be specified.
         * 
         * @return builder
         * 
         */
        public Builder source(@Nullable Output<String> source) {
            $.source = source;
            return this;
        }

        /**
         * @param source Path to file to use as source for the one we are creating.
         * Conflicts with `content`, `sensitive_content` and `content_base64`.
         * Exactly one of these four arguments must be specified.
         * 
         * @return builder
         * 
         */
        public Builder source(String source) {
            return source(Output.of(source));
        }

        public FileState build() {
            return $;
        }
    }

}
