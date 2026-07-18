# CDM Applications API overview

The CDM Applications API manages Configuration Data Management (DevOps Config)
applications: deployables, shared components and libraries, uploads, and
configuration exports.

## Basic usage

Access it through the `Cdm()` namespace:

```go
applications := client.Cdm().Applications()

// Upload component configuration data
status, err := applications.Uploads().Components().Post(context.Background(), uploadRequest, nil)

// Check an upload's status
uploadStatus, err := applications.UploadStatus().ByID("{uploadID}").Get(context.Background(), nil)

// Export a deployable's configuration and fetch the content
export, err := applications.Deployables().Exports().Get(context.Background(), nil)
content, err := applications.Deployables().Exports().ByID("{exportID}").Content().Get(context.Background(), nil)
```

## Available operations

- **Deployables** — `Deployables().Put` / `.Delete` update or remove deployables; `Deployables().Exports()` manages exports (`Get`, `ByID(id).Status().Get`, `ByID(id).Content().Get`).
- **Shared components** — `SharedComponents().Put` / `.Delete`.
- **Shared libraries** — `SharedLibraries().Components().Applications().Get`.
- **Uploads** — `Uploads().Components()` / `.Collections()` / `.Deployables()` POST configuration data.
- **Upload status** — `UploadStatus().ByID(uploadID).Get`.
