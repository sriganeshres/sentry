package golang

import (
	"go/ast"
	"go/types"
	"strconv"

	myTypes "sentry/ast/types"
)

func (p *CustomGoParser) handleStructs(name string, st *ast.StructType) myTypes.TypeDefinition {
	structType := myTypes.TypeDefinition{
		Name: name,
		Kind: "STRUCT",
	}
	for _, vals := range st.Fields.List {
		if len(vals.Names) == 0 {
			pos := p.fset.Position(vals.Pos())
			structType.Embedded = append(structType.Embedded, []string{
				p.getTypeString(vals.Type), strconv.Itoa(pos.Line), strconv.Itoa(pos.Column),
			})
		}
		for _, names := range vals.Names {
			pos := p.fset.Position(names.Pos())
			structType.Fields = append(structType.Fields, []string{names.Name, types.ExprString(vals.Type), strconv.Itoa(pos.Line), strconv.Itoa(pos.Column)})
		}
	}

	return structType
}
