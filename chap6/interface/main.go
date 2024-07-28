package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"chap6/interface/note"
	"chap6/interface/todo"
)

//  save interface needs a valid value with a save method with no arguments and returns an error
type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// type outputtable interface {
// 	Save() error
// 	Display()
// }

// embedded interface
type outputtable interface {
	saver
	Display()
}


func main() {

	printSomething("hello world")
	printSomething(100)
	// printSomething(outputtable)
	
	todoText := getUserInput("Todo text: ")
	title, content := getNoteData()
	
	todo, err := todo.New(todoText)
	printSomething(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}


	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

	// todo.Display()
	// // err = todo.Save()
	// err = saveData(&todo)

	err = outputData(&todo)

	if err != nil {
		return
	}


	fmt.Println("// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// userNote.Display()
	// // err = userNote.Save()
	// err = saveData(&userNote)

	err = outputData(&userNote)

	if err != nil {
		return
	}


}

// one saveData function to save two different types of values
// func saveData(data note.Note) {

// }

//any = interface{}
func printSomething(value any) {

	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String:", stringVal)
		return
	}

	todoVal, ok := value.(todo.Todo)

	if ok {
		fmt.Println("Todo:", todoVal)
		return
	}


	// switch valueType:=value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:",value)
	// case todo.Todo:
	// 	fmt.Println("Todo:",value,valueType)
	// default:
	// 	fmt.Println("Other Value Type:",valueType)
	// }
}

func outputData(data outputtable) error{
	data.Display()
	return saveData(data) //reference
}



func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving failed.")
		return err
	}

	fmt.Println("Saving succeeded!")
	return nil
}


func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}