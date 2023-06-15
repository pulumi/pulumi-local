package provider

import (
	"context"

	pftfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	upstream "github.com/terraform-providers/terraform-provider-local/pulumi-shim"
)

func ShimmedProvider() shim.Provider {
	return pftfbridge.ShimProvider(upstream.PFProvider())
}

func TfbridgeMain(pulumiSchema []byte, bridgeMetadata []byte) {
	meta := pftfbridge.ProviderMetadata{
		PackageSchema:  pulumiSchema,
		BridgeMetadata: bridgeMetadata,
	}
	pftfbridge.Main(context.Background(), "local", Provider(), meta)
}
