package codingame

import (
	"encoding/json"
)

type LoginResponseContainer struct {
	Test    string `json:"test"`
	Data LoginResponse `json:"success"`
}

type LoginResponse struct {
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

type ActionTypes struct {
	DeleteContribution                   ActionType `json:"deleteContribution"`
	ReceiveCareerNotification            ActionType `json:"receiveCareerNotification"`
	DeleteComment                        ActionType `json:"deleteComment"`
	SubmitMultiContribution              ActionType `json:"submitMultiContribution"`
	SendPingEmailToAll                   ActionType `json:"sendPingEmailToAll"`
	AccessCoursesBeta                    ActionType `json:"accessCoursesBeta"`
	EditAcceptedContribution             ActionType `json:"editAcceptedContribution"`
	UpdateAvatar                         ActionType `json:"updateAvatar"`
	UpdateAllAvatar                      ActionType `json:"updateAllAvatar"`
	AccessDisabledLeaderboard            ActionType `json:"accessDisabledLeaderboard"`
	Impersonate                          ActionType `json:"impersonate"`
	EditComment                          ActionType `json:"editComment"`
	EditContribution                     ActionType `json:"editContribution"`
	DenyPuzzleContribution               ActionType `json:"denyPuzzleContribution"`
	SendPingEmailToOptinCodingamers      ActionType `json:"sendPingEmailToOptinCodingamers"`
	DenyClashContribution                ActionType `json:"denyClashContribution"`
	ValidateClashContribution            ActionType `json:"validateClashContribution"`
	SendPrivateMessageToAll              ActionType `json:"sendPrivateMessageToAll"`
	DenyContribution                     ActionType `json:"denyContribution"`
	EditClashContribution                ActionType `json:"editClashContribution"`
	UpdateCover                          ActionType `json:"updateCover"`
	EditPuzzleContribution               ActionType `json:"editPuzzleContribution"`
	SendPrivateMessageToOptinCodingamers ActionType `json:"sendPrivateMessageToOptinCodingamers"`
	EditAcceptedPuzzleContribution       ActionType `json:"editAcceptedPuzzleContribution"`
	UpdateCareerSettings                 ActionType `json:"updateCareerSettings"`
	EditAcceptedClashContribution        ActionType `json:"editAcceptedClashContribution"`
	ValidateContribution                 ActionType `json:"validateContribution"`
	ValidatePuzzleContribution           ActionType `json:"validatePuzzleContribution"`
	UpdateAllCover                       ActionType `json:"updateAllCover"`
}

type ActionType struct {
	AuthorPolicy       string `json:"authorPolicy"`
	MinCodinGamerCount int    `json:"minCodinGamerCount"`
}

func newLoginResponse(body []byte) LoginResponse {
	var responseContainer LoginResponseContainer
	json.Unmarshal(body, &responseContainer)

	var response LoginResponse = responseContainer.Data
	client.CodinGamerID = response.CodinGamer.UserID

	return response
}
