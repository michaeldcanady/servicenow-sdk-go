package documentsapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// CreateRequestBuilderPostRequestConfiguration ...
type CreateRequestBuilderPostRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
	Data    Document
}
