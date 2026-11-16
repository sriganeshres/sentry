package golang

import (
	"go/ast"
	"go/types"

	myTypes "sentry/ast/types"

	"golang.org/x/tools/go/types/typeutil"
)

func (p *CustomGoParser) handleFuncDecls(fun *ast.FuncDecl) myTypes.Function {
	start := p.fset.Position(fun.Pos())
	end := p.fset.Position(fun.End())
	if fun.Recv == nil {
		return myTypes.Function{
			Name:       fun.Name.Name,
			StartPos:   start.Line,
			EndPos:     end.Line,
			TypeParams: p.handleTypeParams(fun.Type.TypeParams),
		}
	} else {
		recvType := types.ExprString(fun.Recv.List[0].Type)
		if p.typesInfo != nil {
			if t := p.typesInfo.TypeOf(fun.Recv.List[0].Type); t != nil {
				recvType = t.String()
			}
		}
		return myTypes.Function{
			Name:       fun.Name.Name,
			StartPos:   start.Line,
			EndPos:     end.Line,
			Receiver:   recvType,
			TypeParams: p.handleTypeParams(fun.Type.TypeParams),
		}
	}
}

func (p *CustomGoParser) handleFuncCalls(call *ast.CallExpr) myTypes.FunctionDependency {
	pos := p.fset.Position(call.Pos())
	dependecy := myTypes.FunctionDependency{
		Callee: types.ExprString(call.Fun),
		Line:   pos.Line,
		Column: pos.Column,
	}
	if p.typesInfo == nil {
		dependecy.Kind = "UNRESOLVED"
		return dependecy
	}

	if tc, ok := p.typesInfo.Types[call.Fun]; ok && tc.IsType() {
		dependecy.Kind = "TYPE_CONVERSION"
		dependecy.Resolved = true
		return dependecy
	}

	if ident, ok := call.Fun.(*ast.Ident); ok {
		if _, isbuiltin := p.typesInfo.Uses[ident].(*types.Builtin); isbuiltin {
			dependecy.Kind = "BUILTIN"
			dependecy.Resolved = true
			return dependecy
		}
	}

	if p.typesInfo != nil {
		if fn := typeutil.StaticCallee(p.typesInfo, call); fn != nil {
			dependecy.Callee = fn.FullName()
			dependecy.Resolved = true
			dependecy.Kind = "FUNCTION"
			if fn.Signature().Recv() != nil {
				dependecy.Kind = "METHOD"
			}
		}

		return dependecy
	}

	if tc, ok := p.typesInfo.Types[call.Fun]; ok && tc.Type != nil {
		if _, isSignature := tc.Type.Underlying().(*types.Signature); isSignature {
			dependecy.Kind = "DYNAMIC_CALL"
			return dependecy
		}
	}
	dependecy.Kind = "UNRESOLVED"
	return dependecy
}
