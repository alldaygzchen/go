package main

import (
	"bufio"
	"chap5/practice/note"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, content:= getNoteData()
	userNote,err:=note.New(title,content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")

	fmt.Println("###########################")
	
	fmt.Println("Read your json file:")
	var filename string
	fmt.Scanln(&filename)
	json_content, err := userNote.Read(filename)
	if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Printf("Memory address of data %p\n", json_content)
	fmt.Println(json_content)
	fmt.Println(json_content.Title)
	fmt.Println(json_content.Content)
	fmt.Println(*json_content)



}

func getNoteData() (string, string) {
	title:= getUserInput("Note title:")

	content:= getUserInput("Note content:")

	// fmt.Println("##title",title,"##content",content)

	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Printf("%v ",prompt)
	// constructor function
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text


	// return value
}