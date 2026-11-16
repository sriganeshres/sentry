package golang

import (
	"go/ast"
	"go/types"
	"strconv"

	myTypes "sentry/ast/types"
)

func (p *CustomGoParser) handleOtherTypes(spec *ast.TypeSpec) myTypes.TypeDefinition {
	alias := spec.Assign.IsValid()

	return myTypes.TypeDefinition{
		Name:           spec.Name.Name,
		Kind:           "TYPE",
		Alias:          alias,
		UnderLyingType: p.getTypeString(spec.Type),
	}
}

func (p *CustomGoParser) getTypeString(expr ast.Expr) string {
	if p.typesInfo != nil {
		if t := p.typesInfo.TypeOf(expr); t != nil {
			return types.TypeString(t, func(pkg *types.Package) string {
				return pkg.Path()
			})
		}
	}

	return types.ExprString(expr)
}

func (p *CustomGoParser) handleTypeParams(params *ast.FieldList) [][]string {
	typeparams := [][]string{}

	if params == nil {
		return typeparams
	}

	for _, field := range params.List {
		constraint := types.ExprString(field.Type)

		for _, name := range field.Names {
			pos := p.fset.Position(name.Pos())

			typeparams = append(typeparams, []string{
				name.Name, constraint, strconv.Itoa(pos.Line), strconv.Itoa(pos.Column),
			})
		}
	}

	return typeparams
}
