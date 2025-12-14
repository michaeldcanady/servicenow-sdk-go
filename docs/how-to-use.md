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

By design, the fluent implementation helps you write cleaner, simpler code with minimal effort.

``` golang {title="Table API"}
client.Now().Table("table_name")
```

``` golang {title="Attachment API"}
client.Now().Attachment2()
```

``` golang {title="Batch API"}
client.Now().Batch()
```

# Standard