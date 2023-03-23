package gql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &gqlProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &gqlProvider{}
}

// gqlProvider is the provider implementation.
type gqlProvider struct{}

// Metadata returns the provider type name.
func (p *gqlProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "gqlschema"
}

// Schema defines the provider-level schema for configuration data.
func (p *gqlProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Configure prepares a gql API client for data sources and resources.
func (p *gqlProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *gqlProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewTypeDataSource,
		NewFieldDataSource,
		NewEnumDataSource,
		NewUnionDataSource,
		NewScalarDataSource,
		NewRootTypeDataSource,
		NewSchemaDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *gqlProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
