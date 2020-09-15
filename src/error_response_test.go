package codingame

import (
	"net/http"
	"reflect"
	"testing"
)

// Test respones
var OKResponse = &http.Response{
	StatusCode: 200,
}
var NotOKResponse = &http.Response{
	StatusCode: 500,
}

// Test bodies to be used
var ErrorCodeString = []byte("{ \"code\": \"UserRequired\", \"message\": \"Only a logged user is authorized to perform this operation\" }")
var ErrorIDString = []byte("{ \"id\": -3, \"message\": \"Service not found: lastactivities.getlastactivities(1)\" }")
var ErrorContainerIDString = []byte("{ \"error\": { \"id\": 334, \"message\": \"Malformed email\" }}")

// Test when TestParseError should return ErrorResponse
func TestParseErrorReturnErrorResponse(t *testing.T) {
	// response.status == 200 && Error message contains ID
	result := parseError(OKResponse, ErrorContainerIDString)
	expected := &ErrorResponse{
		ID:         334,
		Message:    "Malformed email",
		StatusCode: 200,
	}
	testParseErrorPrint(testParseErrorStruct{
		success:  reflect.DeepEqual(result, expected),
		input:    []string{"OKResponse", "ErrorString"},
		expected: expected,
		got:      result,
	}, t)

	// response.status != 200 && Error message contains code
	result = parseError(NotOKResponse, ErrorCodeString)
	expected = &ErrorResponse{
		Code:       "UserRequired",
		Message:    "Only a logged user is authorized to perform this operation",
		StatusCode: 500,
	}
	testParseErrorPrint(testParseErrorStruct{
		success:  reflect.DeepEqual(result, expected),
		input:    []string{"NotOKResponse", "ErrorString"},
		expected: expected,
		got:      result,
	}, t)

	// response.status != 200 && Error message contains ID
	result = parseError(NotOKResponse, ErrorIDString)
	expected = &ErrorResponse{
		ID:         -3,
		Message:    "Service not found: lastactivities.getlastactivities(1)",
		StatusCode: 500,
	}
	testParseErrorPrint(testParseErrorStruct{
		success:  reflect.DeepEqual(result, expected),
		input:    []string{"NotOKResponse", "ErrorString"},
		expected: expected,
		got:      result,
	}, t)
}

// Test when parseError should return nil
func TestParseErrorReturnNil(t *testing.T) {
	result := parseError(OKResponse, []byte("{ \"status\": \"OK\" }"))
	testParseErrorPrint(testParseErrorStruct{
		success:  result == nil,
		input:    []string{"OKResponse", "{ \"status\": \"OK\" }"},
		expected: nil,
		got:      result,
	}, t)
}

type testParseErrorStruct struct {
	success  bool
	input    []string
	expected *ErrorResponse
	got      *ErrorResponse
}

func testParseErrorPrint(input testParseErrorStruct, t *testing.T) {
	method := "parseError(" + input.input[0] + ", " + input.input[1] + ")"
	if input.success {
		t.Logf("%+v success,\n    should return %+v", method, input.expected)
	} else {
		t.Errorf("%+v failed\n    expected: %+v\n         got: %+v", method, input.expected, input.got)
	}
}
