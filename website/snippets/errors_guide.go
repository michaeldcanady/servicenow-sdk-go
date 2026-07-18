//go:build snippets

package snippets

// [START errors_imports]
import (
	"context"
	"errors"
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
)

// [END errors_imports]

func errorsGuideStatus() error {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()

	// [START errors_status_match]
	response, err := client.Now().Table("xSDK_SN_TABLEx").ByID("xSDK_SN_TABLE_SYS_IDx").Get(ctx, nil)
	if err != nil {
		var notFound *core.NotFoundError
		var unauthorized *core.UnauthorizedError

		switch {
		case errors.As(err, &notFound):
			// The record does not exist — treat as absent, not fatal.
		case errors.As(err, &unauthorized):
			// Credentials rejected — re-authenticate or fail fast.
		default:
			return err
		}
	}
	// [END errors_status_match]
	_ = response
	return nil
}

func errorsGuideDetails(err error) {
	// [START errors_details]
	var snErr *core.ServiceNowError
	if errors.As(err, &snErr) {
		mainErr, _ := snErr.GetError()
		message, _ := mainErr.GetMessage()
		detail, _ := mainErr.GetDetail()
		status, _ := mainErr.GetStatus()
		log.Printf("ServiceNow error: message=%v detail=%v status=%v",
			message, detail, status)
	}
	// [END errors_details]
}

func errorsGuideSentinels(err error) {
	// [START errors_sentinels]
	if errors.Is(err, snerrors.ErrNilRequestBuilder) {
		// The builder chain was constructed from a nil client.
	}
	// [END errors_sentinels]
}

func _() {
	_ = errorsGuideStatus()
	errorsGuideDetails(nil)
	errorsGuideSentinels(nil)
}
