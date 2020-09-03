package codinGame

import (
	"fmt"
	"net/http"
)

const baseURL = "https://www.codingame.com/services/"

func main() {
}

/*
 *	Gets the recent activity of the user
 *		Inputs:
 *			userID	-	User's ID you want to get recent activity of
 *			count		- How many activities you want to get
 *			cookies	-	HTTP cookies that are needed for the request
 *
 *		Output:
 *			PuzzleResponse	-	Puzzles that came back from the request
 *			[]http.Cookies	-	Cookies that were returned back by the request
 */
func getLastActivity(userID int, count byte, cookies []*http.Cookie) (PuzzleResponse, error) {
	activityURL := baseURL + "LastActivities/getLastActivities"
	payload := []byte(`["` + fmt.Sprint(userID) + `", "` + fmt.Sprint(count) + `"]`)

	_, body, cookies, err := requestPost(activityURL, payload, cookies)
	if err != nil {
		return nil, err
	}

	puzzleResponse := newPuzzleResponse(body)

	return puzzleResponse, nil
}
