# Resource Owner Password Credential

## Overview

## Example
> [!CAUTION]
> This example has username/password hardcoded for simplicity sake. This is not recommended as it can expose sensitive information!

```go
    import (
    	"context"
    	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
    )

    const (
        ClientID     = "{client id here}"
        ClientSecret = "{client secret here}"
        baseURL      = "{base service-now url here}"
        username     = "{username here}"
        password     = "{password here}"
    )

    userPassProvider := func(ctx context.Context) (string, string, error) {
		return username, password, nil
	}

    cred, err := credentials.NewResourceOwnerPasswordCredential(ClientID, ClientSecret, baseURL, userPassProvider)
```
