package core

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestToQueryMap(t *testing.T) {
	type Params struct {
		A string `url:"a"`
	}
	tests := []struct {
		name   string
		input  any
		expect int
		err    error
	}{
		{"Struct", Params{A: "v"}, 1, nil},
		{"Nil", nil, 0, ErrNilSource},
		{"Map", map[string]string{"key": "value"}, 1, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := ToQueryMap(tt.input)
			if err != tt.err {
				t.Errorf("got error %v, expected %v", err, tt.err)
			}
			if len(res) != tt.expect {
				t.Errorf("got %d, expected %d", len(res), tt.expect)
			}
		})
	}
}

func TestFromJson(t *testing.T) {
	t.Run("Successful", func(t *testing.T) {
		resp := &http.Response{Body: io.NopCloser(strings.NewReader(`{"a":"b"}`))}
		var target map[string]string
		err := FromJson(resp, &target)
		if err != nil {
			t.Errorf("unexpected err %v", err)
		}
	})

	t.Run("Nil Response", func(t *testing.T) {
		err := FromJson[map[string]string](nil, nil)
		if err != ErrNilResponse {
			t.Errorf("expected ErrNilResponse, got %v", err)
		}
	})

	t.Run("Nil Body", func(t *testing.T) {
		resp := &http.Response{Body: io.NopCloser(strings.NewReader(""))}
		var target map[string]string
		err := FromJson(resp, &target)
		if err != ErrNilResponseBody {
			t.Errorf("expected ErrNilResponseBody, got %v", err)
		}
	})
}

type mockResponse struct {
	Body []byte
}

func (m *mockResponse) ParseHeaders(headers http.Header) {}

func TestParseResponse(t *testing.T) {
	resp := &http.Response{
		Body:   io.NopCloser(strings.NewReader(`{"Body":"test"}`)),
		Header: http.Header{},
	}
	var target *mockResponse
	err := ParseResponse(resp, &target)
	if err != nil {
		t.Errorf("unexpected err %v", err)
	}
}

func TestIsPointer(t *testing.T) {
	s := "test"
	if IsPointer(s) { t.Error("failed") }
	if !IsPointer(&s) { t.Error("failed") }
}
