# Authentication

The ServiceNow SDK for Go supports multiple authentication methods. All authentication providers are located in the `github.com/michaeldcanady/servicenow-sdk-go/credentials` package.

## Basic Authentication

Basic authentication is the simplest way to get started. It requires a username and password.

```go
import (
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// Create a username/password credential
cred := credentials.NewUsernamePasswordCredential("username", "password")
```

## OAuth2 Authentication

For more secure applications, you can use OAuth2. The SDK supports the password grant type.

```go
import (
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// Create a token credential
// You'll need your Client ID, Client Secret, and the Base URL of your instance.
cred, err := credentials.NewTokenCredential(
    "your-client-id",
    "your-client-secret",
    "https://your-instance.service-now.com",
    nil, // Use default prompt for username/password or provide your own
)
```

## Custom Credentials

You can implement your own credential provider by satisfying the `core.Credential` interface.

```go
type Credential interface {
	GetAuthentication() (string, error)
}
```

This allows you to integrate with external secret managers or custom authentication flows.
