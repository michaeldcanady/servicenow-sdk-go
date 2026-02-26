package core

import (
	"context"
	"testing"
)

func TestNewRequestBuilder(t *testing.T) {
	tests := []struct {
		name  string
		templ string
	}{
		{"Basic", "t"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := NewRequestBuilder(nil, tt.templ, nil)
			if res == nil {
				t.Fatal("returned nil")
			}
			if res.UrlTemplate != tt.templ {
				t.Error("failed")
			}
		})
	}
}

func TestRequestBuilder_ToPutRequestInformation(t *testing.T) {
	rb := &RequestBuilder{}
	tests := []struct {
		name string
	}{
		{"Basic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := rb.ToPutRequestInformation(nil)
			if res.Method != PUT {
				t.Error("wrong method")
			}
		})
	}
}

func TestRequestBuilder_ToPostRequestInformation(t *testing.T) {
	rb := &RequestBuilder{}
	res, _ := rb.ToPostRequestInformation(nil)
	if res.Method != POST { t.Error("failed") }
}

func TestRequestBuilder_ToDeleteRequestInformation(t *testing.T) {
	rb := &RequestBuilder{}
	res, _ := rb.ToDeleteRequestInformation(nil)
	if res.Method != DELETE { t.Error("failed") }
}

func TestRequestBuilder_ToGetRequestInformation(t *testing.T) {
	rb := &RequestBuilder{}
	res, _ := rb.ToGetRequestInformation(nil)
	if res.Method != GET { t.Error("failed") }
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
	rb := &RequestBuilder{}
	config := &RequestConfiguration{Data: 123} // trigger prepareData error
	
	t.Run("SendGet", func(t *testing.T) {
		if err := rb.SendGet(context.Background(), config); err == nil {
			t.Error("expected error")
		}
	})
	t.Run("SendPost", func(t *testing.T) {
		if err := rb.SendPost(context.Background(), config); err == nil {
			t.Error("expected error")
		}
	})
	t.Run("SendDelete", func(t *testing.T) {
		if err := rb.SendDelete(context.Background(), config); err == nil {
			t.Error("expected error")
		}
	})
	t.Run("SendPut", func(t *testing.T) {
		if err := rb.SendPut(context.Background(), config); err == nil {
			t.Error("expected error")
		}
	})
}
