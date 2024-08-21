package main

import (
	"fmt"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func main() {
	note, err := note.CreateNote()

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()

	err = saveData(note)

	if err != nil {
		return
	}

	todo, err := todo.CreateTodo()

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()

	saveData(todo)
}
