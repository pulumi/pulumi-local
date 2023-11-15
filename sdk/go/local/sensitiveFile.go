// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package local

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-local/sdk/go/local/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-local/sdk/go/local"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := local.NewSensitiveFile(ctx, "foo", &local.SensitiveFileArgs{
//				Content:  pulumi.String("foo!"),
//				Filename: pulumi.String(fmt.Sprintf("%v/foo.bar", path.Module)),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SensitiveFile struct {
	pulumi.CustomResourceState

	// Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
	// Conflicts with `contentBase64` and `source`.
	// Exactly one of these three arguments must be specified.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// Sensitive Content to store in the file, expected to be binary encoded as base64 string.
	// Conflicts with `content` and `source`.
	// Exactly one of these three arguments must be specified.
	ContentBase64 pulumi.StringPtrOutput `pulumi:"contentBase64"`
	// Base64 encoded SHA256 checksum of file content.
	ContentBase64sha256 pulumi.StringOutput `pulumi:"contentBase64sha256"`
	// Base64 encoded SHA512 checksum of file content.
	ContentBase64sha512 pulumi.StringOutput `pulumi:"contentBase64sha512"`
	// MD5 checksum of file content.
	ContentMd5 pulumi.StringOutput `pulumi:"contentMd5"`
	// SHA1 checksum of file content.
	ContentSha1 pulumi.StringOutput `pulumi:"contentSha1"`
	// SHA256 checksum of file content.
	ContentSha256 pulumi.StringOutput `pulumi:"contentSha256"`
	// SHA512 checksum of file content.
	ContentSha512 pulumi.StringOutput `pulumi:"contentSha512"`
	// Permissions to set for directories created (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0700"`.
	DirectoryPermission pulumi.StringOutput `pulumi:"directoryPermission"`
	// Permissions to set for the output file (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0700"`.
	FilePermission pulumi.StringOutput `pulumi:"filePermission"`
	// The path to the file that will be created.
	// Missing parent directories will be created.
	// If the file already exists, it will be overridden with the given content.
	Filename pulumi.StringOutput `pulumi:"filename"`
	// Path to file to use as source for the one we are creating.
	// Conflicts with `content` and `contentBase64`.
	// Exactly one of these three arguments must be specified.
	Source pulumi.StringPtrOutput `pulumi:"source"`
}

