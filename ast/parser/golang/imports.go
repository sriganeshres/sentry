package golang

import (
	"fmt"
	"go/ast"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func (p *CustomGoParser) handleImports(imp *ast.ImportSpec) {
	rawPath, err := strconv.Unquote(imp.Path.Value)
	if err != nil {
		rawPath = imp.Path.Value
	}
	kind := p.classifyImport(rawPath)
	pos := p.fset.Position(imp.Pos())
	fmt.Printf("[%s] Import: %s - At %s\n", kind, rawPath, pos)
}

func (p *CustomGoParser) classifyImport(importPath string) string {
	if p.moduleName != "" && (importPath == p.moduleName || strings.HasPrefix(importPath, p.moduleName+"/")) {
		return "Internal"
	}
	first := strings.SplitN(importPath, "/", 2)[0]
	if !strings.Contains(first, ".") {
		return "STDLIB"
	}
	return "External"
}

func detectModuleName(filePath string) string {
	cmd := exec.Command("go", "env", "GOMOD")
	cmd.Dir = filepath.Dir(filePath)
	out, err := cmd.Output()
	goModPath := strings.TrimSpace(string(out))
	if err != nil || goModPath == "" || goModPath == os.DevNull {
		return ""
	}
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}
	return ""
}
