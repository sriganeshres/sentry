package main

import (
	"sentry/ast/parser/golang"
)

func main() {
	p := golang.NewCustomParser(".")
	p.ParseFile("dummy.go")
}
