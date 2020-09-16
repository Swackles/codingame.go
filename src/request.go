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
 * 			url 		- The url request should be sent to
 *			payload - The payload that gets attached to the request
 * 			cookies	-	HTTP cookies that are needed for the request	(nullable)
 *
 *		Output:
 *			http.Response		-	HTTP response struct for the request
 *			[]byte					-	Response body in byte array
 *			[]http.Cookies	-	Cookies that were returned back by the request
 *			ErrorResponse		-	Returns this if response is error
 */
func requestPost(url string, payload []byte, cookies []*http.Cookie) (*http.Response, []byte, []*http.Cookie, *ErrorResponse) {
	client := createClient(url, cookies)

	res, _ := client.Post(url, "application/json;charset=UTF-8", bytes.NewBuffer(payload))

	for _, newCookie := range res.Cookies() {
		index, cookie := findCookie(cookies, newCookie)
		if index != -1 {
			cookies[index] = cookie
		}
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := parseError(res, body)
	return res, body, cookies, err
}

// Create a new http client
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
