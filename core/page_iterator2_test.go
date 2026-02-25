package core

import (
	"net/http"
	"testing"
)

func TestNewPageIterator2(t *testing.T) {
	client := &mockCoreClient2{}
	page := &mockPage{}
	cf := func(r *http.Response) (CollectionResponse[any], error) { return page, nil }
	tests := []struct {
		name   string
		page   *mockPage
		client Client
		err    error
	}{
		{"Ok", page, client, nil},
		{"NilClient", page, nil, ErrNilClient},
		{"NilPage", nil, client, ErrNilResponse},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewPageIterator2[any](tt.page, tt.client, cf)
			if err != tt.err {
				t.Errorf("got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestPageIterator2_Iterate(t *testing.T) {
	tests := []struct {
		name     string
		iterator *PageIterator2[any]
		callback func(*any) bool
		err      error
	}{
		{"NilCallback", &PageIterator2[any]{}, nil, ErrNilCallback},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.iterator.Iterate(tt.callback, false)
			if err != tt.err {
				t.Errorf("got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestPageIterator2_nextPage(t *testing.T) {
	it := &PageIterator2[any]{}
	tests := []struct {
		name string
		err  bool
	}{
		{"EmptyLink", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := it.nextPage(false)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v", err)
			}
		})
	}
}
