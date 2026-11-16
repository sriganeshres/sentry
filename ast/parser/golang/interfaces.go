package golang

import (
	"go/ast"
	"go/types"
	"strconv"

	myTypes "sentry/ast/types"
)

func (p *CustomGoParser) handleInterfaces(name string, inter *ast.InterfaceType) myTypes.TypeDefinition {
	interfaceType := myTypes.TypeDefinition{
		Name: name,
		Kind: "INTERFACE",
	}
	for _, vals := range inter.Methods.List {
		if len(vals.Names) == 0 {
			pos := p.fset.Position(vals.Pos())
			interfaceType.Embedded = append(interfaceType.Embedded, []string{
				p.getTypeString(vals.Type), strconv.Itoa(pos.Line), strconv.Itoa(pos.Column),
			})
		}
		for _, names := range vals.Names {
			pos := p.fset.Position(names.Pos())
			interfaceType.Methods = append(interfaceType.Methods, []string{names.Name, types.ExprString(vals.Type), strconv.Itoa(pos.Line), strconv.Itoa(pos.Column)})
		}
	}

	return interfaceType
}
