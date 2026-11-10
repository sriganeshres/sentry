# File structure

```
sentry/
│
├── ast/                         # AST analysis subsystem
│   ├── parser/                  # Language-specific parsers
│   │   ├── golang/
│   │   │   ├── parser.go
│   │   │   ├── functions.go
│   │   │   ├── structs.go
│   │   │   ├── interfaces.go
│   │   │   └── imports.go
│   │   │
│   │   └── treesitter/          # Future multi-language support
│   │
│   ├── extractor/               # Extract architectural information
│   │   ├── calls.go
│   │   ├── dependencies.go
│   │   ├── api.go
│   │   └── models.go
│   │
│   ├── analyzer/                # Analyze extracted information
│   │   ├── symbols.go
│   │   └── relationships.go
│   │
│   └── types/                   # AST extraction result structures
│       ├── nodes.go
│       └── metadata.go
│
├── graph/                       # Knowledge graph (future)
├── impact/                      # Blast-radius analysis (future)
├── compression/                 # Context compression (future)
├── reasoning/                   # Qwen reasoning (future)
├── ollama/                      # Ollama integration (future)
│
├── cmd/
│   └── eidos/
│       └── main.go
│
├── go.mod
└── README.md
```
