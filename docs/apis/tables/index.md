# Table API Overview

The Table API enables you to perform Create, Read, Update, and Delete (CRUD) operations on ServiceNow tables. Each endpoint corresponds to a specific operation and supports both fluent and standard usage patterns.

Commonly used for:

- Retrieving records from tables like `incident`, `change_request`, or `sys_user`.
- Creating new records with specific field values.
- Updating existing records by their unique system ID (`sys_id`).
- Deleting records when they're no longer needed.

## Basic Usage

To interact with a table, you start from the `ServiceNowClient`, access the `now` namespace, and then specify the table name using `TableV2`.

```go
client, _ := servicenowsdkgo.NewServiceNowClient2(credential, instance)

// Access the incident table
incidentTable := client.Now2().TableV2("incident")

// Fetch a list of records
result, err := incidentTable.Get(context.Background(), nil)
```

## Available Operations

The Table API supports the following standard REST operations:

- [**List Records**](list.md): Retrieve multiple records from a table with support for filtering, sorting, and pagination.
- [**Get Record**](get.md): Retrieve a single record by its `sys_id`.
- [**Create Record**](create.md): Insert a new record into a table.
- [**Update Record**](update.md): Modify an existing record by its `sys_id`.
- [**Delete Record**](delete.md): Remove a record from a table by its `sys_id`.
