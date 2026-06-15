package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
)

func TestNewBaseRequestBuilder(t *testing.T) {
	ra := mocking.NewMockRequestAdapter()
	tests := []struct {
		name   string
		ra     *mocking.MockRequestAdapter
		templ  string
		params map[string]string
	}{
		{"Standard", ra, "t", map[string]string{"a": "b"}},
		{"NilParams", ra, "t", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := NewBaseRequestBuilder(tt.ra, tt.templ, tt.params)
			if res == nil {
				t.Fatal("NewBaseRequestBuilder returned nil")
			}
			if res.UrlTemplate != tt.templ {
				t.Errorf("got %s, expected %s", res.UrlTemplate, tt.templ)
			}
		})
	}
}

func TestBaseRequestBuilder_GetPathParameters(t *testing.T) {
	params := map[string]string{"a": "b"}
	rb := &BaseRequestBuilder{}
	rb.PathParameters = params
	var nilRB *BaseRequestBuilder

	tests := []struct {
		name     string
		rb       *BaseRequestBuilder
		expected map[string]string
	}{
		{"Ok", rb, params},
		{"NilRB", nilRB, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.rb.GetPathParameters()
			if tt.expected == nil {
				if res != nil {
					t.Error("expected nil")
				}
			} else {
				if res["a"] != tt.expected["a"] {
					t.Errorf("got %v, expected %v", res, tt.expected)
				}
			}
		})
	}
}

func TestBaseRequestBuilder_SetPathParameters(t *testing.T) {
	params := map[string]string{"a": "b"}
	rb := &BaseRequestBuilder{}
	var nilRB *BaseRequestBuilder

	tests := []struct {
		name   string
		rb     *BaseRequestBuilder
		params map[string]string
		err    error
	}{
		{"Ok", rb, params, nil},
		{"NilParams", rb, nil, errors.New("pathParameters is nil")},
		{"NilRB", nilRB, params, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rb.SetPathParameters(tt.params)
			if tt.err != nil {
				if err == nil || err.Error() != tt.err.Error() {
					t.Errorf("got err %v, expected err %v", err, tt.err)
				}
			} else if err != nil {
				t.Errorf("unexpected err %v", err)
			}
		})
	}
}

func TestBaseRequestBuilder_GetRequestAdapter(t *testing.T) {
	ra := mocking.NewMockRequestAdapter()
	rb := &BaseRequestBuilder{}
	rb.RequestAdapter = ra
	var nilRB *BaseRequestBuilder

	tests := []struct {
		name     string
		rb       *BaseRequestBuilder
		expected any
	}{
		{"Ok", rb, ra},
		{"NilRB", nilRB, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.rb.GetRequestAdapter()
			if tt.expected == nil {
				if res != nil {
					t.Error("expected nil")
				}
			} else if res != tt.expected {
				t.Error("got wrong adapter")
			}
		})
	}
}

func TestBaseRequestBuilder_SetRequestAdapter(t *testing.T) {
	ra := mocking.NewMockRequestAdapter()
	rb := &BaseRequestBuilder{}
	var nilRB *BaseRequestBuilder

	tests := []struct {
		name string
		rb   *BaseRequestBuilder
		ra   *mocking.MockRequestAdapter
		err  error
	}{
		{"Ok", rb, ra, nil},
		{"NilRA", rb, nil, errors.New("requestAdapter is nil")},
		{"NilRB", nilRB, ra, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Type conversion needed because SetRequestAdapter expects abstractions.RequestAdapter
			// Wait, the method signature is abstractions.RequestAdapter, let's just pass it
			err := tt.rb.SetRequestAdapter(tt.ra)
			if tt.err != nil {
				if err == nil || err.Error() != tt.err.Error() {
					t.Errorf("got err %v, expected err %v", err, tt.err)
				}
			} else if err != nil {
				t.Errorf("unexpected err %v", err)
			}
		})
	}
}

func TestBaseRequestBuilder_GetURLTemplate(t *testing.T) {
	rb := &BaseRequestBuilder{}
	rb.UrlTemplate = "t"
	var nilRB *BaseRequestBuilder

	tests := []struct {
		name     string
		rb       *BaseRequestBuilder
		expected string
	}{
		{"Ok", rb, "t"},
		{"NilRB", nilRB, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.rb.GetURLTemplate() != tt.expected {
				t.Errorf("got %s, expected %s", tt.rb.GetURLTemplate(), tt.expected)
			}
		})
	}
}

func TestBaseRequestBuilder_SetURLTemplate(t *testing.T) {
	rb := &BaseRequestBuilder{}
	var nilRB *BaseRequestBuilder

	tests := []struct {
		name string
		rb   *BaseRequestBuilder
		t    string
		err  error
	}{
		{"Ok", rb, "t", nil},
		{"Empty", rb, "", errors.New("urlTemplate is empty")},
		{"NilRB", nilRB, "t", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rb.SetURLTemplate(tt.t)
			if tt.err != nil {
				if err == nil || err.Error() != tt.err.Error() {
					t.Errorf("got err %v, expected err %v", err, tt.err)
				}
			} else if err != nil {
				t.Errorf("unexpected err %v", err)
			}
		})
	}
}
