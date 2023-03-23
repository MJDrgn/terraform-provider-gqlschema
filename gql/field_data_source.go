package gql

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource = &fieldDataSource{}
)

func NewFieldDataSource() datasource.DataSource {
	return &fieldDataSource{}
}

// fieldDataSourceModel maps the data source schema data.
type fieldDataSourceModel struct {
	ID         types.String     `tfsdk:"id"`
	Type       types.String     `tfsdk:"type"`
	Arguments  []argumentModel  `tfsdk:"argument"`
	Directives []directiveModel `tfsdk:"directive"`
	Schema     types.String     `tfsdk:"schema"`
}

func (m fieldDataSourceModel) String() string {
	fm := fieldModel{
		ID:         m.ID,
		Type:       m.Type,
		Arguments:  m.Arguments,
		Directives: m.Directives,
	}
	return fm.String()
}

type fieldDataSource struct{}

func (d *fieldDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_field"
}

func (d *fieldDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
			"schema": schema.StringAttribute{
				Computed: true,
				Required: false,
				Optional: false,
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
	}
}

func (d *fieldDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state fieldDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Schema = types.StringValue(state.String())

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
