package codingame

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

/*
 *	Performs the basic POST request
 * 		Inputs:
 * 			url 			- The url request should be sent to
 *			payload 		- The payload that gets attached to the request
 *
 *		Output:
 *			http.Response	- HTTP response struct for the request
 *			[]byte			- Response body in byte array
 *			ErrorResponse	- Returns this if response is error
 */
func requestPost(url string, payload []byte) (*http.Response, []byte, *ErrorResponse) {
	httpClient := createClient(url, client.Cookies)

	res, _ := httpClient.Post(url, "application/json;charset=UTF-8", bytes.NewBuffer(payload))
	
	if client.Cookies == nil {
		// User just logged in and dosen't have any cookies
		client.Cookies = res.Cookies()
	} else {
		for _, newCookie := range res.Cookies() {
			index, _ := findCookie(client.Cookies, newCookie)
			if index != -1 {
				client.Cookies[index] = newCookie
			}
		}
	}


	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := parseError(res, body)
	return res, body, err
}

// Create a new http client for request
func createClient(path string, cookies []*http.Cookie) *http.Client {
	jar, _ := cookiejar.New(nil)
	u, _ := url.Parse(path)
	jar.SetCookies(u, cookies)

	return &http.Client{Jar: jar}
}

// Finds specific cookie from array of cookies
func findCookie(cookies []*http.Cookie, findCookie *http.Cookie) (int, *http.Cookie) {
	for index, cookie := range cookies {
		if cookie.Name == findCookie.Name {
			return index, cookie
		}
	}
	return -1, nil
}
