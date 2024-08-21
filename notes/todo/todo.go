package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	userInput "example.com/notes/user_input"
)

type Todo struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func New(text string) (*Todo, error) {

	if text == "" {
		return &Todo{}, errors.New("todo text cannot be empty")
	}
	CreatedAt := time.Now()
	return &Todo{text, CreatedAt}, nil
}

func (t Todo) Display() {
	fmt.Printf("Todo: %v\n", t.Text)
}

func CreateTodo() (defaultTodo *Todo, err error) {
	userInput, err := userInput.GetUserInput("Enter todo text:")
	defaultTodo = &Todo{}

	if err != nil {
		return defaultTodo, nil
	}

	newTodo, err := New(userInput)

	if err != nil {
		return defaultTodo, err
	}

	return newTodo, nil
}

func (t Todo) Save() error {
	dateTime, _ := time.Parse(time.DateTime, t.CreatedAt.Format(time.DateTime))
	fileName := fmt.Sprintf("todo_%v.json", dateTime)
	jsonData, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}
