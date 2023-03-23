package gql

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"regexp"
	"strings"
)

var (
	_ datasource.DataSource = &enumDataSource{}
)

func NewEnumDataSource() datasource.DataSource {
	return &enumDataSource{}
}

// enumDataSourceModel maps the data source schema data.
type enumDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Values types.List   `tfsdk:"values"`
	Schema types.String `tfsdk:"schema"`
}

type enumDataSource struct{}

func (d *enumDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_enum"
}

func (d *enumDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:   false,
				Required:   true,
				Validators: []validator.String{}, // TODO
			},
			"values": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    false,
				Required:    true,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.SizeAtLeast(1),
					listvalidator.ValueStringsAre(stringvalidator.RegexMatches(regexp.MustCompile(`[A-Z]+`), "Enum values must be upper case")),
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

func (d *enumDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state enumDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var builder strings.Builder

	builder.WriteString("enum ")

	builder.WriteString(state.ID.ValueString())
	builder.WriteString(" {\n")

	for _, f := range state.Values.Elements() {
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
