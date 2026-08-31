// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

// UploadsCollectionsFileRequestBuilderPostQueryParameters represents the POST query parameters for the Uploads Collections File resource.
type UploadsCollectionsFileRequestBuilderPostQueryParameters struct {
	AppName        *string `uriparametername:"appName"`
	CollectionName *string `uriparametername:"collectionName"`
}
