package attachmentapi

import (
	"errors"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
)

type attachmentRequestBuilder2Config struct {
	pathParameters map[string]string
}

type AttachmentRequestBuilder2Option = newInternal.Option[*attachmentRequestBuilder2Config]

func errorOption[T any](err error) newInternal.Option[T] {
	return func(_ T) error {
		return err
	}
}

func WithBaseURL(baseURL string) AttachmentRequestBuilder2Option {
	return WithPathParameters(map[string]string{newInternal.BaseURLkey: baseURL})
}

func WithRawURL(rawURL string) AttachmentRequestBuilder2Option {
	return WithPathParameters(map[string]string{newInternal.RawURLKey: rawURL})
}

func WithPathParameters(pathParameters map[string]string) AttachmentRequestBuilder2Option {
	if len(pathParameters) == 0 {
		return errorOption[*attachmentRequestBuilder2Config](errors.New("pathParameters is empty"))
	}

	return func(config *attachmentRequestBuilder2Config) error {
		if newInternal.IsNil(config) {
			return errors.New("nil config")
		}
		config.pathParameters = pathParameters
		return nil
	}
}
