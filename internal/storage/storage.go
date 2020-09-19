package storage

import(
	"github.com/fatih/color"
	"github.com/Swackles/codingame.go/pkg/codingame"
	"encoding/json"
	"os"
)

type Puzzle struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	LastActivity int64 `json:"lastActivity"`
	Solutions []Solution `json:"solutions"`
}

type Solution struct {
	Language string `json:"language"`
	Code string `json:"code"`
}

func NewSolution(language string, code string) Solution {
	return Solution {
		Language: language,
		Code: code,
	}
}

func NewPuzzle(puzzle codingame.Puzzle, solutions []Solution) Puzzle {
	return Puzzle {
		ID: puzzle.ID,
		Name: puzzle.Title,
		Type: puzzle.Type,
		LastActivity: puzzle.LastActivity,
		Solutions: solutions,
	}
}

func Save(puzzle []Puzzle) {
	file_content, _ := json.MarshalIndent(puzzle, "", " ")

	file, err := os.Create("/metadata.json") 
	if err != nil {
		panic(err)
	}

	defer file.Close()

	written, err := file.Write(file_content)
	if (err != nil) {
		panic(err)
	}

	file.Sync()

	c := color.New(color.FgGreen)
	c.Printf("Wrote %v bytes to metadata.json", written)
}