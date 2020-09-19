package main

import(
	"fmt"
	"github.com/Swackles/codingame.go/pkg/codingame"
	"github.com/Swackles/codingame.go/internal/storage"
)

func main() {
	puzzle := storage.Puzzle {
		ID: 1,
		Name: "Test",
	}

	storage.Save([]storage.Puzzle{puzzle})

	return
	loginResponse, err := codingame.New("username", "password")
	check_err(err)

	puzzleResponses, err := loginResponse.CodinGamer.GetLastActivity(4)
	check_err(err)

	for _, puzzleResponse := range puzzleResponses {
		puzzle := puzzleResponse.Puzzle
		if puzzleResponse.Type == "CLASH" || puzzle.Type == "ARENA" {
			continue
		}
		fmt.Println(puzzle.Title)
		languages, err := puzzle.GetProgrammingLanguages()
		check_err(err)

		for _, language := range languages {
			if language.Solved {
				fmt.Printf("Language: %v\n", language.ID)
				solutions, err := puzzle.GetSolutions(language.ID)
				check_err(err)
				for _, solution := range solutions {
					code, err := solution.GetCode()
					check_err(err)

					fmt.Println(string(code))
				}
			}
		}

		fmt.Println("\n=================\n")
	}
}

func check_err(err *codingame.ErrorResponse) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}