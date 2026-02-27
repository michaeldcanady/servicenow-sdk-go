package core

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestSendMethods(t *testing.T) {
	client := &mockCoreClient{
		SendFunc: func(requestInfo IRequestInformation, errorMapping ErrorMapping) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"Body":"test"}`)),
				Header:     http.Header{},
			}, nil
		},
	}
	rb := NewRequestBuilder(client, "{+baseurl}/test", map[string]string{"baseurl": "http://localhost"})

	t.Run("sendGet", func(t *testing.T) {
		var res *mockResponse
		err := sendGet(rb, nil, nil, &res)
		assert.NoError(t, err)
	})

	t.Run("SendGet2", func(t *testing.T) {
		var res Response = &mockResponse{}
		config := &RequestConfiguration{
			Response: res,
		}
		err := SendGet2(rb, config)
		assert.NoError(t, err)
	})

	t.Run("sendPost", func(t *testing.T) {
		var res *mockResponse
		err := sendPost(rb, map[string]string{"a": "b"}, nil, nil, &res)
		assert.NoError(t, err)
	})

	t.Run("SendPost2", func(t *testing.T) {
		var res Response = &mockResponse{}
		config := &RequestConfiguration{
			Response: res,
			Data:     map[string]string{"a": "b"},
		}
		err := SendPost2(rb, config)
		assert.NoError(t, err)
	})

	t.Run("sendDelete", func(t *testing.T) {
		err := sendDelete(rb, nil, nil)
		assert.NoError(t, err)
	})

	t.Run("sendDelete2", func(t *testing.T) {
		err := sendDelete2(rb, nil)
		assert.Equal(t, errors.New("config is nil"), err)
	})

	t.Run("sendPut", func(t *testing.T) {
		var res *mockResponse
		err := sendPut(rb, map[string]string{"a": "b"}, nil, nil, &res)
		assert.NoError(t, err)
	})

	t.Run("sendPut2", func(t *testing.T) {
		var res Response = &mockResponse{}
		config := &RequestConfiguration{
			Response: res,
			Data:     map[string]string{"a": "b"},
		}
		err := sendPut2(rb, config)
		assert.NoError(t, err)
	})
}

func TestIsPointer(t *testing.T) {
	s := "test"
	if IsPointer(s) {
		t.Error("failed")
	}
	if !IsPointer(&s) {
		t.Error("failed")
	}
}
