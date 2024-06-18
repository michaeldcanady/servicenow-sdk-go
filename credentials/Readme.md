# Credentials

The credential module exists to provide various was to authenticate to the ServiceNow APIs.

## Username and Password Credential

`UsernamePasswordCredential` is a `Basic` Credential type, where the SDK authenticates via the provided username and password.

> **NOTE**: the supplied credentials are stored within the credential and the fields are exported.

```golang
import (
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

func main() {
    cred, err := credentials.NewUsernamePasswordCredential("username", "password")
    if err != nil {
        panic(err)
    }

    ...
}
```

## Token Credential

`TokenCredential` is a `Token` Credential type, where the SDK authenticates via a token retrieved using the provided username and password.

```golang
import (
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

func main() {
    cred, err := credentials.NewTokenCredential("clientID", "clientSecret", "baseURL", prompt)
    if err != nil {
        panic(err)
    }

    ...
}
```

## JWT Token Credential

ServiceNow also supports using [OAuth JWT API}(https://docs.servicenow.com/bundle/washingtondc-platform-security/page/administer/security/task/create-jwt-endpoint.html) to authenticate with your instance seamlessly using the inbound JWT grant type instead of username and password. It uses the same `TokenCredential`, but the SDK authenticates via a token retrieved using the provided JWT assertion. An RSA private key is needed to sign the token so ServiceNow can validate it (see link above).

```golang
import (
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

func main() {
    cred, err := credentials.NewJwtTokenCredential("clientID", "clientSecret", "kid", "baseURL", "user", "privateKeyPath")
    if err != nil {
        panic(err)
    }

    ...
}
```
