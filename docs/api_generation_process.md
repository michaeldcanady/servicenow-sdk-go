# Process Plan: AI-Driven API Support Generation

This document outlines the systematic process for building new ServiceNow API support in the Go SDK using AI, based on OpenAPI specifications.

## 1. Overview
The goal is to automate the generation of Kiota-compliant RequestBuilders, Models, and Configurations from OpenAPI specs stored in the `spec/` directory, ensuring strict adherence to the SDK's established patterns and `internal/new/` abstractions.

## 2. Core Components

### 2.1. New Skill: `openapi-architect`
A specialized skill for mapping OpenAPI constructs (paths, methods, schemas) to Kiota-based SDK structures.
- **Location**: `.gemini/skills/openapi-architect/SKILL.md`
- **Responsibility**: Translating Spec to Blueprint.

### 2.2. New Agent: `api-generator`
A sub-agent tasked with the actual code generation based on the blueprint provided by the `openapi-architect`.
- **Location**: `.gemini/agents/api-generator.md`
- **Responsibility**: Generating Go code and Unit Tests.

## 3. Workflow Phases

### Phase 1: Research & Mapping (Inquiry)
1. **Spec Ingestion**: Read the OpenAPI spec from `spec/`.
2. **Analysis**: Use `openapi-architect` to identify URL hierarchies, query parameters, and data models.
3. **Consistency Check**: Verify mappings against existing modules (e.g., `attachment-api/`).
4. **Output**: A **Technical Blueprint** (Markdown) describing the proposed package structure, RequestBuilders, and Parsable models.

### Phase 2: Strategy & Approval (Directives)
1. **Blueprint Review**: User/Architect reviews the Blueprint.
2. **Kiota Alignment**: `kiota-architect` ensures the design follows `internal/new/` patterns.
3. **Approval**: Obtain user confirmation to proceed with implementation.

### Phase 3: Automated Implementation (Execution)
1. **Module Creation**: `api-generator` creates the new package directory.
2. **Code Generation**:
    - `RequestBuilder` files using `internal/new.BaseRequestBuilder`.
    - `Model` files implementing `serialization.Parsable` with `internal/new.BaseModel`.
    - `Configuration` and `QueryParameters` files.
3. **Boilerplate Integration**: Ensure imports and exports are correctly handled.

### Phase 4: Validation & Quality Assurance
1. **Unit Test Generation**: `qa-engineer` generates tests for all new components.
2. **Mocking**: Use `testify/mock` and `httpmock`.
3. **Execution**: Run `go test ./...` to verify the implementation.
4. **Documentation**: `docs-engineer` updates `Readme.md` and generates API docs.

## 4. Templates & Conventions

### 4.1. RequestBuilder Pattern
```go
type %Name%RequestBuilder struct {
    newInternal.RequestBuilder
}

func New%Name%RequestBuilder(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *%Name%RequestBuilder {
    return &%Name%RequestBuilder{
        newInternal.NewBaseRequestBuilder(requestAdapter, %URLTemplate%, pathParameters),
    }
}
```

### 4.2. Parsable Model Pattern
```go
type %Model% struct {
    newInternal.BaseModel
}

func New%Model%() *%Model% {
    return &%Model%{
        BaseModel: *newInternal.NewBaseModel(),
    }
}
```

## 5. Automation Commands
- `gemini invoke api-generator --spec spec/my_api.json` (Conceptual)
