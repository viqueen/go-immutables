package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	fileName := os.Getenv("GOFILE")

	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("could not parse file: %v", err)
	}

	var structNames []string
	for _, declaration := range file.Decls {
		if genericDeclaration, ok := declaration.(*ast.GenDecl); ok && genericDeclaration.Tok == token.TYPE {
			for _, spec := range genericDeclaration.Specs {
				if typeSpec, typeSpecOk := spec.(*ast.TypeSpec); typeSpecOk {
					structNames = append(structNames, typeSpec.Name.Name)
				}
			}
		}
	}

	fmt.Printf("structs: %s\n", structNames)
}
