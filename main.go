package main

import (
	"context"
	"terraform-provider-gqlschema/gql"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	providerserver.Serve(context.Background(), gql.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/mjdrgn/gqlschema",
	})
}
