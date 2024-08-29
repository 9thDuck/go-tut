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

func printSomethingSwitch (value interface {}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println(value)
	}
}

func printIntStr (value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("int value", intVal)
	} 

	strVal, ok := value.(string)

	if ok {
		fmt.Println("str value", strVal)
	}

}

func main() {

	//Note 
	note, err := note.CreateNote()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(note)

	if err != nil {
		return
	}

	//Todo
	todo, err := todo.CreateTodo()

	if err != nil {
		fmt.Println(err)
		return
	}
	
	err = outputData(todo)
	
	if(err != nil) {
		fmt.Println(err)
	}

	// type checking using switch
	printSomethingSwitch(1)
	printSomethingSwitch('s') //ignored
	printSomethingSwitch("asdf")
	printSomethingSwitch(`asdf`)

	printIntStr(2)
	printIntStr("s")
	printIntStr('s') //ignored
}
