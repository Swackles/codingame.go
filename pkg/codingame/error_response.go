package codingame

import (
	"encoding/json"
	"net/http"
)

type ErrorContainer struct {
	Error *ErrorResponse `json:"error,omitempty"`
}
type ErrorResponse struct {
	ID         int    `json:"id,omitempty"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int
}

// This method tests weather the returned data is an error or not
func parseError(res *http.Response, body []byte) *ErrorResponse {
	var errRes *ErrorResponse
	var errContainer *ErrorContainer

	// Issue with status code
	if res.StatusCode != 200 {
		json.Unmarshal(body, &errRes)

		errRes.StatusCode = res.StatusCode
	} else {
		json.Unmarshal(body, &errContainer)
		errRes = errContainer.Error
		if errRes != nil {
			errRes.StatusCode = res.StatusCode
		}
	}
	return errRes
}
