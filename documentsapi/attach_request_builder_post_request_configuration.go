// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package documentsapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// AttachRequestBuilderPostRequestConfiguration ...
type AttachRequestBuilderPostRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
	Data    Document
}
