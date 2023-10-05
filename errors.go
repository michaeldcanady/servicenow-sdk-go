package servicenowsdkgo

import "fmt"

type ErrorMapping map[string]string

type ApiError struct {
	Message            string
	ResponseStatusCode int
}

func (e *ApiError) Error() string {
	if len(e.Message) > 0 {
		return fmt.Sprint(e.Message)
	} else {
		return "error status code received from the API"
	}
}

type Exception struct {
	Detail  string
	Message string
}

type ServiceNowError struct {
	Exception Exception `json:"error"`
	Status    string
}

func (e *ServiceNowError) Error() string {
	return fmt.Sprintf("%s: %s", e.Exception.Message, e.Exception.Detail)
}
