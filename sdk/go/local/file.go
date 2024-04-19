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
type File struct {
	pulumi.CustomResourceState

	// Content to store in the file, expected to be a UTF-8 encoded string.
	// Conflicts with `sensitiveContent`, `contentBase64` and `source`.
	// Exactly one of these four arguments must be specified.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// Content to store in the file, expected to be binary encoded as base64 string.
	// Conflicts with `content`, `sensitiveContent` and `source`.
	// Exactly one of these four arguments must be specified.
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
	// Default value is `"0777"`.
	DirectoryPermission pulumi.StringOutput `pulumi:"directoryPermission"`
	// Permissions to set for the output file (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0777"`.
	FilePermission pulumi.StringOutput `pulumi:"filePermission"`
	// The path to the file that will be created.
	// Missing parent directories will be created.
	// If the file already exists, it will be overridden with the given content.
	Filename pulumi.StringOutput `pulumi:"filename"`
	// Sensitive content to store in the file, expected to be an UTF-8 encoded string.
	// Will not be displayed in diffs.
	// Conflicts with `content`, `contentBase64` and `source`.
	// Exactly one of these four arguments must be specified.
	// If in need to use *sensitive* content, please use the `SensitiveFile`
	// resource instead.
	//
	// Deprecated: Use the `SensitiveFile` resource instead
	SensitiveContent pulumi.StringPtrOutput `pulumi:"sensitiveContent"`
	// Path to file to use as source for the one we are creating.
	// Conflicts with `content`, `sensitiveContent` and `contentBase64`.
	// Exactly one of these four arguments must be specified.
	Source pulumi.StringPtrOutput `pulumi:"source"`
}

