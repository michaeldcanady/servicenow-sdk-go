package main

import (
	"log"
	"os"
)

func main() {
	var openAPIs []string
	outputDir := "generated"

	switch len(os.Args) {
	case 3:
		openAPIPath := os.Args[1]
		outputDir = os.Args[2]

		openAPIs = append(openAPIs, openAPIPath)
	case 1:
		openAPIs = []string{
			"cmd/generator/attachment_openapi.json",
			"cmd/generator/table_openapi.json",
			"cmd/generator/account_openapi.json",
		}
	default:
		log.Fatal("Usage: generator <openapi_spec> <output_dir> or run without arguments for defaults")
	}

	for _, openAPIPath := range openAPIs {
		if err := processOpenAPI(openAPIPath, outputDir); err != nil {
			log.Fatalf("Error processing %s: %v", openAPIPath, err)
		}
	}

	log.Println("Generation complete!")
}

func processOpenAPI(openAPIPath, outputDir string) error {
	log.Printf("Processing %s into %s...", openAPIPath, outputDir)
	api, schemas, err := ParseOpenAPI(openAPIPath)
	if err != nil {
		return err
	}

	if err := EnsureDir(outputDir); err != nil {
		return err
	}

	rbTemplate := "/workspaces/servicenow-sdk-go/.gemini/skills/software-engineer/references/templates/request_builder.go.tmpl"
	modelTemplate := "/workspaces/servicenow-sdk-go/.gemini/skills/software-engineer/references/templates/model.go.tmpl"

	return GenerateAPI(
		api,
		schemas,
		outputDir,
		rbTemplate,
		modelTemplate,
	)
}
