package tableapi

import (
	"net/http"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/RecoLabs/servicenow-sdk-go/internal"
)

// TablePageIterator is a generic struct in Golang which is used to iterate over pages of entries.
// It embeds the core.PageIterator2 struct and can be used with any type that satisfies the Entry interface.
//
// Fields:
//   - core.PageIterator2[T]: This is an embedded field of type PageIterator2 from the core package.
//     The type T represents the type of Entry that the iterator will return.
//
// Usage:
// You can use this struct to iterate over pages of entries in a table.
// The specific type of Entry depends on what you specify for T when you create an instance of TablePageIterator.
type TablePageIterator[T Entry] struct {
	*core.PageIterator2[T]
}

// constructTableCollection is a generic function in Golang used to construct a collection of table entries.
// It takes an http.Response pointer as input and returns a CollectionResponse of the specified Entry type and an error.
//
// Parameters:
// * response: This is a pointer to an http.Response that contains the server's response to an HTTP request.
//
// Returns:
// * core.CollectionResponse[T]: This is a CollectionResponse of the specified Entry type. It represents the collection of table entries.
// * error: This is an error that will be returned if there is any error while parsing the response.
//
// Usage:
// You can use this function to construct a collection of table entries from an *http.Response.
//
//	func ExampleConstructTableCollectionF() {
//		response, err := http.Get("http://example.com")
//
//		if err != nil {
//	    	log.Fatal(err)
//		}
//
//		collection, err := constructTableCollection[MyEntryType](response)
//
//		if err != nil {
//	    	log.Fatal(err)
//		}
//
//		// Process collection
//	}
func constructTableCollection[T Entry](response *http.Response) (core.CollectionResponse[T], error) {
	resp := &TableCollectionResponse2[T]{}

	err := internal.ParseResponse(response, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NewTablePageIterator is a function that creates a new instance of TablePageIterator.
// It takes a TableCollectionResponse and a Client as input and returns a pointer to a TablePageIterator and an error.
//
// Parameters:
//
//   - collection: This is a pointer to a TableCollectionResponse that contains the collection of table entries.
//   - client: This is a Client that will be used to make requests to the server.
//
// Returns:
//
//   - *TablePageIterator[T]: This is a pointer to a TablePageIterator that can be used to iterate over the pages of entries in the collection.
//   - error: This is an error that will be returned if there is any error while creating the PageIterator.
//
// Usage:
// You can use this function to create a new instance of TablePageIterator.
//
//		func ExampleNewTablePageIteratorF() {
//			// use an existing collection
//			var collection TableCollectionResponse2[T]
//			// use a new or existing client
//			var client core.Client
//			iter, err := NewTablePageIterator[T](collection, client)
//			if err != nil {
//	    		log.Fatal(err)
//			}
//			// Use iter to iterate over the pages of entries
//		}
func NewTablePageIterator[T Entry](collection *TableCollectionResponse2[T], client core.Client) (*TablePageIterator[T], error) {
	pageIterator, err := core.NewPageIterator2[T](collection, client, constructTableCollection[T])
	if err != nil {
		return nil, err
	}

	return &TablePageIterator[T]{
		pageIterator,
	}, nil
}
