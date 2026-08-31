// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package oauth2

import "context"

type DeviceCodeExchanger interface {
	// ExchangeDeviceCode exchanges a device code for an access token.
	ExchangeDeviceCode(ctx context.Context, deviceCode string) (*Token, error)
}
