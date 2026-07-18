//go:build snippets

package snippets

// [START config_imports]
import (
	"log"
	"net/http"
	"time"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

// [END config_imports]

import "github.com/microsoft/kiota-abstractions-go/authentication"

// [START config_logging_middleware]
type myLoggingMiddleware struct{}

func (m *myLoggingMiddleware) Intercept(
	pipeline nethttplibrary.Pipeline, middlewareIndex int, req *http.Request,
) (*http.Response, error) {
	start := time.Now()
	resp, err := pipeline.Next(req, middlewareIndex)
	log.Printf("%s %s -> %v (%s)", req.Method, req.URL.Path, err, time.Since(start))
	return resp, err
}

// [END config_logging_middleware]

// [START config_std_logger]
type stdLogger struct{}

func (stdLogger) Log(message string, args ...interface{}) {
	log.Printf(message, args...)
}

// [END config_std_logger]

func _() {
	var cred authentication.AuthenticationProvider

	// [START config_target]
	servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx")
	// or, for a full URL (e.g. behind a proxy):
	servicenowsdkgo.WithURL("xSDK_SN_URLx")
	// [END config_target]

	// [START config_middleware]
	middleware := append(
		nethttplibrary.GetDefaultMiddlewares(), // keep retries, redirects, compression
		&myLoggingMiddleware{},
	)

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
		servicenowsdkgo.WithMiddleware(middleware...),
	)
	// [END config_middleware]
	if err != nil {
		log.Fatal(err)
	}

	// [START config_http_client]
	servicenowsdkgo.WithHTTPClient(&http.Client{Timeout: 30 * time.Second})
	// [END config_http_client]

	// [START config_logger]
	servicenowsdkgo.WithLogger(stdLogger{})
	// [END config_logger]

	_ = client
}

func _() {
	// [START config_full]
	// Step 1: Credentials (see the Authentication guide for OAuth2 flows)
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")

	// Step 2: Extend the default middleware chain — keep the built-in
	// retries, redirects, and compression, and add your own handler
	middleware := append(
		nethttplibrary.GetDefaultMiddlewares(),
		&myLoggingMiddleware{},
	)

	// Step 3: Compose the client from options
	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
		servicenowsdkgo.WithMiddleware(middleware...),
		servicenowsdkgo.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
		servicenowsdkgo.WithLogger(stdLogger{}),
	)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}

	// The configured client is what every request builder hangs off
	_ = client
	// [END config_full]
}
