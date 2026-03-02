# Attachment API overview

The Attachment API lets you manage files attached to records in ServiceNow. You can use it to list metadata for multiple attachments, retrieve metadata for a specific attachment, or delete an attachment.

Additionally, it provides specialized endpoints for handling the raw file content (binary data) through the [File operations](file/index.md) sub-module.

Commonly used for:
- Retrieving a list of attachments associated with a specific record or table.
- Downloading file content for processing.
- Uploading documents, images, or logs to existing ServiceNow records.
- Cleaning up old or redundant attachments.

## Basic usage

To interact with attachments, you use the `Attachment2` method from the `now` namespace.

```go
client, _ := servicenowsdkgo.NewServiceNowClient2(credential, instance)

// Access the attachment API
attachmentAPI := client.Now2().Attachment2()

// Fetch metadata for all attachments (optionally filtered)
result, err := attachmentAPI.Get(context.Background(), nil)
```

## Available operations

- [**List Attachments**](list.md): Retrieve metadata for multiple attachments.
- [**Get Attachment Metadata**](get.md): Retrieve metadata for a specific attachment using its `sys_id`.
- [**Delete Attachment**](delete.md): Permanently remove an attachment from the system.
- [**File Operations**](file/index.md): Dedicated endpoints for uploading and downloading the actual file content.
- [**Upload Attachment**](upload/create.md): Upload a new file and associate it with a record using `multipart/form-data`.
