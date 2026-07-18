# CDM Changesets API overview

The CDM Changesets API manages Configuration Data Management changesets:
inspect changeset activity, check commit status, and evaluate the impact of a
changeset before committing it.

## Basic usage

Access it through the `Cdm()` namespace:

```go
changesets := client.Cdm().Changesets()

// List changesets
list, err := changesets.Get(context.Background(), nil)

// Check the status of a commit
status, err := changesets.CommitStatus().ByID("{commitID}").Get(context.Background(), nil)

// Impacted deployables for a specific changeset
impact, err := changesets.ByID("{changesetSysID}").ImpactedDeployables().Get(context.Background(), nil)
```

## Available operations

- **List / delete changesets** — `Changesets().Get(ctx, config)` and `.Delete(ctx, config)`.
- **Activity** — `Activity().Get(ctx, config)`.
- **Commit status** — `CommitStatus().ByID(commitID).Get(ctx, config)`.
- **Impact analysis** — `ImpactedSharedComponents().Get`, `ImpactedDeployables().Get`, and per-changeset `ByID(id).ImpactedDeployables().Get`.
