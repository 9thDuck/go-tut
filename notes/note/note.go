package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	userInput "example.com/notes/user_input"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("title and content are required ")
	}

	createdAt := time.Now()

	return &Note{
		title, content, createdAt,
	}, nil
}

func CreateNote() (newNote *Note, err error) {
	title, err := userInput.GetUserInput("Note title:")

	if err != nil {
		return nil, err
	}

	content, err := userInput.GetUserInput("Note content:")

	if err != nil {
		return nil, err
	}

	newNote, err = New(title, content)

	if err != nil {
		return nil, err
	}
	return newNote, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled \"%v\" has the folloing content:\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonData, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)

}
