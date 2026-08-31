// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

// SharedComponentsRequestBuilderDeleteQueryParameters represents query parameters for DELETE /applications/shared_components.
type SharedComponentsRequestBuilderDeleteQueryParameters struct {
	AppName             *string `uriparametername:"appName"`
	SharedComponentName *string `uriparametername:"sharedComponentName"`
	Name                *string `uriparametername:"name"` // Support name too just in case
}
