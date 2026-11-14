// Package golang for golang parser
package golang

import (
	"go/ast"
	"go/token"
	"go/types"
)

type CustomGoParser struct {
	moduleName string
	fset       *token.FileSet
	typesInfo  *types.Info
}

func NewCustomParser(projectDir string, fset *token.FileSet, info *types.Info) *CustomGoParser {
	return &CustomGoParser{
		moduleName: detectModuleName(projectDir),
		fset:       fset,
		typesInfo:  info,
	}
}

func (p *CustomGoParser) ParseFile(file *ast.File) {
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ImportSpec:
			p.handleImports(x)
			return false
		case *ast.FuncDecl:
			caller := p.handleFuncDecls(x)
			if x.Body != nil {
				ast.Inspect(x.Body, func(n ast.Node) bool {
					if call, ok := n.(*ast.CallExpr); ok {
						p.handleFuncCalls(caller, call)
					}
					return true
				})
			}
			return false
		case *ast.TypeSpec:
			switch t := x.Type.(type) {
			case *ast.StructType:
				p.handleStructs(x.Name.Name, t)
			case *ast.InterfaceType:
				p.handleInterfaces(x.Name.Name, t)
			}
			return false
		}
		return true
	})
}
