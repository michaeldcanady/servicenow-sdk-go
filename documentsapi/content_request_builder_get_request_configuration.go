// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package documentsapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// ContentRequestBuilderGetRequestConfiguration ...
type ContentRequestBuilderGetRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
}