// NewSensitiveFile registers a new resource with the given unique name, arguments, and options.
func NewSensitiveFile(ctx *pulumi.Context,
	name string, args *SensitiveFileArgs, opts ...pulumi.ResourceOption) (*SensitiveFile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Filename == nil {
		return nil, errors.New("invalid value for required argument 'Filename'")
	}
	if args.Content != nil {
		args.Content = pulumi.ToSecret(args.Content).(pulumi.StringPtrInput)
	}
	if args.ContentBase64 != nil {
		args.ContentBase64 = pulumi.ToSecret(args.ContentBase64).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"content",
		"contentBase64",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SensitiveFile
	err := ctx.RegisterResource("local:index/sensitiveFile:SensitiveFile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSensitiveFile gets an existing SensitiveFile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSensitiveFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SensitiveFileState, opts ...pulumi.ResourceOption) (*SensitiveFile, error) {
	var resource SensitiveFile
	err := ctx.ReadResource("local:index/sensitiveFile:SensitiveFile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SensitiveFile resources.
type sensitiveFileState struct {
	// Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
	// Conflicts with `contentBase64` and `source`.
	// Exactly one of these three arguments must be specified.
	Content *string `pulumi:"content"`
	// Sensitive Content to store in the file, expected to be binary encoded as base64 string.
	// Conflicts with `content` and `source`.
	// Exactly one of these three arguments must be specified.
	ContentBase64 *string `pulumi:"contentBase64"`
	// Base64 encoded SHA256 checksum of file content.
	ContentBase64sha256 *string `pulumi:"contentBase64sha256"`
	// Base64 encoded SHA512 checksum of file content.
	ContentBase64sha512 *string `pulumi:"contentBase64sha512"`
	// MD5 checksum of file content.
	ContentMd5 *string `pulumi:"contentMd5"`
	// SHA1 checksum of file content.
	ContentSha1 *string `pulumi:"contentSha1"`
	// SHA256 checksum of file content.
	ContentSha256 *string `pulumi:"contentSha256"`
	// SHA512 checksum of file content.
	ContentSha512 *string `pulumi:"contentSha512"`
	// Permissions to set for directories created (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0700"`.
	DirectoryPermission *string `pulumi:"directoryPermission"`
	// Permissions to set for the output file (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0700"`.
	FilePermission *string `pulumi:"filePermission"`
	// The path to the file that will be created.
	// Missing parent directories will be created.
	// If the file already exists, it will be overridden with the given content.
	Filename *string `pulumi:"filename"`
	// Path to file to use as source for the one we are creating.
	// Conflicts with `content` and `contentBase64`.
	// Exactly one of these three arguments must be specified.
	Source *string `pulumi:"source"`
}

type SensitiveFileState struct {
	// Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
	// Conflicts with `contentBase64` and `source`.
	// Exactly one of these three arguments must be specified.
	Content pulumi.StringPtrInput
	// Sensitive Content to store in the file, expected to be binary encoded as base64 string.
	// Conflicts with `content` and `source`.
	// Exactly one of these three arguments must be specified.
	ContentBase64 pulumi.StringPtrInput
	// Base64 encoded SHA256 checksum of file content.
	ContentBase64sha256 pulumi.StringPtrInput
	// Base64 encoded SHA512 checksum of file content.
	ContentBase64sha512 pulumi.StringPtrInput
	// MD5 checksum of file content.
	ContentMd5 pulumi.StringPtrInput
	// SHA1 checksum of file content.
	ContentSha1 pulumi.StringPtrInput
	// SHA256 checksum of file content.
	ContentSha256 pulumi.StringPtrInput
	// SHA512 checksum of file content.
	ContentSha512 pulumi.StringPtrInput
	// Permissions to set for directories created (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0700"`.
	DirectoryPermission pulumi.StringPtrInput
	// Permissions to set for the output file (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0700"`.
	FilePermission pulumi.StringPtrInput
	// The path to the file that will be created.
	// Missing parent directories will be created.
	// If the file already exists, it will be overridden with the given content.
	Filename pulumi.StringPtrInput
	// Path to file to use as source for the one we are creating.
	// Conflicts with `content` and `contentBase64`.
	// Exactly one of these three arguments must be specified.
	Source pulumi.StringPtrInput
}

func (SensitiveFileState) ElementType() reflect.Type {
	return reflect.TypeOf((*sensitiveFileState)(nil)).Elem()
}

type sensitiveFileArgs struct {
	// Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
	// Conflicts with `contentBase64` and `source`.
	// Exactly one of these three arguments must be specified.
	Content *string `pulumi:"content"`
	// Sensitive Content to store in the file, expected to be binary encoded as base64 string.
	// Conflicts with `content` and `source`.
	// Exactly one of these three arguments must be specified.
	ContentBase64 *string `pulumi:"contentBase64"`
	// Permissions to set for directories created (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0700"`.
	DirectoryPermission *string `pulumi:"directoryPermission"`
	// Permissions to set for the output file (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0700"`.
	FilePermission *string `pulumi:"filePermission"`
	// The path to the file that will be created.
	// Missing parent directories will be created.
	// If the file already exists, it will be overridden with the given content.
	Filename string `pulumi:"filename"`
	// Path to file to use as source for the one we are creating.
	// Conflicts with `content` and `contentBase64`.
	// Exactly one of these three arguments must be specified.
	Source *string `pulumi:"source"`
}

// The set of arguments for constructing a SensitiveFile resource.
type SensitiveFileArgs struct {
	// Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
	// Conflicts with `contentBase64` and `source`.
	// Exactly one of these three arguments must be specified.
	Content pulumi.StringPtrInput
	// Sensitive Content to store in the file, expected to be binary encoded as base64 string.
	// Conflicts with `content` and `source`.
	// Exactly one of these three arguments must be specified.
	ContentBase64 pulumi.StringPtrInput
	// Permissions to set for directories created (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0700"`.
	DirectoryPermission pulumi.StringPtrInput
	// Permissions to set for the output file (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0700"`.
	FilePermission pulumi.StringPtrInput
	// The path to the file that will be created.
	// Missing parent directories will be created.
	// If the file already exists, it will be overridden with the given content.
	Filename pulumi.StringInput
	// Path to file to use as source for the one we are creating.
	// Conflicts with `content` and `contentBase64`.
	// Exactly one of these three arguments must be specified.
	Source pulumi.StringPtrInput
}

func (SensitiveFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sensitiveFileArgs)(nil)).Elem()
}

type SensitiveFileInput interface {
	pulumi.Input

	ToSensitiveFileOutput() SensitiveFileOutput
	ToSensitiveFileOutputWithContext(ctx context.Context) SensitiveFileOutput
}

func (*SensitiveFile) ElementType() reflect.Type {
	return reflect.TypeOf((**SensitiveFile)(nil)).Elem()
}

func (i *SensitiveFile) ToSensitiveFileOutput() SensitiveFileOutput {
	return i.ToSensitiveFileOutputWithContext(context.Background())
}

func (i *SensitiveFile) ToSensitiveFileOutputWithContext(ctx context.Context) SensitiveFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensitiveFileOutput)
}

