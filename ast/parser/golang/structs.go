package golang

import (
	"fmt"
	"go/ast"
)

func (p *CustomGoParser) handleStructs(name string, st *ast.StructType) {
	pos := p.fset.Position(st.Pos())
	fmt.Printf("Struct %s - At %s\n", name, pos)
	for _, vals := range st.Fields.List {
		for _, names := range vals.Names {
			fmt.Printf("Fields %s - %s \n", names.Name, vals.Type)
		}
	}
}
