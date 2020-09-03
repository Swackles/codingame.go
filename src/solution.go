package codinGame

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

func (solution *Solution) GetCode() ([]byte, error) {
	if solution.Code != "" {
		return []byte(solution.Code), nil
	}

	newSolution, err := solution.getSolution()
	if err != nil {
		return nil, err
	}

	return []byte(newSolution.Code), nil
}

func (solution *Solution) getSolution() (*Solution, error) {
	solutionURL := baseURL + "Solution/findSolution"
	CodinGamerId, SumbmissionID := fmt.Sprint(solution.CodingamerID), fmt.Sprint(solution.TestSessionQuestionSubmissionID)
	payload := []byte(`[` + CodinGamerId + `, ` + SumbmissionID + `]`)

	_, body, _, err := requestPost(solutionURL, payload, solution.Client.Cookies)
	if err != nil {
		return nil, err
	}

	return newSolution(body, solution.Client), nil
}

func newSolutions(body []byte, client CodinGameClient) []Solution {
	var response []Solution
	json.Unmarshal(body, &response)

	response = append(response, Solution{Client: client})

	return response
}

func newSolution(body []byte, client CodinGameClient) *Solution {
	var response *Solution
	json.Unmarshal(body, &response)

	response.Client = client

	return response
}
