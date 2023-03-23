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
	_ datasource.DataSource = &rootTypeDataSource{}
)

func NewRootTypeDataSource() datasource.DataSource {
	return &rootTypeDataSource{}
}

// rootTypeDataSourceModel maps the data source schema data.
type rootTypeDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Fields types.List   `tfsdk:"fields"`
	Schema types.String `tfsdk:"schema"`
}

type rootTypeDataSource struct{}

func (d *rootTypeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_root_type"
}

func (d *rootTypeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:   false,
				Required:   true,
				Validators: []validator.String{}, // TODO
			},
			"fields": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    false,
				Required:    true,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.SizeAtLeast(1),
					// TODO validate a bit
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

func (d *rootTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state rootTypeDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var builder strings.Builder

	builder.WriteString("type ")
	builder.WriteString(state.ID.ValueString())
	builder.WriteString(" {\n")

	for _, f := range state.Fields.Elements() {
		if f.IsNull() {
			continue
		}
		val := f.(types.String)
		builder.WriteRune('\t')
		builder.WriteString(val.ValueString())
		builder.WriteRune('\n')
	}

	builder.WriteString("}\n")

	state.Schema = types.StringValue(builder.String())

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
