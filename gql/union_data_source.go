package gql

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

var (
	_ datasource.DataSource = &unionDataSource{}
)

func NewUnionDataSource() datasource.DataSource {
	return &unionDataSource{}
}

// unionDataSourceModel maps the data source schema data.
type unionDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Types  types.List   `tfsdk:"types"`
	Schema types.String `tfsdk:"schema"`
}

type unionDataSource struct{}

func (d *unionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_union"
}

func (d *unionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:   false,
				Required:   true,
				Validators: []validator.String{}, // TODO
			},
			"types": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    false,
				Required:    true,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.SizeAtLeast(1),
					// TODO validate type IDs too
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

func (d *unionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state unionDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var stringTypes []string
	state.Types.ElementsAs(ctx, &stringTypes, false)

	state.Schema = types.StringValue(fmt.Sprintf("union %s = %s", state.ID.ValueString(), strings.Join(stringTypes, " | ")))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
