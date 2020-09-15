package codingame

import (
	"encoding/json"
	"fmt"
)

type PuzzleResponse []struct {
	Type          string `json:"type"`
	Puzzle        Puzzle `json:"puzzle,omitempty"`
	LastClashTime int64  `json:"lastClashTime,omitempty"`
}

type Puzzle struct {
	ID                   int         `json:"id"`
	Level                string      `json:"level"`
	Rank                 int         `json:"rank"`
	ThumbnailBinaryID    int64       `json:"thumbnailBinaryId"`
	PreviewBinaryID      int64       `json:"previewBinaryId"`
	CoverBinaryID        int64       `json:"coverBinaryId"`
	LogoBinaryID         int64       `json:"logoBinaryId"`
	Title                string      `json:"title"`
	Description          string      `json:"description"`
	ValidatorScore       int         `json:"validatorScore"`
	AchievementCount     int         `json:"achievementCount"`
	DoneAchievementCount int         `json:"doneAchievementCount"`
	LastActivity         int64       `json:"lastActivity"`
	ForumLink            string      `json:"forumLink"`
	Contributor          Contributor `json:"contributor"`
	ChatRoom             string      `json:"chatRoom"`
	Score                float64     `json:"score"`
	Position             int         `json:"position"`
	Total                int         `json:"total"`
	GlobalTotal          int         `json:"globalTotal"`
	SolvedCount          int         `json:"solvedCount"`
	AttemptCount         int         `json:"attemptCount"`
	XpPoints             int         `json:"xpPoints"`
	Feedback             Feedback    `json:"feedback"`
	OptimCriteriaID      string      `json:"optimCriteriaId"`
	Topics               []Topic     `json:"topics"`
	CreationTime         int64       `json:"creationTime"`
	League               League      `json:"league"`
	Type                 string      `json:"type"`
	PrettyID             string      `json:"prettyId"`
	DetailsPageURL       string      `json:"detailsPageUrl"`
	ReplayIds            []int       `json:"replayIds"`
	TestSessionHandle    string      `json:"testSessionHandle"`
	PuzzleLeaderboardID  string      `json:"puzzleLeaderboardId"`
	CommunityCreation    bool        `json:"communityCreation"`
	Client               CodinGameClient
}

type Contributor struct {
	UserID       int    `json:"userId"`
	Pseudo       string `json:"pseudo"`
	PublicHandle string `json:"publicHandle"`
	Enable       bool   `json:"enable"`
	Avatar       int64  `json:"avatar"`
	Cover        int64  `json:"cover"`
}

type Feedback struct {
	FeedbackID         int   `json:"feedbackId"`
	CodingamerFeedback int   `json:"codingamerFeedback"`
	Feedbacks          []int `json:"feedbacks"`
}

type Topic struct {
	Handle   string  `json:"handle"`
	Value    string  `json:"value"`
	Children []Child `json:"children"`
}

type Child struct {
	Handle   string        `json:"handle"`
	Value    string        `json:"value"`
	Children []interface{} `json:"children"`
}

type League struct {
	DivisionIndex       int `json:"divisionIndex"`
	DivisionCount       int `json:"divisionCount"`
	OpeningLeaguesCount int `json:"openingLeaguesCount"`
	DivisionOffset      int `json:"divisionOffset"`
}

func newPuzzleResponse(body []byte) PuzzleResponse {
	var response PuzzleResponse
	json.Unmarshal(body, &response)
	return response
}

func (puzzle Puzzle) GetProgrammingLanguages(codinGamer CodinGamer) ([]AvailableLanguages, *ErrorResponse) {
	langURL := baseURL + "Puzzle/findAvailableProgrammingLanguages"
	payload := []byte(`[` + fmt.Sprint(puzzle.ID) + `, ` + fmt.Sprint(codinGamer.UserID) + `]`)

	_, body, _, err := requestPost(langURL, payload, puzzle.Client.Cookies)

	availableLanguages := newAvailableLanguages(body)
	return availableLanguages, err
}

func (puzzle Puzzle) GetSolutions(codinGamer CodinGamer, language string) ([]Solution, *ErrorResponse) {
	solutionURL := baseURL + "Solution/findMySolutions"
	gamerId, puzzleId := fmt.Sprint(codinGamer.UserID), fmt.Sprint(puzzle.ID)
	payload := []byte(`[` + gamerId + `, ` + puzzleId + `, "` + language + `"]`)

	_, body, _, err := requestPost(solutionURL, payload, puzzle.Client.Cookies)

	solution := newSolutions(body, puzzle.Client)
	return solution, err
}
