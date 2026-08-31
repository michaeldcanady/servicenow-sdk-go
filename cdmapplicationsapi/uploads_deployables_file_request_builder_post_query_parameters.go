// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

// UploadsDeployablesFileRequestBuilderPostQueryParameters represents the POST query parameters for the Uploads Deployables File resource.
type UploadsDeployablesFileRequestBuilderPostQueryParameters struct {
	AppName        *string `uriparametername:"appName"`
	DeployableName *string `uriparametername:"deployableName"`
}
