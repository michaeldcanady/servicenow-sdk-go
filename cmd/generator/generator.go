package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func GenerateAPI(root *Endpoint, schemas map[string]*Schema, outputDir string, templates map[string]string) error {
	// Derive package name from root
	root.Package = root.NameLower + "api"

	if err := generateEndpointRecursive(root, outputDir, templates, root.Package); err != nil {
		return err
	}

	for name, schema := range schemas {
		if err := generateModel(name, schema, outputDir, templates["model"], root.Package); err != nil {
			return err
		}
	}

	return nil
}

func generateEndpointRecursive(ep *Endpoint, outputDir string, templates map[string]string, pkg string) error {
	fmt.Println(">>> GENERATING ENDPOINT:", ep.Name, "Kind:", ep.Kind, "Path:", ep.FullPath)

	ep.Package = pkg

	// 1. Generate main builder file
	builderFileName := fmt.Sprintf("%s_request_builder.go", ToSnakeCase(ep.Name))
	builderOutFile := filepath.Join(outputDir, builderFileName)
	if err := RenderTemplate(templates["builder"], ep, builderOutFile); err != nil {
		return err
	}

	// 2. Generate per-method files
	for _, m := range ep.Methods {
		// a. Query Parameters
		if len(m.QueryParameters) > 0 {
			qpData := struct {
				Package     string
				BuilderName string
				MethodName  string
				Parameters  []QueryParam
			}{
				Package:     pkg,
				BuilderName: ep.Name,
				MethodName:  m.Name,
				Parameters:  m.QueryParameters,
			}
			qpFileName := fmt.Sprintf("%s_request_builder_%s_query_parameters.go", ToSnakeCase(ep.Name), strings.ToLower(m.Name))
			qpOutFile := filepath.Join(outputDir, qpFileName)
			if err := RenderTemplate(templates["query_parameters"], qpData, qpOutFile); err != nil {
				return err
			}
		}

		// b. Request Configuration
		rcData := struct {
			Package     string
			BuilderName string
			MethodName  string
			HasParams   bool
		}{
			Package:     pkg,
			BuilderName: ep.Name,
			MethodName:  m.Name,
			HasParams:   len(m.QueryParameters) > 0,
		}
		rcFileName := fmt.Sprintf("%s_request_builder_%s_request_configuration.go", ToSnakeCase(ep.Name), strings.ToLower(m.Name))
		rcOutFile := filepath.Join(outputDir, rcFileName)
		if err := RenderTemplate(templates["request_configuration"], rcData, rcOutFile); err != nil {
			return err
		}
	}

	// Generate non-item children
	for i := range ep.Children {
		child := &ep.Children[i]
		if child.Kind != "item" {
			if err := generateEndpointRecursive(child, outputDir, templates, pkg); err != nil {
				return err
			}
		} else {
			fmt.Println("    GENERATING item child:", child.Name)
			if err := generateEndpointRecursive(child, outputDir, templates, pkg); err != nil {
				return err
			}
		}
	}

	// Generate item child
	if ep.HasItemChild && ep.ItemChild != nil {
		fmt.Println("    GENERATING item child:", ep.ItemChild.Name)
		if err := generateEndpointRecursive(ep.ItemChild, outputDir, templates, pkg); err != nil {
			return err
		}
	}

	return nil
}

func generateModel(name string, schema *Schema, outputDir, templatePath, pkg string) error {
	fmt.Println(">>> GENERATING MODEL:", name)

	data := struct {
		Name       string
		Package    string
		Properties map[string]*Schema
	}{
		Name:       name,
		Package:    pkg,
		Properties: schema.Properties,
	}

	fileName := fmt.Sprintf("%s.go", ToSnakeCase(name))
	outFile := filepath.Join(outputDir, fileName)

	return RenderTemplate(templatePath, data, outFile)
}

func ToSnakeCase(s string) string {
	var out []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			out = append(out, '_')
		}
		out = append(out, r)
	}
	return strings.ToLower(string(out))
}

func lower(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]|0x20) + s[1:]
}
