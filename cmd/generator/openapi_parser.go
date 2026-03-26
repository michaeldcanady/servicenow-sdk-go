package main

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
)

type OpenAPI struct {
	Paths      map[string]PathItem `json:"paths"`
	Components Components          `json:"components"`
}

type PathItem struct {
	Get     *Operation `json:"get"`
	Post    *Operation `json:"post"`
	Patch   *Operation `json:"patch"`
	Delete  *Operation `json:"delete"`
	Put     *Operation `json:"put"`
	XGoName string     `json:"x-go-name"`
}

type Components struct {
	Schemas map[string]*Schema `json:"schemas"`
}

type Operation struct {
	Summary             string      `json:"summary"`
	Parameters          []Parameter `json:"parameters"`
	XGoMethodName       string      `json:"x-go-method-name"`
	XGoRequestType      string      `json:"x-go-request-type"`
	XGoResponseType     string      `json:"x-go-response-type"`
	XGoRespFactory      string      `json:"x-go-response-factory"`
	XGoIsStream         bool        `json:"x-go-is-stream"`
	XGoIsMultipart      bool        `json:"x-go-is-multipart"`
	XGoIsPrimitive      bool        `json:"x-go-is-primitive"`
	XGoPrimitiveType    string      `json:"x-go-primitive-type"`
	XGoIsCollection     bool        `json:"x-go-is-collection"`
	XGoIsAttachmentFile bool        `json:"x-go-is-attachment-file"`
}

type Parameter struct {
	Name        string `json:"name"`
	In          string `json:"in"`
	Schema      Schema `json:"schema"`
	Description string `json:"description"`
}

type Schema struct {
	Type        string             `json:"type"`
	Description string             `json:"description"`
	Properties  map[string]*Schema `json:"properties"`
	Items       *Schema            `json:"items"`
	Ref         string             `json:"$ref"`
	Format      string             `json:"format"`
	Enum        []string           `json:"enum"`
}

func ParseOpenAPI(filePath string) (*Endpoint, map[string]*Schema, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	var oai struct {
		OpenAPI
		XGoGenericType string `json:"x-go-generic-type"`
		XGoBaseType    string `json:"x-go-base-type"`
	}
	if err := json.Unmarshal(data, &oai); err != nil {
		return nil, nil, err
	}

	// 1. Collect all paths and sort them by length to process parents before children
	paths := make([]string, 0, len(oai.Paths))
	for path := range oai.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	// 2. Map paths to Endpoints
	endpoints := make(map[string]*Endpoint)

	for _, path := range paths {
		item := oai.Paths[path]
		name := pathToName(path)
		if item.XGoName != "" {
			name = item.XGoName
		}
		ep := &Endpoint{
			Name:        name,
			MethodName:  pathToMethodName(path),
			NameLower:   strings.ToLower(name),
			FullPath:    path,
			Kind:        determineKind(path),
			GenericType: oai.XGoGenericType,
			BaseType:    oai.XGoBaseType,
			ItemKey:     extractItemKey(path),
		}

		if item.Get != nil {
			ep.Methods = append(ep.Methods, mapOperation("Get", "GET", item.Get))
		}
		if item.Post != nil {
			ep.Methods = append(ep.Methods, mapOperation("Post", "POST", item.Post))
		}
		if item.Patch != nil {
			ep.Methods = append(ep.Methods, mapOperation("Patch", "PATCH", item.Patch))
		}
		if item.Delete != nil {
			ep.Methods = append(ep.Methods, mapOperation("Delete", "DELETE", item.Delete))
		}
		if item.Put != nil {
			ep.Methods = append(ep.Methods, mapOperation("Put", "PUT", item.Put))
		}

		endpoints[path] = ep
	}

	// 3. Build hierarchy
	var root *Endpoint
	for _, path := range paths {
		ep := endpoints[path]
		parentPath := getParentPath(path)

		parent, parentExists := endpoints[parentPath]
		if parentExists {
			if ep.Kind == "item" && ep.ItemKey != "" && strings.HasSuffix(path, "{"+ep.ItemKey+"}") {
				parent.HasItemChild = true
				parent.ItemChild = ep
			} else {
				parent.Children = append(parent.Children, *ep)
			}
		} else {
			if root == nil {
				root = ep
			}
		}
	}

	return root, oai.Components.Schemas, nil
}

func mapOperation(defaultName, verb string, op *Operation) Method {
	name := defaultName
	if op.XGoMethodName != "" {
		name = op.XGoMethodName
	}

	m := Method{
		Name:             name,
		HTTPVerb:         verb,
		RequestType:      op.XGoRequestType,
		ResponseType:     op.XGoResponseType,
		ResponseFactory:  op.XGoRespFactory,
		QueryParameters:  mapParams(op.Parameters),
		IsStream:         op.XGoIsStream,
		IsMultipart:      op.XGoIsMultipart,
		IsPrimitive:      op.XGoIsPrimitive,
		PrimitiveType:    op.XGoPrimitiveType,
		IsCollection:     op.XGoIsCollection,
		IsAttachmentFile: op.XGoIsAttachmentFile,
	}

	return m
}

func pathToName(path string) string {
	parts := strings.Split(path, "/")
	name := ""
	for _, part := range parts {
		if part == "" || part == "api" || part == "now" || strings.HasPrefix(part, "v1") {
			continue
		}
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			name += "Item"
			continue
		}
		name += strings.Title(part)
	}
	return name
}

func pathToMethodName(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return ""
	}
	last := parts[len(parts)-1]
	if strings.HasPrefix(last, "{") && strings.HasSuffix(last, "}") {
		return "ByID"
	}
	return strings.Title(last)
}

func extractItemKey(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return ""
	}
	last := parts[len(parts)-1]
	if strings.HasPrefix(last, "{") && strings.HasSuffix(last, "}") {
		return last[1 : len(last)-1]
	}
	return ""
}

func determineKind(path string) string {
	if strings.HasSuffix(path, "}") {
		return "item"
	}
	return "collection"
}

func getParentPath(path string) string {
	if strings.HasSuffix(path, "/file") || strings.HasSuffix(path, "/upload") {
		return strings.TrimSuffix(strings.TrimSuffix(path, "/file"), "/upload")
	}
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash <= 0 {
		return ""
	}
	return path[:lastSlash]
}

func mapParams(params []Parameter) []QueryParam {
	var qps []QueryParam
	for _, p := range params {
		if p.In == "query" {
			qps = append(qps, QueryParam{
				Name:        p.Name,
				Type:        mapType(p.Schema.Type),
				Description: p.Description,
			})
		}
	}
	return qps
}

func mapType(t string) string {
	switch t {
	case "integer":
		return "int"
	case "string":
		return "string"
	case "boolean":
		return "bool"
	default:
		return "string"
	}
}
