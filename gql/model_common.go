package gql

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

type fieldModel struct {
	ID         types.String     `tfsdk:"id"`
	Type       types.String     `tfsdk:"type"`
	Arguments  []argumentModel  `tfsdk:"argument"`
	Directives []directiveModel `tfsdk:"directive"`
}

func (m fieldModel) String() string {
	var builder strings.Builder
	builder.WriteString(m.ID.ValueString())

	if m.Arguments != nil && len(m.Arguments) > 0 {
		builder.WriteRune('(')
		args := make([]string, 0, len(m.Arguments))
		for _, arg := range m.Arguments {
			args = append(args, arg.String())
		}
		builder.WriteString(strings.Join(args, ", "))
		builder.WriteRune(')')
	}

	builder.WriteString(": ")
	builder.WriteString(m.Type.ValueString())

	if m.Directives != nil && len(m.Directives) > 0 {
		for _, dir := range m.Directives {
			builder.WriteRune(' ')
			builder.WriteString(dir.String())
		}
	}

	return builder.String()
}

type argumentModel struct {
	ID         types.String     `tfsdk:"id"`
	Type       types.String     `tfsdk:"type"`
	Default    types.String     `tfsdk:"default"`
	Directives []directiveModel `tfsdk:"directive"`
}

func (m argumentModel) String() string {
	var builder strings.Builder
	builder.WriteString(m.ID.ValueString())
	builder.WriteString(": ")
	builder.WriteString(m.Type.ValueString())
	if !m.Default.IsNull() {
		builder.WriteString(" = ")
		builder.WriteString(m.Default.ValueString())
	}
	if m.Directives != nil && len(m.Directives) > 0 {
		for _, dir := range m.Directives {
			builder.WriteRune(' ')
			builder.WriteString(dir.String())
		}
	}
	return builder.String()
}

type directiveModel struct {
	ID        types.String             `tfsdk:"id"`
	Arguments []directiveArgumentModel `tfsdk:"argument"`
}

func (m directiveModel) String() string {
	if m.Arguments != nil && len(m.Arguments) > 0 {
		args := make([]string, 0, len(m.Arguments))
		for _, arg := range m.Arguments {
			args = append(args, arg.String())
		}
		return fmt.Sprintf("@%s: (%s)", m.ID.ValueString(), strings.Join(args, ", "))
	}
	return fmt.Sprintf("@%s", m.ID.ValueString())
}

type directiveArgumentModel struct {
	ID    types.String `tfsdk:"id"`
	Value types.String `tfsdk:"value"`
}

func (m directiveArgumentModel) String() string {
	return fmt.Sprintf("%s: \"%s\"", m.ID.ValueString(), m.Value.ValueString())
}
