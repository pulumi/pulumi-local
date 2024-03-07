# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetFileResult',
    'AwaitableGetFileResult',
    'get_file',
    'get_file_output',
]

@pulumi.output_type
class GetFileResult:
    """
    A collection of values returned by getFile.
    """
    def __init__(__self__, content=None, content_base64=None, content_base64sha256=None, content_base64sha512=None, content_md5=None, content_sha1=None, content_sha256=None, content_sha512=None, filename=None, id=None):
        if content and not isinstance(content, str):
            raise TypeError("Expected argument 'content' to be a str")
        pulumi.set(__self__, "content", content)
        if content_base64 and not isinstance(content_base64, str):
            raise TypeError("Expected argument 'content_base64' to be a str")
        pulumi.set(__self__, "content_base64", content_base64)
        if content_base64sha256 and not isinstance(content_base64sha256, str):
            raise TypeError("Expected argument 'content_base64sha256' to be a str")
        pulumi.set(__self__, "content_base64sha256", content_base64sha256)
        if content_base64sha512 and not isinstance(content_base64sha512, str):
            raise TypeError("Expected argument 'content_base64sha512' to be a str")
        pulumi.set(__self__, "content_base64sha512", content_base64sha512)
        if content_md5 and not isinstance(content_md5, str):
            raise TypeError("Expected argument 'content_md5' to be a str")
        pulumi.set(__self__, "content_md5", content_md5)
        if content_sha1 and not isinstance(content_sha1, str):
            raise TypeError("Expected argument 'content_sha1' to be a str")
        pulumi.set(__self__, "content_sha1", content_sha1)
        if content_sha256 and not isinstance(content_sha256, str):
            raise TypeError("Expected argument 'content_sha256' to be a str")
        pulumi.set(__self__, "content_sha256", content_sha256)
        if content_sha512 and not isinstance(content_sha512, str):
            raise TypeError("Expected argument 'content_sha512' to be a str")
        pulumi.set(__self__, "content_sha512", content_sha512)
        if filename and not isinstance(filename, str):
            raise TypeError("Expected argument 'filename' to be a str")
        pulumi.set(__self__, "filename", filename)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        Raw content of the file that was read, as UTF-8 encoded string. Files that do not contain UTF-8 text will have invalid UTF-8 sequences in `content`
        replaced with the Unicode replacement character.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> str:
        """
        Base64 encoded version of the file content (use this when dealing with binary data).
        """
        return pulumi.get(self, "content_base64")

    @property
    @pulumi.getter(name="contentBase64sha256")
    def content_base64sha256(self) -> str:
        """
        Base64 encoded SHA256 checksum of file content.
        """
        return pulumi.get(self, "content_base64sha256")

    @property
    @pulumi.getter(name="contentBase64sha512")
    def content_base64sha512(self) -> str:
        """
        Base64 encoded SHA512 checksum of file content.
        """
        return pulumi.get(self, "content_base64sha512")

    @property
    @pulumi.getter(name="contentMd5")
    def content_md5(self) -> str:
        """
        MD5 checksum of file content.
        """
        return pulumi.get(self, "content_md5")

    @property
    @pulumi.getter(name="contentSha1")
    def content_sha1(self) -> str:
        """
        SHA1 checksum of file content.
        """
        return pulumi.get(self, "content_sha1")

    @property
    @pulumi.getter(name="contentSha256")
    def content_sha256(self) -> str:
        """
        SHA256 checksum of file content.
        """
        return pulumi.get(self, "content_sha256")

    @property
    @pulumi.getter(name="contentSha512")
    def content_sha512(self) -> str:
        """
        SHA512 checksum of file content.
        """
        return pulumi.get(self, "content_sha512")

    @property
    @pulumi.getter
    def filename(self) -> str:
        """
        Path to the file that will be read. The data source will return an error if the file does not exist.
        """
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The hexadecimal encoding of the SHA1 checksum of the file content.
        """
        return pulumi.get(self, "id")


class AwaitableGetFileResult(GetFileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileResult(
            content=self.content,
            content_base64=self.content_base64,
            content_base64sha256=self.content_base64sha256,
            content_base64sha512=self.content_base64sha512,
            content_md5=self.content_md5,
            content_sha1=self.content_sha1,
            content_sha256=self.content_sha256,
            content_sha512=self.content_sha512,
            filename=self.filename,
            id=self.id)


def get_file(filename: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFileResult:
    """
    Reads a file from the local filesystem.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_aws as aws
    import pulumi_local as local

    foo = local.get_file(filename=f"{path['module']}/foo.bar")
    shared_zip = aws.s3.BucketObjectv2("sharedZip",
        bucket="my-bucket",
        key="my-key",
        content=foo.content)
    ```
    <!--End PulumiCodeChooser -->


    :param str filename: Path to the file that will be read. The data source will return an error if the file does not exist.
    """
    __args__ = dict()
    __args__['filename'] = filename
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('local:index/getFile:getFile', __args__, opts=opts, typ=GetFileResult).value

    return AwaitableGetFileResult(
        content=pulumi.get(__ret__, 'content'),
        content_base64=pulumi.get(__ret__, 'content_base64'),
        content_base64sha256=pulumi.get(__ret__, 'content_base64sha256'),
        content_base64sha512=pulumi.get(__ret__, 'content_base64sha512'),
        content_md5=pulumi.get(__ret__, 'content_md5'),
        content_sha1=pulumi.get(__ret__, 'content_sha1'),
        content_sha256=pulumi.get(__ret__, 'content_sha256'),
        content_sha512=pulumi.get(__ret__, 'content_sha512'),
        filename=pulumi.get(__ret__, 'filename'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_file)
def get_file_output(filename: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFileResult]:
    """
    Reads a file from the local filesystem.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_aws as aws
    import pulumi_local as local

    foo = local.get_file(filename=f"{path['module']}/foo.bar")
    shared_zip = aws.s3.BucketObjectv2("sharedZip",
        bucket="my-bucket",
        key="my-key",
        content=foo.content)
    ```
    <!--End PulumiCodeChooser -->


    :param str filename: Path to the file that will be read. The data source will return an error if the file does not exist.
    """
    ...
