package main

import (
	"fmt"

	"example.com/notes/note"
)

func main() {
	note, err := note.CreateNote()

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()

	err = note.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}

	fmt.Println("Saving the note succeeded!")
}
