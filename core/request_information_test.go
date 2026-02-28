package core

import (
	"context"
	"net/http"
	"net/url"
	"testing"
)

func TestNewRequestInformation(t *testing.T) {
	ri := NewRequestInformation()
	if ri == nil {
		t.Fatal("returned nil")
	}
}

func TestRequestInformation_AddRequestOptions(t *testing.T) {
	ri := NewRequestInformation()
	ri.options = nil
	tests := []struct {
		name string
		opts []RequestOption
	}{
		{"Nil", nil},
		{"Single", []RequestOption{&mockRequestOption{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ri.AddRequestOptions(tt.opts)
		})
	}
}

func TestRequestInformation_GetRequestOptions(t *testing.T) {
	ri := NewRequestInformation()
	t.Run("Empty", func(t *testing.T) {
		if len(ri.GetRequestOptions()) != 0 {
			t.Error("expected empty")
		}
	})
	t.Run("NilInternal", func(t *testing.T) {
		ri.options = nil
		if len(ri.GetRequestOptions()) != 0 {
			t.Error("expected empty")
		}
	})
}

func TestRequestInformation_SetStreamContent(t *testing.T) {
	ri := NewRequestInformation()
	data := []byte("test")
	ri.SetStreamContent(data)
	if string(ri.Content) != "test" {
		t.Error("failed to set content")
	}
	if ri.Headers.Get("Content-Type") != "application/octet-stream" {
		t.Error("failed to set header")
	}
}

func TestRequestInformation_Url(t *testing.T) {
	ri := NewRequestInformation()
	ri.uri.UrlTemplate = "http://{+baseurl}/t"
	ri.uri.PathParameters["baseurl"] = "test.com"
	res, _ := ri.Url()
	if res != "http://test.com/t" {
		t.Errorf("got %s", res)
	}

	riBad := NewRequestInformation()
	_, err := riBad.Url()
	if err == nil {
		t.Error("expected error for missing params")
	}
}

func TestRequestInformation_ToRequest(t *testing.T) {
	ri := NewRequestInformation()
	ri.uri.UrlTemplate = "http://{+baseurl}/t"
	ri.uri.PathParameters["baseurl"] = "test.com"
	ri.Method = GET
	ri.Headers.Add("A", "B")

	req, err := ri.ToRequest()
	if err != nil {
		t.Errorf("unexpected err %v", err)
	}
	if req.Header.Get("A") != "B" {
		t.Error("header missing")
	}

	riBad := NewRequestInformation()
	_, err = riBad.ToRequest()
	if err == nil {
		t.Error("expected error")
	}
}

func TestRequestInformation_ToRequestWithContext(t *testing.T) {
	ri := NewRequestInformation()
	ri.uri.UrlTemplate = "http://{+baseurl}/t"
	ri.uri.PathParameters["baseurl"] = "test.com"

	_, err := ri.ToRequestWithContext(context.Background())
	if err != nil {
		t.Errorf("unexpected err %v", err)
	}

	riBad := NewRequestInformation()
	_, err = riBad.ToRequestWithContext(context.Background())
	if err == nil {
		t.Error("expected error")
	}
}

func TestRequestInformation_AddHeaders(t *testing.T) {
	ri := NewRequestInformation()
	tests := []struct {
		name  string
		input any
		err   bool
	}{
		{"Header", http.Header{"A": []string{"B"}}, false},
		{"Struct", struct {
			A string `header:"A"`
		}{A: "B"}, false},
		{"BadType", 123, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ri.AddHeaders(tt.input)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestRequestInformation_getContentReader(t *testing.T) {
	ri := NewRequestInformation()
	ri.Content = []byte("test")
	reader := ri.getContentReader()
	if reader == nil {
		t.Fatal("returned nil")
	}
}

func TestRequestInformation_SetUri(t *testing.T) {
	ri := NewRequestInformation()
	u, _ := url.Parse("http://test")
	ri.SetUri(u)
	if ri.uri.PathParameters[rawURLKey] != "http://test" {
		t.Error("failed to set raw url")
	}
}

func TestRequestInformation_AddQueryParameters(t *testing.T) {
	ri := NewRequestInformation()
	err := ri.AddQueryParameters(struct {
		A string `url:"a"`
	}{A: "v"})
	if err != nil {
		t.Errorf("unexpected err %v", err)
	}
}

type mockRequestOption struct{}

func (m *mockRequestOption) GetKey() RequestOptionKey { return RequestOptionKey{Key: "mock"} }
