// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import "errors"

// errNetwork is a stand-in for a transport-level error returned by the adapter.
var errNetwork = errors.New("network error")
