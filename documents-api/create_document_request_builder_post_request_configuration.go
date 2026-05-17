package documentsapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// CreateDocumentRequestBuilderPostRequestConfiguration ...
type CreateDocumentRequestBuilderPostRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
	Data    Document
}
