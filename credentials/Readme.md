# Credentials

The credential module exists to provide various was to authenticate to the ServiceNow APIs.

## Username and Password Credential

`UsernamePasswordCredential` is a `Basic` Credential type, where the SDK authenticates via the provided username and password.

> **NOTE**: the supplied credentials are stored within the credential and the fields are exported.

```golang
import (
    "github.com/RecoLabs/servicenow-sdk-go/credentials"
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
    "github.com/RecoLabs/servicenow-sdk-go/credentials"
)

func main() {
    cred, err := credentials.NewTokenCredential("clientID", "clientSecret", "baseURL", prompt)
    if err != nil {
        panic(err)
    }

    ...
}
```
