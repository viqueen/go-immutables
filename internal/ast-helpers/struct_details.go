package ast_helpers

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/viqueen/go-immutables/internal/collections"
)

type NameInfo struct {
	Upper string
	Lower string
}

type FieldDetails struct {
	Name NameInfo
	Type string
}

type StructDetails struct {
	PackageName string
	Name        NameInfo
	Fields      []FieldDetails
}

func ExtractDetails(packageName string, typeSpec *ast.TypeSpec) StructDetails {
	structName := typeSpec.Name.Name
	structDetails := StructDetails{
		PackageName: packageName,
		Name: NameInfo{
			Upper: typeSpec.Name.Name,
			Lower: strings.ToLower(structName[0:1]) + structName[1:],
		},
		Fields: extractPrivateFields(typeSpec),
	}
	return structDetails
}

func extractPrivateFields(typeSpec *ast.TypeSpec) []FieldDetails {
	var fields []FieldDetails
	if structType, ok := typeSpec.Type.(*ast.StructType); ok {
		publicFields := collections.Filter(structType.Fields.List, func(item *ast.Field) bool {
			return !ast.IsExported(item.Names[0].Name)
		})
		for _, field := range publicFields {
			fieldName := field.Names[0].Name
			fieldType := typeAsString(field.Type)
			fieldDetails := FieldDetails{
				Name: NameInfo{
					Upper: strings.ToUpper(fieldName[0:1]) + fieldName[1:],
					Lower: fieldName,
				},
				Type: fieldType,
			}
			fields = append(fields, fieldDetails)
		}
	}
	return fields
}

func typeAsString(expression ast.Expr) string {
	switch t := expression.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", typeAsString(t.Key), typeAsString(t.Value))
	default:
		return "unknown"
	}
}
