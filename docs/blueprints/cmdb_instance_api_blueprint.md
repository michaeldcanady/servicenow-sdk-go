# Technical Blueprint: CMDB Instance API Support (`cmdbinstanceapi`)

This blueprint outlines the architectural plan for implementing the ServiceNow CMDB Instance API (v1) in the Go SDK, derived from `spec/now_cmdb_instance_api_v1_spec.json`.

## 1. Package Overview
- **Package Name**: `cmdbinstanceapi`
- **Base Path**: `/api/now/v1/cmdb/instance`
- **Internal Alias**: `newInternal` (`github.com/michaeldcanady/servicenow-sdk-go/internal/new`)

## 2. RequestBuilder Hierarchy

| Path | RequestBuilder | Method | Description |
| :--- | :--- | :--- | :--- |
| `/api/now/v1/cmdb/instance` | `CmdbInstanceRequestBuilder` | - | Base entry point |
| `/api/now/v1/cmdb/instance/{className}` | `CmdbClassRequestBuilder` | `Get`, `Post` | Operations on a CMDB class |
| `/api/now/v1/cmdb/instance/{className}/{sys_id}` | `CmdbItemRequestBuilder` | `Get`, `Put`, `Patch` | Operations on a specific CI |
| `/api/now/v1/cmdb/instance/{className}/{sys_id}/relation` | `CmdbRelationRequestBuilder` | `Post` | Create relationship for a CI |

## 3. Operations & Models

### 3.1. Class Query Operation
- **Path**: `/api/now/v1/cmdb/instance/{className}`
- **Method**: `Get`
- **Query Parameters**:
    - `sysparm_query` (string)
    - `sysparm_limit` (int)
    - `sysparm_offset` (int)
- **Response**: `ServiceNowCollectionResponse[CmdbInstance]`

### 3.2. Item Operations
- **Path**: `/api/now/v1/cmdb/instance/{className}/{sys_id}`
- **Methods**: `Get`, `Put`, `Patch`
- **Response**: `ServiceNowItemResponse[CmdbInstance]`

### 3.3. Relation Operation
- **Path**: `/api/now/v1/cmdb/instance/{className}/{sys_id}/relation`
- **Method**: `Post`
- **Response**: `ServiceNowItemResponse[CmdbInstance]` (Or specialized relation model)

## 4. Model Definitions (Proposed)
*Note: Spec schemas are empty; using standard ServiceNow CMDB patterns.*

### 4.1. `CmdbInstance` Model
- `SysID` (string)
- `Name` (string)
- `ClassName` (string)
- (Backed by internal backing store for dynamic fields)

## 5. Implementation Strategy
1. **Scaffold Directory**: Create `cmdb-instance-api/`.
2. **Base Builders**:
    - `CmdbInstanceRequestBuilder2` (root)
    - `CmdbClassRequestBuilder` (by class name)
    - `CmdbItemRequestBuilder` (by sys id)
3. **Configurations**: Generate `RequestConfiguration` and `QueryParameters` for all operations.
4. **Validation**: Add unit tests using table-driven patterns in decentralized files.
