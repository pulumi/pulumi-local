# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['SensitiveFileArgs', 'SensitiveFile']

@pulumi.input_type
class SensitiveFileArgs:
    def __init__(__self__, *,
                 filename: pulumi.Input[str],
                 content: Optional[pulumi.Input[str]] = None,
                 content_base64: Optional[pulumi.Input[str]] = None,
                 directory_permission: Optional[pulumi.Input[str]] = None,
                 file_permission: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SensitiveFile resource.
        :param pulumi.Input[str] filename: The path to the file that will be created.
               Missing parent directories will be created.
               If the file already exists, it will be overridden with the given content.
        :param pulumi.Input[str] content: Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
               Conflicts with `content_base64` and `source`.
               Exactly one of these three arguments must be specified.
        :param pulumi.Input[str] content_base64: Sensitive Content to store in the file, expected to be binary encoded as base64 string.
               Conflicts with `content` and `source`.
               Exactly one of these three arguments must be specified.
        :param pulumi.Input[str] directory_permission: Permissions to set for directories created (before umask), expressed as string in
               [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
               Default value is `"0700"`.
        :param pulumi.Input[str] file_permission: Permissions to set for the output file (before umask), expressed as string in
               [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
               Default value is `"0700"`.
        :param pulumi.Input[str] source: Path to file to use as source for the one we are creating.
               Conflicts with `content` and `content_base64`.
               Exactly one of these three arguments must be specified.
        """
        pulumi.set(__self__, "filename", filename)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if content_base64 is not None:
            pulumi.set(__self__, "content_base64", content_base64)
        if directory_permission is not None:
            pulumi.set(__self__, "directory_permission", directory_permission)
        if file_permission is not None:
            pulumi.set(__self__, "file_permission", file_permission)
        if source is not None:
            pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Input[str]:
        """
        The path to the file that will be created.
        Missing parent directories will be created.
        If the file already exists, it will be overridden with the given content.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: pulumi.Input[str]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
        Conflicts with `content_base64` and `source`.
        Exactly one of these three arguments must be specified.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> Optional[pulumi.Input[str]]:
        """
        Sensitive Content to store in the file, expected to be binary encoded as base64 string.
        Conflicts with `content` and `source`.
        Exactly one of these three arguments must be specified.
        """
        return pulumi.get(self, "content_base64")

    @content_base64.setter
    def content_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_base64", value)

    @property
    @pulumi.getter(name="directoryPermission")
    def directory_permission(self) -> Optional[pulumi.Input[str]]:
        """
        Permissions to set for directories created (before umask), expressed as string in
        [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
        Default value is `"0700"`.
        """
        return pulumi.get(self, "directory_permission")

    @directory_permission.setter
    def directory_permission(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_permission", value)

    @property
    @pulumi.getter(name="filePermission")
    def file_permission(self) -> Optional[pulumi.Input[str]]:
        """
        Permissions to set for the output file (before umask), expressed as string in
        [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
        Default value is `"0700"`.
        """
        return pulumi.get(self, "file_permission")

    @file_permission.setter
    def file_permission(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_permission", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Path to file to use as source for the one we are creating.
        Conflicts with `content` and `content_base64`.
        Exactly one of these three arguments must be specified.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)


@pulumi.input_type
class _SensitiveFileState:
    def __init__(__self__, *,
                 content: Optional[pulumi.Input[str]] = None,
                 content_base64: Optional[pulumi.Input[str]] = None,
                 content_base64sha256: Optional[pulumi.Input[str]] = None,
                 content_base64sha512: Optional[pulumi.Input[str]] = None,
                 content_md5: Optional[pulumi.Input[str]] = None,
                 content_sha1: Optional[pulumi.Input[str]] = None,
                 content_sha256: Optional[pulumi.Input[str]] = None,
                 content_sha512: Optional[pulumi.Input[str]] = None,
                 directory_permission: Optional[pulumi.Input[str]] = None,
                 file_permission: Optional[pulumi.Input[str]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SensitiveFile resources.
        :param pulumi.Input[str] content: Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
               Conflicts with `content_base64` and `source`.
               Exactly one of these three arguments must be specified.
        :param pulumi.Input[str] content_base64: Sensitive Content to store in the file, expected to be binary encoded as base64 string.
               Conflicts with `content` and `source`.
               Exactly one of these three arguments must be specified.
        :param pulumi.Input[str] content_base64sha256: Base64 encoded SHA256 checksum of file content.
        :param pulumi.Input[str] content_base64sha512: Base64 encoded SHA512 checksum of file content.
        :param pulumi.Input[str] content_md5: MD5 checksum of file content.
        :param pulumi.Input[str] content_sha1: SHA1 checksum of file content.
        :param pulumi.Input[str] content_sha256: SHA256 checksum of file content.
        :param pulumi.Input[str] content_sha512: SHA512 checksum of file content.
        :param pulumi.Input[str] directory_permission: Permissions to set for directories created (before umask), expressed as string in
               [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
               Default value is `"0700"`.
        :param pulumi.Input[str] file_permission: Permissions to set for the output file (before umask), expressed as string in
               [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
               Default value is `"0700"`.
        :param pulumi.Input[str] filename: The path to the file that will be created.
               Missing parent directories will be created.
               If the file already exists, it will be overridden with the given content.
        :param pulumi.Input[str] source: Path to file to use as source for the one we are creating.
               Conflicts with `content` and `content_base64`.
               Exactly one of these three arguments must be specified.
        """
        if content is not None:
            pulumi.set(__self__, "content", content)
        if content_base64 is not None:
            pulumi.set(__self__, "content_base64", content_base64)
        if content_base64sha256 is not None:
            pulumi.set(__self__, "content_base64sha256", content_base64sha256)
        if content_base64sha512 is not None:
            pulumi.set(__self__, "content_base64sha512", content_base64sha512)
        if content_md5 is not None:
            pulumi.set(__self__, "content_md5", content_md5)
        if content_sha1 is not None:
            pulumi.set(__self__, "content_sha1", content_sha1)
        if content_sha256 is not None:
            pulumi.set(__self__, "content_sha256", content_sha256)
        if content_sha512 is not None:
            pulumi.set(__self__, "content_sha512", content_sha512)
        if directory_permission is not None:
            pulumi.set(__self__, "directory_permission", directory_permission)
        if file_permission is not None:
            pulumi.set(__self__, "file_permission", file_permission)
        if filename is not None:
            pulumi.set(__self__, "filename", filename)
        if source is not None:
            pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
        Conflicts with `content_base64` and `source`.
        Exactly one of these three arguments must be specified.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> Optional[pulumi.Input[str]]:
        """
        Sensitive Content to store in the file, expected to be binary encoded as base64 string.
        Conflicts with `content` and `source`.
        Exactly one of these three arguments must be specified.
        """
        return pulumi.get(self, "content_base64")

    @content_base64.setter
    def content_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_base64", value)

    @property
    @pulumi.getter(name="contentBase64sha256")
    def content_base64sha256(self) -> Optional[pulumi.Input[str]]:
        """
        Base64 encoded SHA256 checksum of file content.
        """
        return pulumi.get(self, "content_base64sha256")

    @content_base64sha256.setter
    def content_base64sha256(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_base64sha256", value)

    @property
    @pulumi.getter(name="contentBase64sha512")
    def content_base64sha512(self) -> Optional[pulumi.Input[str]]:
        """
        Base64 encoded SHA512 checksum of file content.
        """
        return pulumi.get(self, "content_base64sha512")

    @content_base64sha512.setter
    def content_base64sha512(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_base64sha512", value)

    @property
    @pulumi.getter(name="contentMd5")
    def content_md5(self) -> Optional[pulumi.Input[str]]:
        """
        MD5 checksum of file content.
        """
        return pulumi.get(self, "content_md5")

    @content_md5.setter
    def content_md5(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_md5", value)

    @property
    @pulumi.getter(name="contentSha1")
    def content_sha1(self) -> Optional[pulumi.Input[str]]:
        """
        SHA1 checksum of file content.
        """
        return pulumi.get(self, "content_sha1")

    @content_sha1.setter
    def content_sha1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_sha1", value)

    @property
    @pulumi.getter(name="contentSha256")
    def content_sha256(self) -> Optional[pulumi.Input[str]]:
        """
        SHA256 checksum of file content.
        """
        return pulumi.get(self, "content_sha256")

    @content_sha256.setter
    def content_sha256(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_sha256", value)

    @property
    @pulumi.getter(name="contentSha512")
    def content_sha512(self) -> Optional[pulumi.Input[str]]:
        """
        SHA512 checksum of file content.
        """
        return pulumi.get(self, "content_sha512")

    @content_sha512.setter
    def content_sha512(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_sha512", value)

    @property
    @pulumi.getter(name="directoryPermission")
    def directory_permission(self) -> Optional[pulumi.Input[str]]:
        """
        Permissions to set for directories created (before umask), expressed as string in
        [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
        Default value is `"0700"`.
        """
        return pulumi.get(self, "directory_permission")

    @directory_permission.setter
    def directory_permission(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_permission", value)

    @property
    @pulumi.getter(name="filePermission")
    def file_permission(self) -> Optional[pulumi.Input[str]]:
        """
        Permissions to set for the output file (before umask), expressed as string in
        [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
        Default value is `"0700"`.
        """
        return pulumi.get(self, "file_permission")

    @file_permission.setter
    def file_permission(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_permission", value)

    @property
    @pulumi.getter
    def filename(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the file that will be created.
        Missing parent directories will be created.
        If the file already exists, it will be overridden with the given content.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Path to file to use as source for the one we are creating.
        Conflicts with `content` and `content_base64`.
        Exactly one of these three arguments must be specified.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)


class SensitiveFile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_base64: Optional[pulumi.Input[str]] = None,
                 directory_permission: Optional[pulumi.Input[str]] = None,
                 file_permission: Optional[pulumi.Input[str]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content: Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
               Conflicts with `content_base64` and `source`.
               Exactly one of these three arguments must be specified.
        :param pulumi.Input[str] content_base64: Sensitive Content to store in the file, expected to be binary encoded as base64 string.
               Conflicts with `content` and `source`.
               Exactly one of these three arguments must be specified.
        :param pulumi.Input[str] directory_permission: Permissions to set for directories created (before umask), expressed as string in
               [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
               Default value is `"0700"`.
        :param pulumi.Input[str] file_permission: Permissions to set for the output file (before umask), expressed as string in
               [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
               Default value is `"0700"`.
        :param pulumi.Input[str] filename: The path to the file that will be created.
               Missing parent directories will be created.
               If the file already exists, it will be overridden with the given content.
        :param pulumi.Input[str] source: Path to file to use as source for the one we are creating.
               Conflicts with `content` and `content_base64`.
               Exactly one of these three arguments must be specified.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SensitiveFileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        :param str resource_name: The name of the resource.
        :param SensitiveFileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SensitiveFileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_base64: Optional[pulumi.Input[str]] = None,
                 directory_permission: Optional[pulumi.Input[str]] = None,
                 file_permission: Optional[pulumi.Input[str]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SensitiveFileArgs.__new__(SensitiveFileArgs)

            __props__.__dict__["content"] = None if content is None else pulumi.Output.secret(content)
            __props__.__dict__["content_base64"] = None if content_base64 is None else pulumi.Output.secret(content_base64)
            __props__.__dict__["directory_permission"] = directory_permission
            __props__.__dict__["file_permission"] = file_permission
            if filename is None and not opts.urn:
                raise TypeError("Missing required property 'filename'")
            __props__.__dict__["filename"] = filename
            __props__.__dict__["source"] = source
            __props__.__dict__["content_base64sha256"] = None
            __props__.__dict__["content_base64sha512"] = None
            __props__.__dict__["content_md5"] = None
            __props__.__dict__["content_sha1"] = None
            __props__.__dict__["content_sha256"] = None
            __props__.__dict__["content_sha512"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["content", "contentBase64"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SensitiveFile, __self__).__init__(
            'local:index/sensitiveFile:SensitiveFile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            content: Optional[pulumi.Input[str]] = None,
            content_base64: Optional[pulumi.Input[str]] = None,
            content_base64sha256: Optional[pulumi.Input[str]] = None,
            content_base64sha512: Optional[pulumi.Input[str]] = None,
            content_md5: Optional[pulumi.Input[str]] = None,
            content_sha1: Optional[pulumi.Input[str]] = None,
            content_sha256: Optional[pulumi.Input[str]] = None,
            content_sha512: Optional[pulumi.Input[str]] = None,
            directory_permission: Optional[pulumi.Input[str]] = None,
            file_permission: Optional[pulumi.Input[str]] = None,
            filename: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None) -> 'SensitiveFile':
        """
        Get an existing SensitiveFile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content: Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
               Conflicts with `content_base64` and `source`.
               Exactly one of these three arguments must be specified.
        :param pulumi.Input[str] content_base64: Sensitive Content to store in the file, expected to be binary encoded as base64 string.
               Conflicts with `content` and `source`.
               Exactly one of these three arguments must be specified.
        :param pulumi.Input[str] content_base64sha256: Base64 encoded SHA256 checksum of file content.
        :param pulumi.Input[str] content_base64sha512: Base64 encoded SHA512 checksum of file content.
        :param pulumi.Input[str] content_md5: MD5 checksum of file content.
        :param pulumi.Input[str] content_sha1: SHA1 checksum of file content.
        :param pulumi.Input[str] content_sha256: SHA256 checksum of file content.
        :param pulumi.Input[str] content_sha512: SHA512 checksum of file content.
        :param pulumi.Input[str] directory_permission: Permissions to set for directories created (before umask), expressed as string in
               [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
               Default value is `"0700"`.
        :param pulumi.Input[str] file_permission: Permissions to set for the output file (before umask), expressed as string in
               [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
               Default value is `"0700"`.
        :param pulumi.Input[str] filename: The path to the file that will be created.
               Missing parent directories will be created.
               If the file already exists, it will be overridden with the given content.
        :param pulumi.Input[str] source: Path to file to use as source for the one we are creating.
               Conflicts with `content` and `content_base64`.
               Exactly one of these three arguments must be specified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SensitiveFileState.__new__(_SensitiveFileState)

        __props__.__dict__["content"] = content
        __props__.__dict__["content_base64"] = content_base64
        __props__.__dict__["content_base64sha256"] = content_base64sha256
        __props__.__dict__["content_base64sha512"] = content_base64sha512
        __props__.__dict__["content_md5"] = content_md5
        __props__.__dict__["content_sha1"] = content_sha1
        __props__.__dict__["content_sha256"] = content_sha256
        __props__.__dict__["content_sha512"] = content_sha512
        __props__.__dict__["directory_permission"] = directory_permission
        __props__.__dict__["file_permission"] = file_permission
        __props__.__dict__["filename"] = filename
        __props__.__dict__["source"] = source
        return SensitiveFile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[Optional[str]]:
        """
        Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
        Conflicts with `content_base64` and `source`.
        Exactly one of these three arguments must be specified.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> pulumi.Output[Optional[str]]:
        """
        Sensitive Content to store in the file, expected to be binary encoded as base64 string.
        Conflicts with `content` and `source`.
        Exactly one of these three arguments must be specified.
        """
        return pulumi.get(self, "content_base64")

    @property
    @pulumi.getter(name="contentBase64sha256")
    def content_base64sha256(self) -> pulumi.Output[str]:
        """
        Base64 encoded SHA256 checksum of file content.
        """
        return pulumi.get(self, "content_base64sha256")

    @property
    @pulumi.getter(name="contentBase64sha512")
    def content_base64sha512(self) -> pulumi.Output[str]:
        """
        Base64 encoded SHA512 checksum of file content.
        """
        return pulumi.get(self, "content_base64sha512")

    @property
    @pulumi.getter(name="contentMd5")
    def content_md5(self) -> pulumi.Output[str]:
        """
        MD5 checksum of file content.
        """
        return pulumi.get(self, "content_md5")

    @property
    @pulumi.getter(name="contentSha1")
    def content_sha1(self) -> pulumi.Output[str]:
        """
        SHA1 checksum of file content.
        """
        return pulumi.get(self, "content_sha1")

    @property
    @pulumi.getter(name="contentSha256")
    def content_sha256(self) -> pulumi.Output[str]:
        """
        SHA256 checksum of file content.
        """
        return pulumi.get(self, "content_sha256")

    @property
    @pulumi.getter(name="contentSha512")
    def content_sha512(self) -> pulumi.Output[str]:
        """
        SHA512 checksum of file content.
        """
        return pulumi.get(self, "content_sha512")

    @property
    @pulumi.getter(name="directoryPermission")
    def directory_permission(self) -> pulumi.Output[str]:
        """
        Permissions to set for directories created (before umask), expressed as string in
        [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
        Default value is `"0700"`.
        """
        return pulumi.get(self, "directory_permission")

    @property
    @pulumi.getter(name="filePermission")
    def file_permission(self) -> pulumi.Output[str]:
        """
        Permissions to set for the output file (before umask), expressed as string in
        [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
        Default value is `"0700"`.
        """
        return pulumi.get(self, "file_permission")

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Output[str]:
        """
        The path to the file that will be created.
        Missing parent directories will be created.
        If the file already exists, it will be overridden with the given content.
        """
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[Optional[str]]:
        """
        Path to file to use as source for the one we are creating.
        Conflicts with `content` and `content_base64`.
        Exactly one of these three arguments must be specified.
        """
        return pulumi.get(self, "source")

