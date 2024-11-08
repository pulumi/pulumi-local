package provider

import (
	"context"

	upstream "github.com/terraform-providers/terraform-provider-local/pulumi-shim"

	pftfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
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
