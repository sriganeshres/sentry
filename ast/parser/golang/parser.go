// Package golang for golang parser
package golang

import (
	"go/ast"
	"go/token"
	"go/types"

	myTypes "sentry/ast/types"
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

func (p *CustomGoParser) ParseFile(file *ast.File) myTypes.FileNode {
	fileNode := myTypes.FileNode{
		PackageName: file.Name.Name,
	}
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ImportSpec:
			imp := p.handleImports(x)
			fileNode.Imports = append(fileNode.Imports, imp)
			return false
		case *ast.FuncDecl:
			fn := p.handleFuncDecls(x)
			if x.Body != nil {
				ast.Inspect(x.Body, func(n ast.Node) bool {
					if call, ok := n.(*ast.CallExpr); ok {
						dep := p.handleFuncCalls(call)
						fn.Dependencies = append(fn.Dependencies, dep)
					}
					return true
				})
			}
			fileNode.Functions = append(fileNode.Functions, fn)
			return false
		case *ast.TypeSpec:
			var TypeDefi myTypes.TypeDefinition
			switch t := x.Type.(type) {
			case *ast.StructType:
				TypeDefi = p.handleStructs(x.Name.Name, t)
			case *ast.InterfaceType:
				TypeDefi = p.handleInterfaces(x.Name.Name, t)
			default:
				TypeDefi = p.handleOtherTypes(x)
			}

			TypeDefi.TypeParams = p.handleTypeParams(x.TypeParams)
			fileNode.Types = append(fileNode.Types, TypeDefi)
			return false
		}
		return true
	})

	return fileNode
}
