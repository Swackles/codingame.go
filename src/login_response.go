package codinGame

import (
	"encoding/json"
)

type LoginResponse struct {
	Test    string `json:"test"`
	Success struct {
		CodinGamer                   CodinGamer         `json:"codinGamer"`
		CountryCode                  string             `json:"countryCode"`
		NewFeatures                  []interface{}      `json:"newFeatures"`
		EnabledNotifications         []string           `json:"enabledNotifications"`
		NotificationConfig           NotificationConfig `json:"notificationConfig"`
		ActivatedFeatures            []string           `json:"activatedFeatures"`
		ActionTypes                  ActionTypes        `json:"actionTypes"`
		ChatToken                    string             `json:"chatToken"`
		DisqusSsoPayload             string             `json:"disqusSsoPayload"`
		GaveFeedbackOnCurrentCompany bool               `json:"gaveFeedbackOnCurrentCompany"`
		XpThresholds                 XpThresholds       `json:"xpThresholds"`
		User                         User               `json:"user"`
		Visitor                      bool               `json:"visitor"`
		Admin                        bool               `json:"admin"`
		LanguageID                   int                `json:"languageId"`
		UserID                       int                `json:"userId"`
		UserEmail                    string             `json:"userEmail"`
		Impersonated                 bool               `json:"impersonated"`
	} `json:"success"`
}

type NotificationConfig struct {
	SoundEnabled              bool `json:"soundEnabled"`
	NativeNotificationEnabled bool `json:"nativeNotificationEnabled"`
}

type XpThresholds []struct {
	Level           int             `json:"level"`
	XpThreshold     int             `json:"xpThreshold"`
	CumulativeXp    int             `json:"cumulativeXp"`
	RewardLanguages RewardLanguages `json:"rewardLanguages,omitempty"`
}

type RewardLanguages struct {
	Num1 string `json:"1"`
	Num2 string `json:"2"`
}

type User struct {
	ID         int        `json:"id"`
	Email      string     `json:"email"`
	LanguageID int        `json:"languageId"`
	Status     string     `json:"status"`
	Properties Properties `json:"properties"`
}

func newLoginResponse(body []byte, client CodinGameClient) *LoginResponse {
	var response *LoginResponse
	json.Unmarshal(body, &response)

	client.CodinGamerID = response.Success.CodinGamer.UserID
	response.Success.CodinGamer.Client = client

	return response
}
