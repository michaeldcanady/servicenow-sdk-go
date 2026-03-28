package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func toCamelCase(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func pluckNames(params []QueryParam) []string {
	out := make([]string, len(params))
	for i, p := range params {
		out[i] = p.Name
	}
	return out
}

func RenderTemplate(templatePath string, data any, outputPath string) error {
	tmpl := template.New("rb").Funcs(template.FuncMap{
		"camel":  toCamelCase,
		"pascal": toPascalCase,
		"join":   strings.Join,
		"pluck":  pluckNames,
		"goType": func(s *Schema) string {
			if s.Ref != "" {
				// Assumes Ref is like "#/components/schemas/Account"
				parts := strings.Split(s.Ref, "/")
				return parts[len(parts)-1]
			}
			switch s.Type {
			case "string":
				if s.Format == "date-time" {
					return "*time.Time"
				}
				return "*string"
			case "integer":
				return "*int64"
			case "boolean":
				return "*bool"
			case "number":
				return "*float64"
			case "array":
				if s.Items != nil {
					// Recursively call for items
					// For simple cases, just return []T
					itemType := "string" // default
					if s.Items.Type != "" {
						// Simple mapping for now
						switch s.Items.Type {
						case "string":
							itemType = "string"
						case "integer":
							itemType = "int64"
						}
					} else if s.Items.Ref != "" {
						parts := strings.Split(s.Items.Ref, "/")
						itemType = parts[len(parts)-1]
					}
					return "[]" + itemType
				}
				return "[]string"
			default:
				return "*string"
			}
		},
		"isSlice": func(s *Schema) bool {
			return s.Type == "array"
		},
		"serializeFunc": func(keyExpr string, s *Schema) string {
			if s.Type == "array" {
				return fmt.Sprintf("internalSerialization.SerializeStringToSliceFunc(%s, \" \")", keyExpr)
			}
			switch s.Type {
			case "string":
				if s.Format == "date-time" {
					return fmt.Sprintf("internalSerialization.SerializeStringToTimeFunc(%s, \"2006-01-02 15:04:05\")", keyExpr)
				}
				return fmt.Sprintf("internalSerialization.SerializeStringFunc(%s)", keyExpr)
			case "integer":
				return fmt.Sprintf("internalSerialization.SerializeStringToInt64Func(%s)", keyExpr)
			case "boolean":
				return fmt.Sprintf("internalSerialization.SerializeStringToBoolFunc(%s)", keyExpr)
			case "number":
				return fmt.Sprintf("internalSerialization.SerializeStringToFloat64Func(%s)", keyExpr)
			default:
				return fmt.Sprintf("internalSerialization.SerializeStringFunc(%s)", keyExpr)
			}
		},
		"deserializeFunc": func(s *Schema) string {
			if s.Type == "array" {
				return "internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToPrimitiveSlice(\" \", func(s string) (string, error) { return s, nil }))"
			}
			switch s.Type {
			case "string":
				if s.Format == "date-time" {
					return "internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToTimePtr(\"2006-01-02 15:04:05\"))"
				}
				return "internalSerialization.DeserializeStringFunc()"
			case "integer":
				return "internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)"
			case "boolean":
				return "internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToBoolPtr)"
			case "number":
				return "internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)"
			default:
				return "internalSerialization.DeserializeStringFunc()"
			}
		},
	})

	tmpl, err := tmpl.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, filepath.Base(templatePath), data); err != nil {
		return err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		// If formatting fails, we still write the unformatted source to disk for debugging
		_ = os.WriteFile(outputPath, buf.Bytes(), 0o644)
		return fmt.Errorf("error formatting generated code for %s: %v", outputPath, err)
	}

	return os.WriteFile(outputPath, formatted, 0o644)
}
