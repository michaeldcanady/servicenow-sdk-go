package documentsapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// SyncDownRequestBuilderPostRequestConfiguration ...
type SyncDownRequestBuilderPostRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
	Data    Document
}
