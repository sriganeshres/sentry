package types

// FileNode represents a single language-agnostic file map
type FileNode struct {
	FilePath    string
	PackageName string
	Imports     []string
	Functions   []Function
	Types       []TypeDefinition
}

type Function struct {
	Name     string
	Receiver string // Blank if global function, filled if it's a struct method
	StartPos int    // Byte/Line start for the extractor to target
	EndPos   int    // Byte/Line end
}

type TypeDefinition struct {
	Name   string
	Kind   string   // "struct" or "interface"
	Fields []string // Field or Method signatures
}
