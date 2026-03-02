# Attachments

The Attachment API lets you manage files associated with records in ServiceNow.
You can list, retrieve, upload, and delete attachments using the SDK.

## List attachments

You can retrieve a list of all attachments or filter them based on criteria
such as the table name or table `sys_id`.

```go
{% include-markdown 'snippets/attachments.go' start='// [START attachment_list_guide]' end='// [END attachment_list_guide]' comments=false trailing-newlines=false dedent=true %}
```

## Upload attachments

To upload a file and associate it with a specific record, use the `File`
resource's `Post` method. You must provide the table name, the record's
`sys_id`, and the file name in the query parameters.

```go
{% include-markdown 'snippets/attachments.go' start='// [START attachment_create_guide]' end='// [END attachment_create_guide]' comments=false trailing-newlines=false dedent=true %}
```

## Download attachments

To download the content of an attachment, use the `ByID` and `File` resources
to call the `Get` method.

```go
{% include-markdown 'snippets/attachments.go' start='// [START attachment_download_guide]' end='// [END attachment_download_guide]' comments=false trailing-newlines=false dedent=true %}
```

## Delete attachments

To delete an attachment, use the `ByID` resource and call the `Delete` method.

```go
{% include-markdown 'snippets/attachments.go' start='// [START attachment_delete]' end='// [END attachment_delete]' comments=false trailing-newlines=false dedent=true %}
```

## Next steps

- **[Batch API](batch.md):** Learn how to perform multiple operations, including
  attachment management, in a single request.
- **[Table Operations](tables.md):** Learn more about managing the records these
  attachments belong to.
