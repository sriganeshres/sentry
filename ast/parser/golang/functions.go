package golang

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/types/typeutil"
)

func (p *CustomGoParser) handleFuncDecls(fun *ast.FuncDecl) string {
	start := p.fset.Position(fun.Pos())
	end := p.fset.Position(fun.End())
	name := p.fset.Position(fun.Name.Pos())
	if fun.Recv == nil {
		fmt.Printf("Function %s: name at %s | Block from %s to %s \n", fun.Name.Name, name, start, end)
		return fun.Name.Name
	} else {
		recvType := types.ExprString(fun.Recv.List[0].Type)
		if p.typesInfo != nil {
			if t := p.typesInfo.TypeOf(fun.Recv.List[0].Type); t != nil {
				recvType = t.String()
			}
		}
		fmt.Printf("Function %s - %s: name at %s | Block from %s to %s \n", fun.Name.Name, recvType, name, start, end)
		return fmt.Sprintf("(%s).%s", recvType, fun.Name.Name)
	}
}

func (p *CustomGoParser) handleFuncCalls(caller string, call *ast.CallExpr) {
	pos := p.fset.Position(call.Pos())
	callee, callType := types.ExprString(call.Fun), "UNRESOLVED"
	if p.typesInfo != nil {
		if fn := typeutil.StaticCallee(p.typesInfo, call); fn != nil {
			callee, callType = fn.FullName(), "FUNCTION"
			if fn.Signature().Recv() != nil {
				callType = "METHOD"
			}
		}
	}
	fmt.Printf("Caller: %s | Callee: %s | Type: %s | Position: %s \n", caller, callee, callType, pos)
}
