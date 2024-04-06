package tableapi

type requestBuilder interface {
}

type TableItemRequestBuilder2 struct {
    requestBuilder
}

func NewTableItemRequestBuilder2(client core.Client, pathParameters map[string]string) *TableItemRequestBuilder2 {
   return &TableItemRequestBuilder2{
        core.NewRequestBuilder(client, pathParameters, )
   }
}
