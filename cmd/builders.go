package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/viqueen/go-immutables/internal/builders"
)

func main() {
	fileName := os.Getenv("GOFILE")

	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("could not parse file: %v", err)
	}

	for _, declaration := range file.Decls {
		if genericDeclaration, ok := declaration.(*ast.GenDecl); ok && genericDeclaration.Tok == token.TYPE {
			for _, spec := range genericDeclaration.Specs {
				if typeSpec, typeSpecOk := spec.(*ast.TypeSpec); typeSpecOk {
					builderCode := builders.GenerateStructBuilder(typeSpec)
					fmt.Printf("%s\n", builderCode)
				}
			}
		}
	}
}
