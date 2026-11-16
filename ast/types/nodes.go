// Package types: Supplies types for ast for language-agnostic file
package types

// FileNode represents a single language-agnostic file map
type FileNode struct {
	FilePath    string
	PackageName string
	Imports     []Import
	Functions   []Function
	Types       []TypeDefinition
}

type Function struct {
	Name         string
	Receiver     string // Blank if global function, filled if it's a struct method
	StartPos     int    // Byte/Line start for the extractor to target
	EndPos       int    // Byte/Line end
	TypeParams   [][]string
	Dependencies []FunctionDependency
}

type FunctionDependency struct {
	Callee   string
	Kind     string
	Line     int
	Column   int
	Resolved bool
}

type TypeDefinition struct {
	Name           string
	Kind           string // "struct", "interface", "type"
	Alias          bool
	UnderLyingType string
	TypeParams     [][]string
	Fields         [][]string // Field signatures
	Methods        [][]string // Method signatures
	Embedded [][]string
}

type Import struct {
	Path   string
	Alias  string
	Kind   string // STDLIB, INTERNAL, EXTERNAL
	Line   int
	Column int
}
