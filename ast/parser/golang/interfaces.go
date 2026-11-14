package golang

import (
	"fmt"
	"go/ast"
)

func (p *CustomGoParser) handleInterfaces(name string, inter *ast.InterfaceType) {
	post := p.fset.Position(inter.Pos())
	fmt.Printf("Interface %s: %t - At %s\n", name, inter.Incomplete, post)
	for _, vals := range inter.Methods.List {
		for _, names := range vals.Names {
			fmt.Printf("Methods %s \n", names.Name)
		}
	}
}
