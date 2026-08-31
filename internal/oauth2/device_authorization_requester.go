// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package oauth2

import "context"

type DeviceAuthorizationRequester interface {
	// RequestDeviceAuthorization initiates the device authorization flow.
	RequestDeviceAuthorization(ctx context.Context, scopes []string) (*DeviceAuthorizationResponse, error)
}
