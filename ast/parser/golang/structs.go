package golang

import (
	"fmt"
	"go/ast"
)

func (p *CustomGoParser) handleStructs(name string, st *ast.StructType) {
	fmt.Printf("Struct %s \n", name)
	for _, vals := range st.Fields.List {
		for _, names := range vals.Names {
			fmt.Printf("Fields %s - %s \n", names.Name, vals.Type)
		}
	}
}
