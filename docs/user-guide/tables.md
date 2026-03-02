# Table operations

The Table API lets you perform Create, Read, Update, and Delete (CRUD)
operations on ServiceNow table records.

## Understand the modalities

The SDK offers two distinct ways to build and execute requests.

- **Fluent:** A method-chaining approach that's expressive and easy to read. That's
  preferred for most common tasks.
- **Standard:** A more explicit approach that requires you to manage request
  builders and raw URLs. This is useful when you need to handle dynamic URLs or
  specific configurations not expressed fluently.

## Get a single record

To retrieve a specific record, use the `ById` method in the fluent interface.

```go
{% include-markdown 'snippets/tables.go' start='// [START table_get_fluent]' end='// [END table_get_fluent]' comments=false trailing-newlines=false dedent=true %}
```

## List records

To list records from a table, call `Get` on the table resource.

```go
{% include-markdown 'snippets/tables.go' start='// [START table_list_guide]' end='// [END table_list_guide]' comments=false trailing-newlines=false dedent=true %}
```

## Create records

To create a new record, use the `Post` method and provide a `TableRecord`
containing the field values.

```go
{% include-markdown 'snippets/tables.go' start='// [START table_create_guide]' end='// [END table_create_guide]' comments=false trailing-newlines=false dedent=true %}
```

## Update records

To update an existing record, use the `Put` method with the record's `sys_id`
and a `TableRecord` containing the updated fields.

```go
{% include-markdown 'snippets/tables.go' start='// [START table_update_guide]' end='// [END table_update_guide]' comments=false trailing-newlines=false dedent=true %}
```

## Delete records

To delete a record, use the `Delete` method with the record's `sys_id`.

```go
{% include-markdown 'snippets/tables.go' start='// [START table_delete_guide]' end='// [END table_delete_guide]' comments=false trailing-newlines=false dedent=true %}
```

## Query records

You can filter results by providing a query string in the request
configuration. ServiceNow uses a specific query syntax (Encoded Queries).

```go
{% include-markdown 'snippets/tables.go' start='// [START table_query_guide]' end='// [END table_query_guide]' comments=false trailing-newlines=false dedent=true %}
```

## Next steps

- **[Pagination](pagination.md):** Learn how to handle large result sets.
- **[Attachments](attachments.md):** Attach files to your table records.
