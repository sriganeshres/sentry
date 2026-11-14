package golang

import (
	"fmt"
	"go/ast"
)

func (p *CustomGoParser) handleInterfaces(name string, inter *ast.InterfaceType) {
	post := p.fset.Position(inter.Pos())
	fmt.Printf("Interface %s - At %s\n", name, post)
	for _, vals := range inter.Methods.List {
		for _, names := range vals.Names {
			fmt.Printf("Methods %s \n", names.Name)
		}
	}
}
