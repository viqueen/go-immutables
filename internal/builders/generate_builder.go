package builders

import (
	"go/ast"
	"log"
	"strings"
	"text/template"

	"github.com/viqueen/go-immutables/internal/collections"
)

type FieldInfo struct {
	Name string
	Type string
}

type StructInfo struct {
	Name   string
	Fields []FieldInfo
}

var builderTemplate = `
type {{.Name}}Builder struct {
	target {{.Name}}
}

func New{{.Name}}Builder() *{{.Name}}Builder {
	return &{{.Name}}Builder{ {{.Name}}() }
}

{{range .Fields}}
func (b *{{$.Name}}Builder) With{{.Name}}({{.Name}} {{.Type}}) *{{$.Name}}Builder {
	b.target.{{.Name}} = {{.Name}}
	return b
}
{{end}}

func (b *{{.Name}}Builder) Build() {{.Name}} {
	return b.target
}
`

func GenerateStructBuilder(typeSpec *ast.TypeSpec) string {
	structInfo := StructInfo{
		Name:   typeSpec.Name.Name,
		Fields: extractPublicFields(typeSpec),
	}

	parsed, err := template.New("builder").Parse(builderTemplate)
	if err != nil {
		log.Fatalf("could not parse builder template: %v", err)
	}

	var builder strings.Builder
	err = parsed.Execute(&builder, structInfo)
	if err != nil {
		log.Fatalf("could not execute builder template: %v", err)
	}

	return builder.String()
}

func extractPublicFields(typeSpec *ast.TypeSpec) []FieldInfo {
	var fields []FieldInfo
	if structType, ok := typeSpec.Type.(*ast.StructType); ok {
		publicFields := collections.Filter(structType.Fields.List, func(item *ast.Field) bool {
			return ast.IsExported(item.Names[0].Name)
		})
		for _, field := range publicFields {
			fieldName := field.Names[0].Name
			fieldType := typeAsString(field.Type)
			fields = append(fields, FieldInfo{Name: fieldName, Type: fieldType})
		}
	}
	return fields
}

func typeAsString(expression ast.Expr) string {
	switch t := expression.(type) {
	case *ast.Ident:
		return t.Name
	default:
		return "unknown"
	}
}
