package codingame

import (
	"fmt"
)

type CodinGamer struct {
	UserID                 int        `json:"userId"`
	Email                  string     `json:"email"`
	Pseudo                 string     `json:"pseudo"`
	CountryID              string     `json:"countryId"`
	PublicHandle           string     `json:"publicHandle"`
	PrivateHandle          string     `json:"privateHandle"`
	FormValues             FormValues `json:"formValues"`
	RegisterOrigin         string     `json:"registerOrigin"`
	Enable                 bool       `json:"enable"`
	SchoolID               int        `json:"schoolId"`
	Rank                   int        `json:"rank"`
	FormCachedValues       FormValues `json:"formCachedValues"`
	Avatar                 int64      `json:"avatar"`
	CreationTime           int64      `json:"creationTime"`
	OnlineSince            int64      `json:"onlineSince"`
	HiddenFromFriendFinder bool       `json:"hiddenFromFriendFinder"`
	ShareAllSolutions      bool       `json:"shareAllSolutions"`
	Tagline                string     `json:"tagline"`
	Company                string     `json:"company"`
	City                   string     `json:"city"`
	Level                  int        `json:"level"`
	Xp                     int        `json:"xp"`
	Category               string     `json:"category"`
	PrivateUser            bool       `json:"privateUser"`
	CreatedCoursesCount    int        `json:"createdCoursesCount"`
	AlreadyAnsweredOptin   bool       `json:"alreadyAnsweredOptin"`
}

type FormValues struct {
	Skills    string `json:"skills"`
	Country   string `json:"country"`
	City      string `json:"city"`
	School    string `json:"school"`
	Student   string `json:"student"`
	SchoolID  string `json:"schoolId"`
	Company   string `json:"company"`
	Avatar    string `json:"avatar"`
	Category  string `json:"category"`
	CountryID string `json:"countryId"`
	Pseudo    string `json:"pseudo"`
}

/*
 *	Gets codingamers last activity
 *		Inputs:
 *			count	-	How many activities you want to query
 *			cookies	-	HTTP cookies that are needed for the request
 */
func (codinGamer CodinGamer) GetLastActivity(count byte) (PuzzleResponse, *ErrorResponse) {
	activityURL := baseURL + "LastActivities/getLastActivities"
	payload := []byte(`["` + fmt.Sprint(codinGamer.UserID) + `", "` + fmt.Sprint(count) + `"]`)

	_, body, err := requestPost(activityURL, payload)
	if err != nil {
		return nil, err
	}

	puzzleResponse := newPuzzleResponse(body)

	return puzzleResponse, nil
}