// SensitiveFileArrayInput is an input type that accepts SensitiveFileArray and SensitiveFileArrayOutput values.
// You can construct a concrete instance of `SensitiveFileArrayInput` via:
//
//	SensitiveFileArray{ SensitiveFileArgs{...} }
type SensitiveFileArrayInput interface {
	pulumi.Input

	ToSensitiveFileArrayOutput() SensitiveFileArrayOutput
	ToSensitiveFileArrayOutputWithContext(context.Context) SensitiveFileArrayOutput
}

type SensitiveFileArray []SensitiveFileInput

func (SensitiveFileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SensitiveFile)(nil)).Elem()
}

func (i SensitiveFileArray) ToSensitiveFileArrayOutput() SensitiveFileArrayOutput {
	return i.ToSensitiveFileArrayOutputWithContext(context.Background())
}

func (i SensitiveFileArray) ToSensitiveFileArrayOutputWithContext(ctx context.Context) SensitiveFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensitiveFileArrayOutput)
}

// SensitiveFileMapInput is an input type that accepts SensitiveFileMap and SensitiveFileMapOutput values.
// You can construct a concrete instance of `SensitiveFileMapInput` via:
//
//	SensitiveFileMap{ "key": SensitiveFileArgs{...} }
type SensitiveFileMapInput interface {
	pulumi.Input

	ToSensitiveFileMapOutput() SensitiveFileMapOutput
	ToSensitiveFileMapOutputWithContext(context.Context) SensitiveFileMapOutput
}

type SensitiveFileMap map[string]SensitiveFileInput

func (SensitiveFileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SensitiveFile)(nil)).Elem()
}

func (i SensitiveFileMap) ToSensitiveFileMapOutput() SensitiveFileMapOutput {
	return i.ToSensitiveFileMapOutputWithContext(context.Background())
}

func (i SensitiveFileMap) ToSensitiveFileMapOutputWithContext(ctx context.Context) SensitiveFileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensitiveFileMapOutput)
}

type SensitiveFileOutput struct{ *pulumi.OutputState }

func (SensitiveFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SensitiveFile)(nil)).Elem()
}

func (o SensitiveFileOutput) ToSensitiveFileOutput() SensitiveFileOutput {
	return o
}

func (o SensitiveFileOutput) ToSensitiveFileOutputWithContext(ctx context.Context) SensitiveFileOutput {
	return o
}

// Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
// Conflicts with `contentBase64` and `source`.
// Exactly one of these three arguments must be specified.
func (o SensitiveFileOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitiveFile) pulumi.StringPtrOutput { return v.Content }).(pulumi.StringPtrOutput)
}

// Sensitive Content to store in the file, expected to be binary encoded as base64 string.
// Conflicts with `content` and `source`.
// Exactly one of these three arguments must be specified.
func (o SensitiveFileOutput) ContentBase64() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitiveFile) pulumi.StringPtrOutput { return v.ContentBase64 }).(pulumi.StringPtrOutput)
}

