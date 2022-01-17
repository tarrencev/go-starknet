package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Error struct {
	StatusCode int    `json:"-"`
	Body       []byte `json:"-"`

	Code     string `json:"code"`
	Message  string `json:"message"`
	Problems string `json:"problems"`
}

// Error implements the error interface.
func (e Error) Error() string {
	return fmt.Sprintf("%d: %s %s", e.StatusCode, e.Code, e.Message)
}

// NewError creates a new Error from an API response.
func NewError(resp *http.Response) error {
	apiErr := Error{StatusCode: resp.StatusCode}
	data, err := ioutil.ReadAll(resp.Body)
	if err == nil && data != nil {
		apiErr.Body = data
		if err := json.Unmarshal(data, &apiErr); err != nil {
			apiErr.Code = "unknown_error_format"
			apiErr.Message = string(data)
		}
	}
	return &apiErr
}
