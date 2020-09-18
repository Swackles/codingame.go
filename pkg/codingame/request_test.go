package codingame

import (
	"net/http"
	"testing"
)

var cookie1 = http.Cookie{
	Name:  "TestName1",
	Value: "TestValue1",
}

var cookie2 = http.Cookie{
	Name:  "TestName2",
	Value: "TestValue2",
}

var cookie3 = http.Cookie{
	Name:  "TestName3",
	Value: "TestValue3",
}

var cookies = []*http.Cookie{&cookie1, &cookie2}

func TestFindCookie(t *testing.T) {
	var result *http.Cookie
	var index int

	index, result = findCookie(cookies, &cookie1)
	if result != &cookie1 {
		t.Errorf("when searching for cookie that exists, should return cookie")
	}
	if index != 0 {
		t.Errorf("when searching for cookie that exists, should return index")
	}

	index, result = findCookie(cookies, &cookie3)
	if result != nil {
		t.Errorf("when searching for cookie that dosen't exists, should return nil")
	}
	if index != -1 {
		t.Errorf("when searching for cookie that exists, should return -1")
	}

}
