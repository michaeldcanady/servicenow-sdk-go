# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [unreleased] - (mm/dd/yyyy)

### Added

### Changed

### Deprecated

### Removed

### Fixed

## [1.4.0] - (02/01/2024)

### Added

- Added `PageIterator.Last` ([Issue #61](https://github.com/michaeldcanady/servicenow-sdk-go/issues/61))
- `tableGetRequestConfiguration2[T]` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `tableItemDeleteRequestConfiguration2[T]` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `tableItemGetRequestConfiguration2[T]` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `tableItemPutRequestConfiguration2[T]` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `TableItemResponse2[T]` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `tablePostRequestConfiguration2[T]` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `TableRequestBuilder.Post3` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `TableItemRequestBuilder.Put2` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))

### Changed

- abstract `core.PageIterator` and `core.PageResult`
- `PageIterator.fetchPage` does not parse provided uri ([Issue #83](https://github.com/michaeldcanady/servicenow-sdk-go/issues/83))

### Deprecated

- `TableGetRequestConfiguration` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `TableItemDeleteRequestConfiguration` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `TableItemGetRequestConfiguration` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `TableItemPutRequestConfiguration` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `TableItemResponse` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `TablePostRequestConfiguration` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `TableRequestBuilder.Post2` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))
- `TableItemRequestBuilder.Put` ([Issue #81](https://github.com/michaeldcanady/servicenow-sdk-go/issues/81))

### Removed

### Fixed

## [1.3.1] - (16/12/2023)

### Added

### Changed

### Removed

### Fixed

- `TableRequestBuilder.Post`/`TableRequestBuilder.Post2` sending `PUT` request ([Issue #69](https://github.com/michaeldcanady/servicenow-sdk-go/issues/69))

## [1.2.2] - (15/12/2023)

### Fixed

- Table Iteration ends after the second iteration ([Issue #70](https://github.com/michaeldcanady/servicenow-sdk-go/issues/70))

## [1.2.1] - 10/12/2023

### Added

- Added AttachmentRequestBuilder.File method ([Issue #52](https://github.com/michaeldcanady/servicenow-sdk-go/pull/52))

### Fixed

- Page Interation stops after first page ([Issue #62](https://github.com/michaeldcanady/servicenow-sdk-go/pull/62))