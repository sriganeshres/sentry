package golang

import (
	"fmt"
	"go/ast"
	"go/types"
)

func (p *CustomGoParser) handleFuncDecls(fun *ast.FuncDecl) {
	start := p.fset.Position(fun.Pos())
	end := p.fset.Position(fun.End())
	name := p.fset.Position(fun.Name.Pos())
	if fun.Recv == nil {
		fmt.Printf("Function %s: name at %s | Block from %s to %s \n", fun.Name.Name, name, start, end)
	} else {
		fmt.Printf("Function %s - %s: name at %s | Block from %s to %s \n", fun.Name.Name, types.ExprString(fun.Recv.List[0].Type), name, start, end)
	}
}

func (p *CustomGoParser) handleFuncCalls(call *ast.CallExpr) {
	name := p.fset.Position(call.Pos())
	switch fn := call.Fun.(type) {
	case *ast.Ident:
		fmt.Printf("Local %s: at %s\n", fn.Name, name)
	case *ast.SelectorExpr:
		caller := types.ExprString(fn.X)
		fmt.Printf("Selector Call %s - %s: at %s\n", caller, fn.Sel.Name, name)
	default:
		fmt.Printf("Call %s: at %s\n", types.ExprString(call.Fun), name)
	}
}
