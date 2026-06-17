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
