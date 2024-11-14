package internal

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

type MockParseResponse struct {
	ParseHeadersCalled bool
}

func (m *MockParseResponse) ParseHeaders(headers http.Header) {
	m.ParseHeadersCalled = true
}

type TestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestFromJSON(t *testing.T) {
	testCases := []struct {
		name          string
		response      *http.Response
		expectedValue TestData
		expectedError error
	}{
		{
			name: "Valid JSON",
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader(`{"name":"John","age":30}`)),
				StatusCode: http.StatusOK,
			},
			expectedValue: TestData{Name: "John", Age: 30},
			expectedError: nil,
		},
		{
			name: "Invalid JSON",
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader(`invalid-json`)),
				StatusCode: http.StatusOK,
			},
			expectedValue: TestData{},
			expectedError: errors.New("invalid character 'i' looking for beginning of value"),
		},
		{
			name:          "Nil Response",
			response:      nil,
			expectedValue: TestData{},
			expectedError: ErrNilResponse,
		},
		{
			name: "Nil Response Body",
			response: &http.Response{
				Body:       http.NoBody,
				StatusCode: http.StatusOK,
			},
			expectedValue: TestData{},
			expectedError: ErrNilResponseBody,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result TestData

			err := FromJSON(tc.response, &result)

			if err != nil && tc.expectedError != nil {
				if err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
				}
			} else if err != nil || tc.expectedError != nil {
				t.Errorf("Unexpected error. Expected: %v, got: %v", tc.expectedError, err)
			}

			if result != tc.expectedValue {
				t.Errorf("Expected value: %v, got: %v", tc.expectedValue, result)
			}
		})
	}
}

func TestParseResponse(t *testing.T) {
	testCases := []struct {
		name           string
		response       *http.Response
		value          *MockParseResponse
		expectedError  error
		expectedCalled bool
	}{
		{
			name: "Valid Response",
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader(`{"name":"John","age":30}`)),
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": {"application/json"}},
			},
			value:          &MockParseResponse{},
			expectedError:  nil,
			expectedCalled: true,
		},
		{
			name: "Invalid JSON",
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader(`invalid-json`)),
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": {"application/json"}},
			},
			value:          &MockParseResponse{},
			expectedError:  errors.New("invalid character 'i' looking for beginning of value"),
			expectedCalled: false,
		},
		{
			name:           "Nil Response",
			response:       nil,
			value:          &MockParseResponse{},
			expectedError:  ErrNilResponse,
			expectedCalled: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//nolint:gosec
			err := ParseResponse(tc.response, tc.value)

			if err != nil && tc.expectedError != nil {
				if err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
				}
			} else if err != nil || tc.expectedError != nil {
				t.Errorf("Unexpected error. Expected: %v, got: %v", tc.expectedError, err)
			}

			if tc.expectedCalled != tc.value.ParseHeadersCalled {
				t.Errorf("Expected ParseHeaders to be called: %v, but it was not", tc.expectedCalled)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	testCases := []mocking.Test[bool]{
		{
			Title:       "Map",
			Input:       nil,
			ExpectedErr: nil,
			Expected:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			output := IsNil(tc.Input)
			assert.Equal(t, tc.Expected, output)
		})
	}
}
