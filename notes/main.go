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

}
