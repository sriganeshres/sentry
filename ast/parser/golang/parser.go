// Package golang for golang parser
package golang

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type CustomGoParser struct {
	moduleName string
	fset       *token.FileSet
}

func NewCustomParser(projectDir string) *CustomGoParser {
	return &CustomGoParser{
		moduleName: detectModuleName(projectDir),
		fset:       token.NewFileSet(),
	}
}

func (p *CustomGoParser) ParseFile(filename string) {
	astReturn, err := parser.ParseFile(p.fset, filename, nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ast.Print(fset, astReturn)

	ast.Inspect(astReturn, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ImportSpec:
			p.handleImports(x)
			return false
		case *ast.FuncDecl:
			p.handleFuncDecls(x)
			return true
		case *ast.TypeSpec:
			switch t := x.Type.(type) {
			case *ast.StructType:
				p.handleStructs(x.Name.Name, t)
			case *ast.InterfaceType:
				p.handleInterfaces(x.Name.Name, t)
			}
			return false
		case *ast.CallExpr:
			p.handleFuncCalls(x)
			return true
		}
		return true
	})
}
