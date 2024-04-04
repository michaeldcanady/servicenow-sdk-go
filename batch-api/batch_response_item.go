package batchapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type BatchResponseItem[T any] interface {
	GetBody() (*T, error)
	GetType() reflect.Type
	GetExecutionTime() time.Duration
	GetHeaders() []batchHeader
	GetID() string
	GetRedirectURL() string
	GetStatusCode() int
}

type batchResponseItem struct {
	Body          string        `json:"body"`
	ErrorMessage  string        `json:"error_message"`
	ExecutionTime int64         `json:"execution_time"`
	Headers       []batchHeader `json:"headers"`
	ID            string        `json:"id"`
	RedirectURL   string        `json:"redirect_url"`
	StatusCode    int           `json:"status_code"`
	StatusText    string        `json:"status_text"`
}

func (rI *batchResponseItem) GetBody() (*any, error) {
	var body any

	unencoded, err := base64.StdEncoding.DecodeString(rI.Body)
	if err != nil {
		unencoded = []byte(rI.Body)
	}

	err = json.Unmarshal(unencoded, &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func (rI *batchResponseItem) GetType() reflect.Type {
	body, _ := rI.GetBody()
	rv := reflect.ValueOf(body)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		fmt.Println(rv.Kind(), rv.Type(), rv)
		rv = rv.Elem()
	}

	return rv.Type()
}

func (rI *batchResponseItem) GetExecutionTime() time.Duration {
	return time.Duration(rI.ExecutionTime)
}

func (rI *batchResponseItem) GetHeaders() []batchHeader {
	return rI.Headers
}

func (rI *batchResponseItem) GetID() string {
	return rI.ID
}

func (rI *batchResponseItem) GetRedirectURL() string {
	return rI.RedirectURL
}

func (rI *batchResponseItem) GetStatusCode() int {
	return rI.StatusCode
}
