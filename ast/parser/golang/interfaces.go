package golang

import (
	"fmt"
	"go/ast"
)

func (p *CustomGoParser) handleInterfaces(name string, inter *ast.InterfaceType) {
	fmt.Printf("Interface %s: %t\n", name, inter.Incomplete)
	for _, vals := range inter.Methods.List {
		for _, names := range vals.Names {
			fmt.Printf("Methods %s \n", names.Name)
		}
	}
}
