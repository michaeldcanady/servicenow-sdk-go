package main

type Endpoint struct {
	Package     string
	Name        string
	MethodName  string // Name used for the accessor method in the parent builder
	FullPath    string
	Kind        string // "collection" or "item"
	NameLower   string
	GenericType string // e.g. "T" if this is a generic builder
	BaseType    string // e.g. "model.ServiceNowItem" constraint
	ItemKey     string // The key used for the item parameter, e.g. "sys_id"

	Methods []Method

	HasItemChild bool
	ItemChild    *Endpoint
	Children     []Endpoint

	// Imports needed by this endpoint
	Imports []string
}

type Method struct {
	Name             string // e.g. "Get", "Post"
	HTTPVerb         string // e.g. "GET", "POST"
	RequestType      string // e.g. "*Media", "abstractions.MultipartBody", or "T"
	ResponseType     string // e.g. "*AttachmentCollectionResponseModel" or "newInternal.ServiceNowCollectionResponse[T]"
	ResponseFactory  string // e.g. "CreateAttachmentCollectionResponseFromDiscriminatorValue"
	QueryParameters  []QueryParam
	IsStream         bool
	IsMultipart      bool
	IsPrimitive      bool
	PrimitiveType    string // e.g. "[]byte"
	IsCollection     bool   // If true, ParseHeaders will be called
	IsAttachmentFile bool   // If true, special logic for Attachment File Get will be used
}

func (e *Endpoint) AllQueryParameters() []QueryParam {
	seen := map[string]bool{}
	var out []QueryParam
	for _, m := range e.Methods {
		for _, p := range m.QueryParameters {
			if !seen[p.Name] {
				seen[p.Name] = true
				out = append(out, p)
			}
		}
	}
	return out
}
