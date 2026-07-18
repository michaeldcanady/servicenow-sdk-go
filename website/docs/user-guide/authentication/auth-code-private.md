# Authorization code (private)

The Authorization Code (private) flow works for applications that can
securely store a client secret (for example, server-side services, desktop apps). A user
authenticates interactively, and the SDK exchanges the authorization code for
tokens.

## Objective

Configure and use the Authorization Code (private) OAuth flow with the
Service‑Now SDK using values provided by your ServiceNow administrator.

## Required values

Your administrator must provide:

| Value             | Description                                          |
| ----------------- | ---------------------------------------------------- |
| Service‑Now URL   | Base URL of the instance                             |
| Client ID         | From a ServiceNow OAuth application registry entry   |
| Client Secret     | From the same registry entry                         |
| Redirect URL      | Must match the redirect URL configured in ServiceNow |
| Authorization URL | OAuth authorization endpoint                         |
| Token URL         | OAuth token endpoint                                 |

## SDK flow

```mermaid
flowchart TD
    A[App Code] --> B[NewAuthorizationCodeAuthenticationProvider]
    B --> C{clientSecret provided?}
    C -->|Yes| D[newConfidentialClient]
    D --> E[AuthorizationCodeCredential]

    %% Initial Token Acquisition
    E --> F[retrieveInitialToken]
    F --> G[Generate state UUID]
    G --> H[Start local redirect server on port 5001]
    H --> I[client.getAuthorizationURL]
    I --> J[Open Browser to Authorization URL]
    J --> K[Wait for redirect to local server]
    K --> L[Extract Authorization Code]
    L --> M[client.acquireTokenByCode]
    M --> N[Access Token + Refresh Token]
    N --> O[Cache Token]

    %% Refresh Flow
    O --> P{Token Expired?}
    P -->|Yes| Q[client.acquireTokenByRefreshToken]
    Q --> O

    %% Request Pipeline
    P -->|No| R[Kiota Request Adapter]
    R --> S[HTTP Request to ServiceNow]
```

## Initialize the SDK

```golang
import (
    "log"

    credentials "github.com/michaeldcanady/service-now-sdk/credentials"
    servicenow "github.com/michaeldcanady/service-now-sdk"
)

func main() {
    authority := credentials.NewInstanceAuthority("{instance}")

    cred, err := credentials.NewAuthorizationCodeAuthenticationProvider(
        clientID,
        clientSecret,
        authority,
        []string{string(authority)},
    )
    if err != nil {
        log.Fatal(err)
    }

    clientOpts := []credentials.ServiceNowServiceClientOption{
        servicenow.WithAuthenticationProvider(cred),
        servicenow.WithInstance("{instance}"),
    }

    client, err := servicenow.NewServiceNowServiceClient(clientOpts...)
    if err != nil {
        log.Fatal(err)
    }
}
```
