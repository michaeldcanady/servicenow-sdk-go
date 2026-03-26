package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func GenerateAPI(root *Endpoint, schemas map[string]*Schema, outputDir string, rbTemplatePath, modelTemplatePath string) error {
	// Derive package name from root
	root.Package = root.NameLower + "api"

	if err := generateEndpointRecursive(root, outputDir, rbTemplatePath, root.Package); err != nil {
		return err
	}

	for name, schema := range schemas {
		if err := generateModel(name, schema, outputDir, modelTemplatePath, root.Package); err != nil {
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

	fileName := fmt.Sprintf("%s.go", toSnakeCase(name))
	outFile := filepath.Join(outputDir, fileName)

	return RenderTemplate(templatePath, data, outFile)
}

func toSnakeCase(s string) string {
	var out []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			out = append(out, '_')
		}
		out = append(out, r)
	}
	return strings.ToLower(string(out))
}

func mergeQueryParams(groups ...[]QueryParam) []QueryParam {
	seen := map[string]bool{}
	out := []QueryParam{}

	for _, group := range groups {
		for _, p := range group {
			if !seen[p.Name] {
				seen[p.Name] = true
				out = append(out, p)
			}
		}
	}
	return out
}

func generateEndpointRecursive(ep *Endpoint, outputDir, templatePath, pkg string) error {
	fmt.Println(">>> GENERATING:", ep.Name, "Kind:", ep.Kind, "Path:", ep.FullPath)

	ep.Package = pkg

	fileName := fmt.Sprintf("%s_request_builder.go", toSnakeCase(ep.Name))
	outFile := filepath.Join(outputDir, fileName)

	if err := RenderTemplate(templatePath, ep, outFile); err != nil {
		return err
	}

	// Generate non-item children
	for i := range ep.Children {
		child := &ep.Children[i]
		if child.Kind != "item" {
			// Non-item children generate methods + files
			if err := generateEndpointRecursive(child, outputDir, templatePath, pkg); err != nil {
				return err
			}
		} else {
			// Item children DO NOT generate methods, but DO generate files
			fmt.Println("    GENERATING item child:", child.Name)
			if err := generateEndpointRecursive(child, outputDir, templatePath, pkg); err != nil {
				return err
			}
		}
	}

	// Generate item child
	if ep.HasItemChild && ep.ItemChild != nil {
		fmt.Println("    GENERATING item child:", ep.ItemChild.Name)
		if err := generateEndpointRecursive(ep.ItemChild, outputDir, templatePath, pkg); err != nil {
			return err
		}
	}

	return nil
}

func lower(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]|0x20) + s[1:]
}
