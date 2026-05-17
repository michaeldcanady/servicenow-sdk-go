package documentsapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// VersionActionRequestBuilderPatchRequestConfiguration ...
type VersionActionRequestBuilderPatchRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
	Data    Document
}
