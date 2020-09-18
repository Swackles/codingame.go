package codingame

import (
	"encoding/json"
	"fmt"
)

type Solution struct {
	Pseudo                          string `json:"pseudo"`
	TestSessionQuestionSubmissionID int    `json:"testSessionQuestionSubmissionId"`
	ProgrammingLanguageID           string `json:"programmingLanguageId"`
	CreationTime                    int64  `json:"creationTime"`
	Avatar                          int64  `json:"avatar"`
	CodingamerID                    int    `json:"codingamerId"`
	CommentableID                   int    `json:"commentableId"`
	CommentCount                    int    `json:"commentCount"`
	UpVotes                         int    `json:"upVotes"`
	DownVotes                       int    `json:"downVotes"`
	Shared                          bool   `json:"shared"`
	CodingamerHandle                string `json:"codingamerHandle"`
	Code                            string `json:"code"`
	VotableID                       int    `json:"votableId"`
	Client                          CodinGameClient
}

func (solution *Solution) GetCode() ([]byte, *ErrorResponse) {
	if solution.Code != "" {
		return []byte(solution.Code), nil
	}

	newSolution, err := solution.getSolution()

	return []byte(newSolution.Code), err
}

func (solution *Solution) getSolution() (*Solution, *ErrorResponse) {
	solutionURL := baseURL + "Solution/findSolution"
	CodinGamerId, SumbmissionID := fmt.Sprint(solution.CodingamerID), fmt.Sprint(solution.TestSessionQuestionSubmissionID)
	payload := []byte(`[` + CodinGamerId + `, ` + SumbmissionID + `]`)

	_, body, err := requestPost(solutionURL, payload, )

	return newSolution(body), err
}

func newSolutions(body []byte) []Solution {
	var response []Solution
	json.Unmarshal(body, &response)

	return response
}

func newSolution(body []byte) *Solution {
	var response *Solution
	json.Unmarshal(body, &response)

	return response
}
