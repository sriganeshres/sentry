package golang

import (
	"fmt"
	"go/ast"
)

func (p *CustomGoParser) handleFuncDecls(fun *ast.FuncDecl) {
	if fun.Recv == nil {
		fmt.Printf("Function %s \n", fun.Name.Name)
	} else {
		fmt.Printf("Function %s - %s \n", fun.Name.Name, fun.Recv.List[0].Type)
	}
}
