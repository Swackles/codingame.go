package codingame

const baseURL = "https://www.codingame.com/services/"

var client CodinGameClient

// Initialize new instance of codingame
func New(email string, password string) (*LoginResponse, *ErrorResponse) {
	client := CodinGameClient{
		Email:    email,
		Password: password,
	}

	loginResponse, err := client.Login()
	return loginResponse, err
}
