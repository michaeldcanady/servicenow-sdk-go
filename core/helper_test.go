package core

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yosida95/uritemplate/v3"
)

type MockParseResponse struct {
	ParseHeadersCalled bool
}

func (m *MockParseResponse) ParseHeaders(headers http.Header) {
	m.ParseHeadersCalled = true
}

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (c *MockClient) Send(requestInfo RequestInformation, errorMapping ErrorMapping) (*http.Response, error) {
	url, err := requestInfo.uri.ToUrl()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(http.MethodGet, "", nil)
	req.URL = url
	return c.DoFunc(req)
}

func pick[M ~map[K]V, K comparable, V any](m M) K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(keys))))
	if err != nil {
		panic(err)
	}

	return keys[index.Int64()]
}

type TestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestToQueryMap(t *testing.T) {
	inputs := []struct {
		Input       interface{}
		ShouldError bool
		Expected    map[string]string
		CheckErr    func(error) bool
	}{
		{
			Input: struct {
				Param1 string `url:"param_1"`
				Param2 int    `url:"param_2"`
				Param3 bool   `url:"param_3"`
			}{
				Param1: "value1",
				Param2: 5,
				Param3: true,
			},
			ShouldError: false,
			Expected:    map[string]string{"param_1": "value1", "param_2": "5", "param_3": "true"},
			CheckErr:    nil,
		},
		{
			Input:       nil,
			ShouldError: true,
			Expected:    nil,
			CheckErr:    func(err error) bool { return assert.Equal(t, ErrNilSource, err) },
		},
		{
			Input:       "test",
			ShouldError: true,
			Expected:    nil,
			CheckErr: func(err error) bool {
				return assert.Equal(t, fmt.Errorf("query: Values() expects struct input. Got %v", reflect.String), err)
			},
		},
	}

	for _, input := range inputs {
		paramMap, err := ToQueryMap(input.Input)

		if input.ShouldError {
			if !input.CheckErr(err) {
				t.Errorf("Expected error, got nil")
			}
		} else if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		assert.Equal(t, input.Expected, paramMap)
	}
}

func TestNormalizeVarNames(t *testing.T) {
	input := []string{"VaR1", "vAr2", "VAr3", "vAR4", "Var5", "vaR6", "VAR7"}
	expected := map[string]string{"var1": "VaR1", "var2": "vAr2", "var3": "VAr3", "var4": "vAR4", "var5": "Var5", "var6": "vaR6", "var7": "VAR7"}

	actual := normalizeVarNames(input)

	assert.Equal(t, expected, actual)
}

func TestAddParametersWithOriginalNames(t *testing.T) {
	params := map[string]string{"var1": "val1", "var2": "val2"}
	normalizedNames := map[string]string{"var1": "VaR1", "var2": "vAr2"}

	values := addParametersWithOriginalNames(params, normalizedNames, nil)
	expected := uritemplate.Values{"VaR1": uritemplate.String("val1"), "vAr2": uritemplate.String("val2")}

	assert.Equal(t, expected, values)
}

func TestGetKeyWithOriginalName(t *testing.T) {
	normalizedNames := map[string]string{"var1": "VaR1", "var2": "vAr2", "var3": "VAr3", "var4": "vAR4", "var5": "Var5", "var6": "vaR6", "var7": "VAR7"}

	randomKey := pick(normalizedNames)

	value := getKeyWithOriginalName(randomKey, normalizedNames)

	assert.Equal(t, normalizedNames[randomKey], value)
}

func TestIsPointer(t *testing.T) {
	s := "test"
	i := 42

	f := func() {}

	inputs := []struct {
		Input    interface{}
		Expected bool
	}{
		{
			Input:    &s,
			Expected: true,
		},
		{
			Input:    s, // this is a string value, not a pointer
			Expected: false,
		},
		{
			Input:    i, // this is an int value, not a pointer
			Expected: false,
		},
		{
			Input:    &i, // this is a pointer to an int value
			Expected: true,
		},
		{
			Input:    nil, // this is a nil value, not a pointer
			Expected: false,
		},
		{
			Input:    (*int)(nil), // this is a nil pointer to an int type
			Expected: true,
		},
		{
			Input:    []int{1, 2, 3}, // this is a slice value, not a pointer
			Expected: false,
		},
		{
			Input:    &[3]int{1, 2, 3}, // this is a pointer to an array value
			Expected: true,
		},
		{
			Input:    map[string]int{"a": 1, "b": 2}, // this is a map value, not a pointer
			Expected: false,
		},
		{
			Input:    &map[string]int{"a": 1, "b": 2}, // this is a pointer to a map value
			Expected: true,
		},
		{
			Input:    f, // this is a function value, not a pointer
			Expected: false,
		},
		{
			Input:    &f, // this is a pointer to a function value
			Expected: true,
		},
		{
			Input:    struct{}{}, // this is a struct value, not a pointer
			Expected: false,
		},
		{
			Input:    &struct{}{}, // this is a pointer to a struct value
			Expected: true,
		},
	}

	for _, input := range inputs {
		actual := IsPointer(input.Input)
		assert.Equal(t, input.Expected, actual)
	}
}

func TestFromJson(t *testing.T) {
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

			err := FromJson(tc.response, &result)

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
			err := ParseResponse(tc.response, &tc.value)

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

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type personCollectionResponse struct {
	Result []*person
}

func (r personCollectionResponse) ParseHeaders(header http.Header) {}

func (r personCollectionResponse) ToPage() PageResult[person] {
	return PageResult[person]{
		Result:           r.Result,
		NextPageLink:     "",
		PreviousPageLink: "",
		LastPageLink:     "",
		FirstPageLink:    "",
	}
}

func TestConvertToPage(t *testing.T) {
	tests := []test[PageResult[person]]{
		{
			title: "Valid",
			input: &personCollectionResponse{
				Result: []*person{
					{
						Name: "bob",
						Age:  25,
					},
				},
			},
			expected: PageResult[person]{
				Result: []*person{
					{
						Name: "bob",
						Age:  25,
					},
				},
				NextPageLink:     "",
				PreviousPageLink: "",
				LastPageLink:     "",
				FirstPageLink:    "",
			},
			shouldErr:   false,
			expectedErr: nil,
		},
		{
			title:       "Nil",
			input:       (*personCollectionResponse)(nil),
			expected:    PageResult[person]{},
			shouldErr:   true,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			page, err := convertToPage[person](tt.input.(*personCollectionResponse))

			if tt.shouldErr {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tt.expected, page)
		})
	}
}
