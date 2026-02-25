# Architecture

The ServiceNow SDK for Go is designed to be a thin, high-performance wrapper around the ServiceNow REST APIs. It leverages the **Microsoft Kiota** ecosystem for its core networking and serialization logic.

## High-Level Components

### 1. Request Builders
Request builders provide a fluent API for constructing requests. They are organized to match the ServiceNow API hierarchy (e.g., `client.Now2().TableV2("incident")`).

### 2. Request Adapters
The `RequestAdapter` (from Kiota) is responsible for executing the actual HTTP requests. It handles:
- Authentication
- Serialization of request bodies
- Deserialization of response bodies
- URL template expansion

### 3. Authentication Providers
Located in `credentials/`, these providers manage the `Authorization` header. They can handle simple Basic Auth or complex OAuth2 token refreshes.

### 4. Serialization
The SDK uses Kiota's serialization abstractions. This allows for pluggable support for different content types (JSON, Text, Form-data), though ServiceNow primarily uses JSON.

## Design Patterns

- **Fluent Interface**: The SDK prioritizes a "fluent" style that makes code readable and easy to discover via IDE autocomplete.
- **Generics (V2)**: The V2 implementations (like `TableRequestBuilder2`) use Go generics to provide type-safe responses.
- **Independence**: Each API module (Table, Attachment, Batch) is relatively independent, allowing for modular growth.

## Internal vs. Core

- **`core/`**: Contains the foundational interfaces and common types used throughout the SDK.
- **`internal/`**: Contains helper logic and Kiota-specific wrappers that are not intended for public use.
- **`internal/new/`**: Specifically houses the base request builder and model logic that aligns with Kiota's generator output.
