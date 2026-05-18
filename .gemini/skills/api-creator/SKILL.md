---
name: api-creator
description: Automates the creation of Go SDK modules based on architectural blueprints derived from OpenAPI specifications. Use when Gemini CLI needs to generate Kiota-compliant code, including RequestBuilders, Models, and RequestConfigurations, that leverages the `internal/new/` abstractions of the ServiceNow Go SDK.
---

# 🛠 API Creator

The `api-creator` skill is a specialized tool for automating the generation of ServiceNow Go SDK modules. It ensures that all generated code follows the established Kiota-based architecture, naming conventions, and structural patterns.

## Core Capabilities

- **Code Generation**: Generates `RequestBuilder`, `Model`, `RequestConfiguration`, and `QueryParameters` Go files.
- **Pattern Adherence**: Strictly follows the templates defined in `assets/templates/`.
- **Dependency Management**: Correctly handles imports for Kiota abstractions and internal SDK packages.
- **Unit Testing**: Generates initial unit tests for the generated components using `testify/mock`.

## 🔄 Workflow

1. **Blueprint Ingestion**: Start by reading the Technical Blueprint provided by the `openapi-architect`.
2. **Directory Setup**: Create the target package directory if it doesn't exist.
3. **RequestBuilder Generation**:
    - **One Builder Per Segment**: Implement a dedicated `RequestBuilder` struct for each path segment of the API (e.g., `/now`, `/cmdb` each get their own builder).
    - **Namespace Connections**: ONLY top-level namespaces (e.g., `now`, `sn_cdm`) should be connected directly to the root `ServiceNowClient`.
    - **Hierarchical Linking**: All other path segments must be linked hierarchically through their parent builders (e.g., `Now() -> Cmdb() -> Instance()`).
    - Use `newInternal.NewBaseRequestBuilder`.
    - Implement methods for each HTTP verb following the signature: `func (rB *RB) Method(ctx context.Context, config *Config) (*Response, error)`.
    - **ToVerbRequestInformation**: Implement request information builders that wrap `abstractions.RequestInformation` with `newInternal.KiotaRequestInformation`. Use `internal.IsNil` for null checks and ensure `Accept` headers are set.
4. **Model Generation**:
    - Implement `Parsable` models using `internal/new.BaseModel`.
    - Generate `Serialize` and `GetFieldDeserializers` methods.
5. **Configuration & Parameters**:
    - **QueryParameters**: Generate structs with `url` tags for each parameter (e.g., `type %Name%RequestBuilder%Verb%QueryParameters struct { ... }`). Use `*int`, `*string`, etc., for fields to support `omitempty`.
    - **RequestConfiguration**: Define as a type alias of `abstractions.RequestConfiguration[%QueryParameters%]`. (e.g., `type %Name%RequestBuilder%Verb%RequestConfiguration = abstractions.RequestConfiguration[%QueryParameters%]`).
6. **Verification**: Run `go fmt` and `go vet` on the generated files to ensure basic correctness.

## 📏 Mandatory Patterns

- **Path Hierarchy**: Strictly adhere to the "one builder per segment" rule. Do not skip path segments or flatten the hierarchy.
- **Client Connectivity**: Do not connect sub-paths (e.g., `cmdb`) to the `ServiceNowClient`. Use parent builders to navigate the hierarchy.
- **Package Naming**: Use lowercase package names without underscores (e.g., `tableapi`, `attachmentapi`).
- **Internal Aliases**: Use `newInternal` for `github.com/michaeldcanady/servicenow-sdk-go/internal/new`.
- **URL Templates**: Define URL templates as constants within the RequestBuilder files.
- **Configuration Alias**: Always use the `abstractions.RequestConfiguration[T]` type alias for request configurations.
- **URL Tags**: Always include `url:"parameter_name,omitempty"` tags in query parameter structs.
- **Templates**: Always use the templates located in `assets/templates/`:
    - `request_builder.go.tmpl`: Template for `RequestBuilder` files.
    - `model.go.tmpl`: Template for `Parsable` model files.
    - `query_parameters.go.tmpl`: Template for `QueryParameters` structs.
    - `request_configuration.go.tmpl`: Template for `RequestConfiguration` type aliases.

## ⚖️ Usage Distinctions

- **Use `api-creator` for**: Bulk generation of API support code once a blueprint is approved.
- **Do NOT use for**: Designing the API architecture (`openapi-architect`) or fixing complex logic bugs (`software-engineer`).

## 📚 Resources

- **Process Guide**: Refer to [docs/api_generation_process.md](../../../docs/api_generation_process.md) for detailed templates and examples.