// Base64 encoded SHA256 checksum of file content.
func (o SensitiveFileOutput) ContentBase64sha256() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitiveFile) pulumi.StringOutput { return v.ContentBase64sha256 }).(pulumi.StringOutput)
}

// Base64 encoded SHA512 checksum of file content.
func (o SensitiveFileOutput) ContentBase64sha512() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitiveFile) pulumi.StringOutput { return v.ContentBase64sha512 }).(pulumi.StringOutput)
}

// MD5 checksum of file content.
func (o SensitiveFileOutput) ContentMd5() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitiveFile) pulumi.StringOutput { return v.ContentMd5 }).(pulumi.StringOutput)
}

// SHA1 checksum of file content.
func (o SensitiveFileOutput) ContentSha1() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitiveFile) pulumi.StringOutput { return v.ContentSha1 }).(pulumi.StringOutput)
}

// SHA256 checksum of file content.
func (o SensitiveFileOutput) ContentSha256() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitiveFile) pulumi.StringOutput { return v.ContentSha256 }).(pulumi.StringOutput)
}

// SHA512 checksum of file content.
func (o SensitiveFileOutput) ContentSha512() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitiveFile) pulumi.StringOutput { return v.ContentSha512 }).(pulumi.StringOutput)
}

// Permissions to set for directories created (before umask), expressed as string in
// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
// Default value is `"0700"`.
func (o SensitiveFileOutput) DirectoryPermission() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitiveFile) pulumi.StringOutput { return v.DirectoryPermission }).(pulumi.StringOutput)
}

// Permissions to set for the output file (before umask), expressed as string in
// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
// Default value is `"0700"`.
func (o SensitiveFileOutput) FilePermission() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitiveFile) pulumi.StringOutput { return v.FilePermission }).(pulumi.StringOutput)
}

// The path to the file that will be created.
// Missing parent directories will be created.
// If the file already exists, it will be overridden with the given content.
func (o SensitiveFileOutput) Filename() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitiveFile) pulumi.StringOutput { return v.Filename }).(pulumi.StringOutput)
}

// Path to file to use as source for the one we are creating.
// Conflicts with `content` and `contentBase64`.
// Exactly one of these three arguments must be specified.
func (o SensitiveFileOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitiveFile) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

type SensitiveFileArrayOutput struct{ *pulumi.OutputState }

func (SensitiveFileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SensitiveFile)(nil)).Elem()
}

func (o SensitiveFileArrayOutput) ToSensitiveFileArrayOutput() SensitiveFileArrayOutput {
	return o
}

func (o SensitiveFileArrayOutput) ToSensitiveFileArrayOutputWithContext(ctx context.Context) SensitiveFileArrayOutput {
	return o
}

func (o SensitiveFileArrayOutput) Index(i pulumi.IntInput) SensitiveFileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SensitiveFile {
		return vs[0].([]*SensitiveFile)[vs[1].(int)]
	}).(SensitiveFileOutput)
}

type SensitiveFileMapOutput struct{ *pulumi.OutputState }

func (SensitiveFileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SensitiveFile)(nil)).Elem()
}

func (o SensitiveFileMapOutput) ToSensitiveFileMapOutput() SensitiveFileMapOutput {
	return o
}

func (o SensitiveFileMapOutput) ToSensitiveFileMapOutputWithContext(ctx context.Context) SensitiveFileMapOutput {
	return o
}

func (o SensitiveFileMapOutput) MapIndex(k pulumi.StringInput) SensitiveFileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SensitiveFile {
		return vs[0].(map[string]*SensitiveFile)[vs[1].(string)]
	}).(SensitiveFileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SensitiveFileInput)(nil)).Elem(), &SensitiveFile{})
	pulumi.RegisterInputType(reflect.TypeOf((*SensitiveFileArrayInput)(nil)).Elem(), SensitiveFileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SensitiveFileMapInput)(nil)).Elem(), SensitiveFileMap{})
	pulumi.RegisterOutputType(SensitiveFileOutput{})
	pulumi.RegisterOutputType(SensitiveFileArrayOutput{})
	pulumi.RegisterOutputType(SensitiveFileMapOutput{})
}
