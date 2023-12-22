package main

import (
	"fmt"
	"log"
	"os"

	asthelpers "github.com/viqueen/go-immutables/internal/ast-helpers"
	"github.com/viqueen/go-immutables/internal/builders"
)

func main() {
	fileName := os.Getenv("GOFILE")

	fileDetails, err := asthelpers.ExtractTypeSpecs(fileName)
	if err != nil {
		log.Fatalf("could not extract type specs: %v", err)
	}

	for _, typeSpec := range fileDetails.TypeSpecs {
		builderFileName := builders.GenerateStandardStructBuilder(fileDetails.PackageName, typeSpec)
		if builderFileName != "" {
			fmt.Printf("generated %s\n", builderFileName)
		}
	}
}
