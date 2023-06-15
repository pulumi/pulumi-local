// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package local

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSensitiveFile(ctx *pulumi.Context, args *LookupSensitiveFileArgs, opts ...pulumi.InvokeOption) (*LookupSensitiveFileResult, error) {
	var rv LookupSensitiveFileResult
	err := ctx.Invoke("local:index/getSensitiveFile:getSensitiveFile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSensitiveFile.
type LookupSensitiveFileArgs struct {
	// Path to the file that will be read. The data source will return an error if the file does not exist.
	Filename string `pulumi:"filename"`
}

// A collection of values returned by getSensitiveFile.
type LookupSensitiveFileResult struct {
	// Raw content of the file that was read, as UTF-8 encoded string. Files that do not contain UTF-8 text will have invalid UTF-8 sequences in `content`
	// replaced with the Unicode replacement character.
	Content string `pulumi:"content"`
	// Base64 encoded version of the file content (use this when dealing with binary data).
	ContentBase64 string `pulumi:"contentBase64"`
	// Base64 encoded SHA256 checksum of file content.
	ContentBase64sha256 string `pulumi:"contentBase64sha256"`
	// Base64 encoded SHA512 checksum of file content.
	ContentBase64sha512 string `pulumi:"contentBase64sha512"`
	// MD5 checksum of file content.
	ContentMd5 string `pulumi:"contentMd5"`
	// SHA1 checksum of file content.
	ContentSha1 string `pulumi:"contentSha1"`
	// SHA256 checksum of file content.
	ContentSha256 string `pulumi:"contentSha256"`
	// SHA512 checksum of file content.
	ContentSha512 string `pulumi:"contentSha512"`
	// Path to the file that will be read. The data source will return an error if the file does not exist.
	Filename string `pulumi:"filename"`
	// The hexadecimal encoding of the SHA1 checksum of the file content.
	Id string `pulumi:"id"`
}

func LookupSensitiveFileOutput(ctx *pulumi.Context, args LookupSensitiveFileOutputArgs, opts ...pulumi.InvokeOption) LookupSensitiveFileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSensitiveFileResult, error) {
			args := v.(LookupSensitiveFileArgs)
			r, err := LookupSensitiveFile(ctx, &args, opts...)
			var s LookupSensitiveFileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSensitiveFileResultOutput)
}

// A collection of arguments for invoking getSensitiveFile.
type LookupSensitiveFileOutputArgs struct {
	// Path to the file that will be read. The data source will return an error if the file does not exist.
	Filename pulumi.StringInput `pulumi:"filename"`
}

func (LookupSensitiveFileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSensitiveFileArgs)(nil)).Elem()
}

// A collection of values returned by getSensitiveFile.
type LookupSensitiveFileResultOutput struct{ *pulumi.OutputState }

func (LookupSensitiveFileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSensitiveFileResult)(nil)).Elem()
}

func (o LookupSensitiveFileResultOutput) ToLookupSensitiveFileResultOutput() LookupSensitiveFileResultOutput {
	return o
}

func (o LookupSensitiveFileResultOutput) ToLookupSensitiveFileResultOutputWithContext(ctx context.Context) LookupSensitiveFileResultOutput {
	return o
}

// Raw content of the file that was read, as UTF-8 encoded string. Files that do not contain UTF-8 text will have invalid UTF-8 sequences in `content`
// replaced with the Unicode replacement character.
func (o LookupSensitiveFileResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitiveFileResult) string { return v.Content }).(pulumi.StringOutput)
}

// Base64 encoded version of the file content (use this when dealing with binary data).
func (o LookupSensitiveFileResultOutput) ContentBase64() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitiveFileResult) string { return v.ContentBase64 }).(pulumi.StringOutput)
}

// Base64 encoded SHA256 checksum of file content.
func (o LookupSensitiveFileResultOutput) ContentBase64sha256() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitiveFileResult) string { return v.ContentBase64sha256 }).(pulumi.StringOutput)
}

// Base64 encoded SHA512 checksum of file content.
func (o LookupSensitiveFileResultOutput) ContentBase64sha512() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitiveFileResult) string { return v.ContentBase64sha512 }).(pulumi.StringOutput)
}

// MD5 checksum of file content.
func (o LookupSensitiveFileResultOutput) ContentMd5() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitiveFileResult) string { return v.ContentMd5 }).(pulumi.StringOutput)
}

// SHA1 checksum of file content.
func (o LookupSensitiveFileResultOutput) ContentSha1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitiveFileResult) string { return v.ContentSha1 }).(pulumi.StringOutput)
}

// SHA256 checksum of file content.
func (o LookupSensitiveFileResultOutput) ContentSha256() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitiveFileResult) string { return v.ContentSha256 }).(pulumi.StringOutput)
}

// SHA512 checksum of file content.
func (o LookupSensitiveFileResultOutput) ContentSha512() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitiveFileResult) string { return v.ContentSha512 }).(pulumi.StringOutput)
}

// Path to the file that will be read. The data source will return an error if the file does not exist.
func (o LookupSensitiveFileResultOutput) Filename() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitiveFileResult) string { return v.Filename }).(pulumi.StringOutput)
}

// The hexadecimal encoding of the SHA1 checksum of the file content.
func (o LookupSensitiveFileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitiveFileResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSensitiveFileResultOutput{})
}
