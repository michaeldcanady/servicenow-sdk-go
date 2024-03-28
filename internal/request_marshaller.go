package internal

import (
	"context"
	"net/http"
)

func ConvertToNativeRequest(ctx context.Context, request RequestInformation) (*http.Request, error) {
	return request.ToRequestWithContext(ctx)
}
