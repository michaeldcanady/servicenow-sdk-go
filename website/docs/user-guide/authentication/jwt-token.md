# JWT token authentication

JWT Token authentication lets the SDK authenticate using a signed JSON Web
Token (JWT). ServiceNow validates the signature using a public key configured in
your instance and issues an access token. This method is ideal for secure,
non‑interactive server‑to‑server integrations.

## Objective

Configure and use JWT Token authentication with the Service‑Now SDK using values
provided by your ServiceNow administrator.

## Required values

Your administrator must provide:

| Value           | Description                                   |
| --------------- | --------------------------------------------- |
| Service‑Now URL | Base URL of the instance                      |
| Client ID       | From a ServiceNow OAuth or JWT registry entry |
| Client Secret   | Required by ServiceNow for JWT Bearer flows   |

Your application must also provide:

- **Private key** used to sign the JWT assertion
- A **token provider** capable of generating signed JWT assertions

## SDK flow


```mermaid
flowchart TD
    A[App Code] --> B[NewJWTProvider]
    B --> C[newConfidentialClient]
    C --> D[NewJWTCredential]

    %% Initial Token Acquisition
    D --> E[retrieveInitialToken]
    E --> F[tokenProvider.GetAuthorizationToken<br/>generate signed JWT assertion]
    F --> G[validateJWT<br/>claims, iat, alg]
    G --> H[client.acquireTokenByJWT<br/>exchange assertion]
    H --> I[Access Token]
    I --> J[Cache Token]

    %% Refresh Behavior
    J --> K{Token Expired?}
    K -->|Yes| E
    K -->|No| L[Kiota Request Adapter]

    %% Request Pipeline
    L --> M[HTTP Request to ServiceNow]
```

## Initialize the SDK

```go
import (
    "log"

    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

func main() {
    // tokenProvider must generate signed JWT assertions; it is user-provided
    // and must satisfy kiota's authentication.AccessTokenProvider.
    tokenProvider := myJWTAssertionProvider()

    cred, err := credentials.NewJWTProvider(
        "{clientID}",
        "{clientSecret}",
        tokenProvider,
        credentials.WithInstance("{instance}"),
    )
    if err != nil {
        log.Fatal(err)
    }

    client, err := servicenowsdkgo.NewServiceNowServiceClient(
        servicenowsdkgo.WithAuthenticationProvider(cred),
        servicenowsdkgo.WithInstance("{instance}"),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Client is now authenticated and ready to use
    _ = client
}
```