// NewFile registers a new resource with the given unique name, arguments, and options.
func NewFile(ctx *pulumi.Context,
	name string, args *FileArgs, opts ...pulumi.ResourceOption) (*File, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Filename == nil {
		return nil, errors.New("invalid value for required argument 'Filename'")
	}
	if args.SensitiveContent != nil {
		args.SensitiveContent = pulumi.ToSecret(args.SensitiveContent).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"sensitiveContent",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource File
	err := ctx.RegisterResource("local:index/file:File", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFile gets an existing File resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileState, opts ...pulumi.ResourceOption) (*File, error) {
	var resource File
	err := ctx.ReadResource("local:index/file:File", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering File resources.
type fileState struct {
	// Content to store in the file, expected to be a UTF-8 encoded string.
	// Conflicts with `sensitiveContent`, `contentBase64` and `source`.
	// Exactly one of these four arguments must be specified.
	Content *string `pulumi:"content"`
	// Content to store in the file, expected to be binary encoded as base64 string.
	// Conflicts with `content`, `sensitiveContent` and `source`.
	// Exactly one of these four arguments must be specified.
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
	// Default value is `"0777"`.
	DirectoryPermission *string `pulumi:"directoryPermission"`
	// Permissions to set for the output file (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0777"`.
	FilePermission *string `pulumi:"filePermission"`
	// The path to the file that will be created.
	// Missing parent directories will be created.
	// If the file already exists, it will be overridden with the given content.
	Filename *string `pulumi:"filename"`
	// Sensitive content to store in the file, expected to be an UTF-8 encoded string.
	// Will not be displayed in diffs.
	// Conflicts with `content`, `contentBase64` and `source`.
	// Exactly one of these four arguments must be specified.
	// If in need to use *sensitive* content, please use the `SensitiveFile`
	// resource instead.
	//
	// Deprecated: Use the `SensitiveFile` resource instead
	SensitiveContent *string `pulumi:"sensitiveContent"`
	// Path to file to use as source for the one we are creating.
	// Conflicts with `content`, `sensitiveContent` and `contentBase64`.
	// Exactly one of these four arguments must be specified.
	Source *string `pulumi:"source"`
}

type FileState struct {
	// Content to store in the file, expected to be a UTF-8 encoded string.
	// Conflicts with `sensitiveContent`, `contentBase64` and `source`.
	// Exactly one of these four arguments must be specified.
	Content pulumi.StringPtrInput
	// Content to store in the file, expected to be binary encoded as base64 string.
	// Conflicts with `content`, `sensitiveContent` and `source`.
	// Exactly one of these four arguments must be specified.
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
	// Default value is `"0777"`.
	DirectoryPermission pulumi.StringPtrInput
	// Permissions to set for the output file (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0777"`.
	FilePermission pulumi.StringPtrInput
	// The path to the file that will be created.
	// Missing parent directories will be created.
	// If the file already exists, it will be overridden with the given content.
	Filename pulumi.StringPtrInput
	// Sensitive content to store in the file, expected to be an UTF-8 encoded string.
	// Will not be displayed in diffs.
	// Conflicts with `content`, `contentBase64` and `source`.
	// Exactly one of these four arguments must be specified.
	// If in need to use *sensitive* content, please use the `SensitiveFile`
	// resource instead.
	//
	// Deprecated: Use the `SensitiveFile` resource instead
	SensitiveContent pulumi.StringPtrInput
	// Path to file to use as source for the one we are creating.
	// Conflicts with `content`, `sensitiveContent` and `contentBase64`.
	// Exactly one of these four arguments must be specified.
	Source pulumi.StringPtrInput
}

func (FileState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileState)(nil)).Elem()
}

type fileArgs struct {
	// Content to store in the file, expected to be a UTF-8 encoded string.
	// Conflicts with `sensitiveContent`, `contentBase64` and `source`.
	// Exactly one of these four arguments must be specified.
	Content *string `pulumi:"content"`
	// Content to store in the file, expected to be binary encoded as base64 string.
	// Conflicts with `content`, `sensitiveContent` and `source`.
	// Exactly one of these four arguments must be specified.
	ContentBase64 *string `pulumi:"contentBase64"`
	// Permissions to set for directories created (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0777"`.
	DirectoryPermission *string `pulumi:"directoryPermission"`
	// Permissions to set for the output file (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0777"`.
	FilePermission *string `pulumi:"filePermission"`
	// The path to the file that will be created.
	// Missing parent directories will be created.
	// If the file already exists, it will be overridden with the given content.
	Filename string `pulumi:"filename"`
	// Sensitive content to store in the file, expected to be an UTF-8 encoded string.
	// Will not be displayed in diffs.
	// Conflicts with `content`, `contentBase64` and `source`.
	// Exactly one of these four arguments must be specified.
	// If in need to use *sensitive* content, please use the `SensitiveFile`
	// resource instead.
	//
	// Deprecated: Use the `SensitiveFile` resource instead
	SensitiveContent *string `pulumi:"sensitiveContent"`
	// Path to file to use as source for the one we are creating.
	// Conflicts with `content`, `sensitiveContent` and `contentBase64`.
	// Exactly one of these four arguments must be specified.
	Source *string `pulumi:"source"`
}

// The set of arguments for constructing a File resource.
type FileArgs struct {
	// Content to store in the file, expected to be a UTF-8 encoded string.
	// Conflicts with `sensitiveContent`, `contentBase64` and `source`.
	// Exactly one of these four arguments must be specified.
	Content pulumi.StringPtrInput
	// Content to store in the file, expected to be binary encoded as base64 string.
	// Conflicts with `content`, `sensitiveContent` and `source`.
	// Exactly one of these four arguments must be specified.
	ContentBase64 pulumi.StringPtrInput
	// Permissions to set for directories created (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0777"`.
	DirectoryPermission pulumi.StringPtrInput
	// Permissions to set for the output file (before umask), expressed as string in
	// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
	// Default value is `"0777"`.
	FilePermission pulumi.StringPtrInput
	// The path to the file that will be created.
	// Missing parent directories will be created.
	// If the file already exists, it will be overridden with the given content.
	Filename pulumi.StringInput
	// Sensitive content to store in the file, expected to be an UTF-8 encoded string.
	// Will not be displayed in diffs.
	// Conflicts with `content`, `contentBase64` and `source`.
	// Exactly one of these four arguments must be specified.
	// If in need to use *sensitive* content, please use the `SensitiveFile`
	// resource instead.
	//
	// Deprecated: Use the `SensitiveFile` resource instead
	SensitiveContent pulumi.StringPtrInput
	// Path to file to use as source for the one we are creating.
	// Conflicts with `content`, `sensitiveContent` and `contentBase64`.
	// Exactly one of these four arguments must be specified.
	Source pulumi.StringPtrInput
}

func (FileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileArgs)(nil)).Elem()
}

type FileInput interface {
	pulumi.Input

	ToFileOutput() FileOutput
	ToFileOutputWithContext(ctx context.Context) FileOutput
}

func (*File) ElementType() reflect.Type {
	return reflect.TypeOf((**File)(nil)).Elem()
}

func (i *File) ToFileOutput() FileOutput {
	return i.ToFileOutputWithContext(context.Background())
}

func (i *File) ToFileOutputWithContext(ctx context.Context) FileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileOutput)
}

// FileArrayInput is an input type that accepts FileArray and FileArrayOutput values.
// You can construct a concrete instance of `FileArrayInput` via:
//
//	FileArray{ FileArgs{...} }
type FileArrayInput interface {
	pulumi.Input

	ToFileArrayOutput() FileArrayOutput
	ToFileArrayOutputWithContext(context.Context) FileArrayOutput
}

type FileArray []FileInput

func (FileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*File)(nil)).Elem()
}

func (i FileArray) ToFileArrayOutput() FileArrayOutput {
	return i.ToFileArrayOutputWithContext(context.Background())
}

func (i FileArray) ToFileArrayOutputWithContext(ctx context.Context) FileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileArrayOutput)
}

// FileMapInput is an input type that accepts FileMap and FileMapOutput values.
// You can construct a concrete instance of `FileMapInput` via:
//
//	FileMap{ "key": FileArgs{...} }
type FileMapInput interface {
	pulumi.Input

	ToFileMapOutput() FileMapOutput
	ToFileMapOutputWithContext(context.Context) FileMapOutput
}

