package core

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToQueryMap_Success(t *testing.T) {

	params := struct {
		Param1 string `query:"param_1"`
		Param2 string `query:"param_2"`
	}{
		Param1: "value1",
		Param2: "value2",
	}

	actual, err := ToQueryMap(params)
	if err != nil {
		t.Error(err)
	}

	expected := map[string]string{"param_1": "value1", "param_2": "value2"}

	assert.Equal(t, expected, actual)
}

func TestToQueryMap_Nil(t *testing.T) {

	_, err := ToQueryMap(nil)
	assert.Error(t, err)
}

func TestToQueryMap_Map(t *testing.T) {

	expected := map[string]string{"param_1": "value1", "param_2": "value2"}
	actual, err := ToQueryMap(expected)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, actual)
}

func TestNormalizeVarNames(t *testing.T) {

	varNames := []string{"Var1", "vAr2", "VaR3"}

	normalizedNames := normalizeVarNames(varNames)
	expected := map[string]string{"var1": "Var1", "var2": "vAr2", "var3": "VaR3"}

	assert.Equal(t, expected, normalizedNames)
}

func TestIsPointer(t *testing.T) {

	val1 := "val1"

	isPointer := IsPointer(&val1)
	assert.Equal(t, true, isPointer)

	isPointer = IsPointer(nil)
	assert.Equal(t, false, isPointer)
}

func TestFromJson(t *testing.T) {

	type TestResponse struct {
		Name string `json:"name"`
	}

	responseData := `{"name": "John Doe"}`

	// Create a fake HTTP response
	resp := &http.Response{
		Body:          ioutil.NopCloser(strings.NewReader(responseData)),
		ContentLength: int64(len(responseData)),
		Status:        "200 OK",
		StatusCode:    200,
		Header:        make(http.Header),
	}

	// Test successful case
	result, err := FromJson[TestResponse](resp)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "John Doe", result.Name)

	// Test with a non-pointer response
	nonPtrResp := &http.Response{}
	result, err = FromJson[TestResponse](nonPtrResp)
	assert.Error(t, err)
	assert.Nil(t, result)

	// Test with invalid JSON
	invalidJSONResp := &http.Response{
		Body:          ioutil.NopCloser(strings.NewReader("invalid json")),
		ContentLength: int64(len("invalid json")),
		Status:        "200 OK",
		StatusCode:    200,
		Header:        make(http.Header),
	}
	result, err = FromJson[TestResponse](invalidJSONResp)
	assert.Error(t, err)
	assert.Nil(t, result)
}
