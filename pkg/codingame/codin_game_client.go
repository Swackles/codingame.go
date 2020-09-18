package codingame

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
 *		email			- CodinGame email you want to login with
 *		password		- CodinGame password you want to login with
 *
 *	Output:
 *		LoginResponse	- Response structure
 * 		ErrorResponse	- Error response structure
 */
func (client CodinGameClient) Login() (*LoginResponse, *ErrorResponse) {
	loginUrl := baseURL + "CodingamerRemoteService/loginSiteV2"
	payload := []byte(`["` + client.Email + `", "` + client.Password + `", true]`)

	_, body, err := requestPost(loginUrl, payload)
	if err != nil {
		return nil, err
	}

	LoginResponse := newLoginResponse(body)
	return &LoginResponse, nil
}
