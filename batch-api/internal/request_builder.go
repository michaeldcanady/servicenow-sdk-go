package internal

import (
	"context"
	"github.com/RecoLabs/servicenow-sdk-go/core"
)

type RequestBuilder interface {
	SendPost3(ctx context.Context, xconfig *core.RequestConfiguration) error
}
