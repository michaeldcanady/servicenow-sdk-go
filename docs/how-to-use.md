# How to use
This SDK has two modalities of usage: `fluent` and `standard`.

The following block is the base you'll need for **all** implementation methods:
```golang
import (
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

cred := credentials.NewUsernamePasswordCredential("username", "password")

client := servicenowsdkgo.NewServiceNowClient2(cred, "instance")
```

## Fluent

The fluent implementation is designed for ease of use and more simplistic implementations. 

```golang
client.Now().Table("table_name")
client.Now().Attachment()
client.Now().Batch()
```

# Standard