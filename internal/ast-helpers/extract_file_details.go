package ast_helpers

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type FileDetails struct {
	PackageName string
	TypeSpecs   []*ast.TypeSpec
}

func ExtractTypeSpecs(fileName string) (*FileDetails, error) {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, fileName, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	packageName := file.Name.Name
	var typeSpecs []*ast.TypeSpec

	for _, declaration := range file.Decls {
		if genericDeclaration, ok := declaration.(*ast.GenDecl); ok {
			switch genericDeclaration.Tok {
			case token.TYPE:
				for _, spec := range genericDeclaration.Specs {
					if typeSpec, typeSpecOk := spec.(*ast.TypeSpec); typeSpecOk {
						typeSpecs = append(typeSpecs, typeSpec)
					}
				}
			}
		}
	}

	return &FileDetails{
		PackageName: packageName,
		TypeSpecs:   typeSpecs,
	}, nil
}
