package oauth2

// DeviceAuthorizationResponse represents the response from the device authorization endpoint (RFC 8628).
type DeviceAuthorizationResponse struct {
	// DeviceCode is the short-lived code used by the client to request an access token.
	DeviceCode string `json:"device_code"`
	// UserCode is the code displayed to the user for verification on another device.
	UserCode string `json:"user_code"`
	// VerificationURI is the URL the user should visit to enter the user code.
	VerificationURI string `json:"verification_uri"`
	// VerificationURIComplete is an optional URL that includes the user code for convenience.
	VerificationURIComplete string `json:"verification_uri_complete,omitempty"`
	// ExpiresIn is the lifetime of the device and user codes in seconds.
	ExpiresIn int `json:"expires_in"`
	// Interval is the minimum amount of time in seconds the client should wait between polling requests.
	Interval int `json:"interval"`
}
