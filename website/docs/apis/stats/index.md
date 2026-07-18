# Stats API overview

The Stats API returns aggregate statistics (count, sum, average, minimum,
maximum) computed over the records of a table, without returning the records
themselves.

:::note
This module currently supports the **ungrouped** shape of the API only:
one aggregate result per request. `sysparm_group_by` / `sysparm_having`
are not yet modeled.
:::

## Basic usage

Access it through the `Now()` namespace:

```go
config := &statsapi.StatsRequestBuilderGetRequestConfiguration{
    QueryParameters: &statsapi.StatsRequestBuilderGetQueryParameters{
        Count:     true,
        SumFields: []string{"reassignment_count"},
        Query:     "active=true",
    },
}

response, err := client.Now().Stats("incident").Get(context.Background(), config)
if err != nil {
    log.Fatal(err)
}

result, err := response.GetResult()
if err != nil {
    log.Fatal(err)
}

stats, err := result.GetStats()
if err != nil {
    log.Fatal(err)
}

count, _ := stats.GetCount()
```

## Available operations

- **Get aggregates** — `Stats(tableName).Get(ctx, config)` with query parameters
  selecting the aggregates (`Count`, `SumFields`, `AvgFields`, `MinFields`,
  `MaxFields`) and an encoded `Query` to filter the input records.
