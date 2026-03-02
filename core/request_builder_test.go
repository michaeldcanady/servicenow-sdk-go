package core

import (
	"context"
	"testing"
)

func TestNewRequestBuilder2(t *testing.T) {
	tests := []struct {
		name  string
		templ string
	}{
		{"Basic", "t"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := NewRequestBuilder2(nil, tt.templ, nil)
			if res == nil {
				t.Fatal("returned nil")
			}
			if res.UrlTemplate != tt.templ {
				t.Error("failed")
			}
		})
	}
}

func TestRequestBuilder_ToPutRequestInformation2(t *testing.T) {
	rb := &RequestBuilder{}
	tests := []struct {
		name string
	}{
		{"Basic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := rb.ToPutRequestInformation2(nil)
			if res.Method != PUT {
				t.Error("wrong method")
			}
		})
	}
}

func TestRequestBuilder_ToPostRequestInformation3(t *testing.T) {
	rb := &RequestBuilder{}
	res, _ := rb.ToPostRequestInformation3(nil)
	if res.Method != POST {
		t.Error("failed")
	}
}

func TestRequestBuilder_ToDeleteRequestInformation2(t *testing.T) {
	rb := &RequestBuilder{}
	res, _ := rb.ToDeleteRequestInformation2(nil)
	if res.Method != DELETE {
		t.Error("failed")
	}
}

func TestRequestBuilder_ToGetRequestInformation2(t *testing.T) {
	rb := &RequestBuilder{}
	res, _ := rb.ToGetRequestInformation2(nil)
	if res.Method != GET {
		t.Error("failed")
	}
}

func TestRequestBuilder_prepareData(t *testing.T) {
	rb := &RequestBuilder{}
	tests := []struct {
		name  string
		input any
		err   bool
	}{
		{"Nil", nil, false},
		{"Bytes", []byte("v"), false},
		{"Map", map[string]string{"a": "b"}, false},
		{"Bad", 123, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := rb.prepareData(tt.input)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v", err)
			}
		})
	}
}

func TestRequestBuilder_SendMethods(t *testing.T) {
	// These are harder to unit test without complex mocks
	// but we can test the error propagation when ToRequestInformation fails
	rb := &RequestBuilder{}
	config := &RequestConfiguration{Data: 123} // trigger prepareData error

	t.Run("SendGet3", func(t *testing.T) {
		if err := rb.SendGet3(context.Background(), config); err == nil {
			t.Error("expected error")
		}
	})
	t.Run("SendPost4", func(t *testing.T) {
		if err := rb.SendPost4(context.Background(), config); err == nil {
			t.Error("expected error")
		}
	})
	t.Run("SendDelete3", func(t *testing.T) {
		if err := rb.SendDelete3(context.Background(), config); err == nil {
			t.Error("expected error")
		}
	})
	t.Run("SendPut3", func(t *testing.T) {
		if err := rb.SendPut3(context.Background(), config); err == nil {
			t.Error("expected error")
		}
	})
}

func TestNewRequestBuilder(t *testing.T) {
	res := NewRequestBuilder(nil, "t", nil)
	if res == nil {
		t.Error("failed")
	}
}

func TestRequestBuilder_DeprecatedToRequestInformation(t *testing.T) {
	rb := &RequestBuilder{}
	t.Run("Head", func(t *testing.T) {
		res, _ := rb.ToHeadRequestInformation()
		if res.Method != HEAD {
			t.Error("failed")
		}
	})
}
