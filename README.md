[![Actions Status](https://github.com/pulumi/pulumi-local/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-local/actions)
[![NPM version](https://img.shields.io/npm/v/@pulumi/local)](https://www.npmjs.com/package/@pulumi/local)
[![NuGet version](https://img.shields.io/nuget/v/Pulumi.Local)](https://www.nuget.org/packages/Pulumi.Local)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-local/sdk/go)](https://pkg.go.dev/github.com/pulumi/pulumi-local/sdk/go)
[![License](https://img.shields.io/github/license/pulumi/pulumi-local)](https://github.com/pulumi/pulumi-local/blob/master/LICENSE)

# Local Resource Provider

> [!NOTE]  
> This provider exists to support users converting from Terraform to Pulumi.
> It is not recommended for new users to use this provider.
> Users looking to execute commands in Pulumi should use the [pulumi-command](https://www.pulumi.com/registry/packages/command/) provider.
> It offers a more flexible and powerful way to execute commands in Pulumi.

The Local resource provider for Pulumi lets you use Local resources in your cloud programs.
To use this package, please [install the Pulumi CLI first](https://www.pulumi.com/docs/install/).

>[!NOTE] As of v0.1.6, this provider is DEPRECATED and will no longer be maintained by Pulumi.
> We recommend using the [Local Provider](https://www.pulumi.com/blog/any-terraform-provider/) version of this package,
> which can be generated from the Local Terraform provider as follows:
> `pulumi package add terraform-provider registry.opentofu.org/hashicorp/local <version>`
> and follow the instructions.

## Migration

The currently equivalent upstream version to pulumi-local@v0.1.6 is search.opentofu.org/provider/hashicorp/local v2.5.2.
We recommend that you migrate between these versions of each provider.
To perform the migration of existing Local resources, we recommend running `pulumi import` in a fresh stack
which uses the locally built provider package.

## Reference

For further information, please visit [Local reference documentation](https://example.com/local).

-->
