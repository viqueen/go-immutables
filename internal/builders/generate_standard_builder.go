package builders

import (
	"fmt"
	"go/ast"
	"log"
	"os"
	"strings"
	"text/template"

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

var standardBuilderTemplate = `
package {{.PackageName}}

type {{.Name.Upper}}Builder interface {
    Build() *{{.Name.Upper}}
    {{range .Fields}}
    Set{{.Name.Upper}}(value {{.Type}}) {{$.Name.Upper}}Builder
    {{end}}
}

type {{.Name.Lower}}Builder struct {
    target *{{.Name.Upper}}
}

func New{{.Name.Upper}}Builder() {{.Name.Upper}}Builder {
    return &{{.Name.Lower}}Builder{
        target: &{{.Name.Upper}}{},
    }
}

{{range .Fields}}
func (b *{{$.Name.Lower}}Builder) Set{{.Name.Upper}}(value {{.Type}}) {{$.Name.Upper}}Builder {
    b.target.{{.Name.Lower}} = value
    return b
}
{{end}}

func (b *{{.Name.Lower}}Builder) Build() *{{.Name.Upper}} {
    return b.target
}
`

func GenerateStandardStructBuilder(packageName string, typeSpec *ast.TypeSpec) string {
	structDetails := StructDetails{
		PackageName: packageName,
		Name: NameInfo{
			Upper: typeSpec.Name.Name,
			Lower: strings.ToLower(typeSpec.Name.Name),
		},
		Fields: extractPrivateFields(typeSpec),
	}

	if (len(structDetails.Fields)) == 0 {
		return ""
	}

	parsed, err := template.New("builder").Parse(standardBuilderTemplate)
	if err != nil {
		log.Fatalf("could not parse standard builder template: %v", err)
	}

	base := structDetails.Name.Lower
	fileName := base + "_standard_builder.go"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("could not create standard builder file: %v", err)
	}
	defer file.Close()

	err = parsed.Execute(file, structDetails)
	if err != nil {
		log.Fatalf("could not execute standard builder template: %v", err)
	}

	return fileName
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