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
 * 			path 		- The url request should be sent to
 *			payload - The payload that gets attached to the request
 * 			cookies	-	HTTP cookies that are needed for the request	(nullable)
 *
 *		Output:
 *			http.Response		-	HTTP response struct for the request
 *			[]byte					-	Response body in byte array
 *			[]http.Cookies	-	Cookies that were returned back by the request
 */
func requestPost(url string, payload []byte, cookies []*http.Cookie) (*http.Response, []byte, []*http.Cookie, error) {
	client := createClient(url, cookies)

	res, err := client.Post(url, "application/json;charset=UTF-8", bytes.NewBuffer(payload))
	if err != nil {
		return nil, nil, nil, err
	}

	cookies = res.Cookies()

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return res, body, cookies, nil
}

func createClient(path string, cookies []*http.Cookie) *http.Client {
	jar, _ := cookiejar.New(nil)
	u, _ := url.Parse(path)
	jar.SetCookies(u, cookies)

	return &http.Client{Jar: jar}
}
