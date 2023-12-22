package ast_helpers_test

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	asthelpers "github.com/viqueen/go-immutables/internal/ast-helpers"
)

type testInput struct {
	packageName string
	typeSpec    *ast.TypeSpec
}

func TestExtractStructDetails(t *testing.T) {
	fileDetails, err := asthelpers.ExtractTypeSpecs("struct_examples_test.go")
	require.NoError(t, err)
	require.NotEmpty(t, fileDetails.PackageName)
	require.NotEmpty(t, fileDetails.TypeSpecs)

	inputMap := make(map[string]testInput)
	for _, typeSpec := range fileDetails.TypeSpecs {
		inputMap[typeSpec.Name.Name] = testInput{
			packageName: fileDetails.PackageName,
			typeSpec:    typeSpec,
		}
	}

	tests := []struct {
		name     string
		input    testInput
		expected asthelpers.StructDetails
	}{
		{
			name:  "IntegerTypes",
			input: inputMap["IntegerTypes"],
			expected: asthelpers.StructDetails{
				PackageName: "ast_helpers_test",
				Name:        asthelpers.NameInfo{Upper: "IntegerTypes", Lower: "integerTypes"},
				Fields: []asthelpers.FieldDetails{
					{Name: asthelpers.NameInfo{Upper: "Number8", Lower: "number8"}, Type: "int8"},
					{Name: asthelpers.NameInfo{Upper: "Number16", Lower: "number16"}, Type: "int16"},
					{Name: asthelpers.NameInfo{Upper: "Number32", Lower: "number32"}, Type: "int32"},
					{Name: asthelpers.NameInfo{Upper: "Number64", Lower: "number64"}, Type: "int64"},
					{Name: asthelpers.NameInfo{Upper: "Number", Lower: "number"}, Type: "int"},
				},
			},
		},
		{
			name:  "FloatTypes",
			input: inputMap["FloatTypes"],
			expected: asthelpers.StructDetails{
				PackageName: "ast_helpers_test",
				Name:        asthelpers.NameInfo{Upper: "FloatTypes", Lower: "floatTypes"},
				Fields: []asthelpers.FieldDetails{
					{Name: asthelpers.NameInfo{Upper: "Number32", Lower: "number32"}, Type: "float32"},
					{Name: asthelpers.NameInfo{Upper: "Number64", Lower: "number64"}, Type: "float64"},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := asthelpers.ExtractStructDetails(test.input.packageName, test.input.typeSpec)
			assert.Equal(t, test.expected, actual)
		})
	}
}
