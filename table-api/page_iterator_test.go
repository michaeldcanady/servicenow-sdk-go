package tableapi

import (
	"net/http"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

// Mock client for testing
type mockClient struct{}

func (c *mockClient) Send(requestInformation core.IRequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {
	// Implement mock Send method for testing
	return nil, nil
}

func TestNewPageIteratorWithClient(t *testing.T) {
	// Mock client and current page
	client := &mockClient{}
	currentPage := TableCollectionResponse{
		// Initialize with test data
	}

	pageIterator, err := NewPageIterator(currentPage, client)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if pageIterator == nil {
		t.Error("Expected PageIterator, but got nil")
	}
}

func TestNewPageIteratorWithoutClient(t *testing.T) {
	// Mock current page
	currentPage := TableCollectionResponse{
		// Initialize with test data
	}

	pageIterator, err := NewPageIterator(currentPage, nil)

	assert.Equal(t, (*PageIterator)(nil), pageIterator)
	assert.Equal(t, ErrNilClient, err)
}

func TestPageIteratorNext(t *testing.T) {
	//To Be Added
}

func TestPageIteratorFetchNextPage(t *testing.T) {
	//To Be Added
}

func TestPageIteratorEnumerate(t *testing.T) {
	//To Be Added
}

func TestIterateWithNoCallback(t *testing.T) {
	// Mock PageIterator
	pageIterator := &PageIterator{
		currentPage: PageResult{
			// Initialize with test data
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	err := pageIterator.Iterate(nil)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestIterateWithCallback(t *testing.T) {
	// Mock PageIterator
	pageIterator := &PageIterator{
		currentPage: PageResult{
			// Initialize with test data
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	// Mock callback function
	callback := func(pageItem *TableEntry) bool {
		// Implement your callback logic for testing
		return true
	}

	err := pageIterator.Iterate(callback)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestIterateWithDefaultCallback(t *testing.T) {
	// Mock PageIterator
	pageIterator := &PageIterator{
		currentPage: PageResult{
			// Initialize with test data
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	err := pageIterator.Iterate(defaultCallback)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
