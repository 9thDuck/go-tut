package main

import (
	"fmt"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}



type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }


func outputData (data outputtable) error {
	data.Display()
	return saveData(data)
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

	err = outputData(note)

	if err != nil {
		return
	}

	todo, err := todo.CreateTodo()

	if err != nil {
		fmt.Println(err)
		return
	}
	
	err = outputData(todo)
	
	if(err != nil) {
		fmt.Println(err)
	}
}
