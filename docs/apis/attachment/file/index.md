# Attachment file overview

The Attachment File API lets you interact with the raw binary content of file attachments. Unlike the base Attachment API which handles metadata (like filenames and sizes), the File sub-module handles the actual upload and download of file data.

Commonly used for:
- Uploading a new file and attaching it to a specific record.
- Downloading the content of an existing attachment by its `sys_id`.

## Basic usage

### Uploading a file

To upload a file, you specify the target table and record ID, along with the desired filename and the raw data.

```go
data := []byte("this is example data")
media := attachmentapi.NewMedia("text/plain", data)

config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
    QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
        TableSysID: "YOUR_RECORD_SYS_ID",
        TableName:  "incident",
        FileName:   "logs.txt",
    },
}

result, err := client.Now2().Attachment2().File().Post(context.Background(), media, config)
```

### Downloading a file

To download a file, you use the `ByID` method with the attachment's `sys_id` and then access the `File()` sub-module.

```go
downloadedFile, err := client.Now2().Attachment2().ByID("ATTACHMENT_SYS_ID").File().Get(context.Background(), nil)

content, _ := downloadedFile.GetContent()
// content is a []byte containing the file data
```

## Available operations

- [**Upload File**](create.md): Upload raw binary data as an attachment.
- [**Download File**](get.md): Retrieve the raw binary data of an attachment.
