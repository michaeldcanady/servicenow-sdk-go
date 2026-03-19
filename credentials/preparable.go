package credentials

// Preparable is an interface for authentication providers or token providers
// that need to be initialized with the base URL from the client.
type Preparable interface {
	Initialize(baseURL string)
}
