# Documents API overview

The Documents API manages ServiceNow Document Management: create and link
documents, explore folders, stream document content, and manage versions.

## Basic usage

Access it through the `Now()` namespace:

```go
documents := client.Now().Documents()

// Search documents and folders
results, err := documents.Explore().Get(context.Background(), nil)

// Download a document's content
content, err := documents.Content("{documentSysID}").Get(context.Background(), nil)

// List a document's versions
versions, err := documents.Versions("{documentSysID}").Get(context.Background(), nil)
```

## Available operations

- **Explore** — `Explore().Get(ctx, config)` searches documents and folders.
- **Create** — `Create().Post(ctx, config)` and `CreateDocument().Post(ctx, config)` create or link documents.
- **Content** — `Content(documentSysID).Get(ctx, config)` streams the raw content (`[]byte`).
- **Versions** — `Versions(documentSysID).Get(ctx, config)` lists versions; `VersionState(versionSysID).Get` reads a version's state.
- **Attach** — `Attach(providerID).Post(ctx, config)` attaches content to a document.
- **Actions** — `Action(action).Document(docSysID).Version(versionSysID).Patch(ctx, config)` executes a document action.
- **Sync down** — `SyncDown(documentSysID).Post(ctx, config)`.
- **Delete** — `Delete().Delete(ctx, config)` deletes by query.
