package batchapi

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

type rawBatchResponseItem struct {
	Body          string        `json:"body"`
	ErrorMessage  string        `json:"error_message"`
	ExecutionTime time.Duration `json:"execution_time"` //make duration
	Headers       []batchHeader `json:"headers"`
	ID            string        `json:"id"`
	RedirectURL   string        `json:"redirect_url"`
	StatusCode    int           `json:"status_code"`
	StatusText    string        `json:"status_text"`
}

type BatchResponseItem struct {
	Body          map[string]interface{}
	ErrorMessage  string
	ExecutionTime time.Duration
	Headers       []batchHeader
	ID            string
	RedirectURL   string
	StatusCode    int
	StatusText    string
}

func (r *BatchResponseItem) UnmarshalJSON(data []byte) error {

	rawData := rawBatchResponseItem{}

	err := json.Unmarshal(data, &rawData)
	if err != nil {
		return err
	}

	body, err := base64.StdEncoding.DecodeString(rawData.Body)
	if err != nil {
		return err
	}

	r.ErrorMessage = rawData.ErrorMessage
	r.Headers = rawData.Headers
	r.ID = rawData.ID
	r.RedirectURL = rawData.RedirectURL
	r.StatusCode = rawData.StatusCode
	r.StatusText = rawData.StatusText
	r.ExecutionTime = time.Duration(rawData.ExecutionTime) * time.Second

	err = json.Unmarshal(body, &r.Body)
	if err != nil {
		return err
	}

	return nil
}
