package codingame

import (
	"net/http"
	"reflect"
	"testing"
)

var OKResponse = &http.Response{
	StatusCode: 200,
}
var NotOKResponse = &http.Response{
	StatusCode: 500,
}

var errorCodeString = []byte("{ \"code\": \"code\", \"message\": \"msg\" }")
var errorCodeStringErrorResponse = &ErrorResponse{Code: "code", Message: "msg", StatusCode: 500}

var errorIDString = []byte("{ \"id\": 3, \"message\": \"msg\" }")
var errorIDStringErrorResponse = &ErrorResponse{ID: 3, Message: "msg", StatusCode: 500}

var errorContainerIDString = []byte("{ \"error\": { \"id\": 1, \"message\": \"msg\" }}")
var errorContainerIDStringErrorResponse = &ErrorResponse{ID: 1, Message: "msg", StatusCode: 200}

func TestParseError(t *testing.T) {
	var result *ErrorResponse

	result = parseError(OKResponse, errorContainerIDString)
	if !reflect.DeepEqual(result, errorContainerIDStringErrorResponse) {
		t.Errorf("when received error with 200 statuscode, should return ErrorResponse")
	}

	result = parseError(NotOKResponse, errorCodeString)
	if !reflect.DeepEqual(result, errorCodeStringErrorResponse) {
		t.Errorf("when received error code with 200 statuscode, should return ErrorResponse")
	}

	// response.status != 200 && Error message contains ID
	result = parseError(NotOKResponse, errorIDString)
	if !reflect.DeepEqual(result, errorIDStringErrorResponse) {
		t.Errorf("when received error id with 200 statuscode, should return ErrorResponse")
	}

	result = parseError(OKResponse, []byte("{ \"status\": \"OK\" }"))
	if result != nil {
		t.Errorf("when received correct response with 200 statuscode, should return nil")
	}
}
