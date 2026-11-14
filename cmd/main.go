package main

import (
	"fmt"
	"go/token"
	"log"

	"sentry/ast/parser/golang"

	"golang.org/x/tools/go/packages"
)

func main() {
	fset := token.NewFileSet()

	cfg := &packages.Config{
		Mode: packages.LoadAllSyntax,
		Dir:  ".",
		Fset: fset,
	}

	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		log.Fatal(err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		log.Fatal("failed to load all packages cleanly")
	}

	for _, pkg := range pkgs {
		fmt.Printf("\n==================Package: %s (%s)===================\n", pkg.Name, pkg.PkgPath)
		p := golang.NewCustomParser(".", fset, pkg.TypesInfo)
		for _, file := range pkg.Syntax {
			fmt.Printf("\n---File: %s---\n", fset.Position(file.Pos()).Filename)
			p.ParseFile(file)
		}
	}
}
