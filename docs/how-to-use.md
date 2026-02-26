# How to use

This SDK has two modalities of usage: `fluent` and `standard`.

The following block is the base you'll need for **all** implementation methods:
```go
import (
    {% include-markdown 'snippets/auth.go' start='// [START credentials_import]' end='// [END credentials_import]' comments=false trailing-newlines=false dedent=true %}
)

{% include-markdown 'snippets/auth.go' start='// [START auth_basic]' end='// [END auth_basic]' comments=false trailing-newlines=false dedent=true %}
{% include-markdown 'snippets/auth.go' start='// [START client_init]' end='// [END client_init]' comments=false trailing-newlines=false dedent=true %}
```

## Fluent

By design, the fluent implementation helps you write cleaner, simpler code with minimal effort.

``` go {title="Table API"}
{% include-markdown 'snippets/howtouse.go' start='// [START fluent_table]' end='// [START fluent_table]' comments=false trailing-newlines=false dedent=true %}
```

``` go {title="Attachment API"}
{% include-markdown 'snippets/howtouse.go' start='// [START fluent_attachment]' end='// [START fluent_attachment]' comments=false trailing-newlines=false dedent=true %}
```

``` go {title="Batch API"}
{% include-markdown 'snippets/howtouse.go' start='// [START fluent_batch]' end='// [START fluent_batch]' comments=false trailing-newlines=false dedent=true %}
```

# Standard