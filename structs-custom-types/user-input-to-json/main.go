package main

import (
	"bufio"
	"example.com/user-input-to-json/model"
	"example.com/user-input-to-json/todo"
	"example.com/user-input-to-json/usernote"
	"fmt"
	"os"
	"strings"
)

func main() {
	// #1 - approach 1, without interface
	//text := getUserData("Todo text: ")
	//
	//newTodo, err := todo.New(text)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//newTodo.Show()
	//
	//err = newTodo.Save()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println("Success to save todo.")
	//
	//title := getUserData("Note title: ")
	//content := getUserData("Note content: ")
	//
	//note, err := usernote.New(title, content)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//note.Show()
	//
	//err = note.Save()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Success to save notes.")

	// #2 - approach 2, with interface saver
	//text := getUserData("Todo text: ")
	//
	//newTodo, err := todo.New(text)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//newTodo.Show()
	//err = saveUserInput(newTodo) // use generic function
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//title := getUserData("Note title: ")
	//content := getUserData("Note content: ")
	//
	//note, err := usernote.New(title, content)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//note.Show()
	//err = saveUserInput(note)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// #3 - approach 3, create more general function with interface that
	// include Show() func
	text := getUserData("Todo text: ")

	newTodo, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputUserInput(newTodo)
	if err != nil {
		fmt.Println(err)
		return
	}

	title := getUserData("Note title: ")
	content := getUserData("Note content: ")

	note, err := usernote.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputUserInput(note)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//func saveUserInput(data model.Saver) error {
//	err := data.Save()
//	if err != nil {
//		return err
//	}
//	return nil
//}

func outputUserInput(data model.Outputable) error {
	data.Show()

	err := data.Save()
	if err != nil {
		return err
	}
	return nil
}

func getUserData(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	userData, _ := reader.ReadString('\n') // Reads until newline
	userData = strings.TrimSpace(userData) // trailing extra newline

	return userData
}