type FileMap map[string]FileInput

func (FileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*File)(nil)).Elem()
}

func (i FileMap) ToFileMapOutput() FileMapOutput {
	return i.ToFileMapOutputWithContext(context.Background())
}

func (i FileMap) ToFileMapOutputWithContext(ctx context.Context) FileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileMapOutput)
}

type FileOutput struct{ *pulumi.OutputState }

func (FileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**File)(nil)).Elem()
}

func (o FileOutput) ToFileOutput() FileOutput {
	return o
}

func (o FileOutput) ToFileOutputWithContext(ctx context.Context) FileOutput {
	return o
}

// Content to store in the file, expected to be a UTF-8 encoded string.
// Conflicts with `sensitiveContent`, `contentBase64` and `source`.
// Exactly one of these four arguments must be specified.
func (o FileOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *File) pulumi.StringPtrOutput { return v.Content }).(pulumi.StringPtrOutput)
}

// Content to store in the file, expected to be binary encoded as base64 string.
// Conflicts with `content`, `sensitiveContent` and `source`.
// Exactly one of these four arguments must be specified.
func (o FileOutput) ContentBase64() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *File) pulumi.StringPtrOutput { return v.ContentBase64 }).(pulumi.StringPtrOutput)
}

// Base64 encoded SHA256 checksum of file content.
func (o FileOutput) ContentBase64sha256() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.ContentBase64sha256 }).(pulumi.StringOutput)
}

// Base64 encoded SHA512 checksum of file content.
func (o FileOutput) ContentBase64sha512() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.ContentBase64sha512 }).(pulumi.StringOutput)
}

// MD5 checksum of file content.
func (o FileOutput) ContentMd5() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.ContentMd5 }).(pulumi.StringOutput)
}

// SHA1 checksum of file content.
func (o FileOutput) ContentSha1() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.ContentSha1 }).(pulumi.StringOutput)
}

// SHA256 checksum of file content.
func (o FileOutput) ContentSha256() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.ContentSha256 }).(pulumi.StringOutput)
}

// SHA512 checksum of file content.
func (o FileOutput) ContentSha512() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.ContentSha512 }).(pulumi.StringOutput)
}

// Permissions to set for directories created (before umask), expressed as string in
// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
// Default value is `"0777"`.
func (o FileOutput) DirectoryPermission() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.DirectoryPermission }).(pulumi.StringOutput)
}

// Permissions to set for the output file (before umask), expressed as string in
// [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).
// Default value is `"0777"`.
func (o FileOutput) FilePermission() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.FilePermission }).(pulumi.StringOutput)
}

// The path to the file that will be created.
// Missing parent directories will be created.
// If the file already exists, it will be overridden with the given content.
func (o FileOutput) Filename() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.Filename }).(pulumi.StringOutput)
}

// Sensitive content to store in the file, expected to be an UTF-8 encoded string.
// Will not be displayed in diffs.
// Conflicts with `content`, `contentBase64` and `source`.
// Exactly one of these four arguments must be specified.
// If in need to use *sensitive* content, please use the `SensitiveFile`
// resource instead.
//
// Deprecated: Use the `SensitiveFile` resource instead
func (o FileOutput) SensitiveContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *File) pulumi.StringPtrOutput { return v.SensitiveContent }).(pulumi.StringPtrOutput)
}

// Path to file to use as source for the one we are creating.
// Conflicts with `content`, `sensitiveContent` and `contentBase64`.
// Exactly one of these four arguments must be specified.
func (o FileOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *File) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

type FileArrayOutput struct{ *pulumi.OutputState }

func (FileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*File)(nil)).Elem()
}

func (o FileArrayOutput) ToFileArrayOutput() FileArrayOutput {
	return o
}

func (o FileArrayOutput) ToFileArrayOutputWithContext(ctx context.Context) FileArrayOutput {
	return o
}

func (o FileArrayOutput) Index(i pulumi.IntInput) FileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *File {
		return vs[0].([]*File)[vs[1].(int)]
	}).(FileOutput)
}

type FileMapOutput struct{ *pulumi.OutputState }

func (FileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*File)(nil)).Elem()
}

func (o FileMapOutput) ToFileMapOutput() FileMapOutput {
	return o
}

func (o FileMapOutput) ToFileMapOutputWithContext(ctx context.Context) FileMapOutput {
	return o
}

func (o FileMapOutput) MapIndex(k pulumi.StringInput) FileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *File {
		return vs[0].(map[string]*File)[vs[1].(string)]
	}).(FileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FileInput)(nil)).Elem(), &File{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileArrayInput)(nil)).Elem(), FileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileMapInput)(nil)).Elem(), FileMap{})
	pulumi.RegisterOutputType(FileOutput{})
	pulumi.RegisterOutputType(FileArrayOutput{})
	pulumi.RegisterOutputType(FileMapOutput{})
}
