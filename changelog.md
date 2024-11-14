# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] - (dd/mm/yyyy)

### Added

- `batchapi` module ([Issue #49](https://github.com/RecoLabs/servicenow-sdk-go/issues/49))
- `tableapi.TableItemRequestBuilder2`
- `tableapi.NewTableItemRequestBuilder2`
- `tableapi.TableRequestBuilder2`
- `tableapi.NewTableRequestBuilder2`
- `servicenowsdk.NowRequestBuilder.Table2`
- `tableapi.DisplayValueAll`
- `tableapi.DisplayValueTrue`
- `tableapi.DisplayValueFalse`
- `tableapi.ViewDesktop`
- `tableapi.ViewMobile`
- `tableapi.ViewBoth`

### Changed

### Deprecated

- `tableapi.TableItemRequestBuilder`
- `tableapi.NewTableItemRequestBuilder`
- `tableapi.TableRequestBuilder`
- `tableapi.NewTableRequestBuilder`
- `servicenowsdk.NowRequestBuilder.Table`
- `tableapi.ALL`
- `tableapi.TRUE`
- `tableapi.FALSE`
- `tableapi.VIEWDESKTOP`
- `tableapi.VIEWMOBILE`
- `tableapi.VIEWBOTH`

### Removed

### Fixed

## [1.6.1] - (09/03/2024)

### Fixed

- `core.RequestInformation.ToRequestWithContext` not including headers when converting to request ([Issue #107](https://github.com/RecoLabs/servicenow-sdk-go/issues/107))

## [1.6.0] - (08/03/2024)

### Added

- `core.PageIterator2[T].Previous` ([Issue #60](https://github.com/RecoLabs/servicenow-sdk-go/issues/60))
- `core.PageIterator2[T].First` ([Issue #60](https://github.com/RecoLabs/servicenow-sdk-go/issues/60))

### Deprecated

- `servicenowsdkgo.ServiceNowClient.Credential`
- `servicenowsdkgo.NewServiceNowClient`

### Fixed

- Improperly embedded `core.PageIterator2[T]` in `tableapi.TablePageIterator` ([Issue #105](https://github.com/RecoLabs/servicenow-sdk-go/issues/105))

## [1.5.0] - (02/03/2024)

### Added

- `tableapi.TablePageIterator[T]` ([Issue #91](https://github.com/RecoLabs/servicenow-sdk-go/issues/91))
- `tableapi.NewTablePageIterator[T]` ([Issue #91](https://github.com/RecoLabs/servicenow-sdk-go/issues/91))
- `core.PageIterator2[T]` ([Issue #58](https://github.com/RecoLabs/servicenow-sdk-go/issues/58))
- `core.NewPageIterator2[T]` ([Issue #58](https://github.com/RecoLabs/servicenow-sdk-go/issues/58))

### Deprecated

- `tableapi.NewPageIterator` ([Issue #91](https://github.com/RecoLabs/servicenow-sdk-go/issues/91))
- `tableapi.PageIterator` ([Issue #91](https://github.com/RecoLabs/servicenow-sdk-go/issues/91))
- `core.PageIterator[T]` ([Issue #58](https://github.com/RecoLabs/servicenow-sdk-go/issues/58))
- `core.NewPageIterator[T]` ([Issue #58](https://github.com/RecoLabs/servicenow-sdk-go/issues/58))

### Fixed

- Incorrect marshalling of query parameters ([Issue #68](https://github.com/RecoLabs/servicenow-sdk-go/issues/91))
- `ToRequestInformation3` not reporting values as nil when nil

## [1.4.0] - (02/01/2024)

### Added

- Added `PageIterator.Last` ([Issue #61](https://github.com/RecoLabs/servicenow-sdk-go/issues/61))
- `tableGetRequestConfiguration2[T]` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `tableItemDeleteRequestConfiguration2[T]` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `tableItemGetRequestConfiguration2[T]` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `tableItemPutRequestConfiguration2[T]` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `TableItemResponse2[T]` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `tablePostRequestConfiguration2[T]` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `TableRequestBuilder.Post3` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `TableItemRequestBuilder.Put2` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))

### Changed

- abstract `core.PageIterator` and `core.PageResult`
- `PageIterator.fetchPage` does not parse provided uri ([Issue #83](https://github.com/RecoLabs/servicenow-sdk-go/issues/83))

### Deprecated

- `TableGetRequestConfiguration` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `TableItemDeleteRequestConfiguration` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `TableItemGetRequestConfiguration` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `TableItemPutRequestConfiguration` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `TableItemResponse` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `TablePostRequestConfiguration` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `TableRequestBuilder.Post2` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))
- `TableItemRequestBuilder.Put` ([Issue #81](https://github.com/RecoLabs/servicenow-sdk-go/issues/81))

## [1.3.1] - (16/12/2023)

### Fixed

- `TableRequestBuilder.Post`/`TableRequestBuilder.Post2` sending `PUT` request ([Issue #69](https://github.com/RecoLabs/servicenow-sdk-go/issues/69))

## [1.2.2] - (15/12/2023)

### Fixed

- Table Iteration ends after the second iteration ([Issue #70](https://github.com/RecoLabs/servicenow-sdk-go/issues/70))

## [1.2.1] - 10/12/2023

### Added

- Added AttachmentRequestBuilder.File method ([Issue #52](https://github.com/RecoLabs/servicenow-sdk-go/pull/52))

### Fixed

- Page Interation stops after first page ([Issue #62](https://github.com/RecoLabs/servicenow-sdk-go/pull/62))
