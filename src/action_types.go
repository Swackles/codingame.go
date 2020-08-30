package main

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
