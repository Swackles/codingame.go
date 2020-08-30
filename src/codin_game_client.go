package main

import (
	"net/http"
)

type CodinGameClient struct {
	Email        string
	Password     string
	Cookies      []*http.Cookie
	CodinGamerID int
}

/*
 * Use this to login yourself to CodinGame
 *	Inputs:
 *		email			-	CodinGame email you want to login with
 *		password	-	CodinGame password you want to login with
 *
 *	Output:
 *		LoginResponse		-	Response structure
 *		[]http.Cookies	-	Cookies of the request, this will be needed to
 *											do further requests
 */
func (client CodinGameClient) Login() (*LoginResponse, error) {
	loginUrl := baseURL + "CodingamerRemoteService/loginSiteV2"
	payload := []byte(`["` + client.Email + `", "` + client.Password + `", true]`)

	_, body, cookies, err := requestPost(loginUrl, payload, nil)
	if err != nil {
		return nil, err
	}

	client.Cookies = cookies

	LoginResponse := newLoginResponse(body, client)
	return LoginResponse, nil
}
