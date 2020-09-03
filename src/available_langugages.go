package codinGame

import (
	"encoding/json"
)

type AvailableLanguages struct {
	ID         string `json:"id"`
	Last       bool   `json:"last"`
	Onboarding bool   `json:"onboarding"`
	Solved     bool   `json:"solved"`
}

func newAvailableLanguages(body []byte) []AvailableLanguages {
	var response []AvailableLanguages
	json.Unmarshal(body, &response)
	return response
}
