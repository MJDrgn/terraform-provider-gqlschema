package gql

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

var (
	_ datasource.DataSource = &schemaDataSource{}
)

func NewSchemaDataSource() datasource.DataSource {
	return &schemaDataSource{}
}

// schemaDataSourceModel maps the data source schema data.
type schemaDataSourceModel struct {
	ID         types.String `tfsdk:"id"`
	Components types.List   `tfsdk:"components"`
	Schema     types.String `tfsdk:"schema"`
}

type schemaDataSource struct{}

func (d *schemaDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_schema"
}

func (d *schemaDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:   false,
				Required:   true,
				Validators: []validator.String{}, // TODO
			},
			"components": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    false,
				Required:    true,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.SizeAtLeast(1),
				},
			},
			"schema": schema.StringAttribute{
				Computed: true,
				Required: false,
				Optional: false,
			},
		},
	}
}

func (d *schemaDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state schemaDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var builder strings.Builder

	for _, f := range state.Components.Elements() {
		if f.IsNull() {
			continue
		}
		val := f.(types.String)
		builder.WriteString(val.ValueString())
		builder.WriteRune('\n')
	}

	state.Schema = types.StringValue(builder.String())

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
