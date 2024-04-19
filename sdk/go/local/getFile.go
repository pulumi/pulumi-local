// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package local

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-local/sdk/go/local/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Reads a file from the local filesystem.
func LookupFile(ctx *pulumi.Context, args *LookupFileArgs, opts ...pulumi.InvokeOption) (*LookupFileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFileResult
	err := ctx.Invoke("local:index/getFile:getFile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFile.
type LookupFileArgs struct {
	// Path to the file that will be read. The data source will return an error if the file does not exist.
	Filename string `pulumi:"filename"`
}

// A collection of values returned by getFile.
type LookupFileResult struct {
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

func LookupFileOutput(ctx *pulumi.Context, args LookupFileOutputArgs, opts ...pulumi.InvokeOption) LookupFileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileResult, error) {
			args := v.(LookupFileArgs)
			r, err := LookupFile(ctx, &args, opts...)
			var s LookupFileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileResultOutput)
}

// A collection of arguments for invoking getFile.
type LookupFileOutputArgs struct {
	// Path to the file that will be read. The data source will return an error if the file does not exist.
	Filename pulumi.StringInput `pulumi:"filename"`
}

func (LookupFileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileArgs)(nil)).Elem()
}

// A collection of values returned by getFile.
type LookupFileResultOutput struct{ *pulumi.OutputState }

func (LookupFileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileResult)(nil)).Elem()
}

func (o LookupFileResultOutput) ToLookupFileResultOutput() LookupFileResultOutput {
	return o
}

func (o LookupFileResultOutput) ToLookupFileResultOutputWithContext(ctx context.Context) LookupFileResultOutput {
	return o
}

// Raw content of the file that was read, as UTF-8 encoded string. Files that do not contain UTF-8 text will have invalid UTF-8 sequences in `content`
// replaced with the Unicode replacement character.
func (o LookupFileResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.Content }).(pulumi.StringOutput)
}

// Base64 encoded version of the file content (use this when dealing with binary data).
func (o LookupFileResultOutput) ContentBase64() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.ContentBase64 }).(pulumi.StringOutput)
}

// Base64 encoded SHA256 checksum of file content.
func (o LookupFileResultOutput) ContentBase64sha256() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.ContentBase64sha256 }).(pulumi.StringOutput)
}

// Base64 encoded SHA512 checksum of file content.
func (o LookupFileResultOutput) ContentBase64sha512() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.ContentBase64sha512 }).(pulumi.StringOutput)
}

// MD5 checksum of file content.
func (o LookupFileResultOutput) ContentMd5() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.ContentMd5 }).(pulumi.StringOutput)
}

// SHA1 checksum of file content.
func (o LookupFileResultOutput) ContentSha1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.ContentSha1 }).(pulumi.StringOutput)
}

// SHA256 checksum of file content.
func (o LookupFileResultOutput) ContentSha256() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.ContentSha256 }).(pulumi.StringOutput)
}

// SHA512 checksum of file content.
func (o LookupFileResultOutput) ContentSha512() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.ContentSha512 }).(pulumi.StringOutput)
}

// Path to the file that will be read. The data source will return an error if the file does not exist.
func (o LookupFileResultOutput) Filename() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.Filename }).(pulumi.StringOutput)
}

// The hexadecimal encoding of the SHA1 checksum of the file content.
func (o LookupFileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileResultOutput{})
}
