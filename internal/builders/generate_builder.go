package builders

import (
	"go/ast"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/viqueen/go-immutables/internal/collections"
)

type FieldInfo struct {
	Name string
	Type string
}

type StructInfo struct {
	PackageName string
	Name        string
	Fields      []FieldInfo
}

var builderTemplate = `
package {{.PackageName}}

type {{.Name}}Builder struct {
	target {{.Name}}
}

func New{{.Name}}Builder() *{{.Name}}Builder {
	return &{{.Name}}Builder{ {{.Name}}{} }
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

func GenerateStructBuilder(packageName string, typeSpec *ast.TypeSpec) string {
	structInfo := StructInfo{
		PackageName: packageName,
		Name:        typeSpec.Name.Name,
		Fields:      extractPublicFields(typeSpec),
	}

	if len(structInfo.Fields) == 0 {
		return ""
	}

	parsed, err := template.New("builder").Parse(builderTemplate)
	if err != nil {
		log.Fatalf("could not parse builder template: %v", err)
	}

	base := strings.ToLower(structInfo.Name)
	buildersFileName := base + "_builder.go"
	buildersFile, err := os.Create(buildersFileName)
	if err != nil {
		log.Fatalf("could not create builders file: %v", err)
	}
	defer buildersFile.Close()

	err = parsed.Execute(buildersFile, structInfo)
	if err != nil {
		log.Fatalf("could not execute builder template: %v", err)
	}

	return buildersFileName
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
