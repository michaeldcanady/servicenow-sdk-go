# Attachment upload overview

The Attachment Upload API provides a specialized endpoint for uploading files using `multipart/form-data`. This is often used when you need to send multiple parts or when the client environment specifically supports multipart uploads.

For most standard file uploads, the [File API](../file/index.md) is simpler to use.

Commonly used for:
- Uploading files from web forms or other multipart-compatible sources.
- Associating a file with a record using the multipart format.

## Available operations

- [**Upload Attachment (Multipart)**](create.md): Upload a file to a record in a specified table using `multipart/form-data`.
