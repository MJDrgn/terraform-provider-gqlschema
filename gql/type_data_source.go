package gql

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

var (
	_ datasource.DataSource = &typeDataSource{}
)

func NewTypeDataSource() datasource.DataSource {
	return &typeDataSource{}
}

// typeDataSourceModel maps the data source schema data.
type typeDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Mode   types.String `tfsdk:"mode"`
	Fields []fieldModel `tfsdk:"field"`
	Schema types.String `tfsdk:"schema"`
}

type typeDataSource struct{}

func (d *typeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_type"
}

func (d *typeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:   false,
				Required:   true,
				Validators: []validator.String{}, // TODO
			},
			"mode": schema.StringAttribute{
				Computed: false,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("output", "input", "interface"),
				},
			},
			"schema": schema.StringAttribute{
				Computed: true,
				Required: false,
				Optional: false,
			},
		},
		Blocks: map[string]schema.Block{
			"field": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:   false,
							Required:   true,
							Validators: []validator.String{}, // TODO
						},
						"type": schema.StringAttribute{
							Computed:   false,
							Required:   true,
							Validators: []validator.String{}, // TODO
						},
					},
					Blocks: map[string]schema.Block{
						"argument": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed:   false,
										Required:   true,
										Validators: []validator.String{}, // TODO
									},
									"type": schema.StringAttribute{
										Computed:   false,
										Required:   true,
										Validators: []validator.String{}, // TODO
									},
									"default": schema.StringAttribute{
										Computed:   false,
										Required:   false,
										Optional:   true,
										Validators: []validator.String{}, // TODO
									},
								},
								Blocks: map[string]schema.Block{
									"directive": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed:   false,
													Required:   true,
													Validators: []validator.String{}, // TODO
												},
											},
											Blocks: map[string]schema.Block{
												"argument": schema.ListNestedBlock{
													NestedObject: schema.NestedBlockObject{
														Attributes: map[string]schema.Attribute{
															"id": schema.StringAttribute{
																Computed:   false,
																Required:   true,
																Validators: []validator.String{}, // TODO
															},
															"value": schema.StringAttribute{
																Computed:   false,
																Required:   true,
																Validators: []validator.String{}, // TODO
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"directive": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed:   false,
										Required:   true,
										Validators: []validator.String{}, // TODO
									},
								},
								Blocks: map[string]schema.Block{
									"argument": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed:   false,
													Required:   true,
													Validators: []validator.String{}, // TODO
												},
												"value": schema.StringAttribute{
													Computed:   false,
													Required:   true,
													Validators: []validator.String{}, // TODO
												},
											},
										},
									},
								},
							},
						},
					},
					Validators: []validator.Object{},
				},
			},
		},
	}
}

func (d *typeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state typeDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var builder strings.Builder

	if state.Mode.IsNull() || state.Mode.ValueString() == "output" {
		builder.WriteString("type ")
	} else if state.Mode.ValueString() == "input" {
		builder.WriteString("input ")
	} else if state.Mode.ValueString() == "interface" {
		builder.WriteString("interface")
	} else {
		panic("unexpected Mode: " + state.Mode.ValueString())
	}

	builder.WriteString(state.ID.ValueString())
	builder.WriteString(" {\n")

	for _, f := range state.Fields {
		builder.WriteRune('\t')
		builder.WriteString(f.String())
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
