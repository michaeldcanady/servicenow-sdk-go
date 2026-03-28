package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 3 {
		openAPIPath := os.Args[1]
		outputDir := os.Args[2]

		if err := processOpenAPI(openAPIPath, outputDir); err != nil {
			log.Fatalf("Error processing %s: %v", openAPIPath, err)
		}
	} else if len(os.Args) == 1 {
		openAPIs := []string{
			//"cmd/generator/attachment_openapi.json",
			//"cmd/generator/table_openapi.json",
			"cmd/generator/policy_openapi.json",
			//"cmd/generator/account_openapi.json",
		}

		for _, openAPIPath := range openAPIs {
			if err := processOpenAPI(openAPIPath, "generated"); err != nil {
				log.Fatalf("Error processing %s: %v", openAPIPath, err)
			}
		}
	} else {
		log.Fatal("Usage: generator <openapi_spec> <output_dir> or run without arguments for defaults")
	}

	log.Println("Generation complete!")
}

func processOpenAPI(openAPIPath, outputDir string) error {
	log.Printf("Processing %s...", openAPIPath)
	api, schemas, err := ParseOpenAPI(openAPIPath)
	if err != nil {
		return err
	}

	apiOutputDir := filepath.Join(outputDir, ToSnakeCase(api.Name))
	if err := EnsureDir(apiOutputDir); err != nil {
		return err
	}

	log.Printf("Writing to %s...", apiOutputDir)

	templates := map[string]string{
		"builder":               "/workspaces/servicenow-sdk-go/.gemini/skills/software-engineer/references/templates/builder.go.tmpl",
		"query_parameters":      "/workspaces/servicenow-sdk-go/.gemini/skills/software-engineer/references/templates/query_parameters.go.tmpl",
		"request_configuration": "/workspaces/servicenow-sdk-go/.gemini/skills/software-engineer/references/templates/request_configuration.go.tmpl",
		"model":                 "/workspaces/servicenow-sdk-go/.gemini/skills/software-engineer/references/templates/model.go.tmpl",
	}

	return GenerateAPI(
		api,
		schemas,
		apiOutputDir,
		templates,
	)
}
