package core

import (
	"testing"
)

func TestNewPageIterator(t *testing.T) {
	client := &mockCoreClient{}
	page := &mockPage{}
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
			_, err := NewPageIterator[any, *mockPage](tt.page, tt.client)
			if err != tt.err {
				t.Errorf("got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestPageIterator_Iterate(t *testing.T) {
	tests := []struct {
		name     string
		iterator *PageIterator[any, *mockPage]
		callback func(*any) bool
		err      error
	}{
		{"NilCallback", &PageIterator[any, *mockPage]{}, nil, ErrNilCallback},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.iterator.Iterate(tt.callback)
			if err != tt.err {
				t.Errorf("got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestPageIterator_fetchAndConvertPage(t *testing.T) {
	it := &PageIterator[any, *mockPage]{client: &mockCoreClient{}}
	tests := []struct {
		name string
		link string
		err  bool
	}{
		{"EmptyLink", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := it.fetchAndConvertPage(tt.link)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v", err)
			}
		})
	}
}
